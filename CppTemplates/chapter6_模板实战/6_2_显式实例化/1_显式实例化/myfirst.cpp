//
// Created by jonah on 2017/4/13.
//

#include "myfirst.h"
#include <iostream>
#include <typeinfo>

template <typename T>
void printTypeof(T const& t) {
    std::cout << typeid(t).name() << std::endl;
}

// 显式实例化由关键字template和紧接其后的我们所需要实例化的实体（可以是类、函数、成员函数等组成）
// 而且该声明是一个已经用实参完全替换参数之后的声明。
// 显式实例化需要出现在模板定义之后，给定的实参组合在一个程序中最多只能有一个显式实例化体
template void printTypeof<double>(double const&);

// 基于int显式实例化Myclass<>的构造函数
// template Myclass<int>::Myclass();

// 基于int显式实例化函数模板max
// template int const& max(int const&, int const&);

// 还可以显式实例化类模板，这样就可以同时实例化它的所有类成员
// 对于已经在前面实例化过的成员，就不能再次对它进行实例化
// template class Stack<int>;

// 显式实例化的缺点：必须仔细跟踪每个需要实例化的实体
// 优点：实例化可以在需要的时候才进行。
// 可以避免包含庞大头文件的开销（模板定义跟声明可以分开，模板定义放源文件中）
// 更可以把模板定义的源文件封装起来，但封装之后客户端程序就不能基于其他类型来进行额外的实例化了。




