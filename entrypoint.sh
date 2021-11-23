#!/bin/sh

set -e

result="$1"
url="$2"
command="$3"

echo "$result"
echo "$url"
echo "$command"

chmod +x /notfli

sh -c "/notfli -result $result -url $url $command"
