#!/bin/bash

# Set variables
VM_NAME="db-server"
DB_USER="gocom"
DB_PASSWORD="gocom123"
DB_NAME="gocomdb"
BACKUP_DIR="/backup"

# Connect to the VM and run the backup command
ssh $VM_NAME "PGPASSWORD=$DB_PASSWORD pg_dump -U $DB_USER $DB_NAME > $BACKUP_DIR/backup.sql"

# Copy the backup file from the VM to the local machine
scp $VM_NAME:$BACKUP_DIR/backup.sql ./backup
