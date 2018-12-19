#include <iostream>
#include <string>
#include <set>
#include <string.h>
#include <iterator>

using namespace std;

struct strLess{
   bool operator() (const char *s1, const char *s2) const {
    return strcmp(s1, s2) < 0;
   }
};

void printSet(set<int> s){
	copy(s.begin(), s.end(), ostream_iterator<int>(cout, ", ") );
	/*
	set<int>::iterator iter;
	for (iter = s.begin(); iter != s.end(); iter++)
	    cout<<"set["<<iter-s.begin()<<"]="<<*iter<<", "; //Error
	cout<<*iter<<", ";
	*/
	cout<<endl;
}

int 
main(){
	set<int> s1;
	
    set<const char*, strLess> s2( strLess); 
	
	set<int> s3(s1); 
	
	int iArray[] = {13, 32, 19};
	set<int> s4(iArray, iArray + 3);
	
	const char* szArray[] = {"hello", "dog", "bird" };
	set<const char*, strLess> s5(szArray, szArray + 3, strLess() );
	
	cout<<"s1.insert() : "<<endl;
	for (int i = 0; i <5 ; i++)
        s1.insert(i*10);
    printSet(s1);
	
	cout<<"s1.insert(20).second = "<<endl;
	if (s1.insert(20).second)
        cout<<"Insert OK!"<<endl;
	else
        cout<<"Insert Failed!"<<endl;
	cout<<"s1.insert(50).second = "<<endl;
	if (s1.insert(50).second){
		cout<<"Insert OK!"<<endl;
		printSet(s1);
	}else
		cout<<"Insert Failed!"<<endl;
	cout<<"pair<set<int>::iterator, bool> p;\np = s1.insert(60);\nif (p.second):"<<endl;
	pair<set<int>::iterator, bool> p;
	p = s1.insert(60);
	if (p.second){
		cout<<"Insert OK!"<<endl; 
		printSet(s1);
	}
	else
		cout<<"Insert Failed!"<<endl;

   
	cout<<"\ns1.erase(70) = "<<endl;
	s1.erase(70);
	printSet(s1);
	cout<<"s1.erase(60) = "<<endl;
	s1.erase(60);
	printSet(s1);

    cout<<"set<int>::iterator iter = s1.begin();\ns1.erase(iter) = "<<endl;
    set<int>::iterator iter = s1.begin();
    s1.erase(iter);
    printSet(s1);

    
	cout<<"\ns1.count(10) = "<<s1.count(10)<<", s1.count(80) = "<<s1.count(80)<<endl;
	cout<<"s1.find(10) : ";
	if (s1.find(10) != s1.end()) 
		cout<<"OK!"<<endl;
	else
		cout<<"not found!"<<endl;
	
	cout<<"s1.find(80) : ";
	if (s1.find(80) != s1.end()) 
		cout<<"OK!"<<endl;
	else
		cout<<"not found!"<<endl;
	
	cout<<"\ns1.empty()="<<s1.empty()<<", s1.size()="<<s1.size()<<endl;
	set<int> s9;
	s9.insert(100);
	cout<<"s1.swap(s9) :"<<endl;
	s1.swap(s9);
	cout<<"s1: "<<endl;
	printSet(s1);
	cout<<"s9: "<<endl;
	printSet(s9);
	return 0;
}
