#pragma once

#include <string>
#include <vector>

class Token{
public:
    Token(){}
    ~Token(){}

    virtual void print_token() = 0;
    virtual int perform_action(int &modif) = 0;
};

class Mult : public Token{
public:
    int val1;
    int val2;

    Mult(int v1, int v2) : val1(v1), val2(v2){}
    ~Mult(){}

    void print_token();
    int perform_action(int &modif);
};

class Do : public Token{
public:
    
    Do(){};

    void print_token();
    int perform_action(int &modif);

};

class Dont : public Token{
public:

    Dont(){};

    void print_token();
    int perform_action(int &modif);
};

void tokenize_string(const std::string &, std::vector<Token*> &, int &);

int process_tokens(std::vector<Token*> &, int &);