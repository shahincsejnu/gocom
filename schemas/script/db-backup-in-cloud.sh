#!/bin/bash

# Set variables
LOCAL_BACKUP_DIR="./backup"
VM_NAME="db-server"
VM_USERNAME="We_Dont Know_Devops"
VM_BACKUP_DIR="/backup"
DB_USER="gocom"
DB_PASSWORD="gocom123"
DB_NAME="gocomdb"

# Create a backup on the local machine
pg_dump -U $DB_USER -d $DB_NAME > $LOCAL_BACKUP_DIR/backup.sql

# Copy the backup file to the cloud VM
scp $LOCAL_BACKUP_DIR/backup.sql $VM_USERNAME@$VM_NAME:$VM_BACKUP_DIR/

# Remove the local backup file
rm $LOCAL_BACKUP_DIR/backup.sql
