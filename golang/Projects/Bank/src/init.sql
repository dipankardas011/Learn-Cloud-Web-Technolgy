show databases ;
create database Bank;

use Bank;

show tables;

-- drop table mainBank;

create table if not exists branchBank (
    branchID int primary key,
    name varchar(10)
);

insert into branchBank values(1, 'NEWSQR');
insert into branchBank values(2, 'ORIG');
insert into branchBank values(3, 'PIPEL');
insert into branchBank values(4, 'DBDBD');

create table if not exists mainBank (
    accountNumber bigint PRIMARY KEY,
    name varchar(20),
    balance float,
    branchID int,
    foreign key (branchID) references branchBank(branchID)
);

insert into mainBank values(2000204324, 'USR_1', 40050.02, 3);
insert into mainBank values(2000206324, 'USR_3', 40012.034, 1);
insert into mainBank values(2000204624, 'USR_10', 423450.033, 2);
insert into mainBank values(2000204326, 'USR_4', 4003330.056, 3);
insert into mainBank values(2000204684, 'USR_6', 4002450.024, 4);
insert into mainBank values(2100204323, 'USR_2', 40053240.032, 4);
insert into mainBank values(2110204324, 'USR_8', 40052340.0211, 1);
insert into mainBank values(2002344324, 'USR_9', 50.021, 1);
insert into mainBank values(2002678424, 'USR_5', 21050.024, 1);

select * from branchBank;
select * from mainBank;

select m.accountNumber as ACCNO,
       m.name as NAME,
       m.balance as AVIL_BAL,
       b.name as BRANCH
from branchBank b, mainBank m
where b.branchID = m.branchID
order by accountNumber;

