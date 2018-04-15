FROM alpine

COPY ProxyGrabber .
COPY index.html .
COPY static /static

RUN chmod +x ProxyGrabber &&\
    apk add --no-cache ca-certificates

ENTRYPOINT [ "./ProxyGrabber" ]