const std = @import("std");
const ds = @import("data_structurez.zig");

const Dir = struct {
    up: bool = false,
    down: bool = false,
    left: bool = false,
    right: bool = false,
};

const input_test =
    \\.M.S......
    \\..A..MSMS.
    \\.M.S.MAA..
    \\..A.ASMSM.
    \\.M.S.M....
    \\..........
    \\S.S.S.S.S.
    \\.A.A.A.A..
    \\M.M.M.M.M.
    \\..........
;

pub fn findXmas(x: usize, y: usize, dir: Dir, matrix: ds.Matrix(u8)) usize {
    var found: usize = 0;

    if (dir.up and dir.down and dir.left and dir.right) {
        const up_left = matrix.get(x - 1, y - 1);
        const down_right = matrix.get(x + 1, y + 1);

        const up_right = matrix.get(x + 1, y - 1);
        const down_left = matrix.get(x - 1, y + 1);

        if (up_left + down_right == 'M' + 'S' and up_right + down_left == 'M' + 'S') found += 1;
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

        if (y >= 1) dir_ok.up = true;
        if (y + 1 <= matrix.height - 1) dir_ok.down = true;

        for (0..matrix.width) |x| {
            dir_ok.left = false;
            dir_ok.right = false;

            if (x >= 1) dir_ok.left = true;
            if (x + 1 <= matrix.width - 1) dir_ok.right = true;

            const loc = matrix.get(x, y);
            if (loc == 'A') {
                acc += findXmas(x, y, dir_ok, matrix);
            }
        }
    }

    std.debug.print("{d}\n", .{acc});
}
