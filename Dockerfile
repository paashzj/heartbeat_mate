FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o heartbeat_mate .


FROM ttbb/heartbeat:nake

LABEL maintainer="shoothzj@gmail.com"

COPY docker-build /opt/sh/heartbeat/mate

COPY --from=build /opt/sh/compile/pkg/heartbeat_mate /opt/sh/heartbeat/mate/heartbeat_mate

WORKDIR /opt/sh/heartbeat

CMD ["/usr/bin/dumb-init", "bash", "-vx", "/opt/sh/heartbeat/mate/scripts/start.sh"]
