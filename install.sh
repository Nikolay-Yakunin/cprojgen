#!/bin/bash
# Скрипт установки для Linux/Unix
# Этот скрипт компилирует программу, создаёт папку build и копирует бинарник в /usr/local/bin

set -e

echo "Сборка cprojgen..."
go build -o cprojgen main.go

echo "Создание папки build..."
mkdir -p build

INSTALL_DIR="/usr/local/bin"
echo "Установка cprojgen в ${INSTALL_DIR}..."

if [ "$(id -u)" -ne 0 ]; then
    echo "Необходимы права суперпользователя. Используем sudo..."
    sudo cp cprojgen "${INSTALL_DIR}/"
else
    cp cprojgen "${INSTALL_DIR}/"
fi

echo "Установка завершена!"
