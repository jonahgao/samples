//
// Created by jonah on 2017/4/13.
//

// 包含模型：需要include stack_def.h
#include "stackdef.h"
#include <iostream>

int main() {
    Stack<int> si;
    si.push(1);
    si.push(2);
    std::cout << si.top() << std::endl;
    si.pop();
    std::cout << si.top() << std::endl;

    Stack<std::string> ss;
    ss.push("a");
    ss.push("b");
    std::cout << ss.top() << std::endl;
    ss.pop();
    std::cout << ss.top() << std::endl;
}
