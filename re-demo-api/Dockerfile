# Dockerfile for 're-demo-api'
# build stage
FROM golang:1.21-alpine AS build

# create work directory
RUN mkdir /app

# switch to work directory
WORKDIR /app

# copy all files
ADD . .

# download dependencies
RUN go mod download

# build application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o re-demo-api .

# deploy stage
FROM alpine:3.18 AS deploy

# create work directory
RUN mkdir /app

# switch to work directory
WORKDIR /app

# copy application artifacts to current directory
COPY --from=build /app/re-demo-api .

# expose port
EXPOSE 80

# run application
ENTRYPOINT ["/app/re-demo-api"]
