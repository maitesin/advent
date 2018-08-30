use std::env;
use std::fs::File;
use std::io::prelude::*;


#[derive(Debug)]
enum Direction {
    North,
    East,
    South,
    West
}

impl Direction {
    fn left(&self) -> Direction {
        match self {
            Direction::North => Direction::West,
            Direction::West => Direction::South,
            Direction::South => Direction::East,
            Direction::East => Direction::North,
        }
    }
    fn right(&self) -> Direction {
        match self {
            Direction::North => Direction::East,
            Direction::East => Direction::South,
            Direction::South => Direction::West,
            Direction::West => Direction::North,
        }
    }
}

#[derive(Debug)]
enum Turn {
    Left,
    Right
}

struct Step {
    turn: Turn,
    distance: i64,
}

struct Location {
    direction: Direction,
    x: i64,
    y: i64,
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let mut f = File::open(&args[1]).expect("File not found");

    let mut content = String::new();
    f.read_to_string(&mut content).expect("Something went wrong reading the file");

    let steps = build_steps(content.trim().to_string());
    let location = find_building(steps);
    println!("Distance from original point: {}", location.x + location.y);
}

fn build_steps(content: String) -> Vec<Step> {
    let mut steps = Vec::new();
    let parts = content.split(", ");
    for part in parts {
        let turn = match &part[0..1] {
            "L" => Turn::Left,
            "R" => Turn::Right,
            _ => panic!(),
        };
        steps.push(Step{turn: turn, distance: part[1..].parse::<i64>().unwrap()});
    }
    return steps;
}

fn find_building(steps: Vec<Step>) -> Location {
    let mut location = Location{direction: Direction::North, x: 0, y: 0}; // Initial location
    for step in steps {
        match step.turn {
            Turn::Left => location.direction = location.direction.left(),
            Turn::Right => location.direction = location.direction.right(),
        }
        match location.direction {
            Direction::North => location.y += step.distance,
            Direction::East => location.x += step.distance,
            Direction::South => location.y -= step.distance,
            Direction::West => location.x -= step.distance,
        }
    }
    return location;
}
