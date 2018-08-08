# build container
FROM golang:1.10 as build
ADD . /go/src/go-http-server
RUN CGO_ENABLED=0 go install go-http-server

# runtime container
FROM scratch
COPY --from=build /go/bin/go-http-server /go-http-server
EXPOSE 8080:8080
CMD ["/go-http-server"]
