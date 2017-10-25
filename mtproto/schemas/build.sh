#!/bin/sh
SRC_DIR=.
DST_DIR=.

./codegen.sh
protoc -I=$SRC_DIR --go_out=$DST_DIR/ $SRC_DIR/*.proto
