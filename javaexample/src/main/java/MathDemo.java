import java.util.Arrays;

class MathDemo {
    public static void main(String[] args) {
        //获取随机数
        //double d = Math.random();
        //System.out.println(d);

		/*
		for(int x=0; x<10; x++) {
			//System.out.println(Math.random());
			System.out.println(Math.random()*100);
		}
		*/

        //我们如何获取1-100之间的随机数呢?
        for(int x=0; x<100; x++) {
            int number = (int)(Math.random()*100)+1;
            System.out.println(number);
            System.out.println(Math.round(2.3232));
        }

        System.out.println(
                "ceil:"+Math.ceil(12.04)+"\n"+
                        "floor"+Math.floor(12.04)
        );
        System.out.println(System.currentTimeMillis());



        int[] arr = { 1, 2, 3, 4, 5 };
        int[] arr2 = { 6, 7, 8, 9, 10 };

        System.arraycopy(arr, 1, arr2, 2, 2);

        System.out.println(Arrays.toString(arr));//[1, 2, 3, 4, 5]

        System.out.println(Arrays.toString(arr2));//[6, 7, 2, 3, 10]
    }
}