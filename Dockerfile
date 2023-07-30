# Copyright (c) Xiaodong Liu.
FROM golang:1.20-alpine3.18

WORKDIR /tmp/src

COPY . .

EXPOSE 80

# build app
RUN go mod tidy && \
  go build -o /usr/local/bin/app && \
  rm -rf /tmp/src

RUN rm -rf /tmp/src

WORKDIR /usr/local/bin

# Configure container startup
CMD ["app"]
