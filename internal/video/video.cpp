#include <iostream>

using namespace std;

extern "C" {
void test() { cout << "TEST FUNCTION RAN" << endl; }
}