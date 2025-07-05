const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;
const input = @embedFile("input.txt");

const Lock = struct {
    x: i8,
    y: i8,

    pub fn init() Lock {
        return .{
            .x = 1,
            .y = 1,
        };
    }

    pub fn getButton(self: *Lock) usize {
        return @intCast(self.y * 3 + self.x + 1);
    }

    pub fn move(self: *Lock, dir: u8) void {
        if (dir == 'U') self.y -= 1;
        if (dir == 'R') self.x += 1;
        if (dir == 'D') self.y += 1;
        if (dir == 'L') self.x -= 1;

        if (self.y < 0) self.y = 0;
        if (self.y > 2) self.y = 2;
        if (self.x < 0) self.x = 0;
        if (self.x > 2) self.x = 2;
    }
};

test "get button" {
    var lock = Lock.init();
    try std.testing.expect(lock.getButton() == 5);

    lock.move('R');
    try std.testing.expect(lock.getButton() == 6);

    lock.move('D');
    try std.testing.expect(lock.getButton() == 9);

    lock.move('R');
    try std.testing.expect(lock.getButton() == 9);
}

const test_input =
    \\ULL
    \\RRDDD
    \\LURDL
    \\UUUUD
;

pub fn main() void {
    var lock = Lock.init();

    var lines = std.mem.tokenizeAny(u8, input, "\n");
    while (lines.next()) |line| {
        for (line) |instr| {
            lock.move(instr);
        }
        print("{d}\n", .{lock.getButton()});
    }
}
