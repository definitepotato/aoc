const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;
const input_file = @embedFile("input.txt");
const allocator = std.heap.page_allocator;

fn isValid(id: u64, memo: *std.AutoHashMap(u64, bool)) bool {
    if (memo.get(id)) |is_valid| {
        return is_valid;
    }

    const id_str = std.fmt.allocPrint(allocator, "{d}", .{id}) catch unreachable;
    if (id_str.len % 2 == 1) return true;

    const half_id_len = id_str.len / 2;
    const left = id_str[0..half_id_len];
    const right = id_str[half_id_len..];

    for (0..half_id_len) |idx| {
        memo.*.put(id, true) catch unreachable;
        if (left[idx] != right[idx]) return true;
    }

    memo.*.put(id, false) catch unreachable;
    return false;
}

const test_input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";

pub fn main() void {
    var acc: u64 = 0;
    var memo: std.AutoHashMap(u64, bool) = .init(allocator);

    var ids = std.mem.tokenizeAny(u8, input_file, ",\n");
    while (ids.next()) |id| {
        var range = std.mem.splitScalar(u8, id, '-');
        const start = std.fmt.parseInt(u64, range.next().?, 10) catch unreachable;
        const end = std.fmt.parseInt(u64, range.next().?, 10) catch unreachable;

        for (start..end + 1) |i| {
            if (!isValid(i, &memo)) {
                acc += i;
            }
        }
    }

    print("Part 1: {d}\n", .{acc});
}
