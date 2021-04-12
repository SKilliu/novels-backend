CREATE TABLE IF NOT EXISTS users (
    id varchar(36) primary key,
    name varchar(255) not null,
    hashed_password varchar(255) not null,
    email varchar(255) unique
);

CREATE TABLE IF NOT EXISTS user_events (
    id varchar(36) primary key,
    user_id varchar(36) references users (id) on delete cascade,
    device_id varchar(255) NOT NULL,
    data jsonb,
    time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)