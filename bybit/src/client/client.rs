use std::collections::HashMap;
use reqwest::Client as HttpClient;

pub enum Method{
    POST,
    GET,
}

pub type Params = HashMap<String,String>;

pub struct Request{
    method:Method,
    path:String,
    params:Params
}

pub struct Client{
    key :String,
    secret_key : String,
    http_client :HttpClient,
    is_test_net:bool,
}

impl Client {
    pub fn new(key:&str,secret_key:&str,is_test_net:bool)->Client{
        Client{
            key:key.to_string(),
            secret_key: secret_key.to_string(),
            http_client: HttpClient::new(),
            is_test_net
        }
    }

}