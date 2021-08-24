#!/bin/bash

nohup $HEARTBEAT_HOME/heartbeat -e -c $HEARTBEAT_HOME/heartbeat.yml >$PACKETBEAT_HOME/heartbeat.stdout.log 2>$PACKETBEAT_HOME/heartbeat.stderr.log &