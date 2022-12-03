use std::collections::HashMap;
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
    let alphabet = hashmap![ 'a' => 1, 'b' => 2, 'c' => 3, 'd' => 4, 'e' => 5, 'f' => 6, 'g' => 7, 'h' => 8, 'i' => 9, 'j' => 10, 'k' => 11, 'l' => 12, 'm' => 13, 'n' => 14, 'o' => 15,'p' => 16, 'q' => 17, 'r' => 18, 's' => 19, 't' => 20, 'u' => 21, 'v' => 22, 'w' => 23, 'x' => 24, 'y' => 25, 'z' => 26];
    let file = File::open("data.txt").expect("file not found");
    let buf = BufReader::new(file);
    let data: Vec<String> = buf.lines().map(|l| l.expect("some err")).collect();

    println!("part1: {}", day3part1(&alphabet, &data));
    println!("part2: {}", day3part2(&alphabet, &data));
}

fn day3part1(alphabet: &HashMap<char, i32>, data: &Vec<String>) -> i32 {
    let mut total: i32 = 0;

    for line in data {
        let mid = line.len() / 2;
        let first_half = &line[..mid];
        let second_half = &line[mid..];
        for char in first_half.chars() {
            if second_half.contains(char) {
                let amount = alphabet.get(&char.to_ascii_lowercase()).unwrap();
                if char.is_lowercase() {
                    total += amount;
                } else {
                    total += amount + 26;
                }
                break;
            }
        }
    }
    return total;
}

fn day3part2(alphabet: &HashMap<char, i32>, data: &Vec<String>) -> i32 {
    let mut total: i32 = 0;

    for i in (2..data.len()).step_by(3) {
        for char in data[i].chars() {
            if data[i - 1].contains(char) && data[i - 2].contains(char) {
                let amount = alphabet.get(&char.to_ascii_lowercase()).unwrap();
                if char.is_lowercase() {
                    total += amount;
                } else {
                    total += amount + 26;
                }
                break;
            }
        }
    }
    return total;
}
