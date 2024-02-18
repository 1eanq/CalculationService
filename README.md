# CalculationService
# tg для связи: @leanq_ha

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
