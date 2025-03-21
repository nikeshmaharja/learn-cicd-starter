FROM golang:1.23.0

RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /app

COPY . .

RUN go build -o /usr/bin/notely

CMD ["notely"]