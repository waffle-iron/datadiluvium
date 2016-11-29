FROM golang:latest
# Reference: https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
