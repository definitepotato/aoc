const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;
const input_file = @embedFile("input.txt");
const allocator = std.heap.page_allocator;

fn isValid(id: u64) bool {
    const id_str = std.fmt.allocPrint(allocator, "{d}", .{id}) catch unreachable;
    if (id_str.len % 2 == 1) return true;

    const half_id_len = id_str.len / 2;
    const left = id_str[0..half_id_len];
    const right = id_str[half_id_len..];

    for (0..half_id_len) |idx| {
        if (left[idx] != right[idx]) {
            return true;
        }
    }

    return false;
}

fn repeat(comptime T: type, slice: []const T, n: usize) []T {
    var result = allocator.alloc(T, slice.len * n) catch unreachable;

    for (0..n) |index| {
        @memcpy(result[index * slice.len .. (index + 1) * slice.len], slice);
    }
    return result;
}

const test_input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";

pub fn main() void {
    var part1: u64 = 0;
    var part2: u64 = 0;

    var ids = std.mem.tokenizeAny(u8, test_input, ",\n");
    while (ids.next()) |id| {
        var range = std.mem.splitScalar(u8, id, '-');
        const start = std.fmt.parseInt(u64, range.next().?, 10) catch unreachable;
        const end = std.fmt.parseInt(u64, range.next().?, 10) catch unreachable;

        for (start..end + 1) |i| {
            if (!isValid(i)) {
                part1 += i;
            }
        }

        for (start..end + 1) |i| {
            const id_str = std.fmt.allocPrint(allocator, "{d}", .{i}) catch unreachable;

            for (1..@divFloor(id_str.len, 2) + 1) |j| {
                const r = repeat(u8, id_str[0..j], @divFloor(id_str.len, j));
                if (@mod(id_str.len, j) == 0 and std.mem.eql(u8, r, id_str)) {
                    part2 += i;
                }
            }
        }
    }

    print("Part 1: {d}\n", .{part1});
    print("Part 2: {d}\n", .{part2});
}
