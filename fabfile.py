# -*- coding: UTF-8 -*-

import os
import time
import datetime

from fabric.api import task
from fabric.state import env
from fabric.operations import run, local, prompt, put

from fabconf import REMOTE_HOST, REMOTE_PORT, REMOTE_USER

env.hosts = REMOTE_HOST
env.port = REMOTE_PORT
env.user = REMOTE_USER
app_port = 8081
home_dir = os.getenv('HOME')
project_name = "trudza40"
now = datetime.datetime.now()
current_date = str(now.strftime("%Y%m%d"))
archive_file_name = project_name + ".tar.gz"
project_dir = home_dir + os.path.sep + "go" + os.path.sep + "src" + os.path.sep + project_name
server_os = "freebsd"


@task(default=True)
def start():
    """ Run all in order """
    make_archive()
    send_to_remote_server()
    make_backup()
    app_stop()

    # Не работает никак -
    # app_start()

    delete_old_app()
    deploy_new_app()

    # check_app_running()


@task()
def make_archive():
    """ Make an archive of the project through 'bee pack' """
    print("\n ==> Make archive ...")

    local("cd " + project_dir + "; " +
          home_dir + os.path.sep + "go/bin/bee pack -be=GOOS=" + server_os + ";" +
          "echo; " +
          "ls -la | grep '{0}'".format(archive_file_name)
          )


@task()
def send_to_remote_server():
    """ Send archive file on the remote server """
    print("\n ==> Send archive to remote server ...")

    remote_home_dir = home_dir
    put(
        project_dir + os.path.sep + archive_file_name,
        remote_home_dir + os.path.sep + archive_file_name
    )


@task()
def make_backup():
    """ Back up application on remote server """
    print("\n ==> Мake a backup application on remote server ...")

    remote_backup_dir = home_dir + os.path.sep + "backups" + os.path.sep + project_name + ".ru"
    remote_project_dir = "/usr/local/www/html" + os.path.sep + project_name + "_ru"

    run("/bin/mkdir -p -- " + remote_backup_dir + os.path.sep + current_date)

    run("/usr/bin/tar cvfz - " + remote_project_dir + " > " +
        remote_backup_dir + os.path.sep + current_date + os.path.sep + archive_file_name)
    print("Backup complete")


@task()
def app_stop():
    """ Stopping application on remote server """
    print("\n ==> Stopping application on remote server ...")
    run("/usr/local/sbin/stop_" + project_name)
    time.sleep(5)
    print("Stopping application complete")


@task()
def delete_old_app():
    """ Deleting old application on remote server """
    print("\n ==> Deleting old application on remote server ...")
    run("/usr/local/sbin/delete_" + project_name)
    print("Deleting application complete")


@task()
def deploy_new_app():
    """ Deploy new application on remote server """
    print("\n ==> Deploy new application on remote server ...")
    run("/usr/local/sbin/deploy_" + project_name)
    print("Deploy complete")


@task()
def app_start():
    """ Starting application on remote server """
    print("\n ==> Starting application on remote server ...")
    run("/usr/local/sbin/start_" + project_name + " &")
    time.sleep(5)
    print("Starting application complete")


@task()
def check_app_running():
    """ Check that application is running """
    print("\n ==> Check that application is running ...")
    output = run("netstat -naf inet")

    if str(app_port) in output:
        print("\nApplication started on remote server normally")
    else:
        print("\nError: Application not started on remote server!!!")



# 9.Послать месседж админу
