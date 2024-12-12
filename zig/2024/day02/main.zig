const std = @import("std");
const print = std.debug.print;
const file = @embedFile("input.txt");

fn diffIsSafe(x: i32, y: i32) bool {
    const difference: u32 = @abs(x - y);

    if (difference >= 1 and difference <= 3) {
        return true;
    }

    return false;
}

fn isIncreasing(report: []i32) bool {
    for (0.., report) |idx, _| {
        if (idx == 0) {
            continue;
        }

        if (report[idx - 1] > report[idx]) {
            return false;
        }

        if (diffIsSafe(report[idx - 1], report[idx]) == false) {
            return false;
        }
    }

    return true;
}

fn isDecreasing(report: []i32) bool {
    for (0.., report) |idx, _| {
        if (idx == 0) {
            continue;
        }

        if (report[idx - 1] < report[idx]) {
            return false;
        }

        if (diffIsSafe(report[idx - 1], report[idx]) == false) {
            return false;
        }
    }

    return true;
}

fn dampen(report: []i32, allocator: std.mem.Allocator) !bool {
    var new_report = std.ArrayList(i32).init(allocator);

    for (0.., report) |idx, _| {
        for (0.., report) |i, val| {
            if (i == idx) {
                continue;
            }

            try new_report.append(val);
        }
        const report_slice = try new_report.toOwnedSlice();

        if (isIncreasing(report_slice)) {
            return true;
        }

        if (isDecreasing(report_slice)) {
            return true;
        }
    }

    return false;
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    // defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var report_list = std.ArrayList(i32).init(allocator);
    // defer report_list.deinit();

    var reports = std.mem.tokenize(u8, file, "\n");

    var safe_reports: i32 = 0;
    var safe_reports_dampened: i32 = 0;

    while (reports.next()) |r| {
        var report = std.mem.tokenize(u8, r, " ");
        while (report.next()) |num| {
            const num_as_int = try std.fmt.parseInt(i32, num, 10);
            try report_list.append(num_as_int);
        }

        // toOwnedSlice() clears capacity, this makes the calls to deinit()
        // above redundant.
        const report_slice = try report_list.toOwnedSlice();

        var is_safe: bool = isIncreasing(report_slice);
        if (is_safe == false) {
            is_safe = isDecreasing(report_slice);
        }

        if (is_safe) {
            safe_reports += 1;
        }

        if (is_safe == false) {
            is_safe = try dampen(report_slice, allocator);
            if (is_safe) {
                safe_reports_dampened += 1;
            }
        }
    }

    print("Part 1: {d}\n", .{safe_reports});
    print("Part 2: {d}\n", .{safe_reports + safe_reports_dampened});
}
