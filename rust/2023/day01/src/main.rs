use std::fs;
use std::io::prelude::*;
use std::io::BufReader;

fn lines_from_file(filename: String) -> Vec<String> {
    let file = fs::File::open(filename).expect("no such file");
    let buf = BufReader::new(file);
    buf.lines()
        .map(|l| l.expect("could not parse line"))
        .collect()
}

fn main() {
    let filename = String::from("src/input.txt");
    let lines = lines_from_file(filename);
    let mut total = 0;
    let mut highest = 0;

    for line in lines {
        if line.is_empty() {
            if total > highest {
                highest = total;
            }
            total = 0;
            continue;
        }

        let num: i32 = line.parse().unwrap();
        total += num
    }

    println!("{highest}")
}
