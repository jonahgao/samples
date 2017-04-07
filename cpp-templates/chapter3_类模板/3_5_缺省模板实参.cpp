//
// Created by jonah on 17-4-6.
//
#include <iostream>
#include <vector>
#include <deque>

// 对于类模板，还可以为模板参数定义缺省值，这些值就被称为缺省模板实参
// 而且它们还可以引用之前的模板参数，例如std::vector<T>引用了前面的模板参数T
template <typename T, typename CONT = std::vector<T>>
class Stack {
private:
    CONT rep_;

public:
    void push(T const& e);
    T pop();
    T top() const;
    bool empty() const { return rep_.empty(); }
};

template <typename T, typename CONT>
void Stack<T, CONT>::push(T const &e) {
    rep_.push_back(e);
}

template <typename T, typename CONT>
T Stack<T, CONT>::pop() {
    if (rep_.empty()) {
        throw std::out_of_range("Stack<>::pop(): empty Stack");
    }
    T e = rep_.back();
    rep_.pop_back();
    return e;
}

template <typename T, typename CONT>
T Stack<T, CONT>::top() const {
    if (rep_.empty()) {
        throw std::out_of_range("Stack<>::top(): empty Stack");
    }
    return rep_.back();
}


int main() {
    Stack<int> s1; // 只传递一个模板参数，将使用vector作为底层容器
    s1.push(1);
    s1.push(2);
    std::cout << s1.pop() << std::endl;
    std::cout << s1.top() << std::endl;

    Stack<int, std::deque<int>> s2; // 传递两个模板参数，底层容器使用deque
    s2.push(1);
    s2.push(2);
    std::cout << s2.pop() << std::endl;
    std::cout << s2.top() << std::endl;
}

