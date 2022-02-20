#!/bin/bash
trap "rm server/server;rm client/client;rm nohup.out;kill 0" EXIT

START_PATH=$(pwd)

protoc  -I. -I api -I "$START_PATH"/api/googleapis \
--go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. api/*.proto


go build -o server ./server/server.go &&
go build -o client ./client/client.go &&

nohup ./server/server &
sleep 3
nohup ./client/client &

wait

