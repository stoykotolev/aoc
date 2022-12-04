use std::fs::File;
use std::io::{prelude::*, BufReader};

fn main() {
    let file = File::open("data.txt").expect("file not found");
    let buf = BufReader::new(file);
    let data: Vec<String> = buf.lines().map(|l| l.expect("some err")).collect();

    part1(&data);
    part2(&data);
}

fn part1(data: &Vec<String>) {
    let mut total: i32 = 0;

    for line in data {
        let elements: Vec<Vec<&str>> = line.split(",").map(|el| el.split("-").collect()).collect();

        let range1: Vec<i32> = elements[0]
            .iter()
            .map(|x| x.parse::<i32>().unwrap())
            .collect();
        let range2: Vec<i32> = elements[1]
            .iter()
            .map(|x| x.parse::<i32>().unwrap())
            .collect();

        if range1[0] <= range2[0] && range1[1] >= range2[1] {
            total += 1;
        } else if range2[0] <= range1[0] && range2[1] >= range1[1] {
            total += 1;
        }
    }

    println!("{}", total);
}

fn part2(data: &Vec<String>) {
    let mut total: i32 = 0;

    for line in data {
        let elements: Vec<Vec<&str>> = line.split(",").map(|el| el.split("-").collect()).collect();

        let item1: Vec<i32> = elements[0]
            .iter()
            .map(|x| x.parse::<i32>().unwrap())
            .collect();
        let item2: Vec<i32> = elements[1]
            .iter()
            .map(|x| x.parse::<i32>().unwrap())
            .collect();

        if item1[0] <= item2[1] && item1[1] >= item2[0] {
            total += 1;
        }
    }

    println!("{}", total);
}
