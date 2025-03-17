const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;
const input = @embedFile("input.txt");

const Action = enum {
    On,
    Off,
    Toggle,
};

const Lights = struct {
    const Self = @This();

    yx: [1000][1000]bool,
    brightness: [1000][1000]usize,

    pub fn init() Lights {
        return Lights{
            .yx = std.mem.zeroes([1000][1000]bool),
            .brightness = std.mem.zeroes([1000][1000]usize),
        };
    }

    pub fn count_lights(self: Self) usize {
        var count: usize = 0;

        for (self.yx) |ly| {
            for (ly) |lx| {
                if (lx == true) {
                    count += 1;
                }
            }
        }

        return count;
    }

    pub fn measure_brightness(self: Self) usize {
        var count: usize = 0;

        for (self.brightness) |y| {
            for (y) |x| {
                count += x;
            }
        }

        return count;
    }

    pub fn lights(self: *Self, from: [2]usize, to: [2]usize, action: Action) void {
        const x_start: usize = from[0];
        const x_end: usize = to[0];
        const y_start: usize = from[1];
        const y_end: usize = to[1];

        for (y_start..y_end + 1) |y| {
            for (x_start..x_end + 1) |x| {
                switch (action) {
                    .On => {
                        self.yx[x][y] = true;
                        self.brightness[x][y] += 1;
                    },
                    .Off => {
                        self.yx[x][y] = false;
                        if (self.brightness[x][y] > 0) {
                            self.brightness[x][y] -= 1;
                        }
                    },
                    .Toggle => {
                        self.yx[x][y] = !self.yx[x][y];
                        self.brightness[x][y] += 2;
                    },
                }
            }
        }
    }
};

test "count lights" {
    var lights = Lights.init();

    lights.yx[0][0] = true;
    lights.yx[1][1] = true;

    assert(lights.count_lights() == 2);
}

test "lights on" {
    var lights = Lights.init();

    lights.lights([2]usize{ 0, 0 }, [2]usize{ 2, 2 }, .On);
    assert(lights.count_lights() == 9);
}

test "lights off" {
    var lights = Lights.init();

    lights.lights([2]usize{ 0, 0 }, [2]usize{ 2, 2 }, .On);
    lights.lights([2]usize{ 0, 0 }, [2]usize{ 1, 1 }, .Off);
    assert(lights.count_lights() == 5);
}

test "lights toggle" {
    var lights = Lights.init();

    lights.lights([2]usize{ 0, 0 }, [2]usize{ 2, 2 }, .On);
    lights.lights([2]usize{ 0, 0 }, [2]usize{ 1, 1 }, .Toggle);
    assert(lights.count_lights() == 5);
}

const Instruction = struct {
    action: Action,
    start_x: usize,
    start_y: usize,
    end_x: usize,
    end_y: usize,
};

pub fn parse_coordinates(coord: []const u8) ![2]usize {
    var x: []const u8 = undefined;
    var y: []const u8 = undefined;

    var coordinates = std.mem.tokenizeAny(u8, coord, ",");
    while (coordinates.next()) |c| {
        x = c;
        y = coordinates.next().?;
    }

    const x_usize = try std.fmt.parseInt(usize, x, 10);
    const y_usize = try std.fmt.parseInt(usize, y, 10);

    return [2]usize{ x_usize, y_usize };
}

test "parse coordinates" {
    const xy = try parse_coordinates("628,958");

    assert(xy[0] == 628);
    assert(xy[1] == 958);
}

pub fn process_instruction(instr: []const u8) !Instruction {
    var parsed_instruction = Instruction{
        .action = undefined,
        .start_x = undefined,
        .start_y = undefined,
        .end_x = undefined,
        .end_y = undefined,
    };

    // Get words in line.
    var words = std.mem.tokenizeAny(u8, instr, " ");
    while (words.next()) |word| {

        // Check if toggle.
        if (std.mem.eql(u8, word, "toggle")) {
            parsed_instruction.action = .Toggle;
        }

        if (std.mem.eql(u8, word, "turn")) {
            const next = words.next();
            // Check if turn on.
            if (std.mem.eql(u8, next.?, "on")) {
                parsed_instruction.action = .On;
            }
            // Check if turn off.
            if (std.mem.eql(u8, next.?, "off")) {
                parsed_instruction.action = .Off;
            }
        }

        const start_xy = try parse_coordinates(words.next().?);
        _ = words.next();
        const end_xy = try parse_coordinates(words.next().?);

        parsed_instruction.start_x = start_xy[0];
        parsed_instruction.start_y = start_xy[1];
        parsed_instruction.end_x = end_xy[0];
        parsed_instruction.end_y = end_xy[1];
    }

    return parsed_instruction;
}

pub fn main() !void {
    var lights = Lights.init();

    // Get lines.
    var lines = std.mem.tokenizeAny(u8, input, "\n");
    while (lines.next()) |line| {
        const instr = try process_instruction(line);
        lights.lights([2]usize{ instr.start_x, instr.start_y }, [2]usize{ instr.end_x, instr.end_y }, instr.action);
    }

    print("Part 1: {d}\n", .{lights.count_lights()});
    print("Part 2: {d}\n", .{lights.measure_brightness()});
}
