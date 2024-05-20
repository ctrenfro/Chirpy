# Chirpy

## User resource

```json
{
  "id": "0159b39d-4e28-b19b-198b3ceea0dd",
  "email": "chris@gmail.com",
  "age": 11
}
```

### GET /v1/users

Returns an array of users.

### POST /v1/users

Request body:

```json
{
  "email": "chris@gmail.com",
  "age": 11
}
```

Response body:

```json
{
  "id": "0159b39d-4e28-b19b-198b3ceea0dd",
  "email": "chris@gmail.com",
  "age": 11
}
```
