#!/bin/bash

# Set variables
VM_NAME="my_vm"
DB_USER="db_user"
DB_PASSWORD="db_password"
DB_NAME="my_database"
BACKUP_DIR="/backup"

# Connect to the VM and run the backup command
ssh $VM_NAME "PGPASSWORD=$DB_PASSWORD pg_dump -U $DB_USER $DB_NAME > $BACKUP_DIR/backup.sql"

# Copy the backup file from the VM to the local machine
scp $VM_NAME:$BACKUP_DIR/backup.sql /path/to/local/backup/directory/
