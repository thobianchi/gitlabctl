# -*- coding: utf-8 -*-

import string
import random
from gitlabctl.config import config
import gitlab

__author__ = "Thomas Bianchi"
__copyright__ = "Thomas Bianchi"
__license__ = "mit"

gl = None
PASSWORD_LENGHT = 20


def create_user(email, username, name):
    chars = string.ascii_letters + string.digits
    password = "".join(random.sample(chars, PASSWORD_LENGHT))
    gl.users.create({'email': email,
                     'password': password,
                     'username': username,
                     'name': name})
    print(f"User {email} created with password: {password}")


def main(fn, *args):
    cfg = config.get_config()
    global gl
    gl = gitlab.Gitlab(cfg['url'], private_token=cfg['token'], per_page=50)
    return fn(*args)
