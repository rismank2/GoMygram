create table socialmedia(
Id SERIAL primary key,
Name varchar(100) not null,
Social_media_url varchar(100) not null,
UserId int not null references public.users(id)
);