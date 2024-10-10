# CI/CD Go App API

This project is a simple RESTful API for managing tasks built with Go. It provides endpoints to create and retrieve tasks, utilizing the Gorilla Mux router for handling HTTP requests.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [API Endpoints](#api-endpoints)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- Create new tasks
- Retrieve a list of tasks

## Technologies

- Go (1.23.1)
- Gorilla Mux Router
- JSON for data interchange

## API Endpoints

### Get All Tasks

- **Endpoint**: `/tasks`
- **Method**: `GET`
- **Description**: Retrieve a list of all tasks.

**Response**:  
Returns a JSON array of tasks.

### Create a Task

- **Endpoint**: `/tasks`
- **Method**: `POST`
- **Description**: Create a new task.

**Request Body**:
```json
{
  "title": "Task Title",
  "description": "Task Description",
  "done": false
}
