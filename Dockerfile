FROM golang:1.25

WORKDIR /app

RUN go install github.com/air-verse/air@latest

ENV PATH="/go/bin:${PATH}"

CMD ["air"]

