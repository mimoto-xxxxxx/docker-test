#!/bin/bash
export GOROOT=/root/go
export GOPATH=$GOROOT/gopath
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
go run /root/sel.go $CHROME_PORT_4444_TCP_ADDR $CHROME_PORT_4444_TCP_PORT
