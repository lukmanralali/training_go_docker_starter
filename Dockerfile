FROM golang:1.12-alpine

RUN apk update && \
    apk upgrade && \
    apk add bash git && \
    apk add gcc && \
    apk add musl-dev && \
    apk add curl

# We create an app directory within our image that will hold our application source file
RUN mkdir /go/src/my_application_name

# We copy everything in the root directory into our app directory
COPY . /go/src/my_application_name

# We specify that we now wish to execute any further commands inside our app directory
WORKDIR /go/src/my_application_name