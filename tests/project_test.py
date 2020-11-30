# -*- coding: utf-8 -*-

from gitlabctl.client import Gitlab_client
from gitlabctl.project import get_all_parents_ids

__author__ = "Thomas Bianchi"
__copyright__ = "Thomas Bianchi"
__license__ = "mit"


class Mock_project():
    id = None
    namespace = {'id': None}


class Mock_group():
    id = None
    parent_id = None


a_proj = Mock_project()
a_proj.id = 112
a_proj.namespace = {'id': 332}

b_proj = Mock_project()
b_proj.id = 645
b_proj.namespace = {'id': 554}

a_grp = Mock_group()
a_grp.id = 332
a_grp.parent_id = 554

b_grp = Mock_group()
b_grp.id = 554
b_grp.parent_id = None


def get_proj_by_id(self, id):
    for prj in self.mock_projects:
        if prj.id == id:
            return prj


def get_group_by_id(self, id):
    for grp in self.mock_groups:
        if grp.id == id:
            return grp


def test_get_all_parents_ids(mocker):

    mocker.patch(
        'gitlabctl.client.Gitlab_client.get_project_by_id',
        get_proj_by_id
    )

    mocker.patch(
        'gitlabctl.client.Gitlab_client.get_group_by_id',
        get_group_by_id
    )

    mocker.patch(
        'gitlabctl.client.Gitlab_client.mock_projects',
        [a_proj, b_proj], create=True
    )

    mocker.patch(
        'gitlabctl.client.Gitlab_client.mock_groups',
        [a_grp, b_grp], create=True
    )

    client = Gitlab_client()
    grp_ids = get_all_parents_ids(client, 112)
    assert grp_ids == [332, 554]
