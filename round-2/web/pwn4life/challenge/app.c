#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

char flag[256];
void generate_flag()
{
	const char *flag_template = getenv("FLAG");
	if (!flag_template)
	{
		printf("Flag not found! Please create a ticket in Discord.\n");
		exit(1);
	}

	char randomness[16];
	char randomness_hex[33];
	FILE *fp = fopen("/dev/urandom", "rb");
	if (fp)
	{
		fread(randomness, 1, sizeof(randomness), fp);
		fclose(fp);
		for (int i = 0; i < sizeof(randomness); i++)
		{
			sprintf(&randomness_hex[i * 2], "%02x", (unsigned char)randomness[i]);
		}
	}
	randomness_hex[32] = '\0';

	snprintf(flag, sizeof(flag), flag_template, randomness_hex);
}

void win()
{
	printf("Flag: %s\n", flag);
	exit(0);
}

void handle()
{
	int command;
	uint32_t data;
	uint8_t buffer[256];
	while (true)
	{
		printf("Commands:\n");
		printf("1. Write data\n");
		printf("2. Exit\n");
		scanf("%d", &command);

		int16_t offset;
		switch (command)
		{
		case 1:
			printf("Enter offset: ");
			scanf("%hd", &offset);
			printf("Enter data: ");
			scanf("%d", &data);
			*(int *)(buffer + offset) = data;
			printf("Data written!\n");
			break;
		case 2:
			printf("Bye!\n");
			return;
		default:
			printf("Invalid command!\n");
			continue;
		}
	}
}

int main()
{
	generate_flag();
	handle();
	printf("Keep safe out there! dWdnY2Y6Ly9qamoubGJoZ2hvci5wYnovam5ncHU/aT1NMGpRTXh6V0NOdCZ5dmZnPUNZY1VWTFJmc3Q4SFlYMmZUNEZXNDFCYjFYZ2RkbUFZeHEK\n");
	return 0;
}
