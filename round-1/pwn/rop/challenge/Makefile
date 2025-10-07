CC = gcc
SRCS = main.c
OBJS = chall
CFLAGS = -Wall -O0
NAME = pwn1

.PHONY: clean

$(OBJS): $(SRCS)
	$(CC) $(SRCS) $(CFLAGS) -o $(NAME)
	
clean: 
	rm $(OBJS)