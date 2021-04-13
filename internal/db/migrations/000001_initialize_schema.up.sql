CREATE TABLE IF NOT EXISTS users (
    username            VARCHAR(255)        NOT NULL UNIQUE,
    id                  VARCHAR(36)         PRIMARY KEY,
    hashed_password     VARCHAR(255)        NOT NULL,
    email               VARCHAR(255)        NOT NULL UNIQUE,
    date_of_birth       INTEGER             NOT NULL DEFAULT 0,
    gender              VARCHAR(100)        NOT NULL DEFAULT 'male',
    membership          VARCHAR(200)        NOT NULL DEFAULT 'none',
    avatar_data         VARCHAR(500)        NOT NULL DEFAULT 'none',
    device_id           VARCHAR(200)        NOT NULL DEFAULT 'none',
    rate                INTEGER             NOT NULL DEFAULT 0,
    is_registered       BOOLEAN             NOT NULL DEFAULT false,
    is_verified         BOOLEAN             NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS user_events (
    id varchar(36) primary key,
    user_id varchar(36) references users (id) on delete cascade,
    device_id varchar(255) NOT NULL,
    data jsonb,
    time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)