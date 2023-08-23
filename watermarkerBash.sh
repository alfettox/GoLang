#!/bin/bash

# Your name (change this to your actual name)
my_name="Giovanni De Franceschi"

# Find all Go files in the current directory and its subdirectories
find . -type f -name "*.go" | while read file; do
    # Remove existing comments at the beginning of the file
    sed -i '/^\/\*/{:a;N;/\*\//!ba}; s/\/\*.*\*\///' "$file"

    # Replace $your_name with $my_name in the file
    sed -i "s/\$your_name/$my_name/g" "$file"

    # Insert your name as a comment at the top of the file
    sed -i "1i /*\nAuthor: $my_name\n*/" "$file"

    echo "Watermarked: $file"
done

echo "All watermarked."
