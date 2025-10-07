use actix_web::{get, HttpResponse, Responder};

#[get("/")]
pub async fn home_page() -> impl Responder {
    HttpResponse::Ok().body(include_str!("index.html"))
}
