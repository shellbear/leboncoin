FROM golang:1.16 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o leboncoin .

FROM scratch

COPY --from=builder /app/leboncoin /

ENTRYPOINT ["/leboncoin"]