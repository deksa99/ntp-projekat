use crate::model;

pub async fn man_report(car_service_id: u32) -> Vec<model::AppointmentInfo> {
    get_appointments(car_service_id).await
}

pub async fn get_appointments(car_service_id: u32) -> Vec<model::AppointmentInfo> {
    match get_appointments_for_service_request(car_service_id).await {
        Ok(r) => r,
        Err(e) => {
            println!("{}", e.to_string());
            Vec::new()
        },
    }
}


pub async fn get_appointments_for_service_request(car_service_id: u32) -> Result<Vec<model::AppointmentInfo>, reqwest::Error> {
    let url = format!("http://localhost:8084/api/appointments/requests/service/{}", car_service_id);
    reqwest::get(url).await?.json::<Vec<model::AppointmentInfo>>().await
}