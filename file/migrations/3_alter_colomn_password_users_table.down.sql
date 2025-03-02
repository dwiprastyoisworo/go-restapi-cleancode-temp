alter table users
alter column password type varchar(50) using password::varchar(50);