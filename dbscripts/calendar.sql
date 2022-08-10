-- DROP DATABASE IF EXISTS calendar;
-- CREATE DATABASE calendar;

DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id            UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    login         character varying(20)    NOT NULL UNIQUE,
    password_hash character varying(255)   NOT NULL,
    created_at    timestamp with time zone NOT NULL DEFAULT now(),
    is_deleted    bool                              DEFAULT false
);

DROP TABLE IF EXISTS tokens;
CREATE TABLE tokens
(
    id    serial                 NOT NULL UNIQUE,
    token character varying(255) NOT NULL UNIQUE
);

DROP TABLE IF EXISTS companies;
CREATE TABLE companies
(
    id         UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    name       character varying(255)   NOT NULL UNIQUE,
    profile    character varying(255)   NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    is_deleted bool                              DEFAULT false
);

DROP TABLE IF EXISTS guests;
CREATE TABLE guests
(
    id         UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    firstname  character varying(255)   NOT NULL,
    lastname   character varying(255)   NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    is_deleted bool                              DEFAULT false
);

DROP TABLE IF EXISTS events;
CREATE TABLE events
(
    id          UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    title       character varying(255)   NOT NULL,
    description character varying(255),
    start_time  timestamp with time zone NOT NULL,
    duration    int                      NOT NULL,
    created_at  timestamp with time zone NOT NULL DEFAULT now(),
    is_deleted  bool                              DEFAULT false
);

DROP TABLE IF EXISTS event_guests;
CREATE TABLE event_guests
(
    id         serial                                        NOT NULL UNIQUE,
    event_id   UUID references events (id) ON DELETE CASCADE NOT NULL,
    guest_id   UUID references guests (id) ON DELETE CASCADE NOT NULL,
    is_deleted bool DEFAULT false
);

DROP TABLE IF EXISTS companies_guests;
CREATE TABLE companies_guests
(
    id         serial                                           NOT NULL UNIQUE,
    company_id UUID references companies (id) ON DELETE CASCADE NOT NULL,
    guest_id   UUID references guests (id) ON DELETE CASCADE    NOT NULL,
    is_deleted bool DEFAULT false
);

-- bootstrap
INSERT INTO users (login, password_hash) VALUES ('me', 'hashHere');