# device-sec
Code structure
```
project/
├── cmd/
│   ├── server/
│   │   └── main.go
│   └── agent/
│       └── main.go
├── internal/
│   ├── models/
│   │   └── command.go
│   ├── server/
│   │   ├── api/
│   │   ├── queue/
│   │   └── database/
│   └── agent/
│       ├── websocket/
│       └── handler/
├── go.mod
└── README.md
```