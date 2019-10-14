create table company
(
    id serial not null
        constraint company_pkey
            primary key,
    company_number integer,
    name varchar(256),
    nzbn varchar(256),
    incorporation_date date,
    company_status varchar(32),
    entity_type varchar(32),
    constitution_filed varchar(32)
);

alter table company owner to postgres;

create table director
(
    id serial not null
        constraint director_pkey
            primary key,
    company_id integer,
    full_legal_name varchar(256),
    residential_address text,
    appointment_date date
);

alter table director owner to postgres;

create table previous_name
(
    id serial not null
        constraint previous_name_pk
            primary key,
    company_id integer not null,
    name varchar(256),
    "from" varchar(256),
    "to" varchar(256)
);

alter table previous_name owner to postgres;

create table shareholder
(
    id serial not null
        constraint shareholder_pk
            primary key,
    shareholding_allocation_id integer,
    name varchar(256),
    address text
);

alter table shareholder owner to postgres;

create table shareholding_allocation
(
    id serial not null
        constraint shareholding_pkey
            primary key,
    company_id integer,
    percentage double precision
);

alter table shareholding_allocation owner to postgres;

create table address
(
    id serial not null
        constraint address_pk
            primary key,
    company_id integer,
    registered_office_address varchar,
    address_for_service varchar,
    address_for_shareregister varchar
);