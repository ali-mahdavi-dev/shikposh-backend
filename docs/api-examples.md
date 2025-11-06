# API Request Examples

## Register User

### Endpoint

```
POST /api/v1/public/register
```

### Request Body

```json
{
  "avatar_identifier": "user123",
  "user_name": "johndoe",
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com",
  "password": "SecurePass123"
}
```

### Validation Rules

- `avatar_identifier`: required
- `user_name`: required, minimum 3 characters
- `first_name`: required, minimum 3 characters
- `last_name`: required, minimum 3 characters
- `email`: required, must be a valid email format
- `password`: required, minimum 6 characters

### Success Response (200)

```json
{
  "success": true,
  "data": {
    "user_id": 1
  }
}
```

### Error Response Examples

#### Validation Error (422)

```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "field 'Email' failed validation: must be a valid email address; field 'Password' failed validation: must be at least 6 characters",
    "status": "Unprocessable Entity"
  }
}
```

#### Conflict Error - User Already Exists (409)

```json
{
  "success": false,
  "error": {
    "code": "USER_ALREADY_EXISTS",
    "message": "User already exists",
    "status": "Conflict"
  }
}
```

---

## Login User

### Endpoint

```
POST /api/v1/public/login
```

### Request Body

```json
{
  "user_name": "johndoe",
  "password": "SecurePass123"
}
```

### Validation Rules

- `user_name`: required
- `password`: required

### Success Response (200)

```json
{
  "success": true,
  "data": {
    "access": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MTAwMDAwMDB9.example"
  }
}
```

### Error Response Examples

#### Validation Error (422)

```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "field 'User_name' failed validation: is required",
    "status": "Unprocessable Entity"
  }
}
```

#### Unauthorized Error - Invalid Credentials (401)

```json
{
  "success": false,
  "error": {
    "code": "USER_NOT_FOUND",
    "message": "User not found",
    "status": "Unauthorized"
  }
}
```

#### Not Found Error - User Not Found (404)

```json
{
  "success": false,
  "error": {
    "code": "USER_NOT_FOUND",
    "message": "User not found",
    "status": "Not Found"
  }
}
```

---

## Logout User

### Endpoint

```
POST /api/v1/public/logout
```

### Headers

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### Request Body

No body required. User ID is extracted from JWT token.

### Success Response (200)

```json
{
  "success": true
}
```

### Error Response Examples

#### Unauthorized Error - No Token (401)

```json
{
  "success": false,
  "error": {
    "code": "USER_NOT_FOUND",
    "message": "User not found",
    "status": "Not Found"
  }
}
```

---

## cURL Examples

### Register User

```bash
curl -X POST http://localhost:8000/api/v1/public/register \
  -H "Content-Type: application/json" \
  -d '{
    "avatar_identifier": "user123",
    "user_name": "johndoe",
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "password": "SecurePass123"
  }'
```

### Login User

```bash
curl -X POST http://localhost:8000/api/v1/public/login \
  -H "Content-Type: application/json" \
  -d '{
    "user_name": "johndoe",
    "password": "SecurePass123"
  }'
```

### Logout User

```bash
curl -X POST http://localhost:8000/api/v1/public/logout \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```
