# TaskFlow TODO Application

TaskFlow is a planned TODO management application with a Go REST API and
MySQL persistence. Its end-user workflow is organized around a Kanban
board with To Do, In Progress, and Done columns.

> **Current status:** The Go application and frontend are not implemented.
> This repository contains a black-box API test scaffold, but no Go module,
> database migrations, Docker configuration, or running service. The API
> and Kanban behavior described here are requirements, not shipped features.

## Table of contents

-   [Overview](#overview)
-   [Product scope](#product-scope)
-   [Technology direction](#technology-direction)
-   [Repository contents](#repository-contents)
-   [API design](#api-design)
-   [Data model](#data-model)
-   [Kanban workflow](#kanban-workflow)
-   [Development status](#development-status)
-   [Roadmap](#roadmap)

## Overview

## Product scope

The initial product is a personal task-management application. Users
should be able to create, view, update, search, and delete their tasks,
and move tasks through a three-stage workflow. Account sharing, team
boards, and custom workflow columns are not part of the initial scope.

The intended API capabilities are TODO CRUD, bounded list pagination,
input validation, consistent errors, and health/readiness checks. These
are design goals until implemented and tested.

## Technology direction

-   **Go** for the REST API.
-   **MySQL** for persistent task data.
-   **OpenAPI** for the API contract.
-   **DBML** for the database design.
-   **Docker** is a possible local-development tool, but no Docker
    configuration is currently present.

No Redis, OpenTelemetry, Zipkin, Prometheus, LocalStack, or load-testing
tool is currently configured in this repository.

## Repository contents

| File | Purpose |
| --- | --- |
| `README.md` | Project overview, current design, and status |
| `user_strories.md` | End-user stories and acceptance criteria, including Kanban |
| `swagger.yaml` | Intended API specification; validate and maintain as OpenAPI YAML |
| `db_schema.dbml` | MySQL-oriented database design |
| `tests/todo_api_test.go` | Black-box test scaffold for the TODO API workflow |

There is currently no application source tree or executable setup
command. The API test scaffold requires a running service and
`TODO_API_BASE_URL`; it skips when that variable is not set. Add build,
run, and full test instructions when the Go service and its dependencies
are added.

## API design

The planned API is a versioned REST API. The current API design describes
these routes; they are not live endpoints yet:

| Method | Endpoint | Purpose |
| --- | --- | --- |
| `POST` | `/v1/todos` | Create a task |
| `GET` | `/v1/todos` | List tasks with cursor pagination |
| `GET` | `/v1/todos/{id}` | Retrieve a task |
| `PATCH` | `/v1/todos/{id}` | Update a task name or status |
| `DELETE` | `/v1/todos/{id}` | Delete a task |
| `GET` | `/health` | Check whether the API process is running |
| `GET` | `/ready` | Check whether the API can serve requests |

Authentication requirements have not been defined. The API contract and
database design must be reconciled before implementation: they currently
disagree on task identifiers, task-name fields, and status representation.
See [Data model](#data-model).

## Data model

The DBML design currently defines `users` and `todos`. A TODO has a
numeric ID, `user_id`, `title`, optional description and due date, a
`completed` boolean, and timestamps. The API draft instead describes a
UUID `id`, `name`, and a three-value status (`pending`, `in_progress`,
`completed`). These differences need an explicit design decision before
database migrations or API implementation.

For the Kanban workflow, the intended columns correspond to `pending`
(To Do), `in_progress` (In Progress), and `completed` (Done). The current
DBML boolean cannot represent all three stages, so it must be updated if
the three-column workflow is selected. Keep the DBML, OpenAPI document,
Go model, and migrations consistent.

## Kanban workflow

The planned end-user experience is described in
[`user_strories.md`](user_strories.md). Its MVP stories cover:

-   Viewing tasks in To Do, In Progress, and Done columns.
-   Creating a task, optionally choosing its initial status.
-   Moving tasks between columns and retaining the change after reload.
-   Editing, viewing, searching, and deleting tasks.
-   Showing clear empty, loading, and error states.

The current repository has no frontend or board implementation. The API
and data model must support loading tasks and persisting status changes
for the board to work.

## Development status

There are no application build or run commands because this repository
does not contain Go source code, `go.mod`, SQL migrations, or Docker
configuration. The API test scaffold can be compiled independently and
will exercise the TODO workflow when pointed at a running service.
Expand coverage as implementation begins to include input validation,
pagination, database failures, and the Kanban board's empty and error
states.

## Roadmap

-   [ ] Decide and reconcile task IDs, names, and status representation across the API and DBML.
-   [ ] Update the API and DBML to support the three Kanban workflow stages.
-   [ ] Implement the Go REST API and MySQL persistence.
-   [ ] Add the Kanban board and core task-management interactions.
-   [ ] Add migrations, local development configuration, and automated tests.
-   [ ] Implement cursor pagination and health/readiness endpoints.
-   [ ] Add operational features only when needed and verify them with tests.
