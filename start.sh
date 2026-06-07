#!/bin/bash
PORT=5175
PID=$(lsof -ti :$PORT)
if [ -n "$PID" ]; then
  kill $PID
fi
cd "$(dirname "$0")/paraflow/Screen & Prototype"
python3 -m http.server $PORT