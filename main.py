#!/usr/bin/python3
import sys

with open(sys.argv[1], "r") as target_file:
    js_code = target_file.read()

with open("regex.txt", "r") as regular_expression_list:
    regular_expressions = regular_expression_list.readlines()

print(len(regular_expressions))

"""
here, the file will get tested against all the regular expressions and findings will be outputted similar to js miner
but we need a way to organize regex patterns into categories: example below
    /secr(e|3)t=/ | CATEGORY secret/api key
    /url=http:\/\/localhost | CATEGORY hidden url endpoint
"""