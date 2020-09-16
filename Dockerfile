FROM alpine:3.9

ADD main .

RUN apk add --no-cache bash && apk add ca-certificates

EXPOSE 80
EXPOSE 8080

CMD ["./main"]



