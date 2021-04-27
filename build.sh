#!/usr/bin/env bash

export APP_NAME="alcochange-dtx"
export ENV="dev"
#export PORT=9010
  
# Sensitive environment variables are added in dev.env
# source dev.env

#swag init
# Swag init done

go install
if [ $? != 0 ]; then
  echo "## Build Failed ##"
  exit
fi

echo "Restoring all vendor versions ..."
godep restore
echo "Done."

echo "Doing some cleaning ..."
go clean
echo "Done."

echo "Running goimport ..."
goimports -w=true .
echo "Done."

# echo "Running go vet ..."
# go vet ./internal/...
# if [ $? != 0 ]; then
#   exit
# fi
# echo "Done."

#echo "Running go generate ..."
#go generate ./internal/...
#echo "Done."

echo "Running go format ..."
gofmt -w .
echo "Done."

echo "Running go build ..."
go build -race
if [ $? != 0 ]; then
  echo "## Build Failed ##"
  exit
fi
echo "Done."

#echo "Running unit test ..."
#go test -parallel 1 ./internal/...
#export PG_DATABASE_URL='ER9OaxWlwxYday7oZ7-Wecnq9HNvAHy4h8BW-0uShA3NcMajFtBMahDeO-XI_y92eaDpjH-bt9nItiLfIDsfFyzgLjxZfGn2qbA3WBd2PcztpGMtdCf6QNWbFp-glIY8f-tMVLGP-Gpl4LIue_pH-nh5QO-69eKmp2ORbB4OY_9VZ8tiZAexTGd3'

if [ $? == 0 ]; then
    echo "Done."
	echo "## Starting service ##"
    ./go-alcochange-dtx-ga -conf conf/dev.json
fi
