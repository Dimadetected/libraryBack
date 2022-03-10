Create TABLE authors(
    id serial,
    name text,
    birthday varchar(191)
);

CREATE TABLE books (
    id serial,
    name text,
    description text,
    author_id int references authors(id),
    year text,
    age varchar(191),
    file text not null default ''
);


alter table books
    add constraint books_pk
        primary key (name,author_id, year);

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
    birthday varchar(191),
    user_token text,
    role_id int
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
    page int,
    pages int,
    created timestamp
);

CREATE TABLE reviews(
    id serial,
    user_id int references users(id),
    book_id int references books(id),
    description text,
    grade int,
    positive int,
    negative int
);

CREATE TABLE reviews_grades(
    id serial,
    user_id int,
    book_id int,
    status int
)