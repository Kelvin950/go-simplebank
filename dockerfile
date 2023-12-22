FROM golang:1.20-alpine3.18  AS  builder
WORKDIR  /app 
COPY  . .  
RUN go build  main.go 

FROM alpine:3.18
WORKDIR /app
COPY --from=builder  /app/main  .
COPY  app.env .
EXPOSE 8080
CMD [ "/app/main" ]