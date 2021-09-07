# -*- coding: utf-8 -*-

import os
import subprocess
import sys
import tempfile
from gitlabctl.config import config
import gitlab

__author__ = "Thomas Bianchi"
__copyright__ = "Thomas Bianchi"
__license__ = "mit"

gl = None


def get_all_parents_ids(id):
    """
    returns a list of ids of all parents groups from project id
    """
    project = gl.projects.get(id)
    parent_id = project.namespace.get('id')
    ids = [parent_id]
    while True:
        group = gl.groups.get(parent_id)
        if not group.parent_id:
            break
        parent_id = group.parent_id
        ids.append(parent_id)
    return ids


def get_variable_list(proj):
    variables = []
    variables.extend(proj.variables.list(as_list=False))
    ids = get_all_parents_ids(proj.id)
    for id in ids:
        grp = gl.groups.get(id)
        variables.extend(grp.variables.list(as_list=False))
    variables.extend(gl.variables.list(as_list=False))
    return variables


def create_temp_file(value):
    fd, f_name = tempfile.mkstemp(prefix="gitlabctl.")
    with os.fdopen(fd, 'w') as fdfile:
        fdfile.write(value)
    return f_name


def create_export_string(key, value):
    return "export {}='{}'".format(key, value)


def format_variables(vars_list):
    variables = []
    for v in vars_list:
        if v.variable_type == "file":
            tmp_file = create_temp_file(v.value)
            variables.append(create_export_string(v.key, tmp_file))
        else:
            variables.append(create_export_string(v.key, v.value))
    return variables


def get_project_id_from_git_remote():
    run = subprocess.run(["git", "config", "--get", "remote.origin.url"],
                         capture_output=True, text=True)
    if not run.stdout:
        print("Remote not found")
        sys.exit(1)
    ns_project = run.stdout.split(":")[-1].split(".")[0]
    return gl.projects.get(ns_project)


def get_current_branch():
    run = subprocess.run(["git", "rev-parse", "--abbrev-ref", "HEAD"],
                         capture_output=True, text=True)
    if "fatal: not a git repository " in run.stderr:
        print(run.stderr)
        sys.exit(1)
    return run.stdout.strip('\n')


# Entry point from CLI
def get_env(id):
    proj = None
    if id:
        proj = gl.projects.get(id)
    else:
        proj = get_project_id_from_git_remote()
    vars = get_variable_list(proj)
    return format_variables(vars)


# Entry point from CLI
def run_pipeline(env_vars):
    proj = get_project_id_from_git_remote()
    branch = get_current_branch()
    proj.pipelines.create({'ref': branch, 'variables': env_vars})


def main(fn, *args):
    cfg = config.get_config()
    global gl
    gl = gitlab.Gitlab(cfg['url'], private_token=cfg['token'], per_page=50)
    return fn(*args)
