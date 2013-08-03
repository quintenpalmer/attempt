Webserver
=========

Requisites
----------
Python
Ubuntu

    sudo apt-get install python python-django django-coffeescript coffee-script sqlite

Fedora

    sudo yum install python python-django django-coffeescript coffee-script sqlite

Database setup
--------------

The file init_db.sh will create and sync the database for the project then insert 3 users with game tokens.
It is advised you run the script before running the webserver, and it is required you run the script from
this webserver/ directory.

Run
---

    python manage.py runserver 127.0.0.1:8000
	cd ..
	GOPATH=`pwd` go run gameserver.go
