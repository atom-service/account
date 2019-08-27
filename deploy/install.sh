#! /bin/bash
# 服务端执行的脚本

# 参数
RELEASES_VERSION=$1 # 安装的版本

# 环境变量
TEMPDIR="`mktemp -d`" #缓存目录
SERVICE_NAME="account" # 服务的名字
PACKAGE_NAME="linux_amd64.tar.gz" # 指定包名
SERVICE_PATH="/etc/systemd/user" # uint service 安装目录
INSTALL_PATH="/usr/local/account" # 程序的安装目录

ready() {
  # 准备安装目录
  if [ ! -d $INSTALL_PATH ]; then
    mkdir -p $INSTALL_PATH ;
    if [ $? != 0 ]; then
      echo "创建安装目录失败"
      exit 1
    fi
	# 下载文件
	echo "wget https://github.com/grpcbrick/account/releases/download/$RELEASES_VERSION/$PACKAGE_NAME to $TEMPDIR/$PACKAGE_NAME ;"
	wget -O $TEMPDIR/$PACKAGE_NAME https://github.com/grpcbrick/account/releases/download/$RELEASES_VERSION/$PACKAGE_NAME ;
	if [ $? != 0 ]; then
    echo "安装包下载失败"
    exit 1
	fi
  fi
}


# 安装程序
install() {
  # 解压文件
  tar -zxvf $TEMPDIR/$PACKAGE_NAME -C $INSTALL_PATH ;
  if [ $? != 0 ]; then
    echo "安装包解压失败"
    exit 1
  fi

  # 清理包
  rm -rf $TEMPDIR ;
  if [ $? != 0 ]; then
    echo "安装包清理失败"
    exit 1
  fi
}

# 安装服务文件
installUnitService() {
  if [ ENV_FILE ]; then
    cp -fp ENV_FILE $INSTALL_PATH/$MAIN_FILE.env ;
    if [ $? != 0 ]; then
      echo "env 文件拷贝失败"
      exit 1
    fi
  fi

  echo "[Unit]"                                                      > $SERVICE_PATH/$MAIN_FILE.service ;
  echo "Description=SERVICE_NAME service"                           >> $SERVICE_PATH/$MAIN_FILE.service ;
  echo "[Service]"                                                  >> $SERVICE_PATH/$MAIN_FILE.service ;
  echo "Type=simple"                                                >> $SERVICE_PATH/$MAIN_FILE.service ;
  echo "Restart=always"                                             >> $SERVICE_PATH/$MAIN_FILE.service ;
  echo "ExecStart=$INSTALL_PATH/$MAIN_FILE"                         >> $SERVICE_PATH/$MAIN_FILE.service ;
  echo "EnvironmentFile=$INSTALL_PATH/$MAIN_FILE.env"               >> $SERVICE_PATH/$MAIN_FILE.service ;
  echo "[Install]"                                                  >> $SERVICE_PATH/$MAIN_FILE.service ;
  echo "WantedBy=multi-user.target"                                 >> $SERVICE_PATH/$MAIN_FILE.service ;

  systemctl daemon-reload ;
  if [ $? != 0 ]; then
    echo "服务安装失败"
    exit 1
  fi
}

# 执行
ready ;
install ;
installUnitService ;
