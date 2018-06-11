\l
\c urlshor
\d urls

create database urlshor;
\c urlshor

create table urls(
	`id` serial not null,
	`url` varchar(255) not null,
	`clicks` integer not null,
	`created_at` timestamp not null default current_timestamp;
);
