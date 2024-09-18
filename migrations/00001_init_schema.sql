-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS gophertask;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS gophertask.users
(
    id       UUID DEFAULT uuid_generate_v4(),
    login    VARCHAR(30)  NOT NULL,
    password VARCHAR(256) NOT NULL,
    CONSTRAINT pk_users PRIMARY KEY (id),
    CONSTRAINT users_login_unique UNIQUE (login)
);

CREATE TYPE gophertask.TASK_STATUS AS ENUM ('NEW', 'IN PROGRESS', 'DONE');

CREATE TABLE IF NOT EXISTS gophertask.tasks
(
    id          UUID DEFAULT uuid_generate_v4(),
    user_id     UUID                   NOT NULL,
    name        VARCHAR(256)           NOT NULL,
    description TEXT,
    status      gophertask.TASK_STATUS NOT NULL,
    duration    INTERVAL,
    started_at  TIMESTAMP,
    CONSTRAINT pk_tasks PRIMARY KEY (id),
    CONSTRAINT tasks_to_users_fk
        FOREIGN KEY (user_id) REFERENCES gophertask.users (id)
);

CREATE TABLE IF NOT EXISTS gophertask.epics
(
    id UUID DEFAULT uuid_generate_v4(),
    CONSTRAINT pk_epics PRIMARY KEY (id),
    CONSTRAINT epics_to_tasks_fk
        FOREIGN KEY (id) REFERENCES gophertask.tasks (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS gophertask.subtasks
(
    id      UUID DEFAULT uuid_generate_v4(),
    epic_id UUID NOT NULL,
    CONSTRAINT pk_subtasks PRIMARY KEY (id),
    CONSTRAINT subtasks_to_tasks_fk
            FOREIGN KEY (id) REFERENCES gophertask.tasks (id) ON DELETE CASCADE,
    CONSTRAINT subtasks_to_epics_fk
        FOREIGN KEY (epic_id) REFERENCES gophertask.epics (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS gophertask.subtasks;
DROP TABLE IF EXISTS gophertask.epics;
DROP TABLE IF EXISTS gophertask.tasks;
DROP TABLE IF EXISTS gophertask.users;
-- +goose StatementEnd