#!/bin/sh

OUTPUT=plot.tmp
DIR="$(cd $(dirname "$0"); pwd)"
echo 'training test' > "$OUTPUT"
while read line; do
    echo "$line" >> "$OUTPUT"
done
gnuplot -p \
        -e 'set key outside' \
        -e 'set yrange [0:1]' \
        -e "plot for [col=1:2] '$OUTPUT' using col with lines title columnheader(col)"
rm "$OUTPUT"
