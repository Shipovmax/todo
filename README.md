# todo-frontend — Web Interface for todo-api

> A minimalist frontend for todo-api in macOS style. Bonus project #4.5 — turns a REST API into a working product.

---

## For Recruiters

### What and Why

This is a bonus project between #4 and #5 on the roadmap. The goal is to demonstrate that a backend developer understands the full stack: not just setting up an API, but connecting a UI to it. The frontend is written without frameworks or bundlers — just HTML, CSS, and JavaScript — to show understanding of the fetch API, async programming, and DOM manipulation.

The project is a real tool used daily: a todo manager with a minimalist macOS-style design, connected to a Go backend from project #4.

### What This Project Demonstrates

| Skill | Implementation |
|---|---|
| REST API integration | `fetch()` for all CRUD operations: POST, GET, PATCH, DELETE |
| Async JavaScript | `async/await` with error handling via `try/catch` |
| DOM manipulation | Dynamic task list rendering without frameworks |
| Network error handling | User-facing message when the API is unavailable |
| UI/UX | macOS minimalism: Inter font, smooth transitions, hover effects |
| Zero dependencies | A single `index.html` file, no npm, no bundlers |

### Stack

- **Frontend:** HTML5 + CSS3 + Vanilla JavaScript (ES2022)
- **Font:** Inter via Google Fonts
- **Backend:** todo-api (Go 1.22+, `net/http`)
- **Dependencies:** none
- **Launch:** open `index.html` in a browser or via a Go file server

---

## For Developers

### Architectural Decisions

#### Why a single file instead of separate HTML + CSS + JS?

For a project without a bundler, separate files cause CORS issues when opened via `file://`. A single `index.html` with inline `<style>` and `<script>` opens directly in the browser with no server required — zero setup friction.

#### Why reload the list from the server after each action instead of updating local state?

```js
// Local update — can desync from server on error
todos = todos.filter(t => t.id !== id)
render(todos)

// Reload from server — truth always lives on the backend
await deleteTodo(id)
await loadTodos() // GET /todos
```

The server is the source of truth. This eliminates an entire class of state-sync bugs and accurately reflects how production SPAs with REST APIs work.

#### Why `async/await` instead of `.then().catch()`?

`async/await` reads like synchronous code — easier to follow the flow during review. `try/catch` explicitly marks error-handling boundaries. Functionally identical to promise chains.

#### Why show the delete button only on hover?

CSS `opacity: 0` → `opacity: 1` on card `:hover`. This is a macOS convention: destructive actions are hidden until the user shows explicit intent. The interface stays clean while scanning the list.

### Structure

```
todo-frontend/
├── index.html    # everything in one file: HTML markup + <style> + <script>
└── README.md
```

### Installation and Launch

```bash
# 1. Start the backend
cd ../todo-api
go run .

# 2. Open the frontend
cd ../todo-frontend
open index.html   # macOS
```

Or via a Go file server (if http:// is needed):

```bash
cd todo-frontend
python3 -m http.server 3000
# open http://localhost:3000
```

### Usage

- **Add a task** — type text in the input at the top, press Enter or the + button
- **Mark as done** — click the checkbox to the left of the task
- **Delete** — hover over the card, click × on the right

### Error Handling

```
API unavailable (backend not running)
→ Red banner at the top: "Could not connect to server"

Empty title when adding
→ Field is highlighted, request is not sent

Server returned an error (4xx/5xx)
→ Console + banner with error text from {"error": "..."}
```
