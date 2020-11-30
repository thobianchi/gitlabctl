# -*- coding: utf-8 -*-

import gitlab

__author__ = "Thomas Bianchi"
__copyright__ = "Thomas Bianchi"
__license__ = "mit"

PERPAGEOPT = 50


class Gitlab_client(object):
    def __init__(self, url="", token=""):
        self.gl = gitlab.Gitlab(url, private_token=token, per_page=PERPAGEOPT)

    def get_project_by_id(self, id):
        return self.gl.projects.get(id)

    def get_group_by_id(self, id):
        return self.gl.groups.get(id)

    def get_environemnt_vars(self, prjgrp):
        """
        Given a project or group object return a list of all env vars
        """
        return prjgrp.variables.list(as_list=False)

    def get_instance_vars(self):
        return self.gl.variables.list(as_list=False)

    def search_project(self, keyword):
        return self.gl.projects.list(search=keyword, as_list=False)
