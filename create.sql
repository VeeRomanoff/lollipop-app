CREATE TABLE lollipop_users (
    id bigserial,
    name text,
    age int,
    height float,
    hobbies jsonb,
    description text,
    email text not null unique
);

ALTER TABLE lollipop_users
    ADD COLUMN email text,
    ADD CONSTRAINT unique_email UNIQUE (email);