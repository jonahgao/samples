//
// Created by jonah on 2017/4/13.
//

// 显式初始化模式: 只需包含 stack.h, 然后把 explicit_inst.cpp 一起加入编译单元
#include "stack.h"
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
    std::cout << ss.top() << std::endl;
    // ss.pop(); // 出错，并未显式初始化pop成员函数
}

