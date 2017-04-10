//
// Created by jonah on 17-4-6.
//

#include <iostream>

template <typename T1, typename T2>
class MyClass {
public:
    void Test() {
        std::cout << "basic template class" << std::endl;
    }
};

// 局部特化：两个模板参数具有相同类型
template <typename T>
class MyClass<T, T> {
public:
    void Test() {
        std::cout << "partial specialization: two same type" << std::endl;
    }
};

// 局部特化：第二个模板参数是int
template <typename T>
class MyClass<T, int> {
public:
    void Test() {
        std::cout << "partial specialization: second is integer" << std::endl;
    }
};

// 局部特化： 两个模板参数都是指针类型
template <typename T1,typename T2>
class MyClass<T1*, T2*> {
public:
    void Test() {
        std::cout << "partial specialization: two pointer type" << std::endl;
    }
};

int main() {
    MyClass<int, float> mif; // 使用基本类模板
    mif.Test();

    MyClass<float, float> mff; // 使用MyClass<T, T>
    mff.Test();

    MyClass<float, int> mfi; // 使用Myclass<T, int>
    mfi.Test();

    MyClass<int *, float*> mp; // 使用Myclass<T1*, T2*>
    mp.Test();

    // MyClass<int, int> m;  // 错误：同等匹配 MyClass<T,T> 和 MyClass<T, int>

    // MyClass<int*, int*> m2; // 错误：同等匹配 MyClass<T,T> 和 MyClass<T1*, T2*>
}

