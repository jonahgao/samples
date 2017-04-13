//
// Created by jonah on 2017/4/13.
//

// 本文件是模板的定义

#ifndef CPPTEMPLATES_STACKDEF_H
#define CPPTEMPLATES_STACKDEF_H

#include "stack.h"

template <typename T>
Stack<T>::Stack() {}

template <typename T>
void Stack<T>::push(const T& t) {
    rep_.push_back(t);
}

template <typename T>
void Stack<T>::pop() {
    if (rep_.empty()) {
        throw std::out_of_range("Stack<>::pop empty");
    }
    rep_.pop_back();
}

template <typename T>
T Stack<T>::top() const {
    if (rep_.empty()) {
        throw std::out_of_range("Stack<>::top empty");
    }
    return rep_.back();
}

#endif //CPPTEMPLATES_STACKDEF_H
