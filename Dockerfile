# Build Image
FROM golang:1.9.2-alpine3.7 as build
WORKDIR /go/src/cisco-dnac-ise-healthcheck/
ADD main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Build Container
FROM scratch
COPY --from=build /go/src/cisco-dnac-ise-healthcheck/main /go/src/cisco-dnac-ise-healthcheck/main
WORKDIR /go/src/cisco-dnac-ise-healthcheck/
ENTRYPOINT ["./main"]
