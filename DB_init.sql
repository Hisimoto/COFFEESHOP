create table users (
ID bigserial not null primary key,
email VARCHAR not null unique,
membership_type numeric not null
);
insert into users(email,membership_type) values ('cristi@gmail.com',1), ('ion@gmail.com',2), ('andrei@gmail.com',3), ('igor@gmail.com',1), ('vasea@gmail.com',2), ('mihai@gmail.com',3); 


create table orders (
ID bigserial not null primary key,
userid numeric not null,
coffee_type numeric,
order_date timestamp
);
insert into orders(userid, coffee_type,order_date) values (1,1,current_timestamp - interval '25 hour'), 
														  (1,1,current_timestamp - interval '5 hour'), 
														  (1,1,current_timestamp - interval '1 hour'), 
														  (1,2,current_timestamp - interval '20 hour'), 
														  (1,2,current_timestamp - interval '5 hour'), 
														  (1,2,current_timestamp - interval '1 hour'), 
														  (2,1,current_timestamp - interval '20 hour'), 
														  (2,1,current_timestamp - interval '5 hour'), 
														  (2,1,current_timestamp - interval '1 hour'), 
														  (3,1,current_timestamp - interval '1 hour'), 
														  (3,1,current_timestamp - interval '2 hour'), 
														  (3,1,current_timestamp - interval '3 hour'), 
														  (3,1,current_timestamp - interval '4 hour'), 
														  (3,1,current_timestamp - interval '5 hour'), 
														  (3,2,current_timestamp - interval '1 hour'), 
														  (3,3,current_timestamp - interval '1 hour');


create table coffee_types(
coffee_type numeric,
coffee_desc varchar
);
insert into coffee_types(coffee_type, coffee_desc) values (1, 'Espresso'), (2, 'Americano'), (3, 'Cappuccino');


create table membership (
ID bigserial not null primary key,
membership_id numeric,
membership_name varchar,
coffee_type numeric,
coffee_amount numeric,
timeframe numeric
);
insert into membership(membership_id,membership_name,coffee_type,coffee_amount,timeframe) values (1, 'Basic', 1,2,24), (1, 'Basic', 2,3,24), (1, 'Basic', 3,1,24),
							  (2, 'Coffeelover', 1,5,24), (2, 'Coffeelover', 2,5,24), (2, 'Coffeelover', 3,5,24),
							  (3, 'Espresso Maniac', 1,5,1), (3, 'Espresso Maniac', 2,3,24), (3, 'Espresso Maniac', 3,1,24);

							  

