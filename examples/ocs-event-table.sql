create table ocs_events 
(
event_id number not null,
event_type number not null,
event_date date not null,
msisdn varchar2(100) not null,
expire_date date not null,
value number not null,
initial_amount number not null,
next_pir varchar2(100) null,
start_date date not null,
accountId1 varchar2(100) not null,
accountId2 varchar2(100),
entry_date date not null
) 
