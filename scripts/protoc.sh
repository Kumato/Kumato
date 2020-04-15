#!/bin/bash

SRC_DIR="internal/types"

echo "generate .pb.go files ..."
#protoc -I ${SRC_DIR}/worker/ ${SRC_DIR}/worker/worker.proto --go_out=plugins=grpc:${SRC_DIR}/worker
#protoc -I ${SRC_DIR}/task/ ${SRC_DIR}/task/task.proto --go_out=plugins=grpc:${SRC_DIR}/task
#protoc -I ${SRC_DIR}/controller/ ${SRC_DIR}/controller/controller.proto --go_out=plugins=grpc:${SRC_DIR}/controller
protoc -I ${SRC_DIR}/ ${SRC_DIR}/common.proto ${SRC_DIR}/controller.proto ${SRC_DIR}/task.proto ${SRC_DIR}/worker.proto --go_out=plugins=grpc:${SRC_DIR}
