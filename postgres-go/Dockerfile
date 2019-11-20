FROM golang:1.12-alpine3.10

# Move to the directory where the source code will live
WORKDIR /go/src/hello

# Install git which will be used by go get
RUN apk add git

# Copy the source code into the current directory
COPY . .

# Get any dependencies, and compile the code
RUN CGO_ENABLED=0 go get -v ./...

# Second stage: start from an empty base image
FROM scratch

# Copy the binary from the first stage
COPY --from=0 /go/bin/hello /

# Tell Docker what executable to run by default when starting this container
ENTRYPOINT ["/hello"]
