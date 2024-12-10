const std = @import("std");
const print = std.debug.print;
const file = @embedFile("input.txt");

fn checksum(fs: std.ArrayList(i64)) i64 {
    var sum: i64 = 0;
    for (fs.items, 0..) |item, idx| {
        if (item == -1) {
            continue;
        }
        const idxAsi64: i64 = @intCast(idx);
        sum += item * idxAsi64;
    }
    return sum;
}

fn nextEmptyIdx(fs: std.ArrayList(i64), upto: usize) ?usize {
    for (fs.items, 0..) |_, idx| {
        if (idx >= upto) {
            return null;
        }

        if (fs.items[idx] == -1) {
            return idx;
        }
    }

    return null;
}

fn moveFile(fs: std.ArrayList(i64)) !void {
    var idx = fs.items.len - 1;

    for (fs.items) |_| {
        if (idx == 0) {
            return;
        }

        if (fs.items[idx] == -1) {
            idx -= 1;
            continue;
        }

        const nextEmptyBlock = nextEmptyIdx(fs, idx);

        if (nextEmptyBlock) |res| {
            fs.items[res] = fs.items[idx];
            fs.items[idx] = -1;
            // print("{any} => Swapping {d}:{d} [{d}]\n", .{ fs.items, res, idx, fs.items[res] });
        }

        idx -= 1;
    }
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var eFS1 = std.ArrayList(i64).init(allocator);
    defer eFS1.deinit();

    var eFS2 = std.ArrayList(i64).init(allocator);
    defer eFS2.deinit();

    const file_system = std.mem.trimRight(u8, file, "\n");

    var file_id: i64 = 0;
    for (0.., file_system) |idx, f| {
        const n = try std.fmt.parseInt(i64, &[_]u8{f}, 10);

        var count_items: i64 = 0;

        if (@mod(idx, 2) != 0) {
            while (count_items < n) : (count_items += 1) {
                try eFS1.append(-1);
                try eFS2.append(-1);
            }
            continue;
        }

        while (count_items < n) : (count_items += 1) {
            try eFS1.append(file_id);
            try eFS2.append(file_id);
        }
        file_id += 1;
    }

    // print("{any}\n", .{eFS1.items});
    try moveFile(eFS1);
    print("Part 1: {d}\n", .{checksum(eFS1)});
}
