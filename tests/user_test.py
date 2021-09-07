# -*- coding: utf-8 -*-

from gitlabctl.user import create_user

__author__ = "Thomas Bianchi"
__copyright__ = "Thomas Bianchi"
__license__ = "mit"


class Users:
    def create(self, a):
        pass


class Gl:
    users = Users()


def test_create_user(mocker):
    mocker.patch(
        "gitlabctl.user.gl",
        Gl
    )
    create_user("john.doe@example.com", "john.doe", "John Doe")
    assert True
