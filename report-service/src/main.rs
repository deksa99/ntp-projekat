#[macro_use] 
extern crate rocket;

mod handle;
mod model;

use rocket::{routes, Rocket, Build, Request};
use rocket::http::Status;
use rocket::serde::json::Json;

#[catch(404)]
fn not_found(req: &Request) -> String {
    format!("I couldn't find '{}'. Try something else?", req.uri())
}

#[catch(default)]
fn default(status: Status, req: &Request) -> String {
    format!("{} ({})", status, req.uri())
}

#[get("/manufacturer/<car_service_id>")]
async fn manufacturer(car_service_id: u32) -> Json<Vec<model::AppointmentInfo>> {
    Json(handle::man_report(car_service_id).await)
}

#[get("/service/<car_service_id>")]
fn service(car_service_id: u32) -> Json<String> {
    Json(format!("{{ \"hi\": \"manufacturer - {}\" }}", car_service_id).to_string())
}

#[get("/financial/<car_service_id>")]
fn financial(car_service_id: u32) -> Json<String> {
    Json(format!("{{ \"hi\": \"manufacturer - {}\" }}", car_service_id).to_string())
}

#[launch]
fn rocket() -> Rocket<Build> {
    rocket::build()
    .mount("/api/reports", routes![manufacturer, service, financial])
    .register("/", catchers![not_found, default])
}