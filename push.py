#!/usr/bin/env python
# coding:utf-8

import os


def main():
    os.system('git add -A')
    os.system('git commit')
    os.system('git fetch origin')
    os.system('git rebase origin/master')
    os.system('git push origin master')

if __name__ == '__main__':
    main()

