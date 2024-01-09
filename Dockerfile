FROM golang:alpine AS Buildstage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o binary

EXPOSE 8085

CMD ./binary