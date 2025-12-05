const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;
const ranges = @embedFile("ranges.txt");
const ingredients = @embedFile("ingredients.txt");
var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
const allocator = arena.allocator();

const test_db_input =
    \\3-5
    \\10-14
    \\16-20
    \\12-18
;

const test_ingr_input =
    \\1
    \\5
    \\8
    \\11
    \\17
    \\32
;

fn buildDB(in: []const u8) std.ArrayList([2]usize) {
    var db: std.ArrayList([2]usize) = .empty;

    var toks = std.mem.tokenizeAny(u8, in, "-\n");
    while (toks.next()) |t| {
        const lr = std.fmt.parseInt(usize, t, 10) catch unreachable;
        const rr = std.fmt.parseInt(usize, toks.next().?, 10) catch unreachable;

        db.append(allocator, [2]usize{ lr, rr }) catch unreachable;
    }

    return db;
}

fn lessThan(_: void, lhs: [2]usize, rhs: [2]usize) bool {
    // Sort primarily by the first element (index 0)
    if (lhs[0] != rhs[0]) {
        return lhs[0] < rhs[0];
    }
    // If the first elements are equal, sort by the second element (index 1) as a tie-breaker
    return lhs[1] < rhs[1];
}

pub fn solvePart2(in: std.ArrayList([2]usize)) usize {
    var db = in;
    std.sort.block([2]usize, db.items, {}, lessThan);
    // print("{any}\n", .{db.items});
    var ans: u64 = 0;

    for (db.items, 0..) |_, i| {
        for (db.items, 0..) |_, idx| {
            // if (db.items[i][1] + 1 > db.items[i][0] - 1) continue;

            if (db.items[idx][0] > db.items[i][0] and db.items[idx][0] < db.items[i][1]) {
                db.items[idx][0] = db.items[i][1] + 1;
            }

            if (db.items[idx][1] < db.items[i][1] and db.items[idx][1] > db.items[i][0]) {
                db.items[idx][1] = db.items[i][0] - 1;
            }
        }
    }

    // print("{any}\n", .{db.items});

    for (db.items) |range| {
        if (range[0] > range[1]) continue;
        ans += (range[1] - range[0]) + 1;
    }

    // 361100454171408 <- too high
    return ans;
}

pub fn main() void {
    var part1: usize = 0;
    const db = buildDB(ranges);
    var fresh_db: std.ArrayList([2]usize) = .empty;

    var ingr = std.mem.tokenizeAny(u8, ingredients, "\n");
    while (ingr.next()) |in| {
        var fresh: bool = false;
        const i = std.fmt.parseInt(usize, in, 10) catch unreachable;
        for (db.items) |range| {
            if (i >= range[0] and i <= range[1] and !fresh) {
                fresh = true;
                part1 += 1;

                fresh_db.append(allocator, [2]usize{ range[0], range[1] }) catch unreachable;
            }
        }
    }

    print("Part 1: {d}\n", .{part1});
    print("Part 2: {d}\n", .{solvePart2(db)});
    arena.deinit();
}
