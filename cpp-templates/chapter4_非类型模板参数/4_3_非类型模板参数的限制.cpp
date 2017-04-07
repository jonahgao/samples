//
// Created by jonah on 17-4-7.
//

#include <iostream>

// 非类型模板参数是有限制的
// 通常，它们可以是常整数（包括枚举值）或者指向外部链接对象的指针


// 浮点数不允许作为非类型模板参数
// error: ‘double’ is not a valid type for a template non-type parameter
// template <double VAT>
// void printer() {
//     std::cout << VAT << std::endl;
// }

// 类对象也不能作为非类型模板参数
// error: ‘class std::__cxx11::basic_string<char>’ is not a valid type for a template non-type parameter
// template <std::string name>
// void printer2() {
//    std::cout << name << std::endl;
// }


template <char const* name>
void printer3() {
    std::cout << name << std::endl;
}

const char *s1 = "hello";
extern char const s2[] = "hello";

int main() {
    // 字符串字面值是内部链接对象，
    // 因为两个具有相同但处于不同模块的字符串，是两个完全不同的对象
    // 所以不能使用它们作为模板实参
    // printer3<"123">();

    // 也不能使用全局指针作为模板实参
//    printer3<s1>();

    // 可以使用外部链接对象
    printer3<s2>();
}


