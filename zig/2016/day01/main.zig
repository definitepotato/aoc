const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;
const input = @embedFile("input.txt");

const direction = enum {
    North,
    South,
    East,
    West,
};

const Pos = struct {
    x: i32,
    y: i32,
    dir: direction,

    pub fn turn(self: *Pos, turn_dir: u8) void {
        if (turn_dir == 'R') {
            switch (self.dir) {
                .North => self.dir = .East,
                .South => self.dir = .West,
                .East => self.dir = .South,
                .West => self.dir = .North,
            }
        }

        if (turn_dir == 'L') {
            switch (self.dir) {
                .North => self.dir = .West,
                .South => self.dir = .East,
                .East => self.dir = .North,
                .West => self.dir = .South,
            }
        }
    }

    pub fn move(self: *Pos, steps: i32) void {
        if (self.dir == .North) {
            self.y += steps;
        }

        if (self.dir == .East) {
            self.x += steps;
        }

        if (self.dir == .South) {
            self.y -= steps;
        }

        if (self.dir == .West) {
            self.x -= steps;
        }
    }

    pub fn distance(self: *Pos) u32 {
        return @abs(self.x) + @abs(self.y);
    }
};

test "turn position" {
    var my_pos = Pos{
        .x = 0,
        .y = 0,
        .dir = .North,
    };

    my_pos.turn('R');
    try std.testing.expect(my_pos.dir == .East);
}

test "move position" {
    var my_pos = Pos{
        .x = 0,
        .y = 0,
        .dir = .North,
    };

    my_pos.move(2.0);
    try std.testing.expect(my_pos.y == 2);

    my_pos.turn('L');
    my_pos.move(2.0);
    try std.testing.expect(my_pos.x == -2);
}

const vec = struct {
    x: i32,
    y: i32,
};

const test_input = "R5, L5, R5, R3";

pub fn main() !void {
    var my_pos = Pos{
        .x = 0,
        .y = 0,
        .dir = .North,
    };

    var lines = std.mem.tokenizeAny(u8, input, ", \n");
    while (lines.next()) |line| {
        // const int_move: u8 = line[1] - '0'; // convert u8 to integer
        const int_move = try std.fmt.parseInt(i32, line[1..], 10);

        my_pos.turn(line[0]);
        my_pos.move(int_move);
        // print("{s}: {}\n", .{ line, my_pos });
    }

    print("{d}\n", .{my_pos.distance()});
}
