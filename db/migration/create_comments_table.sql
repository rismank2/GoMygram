create table comment(
Id SERIAL primary key,
User_id int not null references public.users(id),
Photo_id int not null references public.photo(id),
Message varchar(100) not null,
Created_at date not null,
Updated_at date not null
);