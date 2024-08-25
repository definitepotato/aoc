const std = @import("std");
const print = std.debug.print;
const input = @embedFile("input.txt");

fn min(a: i32, b: i32, c: i32) i32 {
    var smallest: i32 = std.math.maxInt(i32);
    if (smallest > a) {
        smallest = a;
    }

    if (smallest > b) {
        smallest = b;
    }

    if (smallest > c) {
        smallest = c;
    }

    return smallest;
}

pub fn main() !void {
    var result1: i32 = 0;
    var result2: i32 = 0;
    var lines = std.mem.tokenize(u8, input, "\n");

    while (lines.next()) |line| {
        var l: i32 = 0;
        var h: i32 = 0;
        var w: i32 = 0;

        var dimensions = std.mem.split(u8, line, "x");

        while (dimensions.next()) |dim| {
            if (l == 0) {
                l = try std.fmt.parseInt(i32, dim, 10);
                continue;
            }

            if (h == 0) {
                h = try std.fmt.parseInt(i32, dim, 10);
                continue;
            }

            if (w == 0) {
                w = try std.fmt.parseInt(i32, dim, 10);
                continue;
            }
        }
        const slack = min(l * w, w * h, h * l);
        const area = 2 * ((l * w) + (w * h) + (h * l));
        result1 += slack + area;

        const ribbon = 2 * min(l + w, w + h, h + l);
        const bow = l * w * h;
        result2 += ribbon + bow;
    }

    print("Part 1: {d}\n", .{result1});
    print("Part 1: {d}\n", .{result2});
}
