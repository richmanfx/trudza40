# -*- coding: UTF-8 -*-

import os

from fabric.api import task
from fabric.state import env
from fabric.operations import run, local, prompt, put

from fabconf import REMOTE_HOST, REMOTE_PORT, REMOTE_USER

env.hosts = REMOTE_HOST
env.port = REMOTE_PORT
env.user = REMOTE_USER
env.shell = "bash"
home_dir = os.getenv('HOME')
project_name = "trudza40"
archive_file_name = project_name + ".tar.gz"
project_dir = home_dir + os.path.sep + "go" + os.path.sep + "src" + os.path.sep + project_name


@task(default=True)
def start():
    """ Запускать всё по очереди """
    make_archive()
    send_to_remote_server()


@task()
def make_archive():
    """ Собрать архив проекта через 'bee pack' """
    print("\n ==> Make archive ...")

    local("cd " + project_dir + "; " +
          home_dir + os.path.sep + "go/bin/bee pack;" +
          "echo; " +
          "ls -la | grep '{0}'".format(archive_file_name)
          )


@task()
def send_to_remote_server():
    """ Отправить архив на дальний сервер """
    print("\n ==> Send archive to remote server ...")

    remote_home_dir = home_dir
    put(
        project_dir + os.path.sep + archive_file_name,
        remote_home_dir + os.path.sep + archive_file_name
    )


# 3.Забекапить старый проект


# 4. Остановить сервер


# 5.Удалить старый проект


# 6.Развернуть новый проект


# 7.Запустить сервер


# 8.Проверить, что сервер жив и отвечает


# 9.Послать месседж админу
