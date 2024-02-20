# CalculationService
# tg для связи: @leanq_ha

## Как запустить
1. Установить Docker Desktop
2. ```
   docker-compose up --build
   ```
3. Помолиться
4. Сервер запускается по адресу http://localhost:8080/static/calculator

##### Если программа не соберется  
В папке Если не запустится лежит версия, где все функции в одном файле, для ее запуска нужно только прописать  
```go run CalculatorRabotaet.go```

## Схема проекта
![image](https://github.com/1eanq/CalculationService/assets/153944563/588278d4-f409-45cf-bc99-b2a25e3e7402)

## Доступные операции:
  "+", "-", "*", "/", скобки

## Пример выражения:
```
((1+3)*6-0.654*(1/26)+2)*0.5
```

## Пример запроса:

### cURL:  
```
curl --location --request POST 'http://localhost:8080/static/calculatot' \
--header 'Expression: ((1+3)*6-0.654*(1/26)+2)*0.5' \
--data ''
```

### HTML:  
```
POST /static/calculatot HTTP/1.1
Host: localhost:8080
Expression: ((1+3)*6-0.654*(1/26)+2)*0.5
```
