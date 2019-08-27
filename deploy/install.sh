#! /bin/bash
# 服务端执行的脚本

# 参数
RELEASES_VERSION=$1 # 安装的版本

# 环境变量
TEMPDIR=`mktemp -d` #缓存目录
SERVICE_NAME="account" # 服务的名字
PACKAGE_NAME="linux-amd64.tar.gz" # 指定包名
SERVICE_PATH="/etc/systemd/system" # uint service 安装目录
INSTALL_PATH="/usr/local/account" # 程序的安装目录

ready() {
  # 准备安装目录
  if [ ! -d $INSTALL_PATH ]; then
    mkdir -p $INSTALL_PATH ;
    if [ $? != 0 ]; then
      echo "创建安装目录失败"
      exit 1
    fi
  fi
	# 下载安装包文件
	echo "wget https://github.com/grpcbrick/account/releases/download/$RELEASES_VERSION/$PACKAGE_NAME to $TEMPDIR/$PACKAGE_NAME ;"
	wget -O $TEMPDIR/$PACKAGE_NAME https://github.com/grpcbrick/account/releases/download/$RELEASES_VERSION/$PACKAGE_NAME ;
	if [ $? != 0 ]; then
    echo "安装包下载失败"
    exit 1
	fi
	echo "安装包下载完成"
}


# 安装程序
install() {
  # 解压安装包
  echo "解压安装包"
  tar -zxvf $TEMPDIR/$PACKAGE_NAME -C $INSTALL_PATH ;
  if [ $? != 0 ]; then
	rm -rf $TEMPDIR ;
    echo "安装包解压失败"
    exit 1
  fi

  # 清理包
  echo "清理安装包"
  rm -rf $TEMPDIR ;
  if [ $? != 0 ]; then
    echo "安装包清理失败"
    exit 1
  fi
}

# 安装服务文件
installUnitService() {
	echo "安装 UnitService"
  echo "[Unit]"                                                      > $SERVICE_PATH/$SERVICE_NAME.service ;
  echo "Description=$SERVICE_NAME service"                          >> $SERVICE_PATH/$SERVICE_NAME.service ;
  echo "[Service]"                                                  >> $SERVICE_PATH/$SERVICE_NAME.service ;
  echo "Type=simple"                                                >> $SERVICE_PATH/$SERVICE_NAME.service ;
  echo "Restart=always"                                             >> $SERVICE_PATH/$SERVICE_NAME.service ;
  echo "ExecStart=$INSTALL_PATH/main"                               >> $SERVICE_PATH/$SERVICE_NAME.service ;
  echo "EnvironmentFile=$INSTALL_PATH/.env"                         >> $SERVICE_PATH/$SERVICE_NAME.service ;
  echo "[Install]"                                                  >> $SERVICE_PATH/$SERVICE_NAME.service ;
  echo "WantedBy=multi-user.target"                                 >> $SERVICE_PATH/$SERVICE_NAME.service ;

  systemctl daemon-reload ;
  if [ $? != 0 ]; then
    echo "服务安装失败"
    exit 1
  fi
}

startUnitService() {
  systemctl restart $SERVICE_NAME.service
  if [ $? != 0 ]; then
    echo "服务启动失败"
    exit 1
  fi
}

# 执行
ready ;
install ;
installUnitService ;
startUnitService
