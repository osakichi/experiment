PROGRAM=hello

.SILENT:

all: golang java cpp python lua

golang: $(PROGRAM).go
	echo "== Golang =="
	go run $<

java: $(PROGRAM).java
	echo "== Java =="
	javac $< && java $(PROGRAM) && rm -f $(PROGRAM).class

cpp: $(PROGRAM).cpp
	echo "== C++ =="
	c++ -std=c++11 -o $(PROGRAM) $< && ./$(PROGRAM) && rm -f $(PROGRAM)

python: $(PROGRAM).py
	echo "== Python =="
	python ./$<

lua: $(PROGRAM).lua
	echo "== Lua =="
	lua ./$<

