use std::fs;

fn main() {
    let input = fs::read_to_string("data.txt").expect("file");

    println!("{}", checker(&input, 4));
    println!("{}", checker(&input, 14));
}

fn checker(input: &String, marker_length: usize) -> usize {
    let mut markers: String = String::from("");

    // Requirements
    // 1. Check prev 4 chars
    // 2. If current char eq one of previous do:
    //  2.1 If first char is equal to curr, drop first
    //  2.2 If last char is eq to curr, drop string and start from curr
    //  2.3 If one of middle chars is eq to curr, drop until the char and continue

    for (idx, char) in input.chars().enumerate() {
        if markers.len() == marker_length {
            println!("{}", markers);
            return idx;
        }
        if markers.contains(char) {
            if markers.chars().next().unwrap() == char {
                markers = markers[1..].to_string();
                markers.push(char);
            } else if markers.chars().last().unwrap() == char {
                markers.clear();
                markers.push(char);
            } else {
                let remove_from = markers.chars().position(|c| c == char).unwrap() + 1;
                markers = markers[remove_from..].to_string();
                markers.push(char);
            }
        } else {
            markers.push(char);
        }
    }

    return 0usize;
}
