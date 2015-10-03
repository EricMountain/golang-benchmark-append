#!/bin/bash

set -e

readonly benchtime=3s
readonly d="results-$(date +%Y.%m.%d)_$(hostname)"

mkdir $d

for i in $(seq 10000 10000 50000) ; do echo -- Size $i ; go test -benchtime $benchtime -bench . -- $i ; done > $d/results.txt 2>&1

cd $d
echo benchtime $benchtime > benchtime.txt
sudo lshw > lshw.txt
go version > go-version.txt
../parse_results.py > plot.dat
gnuplot < ../plot.gp
cd -
