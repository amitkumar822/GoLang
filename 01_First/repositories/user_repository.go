package repositories

import (
	"context"
	"errors"
	"strings"
	"time"

	"user-management-system/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserRepository handles all database operations for users
type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository creates a new user repository
func NewUserRepository(collection *mongo.Collection) *UserRepository {
	repo := &UserRepository{collection: collection}
	repo.createIndexes()
	return repo
}

// createIndexes creates necessary indexes for the user collection
func (r *UserRepository) createIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create unique index on email
	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := r.collection.Indexes().CreateOne(ctx, emailIndex)
	if err != nil {
		// Index might already exist, which is fine
		_ = err
	}
}

// Create inserts a new user into the database
func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, user)
	return err
}

// FindByID finds a user by their ID
func (r *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	var user models.User
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// FindByEmail finds a user by their email
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	// Convert email to lowercase for consistent searching
	email = strings.ToLower(strings.TrimSpace(email))

	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// Update updates an existing user
func (r *UserRepository) Update(ctx context.Context, id string, updateData bson.M) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID")
	}

	updateData["updatedAt"] = time.Now()

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updateData}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

// Delete removes a user from the database
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID")
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

// FindAll retrieves all users with pagination
func (r *UserRepository) FindAll(ctx context.Context, page, limit int) ([]*models.User, int64, error) {
	// Calculate skip value
	skip := (page - 1) * limit

	// Set up pagination options
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}}) // Sort by createdAt descending

	// Find all users
	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var users []*models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, 0, err
	}

	// Get total count
	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// FindAllWithoutLimit retrieves ALL users from database without pagination limit
// Uses concurrent MongoDB cursor processing with channels for optimal performance
func (r *UserRepository) FindAllWithoutLimit(ctx context.Context) ([]*models.User, int64, error) {
	// Set up find options with optimized batch size
	findOptions := options.Find()
	findOptions.SetBatchSize(2000)                             // Increased batch size for better throughput
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}}) // Sort by createdAt descending

	// Find all users with cursor
	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	// Get total count concurrently in separate goroutine
	type countResult struct {
		count int64
		err   error
	}
	countChan := make(chan countResult, 1)

	go func() {
		total, err := r.collection.CountDocuments(ctx, bson.M{})
		countChan <- countResult{count: total, err: err}
	}()

	// Use channels for concurrent cursor processing
	userChan := make(chan *models.User, 2000) // Buffered channel for better throughput
	errChan := make(chan error, 1)

	// Start goroutine to process cursor concurrently
	go func() {
		defer close(userChan)
		for cursor.Next(ctx) {
			var user models.User
			if err := cursor.Decode(&user); err != nil {
				select {
				case errChan <- err:
				case <-ctx.Done():
					return
				}
				return
			}
			select {
			case userChan <- &user:
			case <-ctx.Done():
				return
			}
		}
		// Check for cursor errors
		if err := cursor.Err(); err != nil {
			select {
			case errChan <- err:
			case <-ctx.Done():
			}
		}
	}()

	// Collect users from channel concurrently
	var users []*models.User
	users = make([]*models.User, 0, 10000) // Pre-allocate with reasonable capacity

	// Use a separate goroutine to collect users while cursor is processing
	done := make(chan struct{})
	go func() {
		defer close(done)
		for user := range userChan {
			users = append(users, user)
		}
	}()

	// Wait for collection to complete or error
	select {
	case err := <-errChan:
		return nil, 0, err
	case <-done:
		// Collection completed successfully
	case <-ctx.Done():
		return nil, 0, ctx.Err()
	}

	// Get count result
	countRes := <-countChan
	if countRes.err != nil {
		return nil, 0, countRes.err
	}

	return users, countRes.count, nil
}

// GetTotalCount retrieves the total count of users in the database
func (r *UserRepository) GetTotalCount(ctx context.Context) (int64, error) {
	return r.collection.CountDocuments(ctx, bson.M{})
}
