package com.example.WeatherService.models;

import java.util.List;

import jakarta.persistence.ElementCollection;
import jakarta.persistence.Embeddable;
import jakarta.persistence.Embedded;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;

@Entity
public class WeatherData {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private float elevation;
    private double generationTimeMs;
    private double latitude;
    private double longitude;
    private String timezone;
    private String timezoneAbbreviation;
    private int utcOffsetSeconds;

    @Embedded
    private Hourly hourly;

    @Embedded
    private HourlyUnits hourlyUnits;

    // Nested class for Hourly
    @Embeddable
    public static class Hourly {
        @ElementCollection
        private List<Float> temperature2m;

        @ElementCollection
        private List<String> time;

        public List<Float> getTemperature2m() {
            return temperature2m;
        }

        public void setTemperature2m(List<Float> temperature2m) {
            this.temperature2m = temperature2m;
        }

        public List<String> getTime() {
            return time;
        }

        public void setTime(List<String> time) {
            this.time = time;
        }
    }

    // Nested class for HourlyUnits
    @Embeddable
    public static class HourlyUnits {
        private String temperature2m;
        private String time;

        public String getTemperature2m() {
            return temperature2m;
        }

        public void setTemperature2m(String temperature2m) {
            this.temperature2m = temperature2m;
        }

        public String getTime() {
            return time;
        }

        public void setTime(String time) {
            this.time = time;
        }
    }

    // Getters and setters for the main WeatherData class
    public float getElevation() {
        return elevation;
    }

    public void setElevation(float elevation) {
        this.elevation = elevation;
    }

    public double getGenerationTimeMs() {
        return generationTimeMs;
    }

    public void setGenerationTimeMs(double generationTimeMs) {
        this.generationTimeMs = generationTimeMs;
    }

    public Hourly getHourly() {
        return hourly;
    }

    public void setHourly(Hourly hourly) {
        this.hourly = hourly;
    }

    public HourlyUnits getHourlyUnits() {
        return hourlyUnits;
    }

    public void setHourlyUnits(HourlyUnits hourlyUnits) {
        this.hourlyUnits = hourlyUnits;
    }

    public double getLatitude() {
        return latitude;
    }

    public void setLatitude(double latitude) {
        this.latitude = latitude;
    }

    public double getLongitude() {
        return longitude;
    }

    public void setLongitude(double longitude) {
        this.longitude = longitude;
    }

    public String getTimezone() {
        return timezone;
    }

    public void setTimezone(String timezone) {
        this.timezone = timezone;
    }

    public String getTimezoneAbbreviation() {
        return timezoneAbbreviation;
    }

    public void setTimezoneAbbreviation(String timezoneAbbreviation) {
        this.timezoneAbbreviation = timezoneAbbreviation;
    }

    public int getUtcOffsetSeconds() {
        return utcOffsetSeconds;
    }

    public void setUtcOffsetSeconds(int utcOffsetSeconds) {
        this.utcOffsetSeconds = utcOffsetSeconds;
    }

    @Override
    public String toString() {
        return "WeatherData{" +
                "elevation=" + elevation +
                ", generationTimeMs=" + generationTimeMs +
                ", hourly=" + hourly +
                ", hourlyUnits=" + hourlyUnits +
                ", latitude=" + latitude +
                ", longitude=" + longitude +
                ", timezone='" + timezone + '\'' +
                ", timezoneAbbreviation='" + timezoneAbbreviation + '\'' +
                ", utcOffsetSeconds=" + utcOffsetSeconds +
                '}';
    }
}
