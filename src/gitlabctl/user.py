# -*- coding: utf-8 -*-

from gitlabctl.config import config
import itertools
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


def do_auth(group, user, permission):
    try:
        group.members.create({'user_id': user.id, 'access_level': permission})
    except gitlab.exceptions.GitlabCreateError as e:
        if "Member already exists" in e.args:
            print("Member already exist")
        else:
            raise e


def search_user_by_username(username):
    found_list = gl.users.list(username=username)
    if len(found_list) != 1:
        raise Exception("User not found")
    return found_list[0]


def authorize(username_list, group_list, permission):
    perm_parsed = getattr(gitlab, f"{permission.upper()}_ACCESS")
    user_group = list(itertools.product(username_list, group_list))
    for couple in user_group:
        found_list = gl.users.list(username=couple[0])
        if len(found_list) != 1:
            raise Exception("Error searching for users")
        user = search_user_by_username(couple[0])
        group = gl.groups.get(couple[1])
        print(f"adding member {user.name} with permission {permission} \
to group {group.name}")
        do_auth(group, user, perm_parsed)


def block(username_list):
    for username in username_list:
        user = search_user_by_username(username)
        print(f"Blocking user {user.name}")
        user.block()


def main(fn, *args):
    cfg = config.get_config()
    global gl
    gl = gitlab.Gitlab(cfg['url'], private_token=cfg['token'], per_page=50)
    return fn(*args)
