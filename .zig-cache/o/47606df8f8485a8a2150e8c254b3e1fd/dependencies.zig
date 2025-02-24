pub const packages = struct {
    pub const @"1220972595d70da33d69d519392742482cb9762935cecb99924e31f3898d2a330861" = struct {
        pub const build_root = "/Users/bekian/.cache/zig/p/1220972595d70da33d69d519392742482cb9762935cecb99924e31f3898d2a330861";
        pub const deps: []const struct { []const u8, []const u8 } = &.{};
    };
    pub const @"1220d7bd38dfca49f4485e8d1c376dbf7a70620e2fd029a11621ed9b5be27534302a" = struct {
        pub const build_root = "/Users/bekian/.cache/zig/p/1220d7bd38dfca49f4485e8d1c376dbf7a70620e2fd029a11621ed9b5be27534302a";
        pub const build_zig = @import("1220d7bd38dfca49f4485e8d1c376dbf7a70620e2fd029a11621ed9b5be27534302a");
        pub const deps: []const struct { []const u8, []const u8 } = &.{
            .{ "sqlite", "1220972595d70da33d69d519392742482cb9762935cecb99924e31f3898d2a330861" },
        };
    };
};

pub const root_deps: []const struct { []const u8, []const u8 } = &.{
    .{ "sqlite", "1220d7bd38dfca49f4485e8d1c376dbf7a70620e2fd029a11621ed9b5be27534302a" },
};
