#!/bin/sh

echo "Running database migrations..."
./migrate

if [ $? -eq 0 ]; then
    echo "Migrations completed successfully"
    echo "Starting the application..."
    ./main
else
    echo "Migration failed"
    exit 1
fi