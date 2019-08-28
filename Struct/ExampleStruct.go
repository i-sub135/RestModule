package TypeStruct

type (
	RoleUser struct {
		RoleID   uint   `json:"roleId"`
		UserID   uint   `json:"userId"`
		UserType string `json:"userType"`
	}

	ViewHistory struct {
		IDBooking      uint    `json:"idBooking"`
		IDUser         uint    `json:"idUser"`
		KodeBooking    string  `json:"kodeBooking"`
		TanggalBooking string  `json:"tglBooking"`
		TanggalSelesai string  `json:"tglSelesai"`
		KeluhanService string  `json:"Keluhan"`
		BiayaService   uint    `json:"Biaya"`
		AlasanBatal    string  `json:"alasanBatal"`
		BookStatus     string  `json:"bookStatus"`
		NamaBengkel    string  `json:"namaBengkel"`
		NomorTelepon   string  `json:"telp"`
		Email          string  `json:"email"`
		AlamatLengkap  string  `json:"alamatBengkel"`
		Latitude       float64 `json:"Latitude"`
		Lohitude       float64 `json:"Longitude"`
		PoliceNo       string  `json:"noPolisi"`
		Brand          string  `json:"Brand"`
		ModelName      string  `json:"modelNama"`
		Year           string  `json:"Tahun"`
		VehicleType    string  `json:"typeKendaraan"`
		Pic            string  `json:"Gambar"`
		IDBengkel      uint    `json:"idBengkel"`
		FotoBengkel    string  `json:"gambarBengkel"`
	}
)

func (RoleUser) TableName() string {
	return "role_user"
}
func (ViewHistory) TableName() string {
	return "view_history_2"
}
