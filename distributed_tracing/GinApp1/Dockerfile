FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o GinApp1 .

EXPOSE 8080

CMD [ "./GinApp1" ]