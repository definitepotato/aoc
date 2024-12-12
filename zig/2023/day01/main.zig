const std = @import("std");
const print = std.debug.print;
const file = @embedFile("input.txt");

const test_input1 =
    \\ 1abc2
    \\ pqr3stu8vwx
    \\ a1b2c3d4e5f
    \\ treb7uchet
;

const test_input2 =
    \\ two1nine
    \\ eightwothree
    \\ abcone2threexyz
    \\ xtwone3four
    \\ 4nineeightseven2
    \\ zoneight234
    \\ 7pqrstsixteen
;

fn firstAndLastNumber(line: []const u8) ![2]i32 {
    var first: i32 = undefined;
    var second: i32 = undefined;
    var firstFound: bool = false;

    for (line) |c| {
        if (c > 47 and c < 58) {
            if (firstFound) {
                second = try std.fmt.parseInt(i32, &[_]u8{c}, 10);
                continue;
            }

            first = try std.fmt.parseInt(i32, &[_]u8{c}, 10);
            firstFound = true;
        }
    }

    if (second < 0) {
        second = first;
    }

    return [2]i32{ first, second };
}

fn firstNumber(line: []const u8) !usize {
    const word_numbers = [_][]const u8{
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",
    };

    for (word_numbers, 0..) |word, word_i| {
        if (std.mem.startsWith(u8, line, word)) {
            // print("[DEBUG] {s}: {s}, {d}\n", .{ line, word_numbers[word_i], word_i + 1 });
            return word_i + 1;
        }
    }

    for (line) |c| {
        if (c > 47 and c < 58) {
            return try std.fmt.parseInt(usize, &[_]u8{c}, 10);
        }
    }

    return 0;
}

fn secondNumber(line: []const u8) !usize {
    const word_numbers = [_][]const u8{
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",
    };

    for (word_numbers, 0..) |word, word_i| {
        if (std.mem.endsWith(u8, line, word)) {
            // print("[DEBUG] {s}: {s}, {d}\n", .{ line, word_numbers[word_i], word_i + 1 });
            return word_i + 1;
        }
    }

    var last_number: usize = undefined;
    for (line) |c| {
        if (c > 47 and c < 58) {
            last_number = try std.fmt.parseInt(usize, &[_]u8{c}, 10);
        }
    }

    return last_number;
}

fn sumTheArray(arr: [2]i32) i32 {
    const left = arr[0] * 10;
    return left + arr[1];
}

pub fn main() !void {
    var total: i32 = 0;
    var lines = std.mem.split(u8, file, "\n");

    while (lines.next()) |line| {
        if (line.len == 0) {
            continue;
        }

        const result = try firstAndLastNumber(line);
        const sum = sumTheArray(result);
        total += sum;

        const resultWithWordsFirst = try firstNumber(line);
        const resultWithWordsSecond = try secondNumber(line);
        print("[DEBUG] {s}: => {d}, {d}\n", .{ line, resultWithWordsFirst, resultWithWordsSecond });

        // print("[DEBUG] {s}: {d},{d} => {d}\n", .{ line, result[0], result[1], sum });
    }

    print("Part 1: {d}\n", .{total});
}
