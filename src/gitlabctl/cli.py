# -*- coding: utf-8 -*-
import click

import gitlabctl.project as gitlab_project
from gitlabctl import __version__

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
    Get anchestor environemnt and print export statements.
    """
    vars = gitlab_project.main(gitlab_project.get_env, by_id)
    for v in vars:
        click.echo(v)


@project.command("run-pipeline")
@click.argument("vars", nargs=-1)
def run_pipeline(vars):
    """
    Run pipeline for current checked-out branch passing vars
    in form VAR1=VALUE VAR2=VALUE
    """
    d = [{'key': a.split('=')[0], 'value': a.split('=')[1]} for a in vars]
    gitlab_project.main(gitlab_project.run_pipeline, d)
