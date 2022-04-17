FROM golang:1.17.8 as BUILDER

WORKDIR /app

ADD go.mod go.mod

ADD . .

RUN go mod tidy

RUN go build -o main cmd/dfw/main.go

FROM golang:1.17.8 as RUNNER

EXPOSE 8080

WORKDIR /app

COPY --from=BUILDER /app/main .

COPY templates templates
COPY static static

RUN chmod +x main
RUN ls 

ENTRYPOINT ["./main"]

