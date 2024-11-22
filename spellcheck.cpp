#include <iostream>
#include <fstream>
#include <windows.h>
using namespace std;

int main(int argc, char** argv) {
  if (argc < 3) {
    cerr << "Error: no input provided" << endl;
    return 1;
  } else {
    //filename creation
    char* arg = argv[2];
    char* path = new char[54];
    strcpy(path, "words-");
    strcat(path, arg);
    strcat(path, ".txt");

    //create file
    ofstream word_file(path, ofstream::trunc);



    word_file << argv[1] << endl;
    word_file.close();


    //allow api to check words then delete file
    Sleep(5000);
    remove(path);
    return 0;
  }
}