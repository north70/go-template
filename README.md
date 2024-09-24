# go-template

## Описание

`go-template` - это шаблонный проект на Go, который включает в себя примеры использования gRPC, генерации моков, миграций базы данных и тестирования. Этот проект предназначен для быстрого старта разработки микросервисов на Go.

## Как развернуть локально проект

### Предварительные требования

- [Go](https://golang.org/doc/install) 1.22 или выше
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)

### Установка зависимостей

1. Клонируйте репозиторий:

    ```sh
    git clone https://github.com/yourusername/go-template.git
    cd go-template
    ```

2. Установите необходимые инструменты:

    ```sh
    make install-tools
    ```

3. Получите зависимости Go:

    ```sh
    go mod tidy
    ```

### Запуск проекта

#### С использованием Docker

1. Проверьте наличие файла `app.local.yaml` в директории `configs/`. Если файла нет, создайте его, используя файл `app.example.yaml` в качестве шаблона.

2. Проверьте корректность переменной `POSTGRES_DSN` в Makefile.

3. Запустите Docker-контейнеры:

    ```sh
    make start
    ```

4. Выполните миграции базы данных:

    ```sh
    make migrate-up
    ```

5. Запустите проект:

    ```sh
    go run main.go --config configs/app.local.yaml
    ```

#### Дебаг

Добавьте конфигурацию запуска в IDE
```
package path = github.com/north70/go-template/cmd/go-template
program arguments = -config=configs/app.local.yml
```
## Работа с Makefile

`Makefile` содержит несколько полезных команд для автоматизации задач в проекте. Вот основные из них:

- `make install-tools`: Устанавливает необходимые инструменты для разработки.
- `make get-dependencies`: Клонирует необходимые зависимости, такие как `googleapis` и `grpc-gateway`.
- `make generate-proto`: Генерирует gRPC и OpenAPI файлы из прото-файлов.
- `make lint`: Запускает линтер `golangci-lint` для проверки кода.
- `make start`: Запускает Docker-контейнеры.
- `make stop`: Останавливает Docker-контейнеры.
- `make migrate-up`: Выполняет миграции базы данных вверх.
- `make migrate-down`: Откатывает миграции базы данных вниз.
- `make migrate-status`: Проверяет статус миграций базы данных.
- `make test`: Запускает все тесты в проекте.

## Структура проекта

```plaintext
.
├── bin/                        # Директория для бинарных файлов
├── cmd/                        # Точка входа в приложение
├── configs/                    # Конфигурационные файлы
├── internal/                   # Внутренние пакеты проекта
│   ├── cache/                  # Логика кэширования
│   ├── domain/                 # Определения доменных сущностей
│   ├── gateway/                # Внешние сервисы и API
│   ├── interceptor/            # Перехватчики для gRPC
│   ├── pb/                     # Сгенерированные прото-файлы
│   ├── repository/             # Логика работы с базой данных
│   ├── service/                # Бизнес-логика
│   ├── store/                  # Логика подключения к базам данных
├── migrations/                 # Миграции базы данных
├── proto/                      # Прото-файлы для gRPC
├── .golangci.yml               # Конфигурация для golangci-lint
├── docker-compose.yml          # Конфигурация Docker Compose
├── Makefile                    # Makefile для автоматизации задач
└── README.md                   # Описание проекта
```

## Лицензия

Этот проект лицензирован под лицензией MIT. Подробности см. в файле LICENSE.
