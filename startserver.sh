#!/bin/bash
nohup ../bin/chaos -d true -l ~/logs/ -p gate001,player001,login001 > /dev/null 2>&1 &
