create table groups(
    id serial primary key,
    title varchar(255) unique not null,
    parent_id int
);

create table humans(
    id serial primary key,
    name varchar(255) not null,
    surname varchar(255) not null,
    year_of_birth date not null,
    group_id int references groups(id) on delete cascade not null
);

