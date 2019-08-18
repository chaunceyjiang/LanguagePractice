package com.example.demo.user;

import lombok.Data;
import org.springframework.stereotype.Service;

@Service
@Data
public class Bean {
    private String name;
    private Integer age;
}
