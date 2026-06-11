# Database Design

## users

Stores the application users. Each user can own projects and can be assigned to tasks.

| Column        | Type         | Constraints             | Description                             |
| ------------- | ------------ | ----------------------- | --------------------------------------- |
| id            | UUID         | Primary key             | Unique user identifier                  |
| name          | VARCHAR(120) | NOT NULL                | User full name                          |
| email         | VARCHAR(255) | NOT NULL, UNIQUE        | User email used for login               |
| password_hash | TEXT         | NOT NULL                | Hashed password, never store plain text |
| created_at    | TIMESTAMPTZ  | NOT NULL, DEFAULT now() | Creation timestamp                      |
| updated_at    | TIMESTAMPTZ  | NOT NULL, DEFAULT now() | Last update timestamp                   |

### SQL

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(120) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
```
