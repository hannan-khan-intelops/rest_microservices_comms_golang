FROM golang:latest

WORKDIR /usr/src/app

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./artifact

ENV PORT 4317
EXPOSE 4317

RUN ls

#ENTRYPOINT [ "tail", "-f", "/dev/null" ]

ENTRYPOINT [ "./artifact" ]
