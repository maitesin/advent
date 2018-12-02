use std::iter::Iterator;
use std::iter::FromIterator;
use std::io::BufReader;
use std::io::BufRead;
use std::fs::File;

fn main() {
    let f = File::open("input.txt").unwrap();
    let mut file = BufReader::new(&f);
    let mut contents : Vec<Content> = Vec::new();
    for line in file.lines() {
        let l = line.unwrap();
        let c = content(&l);
        contents.push(c);
    }
    println!("{}", checksum(contents));
}

#[derive(Debug)]
struct Content {
    two: bool,
    three: bool,
}

fn checksum(contents: Vec<Content>) -> i32 {
    let mut three = 0;
    let mut two = 0;
    for content in contents {
        if content.two {
            two += 1;
        }
        if content.three {
            three += 1;
        }
    }
    return three * two;
}

fn sort(input: &str) -> String {
    let mut chars: Vec<char> = input.chars().collect();
    chars.sort();
    return String::from_iter(chars);
}

fn content(input: &str) -> Content {
    let chars: Vec<char> = sort(input).chars().collect();
    let mut two: bool = false;
    let mut three: bool = false;
    let mut pos: usize = 0;
    let max = chars.len();
    loop {
        if pos + 2 <  max {
            if chars[pos] == chars[pos+1] && chars[pos] == chars[pos+2] {
                three = true;
                pos += 2;
            } else {
                if chars[pos] == chars[pos+1] {
                    two = true;
                    pos += 1;
                }
            }
        } else {
            if pos + 1 < max {
                if chars[pos] == chars[pos+1] {
                    two = true;
                    pos += 1;
                }
            } else {
                break;
            }
        }
        pos += 1;
    }
    return Content{two: two, three: three};
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_checksum() {
        let mut input: Vec<Content> = Vec::new();
        input.push(Content{three: false, two: false});
        input.push(Content{three: true, two: true});
        input.push(Content{three: false, two: true});
        input.push(Content{three: true, two: false});
        input.push(Content{three: false, two: true});
        input.push(Content{three: false, two: true});
        input.push(Content{three: true, two: false});
        assert_eq!(checksum(input), 12);
    }

    #[test]
    fn test_string_sort() {
        assert_eq!(sort(""), String::from(""));
        assert_eq!(sort("a"), String::from("a"));
        assert_eq!(sort("ba"), String::from("ab"));
        assert_eq!(sort("bababa"), String::from("aaabbb"));
        assert_eq!(sort("cbaabc"), String::from("aabbcc"));
    }

    #[test]
    fn test_string_content() {
        let mut output = content("abcdef");
        assert_eq!(output.two, false);
        assert_eq!(output.three, false);
        output = content("bababc");
        println!("{:?}", output);
        assert_eq!(output.two, true);
        assert_eq!(output.three, true);
        output = content("abbcde");
        assert_eq!(output.two, true);
        assert_eq!(output.three, false);
        output = content("abcccd");
        assert_eq!(output.two, false);
        assert_eq!(output.three, true);
        output = content("aabcdd");
        assert_eq!(output.two, true);
        assert_eq!(output.three, false);
        output = content("abcdee");
        assert_eq!(output.two, true);
        assert_eq!(output.three, false);
        output = content("ababab");
        assert_eq!(output.two, false);
        assert_eq!(output.three, true);
    }
}
