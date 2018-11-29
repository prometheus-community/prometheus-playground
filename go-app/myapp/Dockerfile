ARG goVersion

FROM golang:${goVersion} AS builder

# Args
ARG goPackage

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/${goPackage}
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY --from=builder /app ./
ENTRYPOINT ["./app"]