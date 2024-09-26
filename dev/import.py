#!/usr/bin/python3

# Imports regex patterns from text file into db file

import sys

regexFile = sys.argv[1]
dbFile = sys.argv[2]
regexType = sys.argv[3]

confirmation = input(f"Are you sure you want to import '{regexType}' regexes from {regexFile} into {dbFile}? [y/n]")
if confirmation != "y":
    print("Cancelling")
    exit()

print("finished")

