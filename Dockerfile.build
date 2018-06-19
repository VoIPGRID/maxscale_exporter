FROM golang:1.8
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go get github.com/skord/maxscale_exporter
RUN make build
RUN mv /go/src/app/maxscale_exporter /bin/maxscale_exporter
ENTRYPOINT ["/bin/maxscale_exporter"]
USER nobody
EXPOSE 9195
