use std::fs;

fn main() {
    let input = fs::read_to_string("data.txt").expect("file");

    let (model, instructions) = input.split_once("\n\n").unwrap();
    let (rows, platform) = model.rsplit_once("\n").unwrap();

    let platform_size: usize = platform.split_whitespace().last().unwrap().parse().unwrap();

    let mut stacks = vec![Vec::new(); platform_size];

    for line in rows.lines().rev() {
        for (idx, chunk) in line
            .chars()
            .collect::<Vec<_>>()
            .chunks(4)
            .into_iter()
            .enumerate()
        {
            let second = chunk.get(1).unwrap();
            if second.is_alphabetic() {
                let letter = second.to_string();
                stacks[idx].push(letter);
            }
        }
    }

    // part1(&mut stacks, &instructions);
    part2(&mut stacks, &instructions);
}
fn part1(stacks: &mut Vec<Vec<String>>, instructions: &str) {
    for line in instructions.lines() {
        let rest = line.strip_prefix("move ").unwrap();
        let (amount, rest) = rest.split_once(" from ").unwrap();
        let (from, to) = rest.split_once(" to ").unwrap();
        let parsed_amount: i32 = amount.parse().unwrap();
        let parsed_from: usize = from.parse().unwrap();
        let parsed_to: usize = to.parse().unwrap();

        for _ in 0..parsed_amount {
            if let Some(removed) = stacks[parsed_from - 1].pop() {
                stacks[parsed_to - 1].push(removed);
            }
        }
    }

    let result = stacks
        .into_iter()
        .filter_map(|el| el.into_iter().last())
        .collect::<Vec<_>>();
    println!("{:?}", result);
}

fn part2(stacks: &mut Vec<Vec<String>>, instructions: &str) {
    for line in instructions.lines() {
        let rest = line.strip_prefix("move ").unwrap();
        let (amount, rest) = rest.split_once(" from ").unwrap();
        let (from, to) = rest.split_once(" to ").unwrap();
        let parsed_amount: usize = amount.parse().unwrap();
        let parsed_from: usize = from.parse().unwrap();
        let parsed_to: usize = to.parse().unwrap();

        if parsed_amount == 1 {
            if let Some(removed) = stacks[parsed_from - 1].pop() {
                stacks[parsed_to - 1].push(removed);
            }
        } else {
            let stack_length = stacks[parsed_from - 1].len();
            if stack_length >= parsed_amount {
                let moved_crates = stacks[parsed_from - 1].split_off(stack_length - parsed_amount);

                stacks[parsed_to - 1].extend(moved_crates);
            }
        }
    }

    let result = stacks
        .into_iter()
        .filter_map(|el| el.into_iter().last())
        .collect::<Vec<_>>();
    println!("{:?}", result);
}
