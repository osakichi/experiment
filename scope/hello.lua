#!/bin/env lua

function foo()
    return "different≠another"
end

x = "Hello World!"
if (true) then
    x = foo()
    print(x)
end
print(x)
