const std = @import("std");
const expect = std.testing.expect;
const print = std.debug.print;
const hash = std.hash;
const input = @embedFile("input.txt");

// returns number of vowels in `s`
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

test "num vowels" {
    try expect(num_vowels("abcdefghijklmnopqrstuvwxyz") == 5);
    try expect(num_vowels("aachzlrnu") == 3);
    try expect(num_vowels("gwjzpxmbrrgsdvwm") == 0);
}

// true if `s` contains at least one letter that appears twice in a row
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

test "has doubles" {
    try expect(has_doubles("ugknbfddgicrmopn") == true);
    try expect(has_doubles("ugknbfxygicrmopn") == false);
}

// true if `s` contains the string 'ab' 'cd' 'pq' 'xy'
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

test "has naughty string" {
    try expect(has_naughty_string("ugknbfddgicrmopn") == false);
    try expect(has_naughty_string("ugknbfdpqicrmopn") == true);
}

// true if `s` contains one letter which repeats with exactly one letter between them
// i.e., `xyx` or `aaa`
fn has_two_nice_letters(s: []const u8) bool {
    for (0..s.len) |idx| {
        if (idx + 2 > s.len - 1) {
            break;
        }

        if (s[idx] == s[idx + 2]) {
            return true;
        }
    }

    return false;
}

test "has two nice letters" {
    try expect(has_two_nice_letters("xyx") == true);
    try expect(has_two_nice_letters("aaa") == true);
    try expect(has_two_nice_letters("abc") == false);
}

// scans `s` for `target` pair starting from index `start`
fn has_pair_from(s: []const u8, start: usize, target: [2]u8) bool {
    var window_left: usize = start;
    var window_right: usize = window_left + 1;

    for (s.len) |_| {
        if (window_right > s.len - 1) {
            break;
        }

        if (s[window_left] == target[0]) {
            if (s[window_right] == target[1]) {
                return true;
            }
        }

        window_left += 1;
        window_right += 1;
    }

    return false;
}

// true if `s` contains a pair of any two letters that appears at least twice
// without overlapping like 'xyxy' or 'aabcdefgaa'
fn has_nice_pair(s: []const u8) bool {
    var window_left: usize = 0;
    var window_right: usize = 1;

    for (s.len) |_| {
        if (window_right > s.len - 1) {
            break;
        }

        const result = has_pair_from(s, window_right + 1, [2]u8{ s[window_left], s[window_right] });
        if (result == true) {
            return true;
        }

        window_left += 1;
        window_right += 1;
    }

    return false;
}

test "has nice pairs" {
    try expect(has_nice_pair("xyxy") == true);
    try expect(has_nice_pair("aabcdefgaa") == true);
    try expect(has_nice_pair("abcdefghijklmnoopp") == false);
    try expect(has_nice_pair("aabcdeffghi") == false);
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

fn is_actually_nice(s: []const u8) bool {
    // Must contain two letters that appear twice without overlapping.
    if (has_nice_pair(s)) {
        // Must contain at least one letter which repeats with exactly one letter between them.
        if (has_two_nice_letters(s)) {
            return true;
        }
    }

    return false;
}

pub fn main() !void {
    var nice_lines: usize = 0;
    var actually_nice_lines: usize = 0;

    var lines = std.mem.tokenize(u8, input, "\n");
    while (lines.next()) |line| {
        if (is_nice(line)) {
            nice_lines += 1;
        }

        if (is_actually_nice(line)) {
            actually_nice_lines += 1;
        }
    }

    print("Part 1: {d}\n", .{nice_lines});
    print("Part 2: {d}\n", .{actually_nice_lines});
}
