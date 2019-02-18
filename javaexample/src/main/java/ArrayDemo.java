public class ArrayDemo {
//    public static void main(String[] args) {
//        //定义一个数组
//        int[] arr = new int[3];
//
//        //赋值，输出。
//        arr[0] = 11;
//        arr[1] = 22;
//        arr[2] = 33;
//        System.out.println(arr);
//        System.out.println("===============");
//        System.out.println(arr[0]);
//        System.out.println(arr[1]);
//        System.out.println(arr[2]);
//
//        //定义第二个数组
//        int[] arr2 =  arr;
//
//        arr2[1] = 100;
//
//        System.out.println(arr);
//        System.out.println(arr[0]);
//        System.out.println(arr[1]); //???
//        System.out.println(arr[2]);
//
//        System.out.println(arr2);
//        System.out.println(arr2[0]);
//        System.out.println(arr2[1]);
//        System.out.println(arr2[2]);
//    }

public static void main(String[] args) {
    //定义一个二维数组
    int[][] arr = new int[3][2];
    //表示arr这个二维数组有三个元素
    //每个元素是一个一维数组
    //每一个一维数组有2个元素

    System.out.println(arr); //[[I@778b3fee
    System.out.println(arr[0]); //[I@57125f92
    System.out.println(arr[1]);
    System.out.println(arr[2]);

    //如何输出元素呢?
    System.out.println(arr[0][1]);
    System.out.println(arr[2][0]);
}
}
