FROM golang:1.16 as build-go
WORKDIR /src
COPY . /src
ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux
RUN go build -o program

FROM scratch
ADD ihm /
COPY --from=build-go /src/program /program
EXPOSE 8000
