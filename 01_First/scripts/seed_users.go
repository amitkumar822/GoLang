package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"user-management-system/config"
	"user-management-system/database"
	"user-management-system/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var (
	firstNames = []string{
		"John", "Jane", "Michael", "Sarah", "David", "Emily", "James", "Jessica",
		"Robert", "Ashley", "William", "Amanda", "Richard", "Melissa", "Joseph", "Deborah",
		"Thomas", "Michelle", "Charles", "Carol", "Christopher", "Lisa", "Daniel", "Nancy",
		"Matthew", "Karen", "Anthony", "Betty", "Mark", "Helen", "Donald", "Sandra",
		"Steven", "Donna", "Paul", "Carolyn", "Andrew", "Ruth", "Joshua", "Sharon",
		"Kenneth", "Laura", "Kevin", "Angela", "Brian", "Brenda", "George", "Emma",
		"Timothy", "Olivia", "Ronald", "Cynthia", "Jason", "Marie", "Edward", "Janet",
		"Jeffrey", "Catherine", "Ryan", "Frances", "Jacob", "Christine", "Gary", "Samantha",
		"Nicholas", "Debra", "Eric", "Rachel", "Jonathan", "Carolyn", "Stephen", "Janet",
		"Larry", "Virginia", "Justin", "Maria", "Scott", "Heather", "Brandon", "Diane",
		"Benjamin", "Julie", "Samuel", "Joyce", "Frank", "Victoria", "Gregory", "Kelly",
		"Raymond", "Christina", "Alexander", "Joan", "Patrick", "Evelyn", "Jack", "Judith",
		"Dennis", "Megan", "Jerry", "Cheryl", "Tyler", "Andrea", "Aaron", "Hannah",
		"Jose", "Jacqueline", "Adam", "Martha", "Nathan", "Gloria", "Henry", "Teresa",
		"Douglas", "Sara", "Zachary", "Janice", "Peter", "Marie", "Kyle", "Julia",
		"Noah", "Grace", "Ethan", "Judy", "Jeremy", "Theresa", "Walter", "Madison",
		"Christian", "Beverly", "Keith", "Denise", "Roger", "Marilyn", "Terry", "Amber",
		"Austin", "Danielle", "Sean", "Rose", "Gerald", "Brittany", "Carl", "Diana",
		"Harold", "Abigail", "Dylan", "Jane", "Wayne", "Lori", "Ralph", "Mason",
	}

	lastNames = []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis",
		"Rodriguez", "Martinez", "Hernandez", "Lopez", "Wilson", "Anderson", "Thomas", "Taylor",
		"Moore", "Jackson", "Martin", "Lee", "Thompson", "White", "Harris", "Sanchez",
		"Clark", "Ramirez", "Lewis", "Robinson", "Walker", "Young", "Allen", "King",
		"Wright", "Scott", "Torres", "Nguyen", "Hill", "Flores", "Green", "Adams",
		"Nelson", "Baker", "Hall", "Rivera", "Campbell", "Mitchell", "Carter", "Roberts",
		"Gomez", "Phillips", "Evans", "Turner", "Diaz", "Parker", "Cruz", "Edwards",
		"Collins", "Reyes", "Stewart", "Morris", "Morales", "Murphy", "Cook", "Rogers",
		"Gutierrez", "Ortiz", "Morgan", "Cooper", "Peterson", "Bailey", "Reed", "Kelly",
		"Howard", "Ramos", "Kim", "Cox", "Ward", "Richardson", "Watson", "Brooks",
		"Chavez", "Wood", "James", "Bennett", "Gray", "Mendoza", "Ruiz", "Hughes",
		"Price", "Alvarez", "Castillo", "Sanders", "Patel", "Myers", "Long", "Ross",
		"Foster", "Jimenez", "Powell", "Jenkins", "Perry", "Russell", "Sullivan", "Bell",
		"Coleman", "Butler", "Henderson", "Barnes", "Gonzales", "Fisher", "Vasquez", "Simmons",
		"Romero", "Jordan", "Patterson", "Alexander", "Hamilton", "Graham", "Reynolds", "Griffin",
		"Wallace", "Moreno", "West", "Cole", "Hayes", "Bryant", "Herrera", "Gibson",
		"Ellis", "Tran", "Medina", "Aguilar", "Stevens", "Murray", "Ford", "Castro",
		"Marshall", "Owens", "Harrison", "Fernandez", "Mcdonald", "Woods", "Washington", "Kennedy",
		"Wells", "Vargas", "Henry", "Chen", "Freeman", "Webb", "Tucker", "Guzman",
		"Burns", "Crawford", "Olson", "Simpson", "Porter", "Hunter", "Gordon", "Mendez",
		"Silva", "Shaw", "Snyder", "Mason", "Dixon", "Munoz", "Hunt", "Hicks",
		"Holmes", "Palmer", "Wagner", "Black", "Robertson", "Boyd", "Rose", "Stone",
		"Salazar", "Fox", "Warren", "Mills", "Meyer", "Rice", "Schmidt", "Garza",
		"Daniels", "Ferguson", "Nichols", "Stephens", "Soto", "Weaver", "Ryan", "Gardner",
		"Payne", "Grant", "Dunn", "Kelley", "Spencer", "Hawkins", "Arnold", "Pierce",
		"Vazquez", "Hansen", "Peters", "Santos", "Hart", "Bradley", "Knight", "Elliott",
		"Cunningham", "Duncan", "Armstrong", "Hudson", "Carroll", "Lane", "Riley", "Andrews",
		"Alvarado", "Ray", "Delgado", "Berry", "Perkins", "Hoffman", "Johnston", "Matthews",
		"Pena", "Richards", "Contreras", "Willis", "Carpenter", "Lawrence", "Sandoval", "Guerrero",
		"George", "Chapman", "Rios", "Estrada", "Ortega", "Watkins", "Greene", "Nunez",
		"Wheeler", "Valdez", "Harper", "Lynch", "Barker", "Maldonado", "Bauer", "Larson",
		"Mack", "Mcdaniel", "Carr", "Townsend", "Calderon", "Ochoa", "Robbins", "Lucas",
		"Floyd", "Bishop", "Bradford", "Curry", "Brewer", "Torres", "Strickland", "Mann",
		"Schneider", "Erickson", "Small", "Lynch", "Marshall", "Walsh", "Jensen", "Horton",
		"Gilbert", "Garrett", "Romero", "Lawrence", "Lawson", "Fields", "Gutierrez", "Ryan",
		"Schmidt", "Carr", "Vasquez", "Castillo", "Wheeler", "Chapman", "Oliver", "Love",
		"Dean", "Snyder", "Morrison", "Kim", "Dunn", "Bradley", "Knight", "Porter",
		"Hunter", "Romero", "Hicks", "Crawford", "Henry", "Boyd", "Mason", "Morales",
		"Kennedy", "Warren", "Dixon", "Ramos", "Reyes", "Burns", "Gordon", "Shaw",
		"Holmes", "Rice", "Robertson", "Hunt", "Black", "Daniels", "Palmer", "Mills",
		"Nichols", "Grant", "Knight", "Ferguson", "Rose", "Stone", "Hawkins", "Dunn",
		"Perkins", "Hudson", "Spencer", "Gardner", "Stephens", "Payne", "Pierce", "Berry",
		"Matthews", "Arnold", "Wagner", "Willis", "Ray", "Watkins", "Olson", "Carroll",
		"Duncan", "Snyder", "Hart", "Cunningham", "Bradley", "Lane", "Andrews", "Ruiz",
		"Harper", "Fox", "Riley", "Armstrong", "Carpenter", "Weaver", "Greene", "Lawrence",
		"Elliott", "Chavez", "Sims", "Austin", "Peters", "Kelley", "Franklin", "Lawson",
		"Fields", "Gutierrez", "Ryan", "Schmidt", "Carr", "Vasquez", "Castillo", "Wheeler",
		"Chapman", "Oliver", "Love", "Dean", "Snyder", "Morrison", "Kim", "Dunn",
	}

	domains = []string{
		"gmail.com", "yahoo.com", "hotmail.com", "outlook.com", "aol.com",
		"icloud.com", "protonmail.com", "mail.com", "zoho.com", "yandex.com",
	}
)

func main() {
	var count int
	flag.IntVar(&count, "count", 1000, "Number of users to insert (1000, 2000, 10000, or 50000 or 100000)")
	flag.Parse()

	if count != 1000 && count != 2000 && count != 10000 && count != 50000 && count != 100000 {
		log.Fatal("Count must be either 1000, 2000, 10000, 50000 or 100000")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to MongoDB
	log.Println("Connecting to MongoDB...")
	if err := database.Connect(cfg.MongoURI, cfg.MongoDB); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer database.Disconnect()

	// Get users collection
	collection := database.GetCollection("users")

	// Hash password once (all users will have password "123456")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	passwordHash := string(hashedPassword)

	// Generate users
	log.Printf("Generating %d users...", count)
	users := generateUsers(count, passwordHash)

	// Insert users in batches for better performance
	batchSize := 500
	totalInserted := 0
	ctx := context.Background()

	for i := 0; i < len(users); i += batchSize {
		end := i + batchSize
		if end > len(users) {
			end = len(users)
		}

		batch := users[i:end]
		documents := make([]interface{}, len(batch))
		for j, user := range batch {
			documents[j] = user
		}

		// Use InsertMany with ordered=false to continue on duplicate key errors
		opts := options.InsertMany().SetOrdered(false)
		result, err := collection.InsertMany(ctx, documents, opts)

		if err != nil {
			// Count successful inserts even if there are errors
			if result != nil {
				inserted := len(result.InsertedIDs)
				totalInserted += inserted
				log.Printf("Batch %d: Inserted %d users, some errors occurred (likely duplicates)",
					(i/batchSize)+1, inserted)
			} else {
				log.Printf("Batch %d error: %v", (i/batchSize)+1, err)
			}
		} else {
			inserted := len(result.InsertedIDs)
			totalInserted += inserted
			log.Printf("Batch %d: Successfully inserted %d users", (i/batchSize)+1, inserted)
		}
	}

	log.Printf("\n✅ Successfully inserted %d out of %d users", totalInserted, count)
	log.Println("All users have password: 123456")
}

func generateUsers(count int, passwordHash string) []models.User {
	rand.Seed(time.Now().UnixNano())
	users := make([]models.User, count)
	emailSet := make(map[string]bool) // Track emails to avoid duplicates

	for i := 0; i < count; i++ {
		// Generate unique email
		var email string
		for {
			firstName := firstNames[rand.Intn(len(firstNames))]
			lastName := lastNames[rand.Intn(len(lastNames))]
			domain := domains[rand.Intn(len(domains))]
			number := rand.Intn(100000) // Increased range for 10k users
			email = fmt.Sprintf("%s.%s%d@%s",
				firstName, lastName, number, domain)
			email = strings.ToLower(email)

			if !emailSet[email] {
				emailSet[email] = true
				break
			}
		}

		// Generate name
		firstName := firstNames[rand.Intn(len(firstNames))]
		lastName := lastNames[rand.Intn(len(lastNames))]
		name := fmt.Sprintf("%s %s", firstName, lastName)

		// Random role (mostly "user", some "admin")
		role := "user"
		if rand.Float32() < 0.1 { // 10% chance of admin
			role = "admin"
		}

		// Random active status (mostly active)
		isActive := true
		if rand.Float32() < 0.05 { // 5% chance of inactive
			isActive = false
		}

		now := time.Now()
		// Add some variation to timestamps
		createdAt := now.Add(-time.Duration(rand.Intn(365)) * 24 * time.Hour)

		users[i] = models.User{
			ID:        primitive.NewObjectID(),
			Name:      name,
			Email:     email,
			Password:  passwordHash,
			Role:      role,
			IsActive:  isActive,
			CreatedAt: createdAt,
			UpdatedAt: createdAt,
		}
	}

	return users
}
