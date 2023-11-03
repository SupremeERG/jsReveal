#!/usr/bin/python3
import sys
import re
import json

"""
here, the file will get tested against all the regular expressions and findings will be outputted similar to js miner
"""



def compile_pattern(pattern, regex_properties):
        flags = 0
        valid_pattern = pattern
        if regex_properties["match_line"] == True:
            valid_pattern = f"{valid_pattern}.*(?:\n|$)"
            #flags |= re.DOTALL
        if regex_properties["case_insensitive"] == True:
            flags |= re.IGNORECASE #flags.append(re.IGNORECASE)

        try:
            return re.compile(valid_pattern, flags)
        except re.error:
            raise ValueError('Invalid Regular Expression: "{}"'.format(pattern))
            

def parse_js():
    # grab our js file
    with open(sys.argv[1], "r") as target_file:
        js_code = target_file.read()

    # grab our regex
    with open("regex.json", "r") as regex_file:
        categories = json.load(regex_file)
        #print(categories)

    for pattern, regex_properties in categories.items():
        pattern = compile_pattern(pattern, regex_properties)

        matches = re.findall(pattern, js_code)
        if matches:
            for match in matches:
                if len(match) > 1000:
                    match = match[:250] # prevents humungous blocks of minified code from being outputted
                print(f'Category: {regex_properties["type"]}\nString: {match}')
        
def main():
    parse_js()

if __name__=="__main__":
    main()
