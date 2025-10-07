use curl::easy::Easy;
use log::debug;

pub fn send_request(url: &str) -> Result<Vec<u8>, curl::Error> {
    let mut handle = Easy::new();
    handle.url(url)?;
    debug!("Sending request to URL: {}", url);
    let mut response_data = Vec::new();
    {
        let mut transfer = handle.transfer();
        transfer.write_function(|data| {
            response_data.extend_from_slice(data);
            Ok(data.len())
        })?;
        transfer.perform()?;
    }
    Ok(response_data.clone())
}
