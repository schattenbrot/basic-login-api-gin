#
# Build go app
#
FROM golang:1.17 AS builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -v -o api ./cmd/api

# 
# Run app
# 
FROM scratch

WORKDIR /run

ENV CORS *

COPY /environment/.env ./environment/.env
COPY --from=builder /go/src/api ./

ENTRYPOINT [ "./api", "-cors=http://localhost:3000" ]