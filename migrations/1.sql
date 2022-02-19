Create TABLE authors(
    id serial,
    name text,
    birthday varchar(191)
);
CREATE TABLE books (
    id serial,
    name text,
    author_id int references authors(id),
    seria text,
    year int,
    page_count int,
    format text,
    type varchar(191),
    weight int,
    age varchar(191)
);
alter table books
    add constraint books_pk
        primary key (name, seria, author_id, year);

create table tags(
    id serial,
    name varchar(191)
);
create table books_tags(
    id serial,
    book_id int references books(id),
    tag_id int references tags(id)
);
CREATE TABLE users(
    id serial,
    name text,
    email text,
    birthday varchar(191)
);
CREATE TABLE favorite_books(
    id serial,
    user_id int references users(id),
    book_id int references books(id)
);
CREATE TABLE processing_books(
    id serial,
    user_id int references users(id),
    book_id int references books(id),
    page int
);