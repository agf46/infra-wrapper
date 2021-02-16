// Program to grep for dockerd logs 
use std::process::{Command, Stdio}; 
use std::io::{BufRead, BufReader, Error, ErrorKind};

fn main() -> Result<(), Error> {
    let stdout = Command::new("journalctl")
        .stdout(Stdio::piped())
        .spawn()?
        .stdout
        .ok_or_else(|| Error::new(ErrorKind::Other, "Could not capture stdout."))?;
    
    // Read the journal logs in a buffered stream and then filter out for dockerd
    let reader = BufReader::new(stdout);
    reader
        .lines()
        .filter_map(|line| line.ok())
        .filter(|line| line.find("dockerd").is_some())
        .for_each(|line| println!("{}", line));

    Ok(())
}