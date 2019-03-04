package com.example.demo;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HelloWorldController {
    @RequestMapping("/hello")
    public String hello(){
        return "Hello world";
    }
    @RequestMapping("/getuser")
    public User getUser(){
        User user = new User();
        user.setAge(1);
        user.setName("abc");
        return user;
    }
}
