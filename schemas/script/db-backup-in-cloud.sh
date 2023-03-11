#!/bin/bash

# Set variables
LOCAL_BACKUP_DIR="/path/to/local/backup/directory"
VM_NAME="cloud_vm"
VM_USERNAME="vm_user"
VM_BACKUP_DIR="/backup"
DB_USER="db_user"
DB_PASSWORD="db_password"
DB_NAME="my_database"

# Create a backup on the local machine
pg_dump -U $DB_USER -d $DB_NAME > $LOCAL_BACKUP_DIR/backup.sql

# Copy the backup file to the cloud VM
scp $LOCAL_BACKUP_DIR/backup.sql $VM_USERNAME@$VM_NAME:$VM_BACKUP_DIR/

# Remove the local backup file
rm $LOCAL_BACKUP_DIR/backup.sql
