#!/bin/bash

SERVER="High-Performance-Online-Bookstore"
BASE_DIR=$PWD
INTERVAL=2

# 命令行参数，需要手动指定
ARGS=""

function start()
{
	if [ "$(pgrep $SERVER -u $UID)" != "" ];then
		echo "$SERVER already running"
		exit 1
	fi

	nohup "$BASE_DIR"/$SERVER "$ARGS" &>/dev/null &

	echo "running in background, please check ./blog/server.log for running log"

	echo "Waiting for server status check..." &&  sleep $INTERVAL

	# check status
	if [ "$(pgrep $SERVER -u $UID)" == "" ];then
		echo "$SERVER start failed"
		exit 1
	fi
	echo "$SERVER start successfully"
}

function status()
{
	if [ "$(pgrep $SERVER -u $UID)" != "" ];then
		echo $SERVER is running
	else
		echo $SERVER is not running
	fi
}

function stop()
{
	if [ "$(pgrep $SERVER -u $UID)" != "" ];then
		kill -9 "$(pgrep $SERVER -u $UID)"
	fi

	echo "Waiting for server status check..." &&  sleep $INTERVAL

	if [ "$(pgrep $SERVER -u $UID)" != "" ];then
		echo "$SERVER stop failed"
		exit 1
	fi

	echo "$SERVER stopped successfully"
}

function version()
{
  ARGS="-v"
  "$BASE_DIR"/$SERVER $ARGS
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
