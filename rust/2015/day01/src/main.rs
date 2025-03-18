use std::fs;
use std::io::BufReader;
use std::io::prelude::*;

fn lines_from_file(filename: String) -> Vec<String> {
    let file = fs::File::open(filename).expect("no such file");
    let buf = BufReader::new(file);
    buf.lines()
        .map(|l| l.expect("could not parse line"))
        .collect()
}

fn main() {
    // let input = r#")())())"#;
    let lines = lines_from_file(String::from("src/input.txt"));

    let mut index: u32 = 0;
    let mut floor: i32 = 0;
    let mut basement_floor: u32 = 0;

    for line in lines {
        for c in line.chars() {
            if basement_floor == 0 {
                index += 1;
            }
            match c {
                ')' => {
                    floor -= 1;
                    if floor < 0 {
                        basement_floor = index;
                    }
                }
                '(' => {
                    floor += 1;
                    if floor < 0 {
                        basement_floor = index;
                    }
                }
                _ => {
                    floor += 0;
                    if floor < 0 {
                        basement_floor = index;
                    }
                }
            }
        }
    }

    println!("Part 1: {}", floor);
    println!("Part 2: {}", basement_floor);
}
