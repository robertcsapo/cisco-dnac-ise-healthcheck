# Build Image
# Build Image
FROM golang:1.9.2-alpine3.7 as build
#FROM golang:onbuild as build

#RUN apk add --no-cache git
#RUN go get github.com/golang/dep/cmd/dep

#COPY talos/ /go/src/cisco-talos-tcpwrapper
#WORKDIR /go/src/cisco-talos-tcpwrapper
#WORKDIR /go/src/
#RUN git clone https://gitlab.com/robertcsapo/cisco-talos-tcpwrapper.git
WORKDIR /go/src/cisco-dnac-ise-healthcheck/
ADD main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Run container
FROM scratch
#COPY --from=build /go/src/ /go/src/
#ADD ca-certificates.crt /etc/ssl/certs/
#COPY --from=build /go/src/cisco-talos-tcpwrapper/main /go/src/cisco-talos-tcpwrapper/main
COPY --from=build /go/src/cisco-dnac-ise-healthcheck/main /go/src/cisco-dnac-ise-healthcheck/main
#COPY --from=build /go/src/cisco-talos-tcpwrapper/cisco-talos-ip-blacklist /go/src/cisco-talos-tcpwrapper/cisco-talos-ip-blacklist
WORKDIR /go/src/cisco-dnac-ise-healthcheck/
#RUN touch cisco-talos-ip-blacklist

#CMD ["./main"]
ENTRYPOINT ["./main"]
