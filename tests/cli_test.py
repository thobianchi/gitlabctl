# -*- coding: utf-8 -*-

import click
import pytest

from click.testing import CliRunner

from gitlabctl.cli import project_get_env
from gitlabctl.cli import run_pipeline

__author__ = "Thomas Bianchi"
__copyright__ = "Thomas Bianchi"
__license__ = "mit"


def main_get_env(func_name, id):
    return [id]


def main_run_pipeline(func_name, d):
    click.echo(d)


get_env_by_id_expections = [
    pytest.param(['--by-id', '1123'], '1123\n', id="full"),
    pytest.param(None, '\n', id="no-id"),
]

run_pipeline_expections = [
    pytest.param(['NOPROD=1'], "[{'key': 'NOPROD', 'value': '1'}]\n",
                 id="spaced single param"),
    pytest.param(['NOPROD=1', 'PROVA=2'],
                 "[{'key': 'NOPROD', 'value': '1'}, {'key': 'PROVA', 'value': '2'}]\n",
                 id="spaced multiple params"),
    # pytest.param(['NOPROD=1,PROVA=2'], pytest.raises(click.BadArgumentUsage),
    #              id="non spaced"),
]


@pytest.mark.parametrize("a,expected", get_env_by_id_expections)
def test_project_get_env(mocker, a, expected):
    mocker.patch(
        'gitlabctl.project.main',
        main_get_env)

    runner = CliRunner()
    result = runner.invoke(project_get_env, a)
    assert expected == result.output


@pytest.mark.parametrize("a,expected", run_pipeline_expections)
def test_run_pipeline(mocker, a, expected):
    mocker.patch(
        'gitlabctl.project.main',
        main_run_pipeline)
    runner = CliRunner()
    result = runner.invoke(run_pipeline, a)
    assert expected == result.output
