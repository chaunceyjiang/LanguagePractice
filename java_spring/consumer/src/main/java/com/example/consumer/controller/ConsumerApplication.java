package com.example.consumer.controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;
import java.util.Map;


@RestController
public class    ConsumerApplication {
//    @Autowired
//    RestTemplate restTemplate;
//
//
//    @LoadBalanced
//    @Bean
//    public RestTemplate rest() {
//        return new RestTemplate();
//    }
//

    /**
     * Rest服务端使用RestTemplate发起http请求,然后得到数据返回给前端
     *
     * @return
//     */
//    @GetMapping(value = "/getUser")
//    @ResponseBody
//    public Map<String, Object> getUser() {
//        Map<String, Object> data = restTemplate.getForObject("http://service-provider/buy", Map.class);
//        return data;
//    }

    @GetMapping(value = "/info")
    @ResponseBody
    public Map<String,Object> getInfo(){
        Map<String,Object> data = new HashMap<>();
        data.put("1",1);
        data.put("2",2);
        return data;
    }
    @GetMapping(value = "/buy")
    @ResponseBody
    public Map<String,Object> getUser(){
        Map<String,Object> data = new HashMap<>();
//        data.put("id",id);
        data.put("size","7#");
        data.put("from","阳光篮球场");
        return data;
    }

}
