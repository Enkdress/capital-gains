FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN make build

FROM golang:1.22 as tester
WORKDIR /app
COPY . .
RUN go mod download
RUN ["make" ,"test"]

FROM alpine:latest
RUN apk add make
COPY --from=builder /app/bin/capitalgains /app/bin/capitalgains
COPY --from=builder /app/input.txt /app/input.txt

# See if I can get the tests run with this
# If the build completes means that there
# are no error with the tests
COPY --from=tester /app/Makefile /app/Makefile

WORKDIR /app
CMD ["make", "run"]


