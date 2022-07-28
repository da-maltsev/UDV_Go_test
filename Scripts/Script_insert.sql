insert into
    Publishers (Title, Adress, Contacts)
values
    ('First Publisher', 'Petrovsko 12, 541', 'www.leningrad.ru'),
    ('Second Publisher', 'Tolstogo 42, 987', '89321231212')
;
insert into
    Books (Title, Description, Authors, PublishYear, Publisher_id, Edition)
values
    ('War and Peace', 'The world-famous russian masterpiece', 'Leo Tolstoy', '1993-08-24', 1, 'best edition'),
    ('War and Piece', 'Very dumm mistake', 'Teo Lolstoy', '2007-07-07', 2, '1 edition'),
    ('Some random book', 'not very interesting', 'Random author', '2013-12-01', 1, 'first edition'),
    ('Fourth book', 'it is the 4th book', 'Anon', '2000-01-10', 2, 'we did our best'),
    ('Live to win', 'til u die', 'Paul Stanley', '1971-09-24', 1, 'deluxe edition'),
    ('Not the last book', 'not first not last', 'Ben Kenobi', '1980-05-04', 2, 'force edition'),
    ('The last book', 'definitely last book', 'Author DM', '1999-09-09', 1, '9th edition')
;
insert into
    Items (Quantity, Book_id)
values
    (3, 1),
    (9, 2),
    (7, 3),
    (5, 4),
    (14, 5),
    (11, 6),
    (9, 7)
;
insert into
    Positions (Item_id, Room, Shelf)
values
    (1, '12b', '3a'),
    (2, '1c', '41'),
    (3, '2a', '7a'),
    (4, '12b', '7c'),
    (5, '7', '90'),
    (6, '1b', '5g'),
    (7, '99', '9d')

;
insert into
    Customers (Name, Contacts)
values
    ('Antony Draskauses', '8-800-555-35-35'),
    ('Anakin Skywalker', 'badboy007@empire.org'),
    ('Some Random Guy', '0000001111'),
    ('Julia Roberts', 'whynot@gmail.com')
;
insert into
    Status (Item_id, Available, Customer_id, Given_at, Recieved_at)
values
    (1, true, 2, '2020-01-22', '2022-02-01'),
    (2, false, 2, '2020-01-22', '2022-02-01'),
    (3, false, 1, '2021-11-02', '2022-03-05'),
    (4, true, 4, '2020-01-29', '2022-02-01'),
    (7, false, 3, '2020-06-12', '2022-02-01'),
    (6, false, 1, '2021-03-23', '2022-03-05'),
    (5, true, 4, '2019-01-22', '2022-02-01')
;