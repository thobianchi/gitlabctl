# -*- coding: utf-8 -*-

import os
import subprocess
import sys
import tempfile

from iterfzf import iterfzf

__author__ = "Thomas Bianchi"
__copyright__ = "Thomas Bianchi"
__license__ = "mit"


def get_all_parents_ids(client, id):
    """
    returns a list of ids of all parents groups from project id
    """
    project = client.get_project_by_id(id)
    parent_id = project.namespace.get('id')
    ids = [parent_id]
    while True:
        group = client.get_group_by_id(parent_id)
        if not group.parent_id:
            break
        parent_id = group.parent_id
        ids.append(parent_id)
    return ids


# get all vars
def get_variable_list(client, id):
    variables = []
    proj = client.get_project_by_id(id)
    variables.extend(client.get_environemnt_vars(proj))
    ids = get_all_parents_ids(client, id)
    for id in ids:
        grp = client.get_group_by_id(id)
        variables.extend(client.get_environemnt_vars(grp))
    variables.extend(client.get_instance_vars())
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


def yield_project_with_ns(search_list):
    for proj in search_list:
        yield "{} {}".format(proj.id, proj.name_with_namespace)


def multiple_projects_found(search_list):
    return iterfzf(yield_project_with_ns(search_list)).split(" ")[0]


def get_project_id_from_git_remote(client):
    run = subprocess.run(["git", "config", "--get", "remote.origin.url"],
                         capture_output=True, text=True)
    if not run.stdout:
        print("Remote not found")
        sys.exit(1)
    repo_name = run.stdout.split("/")[-1].split(".")[0]
    search_list = client.search_project(repo_name)
    if len(search_list) == 1:
        return search_list[0].id
    else:
        return multiple_projects_found(search_list)


def get_env(client, id):
    if not id:
        id = get_project_id_from_git_remote(client)
    vars = get_variable_list(client, id)
    return format_variables(vars)
