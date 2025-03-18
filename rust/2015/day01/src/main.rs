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

    let mut floor: i32 = 0;

    for line in lines {
        for c in line.chars() {
            match c {
                ')' => {
                    floor -= 1;
                }
                '(' => {
                    floor += 1;
                }
                _ => {
                    floor += 0;
                }
            }
        }
    }

    println!("Part 1: {}", floor);
}
