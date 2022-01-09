#!/bin/bash

mkdir $HEARTBEAT_HOME/logs
nohup $HEARTBEAT_HOME/heartbeat -e -c $HEARTBEAT_HOME/heartbeat.yml >>$HEARTBEAT_HOME/logs/heartbeat.stdout.log 2>>$HEARTBEAT_HOME/logs/heartbeat.stderr.log &

