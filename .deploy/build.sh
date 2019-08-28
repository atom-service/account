#! /bin/bash

rootPath=`pwd`
outPath="./build"
platforms=('linux-386' 'linux-amd64' 'linux-arm' 'darwin-386' 'darwin-amd64' 'freebsd-386' 'freebsd-amd64' 'windows-386' 'windows-amd64')

installVendor() {
  go get -v -u -d google.golang.org/grpc
  go get -v -u -d github.com/jmoiron/sqlx
  go get -v -u -d github.com/go-sql-driver/mysql
  go get -v -u -d github.com/yinxulai/goutils/...
  go get -v -u -d github.com/golang/protobuf/proto
}

build() {
  cd $rootPath
  mkdir $outPath
  for platform in ${platforms[@]}; do

      tempDir=`mktemp -d`
      echo "build start cgoEnabled, platform $platform."
      GOOS=${platform%%-*} GOARCH=${platform##*-} go build -o $tempDir/main ./main.go ;

      if [ $? == 0 ]; then
        tar -zcvf ${outPath}/${platform}.tar.gz -C $tempDir ./main;
        echo "build done"
      else
        echo "build error"
      fi
      rm -rf $tempDir
  done
}

if [ $1 == 'installVendor' ];then
  installVendor
fi

if [ $1 == 'build' ];then
  build
fi
