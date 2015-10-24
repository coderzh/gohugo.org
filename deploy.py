#!/usr/bin/env python
# coding:utf-8

import os
import sys
import glob
import shutil
import subprocess

__author__ = 'coderzh'


class ChDir:
    """Context manager for changing the current working directory"""
    def __init__(self, new_path):
        self.newPath = os.path.expanduser(new_path)

    def __enter__(self):
        self.savedPath = os.getcwd()
        os.chdir(self.newPath)

    def __exit__(self, exception_type, exception_value, traceback):
        os.chdir(self.savedPath)


def deploy():
    current_dir = os.path.dirname(os.path.abspath(__file__))
    parent_dir = os.path.abspath(os.path.join(current_dir, '..'))

    with ChDir(current_dir):
        # step1 clean
        os.system('git fetch origin')
        os.system('git checkout master')
        os.system('git reset --hard origin/master')
        os.system('git clean -fdx')
        # step2 build
        os.system('hugo -v -cacheDir="./cache"')

    deploy_dir = os.path.join(parent_dir, 'deploy')

    # step3 create if not exists
    if not os.path.exists(deploy_dir):
        os.makedirs(deploy_dir)

    with ChDir(deploy_dir):
        # step4 remove all files
        for f in os.listdir('.'):
            if f != 'index.html':
                if os.path.isfile(f):
                    os.remove(f)
                elif os.path.isdir(f):
                    shutil.rmtree(f)

        # step5 copy new files
        from_dir = os.path.join(current_dir, 'public')
        for f in os.listdir(from_dir):
            file_path = os.path.join(from_dir, f)
            if os.path.isfile(file_path):
                shutil.copy(file_path, '.')
            elif os.path.isdir(file_path):
                shutil.copytree(file_path, f)

if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == '--auto':
        deploy()
    else:
        print 'please use --auto flag.'

