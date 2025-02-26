const std = @import("std");
const json = std.json;
// I'm calling this done as of 2/5/'25 cause idk how to import the popular zig sqlite libs and im not gonna bother figuring it out.
// I want to work on other projects.

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
    //std.debug.print("Response body: {s}\n", .{body});

    // parse json
    const forecast = try json.parseFromSlice(Forecast, allocator, body, .{});
    defer forecast.deinit();
    const val = forecast.value;

    std.debug.print("{any}", .{val});
}
