#include "day4.h"
#include <iostream>

using namespace std;


int check_diagonal(vector<string> lines, int i, int j){
    int diag_matchs = 0;
    
    if (j+3 < lines[i].size()){

        // check right up diagonal

        if (i+3 < lines.size()){
            if (lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S'){
                diag_matchs += 1;
            }
        }

        // check right down diagonal
        if (i-3 >= 0){
            if (lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S'){
                diag_matchs += 1;
            }
        }
        
    }

    if (j-3 >=0){

        // check left up diagonal
        if (i+3 < lines.size()){
            if (lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S'){
                diag_matchs += 1;
            }
        }


        // check left down diagonal
        if (i-3 >= 0){
            if (lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S'){
                diag_matchs += 1;
            }
        }
    }

    return diag_matchs;
    
}

int check_horizontal(vector<string> lines, int i, int j){
    int diag_matchs = 0;
    // check right

    if (j+3 <= lines[i].size()){
        
        if (lines[i][j+1] == 'M' && lines[i][j+2] == 'A' && lines[i][j+3] == 'S'){
            diag_matchs += 1;
        }

    }
    // check left
    if (j-3 >= 0){
        if (lines[i][j-1] == 'M' && lines[i][j-2] == 'A' && lines[i][j-3] == 'S'){
            diag_matchs += 1;
        }
        
    }

    return diag_matchs;
    
}

int check_vertical(vector<string> lines, int i, int j){
    int diag_matchs = 0;
    // check up

    if(i+3 <= lines.size()){
        if (lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S'){
            diag_matchs += 1;
        }
    }

    // check down
    if (i-3 >= 0){
        if (lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S'){
            diag_matchs += 1;
        }
    }

    return diag_matchs;
    
}

void count_xmas(int &sum, vector<string> lines){
    
    for (size_t i = 0; i < lines.size(); i++){
        for (size_t j = 0; j < lines[i].size(); j++){
            
            if (lines[i][j] == 'X'){
                
                sum += check_diagonal(lines, i, j);

                sum += check_horizontal(lines, i, j);

                sum += check_vertical(lines, i, j);
            }
            

        }
    }
      
}





void count_x_mas(int &sum, vector<string> lines){
    for (int i = 0; i < lines.size(); i++){
        for (int j = 0; j < lines[i].size(); j++){

            if (i-1 < 0 || i+1 >= lines.size() || j-1 < 0 || j+1 >= lines[i].size()){
                continue;
            }
            
            if (lines[i][j] == 'A'){
                
                if(lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S'){

                    if(lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S'){
                        sum += 1;
                    } else if ( lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M' ) {
                        sum += 1;
                    }
                    
                } else if ( lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M' ){

                    if(lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S'){
                        sum += 1;
                    } else if ( lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M' ) {
                        sum += 1;
                    }

                }

            }
            

        }
    }
}

