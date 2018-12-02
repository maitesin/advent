use std::iter::Iterator;
use std::io::BufReader;
use std::io::BufRead;
use std::fs::File;

fn main() {
    let f = File::open("input.txt").unwrap();
    let file = BufReader::new(&f);
    let mut s_ids: Vec<String> = Vec::new();
    for line in file.lines() {
        let l = line.unwrap();
        s_ids.push(l);
    }

    let mut ids: Vec<ID> = Vec::new();
    for id1 in &s_ids {
        for id2 in &s_ids {
            ids.push(ID{s1: id1, s2: id2, diff: diff(id1, id2)});
        }
    }

    for cmp in ids {
        if cmp.diff == 1 {
            println!("{} {}", cmp.s1, cmp.s2);
        }
    }
}

struct ID<'a> {
    s1: &'a str,
    s2: &'a str,
    diff: i32,
}

fn diff(s1: &str, s2: &str) -> i32 {
    if s1.len() != s2.len() {
        return -1;
    }
    let v1: Vec<char> = s1.chars().collect();
    let v2: Vec<char> = s2.chars().collect();
    let mut count = 0;
    for pos in 0..s1.len() {
        if v1[pos] != v2[pos] {
            count += 1;
        }
    }
    return count;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_checksum() {
        assert_eq!(12, 12);
    }
}
