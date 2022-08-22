use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize)]
pub struct AppointmentInfo {
	#[serde(rename = "Id")] pub id: u32,
	#[serde(rename = "UserId")] pub user_id: u32,
	#[serde(rename = "FirstName")] pub first_name: String,
	#[serde(rename = "LastName")] pub last_name: String,
	#[serde(rename = "Email")] pub email: String,
	#[serde(rename = "VehicleId")] pub vehicle_id: u32,
	#[serde(rename = "Manufacturer")] pub manufacturer: String,
	#[serde(rename = "CarModel")] pub car_model: String,
	#[serde(rename = "LicencePlate")] pub licence_plate: String,
	#[serde(rename = "ChassisNumber")] pub chassis_number: String,
	#[serde(rename = "ServiceId")] pub service_id: u32,
	#[serde(rename = "ServiceName")] pub service_name: String,
	#[serde(rename = "ServicePrice")] pub service_price: f32,
	#[serde(rename = "CarServiceId")] pub car_service_id: u32,
	#[serde(rename = "CarServiceName")] pub car_service_name: String,
	#[serde(rename = "Street")] pub street: String,
	#[serde(rename = "StartTime")] pub start_time: String,
	#[serde(rename = "FinishTime")] pub finish_time: String,
	#[serde(rename = "Status")] pub status: String,
}

#[derive(Serialize)]
pub struct ManufacturerReport {
	#[serde(rename = "Manufacturer")] pub manufacturer: String,
	#[serde(rename = "Service")] pub service: String,
	#[serde(rename = "Number")] pub number: u32,
	#[serde(rename = "TotalPrice")] pub total_price: f32
}

#[derive(Serialize)]
pub struct ServiceReport {
	#[serde(rename = "Service")] pub service: String,
	#[serde(rename = "Number")] pub number: u32,
	#[serde(rename = "TotalPrice")] pub total_price: f32
}

#[derive(Serialize)]
pub struct FinancialReport {
	#[serde(rename = "Number")] pub number: u32,
	#[serde(rename = "TotalPrice")] pub total_price: f32
}

#[derive(Serialize)]
pub struct StatusReport {
	#[serde(rename = "ServiceStatus")] pub service_status: String,
	#[serde(rename = "Number")] pub number: u32
}