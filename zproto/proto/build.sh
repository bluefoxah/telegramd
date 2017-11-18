#!/bin/sh
SRC_DIR=.
DST_DIR=..

protoc -I=$SRC_DIR --go_out=plugins=grpc:$DST_DIR/ $SRC_DIR/*.proto

