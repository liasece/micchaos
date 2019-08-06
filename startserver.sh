#!/bin/bash
nohup ../bin/chaos -d true -l ~/logs/ -p gate001,gate002,player001 > /dev/null 2>&1 &
