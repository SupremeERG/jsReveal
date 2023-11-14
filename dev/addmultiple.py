import json, sys
from os import path

# this should add all regexes from a file (the file that you've been using for testing)


sourceFile = './regex.txt' # or the file with your regexpressions in a .txt file
regexFile = './regex.json' # the file with formatted, organized regex objects


with open(sourceFile, "r") as filewithregexes:
    regexes = filewithregexes.readlines()

if path.isfile(regexFile) == False:
    raise Exception("regex file not found " + f"'{regexFile}'")

with open(sourceFile, "r") as filewithregexes: # get source regexes
    regexes = filewithregexes.readlines()

with open(regexFile) as regexfileobject: # get the json object
    
    regexObj = json.load(regexfileobject)

for newRegex in regexes:
    regexObj.update({
        newRegex: {
            "type": "AddedFromAdd.py",
            "case_insensitive": False,
            "confidence": "low",
            "match_line": False
        }
    })

with open(regexFile, 'w') as regexfileobject:
    json.dump(regexObj, regexfileobject, indent=4, separators=(',',': '))


