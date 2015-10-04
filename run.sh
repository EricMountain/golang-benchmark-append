#!/bin/bash

set -e

# GC setting: default is 100 (https://golang.org/pkg/runtime/); off to disable.
readonly gc=off
readonly benchtime=1s

readonly d="results-$(date +%Y.%m.%d)_$(hostname)"

mkdir $d

for i in $(seq 10000 10000 10000000) ; do
    echo -- Size $i
    for b in IntIdx IntAppend PoloIdx PoloAppend ; do
        GOGC=$gc go test -benchtime $benchtime -bench $b -- $i
    done
done > $d/results.txt 2>&1

cd $d
echo GOGC=$gc > gc.txt
echo benchtime $benchtime > benchtime.txt
sudo lshw > lshw.txt
go version > go-version.txt
../parse_results.py > plot.dat
gnuplot < ../plot.gp
cd -
