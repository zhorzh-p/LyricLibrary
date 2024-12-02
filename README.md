Реализация онлайн библиотеки песен

## swagger
Для генерации swagger документации api используется библиотеки "github.com/swaggo/gin-swagger" и "github.com/swaggo/files".
Для генерации файлов документации необходимо добавить соответствующие комментарии и выполнить:
```bash
swag init -g handler.go --dir ./cmd/server/api/
```

## Тесты
### Mockery
В проекте используется `mockery` для генерации моков интерфейсов. Это упрощает тестирование.

Установите `mockery` через Go CLI:
```bash
go install github.com/vektra/mockery/v2@latest
```

Конфигурация расположена в файле [\.mockery.yml.](.mockery.yml.). Для генерации моков необходимо выполнить в корне проекта

```bash
mockery
```
