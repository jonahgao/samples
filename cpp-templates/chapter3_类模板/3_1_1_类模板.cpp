//
// Created by jonah on 17-4-6.
//

#include <iostream>
#include <vector>
#include <stdexcept>

template <class T>
class Stack {
private:
   std::vector<T> rep_;

public:
    // 当使用类名而不是类的类型时，就应该使用Stack
    // 譬如，当你指定类的名称、类的构造函数、析构函数时，就应该使用Stack
    Stack() = default;

    // 当在声明中需要使用该类的类型时，必须使用Stack<T>
    Stack<T>& operator=(Stack<T> const&);

    void push(T const&);
    T pop();

    // 实现为内联函数，定义于类声明里面
    T top() const {
        if (rep_.empty()) {
            throw std::out_of_range("Stack<>::top(): empty Stack");
        }
        return rep_.back();
    }

    // 只有那些被调用的成员函数，才会产生这些函数的实例化代码
    // 对于类模板，只有在被使用的时候才会初始化
    void unusedMethod();
};

// 为了定义类模板的成员函数，必须指定该成员函数是一个函数模板
// 还需要使用这个类模板的完整类型限定符 Stack<T>::
template <typename T>
void Stack<T>::push(T const & e) {
    rep_.push_back(e);
}

template <typename T>
T Stack<T>::pop() {
    if (rep_.empty()) {
        throw std::out_of_range("Stack<>::pop(): empty Stack");
    }
    T e = rep_.back();
    rep_.pop_back();
    return e;
}

int main() {
    Stack<int> s;
    s.push(1);
    s.push(2);

    std::cout << s.pop() << std::endl;
    std::cout << s.top() << std::endl;

    return 0;
}

