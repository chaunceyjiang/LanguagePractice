package com.example.server.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;
import java.util.Map;

@RestController
public class BasketballController {
    /**
     * 假如这个客户端要提供一个getUser的方法
     * @return
     */
    @GetMapping(value = "/buy")
    @ResponseBody
    public Map<String,Object> getUser(){
        Map<String,Object> data = new HashMap<>();
//        data.put("id",id);
        data.put("size","7#");
        data.put("from","阳光篮球场");
        return data;
    }
    @GetMapping(value = "/user")
//    @ResponseBody
    public Map<String,Object> getInfo(){
        Map<String,Object> data = new HashMap<>();
        data.put("xx","xx");
        data.put("xxxxx","xxxx");
        return data;
    }
}