-- Write your migrate up statements here

create table plan_types (
  id serial primary key,
  plan_name varchar(255) not null,
  number_of_days smallint not null,
  times_per_day smallint not null,
  amount smallint not null,
  max_allowed_messout_days smallint not null
);

insert into plan_types (plan_name, number_of_days, times_per_day, amount, max_allowed_messout_days) values
('Super Saver Pack', 30, 3, 3999, 5),
('Flexi Meal', 30, 2, 2999, 5),
('Solo Meal', 30, 1, 1999, 5),
('Quick Start Trial', 10, 3, 1599, 3);

---- create above / drop below ----

drop table plan_types;
