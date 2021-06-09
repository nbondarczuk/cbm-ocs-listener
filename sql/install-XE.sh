#/bin/bash

# make user CBMADM to handle the backend
ORA='SYSTEM/oracle@XE'
sqlplus ${ORA} @create_user_cbmadm.sql
sqlplus ${ORA} @create_user_cbmgtw.sql

# make tables in CGSYSADM schema to be accessed by CBMADM
ORA="CBMADM/CBMADM@XE"
sqlplus ${ORA} @create_ocs_event.sql
