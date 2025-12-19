# User Seeding Script

This script allows you to bulk insert user records into your MongoDB database with dummy data.

## Features

- Insert 1,000 or 2,000 users at a time
- All users have password: `123456`
- Generates realistic dummy data (names, emails)
- Uses bulk insert for optimal performance
- Handles duplicate email errors gracefully

## Prerequisites

1. MongoDB must be running and accessible
2. Environment variables must be set (or use defaults):
   - `MONGO_URI` (default: `mongodb://localhost:27017`)
   - `MONGO_DB` (default: `userdb`)

## How to Run

### Option 1: Insert 1,000 users

```bash
cd c:\Users\ak777\Desktop\GoLang\01_First
go run scripts/seed_users.go -count=1000
```

### Option 2: Insert 2,000 users

```bash
cd c:\Users\ak777\Desktop\GoLang\01_First
go run scripts/seed_users.go -count=2000
```

### Option 3: Using PowerShell (Windows)

```powershell
cd c:\Users\ak777\Desktop\GoLang\01_First
go run scripts/seed_users.go -count=1000
```

or

```powershell
cd c:\Users\ak777\Desktop\GoLang\01_First
go run scripts/seed_users.go -count=2000
```

## Environment Variables

The script uses the same configuration as your main application. You can set these in a `.env` file or as environment variables:

```env
MONGO_URI=mongodb://localhost:27017
MONGO_DB=userdb
```

If not set, the script will use the default values from `config/config.go`.

## Generated User Data

- **Names**: Random combinations of first and last names
- **Emails**: Format: `firstname.lastname[number]@domain.com`
- **Password**: All users have password `123456` (hashed with bcrypt)
- **Role**: 90% users, 10% admins
- **IsActive**: 95% active, 5% inactive
- **Timestamps**: Random creation dates within the last year

## Example Output

```
Connecting to MongoDB...
✅ MongoDB connected successfully
Generating 1000 users...
Batch 1: Successfully inserted 500 users
Batch 2: Successfully inserted 500 users

✅ Successfully inserted 1000 out of 1000 users
All users have password: 123456
```

## Notes

- The script inserts users in batches of 500 for optimal performance
- If duplicate emails are encountered, the script will continue inserting other users
- All passwords are hashed using bcrypt (same as your application)
- The script respects the unique email index in your database

## Troubleshooting

1. **Connection Error**: Make sure MongoDB is running and accessible
2. **Duplicate Emails**: If you run the script multiple times, some emails may already exist. The script will skip duplicates and continue.
3. **Permission Error**: Ensure your MongoDB user has write permissions
