FROM registry.access.redhat.com/ubi8/go-toolset:latest AS build

WORKDIR /opt/app-root/src
COPY . .
RUN go build -o http-mirror

FROM registry.access.redhat.com/ubi9/ubi-micro:latest AS bin
#FROM scratch AS bin

USER 0

RUN microdnf update -y && \
    microdnf clean all && \
    rm -rf /var/cache/yum

USER 1001

COPY --from=build /opt/app-root/src/http-mirror /usr/local/bin/
COPY container_root/ /

EXPOSE 8080

CMD [ "/start.sh" ]
