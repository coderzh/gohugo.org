#!/usr/bin/env python
# coding:utf-8

import os
import sys
import glob
import locale
import shutil
import argparse
import subprocess

__author__ = 'coderzh'

GIT_REPO = [
    ['origin',  'gh-pages', 'git@github.com:coderzh/gohugo-deployed.git'],
    ['gitcafe', 'gitcafe-pages', 'git@gitcafe.com:coderzh/gohugo.org.git'],
]

DEPLOY_DIR = 'deploy'


class ChDir:
    """Context manager for changing the current working directory"""
    def __init__(self, new_path):
        self.newPath = os.path.expanduser(new_path)

    def __enter__(self):
        self.savedPath = os.getcwd()
        os.chdir(self.newPath)

    def __exit__(self, exception_type, exception_value, traceback):
        os.chdir(self.savedPath)


def deploy(args):
    current_dir = os.path.dirname(os.path.abspath(__file__))
    parent_dir = os.path.abspath(os.path.join(current_dir, '..'))
    public_dir = os.path.join(current_dir, 'public')
    commit_msg = ''

    with ChDir(current_dir):
        # step1 clean
        if os.path.exists(public_dir):
            shutil.rmtree(public_dir)

        if args.type == 'auto':
            subprocess.call('git fetch origin', shell=True)
            subprocess.call('git checkout master', shell=True)
            subprocess.call('git reset --hard origin/master', shell=True)
            subprocess.call('git clean -fdx', shell=True)
            
        # on windows set TERM=msys
        s = subprocess.Popen('git log -1 --pretty=format:"%s"',
                             shell=True, stdout=subprocess.PIPE)
        commit_msg = s.communicate()[0].decode('utf-8').encode(locale.getpreferredencoding())
        # step2 build
        subprocess.call('hugo -v --cacheDir="./cache"', shell=True)

    deploy_dir = os.path.join(parent_dir, DEPLOY_DIR)

    # step3 clone if not exists
    if not os.path.exists(deploy_dir):
        os.makedirs(deploy_dir)
        with ChDir(deploy_dir):
            subprocess.call('git init', shell=True)
            for repo in GIT_REPO:
                subprocess.call('git remote add {0} {1}'.format(repo[0], repo[2]), shell=True)
    elif args.type == 'first':
        with ChDir(deploy_dir):
            subprocess.call('git init', shell=True)
            for repo in GIT_REPO:
                subprocess.call('git remote add {0} {1}'.format(repo[0], repo[2]), shell=True)

    with ChDir(deploy_dir):
        # step4 clean and pull
        if len(GIT_REPO) > 0:
            subprocess.call('git fetch {0}'.format(GIT_REPO[0][0]), shell=True)
            if args.type == 'first':
                subprocess.call('git checkout --orphan temp', shell=True)
                subprocess.call('git rm --cached -r .', shell=True)
                subprocess.call('git clean -fdx', shell=True)
                subprocess.call('git branch -D {0}'.format(GIT_REPO[0][1]), shell=True)
                subprocess.call('git checkout -b {0}'.format(GIT_REPO[0][1]), shell=True)
            else:
                subprocess.call('git checkout {0}'.format(GIT_REPO[0][1]), shell=True)

            subprocess.call('git reset --hard {0}/{1}'.format(GIT_REPO[0][0], GIT_REPO[0][1]), shell=True)
            subprocess.call('git clean -fdx', shell=True)

        # step5 remove all files
        for f in os.listdir('.'):
            if f != '.git':
                if os.path.isfile(f):
                    os.remove(f)
                elif os.path.isdir(f):
                    shutil.rmtree(f)

        # step6 copy new files
        for f in os.listdir(public_dir):
            file_path = os.path.join(public_dir, f)
            if os.path.isfile(file_path):
                shutil.copy(file_path, '.')
            elif os.path.isdir(file_path):
                shutil.copytree(file_path, f)

        # step7 commit and push
        if len(GIT_REPO) > 0:
            subprocess.call('git add --all', shell=True)
            subprocess.call('git commit -a -m "{0}"'.format(commit_msg), shell=True)
            for repo in GIT_REPO:
                if args.test:
                    print('git push {0} {1}:{2} -u'.format(repo[0], GIT_REPO[0][1], repo[1]))
                else:
                    subprocess.call('git push {0} {1}:{2} -u'.format(repo[0], GIT_REPO[0][1], repo[1]), shell=True)

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='deploy hugo')
    parser.add_argument('type', help='auto or manual')
    parser.add_argument('-t', dest='test', action='store_true', help='for test')
    args = parser.parse_args()

    if args.type in ['auto', 'manual', 'first']:
        deploy(args)

