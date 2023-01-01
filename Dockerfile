FROM golang:1.19.4-alpine

RUN go install github.com/cosmtrek/air@latest

WORKDIR /workspace/app/bingo

CMD ["air"]