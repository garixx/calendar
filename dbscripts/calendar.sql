DROP DATABASE IF EXISTS calendar;
CREATE DATABASE calendar;

DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id           UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    login        character varying(20)    NOT NULL UNIQUE,
    passwordHash character varying(255)   NOT NULL,
    createdAt    timestamp with time zone NOT NULL DEFAULT now(),
    isDeleted    bool                              DEFAULT false
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
    id        UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    name      character varying(255)   NOT NULL UNIQUE,
    profile   character varying(255)   NOT NULL,
    createdAt timestamp with time zone NOT NULL DEFAULT now(),
    isDeleted bool                              DEFAULT false
);

DROP TABLE IF EXISTS guests;
CREATE TABLE guests
(
    id        UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    firstName character varying(255)   NOT NULL,
    lastName  character varying(255)   NOT NULL,
    createdAt timestamp with time zone NOT NULL DEFAULT now(),
    isDeleted bool                              DEFAULT false
);

DROP TABLE IF EXISTS events;
CREATE TABLE events
(
    id          UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    title       character varying(255)   NOT NULL,
    description character varying(255),
    startTime   timestamp with time zone NOT NULL,
    duration    int                      NOT NULL,
    createdAt   timestamp with time zone NOT NULL DEFAULT now(),
    isDeleted   bool                              DEFAULT false
);

DROP TABLE IF EXISTS event_guests;
CREATE TABLE event_guests
(
    id        serial                                       NOT NULL UNIQUE,
    eventId   int references events (id) ON DELETE CASCADE NOT NULL,
    guestId   int references guests (id) ON DELETE CASCADE NOT NULL,
    isDeleted bool DEFAULT false
);

DROP TABLE IF EXISTS companies_guests;
CREATE TABLE companies_guests
(
    id        serial                                          NOT NULL UNIQUE,
    companyId int references companies (id) ON DELETE CASCADE NOT NULL,
    guestId   int references guests (id) ON DELETE CASCADE    NOT NULL,
    isDeleted bool DEFAULT false
);
