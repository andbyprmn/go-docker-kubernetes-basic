FROM golang:alpine

RUN apk update && apk add git

ARG DEFAULT_PORT=9024
ARG DEFAULT_INSTANCE_ID="--build-arg Dockerfile"

ENV PORT=${DEFAULT_PORT}
ENV INSTANCE_ID=${DEFAULT_INSTANCE_ID}

## debug cmd
RUN echo "Isi dari build argument DEFAULT_PORT adalah ${DEFAULT_PORT}"
RUN echo "Isi dari build argument DEFAULT_INSTANCE_ID adalah ${DEFAULT_INSTANCE_ID}"
RUN echo "Isi dari env var PORT adalah ${PORT}"
RUN echo "Isi dari env var INSTANCE_ID adalah ${INSTANCE_ID}"

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o binary

ENTRYPOINT [ "./binary" ]