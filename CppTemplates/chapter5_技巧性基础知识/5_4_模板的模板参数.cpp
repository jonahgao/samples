//
// Created by jonah on 17-4-9.
//
// 在stack的例子中，如果我们要指定底层容器的类型，就必须指定两次元素类型：
// Stack<int, std::vector<int> > vStack; // 两次int

// 借助于模板的模板参数，可以只指定容器的类型而不需要指定所含元素的类型 就像这样:
// Stack<int, std::vector> vStack;
// 为了获得这个特性，必须把第2个模板参数指定为模板的模板参数

// 注： 函数模板并不支持模板的模板参数 (template template argument)

#include <deque>
#include <vector>
#include <iostream>

template <typename T, template <typename ELEM, typename ALLOC = std::allocator<ELEM> > class CONT = std::deque >
//template <typename T, template <typename ELEM, typename ALLOC = std::allocator<ELEM> > typename CONT = std::deque >
class Stack {
private:
    CONT<T> rep_;
public:
    void push(T const& e) {
        rep_.push_back(e);
    }

    T top() const {
        if (rep_.empty()) {
            throw std::out_of_range("Stack<>::top() stack is empty");
        }
        return rep_.back();
    }
};

int main(void) {
    Stack<int, std::vector> s;
    s.push(1);
    std::cout << s.top() << std::endl;
}
