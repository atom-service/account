#! /bin/bash

MAIN_FILE=$1 # 入口文件名称

systemctl restart $MAIN_FILE.service
if [ $? != 0 ]; then
  echo "服务启动失败"
  exit 1
fi
