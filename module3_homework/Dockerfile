FROM golang:1.16-alpine AS build

WORKDIR /app
COPY go.mod go.sum *.go ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o /server

FROM scratch
COPY --from=build /server /server
EXPOSE 8000
ENTRYPOINT ["/server"]