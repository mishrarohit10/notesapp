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

server/
├── cmd/
│   └── main.go
├── internal/
│   ├── user/
│   │   ├── handler.go
│   │   ├── handler_test.go   # Tests for handlers
│   │   ├── service.go
│   │   └── service_test.go   # Tests for services
│   ├── order/
│   │   ├── handler.go
│   │   ├── handler_test.go   # Tests for handlers
│   │   ├── service.go
│   │   └── service_test.go   # Tests for services
│   └── product/
│       ├── handler.go
│       ├── handler_test.go   # Tests for handlers
│       ├── service.go
│       └── service_test.go   # Tests for services
├── pkg/
│   ├── middleware/
│   │   └── auth.go
│   │   └── auth_test.go      # Tests for middleware
│   └── utils/
│       └── helpers.go
│       └── helpers_test.go   # Tests for utilities
├── go.mod
├── go.sum
└── README.md
