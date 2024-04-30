FROM golang:alpine
WORKDIR /app
COPY main.go .
RUN apk --no-cache add curl && apk add pingu
RUN go mod init main
RUN go build ./main.go
EXPOSE 80
CMD ["./main"]