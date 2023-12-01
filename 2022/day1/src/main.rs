use std::fs::File;
use std::io::{prelude::*, BufReader};
fn main() {
    println!("{}", get_files())
}

fn get_files() -> i32 {
    let mut curr_cals: i32 = 0;
    let mut top_elves: Vec<i32> = Vec::new();
    let file = File::open("data.txt").expect("no such file");
    let buf = BufReader::new(file);
    let data: Vec<String> = buf
        .lines()
        .map(|l| l.expect("Could not parse line"))
        .collect();

    for line in data {
        if line.is_empty() {
            if top_elves.len() < 3 {
                top_elves.push(curr_cals);
            }

            for i in 0..top_elves.len() {
                if curr_cals > top_elves[i] && !top_elves.contains(&curr_cals) {
                    top_elves.remove(top_elves.len() - 1);
                    top_elves.push(curr_cals);
                }
            }
            curr_cals = 0;
            top_elves.sort_by(|a, b| b.cmp(a));
        } else {
            curr_cals += line.parse::<i32>().unwrap();
        }
    }
    return top_elves.iter().sum();
}
