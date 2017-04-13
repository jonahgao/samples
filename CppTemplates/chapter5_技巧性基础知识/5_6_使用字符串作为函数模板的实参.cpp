//
// Created by jonah on 2017/4/13.
//
// 对于非引用类型的参数，在实参演绎的过程中，会出现数组到指针的(array-to-pointer)的类型转换
// 这种类型转换也被称为 decay

#include <typeinfo>
#include <iostream>

template <typename T>
void ref(T const& x) {
    std::cout << "x in ref(T const&): " << typeid(x).name() << std::endl;
}

template <typename T>
void unref(T x) {
    std::cout << "x int unref(T): " << typeid(x).name() << std::endl;
}

int main() {
    ref("Hello"); // char[6]
    ref("test"); // char[4]
    unref("Hello"); // const char *
    return 0;
}

