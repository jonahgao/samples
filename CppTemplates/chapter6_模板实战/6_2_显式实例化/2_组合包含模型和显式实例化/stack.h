//
// Created by jonah on 2017/4/13.
//

// 为了让程序员能根据实际情况自由地选择包含模型或者显式实例化，
// 我们可以把模板的定义和模板的声明放在两个不同的文件中
// 通常的做法是使用头文件来表示这两个文件 stack.h: 声明  stackdef.h: 定义

// 如果我们希望使用包含模型，那么只要 #include stackdef.h 就行
// 反之，如果希望显式实例化模板，就应该 #include  stack.h ,然后再提供一个含有需要显式实例化的 .cpp 源文件

// 本文件是模板的声明

#ifndef CPPTEMPLATES_STACK_H
#define CPPTEMPLATES_STACK_H

#include <vector>

template <typename T>
class Stack {
private:
    std::vector<T> rep_;
public:
    Stack();
    void push(T const&);
    void pop();
    T top() const;
};


#endif //CPPTEMPLATES_STACK_H
