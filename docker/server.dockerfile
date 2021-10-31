FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY ./src ./src

RUN go build

ARG PORT
EXPOSE ${PORT}

CMD [ "./devices" ]