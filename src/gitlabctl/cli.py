# -*- coding: utf-8 -*-
import click

import gitlabctl.project as gitlab_project
from gitlabctl import __version__
from gitlabctl.client import Gitlab_client
from gitlabctl.config import config

__author__ = "Thomas Bianchi"
__copyright__ = "Thomas Bianchi"
__license__ = "mit"


@click.group()
@click.version_option(version=__version__)
def cli():
    """Gitlab CLI
    Interacts with a gitlab installation: gets environment of a project or launch
    pipeline and see output.
    """


@cli.group()
def project():
    """Manages project."""


@project.command("get-env")
@click.option("--by-id", type=int)
def project_get_env(by_id):
    """
    Get project and anchestor environemnt and print export statements.
    """
    cfg = config.get_config()
    client = Gitlab_client(cfg['url'], cfg['token'])
    vars = gitlab_project.get_env(client, by_id)
    for v in vars:
        click.echo(v)
