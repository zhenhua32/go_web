#!/bin/bash

SERVER="web"
BASE_DIR=$PWD
INTERVAL=2

# 命令行参数，需要手动指定, 这是在 docker 容器中运行的参数
ARGS="-c $BASE_DIR/conf/config_docker.yaml"

function start()
{
	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		echo "$SERVER already running"
		exit 1
	fi

	nohup $BASE_DIR/$SERVER $ARGS >/dev/null 2>&1 &
	echo "sleeping..." &&  sleep $INTERVAL

	# check status
	if [ "`pgrep $SERVER -u $UID`" == "" ];then
		echo "$SERVER start failed"
		exit 1
  else
    echo "start success"
	fi
}

function status() 
{
	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		echo $SERVER is running
	else
		echo $SERVER is not running
	fi
}

function stop() 
{
	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		kill `pgrep $SERVER -u $UID`
	fi

	echo "sleeping..." &&  sleep $INTERVAL

	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		echo "$SERVER stop failed"
		exit 1
  else
    echo "stop success"
	fi
}

function version()
{
  $BASE_DIR/$SERVER $ARGS version
}

case "$1" in
	'start')
	start
	;;  
	'stop')
	stop
	;;  
	'status')
	status
	;;  
	'restart')
	stop && start
	;;  
  'version')
  version
  ;;
	*)  
	echo "usage: $0 {start|stop|restart|status|version}"
	exit 1
	;;  
esac