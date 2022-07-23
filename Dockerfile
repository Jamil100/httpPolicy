FROM golang:1.16-alpine


#RUN apk add git
RUN mkdir src/app
WORKDIR /src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

#COPY cmd/web/server.go ./
#COPY pkg/handlers/handler.go ./

#ADD . /app
#WORKDIR /app
#RUN go build -o cmd/web/server.go .
#CMD ["/app/cmd/web/server"]

RUN go build -o server cmd/web/server.go

CMD [ "/src/app/server" ]