#!/bin/bash

# Script to automatically commit new LeetCode solutions
# Usage: ./commit-solution.sh "Problem Name"

cd /Users/jn/leetcode-solutions

if [ $# -eq 0 ]; then
    echo "Usage: $0 \"Problem Name\""
    echo "Example: $0 \"Two Sum\""
    exit 1
fi

PROBLEM_NAME="$1"

# Add all new files
git add .

# Check if there are changes to commit
if git diff --staged --quiet; then
    echo "No changes to commit."
    exit 0
fi

# Commit with the problem name
git commit -m "Add solution: $PROBLEM_NAME"

echo "Solution committed successfully!"
echo "Current status:"
git status --short
