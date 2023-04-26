create table countries
(
    id     bigserial,
    name   text not null,
    icon   text,
    active boolean default true,
    primary key (id)
);


create table languages(
    id     bigserial,
    name   text not null,
    icon   text not null,
    active boolean default true,
    primary key (id)
);

create table sys_messages
(
    id     bigserial,
    name   text not null,
    active boolean default true,
    primary key (id)
);

create table currencies(
    id   bigserial,
    name text not null,
    icon text,
    primary key (id)
);

create table transfers(
    id    bigserial primary key ,
    entity text,
    entity_id int,
    lang_id int,
    keyfield int,
    value text,

--     foreign key (entity, entity_id) references countries(name, id),
    foreign key (lang_id) references languages(id)

);
create table agents(
    id bigserial primary key,
    name       text,
    legal_name text,
    active     boolean default true
);


create table account_agents
(
    id       bigserial,
    agent_id bigint not null,
    curr_id  bigint not null,
    active   boolean default true,
    is_default   boolean default true,
    type     text not null,
    primary key (id),
    foreign key (agent_id) references agents(id),
    foreign key (curr_id) references currencies(id)

);

create table userinfo (
    id bigserial primary key,
    name text,
    icon text,
    active bool default true,
    sort int
);

create table vendor (
  id bigserial primary key,
  name text,
  active bool
);

create table services (
    id bigserial primary key,
    vendor_id bigint references vendor,
    name text,
    active bool,
    type text
);

create table services_country(
    service_id bigint references services,
    country_id bigint references countries,
    active bool
);

create table services_rules(
    id bigserial primary key,
    name text,
    type text
);

create table Payment_type(
    id bigserial primary key,
    name text
);


--

create table registrator(
                            id bigserial,
                            request text,
                            name text,
                            icon text,
                            active bool,
                            entity text,
                            entity_id int,
                            lang_id int,
                            keyfield int,
                            value text,
                            legal_name text,
                            agent_id bigint not null,
                            curr_id  bigint not null,
                            is_default   boolean default true,
                            type     text   not null,
                            sort int,
                            status text
);

