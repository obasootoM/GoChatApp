#Build stage

FROM golang:1-alpine3.15 As builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Run stage

FROM alpine
WORKDIR /app 
COPY --from=builder /app/main .
CMD [ "/app/main" ]