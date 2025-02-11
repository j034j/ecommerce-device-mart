package repo

import (
	"fmt"

	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{
		DB: DB,
	}

}

func (ad *adminDatabase) FindAdminCredentials() (config.AdminCredentials, error) {
	var adminCredentials = config.GetAdminCredentials()
	if adminCredentials.AdminUsername == "" || adminCredentials.AdminPassword == "" {
		return adminCredentials, fmt.Errorf("failed to fetch admin credentials")
	}

	return adminCredentials, nil
}

func (ad *adminDatabase) FetchAllUserData() ([]response.UserData, error) {
	var ListOfAllUsers = make([]response.UserData, 0)
	query := "SELECT Id,user_name ,email,phone ,is_blocked FROM users ORDER BY Id"
	err := ad.DB.Raw(query).Scan(&ListOfAllUsers).Error
	return ListOfAllUsers, err
}

func (ad *adminDatabase) BlockUserByID(userID int) error {
	var BlockedUser response.UserData
	status := true
	query := "UPDATE Users SET Is_blocked =$1  WHERE Id =$2 RETURNING *"
	err := ad.DB.Raw(query, status, userID).Scan(&BlockedUser).Error
	fmt.Println(BlockedUser)
	return err
}

func (ad *adminDatabase) UnblockUserByID(userID int) error {
	var BlockedUser response.UserData
	status := false
	query := "UPDATE Users SET Is_blocked =$1 WHERE id =$2 RETURNING *"
	err := ad.DB.Raw(query, status, userID).Scan(&BlockedUser).Error
	fmt.Println(BlockedUser)
	return err
}

func (ad *adminDatabase) FindUsersByName(name string) ([]response.UserData, error) {
	var users []response.UserData
	query := "SELECT * FROM Users WHERE User_name ILIKE `%$1%`"
	err := ad.DB.Raw(query, name).Scan(&users).Error
	return users, err
}

func (ad *adminDatabase) SetupDB() {
	// ad.DropTable()
	ad.InsertOrderStatus()
	ad.InsertPaymentMethods()
	ad.InsertStates()

}

func (ad *adminDatabase) InsertStates() {
	query := `INSERT INTO states (name) VALUES
('Andhra Pradesh'),
('Arunachal Pradesh'),
('Assam'),
('Bihar'),
('Chhattisgarh'),
('Goa'),
('Gujarat'),
('Haryana'),
('Himachal Pradesh'),
('Jharkhand'),
('Karnataka'),
('Kerala'),
('Madhya Pradesh'),
('Maharashtra'),
('Manipur'),
('Meghalaya'),
('Mizoram'),
('Nagaland'),
('Odisha'),
('Punjab'),
('Rajasthan'),
('Sikkim'),
('Tamil Nadu'),
('Telangana'),
('Tripura'),
('Uttar Pradesh'),
('Uttarakhand'),
('West Bengal');`

	ad.DB.Exec(query)
}

func (ad *adminDatabase) InsertPaymentMethods() {
	query := `INSERT INTO payment_methods (method_name) VALUES ('cash on delivery') , ('online payment') , ('Wallet')`

	ad.DB.Exec(query)
}

// func (ad *adminDatabase) DropTable() {
// 	query := `drop table order_lines; drop table order_statuses ; drop table payment_methods;`

// 	ad.DB.Exec(query)
// }

func (ad *adminDatabase) InsertOrderStatus() {
	query := `INSERT INTO order_statuses (status) VALUES ('Pending'),('Shipped'), ('Delivered'), ('Cancelled'), ('Returned')`

	ad.DB.Exec(query)
}
