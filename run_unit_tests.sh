#!/bin/sh -e
main=${PWD##*/} 
go test -v -cover github.com/renaudcalmont/$main/businesslogic
