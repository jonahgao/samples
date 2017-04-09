#include <deque>
#include <stdexcept>
#include <iostream>

//
// Created by jonah on 17-4-9.
//
// 类成员也可以是模板。嵌套类和成员函数都可以作为模板
template <typename T>
class Stack {
private:
    std::deque<T> rep_;

public:
    void push(T const& e)  {
        rep_.push_back(e);
    }

    void pop() {
        if (rep_.empty()) {
            throw std::out_of_range("Stack<> pop: stack is empty");
        }
        rep_.pop_back();
    }

    T top() const {
        if (rep_.empty()) {
            throw std::out_of_range("Stack<> top: stack is empty");
        }
        return rep_.back();
    }

    bool empty() const {
        return rep_.empty();
    }

    // 使用元素类型为T2的栈进行赋值
    template <typename T2>
    Stack<T>& operator=(Stack<T2> const &);
};

template <typename T>
template <typename T2>
Stack<T>& Stack<T>::operator=(Stack<T2> const &s2) {
    // 检查是否是赋值给自身
    if (reinterpret_cast<const void *>(this) == reinterpret_cast<const void *>(&s2)) {
        return *this;
    }

    // 复制一份
    Stack<T2> tmp(s2);
    rep_.clear();
    while (!tmp.empty()) {
        this->push(tmp.top());
        tmp.pop();
    }
    return *this;
}

int main(void) {
    Stack<int> si;
    for (int i = 0; i < 10; i++) {
       si.push(i);
    }

    Stack<float> sf;
    sf = si;
    while (!sf.empty()) {
        std::cout << sf.top() << " ";
        sf.pop();
    }
    std::cout << std::endl;
}

