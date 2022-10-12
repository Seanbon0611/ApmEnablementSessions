FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o example_app .

EXPOSE 80

CMD [ "./example_app" ]