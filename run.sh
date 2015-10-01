#!/bin/bash

set -e

d="results-$(date +%Y.%m.%d)_$(hostname)"

mkdir $d

for i in $(seq 10000 10000 50000) ; do echo -- Size $i ; go test -bench . -- $i ; done > $d/results.txt 2>&1

cd $d
../parse_results.py > plot.dat
gnuplot < ../plot.gp
cd -
