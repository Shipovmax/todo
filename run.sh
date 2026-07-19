#!/bin/bash
set -e

# Move to the project directory (works from any working directory)
cd "$(dirname "$0")"

# Check that the port is free
if lsof -ti:8080 > /dev/null 2>&1; then
  echo "Port 8080 is already in use. Stopping the old process..."
  kill $(lsof -ti:8080) 2>/dev/null || true
  sleep 0.5
fi

echo "Starting server..."
go run . &
SERVER_PID=$!

# Wait for the server to come up (up to 5 seconds)
for i in $(seq 1 20); do
  if curl -s http://localhost:8080/todos > /dev/null 2>&1; then
    break
  fi
  sleep 0.25
done

echo "Opening browser -> http://localhost:8080"
open http://localhost:8080

echo "Server is running (PID $SERVER_PID). Press Ctrl+C to stop."

# Stop the server on exit
trap "kill $SERVER_PID 2>/dev/null; echo 'Server stopped.'" EXIT
wait $SERVER_PID
