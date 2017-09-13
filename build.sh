#!/bin/bash
#set -x
peername='imagesvr'
nowtime=$(date +%Y%m%d%H%M)
appname=$peername.$nowtime
dir=$(pwd)


###  initalize #####
echo "initalize ..."
mkdir -p  output/bin
mkdir -p  output/conf
mkdir -p  output/logs

###  build      ####
echo "start build ..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

### copy files  ####
echo "copy to destination dir"
mv main                  output/bin/$appname
cp ./conf/config.json ./output/conf/config.json

### shell script ####
shell_header='#!bin/bash'
load_sh='nohup ./bin/'$appname' >./logs/'$appname.log' 2>./logs/error.log  &'

echo $shell_header  >  output/load.sh
echo $load_sh >> output/load.sh

args='$2'
load="ps -ef | grep $peername | grep -v grep | awk '{print $args}'"
pscount="ps -ef | grep $peername | grep -v grep | wc -l"
gpid='pid=$('$load')'

echo $shell_header  >  output/kill.sh
echo 'pid=$('$load')' >> output/kill.sh
echo 'count=$('$pscount')' >> output/kill.sh
echo 'if [ $count -ne 0 ];then kill $pid' >> output/kill.sh
echo "fi" >> output/kill.sh
echo 'echo "count:"$count' >> output/kill.sh

### tar ############
echo "tar ..."
cd output
tar -czf $peername.tar.gz ./bin ./conf ./logs ./load.sh ./kill.sh
mv ./$peername.tar.gz $dir/
rm -r $dir/output/

### done ############
echo "done,done,done"
