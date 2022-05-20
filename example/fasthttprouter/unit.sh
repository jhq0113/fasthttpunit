#!/bin/bash
SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)

cd $SCRIPTPATH

go test . -v -coverpkg=... -coverprofile=$SCRIPTPATH/unitout/app.out
go tool cover -func=$SCRIPTPATH/unitout/app.out -o $SCRIPTPATH/unitout/coverage.txt
go tool cover -html=$SCRIPTPATH/unitout/app.out -o $SCRIPTPATH/unitout/coverage.html