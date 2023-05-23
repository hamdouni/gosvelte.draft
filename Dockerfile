FROM golang:1.20 as build-go
WORKDIR /src
COPY . /src
ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux
RUN go build -o ./build/ cmd/srv/*.go

FROM scratch
COPY --from=build-go server /
CMD ["/server"]
EXPOSE 80
