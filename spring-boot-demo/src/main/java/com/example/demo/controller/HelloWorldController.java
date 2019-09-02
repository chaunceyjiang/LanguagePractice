package com.example.demo.controller;

import com.example.demo.user.User;
import com.wordnik.swagger.annotations.Api;
import com.wordnik.swagger.annotations.ApiOperation;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;
import java.util.Map;

@Slf4j
@RestController
@Api(tags = "Hello")
public class HelloWorldController {
    @Autowired
    private User user;
    @RequestMapping("/hello")
    public String hello(){
        return "Hello world";
    }

    @ApiOperation(value = "getsuer",notes = "获取用户信息")
    @RequestMapping("/getuser")
    public User getUser(){
        log.info("ssss"+user.toString());
        user.setAge(1);
        user.setName("abc");
        return user;
    }
    @RequestMapping("/info")
    public Map<String, Object> getInfo(){
        log.info("ssss"+user.toString());
        Map<String,Object> info = new HashMap<>();
        info.put("zz","1");
        info.put("yy","2");
        return info;
    }
}
