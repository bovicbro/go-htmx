FROM golang:1.22.1-alpine3.18

WORKDIR /app

COPY . .

RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go build -o /godocker

EXPOSE 3000 

CMD ["/godocker"]
