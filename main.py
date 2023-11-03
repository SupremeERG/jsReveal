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

def parse_js():
    # grab our js file
    with open(sys.argv[1], "r") as target_file:
        js_code = target_file.read()

    # grab our regex
    with open("regex.json", "r") as regex_file:
        categories = json.load(regex_file)
        #print(categories)

    for pattern, regex_properties in categories.items():
        # run regex through transformations based on JSON file properties
        flags = 0
        if regex_properties["match_line"] == True:
            pattern = f"{pattern}.*(?:\n|$)"
            #flags |= re.DOTALL
        if regex_properties["case_sensitive"] == True:
            flags |= re.IGNORECASE #flags.append(re.IGNORECASE)
        pattern = re.compile(pattern, flags)


        matches = re.findall(pattern, js_code)
        if matches:
            for match in matches:
                print(f'Category {regex_properties["type"]}\nString: {match[:150]}') # limiting this to 150 characters because if the code is obfuscated it will print a giant ass unreadable block
                print("\n\n") # just for extra space, so i can read easier while debugging
        
def main():
    parse_js()

if __name__=="__main__":
    main()
