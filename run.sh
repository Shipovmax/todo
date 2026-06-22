#!/bin/bash
set -e

# Переходим в папку проекта (работает из любой директории)
cd "$(dirname "$0")"

# Проверяем что порт свободен
if lsof -ti:8080 > /dev/null 2>&1; then
  echo "Порт 8080 уже занят. Останавливаю старый процесс..."
  kill $(lsof -ti:8080) 2>/dev/null || true
  sleep 0.5
fi

echo "Запускаю сервер..."
go run . &
SERVER_PID=$!

# Ждём пока сервер поднимется (до 5 секунд)
for i in $(seq 1 20); do
  if curl -s http://localhost:8080/todos > /dev/null 2>&1; then
    break
  fi
  sleep 0.25
done

echo "Открываю браузер → http://localhost:8080"
open http://localhost:8080

echo "Сервер работает (PID $SERVER_PID). Нажми Ctrl+C для остановки."

# Останавливаем сервер при выходе
trap "kill $SERVER_PID 2>/dev/null; echo 'Сервер остановлен.'" EXIT
wait $SERVER_PID
