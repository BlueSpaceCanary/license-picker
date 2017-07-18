use std::io;

trait tree {
    fn add_child(&mut self, fork) -> Self;
    fn get_match(&self) -> Option<fork>;
}

struct fork {
    node_text: &'static str,
    choice_text: &'static str,
    children: Vec<fork>
}

impl fork {
    fn new(node_text: &'static str, choice_text: &'static str) -> fork {
        fork {
            node_text: node_text,
            choice_text: choice_text,
            children: std::vec::Vec::new()
        }
    }
    fn add_child(self: &mut Self, child: fork) {
        self.children.push(child);
    }
    fn get_match(self: &Self, s: String) -> Option<&fork> {
        for n in self.children.iter() {
            if format!("{}\n", n.choice_text) == s {
                return Some(n);
            }
        }
        return None;
    }
}

fn walk(f: &fork) {
    println!("{}? (or \"help\" to list options): ", f.node_text);
    let mut txt = String::new();
    io::stdin().read_line(&mut txt).unwrap();

    match f.get_match(txt) {
        Some(m) =>
            if m.children.len() > 0 {
                walk(m);
            } else {
                println!("****{}****", m.node_text);
            },
        None => {
            println!("Options: ");
            for n in f.children.iter() {
                println!("{}", n.choice_text);
            }
            walk(f);
        }
    }
}

fn lolsetup() -> fork {
    let mut co_bs = fork::new("Can companies say you like their product?", "NOPE");
    co_bs.add_child(fork::new("BSD 3 clause", "What OMG no!"));
    co_bs.add_child(fork::new("MIT/X11/zlib", "Why are you still asking me questions?"));

    let mut v2 = fork::new("Bernie or Lenin", "YES");
    v2.add_child(fork::new("LGPLv2", "Bernie"));
    v2.add_child(fork::new("GPLv2", "Lenin"));

    // same
    let mut copyleft_ip_anarchy = fork::new("Copyleft", "LOL");
    copyleft_ip_anarchy.add_child(co_bs);
    copyleft_ip_anarchy.add_child(v2);

    let mut v3 = fork::new("Bernie or Lenin", "YES");
    v3.add_child(fork::new("LGPLv3", "Bernie"));
    v3.add_child(fork::new("GPLv3", "Lenin"));
    v3.add_child(fork::new("aGPLv3", "Mao"));

    let apache = fork::new("Apache 2.0", "NOPE");

    let mut copyleft_ip_tyrrany = fork::new("Copyleft", "oh shit");
    copyleft_ip_tyrrany.add_child(apache);
    copyleft_ip_tyrrany.add_child(v3);

    let mut patents = fork::new("patents?", "sure");
    patents.add_child(copyleft_ip_anarchy);
    patents.add_child(copyleft_ip_tyrrany);

    let pd = fork::new("Public domain", "HAHA NO");

    let mut us_is_god = fork::new("About non-US folks?", "sure");
    us_is_god.add_child(pd);
    us_is_god.add_child(patents);

    let the_good_ones = fork::new("WTFPL/CC0", "NO");

    let mut meirl = fork::new("Do you care?", "Let's go");

    meirl.add_child(the_good_ones);
    meirl.add_child(us_is_god);

    let mut root = fork::new("How to pick an open source license", "");
    root.add_child(meirl);
    root
}

fn main() {
    walk(&lolsetup());
}
