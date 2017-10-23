public class hello {
    private static String foo() {
        return "different≠another";
    }

    public static void main (String[] args) {
	String x = "Hello World!";
        if (true) {
	    x = foo();
            System.out.println(x);
        }
        System.out.println(x);
    }
}
