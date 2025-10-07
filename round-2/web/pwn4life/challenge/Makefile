SRC = app.c
NAME = pwn4life
CFLAGS = -no-pie -fno-stack-protector

$(NAME): $(SRC)
	gcc $(CFLAGS) $(SRC) -o $(NAME)

clean :
	rm -rf $(NAME)
