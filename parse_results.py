#!/usr/bin/env python3

import re

def parseBench(line):
    #BenchmarkDoIdx-2            1000           1967567 ns/op
    match = re.fullmatch('Benchmark.+-.\s+\d+\s+(\d+) ns/op\n', line)
    if match is not None:
        return int(match.group(1))


t = dict()
with open('results.txt', 'r') as f:
    for line in f:
        if line.startswith('-- Size '):
            size = int(line[len('-- Size '):-1])
            r = dict()
            t[size] = r
            continue
        if line.startswith('PASS') or line.startswith('ok'):
            continue
        if line.startswith('BenchmarkDoIdx'):
            ns = parseBench(line)
            r['Idx'] = ns
            continue
        if line.startswith('BenchmarkDoAppend'):
            ns = parseBench(line)
            r['Append'] = ns
            continue
        if line.startswith('BenchmarkPoloAppend'):
            ns = parseBench(line)
            r['PoloAppend'] = ns
            continue
        if line.startswith('BenchmarkPolo'):
            ns = parseBench(line)
            r['Polo'] = ns

print('Size Idx Append Polo PoloAppend')
sizes = sorted(t.keys())
for size in sizes:
    b = t[size]
    print(size, b['Idx'], b['Append'], b['Polo'], b['PoloAppend'])
