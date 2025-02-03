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

pub fn main() {
    let lines = lines_from_file(String::from("src/input.txt"));
    // let mut left: Vec<String> = vec![];
    // let mut right: Vec<String> = vec![];
    let mut left = Vec::new();
    let mut right = Vec::new();
    let mut sum = Vec::new();

    let mut count = 0;
    for line in lines {
        for r in line.split("   ") {
            if count % 2 == 0 {
                right.push(r.to_string());
            } else {
                left.push(r.to_string());
            }

            count += 1;
        }
    }

    // Vec<String> to Vec<i32>.
    let mut int_left: Vec<i32> = left.iter().filter_map(|s| s.parse().ok()).collect();
    let mut int_right: Vec<i32> = right.iter().filter_map(|s| s.parse().ok()).collect();

    // Sort the vectors.
    int_left.sort();
    int_right.sort();

    // Subtract each pair in ascending order and push to new Vec<i32>.
    for (idx, int) in int_left.clone().into_iter().enumerate() {
        sum.push(int - int_right[idx])
    }

    let mut ans: i32 = 0;

    // borrow `sum`
    for value in &sum {
        ans += value.abs();
    }

    // println!("{:?} {:?}\n{:?}\n{:?}", int_left, int_right, sum, ans);
    println!("Part 1: {}", ans);
}
