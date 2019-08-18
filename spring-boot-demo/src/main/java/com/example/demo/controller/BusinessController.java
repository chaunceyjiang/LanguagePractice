package com.example.demo.controller;

import com.example.demo.user.Bean;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(value = "/business")
public class BusinessController {
    @Autowired
    private Bean bean;

    @GetMapping(value = "/user1")
    public Bean getUser1() {
        bean.setAge(10);
        bean.setName("user");
        return bean;
    }
}
