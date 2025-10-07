NAME := fastbin
PORT := 30010
CFLAGS := 
all: $(NAME).c
	gcc $(CFLAGS) $(NAME).c -o $(NAME)

clean: .PHONY
	docker container rm -f $(NAME)
	docker rmi -f $(NAME)
	rm -f $(NAME) $(NAME).zip libc-2.23.so ld-2.23.so

all:
	gcc $(CFLAGS) $(NAME).c -o $(NAME)

docker: Dockerfile
	docker build -t $(NAME) --build-arg name=$(NAME) --build-arg port=$(PORT) .

run: docker
	docker run -d --read-only --restart=always --name $(NAME) -p $(PORT):$(PORT) $(NAME)

#root@adb766e6f058:/pwn# ls -la /lib/x86_64-linux-gnu/libc.so.6 
#lrwxrwxrwx 1 root root 12 Apr 30 18:20 /lib/x86_64-linux-gnu/libc.so.6 -> libc-2.31.so
#root@adb766e6f058:/pwn# ls -la /lib64/ld-linux-x86-64.so.2
#lrwxrwxrwx 1 root root 32 Apr 30 18:20 /lib64/ld-linux-x86-64.so.2 -> /lib/x86_64-linux-gnu/ld-2.31.so


extract: docker
	$(eval id := $(shell docker create $(NAME)))
	docker cp $(id):/pwn/$(NAME) $(NAME)
	#docker cp $(id):/lib/x86_64-linux-gnu/libc.so.6 libc.so.6
	docker cp $(id):/lib/x86_64-linux-gnu/libc-2.23.so libc-2.23.so
	docker cp $(id):/lib/x86_64-linux-gnu/ld-2.23.so ld-2.23.so
	docker rm -v $(id)

dist: extract
	zip $(NAME).zip $(NAME) libc.so.6

.PHONY:
