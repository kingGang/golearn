use 
struct Point{
    x: i32,
    y: i32
}

fn main() {
    // simple_ownership();
    mutable();
}

//所有权
fn simple_ownership() {
    let  p1 = Point{x:25,y:25};
    let p2 = p1;
    println!("{} {}",p2.x,p2.y);
}

fn mutable(){
    let mut p1 =Point{x:25,y:25};
    p1.x=30;
    println!("{} {}",p1.x,p1.y);
}

fn new_thread(){
    let p =Rc:new(Point(x:25,y:25))
}