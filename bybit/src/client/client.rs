use std::collections::HashMap;
use std::fmt::format;
use reqwest::{Client as HttpClient, Response};
use serde_json::Value::String;
use url::form_urlencoded;

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

    pub async fn get(&self,path:&str,params:Params)->Result<Response,reqwest::Error>{
       self.do_request(Method::GET,path,params)
    }
    pub async fn post(&self,path:&str,params:Params)->Result<Response,reqwest::Error>{
        self.do_request(Method::GET,path,params)
    }

    async fn do_request(&self,method:Method,path :&str,params:Params)->Result<Response,reqwest::Error>{
        let base_url = if self.is_test_net{
            "https://api-testnet.bybit.com"
        }else {
            "https://api.bybit.com"
        };
        let full_url = format!("{}/{}",base_url,path);
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

        self.set_common_headers(&mut http_request, &params);

        http_request.send().await
    }


}