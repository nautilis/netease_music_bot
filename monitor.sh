#!/bin/bash

# keep the pro alive
BASE_DIR="$(dirname $(readlink -f $0))"
PROG=${BASE_DIR##*/}
DINGTALK_API="https://api.telegram.org/bot5629309539:AAHX835Xk8obrPUzepa83c4sQ0wxEgMRVWE/sendMessage"
DINGTALK_CTYPE="Content-Type: application/json"
MACHINE=`hostname -s`

ulimit -c unlimited
ps -ef|grep ${PROG}|grep -v grep|grep -v monitor
DEAMON_NUM=`ps -ef|grep ${PROG}|grep -v grep|grep -v monitor|wc -l`

if [ $DEAMON_NUM -lt 1 ] ; then
  MSG="Try to Restart ${PROG} in ${MACHINE}"
  echo ${MSG}
  curl -s "${DINGTALK_API}" -H "${DINGTALK_CTYPE}" -d "{\"chat_id\":\"-1001877978872\",\"text\":\"${MSG}\"}"
  cd `dirname "$0"`
  nohup sh ./run.sh 2>&1 &
  sleep 1
  ps -ef|grep ${PROG}|grep -v grep|grep -v monitor
fi