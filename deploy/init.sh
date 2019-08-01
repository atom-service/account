#! /bin/bash

RELEASES_VERSION=$1 # 安装的版本
INSTALL_PATH=$2 # 安装目录
FILE_NAME=$3 # 包名
MYSQL=$2 # mysql 配置
PORT=$1 # 服务端口

# 安装程序
install() {
  tmpdir=`mktemp`
  # 下载文件
  wget https://github.com/grpcbrick/account/releases/download/$RELEASES_VERSION/$FILE_NAME $tmpdir
  if [ $? != 0 ]; then
    exit 1
  fi

  tar -zxvf $tmpdir/$FILE_NAME $INSTALL_PATH
  rm -rf $tmpdir/$FILE_NAME

}

# 安装服务文件
installUnitService() {

  # 安装目录
  # systemctl daemon-reload
  # systemctl start comm
  # /etc/systemd/user/
}
