CREATE table if not exists Publishers (
ID serial primary key,
Title varchar(50) not null,
Adress varchar(100) not null,
Contacts varchar(50) not null
);
CREATE table if not exists Books (
ID serial primary key,
Title varchar(50) not null,
Description varchar(200) not null,
Authors varchar(100) not null,
PublishYear date,
Publisher_id integer not null references Publishers(ID),
Edition varchar(50) not null
);
CREATE  table if not exists Items (
ID serial primary key,
Quantity integer,
Book_id integer not null references Books(ID)
);
CREATE  table if not exists Positions (
ID serial primary key,
Item_id integer not null references Items(ID),
Room varchar(20) not null,
Shelf varchar(10) not null
);
CREATE  table if not exists Customers (
ID serial primary key,
Name varchar(50) not null,
Contacts varchar(50) not null
);
CREATE  table if not exists Status (
ID serial primary key,
Item_id integer not null references Items(ID),
Available boolean not null,
Customer_id integer references Customers(ID),
Given_at date,
Recieved_at date
);