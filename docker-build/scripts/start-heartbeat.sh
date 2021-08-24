#!/bin/bash

nohup $HEARTBEAT_HOME/heartbeat -e -c $HEARTBEAT_HOME/heartbeat.yml >$HEARTBEAT_HOME/heartbeat.stdout.log 2>$HEARTBEAT_HOME/heartbeat.stderr.log &