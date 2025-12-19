# Quick Start Guide - User Seeding Script

## Quick Commands

### Insert 1,000 users:
```powershell
cd c:\Users\ak777\Desktop\GoLang\01_First
go run scripts/seed_users.go -count=1000
```

### Insert 2,000 users:
```powershell
cd c:\Users\ak777\Desktop\GoLang\01_First
go run scripts/seed_users.go -count=2000
```

## What You Need

1. **MongoDB running** (default: `mongodb://localhost:27017`)
2. **Database name** (default: `userdb`)
3. **Go installed** and in your PATH

## What Gets Created

- **1,000 or 2,000 user records** (based on your choice)
- **All passwords**: `123456` (hashed with bcrypt)
- **Realistic dummy data**: Random names and unique emails
- **Mixed roles**: 90% users, 10% admins
- **Mixed status**: 95% active, 5% inactive

## Example Emails Generated

- `john.smith1234@gmail.com`
- `jane.doe5678@yahoo.com`
- `michael.johnson9012@hotmail.com`

## Verify Insertion

After running the script, you can verify users were inserted by:
1. Using MongoDB Compass or mongo shell
2. Checking your application's user list endpoint
3. The script will show a summary of inserted users

## Troubleshooting

- **Connection error**: Make sure MongoDB is running
- **Duplicate emails**: Script will skip duplicates and continue
- **Permission error**: Check MongoDB user permissions
