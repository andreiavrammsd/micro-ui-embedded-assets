#!/usr/bin/env bash

echo
read -p "Confirm deploy (y/n): " CONTINUE
echo

if [ ! "$CONTINUE" = "y" ]; then
  exit 1
fi

LOCAL_PATH="/tmp/micro-ui"
REMOTE_ADDR="user@host"
REMOTE_TMP_PATH="/tmp/micro-ui-tmp"
REMOTE_PATH="~/micro-ui"

go test
if [[ "${PIPESTATUS[0]}" != "0" ]]; then
    exit
fi
echo

go-bindata template/... && go build -o $LOCAL_PATH

scp $LOCAL_PATH $REMOTE_ADDR:$REMOTE_TMP_PATH
rm $LOCAL_PATH

CMD=""

PID=`ssh $REMOTE_ADDR "screen -ls" | awk '/\.micro-ui\t/ {print $1+0}'`
if [[ ! -z "$PID" ]]; then
    CMD="kill $PID &&"
fi

CMD="$CMD cp $REMOTE_TMP_PATH $REMOTE_PATH && rm $REMOTE_TMP_PATH"
CMD="$CMD && screen -AdmS micro-ui sh -c 'while true; do ~/micro-ui; sleep 5; done'"

ssh $REMOTE_ADDR "$CMD"

echo
