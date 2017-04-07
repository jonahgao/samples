//
// Created by jonah on 17-4-7.
//
#include <stdexcept>
#include <iostream>

template <typename T, int MAXSIZE>
//template <typename T = int, int MAXSIZE = 100>  // 也可以为模板指定缺省值
class Stack {
private:
    T elems_[MAXSIZE];
    int numElems_;

public:
    Stack();
    void push(T const& e);
    void pop();
    T top() const;
};

template <typename T, int MAXSIZE>
Stack<T, MAXSIZE>::Stack() : numElems_(0) {}

template <typename T, int MAXSIZE>
void Stack<T, MAXSIZE>::push(T const &e) {
    if (MAXSIZE == numElems_) {
        throw std::out_of_range("Stack<>::push(): stack is full");
    }
    elems_[numElems_++] = e;
}

template <typename T, int MAXSIZE>
void Stack<T, MAXSIZE>::pop() {
    if (numElems_ <= 0) {
        throw std::out_of_range("Stack<>::pop(): stack is empty");
    }
    --numElems_;
};

template <typename T, int MAXSIZE>
T Stack<T, MAXSIZE>::top() const {
    if (numElems_ <= 0) {
        throw std::out_of_range("Stack<>::top(): stack is empty");
    }
    return elems_[numElems_-1];
};

int main() {
    Stack<int, 10> s;
    s.push(1);
    s.push(2);
    std::cout << s.top() << std::endl;
    s.pop();
    std::cout << s.top() << std::endl;

    // 每个模板实例都有自己的类型。
    // s 和 s2属于不同的类型，两种类型不存在显式或隐式的转换
    // 它们之间不能相互替换，更不能相互赋值
    Stack<int, 20> s2;
}


