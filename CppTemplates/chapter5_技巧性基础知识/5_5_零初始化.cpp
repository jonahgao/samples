//
// Created by jonah on 17-4-12.
//
#include <iostream>

// 对于 int、 double 或者指针等基本类型，任何未被初始化的局部变量都具有一个不确定值。
// 应该显式地调用内建类型的缺省构造函数，设置缺省值为零值
template <typename T>
void printZero() {
    T t1 = T();
    T t2{}; // C++11

    std::cout << t1 << "\t" << t2 << "\t" << std::endl;
}

// 对于类模板，在用某种类型实例化该模板后，为了确认它所有的成员都已经初始化完毕，
// 需要定义一个缺省构造函数，通过初始化列表来初始化类模板成员
template <typename T>
class Foo {
private:
    T t_;
public:
    Foo() : t_() { // 确保t_已经初始化，内建类型对象也是如此
    }

    void print() const { std::cout << t_ << std::endl; }
};

int main() {
    printZero<int>();
    printZero<double>();

    Foo<int> fi;
    fi.print();
    Foo<std::string> fs;
    fs.print();
}
