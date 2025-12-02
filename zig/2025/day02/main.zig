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
        if (left[idx] != right[idx]) return true;
    }

    return false;
}

const test_input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";

pub fn main() void {
    var acc: u64 = 0;

    var ids = std.mem.tokenizeAny(u8, input_file, ",\n");
    while (ids.next()) |id| {
        var start_end = std.mem.splitScalar(u8, id, '-');
        const start = start_end.next().?;
        const end = start_end.next().?;

        const start_int = std.fmt.parseInt(u64, start[0..], 10) catch unreachable;
        const end_int = std.fmt.parseInt(u64, end[0..], 10) catch unreachable;

        for (start_int..end_int + 1) |i| {
            if (!isValid(i)) {
                acc += i;
            }
        }
    }

    print("Part 1: {d}\n", .{acc});
}
