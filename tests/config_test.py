# -*- coding: utf-8 -*-

import configparser
from pathlib import Path

import pytest

from gitlabctl.config import config

__author__ = "Thomas Bianchi"
__copyright__ = "Thomas Bianchi"
__license__ = "mit"

CONTENT = """[DEFAULT]
current-context = gitTools

[gitTools]
url = https://ahahah.com
token = kkkkkk

"""

CONTENT2 = """[gitTools]
url = https://ahahah.com
token = kkkkkk
"""

CONTENT3 = """[DEFAULT]
current-context = gitTools
[gitTools]
url = https://ahahah.com
token = kkkkkk
[another]
url = https://ssss.com
token = sss
"""

CONTENT4 = """[DEFAULT]
current-context = gitTools
"""


def write_config_example(tmp_path, content):
    config_file = Path.joinpath(tmp_path, ".gitlabctl.ini")
    with open(config_file, "w") as f:
        f.write(content)
    return config_file


test_read_config_expections = [
    pytest.param(CONTENT, {
             'url': 'https://ahahah.com',
             'context': 'gitTools',
             'token': 'kkkkkk'
             }, id='single-section'),
    pytest.param(CONTENT3, {
             'url': 'https://ahahah.com',
             'context': 'gitTools',
             'token': 'kkkkkk'
             }, id='multiple-sections'),
    pytest.param(CONTENT2, None, id='no-current'),
]

test_get_contexts_expections = [
    pytest.param(CONTENT3, ["gitTools", "another"], id="two-sections"),
    pytest.param(CONTENT4, [], id="no-sections"),
]


@pytest.mark.parametrize("a,expected", test_read_config_expections)
def test_read_config(tmp_path, a, expected):
    config_file = write_config_example(tmp_path, a)
    # Re-Init config otherwise one tests collides
    config.config = configparser.ConfigParser()
    config.set_filepath(config_file)
    info = config.get_config()
    print("info:", info)
    assert info == expected


def test_set_context(tmp_path):
    # Re-Init config otherwise one tests collides
    config.config = configparser.ConfigParser()
    config.set_filepath(Path.joinpath(tmp_path, ".gitlabctl.ini"))
    config.set_context('testSet', 'https://lll', 'eeeee')
    assert 'testSet' in config.config.sections()
    assert config.config['testSet']['url'] == 'https://lll'
    assert config.config['testSet']['token'] == 'eeeee'


def test_set_current_context(tmp_path):
    # Re-Init config otherwise one tests collides
    config.config = configparser.ConfigParser()
    config.set_filepath(Path.joinpath(tmp_path, ".gitlabctl.ini"))
    config.set_current_context('currcontextA')
    assert config._get_current_context() == 'currcontextA'
    assert config.config['DEFAULT']['current-context'] == 'currcontextA'


def test_not_existing(tmp_path):
    # Re-Init config otherwise one tests collides
    config.config = configparser.ConfigParser()
    config.set_filepath(Path.joinpath(tmp_path, "NonEsiste"))
    assert not config.get_config()


def test_save(tmp_path):
    test_file = Path.joinpath(tmp_path, ".gitlabctl.ini")
    # Re-Init config otherwise one tests collides
    config.config = configparser.ConfigParser()
    config.set_filepath(test_file)
    config.config.read_string(CONTENT)
    config.save()
    with open(test_file) as f:
        assert f.read() == CONTENT


@pytest.mark.parametrize("a,expected", test_get_contexts_expections)
def test_get_contexts(a, expected):
    # Re-Init config otherwise one tests collides
    config.config = configparser.ConfigParser()
    config.config.read_string(a)
    assert config.get_contexts() == expected
