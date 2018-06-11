create database urlshor;
\c urlshor

create table urls(
	id integer not null,
	url varchar(255) not null,
	clicks integer not null,
	created_at timestamp not null default current_timestamp,
	constraint pk_urls_id primary key (id)
);
