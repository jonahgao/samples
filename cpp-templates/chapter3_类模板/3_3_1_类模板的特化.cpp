//
// Created by jonah on 17-4-6.
//
#include <iostream>
#include <vector>
#include <string>
#include <stdexcept>
#include <deque>

template <class T>
class Stack {
private:
    std::vector<T> rep_;

public:
    void push(T const& e) { rep_.push_back(e); }

    T pop() {
        if (rep_.empty()) {
            throw std::out_of_range("Stack<>::pop(): empty Stack");
        }
        T e = rep_.back();
        rep_.pop_back();
        return e;
    }

    T top() const {
        if (rep_.empty()) {
            throw std::out_of_range("Stack<>::top(): empty Stack");
        }
        return rep_.back();
    }
};

// 特化一个类模板：在起始处声明一个 template<>
// 接下来声明用来特化类模板的类型，这个类型被用作模板实参，且必须在类名的后面直接指定
template <>
class Stack<std::string> {
private:
    std::deque<std::string> rep_; // 实现可以跟基本类模板完全不同

public:
    void push(std::string const& e);

    std::string pop();
    std::string top() const;
};

// 进行类模板的特化时，每个成员函数都必须重新定义为普通函数
// 原来模板函数中的每个T也相应地被进行特化的类型取代
void Stack<std::string>::push(std::string const& e) {
    rep_.push_back(e);
}

std::string Stack<std::string>::pop() {
    if (rep_.empty()) {
        throw std::out_of_range("Stack<std::string>::pop(): empty Stack");
    }
    auto e = rep_.back();
    rep_.pop_back();
    return e;
}

std::string Stack<std::string>::top() const {
    if (rep_.empty()) {
        throw std::out_of_range("Stack<std::string>::top(): empty Stack");
    }
    return rep_.back();
}

int main() {
    Stack<int> s;
    s.push(1);
    s.push(2);
    std::cout << s.pop() << std::endl;
    std::cout << s.top() << std::endl;

    Stack<std::string> s2;
    s2.push("aa");
    s2.push("bb");
    std::cout << s2.pop() << std::endl;
    std::cout << s2.top() << std::endl;

    return 0;
}
