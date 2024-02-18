# Используем образ Golang для сборки приложения
FROM golang:1.17-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файл go.mod и go.sum
COPY go.mod ./
COPY go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем все остальные файлы проекта
COPY . .

# Собираем проект
RUN go build -o CalculationService cmd/main.go


# Используем минимальный образ alpine в качестве базового образа для запуска приложения
FROM alpine:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем исполняемый файл из предыдущего этапа
COPY --from=builder /app/CalculationService .

# Запускаем наше приложение при старте контейнера
CMD ["./CalculationService"]
