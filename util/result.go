package util

type Result struct {
	Age                       string `json:"forma-age"`
	Gender                    string `json:"forma-gender"`
	IndustrialSector          string `json:"forma-industrial_sector"`
	MonthlyHouseholdIncome    string `json:"forma-monthly_household_income"`
	SocialStatusAndOccupation string `json:"forma-social_status_and_occupation"`

	CarJeepVan   string `json:"forma-forma5-car_jeep_van"`
	Motorcycle   string `json:"forma-forma5-motorcycle"`
	PickUp       string `json:"forma-forma5-pick_up"`
	ThreeWheeler string `json:"forma-forma5-three_wheeler"`
	Truck        string `json:"forma-forma5-truck"`
	Others       string `json:"forma-forma5-others"`

	GNDivisions  string `json:"forma-forma7-g_n_divisions"`
	DSDivision   string `json:"forma-forma7-d_s_division"`
	HomeDistrict string `json:"forma-forma7-home_district"`
	Town         string `json:"forma-forma7-town"`

	TransportMean           string `json:"formb-transport_mean"`
	WalkDurationDestination string `json:"formb-walk_duration_destination"`
	WalkDurationTransport   string `json:"formb-walk_duration_transport"`

	FormdFuelCost             string `json:"formd-fuel_cost"`
	FormdTransfer             string `json:"formd-transfer"`
	FormdTransportMeanCommute string `json:"formd-transport_mean_commute"`
	FormdTsStation            string `json:"formd-ts_station"`

	HeaderDate             string `json:"header-date"`
	HeaderFormID           string `json:"header-form_id"`
	HeaderLocation         string `json:"header-location"`
	HeaderLocationAccuracy string `json:"header-location-accuracy"`
	HeaderLocationAltitude string `json:"header-location-altitude"`
	HeaderSurveyorID       string `json:"header-surveyor_id"`
	HeaderSurveyorName     string `json:"header-surveyor_name"`
	HeaderTime             string `json:"header-time"`
	HeaderTravelDistance   string `json:"header-travel_distance"`
	HeaderTripPurpose      string `json:"header-trip_purpose"`

	MetaAudit      string `json:"meta-audit"`
	MetaInstanceID string `json:"meta-instanceID"`

	ModeChoice1 string `json:"mode_choice_1"`
	ModeChoice2 string `json:"mode_choice_2"`
	ModeChoice3 string `json:"mode_choice_3"`
	ModeChoice4 string `json:"mode_choice_4"`
	ModeChoice5 string `json:"mode_choice_5"`
}
