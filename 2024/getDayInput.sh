#!/usr/sh

# Usage: ./getDayInput.sh <day_number>
# Ensure that you have a session.txt file in the same directory as this script
# Should get from adventofcode website

if [ -z "$1" ]
then
    echo "Please provide the day number"
    exit 1
fi

if [ -e "session.txt" ]
then
    echo "Session file exists"
else
    echo "Session file does not exist"
    exit 1
fi

# check if input is a number
if [ "$1" -eq "$1" ] 2>/dev/null
then
    echo "Day number is a number"
else
    echo "Day number should be a number"
    exit 1
fi

# check if input is between 1 and 25
if [ "$1" -lt 1 ] || [ "$1" -gt 25 ]
then
    echo "Day number should be between 1 and 25"
    exit 1
fi

dayDir=$(printf "inputs/day%02d" "$1")

mkdir -p "$dayDir"

curl -X GET https://adventofcode.com/2024/day/"$1"/input -o "$dayDir/input.txt" --cookie "$(cat session.txt)"
