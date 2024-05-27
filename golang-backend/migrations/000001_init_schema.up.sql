CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username varchar(50),
    created_at timestamp
);

CREATE TABLE folders (
    id SERIAL PRIMARY KEY,
    name varchar(50),
    owner integer references users(id) on DELETE CASCADE,
    parent integer DEFAULT NULL,
    created_at timestamp
);

CREATE TABLE users_to_main_folders (
    "user" integer references users(id) on DELETE CASCADE,
    folder integer references folders(id) on DELETE CASCADE,
    PRIMARY KEY ("user", folder)
);

CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    owner integer references users(id) on DELETE CASCADE,
    folder integer references folders(id) on DELETE CASCADE,
    name varchar(50),
    created_at timestamp,
    is_public bool DEFAULT false
);

CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    note integer references notes(id) on DELETE CASCADE,
    owner integer references users(id) on DELETE CASCADE,
    body text,
    answer text,
    is_public bool DEFAULT false
);

CREATE TABLE public_notes (
    id integer references notes(id) on DELETE CASCADE,
    author integer references users(id) on DELETE CASCADE,
    name varchar(50),
    descr text,
    stars integer,
    created_at timestamp
);

CREATE TABLE tags (
  id SERIAL PRIMARY KEY,
  name varchar(50)
);

CREATE TABLE tags_notes (
    note integer references public_notes(id),
    tag integer references tags(id),
    PRIMARY KEY (note, tag)
);

CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    admin integer references users(id),
    is_private bool,
    link varchar(255),
    name varchar(50),
    descr text
);

CREATE TABLE groups_tags (
    "group" integer references groups(id),
    tag integer references tags(id),
    PRIMARY KEY ("group", tag)
);

CREATE TABLE groups_notes (
    "group" integer references groups(id),
    note integer references notes(id),
    PRIMARY KEY ("group", note)
);
