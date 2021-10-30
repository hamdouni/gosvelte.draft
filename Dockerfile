FROM golang:1.17 as build-go
WORKDIR /src
COPY . /src
ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux
RUN go build -o server ./cmd/server

FROM scratch
COPY --from=build-go /src/server /
EXPOSE 8000
