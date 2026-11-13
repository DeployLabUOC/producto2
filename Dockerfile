FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o goapp .

FROM alpine:3.20
COPY --from=build /app/goapp /goapp
EXPOSE 8080
CMD ["/goapp"]