#!/bin/bash
#
# restart --- start or retart the servers
#

tmp=$IFS
IFS='
'

for var  in $(ps -u $(basename $HOME) | grep "src") 
do
	pid=$(echo $var | cut -c1-5)
	pname=$(echo $var | cut -c25-)	

	if  kill -2 $pid
	then
		echo "$pname stoped"
	else
		echo "$pname can't be stoped"
	fi
done

cond=$(ps -u $(basename $HOME) | grep "src" | wc -l)
while [ $cond -gt 0 ]
do
	sleep 1
	cond=$(ps -u $(basename $HOME) | grep "src" | wc -l)
	echo "ServerNum:$cond"
done

