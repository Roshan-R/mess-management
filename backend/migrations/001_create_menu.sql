-- Create the table for menu
create table menus(
  id serial PRIMARY KEY,
  active_date DATE NOT NULL,
  breakfast VARCHAR NOT NULL,
  lunch VARCHAR NOT NULL,
  dinner VARCHAR NOT NULL
);

insert into menus (active_date, breakfast, lunch, dinner) values
('2024-11-21', 'Pancakes with syrup', 'Grilled Chicken Sandwich', 'Salmon with asparagus'),
('2024-11-22', 'Omelette with cheese', 'Caesar Salad', 'Beef Stew with mashed potatoes'),
('2024-11-23', 'French Toast', 'Vegetable Wrap', 'Chicken Alfredo'),
('2024-11-24', 'Smoothie with fruit', 'Turkey Club Sandwich', 'Pasta Primavera'),
('2024-11-25', 'Bagel with cream cheese', 'Tuna Salad', 'Steak with roasted vegetables');

---- create above / drop below ----
drop table menus;
