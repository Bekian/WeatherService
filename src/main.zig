const std = @import("std");
const json = std.json;
const sqlite = @import("sqlite");
const zdt = @import("zdt");

const Forecast = struct {
    latitude: f64,
    longitude: f64,
    generationtime_ms: f64,
    utc_offset_seconds: i32,
    timezone: []const u8,
    timezone_abbreviation: []const u8,
    elevation: f64,
    hourly_units: struct {
        time: []const u8,
        temperature_2m: []const u8,
    },
    hourly: struct {
        time: []const []const u8,
        temperature_2m: []const f32,
    },
};

pub fn main() !void {
    //// Setup
    // (de)init allocator
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer _ = gpa.deinit();
    // (de)init http client
    var client = std.http.Client{ .allocator = allocator };
    defer client.deinit();
    // init db
    var db = try sqlite.Db.init(.{
        .mode = sqlite.Db.Mode{ .File = "forecasts.db" },
        .open_flags = .{
            .write = true,
            .create = true,
        },
    });
    const create_table =
        \\CREATE TABLE forecast (
        \\id SERIAL PRIMARY KEY,
        \\latitude DOUBLE PRECISION NOT NULL,
        \\longitude DOUBLE PRECISION NOT NULL,
        \\generationtime_ms DOUBLE PRECISION NOT NULL,
        \\utc_offset_seconds INTEGER NOT NULL,
        \\timezone TEXT NOT NULL,
        \\timezone_abbreviation TEXT NOT NULL,
        \\elevation DOUBLE PRECISION NOT NULL
        \\);
        \\
        \\CREATE TABLE hourly_units (
        \\forecast_id INTEGER REFERENCES forecast(id) ON DELETE CASCADE,
        \\time TEXT NOT NULL,
        \\temperature_2m TEXT NOT NULL,
        \\PRIMARY KEY (forecast_id)
        \\);
        \\
        \\CREATE TABLE hourly (
        \\id SERIAL PRIMARY KEY,
        \\forecast_id INTEGER REFERENCES forecast(id) ON DELETE CASCADE,
        \\time TEXT NOT NULL,
        \\temperature_2m REAL NOT NULL
        \\);
    ;

    try db.exec(create_table);

    //// Make the web request
    // create uri and header buffer
    const uri = try std.Uri.parse("https://api.open-meteo.com/v1/forecast?latitude=35&longitude=139&hourly=temperature_2m");
    const headerBuffer = try allocator.alloc(u8, 2048);
    defer allocator.free(headerBuffer);
    // open a blocking request using client.open() ; non-blocking requests are created using a std.http.Client.Connection
    var req = try client.open(.GET, uri, .{ .server_header_buffer = headerBuffer });
    defer req.deinit();
    // call methods in order
    try req.send();
    try req.finish();
    try req.wait();
    // read and print response body
    const body = try req.reader().readAllAlloc(allocator, 1024 * 1024);
    defer allocator.free(body);

    // parse json
    const forecast = try json.parseFromSlice(Forecast, allocator, body, .{});
    defer forecast.deinit();
    const val = forecast.value;
    // test parse to datetime which can be converted to string (even though it starts as a string)
    //const timeS = try zdt.Datetime.fromISO8601(val.hourly.time[0]);

}
