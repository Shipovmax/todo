# Task #4.5 — Todo Frontend

## Цель

Написать минималистичный веб-интерфейс для todo-api из проекта #4. Один файл `index.html` без фреймворков и сборщиков — чистый HTML + CSS + Vanilla JS с `fetch()`. Главная учебная цель — понять как фронтенд взаимодействует с REST API: все CRUD операции через fetch, асинхронность через async/await, источник правды на сервере.

---

## Acceptance Criteria

- [ ] Открыть `index.html` в браузере — видны все задачи из API (`GET /todos`)
- [ ] Ввести текст + Enter → задача появляется в списке (`POST /todos`)
- [ ] Нажать чекбокс → задача отмечается выполненной визуально и на сервере (`PATCH /todos/{id}`)
- [ ] Навести на задачу → появляется кнопка удаления; нажать → задача исчезает (`DELETE /todos/{id}`)
- [ ] Пустой input → запрос не отправляется, поле подсвечивается
- [ ] API недоступен → показывается сообщение об ошибке пользователю
- [ ] После каждого действия список перезагружается с сервера
- [ ] Выполненные задачи визуально отличаются (зачёркнутый текст, приглушённый цвет)
- [ ] Никаких внешних JS-зависимостей — только один шрифт через Google Fonts

---

## Технические требования

### Обязательно

| Требование | Детали |
|---|---|
| Структура | один файл `index.html` с `<style>` и `<script>` внутри |
| HTTP-запросы | `fetch()` с `async/await` для всех операций |
| Метод запроса | явно указывать `method`, `headers`, `body` в fetch |
| Ошибки сети | `try/catch` вокруг каждого fetch, сообщение пользователю |
| Обновление после действия | после POST/PATCH/DELETE вызывать `loadTodos()` |
| Пустой список | если задач нет — показывать placeholder текст, не пустой экран |
| Content-Type | `"Content-Type": "application/json"` в headers для POST/PATCH |
| API URL | константа `const API = "http://localhost:8080"` вверху скрипта |

### Запрещено

- Сторонние JS-библиотеки (React, Vue, jQuery, axios)
- npm, node_modules, сборщики (webpack, vite)
- Локальное изменение массива задач вместо перезагрузки с сервера
- Inline стили через `element.style` — только CSS классы

---

## Темы и навыки которые прокачиваешь

> Это не просто список — это checklist того, что **обязан использовать** в реализации.

- **`fetch()` с конфигурацией** — `{method, headers, body: JSON.stringify(...)}` для POST/PATCH/DELETE
- **`async/await`** — все функции работы с API асинхронные: `async function loadTodos()`
- **`try/catch`** — обёртка вокруг каждого fetch для обработки сетевых ошибок
- **`response.json()`** — парсинг JSON ответа от API
- **`response.ok`** — проверка что статус 2xx перед обработкой ответа
- **DOM API** — `document.createElement`, `element.appendChild`, `element.innerHTML`
- **Event listeners** — `addEventListener('click', ...)`, `addEventListener('keydown', ...)`
- **CSS переменные** — `--color-accent`, `--color-text` для единой цветовой схемы
- **CSS transitions** — `transition: opacity 0.2s` для появления кнопки удаления при hover

---

## Структура файлов

```
todo-frontend/
├── index.html    # HTML + <style> + <script> — всё в одном файле
└── README.md
```

---

## Подсказки по архитектуре

```js
// Константы вверху скрипта
const API = 'http://localhost:8080'

// Функции для каждой операции с API
async function loadTodos()           // GET /todos → рендер списка
async function createTodo(title)     // POST /todos
async function toggleTodo(id, done)  // PATCH /todos/{id}
async function deleteTodo(id)        // DELETE /todos/{id}

// Рендер — пересобирает список с нуля
function renderTodos(todos)          // очищает контейнер, создаёт карточки

// Обработка ошибок — единая функция
function showError(message)          // показывает баннер с текстом ошибки
```

> Каждая API-функция после успешного выполнения вызывает `loadTodos()` — список всегда актуален относительно сервера.

---

## Definition of Done

1. Все Acceptance Criteria выполнены
2. Код запушен на GitHub в репозиторий `todo-frontend`
3. README.md в репозитории соответствует шаблону проекта
4. Ты можешь объяснить каждую строку кода вслух без подглядывания

---

## Следующий шаг после сдачи

После ревью переходим к **Task #5 — Конкурентный воркер**: goroutines, channels, `sync.WaitGroup`, `context.WithCancel` — fan-out/fan-in паттерн, самая важная тема Go для BigTech собеседований.
