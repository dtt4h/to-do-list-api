# 🚀 Quick Start  
Clone and run  
`git clone https://github.com/d0n77ryth1s4th0m3/to-do-list-api.git`  
`cd to-do-list-api`    
`go run main.go ` 

API will be available at:  
http://localhost:8080

# 🛠️Tech Stack  
Backend:  
- Go 1.25+ - Main server language    
- Gin Web Framework - Routing and middleware
  
Database:  
- SQLite3 - Lightweight embedded database  
- modernc.org/sqlite - Pure Go driver (no CGO required)
    
API:  
- RESTful architecture  
- JSON data format  
- CRUD operations
- JWT Authentication

# 📚Endpoints  
GET /tasks - get all tasks  
GET /tasks/:id - get task by ID  
POST /tasks - create new task  
PATCH /tasks/:id - update task  
DELETE /tasks/:id - delete task

