/*
 * This Java source file was generated by the Gradle 'init' task.
 */
package com.example.payservice;

import com.example.payservice.impl.PayServiceImpl;
import com.example.payservice.interceptor.MetadataInterceptor;
import io.grpc.Server;
import io.grpc.ServerBuilder;

import java.io.IOException;

public class App {


    public static void main(String[] args) throws IOException, InterruptedException {

        Server server = ServerBuilder.forPort(10083)
                .addService(new PayServiceImpl())
                .intercept(new MetadataInterceptor())
                .build();

        server.start();
        server.awaitTermination();
    }
}
