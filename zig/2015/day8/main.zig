const std = @import("std");
const print = std.debug.print;
const input = @embedFile("input.txt");
const expect = std.testing.expect;

const test_input = "aaa\"aaa";

pub fn is_char(c: u8) bool {
    // lowercase [a-z]
    if (c >= 97 and c <= 122) {
        return true;
    }

    // uppercase [A-Z]
    if (c >= 65 and c <= 90) {
        return true;
    }

    return false;
}

pub fn is_num(c: u8) bool {
    return c >= 48 and c <= 57; // [0-9]
}

pub fn main() void {
    // const result = is_char('1');
    // print("{any}\n", .{result});
    // print("{s}\n", .{test_input});

    var char_count: i32 = 0;
    for (test_input, 0..) |char, idx| {
        // space
        if (char == 32) {
            continue;
        }

        // a-z
        if (is_char(char)) {
            char_count += 1;
            continue;
        }

        print("{d}: {c}\n", .{ idx, char });
    }

    print("{d}\n", .{char_count});
}
