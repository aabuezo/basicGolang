import math
import time

number = 10000000
lst = []
total = 0


def fill_primes_list(num):
    global lst, total
    for i in range(1, num):
        if is_prime(i):
            total += 1
            lst.append(i)


def is_prime(num):
    sqr = int(math.sqrt(num))
    for i in range(2, sqr + 1):
        if num % i == 0:
            return False
    return True


def print_list_of_primes(n):
    for i in range(0, n):
        print(lst[i], end=", ")
    print("...")


if __name__ == '__main__':
    print(f"Prime numbers in {number}\nPython")
    start = time.time()
    fill_primes_list(number)
    end = time.time()
    print("took: ", end - start, "seconds")
    print_list_of_primes(20)
    print(f"Total primes in {number}: {total}")
