#!/bin/env python
# -*- coding: utf-8 -*-

def foo():
    return "different≠another"

if __name__ == "__main__":
    x = "Hello World!"
    if (True):
        x = foo()
        print(x)
    print(x)
