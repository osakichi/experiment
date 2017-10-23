#include <string>
#include <iostream>

std::string foo() {
    return "different≠another";
}

main() {
    std::string x = "Hello World!";
    if (true) {
        x = foo();
        std::cout << x << std::endl;
    }
    std::cout << x << std::endl;
}
