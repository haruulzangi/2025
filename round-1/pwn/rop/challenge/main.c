#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    double age;
    double kg;
    char name[50];
} Human;

void dummy() {
    asm("pop %rdi; ret;");
}

void init(){
    setbuf(stdin, NULL);
    setbuf(stdout, NULL);
    setbuf(stderr, NULL);
    alarm(180);
}

double calculate_person(Human person[], int num_people) {
    double total_calories = 0.0;
    for (int i = 0; i <= num_people; i++) {
        total_calories += person[i].age * person[i].kg;
    }
    return total_calories;
}

int main() {
    init();

    Human people[3];
    printf("%p\n", (void *)printf);

    for (int i = 0; i <= 3; i++) {
        printf("\nName %d: ", i + 1);
        scanf("%s", people[i].name);

        printf("Age %s: ", people[i].name);
        scanf("%lf", &people[i].age);

        printf("Kg %s: ", people[i].name);
        scanf("%lf", &people[i].kg);
    }

    double total_people = calculate_person(people, 3);

    printf("\nTotal: %.2f kcal\n", total_people);

    return 0;
}