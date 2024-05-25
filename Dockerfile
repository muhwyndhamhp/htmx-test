FROM golang:1.22-alpine3.19 as builder

WORKDIR /app

COPY . ./

# Do dep installs outside, due to private git modules
# RUN make dep

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

RUN apk update && apk add libwebp

COPY --from=builder /app/main /app/
COPY --from=builder /app/scripts /app/scripts
COPY --from=builder /app/src /app/src

EXPOSE 8080

CMD [ "/app/main" ]

