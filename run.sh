#!/bin/bash
rm -rf ~/logs/*
nohup ../bin/src > /dev/null 2>&1 &
