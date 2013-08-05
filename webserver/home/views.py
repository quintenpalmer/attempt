# Create your views here.
from django.http import HttpResponse
from django.http import HttpResponseRedirect
from django.shortcuts import render_to_response
from django.template import RequestContext

def index(request):
    if not request.user.is_authenticated():
        return HttpResponseRedirect("/account/login/")
    rc = RequestContext(request, {"user": request.user})
    return render_to_response("play.html", rc)
