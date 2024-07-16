FROM golang:1.22.2
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /bin/app ./cmd
WORKDIR /src

COPY --from=0 /bin/app /bin/app
COPY --from=0 /src/asserts /src/asserts

ENTRYPOINT ["./bin/app"]