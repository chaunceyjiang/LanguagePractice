package com.example.demo;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import springfox.documentation.builders.ApiInfoBuilder;
import springfox.documentation.builders.PathSelectors;
import springfox.documentation.builders.RequestHandlerSelectors;
import springfox.documentation.service.ApiInfo;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

@Configuration
@EnableSwagger2
public class Swagger2Configration {

        @Bean
        public Docket createRestApi() {

            return new Docket(DocumentationType.SWAGGER_2)

                    .apiInfo(apiInfo())

                    .select()

                    .apis(RequestHandlerSelectors.basePackage("com.example"))

                    .paths(PathSelectors.any())

                    .build();

        }

        private ApiInfo apiInfo() {

            return new ApiInfoBuilder().title("SIMO Report RESTful APIs")

                    .description("SIMO Report 的 管理/测试 API")

                    .version("SIMO 2.4")

                    .build();

        }

}
