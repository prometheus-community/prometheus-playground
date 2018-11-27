FROM golang:1.10 AS builder

# Args
ARG goPackage

# Get Dep
RUN go get -u github.com/golang/dep/cmd/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/${goPackage}
COPY . ./
RUN dep ensure --vendor-only
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY --from=builder /app ./
ENTRYPOINT ["./app"]