# -*- coding: utf-8 -*-

import pytest

from click.testing import CliRunner

from gitlabctl.cli import project_get_env

__author__ = "Thomas Bianchi"
__copyright__ = "Thomas Bianchi"
__license__ = "mit"


def mock_get_env(client, id):
    if not id:
        print('None')
        return
    print(id)


get_env_by_id_expections = [
    pytest.param(['--by-id', '1123'], '1123\n', id="full"),
    pytest.param(None, 'None\n', id="None"),
]


@pytest.mark.parametrize("a,expected", get_env_by_id_expections)
def test_get_env_by_id(mocker, a, expected):
    mocker.patch(
        'gitlabctl.project.get_env',
        mock_get_env)

    runner = CliRunner()
    result = runner.invoke(project_get_env, a)
    assert expected == result.output
