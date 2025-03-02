alter table users
alter column password type varchar(100) using password::varchar(100);