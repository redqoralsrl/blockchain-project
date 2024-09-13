#!/bin/bash

sudo kill $(ps aux | grep 'active' | awk '{print $2}')

PIDS=$(pgrep -f './main')

if [ -z "$PIDS" ]; then
  echo "No './main' process running."
else
  echo "Killing './main' processes..."
  # './main' 프로세스를 종료
  pkill -f './main'

  if [ $? -eq 0 ]; then
    echo "'./main' processes killed successfully."
  else
    echo "Failed to kill './main' processes."
  fi
fi

tmux kill-session -t gmwallet 2>/dev/null

rm -rf /home/ubuntu/gmwallet-connect-go/*
