from django.conf.urls import patterns, include, url

from . import views

urlpatterns = patterns('',
    url(r'^$',views.login),
    url(r'^login/$',views.login),
    url(r'^logout/$',views.logout),
)
