create database urlshor;
\c urlshor

create table urls(
	id integer not null,
	url varchar(255) not null,
	encoded varchar(255) not null,
	clicks integer not null,
	created_at timestamp not null default current_timestamp,
	constraint pk_urls_id primary key (id)
);

create sequence seq_urls_id;

create or replace function increment_clicks_counter(VARCHAR)
	return void as
	"
		UPDATE urls SET clicks = 1 + (
			SELECT clicks FROM urls WHERE encode = $1
		) WHERE encode = $1;
	"
language 'sql';