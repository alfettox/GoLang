#!/bin/bash
new_files=$(git status -s | grep '^??' | sed 's/^?? //')

last_commit=$(git log --oneline | wc -l | tr -d ' ')
next_commit=$((last_commit + 1))
git add .
commit_message="$next_commit: $commit_message ($new_files)"
git commit -m "$commit_message"
git push

echo "Committed to github"

