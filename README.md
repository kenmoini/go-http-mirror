# Golang HTTP Mirror

This application will take a list of URLs, download them to a specific path, optionally overwrite existing files, and then mirror them via an HTTP server.  This is handy for ZTP where you may need to mirror RHCOS ISOs and Root Filesystem blobs.

## Example Configuration

You can find some example configuration in the `container_root/etc/http-mirror/config.yml` file.

The default configuration has files from `/tmp/server/pub` served to `http://localhost:8080/pub/`

## Building & Running

### Go

```bash
## Get the modules
go mod tidy

## Build the binary
go build -o http-mirror

## Run the binary with the example configuration
./http-mirror -config=./container_root/etc/http-mirror/config.yml
```

### Container

```bash
## Build the container
podman build -t http-mirror .

## or just pull it from the pre-built image
podman pull quay.io/kenmoini/go-http-mirror:latest

## Run the container with the example configuration
podman run -d --rm --name http-mirror -p 8080:8080 http-mirror
podman run -d --rm --name http-mirror -p 8080:8080 -v ./container_root/etc/ztp-mirror:/etc/http-mirror http-mirror
```

## Deploying to OpenShift

You can easily deploy to OpenShift with the pre-provided YAML manifests.  Take note of the commented-out portions of the `deploy/03-deployment.yaml` file for suggestions on how to mount custom Root CAs and Proxy Configuration.

```bash
## Create a project
oc new-project http-mirror

## Deploy the ConfigMaps, PVC, Deployment, Service, and Route
oc apply -R -f deploy/
```
