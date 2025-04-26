# Authentication Application using Golang

- Gorilla Mux (for routing)
- JWT (JSON Web Tokens) for token-based auth
- bcrypt for secure password hashing
- Optionally PostgreSQL or SQLite for user storage

```bash
/auth-app
│
├── main.go
├── handlers/
│   └── auth.go
├── models/
│   └── user.go
├── middleware/
│   └── auth.go
```

## concepts
- ✅ Setting up your Go project
- 🔐 Creating user registration with hashed passwords
- 🔑 Logging in and generating JWT tokens
- 🛡️ Middleware to protect private routes
- 🔄 Token validation and refresh (optional)
