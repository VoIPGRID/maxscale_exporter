FROM golang:1.15 AS build

WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN make build

FROM alpine

COPY --from=build /go/src/app/maxscale_exporter /bin/maxscale_exporter
USER nobody
EXPOSE 9195
ENTRYPOINT ["/bin/maxscale_exporter"]
