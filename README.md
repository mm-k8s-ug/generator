# K8smm Content Generator Pod for lab

[![Go](https://github.com/mm-k8s-ug/generator/workflows/Go/badge.svg)](https://github.com/mm-k8s-ug/generator/actions?query=workflow%3AGo")
[![Docker Repository on Quay](https://quay.io/repository/dther/hola/status "Docker Repository on Quay")](https://quay.io/repository/dther/hola?tab=tags)

## Testing container for pod. writing with golang and lightweghit container.
It just generate content to /var/htdocs as index.html.

```bash
$ docker built -t generator .
$ docker run -d -p 8080:8080 generator
$ docker images

```
