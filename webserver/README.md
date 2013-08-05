Webserver
=========

Requisites
----------
Python
Ubuntu

    sudo apt-get install python python-django django-coffeescript coffee-script sqlite

Fedora

    sudo yum install python python-django django-coffeescript coffee-script sqlite

It is very possible your package manager will not know about django-coffeescript. If this
is the case you can install pip, the Python package manager, via your package manager, then run:

    sudo pip install django-coffeescript

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
