-- Write your migrate up statements here

CREATE TABLE plans (
    id serial primary key,
    start_date date not null,
    end_date date not null,
    plan_type_id int not null,  -- this will reference the plan_types table
    is_active boolean not null,
    has_paid boolean not null,
    messout_days date[],  -- array of dates for messout days
    constraint fk_plan_type_id foreign key (plan_type_id)
    references plans(id)
);


---- create above / drop below ----

drop table plans;

