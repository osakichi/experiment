#include <string>
#include <iostream>

std::string foo() {
    return "different≠another";
}

int main() {
    auto x = std::string("Hello World!");
    if (true) {
        x = foo();
        std::cout << x << std::endl;
    }
    std::cout << x << std::endl;
    return 0;
}
