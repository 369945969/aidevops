#!/bin/bash
PORT=8080
PID=$(lsof -ti :$PORT 2>/dev/null)

if [ -n "$PID" ]; then
  echo "Port $PORT is already in use (PID: $PID), killing process..."
  kill -9 $PID
  sleep 1
  echo "Process killed."
fi

cd "$(dirname "$0")/golang"
go run .