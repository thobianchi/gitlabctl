# -*- coding: utf-8 -*-
import click

import gitlabctl.project as gitlab_project
import gitlabctl.user as gitlab_user
from gitlabctl.config import config

from gitlabctl import __version__

__author__ = "Thomas Bianchi"
__copyright__ = "Thomas Bianchi"
__license__ = "mit"


def set_config(ctx, param, value):
    if not value or ctx.resilient_parsing:
        return
    config.set_filepath(click.format_filename(value))


@click.group()
@click.version_option(version=__version__)
@click.option("--conf", type=click.Path(exists=True), callback=set_config,
              is_eager=True, help="Set configuration file fullpath.")
def cli(conf):
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
    in form VAR1=VALUE VAR2=VALUE.
    """
    d = [{'key': a.split('=')[0], 'value': a.split('=')[1]} for a in vars]
    for _, v in d:
        if any(elem in v for elem in r", "):
            # FIXME non working
            raise click.BadArgumentUsage("parameters should be in form 'VAR=1 RAV=2'")
    gitlab_project.main(gitlab_project.run_pipeline, d)


@cli.group()
def user():
    """Manages user."""
    pass


@user.command("create")
@click.option("-e", "--email", "email", required=True, type=str)
@click.option("-u", "--username", "username", required=True, type=str)
@click.option("-n", "--name", "name", required=True, type=str)
def create_user(email, username, name):
    """
    Create user
    """
    gitlab_user.main(gitlab_user.create_user, email, username, name)
