#include <iostream>
using namespace std;
int main(int argc, char const *argv[])
{
    cout << "hello world" <<endl;
    string str="12345678901234567890";
    const char *k=str.c_str();
    const char *t=str.data();
    printf("%s,%s",k,t);
    cout<<k<<t<<endl;
    printf("\n--------------------\n");
    char *data;
    int len=str.length();
    data=(char *)malloc((len+1)*sizeof(char));
    str.copy(data,len,0);
    printf("%s\n",data);
    cout<<data;
    return 0;
}
