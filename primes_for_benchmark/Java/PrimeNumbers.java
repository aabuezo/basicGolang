import java.util.ArrayList;
import java.util.List;

public class PrimeNumbers {

    private static final List<Integer> listOfPrimes;
    private static int count;

    static {
        listOfPrimes = new ArrayList<>();
    }

    public static void main(String[] args) {

        int p = 10000000;
        System.out.println("Prime numbers in " + p + "\nJava");
        long init = System.currentTimeMillis();

        fillListOfPrimes(p);

        System.out.printf("took: %f seconds\n", (float)(System.currentTimeMillis() - init)/1000);

        printListOfPrimes(20);

        System.out.printf("Total primes in %d: %d\n", p, getCount());
    }

    public static void fillListOfPrimes(int number) {

        for (int i = 1; i <= number; i++) {
            if (isPrime(i)) {
                listOfPrimes.add(i);
            }
        }
    }

    public static void printListOfPrimes(int num) {
        for (int i = 0; i < num; i++) {
            System.out.print(listOfPrimes.get(i) + ", ");
        }
        System.out.println("...");
    }

    public static boolean isPrime(int number) {

        int sqr = (int)Math.sqrt(number);
        for (int i = 2; i <= sqr; i++) {
            if (number % i == 0) return false;
        }
        ++count;
        return true;
    }

    public static int getCount() {
        return count;
    }
}

// Q      Java      Go
// 100k-    11 -   4.8 - ms
// 1M  -   143 -   166 - ms
// 10M -  3017 -  3458 - ms
// 20M -  7533 -  8574 - ms
// 50M - 27473 - 31521 - ms ... JIT optimizations?
