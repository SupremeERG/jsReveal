#!/usr/bin/python3
import sys

with open(sys.argv[1], "r") as target_file:
    print(target_file.read())