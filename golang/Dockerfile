# Шаг 1: Сборка бинарного файла
FROM golang:1.23 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum и устанавливаем зависимости
COPY go.mod ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка бинарного файла
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service .

# Шаг 2: Создание минимального образа для запуска
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем собранный бинарный файл из предыдущего шага
COPY --from=builder /app/service .

# Открываем порт, на котором работает сервис
EXPOSE 3001

# Запуск приложения
CMD ["./service"]
