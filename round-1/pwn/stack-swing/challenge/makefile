SRC = main.c
NAME = pwn2
CFLAGS = -no-pie -fno-stack-protector

$(NAME): $(SRC)
	gcc $(CFLAGS) $(SRC) -o $(NAME)

clean :
	rm -rf $(NAME)