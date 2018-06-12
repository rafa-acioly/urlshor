create database urlshor;
\c urlshor

create table urls(
	id integer not null,
	url varchar(255) not null,
	encoded varchar(255) not null,
	clicks integer not null DEFAULT 0,
	created_at timestamp not null DEFAULT current_timestamp,
	constraint pk_urls_id primary key (id)
);

create sequence seq_urls_id;

create or replace function increment_clicks_counter(VARCHAR)
	return void as
	"
		UPDATE urls SET clicks = clicks + 1 WHERE encode = $1
	"
language 'sql';