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
	c++ -o $(PROGRAM) $< && ./$(PROGRAM) && rm -f $(PROGRAM)

python: $(PROGRAM).py
	echo "== Python =="
	./$<

lua: $(PROGRAM).lua
	echo "== Lua =="
	./$<

