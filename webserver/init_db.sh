#!/bin/sh
echo "Hopefully you have sqlite3 installed!"

echo "Creating the sqlite3 database"
sqlite3 djangodb.db ""
echo "Database created"

echo "Syncing database with Django, make sure to create a superuser account"
python manage.py syncdb

echo "Populating database with basic user info"
python manage.py populatedb

echo "Database setup completed!"
