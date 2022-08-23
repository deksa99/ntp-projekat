use std::collections::HashMap;

use crate::model;

pub async fn man_report(car_service_id: u32) -> Vec<model::ManufacturerReport> {
    match get_appointments_for_service_request(car_service_id).await {
        Ok(data) => {
            let filtered = data.iter().filter(|a| a.status == "Finished");
            let mut map: HashMap<(String, String), model::ManufacturerReport> = HashMap::new();
            for d in filtered {
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

pub async fn service_report(car_service_id: u32) -> Vec<model::ServiceReport> {
    match get_appointments_for_service_request(car_service_id).await {
        Ok(data) => {
            let filtered = data.iter().filter(|a| a.status == "Finished");
            let mut map: HashMap<String, model::ServiceReport> = HashMap::new();
            for d in filtered {
                let record = map
                    .entry(d.service_name.clone())
                    .or_insert(model::ServiceReport {
                        service: d.service_name.clone(),
                        number: 0,
                        total_price: 0.0,
                    });
                record.number += 1;
                record.total_price += d.service_price;
            };
            let report: Vec<model::ServiceReport> = map.into_iter().map(|(_id, rep)| rep).collect();
            report
        },
        Err(e) => {
            println!("{}", e.to_string());
            Vec::new()
        },
    }
}

pub async fn financial_report(car_service_id: u32) -> model::FinancialReport {
    match get_appointments_for_service_request(car_service_id).await {
        Ok(data) => {
            let report: Vec<model::FinancialReport> = data.into_iter().filter(|a| a.status == "Finished").map(|d| {
                model::FinancialReport {
                    number: 0,
                    total_price: d.service_price
                }
            }).collect();
            let init = model::FinancialReport {
                number: 0,
                total_price: 0.0
            };
            report.iter().fold(init, |acc, x| {
                model::FinancialReport {
                    number: acc.number + 1,
                    total_price: acc.total_price + x.total_price
                }
            })
        },
        Err(e) => {
            println!("{}", e.to_string());
            model::FinancialReport {
                number: 0,
                total_price: 0.0
            }
        },
    }
}

pub async fn status_report(car_service_id: u32) -> Vec<model::StatusReport> {
    match get_appointments_for_service_request(car_service_id).await {
        Ok(data) => {
            let mut map: HashMap<String, model::StatusReport> = HashMap::new();
            for d in data.iter() {
                let record = map
                    .entry(d.status.clone())
                    .or_insert(model::StatusReport {
                        service_status: d.status.clone(),
                        number: 0,
                    });
                record.number += 1;
            };
            let report: Vec<model::StatusReport> = map.into_iter().map(|(_id, rep)| rep).collect();
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