use reqwest::Client as HttpClient;
use reqwest::Response;
use std::collections::HashMap;
use std::time::{SystemTime, UNIX_EPOCH};
use hmac::{Hmac, Mac};
use sha2::Sha256;
use hex;
use serde_json;
use url::form_urlencoded;

pub enum Method {
    POST,
    GET,
}

pub type Params = HashMap<String, String>;

pub struct Request {
    method: Method,
    path: String,
    params: Params,
}

pub struct RustClient {
    key: String,
    secret_key: String,
    http_client: HttpClient,
    is_test_net: bool,
}

impl RustClient {
    pub fn new(key: &str, secret_key: &str, is_test_net: bool) -> RustClient {
        RustClient {
            key: key.to_string(),
            secret_key: secret_key.to_string(),
            http_client: HttpClient::new(),
            is_test_net,
        }
    }

    pub async fn get(&mut self, path: &str, params: Params) -> Result<Response, reqwest::Error> {
        self.do_request(Method::GET, path, params).await
    }

    pub async fn post(&mut self, path: &str, params: Params) -> Result<Response, reqwest::Error> {
        self.do_request(Method::POST, path, params).await
    }

    async fn do_request(&mut self, method: Method, path: &str, params: Params) -> Result<Response, reqwest::Error> {
        let base_url = if self.is_test_net {
            "TestnetBaseURL" // replace with the actual TestnetBaseURL
        } else {
            "BaseURL" // replace with the actual BaseURL
        };

        let full_url = format!("{}/{}", base_url, path);
        let mut http_request = match method {
            Method::GET => {
                let query = form_urlencoded::Serializer::new(String::new())
                    .extend_pairs(params.iter())
                    .finish();
                self.http_client.get(&format!("{}?{}", full_url, query))
            },
            Method::POST => {
                let body = serde_json::to_string(&params).unwrap();
                self.http_client.post(&full_url).body(body)
            },
        };

       http_request = self.set_common_headers(&mut http_request, &params);

        http_request.send().await
    }

    fn set_common_headers(&self, request: &mut reqwest::RequestBuilder, params: &Params) ->reqwest::RequestBuilder {
        let timestamp = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_millis().to_string();
        let signature = generate_signature(&self.secret_key, &self.key, &params, &timestamp);

      return  request.header("X-BAPI-SIGN-TYPE", "2")
            .header("X-BAPI-SIGN", signature)
            .header("X-BAPI-API-KEY", &self.key)
            .header("X-BAPI-TIMESTAMP", timestamp)
            .header("X-BAPI-RECV-WINDOW", "50000")
            .header("Content-Type", "application/json");
    }
}

fn generate_signature(secret_key: &str, api_key: &str, params: &Params, timestamp: &str) -> String {
    let mut base_string = timestamp.to_string() + api_key + "50000";

    let mut keys: Vec<_> = params.keys().collect();
    keys.sort();

    for key in keys {
        base_string.push_str(&format!("{}={}", key, params.get(key).unwrap()));
    }

    let mut mac = Hmac::<Sha256>::new_from_slice(secret_key.as_bytes()).unwrap();
    mac.update(base_string.as_bytes());

    hex::encode(mac.finalize().into_bytes())
}
