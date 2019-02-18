class MethodDemo {
    public static void main(String[] args) {
        //需求：请写一个功能，实现两个数据的求和
        System.out.println(sum(10, 20));

        //需求：请写一个功能，实现三个数据的求和
        //System.out.println(sum2(10,20,30));
        System.out.println(sum(10, 20, 30));

        System.out.println(sum(1.5, 2.5));
    }

    public static int sum(int a, int b) {
        return a + b;
    }

	/*
	public static int sum2(int a,int b,int c) {
		return a + b + c;
		//方法的嵌套调用
		//return sum(sum(a,b),c);
	}
	*/

    //由于方法名要表达的其实是该方法的作用
    //所以，sum2这个名字不好，还是要用sum 比较好
    public static int sum(int a, int b, int c) {
        return a + b + c;
    }

    public static double sum(double a, double b) {
        return a + b;
    }
}