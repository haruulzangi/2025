use url::Url;
use serde::Deserialize;
use actix_web::{post, web, HttpResponse};

use crate::request::dispatcher::send_request;

#[derive(Deserialize)]
struct FormData {
    url: String,
}

#[post("/request")]
pub async fn request_handler(form: web::Form<FormData>) -> HttpResponse {
    if Url::parse(&form.url).is_err() {
        return HttpResponse::BadRequest().body("Invalid URL");
    }
    match send_request(&form.url) {
        Ok(response) => HttpResponse::Ok().body(response),
        Err(_) => HttpResponse::InternalServerError().body("Failed to send request"),
    }
}
