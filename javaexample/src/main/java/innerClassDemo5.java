class Outerx {
    public int num  = 10;

    class Inner {
        public int num = 20;
        public void show() {
            int num = 30;
            System.out.println(num);
            System.out.println(this.num);
            //System.out.println(new Outer().num);
            System.out.println(Outerx.this.num);
        }
    }

}
class InnerClassTest {
    public static void main(String[] args) {
        Outerx.Inner oi = new Outerx().new Inner();
        oi.show();
    }
}