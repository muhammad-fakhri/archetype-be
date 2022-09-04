### STAGE 1 : Build the go source code into binary
FROM golang:1.14 as builder


ENV APP_DIR /go/src/github.com/muhammad-fakhri/archetype-be 
ENV GOFLAGS -mod=vendor

## Copy source code from local machine into container
RUN mkdir -p ${APP_DIR}
COPY . ${APP_DIR}

# Compile the binary and statically link
RUN cd $APP_DIR && CGO_ENABLED=0 go build -ldflags '-d -w -s'

### STAGE 2 : Package the binary in a minimal alpine base image
FROM alpine:3.13

ENV APP_DIR github.com/muhammad-fakhri/archetype-be 

COPY --from=builder /go/src/${APP_DIR}/archetype-be .