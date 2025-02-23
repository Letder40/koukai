use std::{env, fmt::Write as WriteFmt, fs::File, process::exit};
use std::io::Write;
use rand::{RngCore, rng};
use sha2::Sha512;
use hmac::{Hmac, Mac};

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() != 3 {
        println!("Usage: {} [ SALT ] [ OUTPUT-FILE ]", &args[0]);
        exit(1)
    }

    let salt = &args[1];
    let output_file_path = &args[2];

    let mut rng = rng();
    let mut access_key: [u8; 128] = [0; 128];
    rng.fill_bytes(&mut access_key);
    
    let mut access_key_hex_lower = String::new();
    for byte in access_key {
        write!(&mut access_key_hex_lower, "{:x}", byte).expect("Error writing hex to string buf");
    }

    let mut hasher = Hmac::<Sha512>::new_from_slice(salt.as_bytes()).unwrap();
    hasher.update(access_key_hex_lower.as_bytes()); 
    let result = hasher.finalize().into_bytes();

    let mut access_key_hashed = String::new();
    for byte in result {
        write!(&mut access_key_hashed, "{:x}", byte).expect("Error writing hex to string buf");
    }

    let mut output_file = File::create(output_file_path).expect("Error opening the provided output file");  
    output_file.write_all(access_key_hex_lower.as_bytes()).unwrap();
    print!("{access_key_hashed}");
}
