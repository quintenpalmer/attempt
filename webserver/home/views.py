# Create your views here.
from django.http import HttpResponse
from django.http import HttpResponseRedirect
from django.shortcuts import render_to_response
from django.template import RequestContext

from players.models import GamePlayer

def index(request):
    if not request.user.is_authenticated():
        return HttpResponseRedirect("/account/login/")
    profile = GamePlayer.objects.get(user = request.user)
    rc = RequestContext(request, {"user": request.user,
                                  "username": request.user.username,
                                  "token": profile.token})
    return render_to_response("play.html", rc)
