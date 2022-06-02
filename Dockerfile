FROM registry.access.redhat.com/ubi8/go-toolset:1.17.7-13 AS build

ENV something something
WORKDIR /opt/app-root/src
COPY . .
RUN go build -o ztp-mirror

FROM scratch AS bin

COPY --from=build /opt/app-root/src/ztp-mirror /usr/local/bin/
COPY container_root/ /

EXPOSE 8080

CMD [ "ztp-mirror -config /etc/ztp-mirror/config.yml" ]