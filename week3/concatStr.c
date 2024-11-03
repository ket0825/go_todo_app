#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char* concatStr(const char* str1, const char* str2) {    
    char* result = (char*)malloc(strlen(str1) + strlen(str2) + 1);

    strcpy(result, str1);
    strcat(result, str2);
    
    return result;
}