//
// Created by jonah on 17-4-7.
//
#include <iostream>
#include <vector>
#include <algorithm>

// 也可以为函数模板定义非类型参数
template <typename T, int VAL>
T addValue(T const& x) {
    return x + VAL;
}

// 如果需要把函数或者操作 用作参数的话，这类函数相当有用
int main() {
    std::vector<int> vec{1, 2, 3, 4};
    // vector每个元素加10,
    const int a = 10; // 可以是常量，不一定非是字面值
    std::transform(vec.begin(), vec.end(), vec.begin(), addValue<int, a>);
//  虽然使用lambda更方便:)
//  std::transform(vec.begin(), vec.end(), vec.begin(), [](int a){ return a + 10; });
    for (auto x : vec) {
        std::cout << x << " " << std::endl;
    }
    std::cout << std::endl;
    return 0;
}

