#Build stage

FROM golang:1-alpine3.15 As builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Run stage

FROM alpine:3.15
WORKDIR /app 
COPY --from=builder /app/main .
COPY --from=builder /app/migration ./migration
COPY app.env .
COPY db/migration ./migration
EXPOSE 8888
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/main" ]