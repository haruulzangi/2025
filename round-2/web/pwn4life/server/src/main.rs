use actix_web::{App, HttpServer};
use log::error;

mod request;
mod routes;

use routes::binary::binary_route;
use routes::index::home_page;
use routes::request::request_handler;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    env_logger::init();
    if std::env::var("BINARY_PATH").is_err() {
        error!("BINARY_PATH environment variable is not set.");
        return Ok(());
    }
    HttpServer::new(|| {
        App::new()
            .service(home_page)
            .service(request_handler)
            .service(binary_route)
    })
    .bind(("0.0.0.0", 8080))?
    .run()
    .await
}
