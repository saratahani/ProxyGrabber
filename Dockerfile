FROM alpine

COPY ProxyGrabber .
COPY ./static /static

RUN apk update && apk add --no-cache ca-certificates

ENTRYPOINT [ "./ProxyGrabber" ]