FROM alpine

COPY ProxyGrabber .
COPY ./code/template/static /template/static
COPY ./code/template/index.html /template

RUN chmod +x ProxyGrabber &&\
    apk add --no-cache ca-certificates

ENTRYPOINT [ "./ProxyGrabber" ]