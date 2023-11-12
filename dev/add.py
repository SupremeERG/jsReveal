import json, sys
from os import path

newRegex = sys.argv[1]
regexFile = './regex.json'
regexObj = []

if path.isfile(regexFile) == False:
    raise Exception("regex file not found " + f"'{regexFile}'")

with open(regexFile) as regexfileobject:
    
    regexObj = json.load(regexfileobject)


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


