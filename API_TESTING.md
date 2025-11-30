# API Testing Guide

## Base URL
```
http://localhost:8080
```

## 1. Authentication Endpoints

### 1.1 Register User
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

Response:
```json
{
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "username": "testuser",
  "role": "user",
  "access_token": "xxxxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  "token_expires_at": "2024-02-01T12:00:00Z"
}
```

### 1.2 Login
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }'
```

Response:
```json
{
  "user_id": "admin-uuid",
  "access_token": "xxxxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  "token_expires_at": "2024-02-01T12:00:00Z"
}
```

### 1.3 Get Profile
```bash
curl http://localhost:8080/api/auth/profile \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

Response:
```json
{
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "username": "testuser",
  "role": "user"
}
```

## 2. Game Endpoints (User Auth Required)

### 2.1 Get Active Stages
```bash
curl http://localhost:8080/api/stages \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

Response:
```json
[
  {
    "stage_id": "stage-001",
    "name": "Java Basics",
    "difficulty": "easy"
  },
  {
    "stage_id": "stage-002",
    "name": "Python Functions",
    "difficulty": "medium"
  }
]
```

### 2.2 Get Stage Detail with Phrases
```bash
curl http://localhost:8080/api/stage/stage-001 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

Response:
```json
{
  "stage_id": "stage-001",
  "name": "Java Basics",
  "theme": "Programming",
  "difficulty": "easy",
  "phrases": [
    {
      "phrase_id": "phrase-001",
      "text": "public class HelloWorld",
      "sequence_number": 1,
      "multiplier": 1.0
    },
    {
      "phrase_id": "phrase-002",
      "text": "System.out.println(\"Hello\");",
      "sequence_number": 2,
      "multiplier": 1.2
    }
  ]
}
```

### 2.3 Submit Score
```bash
curl -X POST http://localhost:8080/api/score/submit \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "stage_id": "stage-001",
    "total_time_ms": 15000,
    "total_errors": 2
  }'
```

Response:
```json
{
  "status": "UPSERTED",
  "final_score": 156.50
}
```

Status values:
- `UPSERTED`: Score berhasil disimpan atau diupdate (score lebih baik)
- `IGNORED`: Score tidak diupdate (score tidak lebih baik dari yang sudah ada)

### 2.4 Get Leaderboard
```bash
curl "http://localhost:8080/api/leaderboard?stage_id=stage-001&limit=10" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

Response:
```json
[
  {
    "username": "user1",
    "final_score": 250.75,
    "total_time_ms": 12000
  },
  {
    "username": "user2",
    "final_score": 180.50,
    "total_time_ms": 15000
  }
]
```

## 3. Admin Endpoints (Admin Auth Required)

### 3.1 Create Stage
```bash
curl -X POST http://localhost:8080/admin/stage \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "JavaScript Basics",
    "theme": "Programming",
    "difficulty": "easy",
    "is_active": true
  }'
```

Response:
```json
{
  "stage_id": "generated-uuid",
  "name": "JavaScript Basics",
  "theme": "Programming",
  "difficulty": "easy",
  "is_active": true
}
```

### 3.2 Update Stage
```bash
curl -X PUT http://localhost:8080/admin/stage/stage-001 \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Java Basics Updated",
    "theme": "Programming",
    "difficulty": "medium",
    "is_active": true
  }'
```

### 3.3 Delete Stage
```bash
curl -X DELETE http://localhost:8080/admin/stage/stage-001 \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN"
```

Response:
```json
{
  "message": "stage deleted successfully"
}
```

### 3.4 Get All Stages (including inactive)
```bash
curl http://localhost:8080/admin/stages \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN"
```

### 3.5 Create Phrase
```bash
curl -X POST http://localhost:8080/admin/phrase \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "stage_id": "stage-001",
    "text": "console.log(\"Hello World\");",
    "sequence_number": 1,
    "base_multiplier": 1.5
  }'
```

Response:
```json
{
  "phrase_id": "generated-uuid",
  "stage_id": "stage-001",
  "text": "console.log(\"Hello World\");",
  "sequence_number": 1,
  "multiplier": 1.5
}
```

### 3.6 Update Phrase
```bash
curl -X PUT http://localhost:8080/admin/phrase/phrase-001 \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "stage_id": "stage-001",
    "text": "console.log(\"Updated\");",
    "sequence_number": 1,
    "base_multiplier": 2.0
  }'
```

### 3.7 Delete Phrase
```bash
curl -X DELETE http://localhost:8080/admin/phrase/phrase-001 \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN"
```

### 3.8 Get Phrases by Stage
```bash
curl "http://localhost:8080/admin/phrases?stage_id=stage-001" \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN"
```

## 4. Health Check
```bash
curl http://localhost:8080/health
```

Response:
```json
{
  "status": "ok"
}
```

## Score Calculation

Formula:
```
Final Score = (Σ(phrase_length × multiplier) / time_in_seconds) - (errors × 50)
```

Example:
- Phrase 1: "Hello World" (11 chars) × 1.0 = 11
- Phrase 2: "Testing" (7 chars) × 1.5 = 10.5
- Total weighted chars = 21.5
- Time taken = 15000ms = 15 seconds
- Errors = 2
- Score = (21.5 / 15) - (2 × 50) = 1.43 - 100 = -98.57 (clamped to 0)

Better example:
- Phrase 1: "public class HelloWorld" (23 chars) × 1.0 = 23
- Phrase 2: "System.out.println(\"Hello\");" (28 chars) × 1.2 = 33.6
- Phrase 3: "int number = 42;" (16 chars) × 1.0 = 16
- Total weighted chars = 72.6
- Time taken = 15000ms = 15 seconds
- Errors = 2
- Score = (72.6 / 15) - (2 × 50) = 4.84 - 100 = -95.16 (clamped to 0)

For positive score:
- Time taken = 10000ms = 10 seconds
- Errors = 1
- Score = (72.6 / 10) - (1 × 50) = 7.26 - 50 = -42.74 (still negative)

Very fast typing:
- Time taken = 5000ms = 5 seconds
- Errors = 0
- Score = (72.6 / 5) - (0 × 50) = 14.52 - 0 = 14.52 ✓

## Error Responses

### 400 Bad Request
```json
{
  "error": "validation error message"
}
```

### 401 Unauthorized
```json
{
  "error": "unauthorized"
}
```

### 403 Forbidden
```json
{
  "error": "admin access required"
}
```

### 404 Not Found
```json
{
  "error": "stage not found"
}
```

### 409 Conflict
```json
{
  "error": "username already exists"
}
```

### 500 Internal Server Error
```json
{
  "error": "internal server error"
}
```

## Testing Flow

1. **Login sebagai admin**
   ```bash
   TOKEN=$(curl -s -X POST http://localhost:8080/api/auth/login \
     -H "Content-Type: application/json" \
     -d '{"username":"admin","password":"admin123"}' | jq -r '.access_token')
   ```

2. **Create stage**
   ```bash
   curl -X POST http://localhost:8080/admin/stage \
     -H "Authorization: Bearer $TOKEN" \
     -H "Content-Type: application/json" \
     -d '{"name":"Test Stage","theme":"Test","difficulty":"easy","is_active":true}'
   ```

3. **Register user**
   ```bash
   USER_TOKEN=$(curl -s -X POST http://localhost:8080/api/auth/register \
     -H "Content-Type: application/json" \
     -d '{"username":"testuser","password":"test123"}' | jq -r '.access_token')
   ```

4. **Get stages as user**
   ```bash
   curl http://localhost:8080/api/stages \
     -H "Authorization: Bearer $USER_TOKEN"
   ```

5. **Submit score**
   ```bash
   curl -X POST http://localhost:8080/api/score/submit \
     -H "Authorization: Bearer $USER_TOKEN" \
     -H "Content-Type: application/json" \
     -d '{"stage_id":"stage-001","total_time_ms":10000,"total_errors":0}'
   ```

6. **Check leaderboard**
   ```bash
   curl "http://localhost:8080/api/leaderboard?stage_id=stage-001" \
     -H "Authorization: Bearer $USER_TOKEN"
   ```

