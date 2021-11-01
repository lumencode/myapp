FROM golang:1.17-alpine

COPY . /app

WORKDIR /app

RUN go build -o main .

EXPOSE 4000

CMD ./main
