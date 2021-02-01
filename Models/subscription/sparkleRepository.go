package subscription

import (
	"database/sql"
	"fmt"
	"log"

	// "github.com/engajerest/auth/Models/users"
	"github.com/engajerest/auth/utils/Errors"
	database "github.com/engajerest/auth/utils/dbconfig"
	"golang.org/x/crypto/bcrypt"
)

const (
	getAllCategoryQuery            = "SELECT categoryid,categoryname,categorytype,sortorder,status FROM app_category WHERE STATUS='Active'"
	getAllSubCategoryQuery         = "SELECT  subcategoryid,categorytypeid,categoryid,subcategoryname,status,icon FROM app_subcategory WHERE statuscode=1"
	getAllPackageQuery             = "SELECT packageid,moduleid,packagename,packageamount,paymentmode,packagecontent,packageicon FROM app_package  WHERE STATUS='Active'"
	insertTenantInfoQuery          = "INSERT INTO tenantinfo (createdby,registrationno,tenantname,primaryemail,primarycontact,bizcategoryid,bizsubcategoryid,Address,state,city,latitude,longitude,postcode,countrycode,timezone,currencycode) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	insertTenantLocationQuery      = "INSERT INTO tenantlocation (tenantid,locationname,email,contactno,address,state,city,latitude,longitude,postcode,countrycode,opentime,closetime,createdby) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	insertTenantSubscription       = "INSERT INTO tenantsubscription (tenantid,transactiondate,packageid,moduleid,currencyid,subscriptionprice,quantity,taxid,taxamount,totalamout,paymentstatus,paymentid) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)"
	getSubscribedDataQuery         = "SELECT a.tenantid, a.tenantname ,b.moduleid, c.name FROM tenantinfo a,tenantsubscription b,app_module c WHERE a.tenantid=b.tenantid AND b.moduleid=c.moduleid AND a.tenantid=?"
	createLocationQuery            = "INSERT INTO tenantlocation (tenantid,locationname,email,contactno,address,state,city,latitude,longitude,postcode,countrycode,opentime,closetime,createdby) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	getLocationbyid                = "SELECT  locationid,locationname,address,city,state,postcode,latitude,longitude,countrycode,opentime,closetime,createdby,status FROM tenantlocation WHERE status='Active' AND locationid=? "
	getAllLocations                = "SELECT  locationid,locationname,tenantid,email,contactno,address,city,state,postcode,latitude,longitude,countrycode,opentime,closetime,createdby,status FROM tenantlocation WHERE status='Active' AND createdby=? "
	createTenantUserQuery          = "INSERT INTO app_users (authname,password,hashsalt,contactno,roleid,referenceid) VALUES(?,?,?,?,?,?)"
	insertTenantUsertoProfileQuery = "INSERT INTO app_userprofile (userid,firstname,lastname,email,contactno,userlocationid) VALUES(?,?,?,?,?,?)"
	getAllTenantUsers              = "SELECT a.firstname,a.lastname,a.userlocationid,a.userid,a.created,a.contactno,a.email,a.status,b.locationname,c.referenceid FROM app_userprofile a, tenantlocation b, app_users c WHERE  a.userid=c.userid AND a.userlocationid=b.locationid AND c.referenceid=b.tenantid AND b.tenantid=?"
    updateTenantUser  ="UPDATE app_users a, app_userprofile b  SET  a.authname=?,a.contactno=?,b.firstname=?,b.lastname=?,b.email=?,b.contactno=?,b.userlocationid=? WHERE a.userid=b.userid AND a.userid=?"
getAllTenantUserByLocationId=""
updateTenantBusiness ="UPDATE tenantinfo SET brandname=?,tenantaccid=?,tenantinfo=?,paymode1=?,paymode2=? WHERE tenantid=?"
insertSocialInfo="INSERT INTO tenantsocial (tenantid,socialprofile,sociallink,socialicon) VALUES(?,?,?,?)"
updateauthuser="UPDATE  app_users a,app_userprofile b SET a.referenceid=?,b.userlocationid=? WHERE a.userid=b.userid AND a.userid=?"
)

func GetAllCategory() []Category {
	print("st1")
	stmt, err := database.Db.Prepare(getAllCategoryQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var categorylist []Category

	for rows.Next() {
		var category Category
		err := rows.Scan(&category.CategoryID, &category.Name, &category.Typeid, &category.SortOrder, &category.Status)
		if err != nil {
			log.Fatal(err)
		}
		categorylist = append(categorylist, category)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return categorylist

}

func GetAllSubCategory() []SubCategory {
	stmt, err := database.Db.Prepare(getAllSubCategoryQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var subcategorylist []SubCategory

	for rows.Next() {
		var subcategory SubCategory
		err := rows.Scan(&subcategory.SubCategoryID, &subcategory.Typeid, &subcategory.CategoryID, &subcategory.Name, &subcategory.Status, &subcategory.Icon)
		if err != nil {
			log.Fatal(err)
		}
		subcategorylist = append(subcategorylist, subcategory)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return subcategorylist
}
func GetAllPackages() []Packages {
	stmt, err := database.Db.Prepare(getAllPackageQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var packageList []Packages

	for rows.Next() {
		var pack Packages
		err := rows.Scan(&pack.PackageID, &pack.ModuleID, &pack.Name, &pack.PackageAmount, &pack.PaymentMode, &pack.PackageContent, &pack.PackageIcon)
		if err != nil {
			log.Fatal(err)
		}
		packageList = append(packageList, pack)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return packageList
}

func (info *SubscriptionData) CreateTenant(userid int) (int64, error) {

	fmt.Println("0")
	print(userid)
	statement, err := database.Db.Prepare(insertTenantInfoQuery)
	print(statement)
	fmt.Println("1")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer statement.Close()

	fmt.Println("2")
	res, err := statement.Exec(userid, &info.Info.Regno, &info.Info.Name, &info.Info.Email, &info.Info.Mobile, &info.Info.CategoryId, &info.Info.SubCategoryID,
		&info.Address.Address, &info.Address.State, &info.Address.Suburb, &info.Address.Latitude, &info.Address.Longitude, &info.Address.Zip, &info.Address.Countrycode,&info.Address.TimeZone,&info.Address.CurrencyCode)
	if err != nil {
		log.Fatal(err)

	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())

	}
	log.Print("Row inserted!")
	return id, nil
}
func (info *SubscriptionData) InsertTenantLocation(tenantid int64,userid int) int64 {
	statement, err := database.Db.Prepare(insertTenantLocationQuery)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	res, err := statement.Exec(tenantid,&info.Info.Name,&info.Info.Email,&info.Info.Mobile, &info.Address.Address, &info.Address.State, &info.Address.Suburb, &info.Address.Latitude, &info.Address.Longitude, &info.Address.Zip, &info.Address.Countrycode,&info.Address.OpenTime,&info.Address.CloseTime,userid)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted in tenant location!")
	return id
}

func (info *TenantSubscription) InsertSubscription(tenantid int64) int64 {
	statement, err := database.Db.Prepare(insertTenantSubscription)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	res, err := statement.Exec(tenantid, &info.Date, &info.PackageId, &info.ModuleId, &info.CurrencyId, &info.Price, &info.Quantity, &info.TaxId, &info.TaxAmount,
		&info.TotalAmount, &info.PaymentStatus, &info.PaymentId)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted in tenant subscription!")
	return id
}

func (info *SubscribedData) GetSubscribedData(tenantid int64) (*SubscribedData, error) {
	fmt.Println("enrty in getsubscription")
	print(tenantid)
	var data SubscribedData
	stmt, err := database.Db.Prepare(getSubscribedDataQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(tenantid)
	err = row.Scan(&data.TenantID, &data.TenantName, &data.ModuleID, &data.ModuleName)
	print(err)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
			data1 := Errors.RestError{}
			data1.Error = err
			return &data, err
		} else {
			log.Fatal(err)
			fmt.Println("nodata")
			var data1 *Errors.RestError
			data1.Error = err
			return &data, err
		}

	}
	// user.Check=true
	return &data, err
}

func (loco *Location) CreateLocation(id int64) (int64, error) {
	statement, err := database.Db.Prepare(createLocationQuery)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	res, err := statement.Exec(&loco.TenantID, &loco.LocationName,&loco.Email,&loco.Mobile, &loco.Address, &loco.State, &loco.Suburb, &loco.Latitude, &loco.Longitude, &loco.Zip, &loco.Countrycode, &loco.OpeningTime, &loco.ClosingTime, id)
	if err != nil {
		log.Fatal(err)
	}
	id, err1 := res.LastInsertId()
	if err1 != nil {
		log.Fatal("Error:", err1.Error())
	}
	log.Print("Row inserted in createlocation!")
	return id, nil
}

func (loco *Location) GetLocationById(id int64) (*Location, error) {
	fmt.Println("enrty in getlocation")
	print(id)
	var data Location
	stmt, err := database.Db.Prepare(getLocationbyid)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	// print(row)
	err = row.Scan(&data.LocationId, &data.LocationName, &data.Address, &data.Suburb, &data.State, &data.Zip, &data.Latitude, &data.Longitude, &data.Countrycode, &data.OpeningTime, &data.ClosingTime, &data.Createdby, &data.Status)
	print(err)
	fmt.Println("2")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
		} else {
			log.Fatal(err)
		}
	}
	// fmt.Println(user)

	fmt.Println("completed")
	return &data, nil
}
func GetAllTenantUsersLocation(id int) []Location {
	stmt, err := database.Db.Prepare(getAllLocations)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var locationlist []Location
	for rows.Next() {
		var data Location
		err = rows.Scan(&data.LocationId, &data.LocationName, &data.TenantID,&data.Email,&data.Mobile, &data.Address, &data.Suburb, &data.State, &data.Zip, &data.Latitude, &data.Longitude, &data.Countrycode, &data.OpeningTime, &data.ClosingTime, &data.Createdby, &data.Status)
		if err != nil {
			log.Fatal(err)
		}
		
		locationlist = append(locationlist, data)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return locationlist
}
func (user *TenantUser) CreateTenantUser() int64 {
	fmt.Println("0")
	statement, err := database.Db.Prepare(createTenantUserQuery)
	fmt.Println("1")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	hashedPassword, err := HashPassword(user.Password)
	fmt.Println("2")

	res, err := statement.Exec(&user.Email, &user.Password, &hashedPassword, &user.Mobile, &user.RoleId, &user.TenantID)
	if err != nil {

		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted in tenantuser!")
	return id
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (user *TenantUser) InsertTenantUserintoProfile(id int64) int64 {
	statement, err := database.Db.Prepare(insertTenantUsertoProfileQuery)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	res, err := statement.Exec(id, &user.FirstName, &user.LastName, &user.Email, &user.Mobile, &user.Locationid)
	if err != nil {
		log.Fatal(err)
	}
	id, err1 := res.LastInsertId()
	if err1 != nil {
		log.Fatal("Error:", err1.Error())
	}
	log.Print("Row inserted in tenant user profile!")
	return id
}
func GetAllTenantUsers(id int) []TenantUser {
	stmt, err := database.Db.Prepare(getAllTenantUsers)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var tenantuserlist []TenantUser
	for rows.Next() {
		var data TenantUser
		err = rows.Scan(
			&data.FirstName,
			&data.LastName,
			&data.Locationid,&data.Userid,&data.Created,&data.Mobile,&data.Email,
			&data.Status,&data.Locationname,&data.Referenceid)
		if err != nil {
			log.Fatal(err)
		}

		tenantuserlist = append(tenantuserlist, data)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return tenantuserlist
}
func (user *TenantUser) UpdateTenantUser() bool {
	statement, err := database.Db.Prepare(updateTenantUser)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	_, err = statement.Exec(user.Email,user.Mobile,user.FirstName,user.LastName,user.Email,user.Mobile,user.Locationid,user.Userid)
	if err != nil {
		log.Fatal(err)
	}
	
	log.Print("Row updated in tenant user profile!")
	return true
}
func (user *BusinessUpdate) UpdateTenantBusiness() bool {
	statement, err := database.Db.Prepare(updateTenantBusiness)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	_, err = statement.Exec(user.Brandname,user.TenantaccId,user.About,user.Paymode1,user.Paymode2,&user.TenantID)
	if err != nil {
		log.Fatal(err)
	}
	
	log.Print("Row updated in business user profile!")
	return true
}
func (info *BusinessUpdate) InsertTenantSocial() int64 {
	statement, err := database.Db.Prepare(insertSocialInfo)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	res, err := statement.Exec(info.TenantID,info.SociaProfile,info.SocialLink,info.SocialIcon)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted in tenant social!")
	return id
}
func (info *AuthUser) UpdateAuthUser(userid int) bool {
	statement, err := database.Db.Prepare(updateauthuser)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	_, err1 := statement.Exec(&info.TenantID,&info.LocationId,userid)

	if err1 != nil {
		log.Fatal(err1)
	
	}
	
	log.Print("Row updated in auth user profile!")
	return true
	
}