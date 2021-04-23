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

CREATE TABLE IF NOT EXISTS reset_password_requests (
    id      VARCHAR(36)         PRIMARY KEY,
    user_id VARCHAR(36)         REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_socials (
    id          VARCHAR(36)         PRIMARY KEY,
    user_id     VARCHAR(36)         REFERENCES users (id) ON DELETE CASCADE,
    social      VARCHAR(100)        NOT NULL,
    social_id   VARCHAR(36)         NOT NULL
);

CREATE TABLE IF NOT EXISTS novels (
    id                          VARCHAR(36)         PRIMARY KEY,
    user_id                     VARCHAR(36)         REFERENCES users (id) ON DELETE CASCADE,
    title                       VARCHAR(255)        NOT NULL,
    data                        VARCHAR             NOT NULL,
    participated_in_competiton  BOOLEAN             NOT NULL DEFAULT false,
    voting_result               INTEGER             NOT NULL DEFAULT 0,
    created_at                  INTEGER             NOT NULL DEFAULT 0,
    updated_at                  INTEGER             NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS novels_pool (
    id                      VARCHAR(36)         PRIMARY KEY,
    novel_one_id            VARCHAR(36)         NOT NULL,
    novel_two_id            VARCHAR(36)         NOT NULL,
    user_one_id             VARCHAR(36)         NOT NULL,
    user_two_id             VARCHAR(36)         NOT NULL,
    competition_started_at  INTEGER             NOT NULL,
    status                  VARCHAR(50)         NOT NULL,
    novel_one_votes         FLOAT               NOT NULL,
    novel_two_votes         FLOAT               NOT NULL,
    created_at              INTEGER             NOT NULL,
    updated_at              INTEGER             NOT NULL        
);

CREATE TABLE IF NOT EXISTS ready_for_vote (
    id              VARCHAR(36)         PRIMARY KEY,
    user_id         VARCHAR(36)         NOT NULL,
    novels_pool_id  VARCHAR(36)         REFERENCES novels_pool (id) ON DELETE CASCADE,
    views_amount    INTEGER             NOT NULL,
    is_voted        BOOLEAN             NOT NULL
);