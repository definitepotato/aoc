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

pub fn main() {
    let lines = lines_from_file("src/input.txt".to_string());
    let mut sum: u32 = 0;

    for line in lines {
        let parts: Vec<u32> = line.split("x").map(|n| n.parse::<u32>().unwrap()).collect();
        let (l, w, h) = (parts[0], parts[1], parts[2]);

        let areas = [l * w, w * h, h * l];
        let smallest = areas.iter().min().unwrap();

        sum += 2 * ((l * w) + (w * h) + (h * l)) + smallest;
    }

    println!("Part 1: {:?}", sum);
}
