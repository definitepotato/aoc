const std = @import("std");
const ds = @import("data_structurez.zig");

const Dir = struct {
    up: bool = false,
    down: bool = false,
    left: bool = false,
    right: bool = false,
};

const input_test =
    \\MMMSXXMASM
    \\MSAMXMSMSA
    \\AMXSXMAAMM
    \\MSAMASMSMX
    \\XMASAMXAMM
    \\XXAMMXXAMA
    \\SMSMSASXSS
    \\SAXAMASAAA
    \\MAMMMXMMMM
    \\MXMXAXMASX
;

pub fn findXmas(x: usize, y: usize, dir: Dir, matrix: ds.Matrix(u8)) usize {
    var found: usize = 0;

    if (dir.up) {
        if (matrix.get(x, y - 1) == 'M' and matrix.get(x, y - 2) == 'A' and matrix.get(x, y - 3) == 'S') found += 1;
    }

    if (dir.down) {
        if (matrix.get(x, y + 1) == 'M' and matrix.get(x, y + 2) == 'A' and matrix.get(x, y + 3) == 'S') found += 1;
    }

    if (dir.left) {
        if (matrix.get(x - 1, y) == 'M' and matrix.get(x - 2, y) == 'A' and matrix.get(x - 3, y) == 'S') found += 1;
    }

    if (dir.right) {
        if (matrix.get(x + 1, y) == 'M' and matrix.get(x + 2, y) == 'A' and matrix.get(x + 3, y) == 'S') found += 1;
    }

    if (dir.up and dir.right) {
        if (matrix.get(x + 1, y - 1) == 'M' and matrix.get(x + 2, y - 2) == 'A' and matrix.get(x + 3, y - 3) == 'S') found += 1;
    }

    if (dir.up and dir.left) {
        if (matrix.get(x - 1, y - 1) == 'M' and matrix.get(x - 2, y - 2) == 'A' and matrix.get(x - 3, y - 3) == 'S') found += 1;
    }

    if (dir.down and dir.right) {
        if (matrix.get(x + 1, y + 1) == 'M' and matrix.get(x + 2, y + 2) == 'A' and matrix.get(x + 3, y + 3) == 'S') found += 1;
    }

    if (dir.down and dir.left) {
        if (matrix.get(x - 1, y + 1) == 'M' and matrix.get(x - 2, y + 2) == 'A' and matrix.get(x - 3, y + 3) == 'S') found += 1;
    }

    return found;
}

pub fn main() !void {
    const allocator = std.heap.page_allocator;

    // var matrix = try ds.Matrix(u8).initFromText(allocator, input_test);
    var matrix = try ds.Matrix(u8).initFromFile(allocator, "input.txt");
    defer matrix.deinit();

    var acc: usize = 0;
    var dir_ok = Dir{};
    for (0..matrix.height) |y| {
        dir_ok.up = false;
        dir_ok.down = false;

        if (y >= 3) dir_ok.up = true;
        if (y + 3 <= matrix.height - 1) dir_ok.down = true;

        for (0..matrix.width) |x| {
            dir_ok.left = false;
            dir_ok.right = false;

            if (x >= 3) dir_ok.left = true;
            if (x + 3 <= matrix.width - 1) dir_ok.right = true;

            const loc = matrix.get(x, y);
            if (loc == 'X') {
                acc += findXmas(x, y, dir_ok, matrix);
            }
        }
    }

    std.debug.print("{d}\n", .{acc});
}
