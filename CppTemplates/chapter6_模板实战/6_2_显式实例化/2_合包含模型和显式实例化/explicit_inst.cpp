//
// Created by jonah on 2017/4/13.
//
#include "stackdef.h"
#include <string>

// 为int类型显式实例化
template class Stack<int>;

// 为string类型显式显例化一些成员函数
template Stack<std::string>::Stack();
template void Stack<std::string>::push(std::string const &);
template std::string Stack<std::string>::top() const;
