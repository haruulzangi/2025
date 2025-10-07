use actix_files as fs;
use actix_web::http::header::{ContentDisposition, DispositionType};
use actix_web::{Error, get};

#[get("/binary")]
pub async fn binary_route() -> Result<fs::NamedFile, Error> {
    let binary_path =
        std::env::var("BINARY_PATH").expect("BINARY_PATH environment variable not set");
    let file = fs::NamedFile::open(&binary_path)?;
    Ok(file
        .use_last_modified(true)
        .set_content_disposition(ContentDisposition {
            disposition: DispositionType::Attachment,
            parameters: vec![],
        }))
}
