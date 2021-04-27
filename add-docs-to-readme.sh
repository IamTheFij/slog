#! /bin/bash
set -e

slogdir=$(dirname "$0")
readme="$slogdir/README.md"

awk '/## Documentation/ {print ; exit} {print}' "$readme" > "$readme.tmp" && go doc -all slog/v2 | sed "s/^/    /;s/[ \t]*$//" >> "$readme.tmp"
mv "$readme.tmp" "$readme"
