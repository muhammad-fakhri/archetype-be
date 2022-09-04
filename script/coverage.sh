#!/bin/bash

set -e

output=$1
echo 'mode: count' > profile.cov

for dir in $(find . -maxdepth 10 -not -path './.git*' -not -path '*/_*' -not -path './vendor*' -not -path './tests*' -type d);
do
if ls $dir/*.go &> /dev/null; then
    go test -short -covermode=count -coverprofile=$dir/profile.tmp $dir
    if [ -f $dir/profile.tmp ]
    then
        cat $dir/profile.tmp | tail -n +2 >> profile.cov
        rm $dir/profile.tmp
    fi
fi
done

if [ "$output" = "html" ]; then
    go tool cover -html=profile.cov 
else
    go tool cover -func profile.cov
fi
