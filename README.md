# geoip (Go demonstration)
A minimalistic *HTTP + JSON* (let insist on it: __not__ *RESTful*) web service with an enterprise-grade code organisation.

The payload is borrowed from different examples, see the references below for details.

This project focuses on good practices for maintainability.

## What is that actually?

## What is it not?

## How to try it?

### See available parameters:
```shell 
$ docker run --rm geoip -h
```
Launch and immediately destroy a container to show the quick help about CLI parameters.

### Example lunch command:
```shell
$ docker run -d -u 33 -p 5002:5000 geoip
```
This will start a container in the background with the `www-data` identity on Debian-derived systems. The container will accept queries on localhost 5002/TCP port.

### Example query:
```shell
$ curl -XPOST -H "Content-Type: application/json" -d '{"address":"<any>"}' http://localhost:5002/country
```
Replace `<any>` with your public IPv4 or IPv6 address or a purposely misspelled one.  
