use std::fs::File;
use std::io::{prelude::*, BufReader};

macro_rules! hashmap {
    ($( $key: expr => $val: expr ),*) => {{
         let mut map = ::std::collections::HashMap::new();
         $( map.insert($key, $val); )*
         map
    }}
}

fn main() {
    println!("{}", day2())
}

fn day2() -> i32 {
    let scores = hashmap!['A' => 1, 'B' => 2, 'C' => 3];
    let results = hashmap!["win" => 6, "loss" => 0, "draw" => 3 ];

    let mut total: i32 = 0;
    let file = File::open("data.txt").expect("file not found");
    let buf = BufReader::new(file);
    let data: Vec<String> = buf.lines().map(|l| l.expect("some err")).collect();

    // A = Rock
    // B = Paper
    // C = Scissors
    // X = lose
    // Y = draw
    // Z = win

    for line in data {
        if line.chars().nth(0).unwrap() == 'A' {
            if line.chars().nth(2).unwrap() == 'Y' {
                total += scores.get(&'A').cloned().unwrap() + results.get("draw").cloned().unwrap();
            } else if line.chars().nth(2).unwrap() == 'X' {
                total += scores.get(&'C').cloned().unwrap() + results.get("loss").cloned().unwrap();
            } else {
                total += scores.get(&'B').cloned().unwrap() + results.get("win").cloned().unwrap();
            }
        } else if line.chars().nth(0).unwrap() == 'B' {
            if line.chars().nth(2).unwrap() == 'Y' {
                total += scores.get(&'B').cloned().unwrap() + results.get("draw").cloned().unwrap();
            } else if line.chars().nth(2).unwrap() == 'X' {
                total += scores.get(&'A').cloned().unwrap() + results.get("loss").cloned().unwrap();
            } else {
                total += scores.get(&'C').cloned().unwrap() + results.get("win").cloned().unwrap();
            }
        } else {
            if line.chars().nth(2).unwrap() == 'Y' {
                total += scores.get(&'C').cloned().unwrap() + results.get("draw").cloned().unwrap();
            } else if line.chars().nth(2).unwrap() == 'X' {
                total += scores.get(&'B').cloned().unwrap() + results.get("loss").cloned().unwrap();
            } else {
                total += scores.get(&'A').cloned().unwrap() + results.get("win").cloned().unwrap();
            }
        }
    }

    return total;
}
