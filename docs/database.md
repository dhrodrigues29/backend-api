# Database Design

## users

Stores application users.

| Column        | Type         | Constraints                            | Description               |
| ------------- | ------------ | -------------------------------------- | ------------------------- |
| id            | UUID         | Primary key, DEFAULT gen_random_uuid() | Unique user identifier    |
| name          | VARCHAR(120) | NOT NULL                               | User full name            |
| email         | VARCHAR(255) | NOT NULL, UNIQUE                       | User email used for login |
| password_hash | TEXT         | NOT NULL                               | Hashed password           |
| created_at    | TIMESTAMPTZ  | NOT NULL, DEFAULT now()                | Creation timestamp        |
| updated_at    | TIMESTAMPTZ  | NOT NULL, DEFAULT now()                | Last update timestamp     |

## projects

Stores projects owned by users.

| Column      | Type         | Constraints                               | Description               |
| ----------- | ------------ | ----------------------------------------- | ------------------------- |
| id          | UUID         | Primary key, DEFAULT gen_random_uuid()    | Unique project identifier |
| name        | VARCHAR(120) | NOT NULL                                  | Project name              |
| description | TEXT         | NOT NULL                                  | Project description       |
| owner_id    | UUID         | NOT NULL, FK users(id), ON DELETE CASCADE | User who owns the project |
| created_at  | TIMESTAMPTZ  | NOT NULL, DEFAULT now()                   | Creation timestamp        |
| updated_at  | TIMESTAMPTZ  | NOT NULL, DEFAULT now()                   | Last update timestamp     |

## tasks

Stores tasks inside projects.

| Column      | Type         | Constraints                                  | Description                |
| ----------- | ------------ | -------------------------------------------- | -------------------------- |
| id          | UUID         | Primary key, DEFAULT gen_random_uuid()       | Unique task identifier     |
| project_id  | UUID         | NOT NULL, FK projects(id), ON DELETE CASCADE | Project that owns the task |
| title       | VARCHAR(120) | NOT NULL, UNIQUE per project                 | Task title                 |
| description | TEXT         | Nullable                                     | Task description           |
| status      | VARCHAR(20)  | NOT NULL, DEFAULT 'todo'                     | Task status                |
| assignee_id | UUID         | Nullable, FK users(id), ON DELETE SET NULL   | User assigned to the task  |
| created_by  | UUID         | NOT NULL, FK users(id), ON DELETE CASCADE    | User who created the task  |
| due_date    | DATE         | Nullable                                     | Task due date              |
| created_at  | TIMESTAMPTZ  | NOT NULL, DEFAULT now()                      | Creation timestamp         |
| updated_at  | TIMESTAMPTZ  | NOT NULL, DEFAULT now()                      | Last update timestamp      |

## SQL

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(120) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(120) NOT NULL,
    description TEXT NOT NULL,
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    title VARCHAR(120) NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'todo' CHECK (
        status IN ('todo', 'in_progress', 'done')
    ),
    assignee_id UUID REFERENCES users(id) ON DELETE SET NULL,
    created_by UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    due_date DATE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    UNIQUE (project_id, title)
);
```
