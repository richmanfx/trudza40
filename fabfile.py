# -*- coding: UTF-8 -*-

import os
import getpass

from fabric.api import task
from fabric.operations import run, local, prompt, put
from fabric.state import env

env.hosts = 'r5am.ru'
env.port = '35712'
env.user = "zoer" 
# env.password = "123"
env.shell = "bash"
home_dir = os.getenv('HOME')
project_dir = home_dir + os.path.sep + "go" + os.path.sep + "src" + os.path.sep + "trudza40"


@task(default=True)
def start():
    make_archive()


@task()
def make_archive():
    """ 1.Собрать архив проекта через bee pack """
    print("\n ==> Make archive ...")
    
    local("cd " + project_dir + "; " + 
            home_dir + os.path.sep + "go/bin/bee pack;" + 
            "echo; " + 
            "ls -la | grep 'tar.gz'"
    )
    

# 2.Отправить архив на дальний сервер


# 3.Забекапить старый проект


# 4. Остановить сервер


# 5.Удалить старый проект


# 6.Развернуть новый проект


# 7.Запустить сервер


# 8.Проверить, что сервер жив и отвечает


# 9.Послать месседж админу


