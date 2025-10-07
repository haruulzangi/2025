use actix_web::{
    App, HttpRequest, HttpResponse, HttpServer, Responder, get, http::header, middleware::Logger,
};
use crc16::{ARC, State};

mod magic;

#[get("/flag")]
async fn get_flag(req: HttpRequest) -> impl Responder {
    let headers = req.headers();
    let client_ip = headers.get("x-client-ip"); // I think Traefik is smart enough to handle this :)
    if client_ip.is_none() {
        return HttpResponse::Forbidden().body("Access denied");
    }
    let hash = State::<ARC>::calculate(client_ip.unwrap().as_bytes());
    if hash != 0xBEEF {
        return HttpResponse::Forbidden().body("Access denied");
    }

    let flag = magic::generate_flag();
    return HttpResponse::Ok().body(flag);
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    env_logger::init_from_env(env_logger::Env::new().default_filter_or("info"));
    let port = std::env::var("PORT").unwrap_or_else(|_| "8080".to_string());
    HttpServer::new(|| App::new().wrap(Logger::default()).service(get_flag))
        .bind(("0.0.0.0", port.parse::<u16>().unwrap()))?
        .run()
        .await
}
