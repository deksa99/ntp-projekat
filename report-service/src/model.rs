use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize)]
pub struct AppointmentInfo {
	#[serde(rename = "Id")]
	id: u32,
	#[serde(rename = "UserId")]
	user_id: u32,
	#[serde(rename = "FirstName")]
	first_name: String,
	#[serde(rename = "LastName")]
	last_name: String,
	#[serde(rename = "Email")]
	email: String,
	#[serde(rename = "VehicleId")]
	vehicle_id: u32,
	#[serde(rename = "Manufacturer")]
	manufacturer: String,
	#[serde(rename = "CarModel")]
	car_model: String,
	#[serde(rename = "LicencePlate")]
	licence_plate: String,
	#[serde(rename = "ChassisNumber")]
	chassis_number: String,
	#[serde(rename = "ServiceId")]
	service_id: u32,
	#[serde(rename = "ServiceName")]
	service_name: String,
	#[serde(rename = "ServicePrice")]
	service_price: f32,
	#[serde(rename = "CarServiceId")]
	car_service_id: u32,
	#[serde(rename = "CarServiceName")]
	car_service_name: String,
	#[serde(rename = "Street")]
	street: String,
	// #[serde(rename = "StartTime")]
	// start_time: String,
	// #[serde(rename = "FinishTime")]
	// finish_time: String,
	// #[serde(rename = "Status")]
	// status: String,
}