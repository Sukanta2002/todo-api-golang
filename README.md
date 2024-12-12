# Todo API in Go

This is a simple Todo API built in Go. It allows you to create, get all, get by id and delete todos.

## Endpoints

### POST /todo

Create a new todo.

Example Request Body:

```json language=JSON
{
  "title": "Buy groceries",
  "description": "Milk, eggs, bread, and vegetables"
}
```

### GET /todos

Get all todos.

### GET /todo/{id}

Get a todo by id.

### DELETE /todo/{id}

Delete a todo by id.
