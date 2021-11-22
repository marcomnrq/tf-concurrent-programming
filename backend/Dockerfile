FROM golang:latest

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY temperature.csv ./
COPY testing.csv ./
COPY training.csv ./
RUN go mod download

COPY *.go ./

EXPOSE 4000

CMD ["go", "run", "api.go"]
