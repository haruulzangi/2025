#include <stdio.h>
#include <stdlib.h>

char *array[0x10]={};
size_t size_array[0x10]={};

void menu(){
	puts("1.\tAdd new note");
	puts("2.\tDelete note");
	puts("3.\tEdit note");
	puts("4.\tView note");
	puts("5.\tExit");
	printf(">");
}
int  readint(){
	char buf[0x10];
	read(0,buf,0x10);
	return atoi(buf);
}
void add(){
	size_t i=0;
	for(;i<0x10;i++){
		if(array[i]==0)
			break;
	}
    puts("Size:");
    size_t size = readint();
    if(size < 0x2000){
	    char *ptr=malloc(size);
        if(ptr > 0)
		{
            array[i] = ptr;
            size_array[i] = size;
        }
        else
            return ;
    }
}
void delete(){
	puts("Note idx?");
	printf(">");
	size_t i= readint();
	if(i<0x10)
	{
		free(array[i]);
		// array[i] = 0 ;
        // size_array[i] = 0 ;
	}
}
void edit(){
	puts("Edit idx?");
	printf(">");
	size_t i= readint();
	if(i<0x10){
		puts("Data:");
		read(0,array[i],size_array[i]);
	}
}
void show(){
	puts("Read idx?");
	printf(">");
    size_t i= readint();
	if(i<0x10){
		puts("Data:");
		write(1,array[i],size_array[i]);
	}
}
void init()
{
	setvbuf(stdin,0,2,0);
	setvbuf(stdout,0,2,0);
	setvbuf(stderr,0,2,0);
}
int main()
{
	init();
	while(1)
	{
		menu();
		int cmd = readint();
		if(cmd==1) add();
		else if(cmd==2) delete();
		else if(cmd==3) edit();
        else if(cmd==4) show();
		else break;
	}
    return 0;
}
