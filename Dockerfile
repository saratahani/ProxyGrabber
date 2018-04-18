FROM alpine

COPY ./code/code .
COPY ./code/template/static /template/static
COPY ./code/template/index.html /template

RUN chmod +x code &&\
    apk add --no-cache ca-certificates

ENTRYPOINT [ "./code" ]