## We specify the base image we need for our
## go application
FROM golang:1.14.0-alpine
## We create an /product-api directory within our
## image that will hold our application source
## files
RUN mkdir /product-api
## We copy everything in the root directory
## into our /app directory
ADD . /product-api
## We specify that we now wish to execute 
## any further commands inside our /product-api
## directory
WORKDIR /product-api
## we run go build to compile the binary
## executable of our Go program
RUN go build -o main .
## Our start command which kicks off
## our newly created binary executable
CMD ["/product-api/main"]