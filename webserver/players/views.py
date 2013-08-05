# Create your views here.
from django.contrib.auth.models import User
from django.http import HttpResponse
from django.http import HttpResponseRedirect
from django.shortcuts import render_to_response
from django.contrib import auth
from django.template import RequestContext

from players.models import GamePlayer
from players.models import createToken

def login(request):
    if request.user.is_authenticated():
        return HttpResponseRedirect("/play/")
    if request.method == 'GET':
        return render_to_response("login.html", RequestContext(request, {}))
    username = request.POST.get('username', '')
    password = request.POST.get('password', '')
    user = auth.authenticate(username=username, password=password)
    if user is not None and user.is_active:
        auth.login(request, user)
        return HttpResponseRedirect("/play/")
    rc = RequestContext(request, {"failed_login": True})
    return render_to_response("login.html", rc)

def logout(request):
    auth.logout(request)
    return HttpResponseRedirect("/account/login/")
