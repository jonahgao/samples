#include <iostream>

using namespace std;

template <class T>
void test(T) {
    cout << "template" << endl;
}

void test(int) {
    cout << "normal" << endl;
}

int main() {
    test(1);  // 优先选择非函数模板 output:"normal"
    test('a'); // 函数模板可以产生更好的匹配，因此选择函数模板 output: "template"
    test<>(1); // 显式指定一个空的模板实参列表，告诉编译器使用模板 output:"template"
    return 0;
}