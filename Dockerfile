FROM golang:1.19 as BUILDER

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=$GOOS
ENV GOARCH=$GOARCH

WORKDIR /app
COPY . .

RUN go mod download && go mod verify && go build -o build/app main.go

FROM scratch AS FINAL

WORKDIR /main
COPY --from=BUILDER /app/build/app .
EXPOSE 8080
CMD ["./app"]
