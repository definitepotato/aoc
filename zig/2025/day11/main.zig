const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;
const ds = @import("ds.zig");
const input_file = @embedFile("input.txt");
var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
const allocator = arena.allocator();

const test_input =
    \\aaa: you hhh
    \\you: bbb ccc
    \\bbb: ddd eee
    \\ccc: ddd eee fff
    \\ddd: ggg
    \\eee: out
    \\fff: out
    \\ggg: out
    \\hhh: ccc fff iii
    \\iii: out
;

pub fn main() void {
    var stack: ds.Stack([]const u8) = .init(allocator);
    var map: std.StringHashMap([]const u8) = .init(allocator);
    var part1: usize = 0;

    // create map of neighbors.
    var tok_it = std.mem.tokenizeAny(u8, input_file, ":\n");
    while (tok_it.next()) |tok| {
        map.put(tok, tok_it.next().?) catch unreachable;
    }

    // test iterate map.
    // var map_it = map.iterator();
    // while (map_it.next()) |item| {
    //     print("Key: {any}, Value: {any}\n", .{ item.key_ptr.*, item.value_ptr.* });
    // }

    // start from "you" and put its neighbors on a stack.
    const you = map.get("you");
    var neighbors = std.mem.tokenizeAny(u8, you.?, " ");
    while (neighbors.next()) |neighbor| {
        stack.push(neighbor) catch unreachable;
    }

    // as long as there are neighbors, check all paths to "out".
    while (stack.top > 0) {
        // get the next neighbor.
        const neighbor = stack.pop().?;

        if (std.mem.eql(u8, neighbor, "out")) {
            part1 += 1;
            continue;
        }

        // fetch its neighbors and put onto the stack.
        const nodes = map.get(neighbor);
        if (nodes) |node| {
            var nk = std.mem.tokenizeAny(u8, node, " ");
            while (nk.next()) |nv| {
                stack.push(nv) catch unreachable;
            }
        }
    }

    print("Part 1: {d}\n", .{part1});
    arena.deinit();
}
