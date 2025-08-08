enum Currency {
    Rupee,
    Taka,
    Yen,
    Yuan,
}

impl Currency {
    fn currency_rate(&self) -> u32 {
        match self {
            Currency::Rupee => 87,
            Currency::Taka => 220,
            Currency::Yen => 1000,
            Currency::Yuan => 7,
        }
    }
    
    fn as_str(&self) -> &str {
        match self {
            Currency::Rupee => "Rupee",
            Currency::Taka => "Taka",
            Currency::Yen => "Yen",
            Currency::Yuan => "Yuan",
        }
    }
}

fn main() {
    let current_check = Currency::Yen;
    println!(
        "Dollar to Currency conversion for {} is {}",
        current_check.as_str(),
        current_check.currency_rate()
    );
}
