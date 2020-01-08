.PHONY: build
build:
	mkdir -p ./bin
	go build -o bin/ptool ./
	go test -c -o ./bin/bench ./benchs/

run-100:
	./collect.sh 100 1000

run-1000:
	./collect.sh 1000 1000

run-10000:
	./collect.sh 10000 1000

run-40000:
	./collect.sh 40000 1000

run-all: run-100 run-1000 run-10000 run-40000

plot-100:
	./plot.sh 100

plot-1000:
	./plot.sh 1000

plot-10000:
	./plot.sh 10000

plot-40000:
	./plot.sh 40000

plot-all: plot-100 plot-1000 plot-10000 plot-40000

view-all:
	sensible-browser ./_assets/*.png
