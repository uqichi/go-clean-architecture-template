#!/bin/bash -x

echo "# Changes"

LAST_RELEASE_TAG=$(git tag -l 'v[0-9].[0-9].[0-9]' --sort=-v:refname | head -1)

RELEASE_LIST=$(git log --oneline --no-merges "${LAST_RELEASE_TAG}..master")

IFS=$'\n'
for line in $(echo "$RELEASE_LIST"); do
  echo "- $line"
done

echo "\nDiff: https://github.com/alumni-inc/delhi-server/compare/${LAST_RELEASE_TAG}...master"
