FROM alpine

COPY ProxyGrabber .
COPY ./code/template/static /code/template/static
COPY ./code/template/index.html /code/template

RUN chmod +x ProxyGrabber &&\
    apk add --no-cache ca-certificates

ENTRYPOINT [ "./ProxyGrabber" ]