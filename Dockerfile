FROM registry.access.redhat.com/ubi8/go-toolset:1.17.7-13 AS build

WORKDIR /opt/app-root/src
COPY . .
RUN go build -o http-mirror

FROM registry.access.redhat.com/ubi9/ubi-micro:9.0.0-11 AS bin
#FROM scratch AS bin

COPY --from=build /opt/app-root/src/http-mirror /usr/local/bin/
COPY container_root/ /

EXPOSE 8080

CMD [ "/start.sh" ]