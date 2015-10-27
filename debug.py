#!/usr/bin/env python
# coding:utf-8

import os
import shutil


def main():
    if os.path.exists('public'):
        shutil.rmtree('public')

    os.system('hugo server -w')


if __name__ == '__main__':
    main()
