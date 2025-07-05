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
    allocator: std.mem.Allocator,
    visited: std.ArrayList(vec),
    overlap: std.ArrayList(vec),

    pub fn init(allocator: std.mem.Allocator) Pos {
        return .{
            .x = 0,
            .y = 0,
            .dir = .North,
            .allocator = allocator,
            .visited = std.ArrayList(vec).init(allocator),
            .overlap = std.ArrayList(vec).init(allocator),
        };
    }

    pub fn deinit(self: *Pos) void {
        self.visited.deinit();
        self.overlap.deinit();
    }

    pub fn seen(self: *Pos, pos: vec) bool {
        if (self.visited.items.len == 0) {
            return false;
        }

        for (self.visited.items) |item| {
            if (std.meta.eql(item, pos)) {
                return true;
            }
        }

        return false;
    }

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

    pub fn move(self: *Pos, steps: usize) !void {
        if (self.dir == .North) {
            for (0..steps) |_| {
                self.y += 1;
                const pos = vec{ .x = self.x, .y = self.y };
                if (self.seen(pos)) {
                    try self.overlap.append(pos);
                }
                try self.visited.append(pos);
            }
        }

        if (self.dir == .East) {
            for (0..steps) |_| {
                self.x += 1;
                const pos = vec{ .x = self.x, .y = self.y };
                if (self.seen(pos)) {
                    try self.overlap.append(pos);
                }
                try self.visited.append(pos);
            }
        }

        if (self.dir == .South) {
            for (0..steps) |_| {
                self.y -= 1;
                const pos = vec{ .x = self.x, .y = self.y };
                if (self.seen(pos)) {
                    try self.overlap.append(pos);
                }
                try self.visited.append(pos);
            }
        }

        if (self.dir == .West) {
            for (0..steps) |_| {
                self.x -= 1;
                const pos = vec{ .x = self.x, .y = self.y };
                if (self.seen(pos)) {
                    try self.overlap.append(pos);
                }
                try self.visited.append(pos);
            }
        }
    }

    pub fn distance(self: *Pos) u32 {
        return @abs(self.x) + @abs(self.y);
    }
};

test "turn position" {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var my_pos = Pos.init(allocator);
    defer my_pos.deinit();

    my_pos.turn('R');
    try std.testing.expect(my_pos.dir == .East);
}

test "move position" {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var my_pos = Pos.init(allocator);
    defer my_pos.deinit();

    try my_pos.move(2.0);
    try std.testing.expect(my_pos.y == 2);

    my_pos.turn('L');
    try my_pos.move(2.0);
    try std.testing.expect(my_pos.x == -2);
}

const vec = struct {
    x: i32,
    y: i32,

    pub fn distance(self: *vec) u32 {
        return @abs(self.x) + @abs(self.y);
    }
};

const test_input = "R8, R4, R4, R8";

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var my_pos = Pos.init(allocator);
    defer my_pos.deinit();

    var lines = std.mem.tokenizeAny(u8, input, ", \n");
    while (lines.next()) |line| {
        const int_move = try std.fmt.parseInt(usize, line[1..], 10);

        my_pos.turn(line[0]);
        try my_pos.move(int_move);
    }

    print("{d}\n", .{my_pos.distance()});
    print("{d}\n", .{my_pos.overlap.items[0].distance()});
}
