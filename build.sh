#!/bin/bash

# Set the source and destination directories
SOURCE_DIR="./backend"
DEST_DIR="./build"

# Create the destination directory if it doesn't exist
mkdir -p "$DEST_DIR"

# Remove all files in the destination directory
rm -rf "$DEST_DIR"/*

# Find all files in the source directory
find "$SOURCE_DIR" -type f | while read -r file; do
    # Get the relative path from the source directory
    relative_path="${file#"$SOURCE_DIR/"}"

    # Replace slashes with underscores
    new_filename=$(echo "$relative_path" | tr '/' '_')

    # Create the full destination path
    new_file_path="$DEST_DIR/$new_filename"

    # Create the destination directory if it doesn't exist
    mkdir -p "$(dirname "$new_file_path")"

    # Copy the file to the new location
    cp "$file" "$new_file_path"
done