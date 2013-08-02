from django.core.management.base import BaseCommand, CommandError
from django.contrib.auth.models import User

from players.models import GamePlayer
from players.models import createToken

class Command(BaseCommand):
    def handle(self, *args, **kwargs):
        print "Creating Kyle's user and game player"
        kphelpsUser = User.objects.create_user(username="kphelps",
                                               email="tsunami.ownz@gmail.com",
                                               password="kphelps")
        kphelpsPlayer = GamePlayer(user=kphelpsUser,
                                   token=createToken())
        kphelpsUser.save()
        kphelpsPlayer.save()

        print "Creating Quinten's user and game player"
        quintenUser = User.objects.create_user(username="quinten",
                                               email="quintenpalmer@gmail.com",
                                               password="quinten")
        quintenPlayer = GamePlayer(user=quintenUser, token=createToken())
        quintenUser.save()
        quintenPlayer.save()

        print "Creating Eric's user and game player"
        ericUser = User.objects.create_user(username="edbrown",
                                               email="eric.d.brown23@gmail.com",
                                               password="edbrown")
        ericPlayer = GamePlayer(user=ericUser, token=createToken())
        ericUser.save()
        ericPlayer.save()

        print "Done creating users"
