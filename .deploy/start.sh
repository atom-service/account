#! /bin/bash
# if [ $1 -eq "debug" ]; then
    go run ./main.go -rpc_port :3000  -mysql_url "root:root@tcp(localhost:3306)/default?charset=utf8mb4"
    echo 1
    cd ./standard && letmegrpc --addr=localhost:3000 --port=8080 ./standard.proto  && cd ../ 
    echo 2
# if