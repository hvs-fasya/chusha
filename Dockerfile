FROM alpine:3.7
RUN apk update && apk add --no-cache ca-certificates tzdata

COPY VERSION /opt/VERSION
COPY chusha /opt/app
COPY front/stage /opt/front/stage
COPY front/dist /opt/front/dist

ENTRYPOINT ["/opt/app"]
