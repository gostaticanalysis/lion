#!/bin/sh

 GO111MODULE=off GOPATH=`pwd`/../.. go test -coverprofile=coverprofile .
