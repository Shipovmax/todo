# Task #4.5 — Todo Frontend

## Goal

Write a minimalist web interface for the todo-api from project #4. A single `index.html` file, no frameworks or bundlers — plain HTML + CSS + Vanilla JS with `fetch()`. The main learning goal is to understand how a frontend talks to a REST API: all CRUD operations via fetch, asynchrony via async/await, with the server as the source of truth.

---

## Acceptance Criteria

- [ ] Open `index.html` in a browser — all tasks from the API are visible (`GET /todos`)
- [ ] Type text + Enter -> the task appears in the list (`POST /todos`)
- [ ] Click the checkbox -> the task is marked done both visually and on the server (`PATCH /todos/{id}`)
- [ ] Hover over a task -> a delete button appears; click it -> the task disappears (`DELETE /todos/{id}`)
- [ ] Empty input -> the request is not sent, the field is highlighted
- [ ] API unavailable -> an error message is shown to the user
- [ ] After each action the list is reloaded from the server
- [ ] Completed tasks are visually distinct (strikethrough text, muted color)
- [ ] No external JS dependencies — only one font via Google Fonts

---

## Technical Requirements

### Mandatory

| Requirement | Details |
|---|---|
| Structure | a single `index.html` file with `<style>` and `<script>` inline |
| HTTP requests | `fetch()` with `async/await` for all operations |
| Request method | explicitly specify `method`, `headers`, `body` in fetch |
| Network errors | `try/catch` around every fetch, message shown to the user |
| Refresh after action | call `loadTodos()` after every POST/PATCH/DELETE |
| Empty list | show placeholder text instead of a blank screen when there are no tasks |
| Content-Type | `"Content-Type": "application/json"` in headers for POST/PATCH |
| API URL | a `const API = "http://localhost:8080"` constant at the top of the script |

### Forbidden

- Third-party JS libraries (React, Vue, jQuery, axios)
- npm, node_modules, bundlers (webpack, vite)
- Mutating the local task array instead of reloading from the server
- Inline styles via `element.style` — CSS classes only

---

## Topics and Skills Being Practiced

> This is not just a list — it is a checklist of things the implementation **must** use.

- **`fetch()` with configuration** — `{method, headers, body: JSON.stringify(...)}` for POST/PATCH/DELETE
- **`async/await`** — all API functions are async: `async function loadTodos()`
- **`try/catch`** — wraps every fetch call to handle network errors
- **`response.json()`** — parsing the JSON response from the API
- **`response.ok`** — checking for a 2xx status before processing the response
- **DOM API** — `document.createElement`, `element.appendChild`, `element.innerHTML`
- **Event listeners** — `addEventListener('click', ...)`, `addEventListener('keydown', ...)`
- **CSS variables** — `--color-accent`, `--color-text` for a consistent color scheme
- **CSS transitions** — `transition: opacity 0.2s` for the delete button appearing on hover

---

## File Structure

```
todo-frontend/
├── index.html    # HTML + <style> + <script> — everything in one file
└── README.md
```

---

## Architecture Hints

```js
// Constants at the top of the script
const API = 'http://localhost:8080'

// Functions for each API operation
async function loadTodos()           // GET /todos -> render the list
async function createTodo(title)     // POST /todos
async function toggleTodo(id, done)  // PATCH /todos/{id}
async function deleteTodo(id)        // DELETE /todos/{id}

// Render — rebuilds the list from scratch
function renderTodos(todos)          // clears the container, creates cards

// Error handling — a single function
function showError(message)          // shows a banner with the error text
```

> Every API function calls `loadTodos()` after a successful operation — the list is always in sync with the server.

---

## Definition of Done

1. All acceptance criteria are met
2. Code is pushed to GitHub in the `todo-frontend` repository
3. README.md in the repository follows the project template
4. You can explain every line of code out loud without looking it up

---

## Next Step After Submission

After review we move on to **Task #5 — Concurrent Worker**: goroutines, channels, `sync.WaitGroup`, `context.WithCancel` — the fan-out/fan-in pattern, the most important Go topic for BigTech interviews.
