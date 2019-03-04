#!/bin/sh -e
main=${PWD##*/} 
go test -v -cover $main/businesslogic