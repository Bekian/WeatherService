package com.example.WeatherService.config;

import org.hibernate.community.dialect.SQLiteDialect;

public class CustomSQLiteDialect extends SQLiteDialect {
    public CustomSQLiteDialect() {
        super();
    }
}
