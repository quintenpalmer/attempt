attempt
=======

Role Playing Game using Go/Django/HTML5

Django Help
===========

Check out the README.md file in webserver/ for extra notes on the Django app.

Go Dependencies
============

Installing the following dependances will require mercurial, which you should be able to install from your favorite
package manager.

If you would like to install the Go dependancies for this project you can simply run "go get" from the project root.
If you would prefer to install them one at a time, use the urls and directions below. 

The Go server for this project uses several open source packages which must be installed before the server can be run.
If you have Go, and especially the "go" command line tool, properly installed, you should be able to "go get <pkg name>"
these packages with ease:

- cgl.tideland.biz/applog
  - A level based logging system for fun and profit
- github.com/garyburd/go-websocket/websocket
  - A Go websocket library for communicating with the front end
