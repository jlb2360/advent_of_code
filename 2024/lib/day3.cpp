#include "day3.h"
#include <string>
#include <iostream>
#include <vector>

void Mult::print_token(){
    std::cout << "(val1, val2): (" << this->val1 << ", " << this->val2 << ")\n";
}

int Mult::perform_action(int &modif){
    return (this->val1 * this->val2) * modif;
}

void Do::print_token(){
    std::cout << "Activating multiplication\n";
}

int Do::perform_action(int &modif){
    modif = 1;
    return 0;
}

void Dont::print_token(){
    std::cout << "Deactivating multiplication\n";
}

int Dont::perform_action(int &modif){
    modif = 0;
    return 0;
}


void find_val(const std::string &line, std::vector<Token*> &toks, int &idx, std::string &val){

    if (idx > line.size()){
        return;
    }

    
    if (std::isdigit(line[idx])){
        val.push_back(line[idx]);

        idx++;
        find_val(line, toks, idx, val);
    } else if (line[idx] == ',' || line[idx] == ')'){
        return;
    } else{
        val = "no";
        return;
    }

}

void l_token(const std::string &line, std::vector<Token*> &toks, int &idx, int &v1, int &v2){
    if (idx > line.size()){
        return;
    }
    
    if (line[idx] == '('){
        
        idx++;

        std::string val;

        find_val(line, toks, idx, val);

        try {
            v1 = std::stoi(val);
        } catch (const std::invalid_argument& e) {
            return;
        } catch (const std::out_of_range& e) {
            return;
        }

        l_token(line, toks, idx, v1, v2);

    } else if (line[idx] == ','){
        
        idx++;

        std::string val;
        
        find_val(line, toks, idx, val);

        try {
            v2 = std::stoi(val);
        } catch (const std::invalid_argument& e) {
            return;
        } catch (const std::out_of_range& e) {
            return;
        }

        l_token(line, toks, idx, v1, v2);
    
    } else if (line[idx] == ')'){
        Token *new_tok = new Mult(v1, v2);
        toks.push_back(new_tok);
    }

}

void ul_token(const std::string &line, std::vector<Token*> &toks, int &idx, int &v1, int &v2){
    if (idx > line.size()){
        return;
    }

    else if (line[idx] == 'l'){

        idx++;
        
        l_token(line, toks, idx, v1, v2);
    }
}


void mul_token(const std::string &line, std::vector<Token*> &toks, int &idx, int &v1, int &v2){
    if (idx > line.size()){
        return;
    }
    

    if(line[idx] == 'u'){

        idx++;
        
        ul_token(line, toks, idx, v1, v2);
    
    }

}

void dont_token(const std::string &line, std::vector<Token*> &toks, int &idx){
    if (idx > line.size()){
        return;
    }

    if (line[idx] == '\''){

        idx++;

        if(line[idx] != 't'){
            return;
        }

        idx++;

        if(line[idx] != '('){
            return;
        }

        idx++;

        if (line[idx] != ')'){
            return;
        }

        Token* new_tok = new Dont();
        toks.push_back(new_tok);

        idx++;

    }
}

void do_or_dont_token(const std::string &line, std::vector<Token*> &toks, int &idx){
    if (idx > line.size()){
        return;
    }

    if (line[idx] == 'o'){
        idx++;

        if (line[idx] == 'n'){

            idx++;
            dont_token(line, toks, idx);

        } else if (line[idx] == '('){
            idx++;
            if (line[idx] == ')'){

                Token* new_tok = new Do();
                toks.push_back(new_tok);

            } else {

                return;

            }
        } else {

            return;

        }
    }

}

void tokenize_string(const std::string &line, std::vector<Token*> &toks, int &idx){
    
    if (idx > line.size()){
        return;
    }

    if (line[idx] == 'm'){
        idx++;
        int v1;
        int v2;
        mul_token(line, toks, idx, v1, v2);

        tokenize_string(line, toks, idx);
    }

    if (line[idx] == 'd'){
        idx++;
        do_or_dont_token(line, toks, idx);

        tokenize_string(line, toks, idx);
    }

    idx++;

    tokenize_string(line, toks, idx);

}

int process_tokens(std::vector<Token*> &toks, int &modif){
    int res = 0;
    
    for (auto token : toks){
        res += token->perform_action(modif);

        delete token;
    }

    toks.clear();

    return res;
    
}