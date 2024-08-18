# notesapp under development

something like google keep

frontend - react ts
backend - golang gin

features ->

   * login and signup 
   * CRUD notes
   * add notes with images
   * save versions of notes
   * light dark mode
   * email notification for reminders -> SendGrid, gomail

-- project structure

├── cmd
│   └── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── auth
│   ├── db
│   │   └── connection.go
│   └── user
│       ├── handler.go
│       ├── handler_test.go
│       ├── service.go
│       └── service_test.go
└── pkg
    ├── middleware
    └── utils
