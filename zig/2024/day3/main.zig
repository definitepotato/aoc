const std = @import("std");
const print = std.debug.print;
const file = @embedFile("test.txt");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var idx_list = std.ArrayList(usize).init(allocator);
    defer idx_list.deinit();

    var instructions = std.mem.tokenize(u8, file, "\n");

    // collect the index of the first number of each valid 'mul(' sequence
    while (instructions.next()) |instr| {
        for (0.., instr) |idx, _| {
            // we don't want to scan past the len of the instruction
            if (idx + 3 > instr.len) {
                continue;
            }

            // starting from any 'm' char, look for 'mul(' pattern and store the index of the first digit
            if (instr[idx] == 'm') {
                if (instr[idx + 1] == 'u' and instr[idx + 2] == 'l' and instr[idx + 3] == '(') {
                    try idx_list.append(idx + 4);

                    // since we're here, capture the first number up to ',' then the second number up to ')', and convert + math
                    // maybe we don't need to allocate for this, just process in main
                }
            }
        }
    }
    // print("{any}\n", .{idx_list.items});
}
