#/bin/bash

# make user CBMADM to handle the backend
ORA='GSYSADM/cgsysadm17@t17bscs'
sqlplus ${ORA} @create_user_cbmgtw.sql

# make tables in CGSYSADM schema to be accessed by CBMADM
# ORA="CBMADM/cbmadm17@t17bscs"
# sqlplus ${ORA} @create_ocs_event.sql
