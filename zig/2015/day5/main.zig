const std = @import("std");
const print = std.debug.print;
const hash = std.hash;
const input = @embedFile("input.txt");

fn num_vowels(s: []const u8) u32 {
    var vowels_count: u32 = 0;

    for (s) |letter| {
        switch (letter) {
            'a' => vowels_count += 1,
            'e' => vowels_count += 1,
            'i' => vowels_count += 1,
            'o' => vowels_count += 1,
            'u' => vowels_count += 1,
            else => continue,
        }
    }
    return vowels_count;
}

fn has_doubles(s: []const u8) bool {
    for (0..s.len) |idx| {
        if (idx + 1 > s.len - 1) {
            break;
        }

        if (s[idx] == s[idx + 1]) {
            return true;
        }
    }

    return false;
}

fn has_naughty_string(s: []const u8) bool {
    for (0..s.len) |idx| {
        if (idx + 1 > s.len - 1) {
            break;
        }

        if (s[idx] == 'a') {
            if (s[idx + 1] == 'b') {
                return true;
            }
        }

        if (s[idx] == 'c') {
            if (s[idx + 1] == 'd') {
                return true;
            }
        }

        if (s[idx] == 'p') {
            if (s[idx + 1] == 'q') {
                return true;
            }
        }

        if (s[idx] == 'x') {
            if (s[idx + 1] == 'y') {
                return true;
            }
        }
    }

    return false;
}

fn is_nice(s: []const u8) bool {
    // Must contain at least 3 vowels.
    if (num_vowels(s) < 3) {
        return false;
    }

    // Cannot contain 'ab', 'cd', 'pq', 'xy'.
    if (has_naughty_string(s)) {
        return false;
    }

    // Must have at least one double.
    if (has_doubles(s) == false) {
        return false;
    }

    return true;
}

pub fn main() !void {
    var nice_lines: i32 = 0;

    var lines = std.mem.tokenize(u8, input, "\n");
    while (lines.next()) |line| {
        if (is_nice(line)) {
            nice_lines += 1;
        }
    }

    print("Part 1: {d}\n", .{nice_lines});
}
