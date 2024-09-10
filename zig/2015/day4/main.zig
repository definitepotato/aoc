const std = @import("std");
const print = std.debug.print;
const crypto = std.crypto;
const expect = std.testing.expect;

fn mine_string(input: []const u8, target: []const u8) !usize {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const md5 = crypto.hash.Md5;

    var dec: usize = 1;
    while (true) : (dec += 1) {
        const b = try std.fmt.allocPrint(allocator, "{s}{d}", .{ input, dec });
        defer allocator.free(b);

        var buf: [md5.digest_length]u8 = undefined;
        md5.hash(b, &buf, .{});

        const hex = try std.fmt.allocPrint(allocator, "{}", .{std.fmt.fmtSliceHexLower(&buf)});
        defer allocator.free(hex);

        if (std.mem.eql(u8, hex[0..target.len], target)) {
            // print("{s}: {d}\n", .{ hex, dec });
            // break;
            return dec;
        }
    }
}

pub fn main() !void {
    // const bingo = try mine_string("bgvyzdsv", "00000");
    // print("Part 1: {d}\n", .{bingo});
    //
    // const bigger_bingo = try mine_string("bgvyzdsv", "000000");
    // print("Part 2: {d}\n", .{bigger_bingo});

    const input = "bgvyzdsv";

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const md5 = crypto.hash.Md5;

    var dec: usize = 1;
    var part1: bool = true;
    var part2: bool = false;
    while (true) : (dec += 1) {
        const b = try std.fmt.allocPrint(allocator, "{s}{d}", .{ input, dec });
        defer allocator.free(b);

        var buf: [md5.digest_length]u8 = undefined;
        md5.hash(b, &buf, .{});

        const hex = try std.fmt.allocPrint(allocator, "{}", .{std.fmt.fmtSliceHexLower(&buf)});
        defer allocator.free(hex);

        if (part1) {
            if (std.mem.eql(u8, hex[0..5], "00000")) {
                part1 = false;
                part2 = true;
                print("Part 1: {d}\n", .{dec});
            }
        }

        if (part2) {
            if (std.mem.eql(u8, hex[0..6], "000000")) {
                print("Part 2: {d}\n", .{dec});
                break;
            }
        }
    }
}
