create table users(
Id SERIAL primary key,
Username varchar(100) unique not null,
Email varchar(50) unique not null,
Password varchar(100)not null,
Age int not null,
Created_at date not null,
Updated_at date not null
);
