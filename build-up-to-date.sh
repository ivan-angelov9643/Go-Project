#!/bin/bash

#NOT USED AT THE MOMENT

BINARY_FILE_PATH="bin/main"
TODO_APP_DIR="todo-app"

echo "Checking if binary is up to date..."
if [ ! -f "$BINARY_FILE_PATH" ]; then
    echo "Binary does not exist."
    exit 1
fi

BINARY_MTIME=$(stat -f %m "$BINARY_FILE_PATH")
NEWER_FILE_FOUND=0

for file in $(find "$TODO_APP_DIR" -type f); do
    FILE_MTIME=$(stat -f %m "$file")
    if [ "$FILE_MTIME" -gt "$BINARY_MTIME" ]; then
        echo "File $file is newer than the binary."
        NEWER_FILE_FOUND=1
        break
    fi
done

if [ "$NEWER_FILE_FOUND" -eq 1 ]; then
    echo "Binary is not up to date."
    exit 1
else
    echo "Binary is up to date."
    exit 0
fi

