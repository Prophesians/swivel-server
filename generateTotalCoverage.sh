#!/usr/bin/env bash

export IFS=";"

coveragesum=0
totalPackages=0
allcoverages=$(go test -cover ./... | cut -d':' -f2 | awk ' {print $1} ' | tr -d ? | tr -d '\n' | tr '%' ';' | sed 's/.$//')

for word in $allcoverages; do
    pre=$(echo $word | cut -d'.' -f1)
    coveragesum=$(( $coveragesum + $pre ))
    totalPackages=$(( $totalPackages + 1 ))
done

echo "total coverage $(($coveragesum/$totalPackages))%"
