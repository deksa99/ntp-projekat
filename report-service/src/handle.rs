use std::collections::HashMap;

use crate::model;

pub async fn man_report(car_service_id: u32) -> Vec<model::ManufacturerReport> {
    match get_appointments_for_service_request(car_service_id).await {
        Ok(data) => {
            let mut map: HashMap<(String, String), model::ManufacturerReport> = HashMap::new();
            for d in data.iter() {
                let record = map
                    .entry((d.manufacturer.clone(), d.service_name.clone()))
                    .or_insert(model::ManufacturerReport {
                        manufacturer: d.manufacturer.clone(),
                        service: d.service_name.clone(),
                        number: 0,
                        total_price: 0.0,
                    });
                record.number += 1;
                record.total_price += d.service_price;
            };
            let report: Vec<model::ManufacturerReport> = map.into_iter().map(|(_id, rep) | rep).collect();
            report
        },
        Err(e) => {
            println!("{}", e.to_string());
            Vec::new()
        },
    }
}



pub async fn get_appointments_for_service_request(car_service_id: u32) -> Result<Vec<model::AppointmentInfo>, reqwest::Error> {
    let url = format!("http://localhost:8084/api/appointments/car-service/{}", car_service_id);
    reqwest::get(url).await?.json::<Vec<model::AppointmentInfo>>().await
}