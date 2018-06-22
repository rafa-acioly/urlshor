create database urlshor;
\c urlshor

create table urls(
	id bigint not null,
	url varchar(255) not null,
	encoded varchar(255) not null UNIQUE,
	clicks bigint DEFAULT 0,
	created_at timestamp not null DEFAULT current_timestamp,
	constraint pk_urls_id primary key (id)
);

create sequence seq_urls_id;
create index idx_encoded on urls(encoded);
