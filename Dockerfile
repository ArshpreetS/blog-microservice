FROM golang

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build -o blog_service

CMD ["./blog_service"]
