# K8smm Pod for lab

[![Docker](https://github.com/mm-k8s-ug/generator/workflows/Docker/badge.svg)](https://github.com/mm-k8s-ug/generator/actions?query=workflow%3ADocker)
[![Docker Repository on Quay](https://quay.io/repository/dther/hola/status "Docker Repository on Quay")](https://quay.io/repository/dther/hola:generator)

## Testing container for pod. writing with golang and lightweghit container.
It just generate content to /var/htdocs as index.html.

```bash
$ docker built -t generator .
$ docker run -d -p 8080:8080 generator
$ docker images

```
