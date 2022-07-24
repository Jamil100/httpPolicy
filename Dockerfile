FROM golang:1.16-alpine

RUN mkdir src/app
WORKDIR /src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o server cmd/web/server.go

CMD [ "/src/app/server" ]