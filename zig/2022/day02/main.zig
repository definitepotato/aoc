const std = @import("std");
const print = std.debug.print;

fn shoot_again(a: u8, b: u8) u32 {
    var score: u32 = 0;

    switch (b) {
        'X' => score += 0, // lose
        'Y' => score += 3, // draw
        'Z' => score += 6, // win
        else => unreachable,
    }

    if (a == 'A') { // rock
        switch (b) {
            'X' => score += 3, // scissors
            'Y' => score += 1, // rock
            'Z' => score += 2, // paper
            else => unreachable,
        }
    }

    if (a == 'B') { // paper
        switch (b) {
            'X' => score += 1, // rock
            'Y' => score += 2, // paper
            'Z' => score += 3, // scissors
            else => unreachable,
        }
    }

    if (a == 'C') { // scissors
        switch (b) {
            'X' => score += 2, // paper
            'Y' => score += 3, // scissors
            'Z' => score += 1, // rock
            else => unreachable,
        }
    }

    // print("{c}:{c} => {d}\n", .{ a, b, score });
    return score;
}

fn shoot(a: u8, b: u8) u32 {
    var score: u32 = 0;

    switch (b) {
        'X' => score += 1, // rock
        'Y' => score += 2, // paper
        'Z' => score += 3, // scissor
        else => unreachable,
    }

    if (a == 'A') { // rock
        switch (b) {
            'X' => score += 3, // rock,draw
            'Y' => score += 6, // paper,win
            'Z' => score += 0, // scissors,lose
            else => unreachable,
        }
    }

    if (a == 'B') { // paper
        switch (b) {
            'X' => score += 0, // rock,lose
            'Y' => score += 3, // paper,draw
            'Z' => score += 6, // scissors,win
            else => unreachable,
        }
    }

    if (a == 'C') { // scissors
        switch (b) {
            'X' => score += 6, // rock,win
            'Y' => score += 0, // paper,lose
            'Z' => score += 3, // scissors,draw
            else => unreachable,
        }
    }

    // print("{c}:{c} => {d}\n", .{ a, b, score });
    return score;
}

fn read_file(allocator: std.mem.Allocator, target: []const u8) ![]u8 {
    const file = try std.fs.cwd().openFile(target, .{});
    defer file.close();

    const stat = try file.stat();
    const buff = try file.readToEndAlloc(allocator, stat.size);
    return buff;
}

pub fn main() !void {
    var score_part1: u32 = 0;
    var score_part2: u32 = 0;
    const allocator = std.heap.page_allocator;

    // allocate file to heap.
    const buff = try read_file(allocator, "input.txt");
    defer allocator.free(buff);

    // split file by newline into iterator.
    var lines = std.mem.split(u8, buff, "\n");

    // iterate split file "lines".
    while (lines.next()) |line| {
        if (line.len > 0) {
            score_part1 += shoot(line[0], line[2]);
            score_part2 += shoot_again(line[0], line[2]);
        }
    }

    print("Part 1: {d}\n", .{score_part1});
    print("Part 2: {d}\n", .{score_part2});
}
