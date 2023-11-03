#!/usr/bin/python3
import sys
import re
import json

"""
here, the file will get tested against all the regular expressions and findings will be outputted similar to js miner
but we need a way to organize regex patterns into categories: example below
    /secr(e|3)t=/ | CATEGORY secret/api key
    /url=http:\/\/localhost | CATEGORY hidden url endpoint
"""

def parse_json():
    # grab our json file
    with open(sys.argv[1], "r") as target_file:
        js_code = target_file.read()

    # grab our regex
    with open("regex.json", "r") as regex_file:
        categories = json.load(regex_file)

    for category, patterns in categories.items():
        for pattern in patterns:
            
            #print our matching patterns
            matches = re.findall(pattern, js_code)
            if matches:
                print(f"CATEGORY {category}: {', '.join(matches)}")

def main():
    parse_json()

if __name__=="__main__":
    main()
