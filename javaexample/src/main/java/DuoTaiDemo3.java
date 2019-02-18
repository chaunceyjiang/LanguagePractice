class  Animal {
    public void eat(){}
}

class Dog extends Animal {
    public void eat(){
        System.out.println("狗吃骨头");
    }

    public void help() {
        System.out.println("狗可以帮助警察叔叔抓小偷");
    }
}

class Cat extends Animal {
    public void eat(){
        System.out.println("猫吃老鼠");
    }
}


class DuoTaiDemo2 {
    public static void main(String[] args) {
        //多态
        Animal a = new Dog();    //向上转型
        a.eat();
        //a.help(); //编译看左边
        //向下转型
        Dog d = (Dog) a;
        d.eat();
        d.help();
        System.out.println("-------------");

        //会报错吗?
        a = new Cat(); ////向上转型
        a.eat();

        Cat c = (Cat) a;
        c.eat();

        //Dog dd = (Dog)a; //ClassCastException
        //dd.eat();
        //dd.help();
    }
}