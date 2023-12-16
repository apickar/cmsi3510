#!/bin/sh

go install --gcflags="-N -l" github.com/apickar/hello
go install github.com/apickar/debugger