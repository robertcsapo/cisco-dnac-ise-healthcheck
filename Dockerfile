# Build Binary
FROM golang:1.17.6-bullseye as build
WORKDIR /go/src/cisco-dnac-ise-healthcheck/
ADD . .
RUN /bin/bash build.sh

# Build Container Image
FROM scratch
COPY --from=build /go/src/cisco-dnac-ise-healthcheck/bin/cisco-dnac-ise-healthcheck-linux-amd64 /app/
WORKDIR /app/
ENTRYPOINT ["./cisco-dnac-ise-healthcheck-linux-amd64"]
