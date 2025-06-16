const std = @import("std");
const print = std.debug.print;
const input = @embedFile("input.txt");

const sample_input =
    \\ 123 -> x
    \\ 456 -> y
    \\ x AND y -> d
    \\ x OR y -> e
    \\ x LSHIFT 2 -> f
    \\ y RSHIFT 2 -> g
    \\ NOT x -> h
    \\ NOT y -> i
;

const OPS = enum {
    OP_ERR,
    OP_AND,
    OP_OR,
    OP_NOT,
    OP_LSHIFT,
    OP_RSHIFT,
    OP_STORE,
};

pub fn main() void {
    var line = std.mem.tokenizeAny(u8, sample_input, "\n");

    while (line.next()) |t| {
        var ins = std.mem.tokenizeAny(u8, t, " ");

        while (ins.next()) |i| {
            print("{s}\n", .{i});
        }
    }

    // var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    // defer _ = gpa.deinit();
    // const allocator = gpa.allocator();
    //
    // var map = std.StringHashMap(bool).init(allocator);
    // defer map.deinit();
    //
    // map.put("test1", true) catch unreachable;
    // map.put("test2", true) catch unreachable;
    //
    // const some_value: []const u8 = "test1";
    // std.debug.print("{any}\n", .{map.get(some_value)});
    //
    // var map_iter = map.iterator();
    // while (map_iter.next()) |entry| {
    //     std.debug.print("{s} -> {}\n", .{ entry.key_ptr.*, entry.value_ptr.* });
    // }
}
