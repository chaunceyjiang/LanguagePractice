import ognl.Ognl;
import ognl.OgnlContext;
import ognl.OgnlException;

import java.util.HashMap;
import java.util.Map;

public class Test {
    public static void main(String[] args) throws OgnlException {
        // 准备ONGL的上下文
        // 1、准备Root --- 将一个User对象放入到Root栈中
//        User root_user = new User("张三", 18);
//        // 2、准备context --- 将User对象以键值对的形式存入Context栈中
//        Map<String, User> context = new HashMap<String, User>();
//        context.put("User1", new User("jeck", 18));
//        context.put("User2", new User("rose", 22));
//        OgnlContext og = new OgnlContext();
//        og.setRoot(root_user);
//        og.setValues(context);
//
//        // 书写ONGL表达式，此表达式目的---取出root_user中的name、age的值
//        String name = (String)Ognl.getValue("name", og, og.getValues());
//        Integer age = (Integer) Ognl.getValue("age", og, og.getValues());
//        System.out.println(name + "," + age);
        useMap();
    }

    public static void useMap() throws OgnlException {
        Map<Object, Object> map = new HashMap<Object, Object>();
        map.put("name", "张三");
        map.put("age", 18);
        map.put("MesgRefId", "       sdiuagslucasdaiushlaiu   ");
        map.put("CnccOrigSendDate", "10000000");
//        map.put("CnccOrigSendTime",null);
        map.put("sendtime", "");
        map.put("RRA", "req");
        Map<String, Map<Object, Object>> context = new HashMap<String, Map<Object, Object>>();
        context.put("user", new HashMap<Object, Object>());

        OgnlContext og = new OgnlContext();
        og.setRoot(map);
        og.setValues(context);
        String name = (String) Ognl.getValue("name.substring(0, 1) in \"张\" ? \"三\" : \"a\"", og, og.getRoot());
        Integer age = (Integer) Ognl.getValue("age", og, og.getRoot());
        String MesgRefId = (String) Ognl.getValue("MesgRefId.trim()", og, og.getRoot());
        String sendtime = (String) Ognl.getValue("sendtime != \"\" ? sendtime :  (CnccOrigSendDate and CnccOrigSendTime) ?(CnccOrigSendDate+CnccOrigSendTime) : \"--\"", og, og.getRoot());
        Boolean rra = (Boolean) Ognl.getValue("RRA==\"req\"", og, og.getRoot());
        System.out.println(name + "," + age + "," + MesgRefId + "," + sendtime + "," + rra.toString());

    }
}
