#!/bin/sh

cat /tmp/data.txt | while read line; do curl -X POST -kq https://numa-realtime-analytics-pipeline-in:8443/vertices/in -d "$line"; echo $line; sleep 3; done
