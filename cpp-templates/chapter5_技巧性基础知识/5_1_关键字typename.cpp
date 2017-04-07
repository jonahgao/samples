//
// Created by jonah on 17-4-7.
//

#include <iostream>
#include <vector>
#include <bitset>

// typename的另一种用法：指定某个依赖模板参数的名称是一个类型
// 例如 T::SubType * ptr 会被看作是类T的静态成员和ptr的乘积，此时需要使用typename指明T::SubType是一个类型
template <class T>
void printVector(T const& t) {
    typename T::const_iterator pos;
    typename T::const_iterator end(t.end());

    for (pos = t.begin(); pos != end; ++pos) {
        std::cout << *pos << " ";
    }
    std::cout << std::endl;
}

template <int N>
void printBitset(std::bitset<N> const& bs) {
    std::cout << bs.template to_string<char, std::char_traits<char>, std::allocator<char> >();
}

int main() {
    std::vector<int> vec{1, 2, 3, 4, 5};
    printVector(vec);

    std::bitset<64> bs;
    bs.set(10, true);
    printBitset<64>(bs);
}



