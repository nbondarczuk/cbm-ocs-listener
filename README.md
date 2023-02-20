# CBM-GATEWAY

## Pupose

The microservice is used as CBM flow gateway. It listens on a dedicated Tibco 
EMS queue. Upon event of receving a json message from OCS system it parses 
the payload and the json structure gets mapped to CBM flat structure. 
The record is inserted to the backend table OCS_EVENT. The validation is done
by database constraints, mostly NOT NULL on most of the fields. 
The parameters are ll mandatory with exception of account id where onlu first
one must be provided in json parameter list.

## OCS event json structure, an example

```
{
	"eventId": "11111111",
	"eventType": "6010",
	"eventDescription": "credit-alert",
	"sourceSystem": "OCS",
	"sourceDate": "2019-09-25T11:39:45.430Z",
	"parameters": {
		"parameter": [
			{ "name": "language", "value": "2002" },
			{ "name": "service_retailer", "value": "TMOBILE" },
			{ "name": "header1", "value": "0048699999984" },
			{ "name": "eventname", "value": "credit-alert" },
			{ "name": "notificationcode", "value": "6010" },
			{ "name": "type", "value": "PIR" },
			{ "name": "logInteraction", "value": "true" },
			{ "name": "accountType", "value": "2102" },
			{ "name": "expireDate", "value": "03/10/2019 00:00:00" },
			{ "name": "value", "value": "0" },
			{ "name": "initialAmount", "value": "2500" },
			{ "name": "nextPIR", "value": "TaurusPIR3Unlimited" },
			{ "name": "startDate", "value": "03/09/2019 15:19:59" },
			{ "name": "messageTemplate", "value": "100% ID 2102 PIR3_Taurus_Used_Up" },
			{ "name": "accountInstanceIDList", "value": "701800000000053594" },
			{ "name": "purchaseSerialNumberList", "value": "SO_ID:10182" },
			{ "name": "sessionID", "value": "auto" },
			{ "name": "sessionID", "value": "1499186203" },
			{ "name": "sessionID", "value": "0000059747" },
			{ "name": "sessionID", "value": "0000000001" }]}
}
```

## Target CBM OCS event structure, an example of DDL

```
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
```

## Mapping from OCS to CBM OCS records

 - **EventId**       <- Int(EventId) (primary key), mandatory
 - **EventType**     <- Int(EventType), mandatory
 - **EventDate**     <- SourceDate, mandatory
 - **Msisdn**        <- Parameter with name "header1", mandatory
 - **ExpireDate**    <- Parameter with name "expireDate", mandatory
 - **Value**         <- Parameter with name "value", mandatory
 - **InitialAmount** <- Parameter with name "initialAmount", mandatory
 - **NextPir**       <- Parameter with name "nextPIR", mandatory
 - **StartDate**     <- Parameter with name "startDate", mandatory
 - **AccouintId1**   <- 1st Parameter with name "accountInstanceIDList", mandatory
 - **AccouintId2**   <- 2nd Parameter with name "accountInstanceIDList", optional
 - **EntryDate**     <- SYSDATE, mandatory

## Configuration, an example

```
{
	"RunPath"               : ".",
	"Debug"                 : "0",
	"OracleDBUser"          : "CBMADM",
	"OracleDBPassword"      : "*****",
	"OracleServiceName"     : "t17bscs",
	"QueueServerUrl"        : "tcp://1.2.3.4:7222",	
	"QueueName"             : "taurus.cbm.ocs.alertNotification.EVENT",
	"QueueUser"             : "cbm",
	"QueuePassword"         : "******",
	"CbmServerUrl"          : "http://localhost:80"
}
```

## Error handling

n/a

## Environment

go version go1.10.4 linux/amd64


