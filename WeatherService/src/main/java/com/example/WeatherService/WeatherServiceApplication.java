package com.example.WeatherService;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

import org.springframework.boot.autoconfigure.SpringBootApplication;

import com.example.WeatherService.models.WeatherData;
import com.fasterxml.jackson.databind.ObjectMapper;

@SpringBootApplication
public class WeatherServiceApplication {

	public static void main(String[] args) {
        // SpringApplication.run(WeatherServiceApplication.class, args);

        try {
            // Create an HttpClient instance
            HttpClient client = HttpClient.newHttpClient();

            // Create a GET request to the URL
            HttpRequest request = HttpRequest.newBuilder()
                    .uri(new URI(
                            "https://api.open-meteo.com/v1/forecast?latitude=44.94&longitude=-93.10&hourly=temperature_2m&models=gfs_seamless"))
                    .build();

            // Send the request and get the response
            HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
            ObjectMapper objectMapper = new ObjectMapper();
            WeatherData weatherData = objectMapper.readValue(response.body(), WeatherData.class);

            // Print the response status and body
            System.out.println("Response Code: " + response.statusCode());
            System.out.println("Response Body: " + response.body());
        } catch (Exception e) {
            e.printStackTrace();
        }

	}

}
