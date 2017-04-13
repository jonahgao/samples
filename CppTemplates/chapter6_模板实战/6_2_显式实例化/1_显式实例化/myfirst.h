//
// Created by jonah on 2017/4/13.
//

// 包含模型（inclusion model, 头文件中同时包含模板声明和模板定义）能够确保所有需要的模板都已经实例化。
// 这是因为：当需要进行实例化的时候，C++编译系统会自动产生所对应的实例化体

// 另外C++标准还提供了一种手工实例化的机制：显式实例化指示符（explicit instantiation directive)

#ifndef CPPTEMPLATES_MYFIRST_H
#define CPPTEMPLATES_MYFIRST_H

template <typename T>
void printTypeof(T const& t);


#endif //CPPTEMPLATES_MYFIRST_H
