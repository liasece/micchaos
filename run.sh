#!/bin/bash
rm -rf ~/logs/*
nohup ../bin/chaos >/dev/null 2>&1 &
