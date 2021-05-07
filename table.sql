
drop database if exists go_farmer;
create database go_farmer;
use go_farmer;
create table fsql_test
(
    id    bigint default 0,
    name    varchar(256) not null,
    primary key(id)
) comment="";
