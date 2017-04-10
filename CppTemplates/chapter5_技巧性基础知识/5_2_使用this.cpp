//
// Created by jonah on 17-4-9.
//
#include <iostream>

using namespace std;

void test() {
    cout << "::test" << endl;
}

template <typename T>
class Base {
public:
    void test() {
        cout << "Base::test" << endl;
    }

    void test2() {
        cout << "Base::test2" << endl;
    }
};

// 对于有基类的模板，自身使用名字x，并不一定等同于this->x，即使该x是从基类继承获得的。
// 对于那些在基类中声明并且依赖于模板参数的符号（函数或变量），应该在它们前面使用this->x或者Base<T>::x
template <typename T>
class Derived: public Base<T> {
public:
    void foo() {
        test();     // 不指定this，调用了外部的test() 导致错误的结果
        this->test();   // 指定了this，调用了基类的test()
//      test2();  // 报错。there are no arguments to ‘test2’ that depend on a template parameter, so a declaration of ‘test2’ must be available
        Base<T>::test2(); // 使用Base<T>，调用基类的test2()
    }
};

int main() {
    Derived<int> d;
    d.foo();
}

