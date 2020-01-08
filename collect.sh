#!/bin/bash

set -e

step="$1"
times="$2"x

./bin/bench -output=./_assets/ -test.run=xx -test.bench=BenchmarkUrkelBlock$step$ -test.v -test.benchmem -test.benchtime=$times -test.timeout=60m

./bin/bench -output=./_assets/ -test.run=xx -test.bench=BenchmarkPatriciaBlock$step$ -test.v -test.benchmem -test.benchtime=$times -test.timeout=60m

./bin/bench -output=./_assets/ -test.run=xx -test.bench=BenchmarkIavlTreeBlock$step$ -test.v -test.benchmem -test.benchtime=$times -test.timeout=60m
