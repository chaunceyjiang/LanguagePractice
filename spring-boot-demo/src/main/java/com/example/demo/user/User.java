package com.example.demo.user;


import com.wordnik.swagger.annotations.ApiModel;
import lombok.AllArgsConstructor;
import lombok.Data;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Service;

@Data
@ApiModel(value = "用户对象")
@Service
public class User {
    private String name;
    private Integer age;
}
