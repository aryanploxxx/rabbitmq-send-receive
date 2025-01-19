# RabbitMQ - Message Broker

## Directory Structure 
```
rabbitmq-send-receive/
├── recieve/
    ├── recieve.go
├── send/
    ├── send.go
├── README.md
├── docker-compose.yml
├── go.mod
└── go.sum
```

## Project Setup

1. Setup Dependencies
``` golang
  go get "github.com/rabbitmq/amqp091-go"
```
2. Initialize RabbitMQ Image
``` golang
  docker compose up
```
3. Run Producer and Consumer Servers
``` golang
  cd send
  go run send.go

  cd ../recieve
  go run recieve.go

```

## Previews
![image](https://github.com/user-attachments/assets/0492066a-48c3-46c1-98b6-235da4d8c4ad)
![image](https://github.com/user-attachments/assets/42fb880d-44e6-41b1-8506-751c4e62b5a7)
