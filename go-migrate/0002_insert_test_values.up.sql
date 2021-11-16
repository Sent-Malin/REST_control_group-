insert into groups(title, parent_id) 
values ('Distribution', NULL),
       ('Football', NULL),
       ('Fans', 2),
       ('Players', 2),
       ('University', NULL),
       ('Teachers', 5),
       ('Students', 5);

insert into humans(name, surname, year_of_birth, group_id) 
values ('Igor', 'Vityaev', '1998-04-10', (select id from groups where title = 'Players')),
       ('Viktor', 'Petrov', '1990-01-22', (select id from groups where title = 'Football')),
       ('Aleksey', 'Lomtik', '2000-06-30', (select id from groups where title = 'Players')),
       ('Fedor', 'Istec', '1975-09-28', (select id from groups where title = 'Students')),
       ('Gennadii', 'Opozdal', '1999-12-22', (select id from groups where title = 'Distribution')),
       ('Kirill', 'Serov', '1992-04-04', (select id from groups where title = 'Students')),
       ('Nikolai', 'Peskunov', '1977-03-11', (select id from groups where title = 'Distribution')),
       ('Daniil', 'Krimov', '1989-07-25', (select id from groups where title = 'Teachers')),
       ('Nikita', 'Udacha', '2000-02-07', (select id from groups where title = 'Fans')),
       ('Ivan', 'Karkarov', '1995-01-19', (select id from groups where title = 'Fans')),
       ('Michail', 'Dyachuk', '1990-11-05', (select id from groups where title = 'University'));