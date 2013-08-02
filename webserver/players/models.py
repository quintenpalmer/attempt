from django.db import models
from django.contrib.auth.models import User

import uuid

def createToken():
    u = uuid.uuid4()
    return u.bytes.encode("base64")[:16]

# Create your models here.
class GamePlayer(models.Model):
    user = models.OneToOneField(User)
    token = models.CharField(max_length=16)
