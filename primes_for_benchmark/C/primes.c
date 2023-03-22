#include <stdio.h>
#include <time.h>
#include <math.h>

#define TRUE	1
#define FALSE	0
#define SIZE	700000	// There are no more than ~7% of primes in N

int number = 10000000;	// N
int total = 0;
int lst[SIZE];


void init_array() {
	for (int i = 0; i < SIZE; i++) {
		lst[i] = 0;
	}
}


int is_prime(int num) {
	int sqr = (int)sqrt(num);
	for (int i = 2; i <= sqr; i++) {
		if (num % i == 0) {
			return FALSE;
		} 
	}
	return TRUE;
}


void fill_list_of_primes(int number) {
	for (int i = 1; i <= number; i++) {
		if (is_prime(i)) {
			lst[total++] = i;
		}
	}
}


void print_list_of_primes(int q) {
	for (int i = 0; i < q; i++) {
		printf("%d, ", lst[i]);
	}
	printf("...\n");
}


int main() {

	printf("Prime numbers in %d\nC language\n", number);
	clock_t start, end;
    double cpu_time_used;

    start = clock();
 
 	init_array();	
	fill_list_of_primes(number);
	 
	end = clock();
    cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
	printf("took: %f seconds\n", cpu_time_used);

	print_list_of_primes(20);
	printf("Total primes in %d: %d\n", number, total);
}
