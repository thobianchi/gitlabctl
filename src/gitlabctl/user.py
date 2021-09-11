# -*- coding: utf-8 -*-

from gitlabctl.config import config
import gitlab

__author__ = "Thomas Bianchi"
__copyright__ = "Thomas Bianchi"
__license__ = "mit"

gl = None


def create_user(email, username, name):
    gl.users.create({'email': email,
                     'username': username,
                     'name': name,
                     'reset_password': True})  # Send user password reset link
    print(f"User {email} created")


def main(fn, *args):
    cfg = config.get_config()
    global gl
    gl = gitlab.Gitlab(cfg['url'], private_token=cfg['token'], per_page=50)
    return fn(*args)
