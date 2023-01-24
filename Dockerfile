FROM golang:alpine

# Setup go environment variables
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Configure application working directories
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Change working directory
WORKDIR $GOPATH/app

# Install dependencies
ENV GO111MODULE=on
COPY . ./

ENV PORT 8080
ENV GIN_MODE release
EXPOSE 8080

RUN go build -o app
CMD ["/app"]