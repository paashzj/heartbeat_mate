#!/bin/bash

mkdir $HEARTBEAT_HOME/logs
nohup $HEARTBEAT_HOME/mate/heartbeat_mate >>$HEARTBEAT_HOME/logs/heartbeat_mate.stdout.log 2>>$HEARTBEAT_HOME/logs/heartbeat_mate.stderr.log

