# Effective Mobile — Subscription Service

REST-сервис для агрегации данных об онлайн подписках пользователей.

```bash
# Запуск
make infra
```

Swagger: http://localhost:8080/swagger/index.html


## Логи
Логи в JSON формате в `stdout`.

**В Docker:**
```bash
# Все логи
docker logs effective-mobile-app

# В реальном времени
docker logs -f effective-mobile-app

# PostgreSQL
docker logs effective-mobile-postgres
```

**Локально (`make run`):** логи в терминал.
