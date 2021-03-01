package subscription

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/engajerest/auth/Models/users"
	"github.com/engajerest/auth/utils/Errors"
	"github.com/engajerest/auth/utils/dbconfig"
	database "github.com/engajerest/auth/utils/dbconfig"
	"github.com/engajerest/sparkle/graph/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	getAllCategoryQuery            = "SELECT categoryid,categoryname,categorytype,sortorder,status FROM app_category WHERE STATUS='Active'"
	getAllSubCategoryQuery         = "SELECT  subcategoryid,categorytypeid,categoryid,subcategoryname,status,icon FROM app_subcategory WHERE statuscode=1"
	getAllPackageQuery             = "SELECT a.packageid,a.moduleid,a.packagename,a.packageamount,a.paymentmode,a.packagecontent,a.packageicon,b.modulename FROM app_package a,app_module b  WHERE a.STATUS='Active' AND a.moduleid=b.moduleid"
	insertTenantInfoQuery          = "INSERT INTO tenants (createdby,registrationno,tenantname,primaryemail,primarycontact,bizcategoryid,bizsubcategoryid,Address,state,city,latitude,longitude,postcode,countrycode,timezone,currencycode,tenanttoken) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	insertTenantLocationQuery      = "INSERT INTO tenantlocation (tenantid,locationname,email,contactno,address,state,city,latitude,longitude,postcode,countrycode,opentime,closetime,createdby) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	insertTenantSubscription       = "INSERT INTO tenantsubscription (tenantid,transactiondate,packageid,moduleid,currencyid,subscriptionprice,quantity,taxid,taxamount,totalamout,paymentstatus,paymentid) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)"
	getSubscribedDataQuery         = "SELECT a.tenantid, a.tenantname ,b.moduleid, c.name FROM tenants a,tenantsubscription b,app_module c WHERE a.tenantid=b.tenantid AND b.moduleid=c.moduleid AND a.tenantid=?"
	createLocationQuery            = "INSERT INTO tenantlocation (tenantid,locationname,email,contactno,address,state,city,latitude,longitude,postcode,countrycode,opentime,closetime,createdby) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	getLocationbyid                = "SELECT  locationid,locationname,address,city,state,postcode,latitude,longitude,countrycode,opentime,closetime,createdby,status FROM tenantlocation WHERE status='Active' AND locationid=? "
	getAllLocations                = "SELECT  locationid,locationname,tenantid,email,contactno,address,city,state,postcode,latitude,longitude,countrycode,opentime,closetime,createdby,status FROM tenantlocation WHERE status='Active' AND tenantid=? "
	createTenantUserQuery          = "INSERT INTO app_users (authname,password,hashsalt,contactno,roleid,referenceid) VALUES(?,?,?,?,?,?)"
	insertTenantUsertoProfileQuery = "INSERT INTO app_userprofiles (userid,firstname,lastname,email,contactno,userlocationid) VALUES(?,?,?,?,?,?)"
	getAllTenantUsers              = "SELECT a.firstname,a.lastname,a.userlocationid,a.userid,a.created,a.contactno,a.email,a.status,b.locationname,c.referenceid FROM app_userprofiles a, tenantlocation b, app_users c WHERE  a.userid=c.userid AND a.userlocationid=b.locationid AND c.referenceid=b.tenantid AND b.tenantid=?"
	updateTenantUser               = "UPDATE app_users a, app_userprofiles b  SET  a.authname=?,a.contactno=?,b.firstname=?,b.lastname=?,b.email=?,b.contactno=?,b.userlocationid=? WHERE a.userid=b.userid AND a.userid=?"
	getAllTenantUserByLocationId   = ""
	updateTenantBusiness           = "UPDATE tenant SET brandname=?,tenantaccid=?,tenantinfo=?,paymode1=?,paymode2=? WHERE tenantid=?"
	insertSocialInfo               = "INSERT INTO tenantsocial (tenantid,socialprofile,sociallink,socialicon) VALUES"
	updatesocialinfo               = "UPDATE tenantsocial SET socialprofile=?, sociallink=?,socialicon=? WHERE tenantid=? AND socialid=? "
	updateauthuser                 = "UPDATE  app_users a,app_userprofiles b SET a.referenceid=?,b.userlocationid=? WHERE a.userid=b.userid AND a.userid=?"
	getBusinessbyid                = "SELECT tenantid,IFNULL(brandname,'') AS brandname,IFNULL(tenantinfo,'') AS tenantinfo,IFNULL(paymode1,0) AS paymode1,IFNULL(paymode2,0) AS paymode2,IFNULL(tenantaccid,0) AS tenantaccid,IFNULL(address,'') AS address,IFNULL(primaryemail,'') AS primaryemail,IFNULL(primarycontact,'') AS  primarycontact,IFNULL(tenanttoken,'') AS tenanttoken FROM tenants WHERE tenantid=?"
	getAllSocial                   = "SELECT socialid, IFNULL(socialprofile,'') AS socialprofile , IFNULL(sociallink,'') AS sociallink, IFNULL(socialicon,'') AS socialicon FROM tenantsocial WHERE tenantid= ?"
	userAuthentication             = "SELECT a.userid,b.firstname,b.lastname,b.email,b.contactno,b.status,b.created FROM app_users a, app_userprofiles b WHERE a.userid=b.userid AND a.status ='Active' AND a.userid=?"
	Getpromotions          = "SELECT a.promotionid,a.promotiontypeid,a.tenantid,a.promoname,a.promocode,a.promoterms,a.promovalue,a.startdate,a.enddate,a.status,b.typename,b.tag, c.tenantname FROM promotions a, promotiontypes b,tenants c WHERE a.promotiontypeid=b.promotiontypeid AND a.tenantid=c.tenantid AND a.`status`='Active' AND a.tenantid=?"
    createpromotion ="INSERT INTO promotions (promotiontypeid,tenantid,promoname,promocode,promoterms,promovalue,startdate,enddate,createdby) VALUES(?,?,?,?,?,?,?,?,?)"
    insertsequence = "INSERT INTO ordersequence (tenantid,tablename,seqno,prefix,subprefix) VALUES(?,?,?,?,?)"
insertcharge = "INSERT INTO tenantcharges (tenantid,locationid,chargeid,chargetype,chargevalue,createdby) VALUES"
insertdelivery = "INSERT INTO tenantsettings (tenantid,locationid,slabtype,slab,slablimit,slabcharge,createdby) VALUES"
updatecharge = "UPDATE tenantcharges SET locationid=?,chargeid=?,chargetype=?, chargevalue=? WHERE tenantchargeid=? AND tenantid=?"
updatedelivery="UPDATE  tenantsettings SET locationid=?,slabtype=?,slab=?,slablimit=?,slabcharge=? WHERE settingsid=? AND tenantid=?"
deletecharge = "DELETE FROM tenantcharges WHERE tenantchargeid=?"
deletedelivery = "DELETE FROM  tenantsettings WHERE settingsid=?"

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
		err := rows.Scan(&pack.PackageID, &pack.ModuleID, &pack.Name, &pack.PackageAmount, &pack.PaymentMode, &pack.PackageContent, &pack.PackageIcon, &pack.ModuleName)
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
		&info.Address.Address, &info.Address.State, &info.Address.Suburb, &info.Address.Latitude, &info.Address.Longitude, &info.Address.Zip, &info.Address.Countrycode, &info.Address.TimeZone, &info.Address.CurrencyCode,&info.Info.Tenanttoken)
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
func (info *SubscriptionData) InsertTenantLocation(tenantid int64, userid int) int64 {
	statement, err := database.Db.Prepare(insertTenantLocationQuery)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	res, err := statement.Exec(tenantid, &info.Info.Name, &info.Info.Email, &info.Info.Mobile, &info.Address.Address, &info.Address.State, &info.Address.Suburb, &info.Address.Latitude, &info.Address.Longitude, &info.Address.Zip, &info.Address.Countrycode, &info.Address.OpenTime, &info.Address.CloseTime, userid)
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
	res, err := statement.Exec(&loco.TenantID, &loco.LocationName, &loco.Email, &loco.Mobile, &loco.Address, &loco.State, &loco.Suburb, &loco.Latitude, &loco.Longitude, &loco.Zip, &loco.Countrycode, &loco.OpeningTime, &loco.ClosingTime, id)
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
		err = rows.Scan(&data.LocationId, &data.LocationName, &data.TenantID, &data.Email, &data.Mobile, &data.Address, &data.Suburb, &data.State, &data.Zip, &data.Latitude, &data.Longitude, &data.Countrycode, &data.OpeningTime, &data.ClosingTime, &data.Createdby, &data.Status)
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
func LocationTest(id int) []Tenantlocation {

	// var data Location

	// dsn := "cegxbczruu:Package@123#@tcp(139.59.69.8:3306)/cegxbczruu?charset=utf8mb4&parseTime=True&loc=Local"
	// // // DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// // DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// sqlDB, err := sql.Open("mysql", dsn)
	DB, err := gorm.Open(mysql.New(mysql.Config{Conn: dbconfig.Db}), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")

	} else {
		log.Println("Connection Established")
	}
	// var user []Tenantlocation
	// err = DB.Debug().Where(&Tenantlocation{Tenantid:37}).Find(&user).Error
	// if err != nil {
	//     fmt.Println("Query failed")
	// }
	// fmt.Println(user)

	// var emails []App_userprofile
	// err = DB.Debug().Model(&user[1]).Related(&emails, "Appuserprofiles").Error
	// if err != nil {
	//     fmt.Println("Query failed")
	// }
	// fmt.Println(emails)

	// // user.Appuserprofiles = emails
	// fmt.Println(user)
	var orders []Tenantlocation

	DB.Table("tenantlocation").Preload("Appuserprofiles").Where("tenantid=?", id).Find(&orders)

	fmt.Println(orders)
	for index, value := range orders {
		fmt.Println(index, " = ", value)

	}

	// Tlocations := &Tenantlocations{}

	// rows, err := DB.Table("tenantlocation").Where("tenantlocation.tenantid= ? and tenantlocation.status = ?", id, "Active").
	// 	Joins("Join app_userprofile on app_userprofile.userlocationid = tenantlocation.locationid").
	// 	// Joins("Join items on items.id = order_items.id").
	// 	Select("tenantlocation.locationid,tenantlocation.locationname,tenantlocation.tenantid,tenantlocation.email,tenantlocation.contactno,tenantlocation.address,tenantlocation.city,tenantlocation.state,tenantlocation.postcode,tenantlocation.latitude,tenantlocation.longitude,tenantlocation.countrycode,tenantlocation.opentime,tenantlocation.closetime,tenantlocation.createdby,tenantlocation.status,app_userprofile.userid,app_userprofile.firstname,app_userprofile.lastname,app_userprofile.email,app_userprofile.contactno,app_userprofile.userlocationid").Rows()
	// if err != nil {
	// 	log.Panic(err)
	// }

	// defer rows.Close()
	// // Values to load into
	// print("1.1")
	// tentlocation := &Tenantlocation{}
	// tentlocation.Tenantusers = make([]Appprofile, 0)
	// print("1.2")
	// for rows.Next() {
	// 	print("1.3")
	// 	profile := Appprofile{}
	// 	// item := Item{}
	// 	err = rows.Scan(&tentlocation.LocationId, &tentlocation.LocationName, &tentlocation.TenantID, &tentlocation.Email, &tentlocation.Mobile, &tentlocation.Address, &tentlocation.Suburb, &tentlocation.State, &tentlocation.Zip, &tentlocation.Latitude, &tentlocation.Longitude, &tentlocation.Countrycode, &tentlocation.OpeningTime, &tentlocation.ClosingTime, &tentlocation.Createdby, &tentlocation.Status,&profile.Userid,&profile.FirstName,&profile.LastName,&profile.Email,&profile.Mobile,&profile.Userlocationid)
	// 	if err != nil {
	// 		print("1.4")
	// 		log.Panic(err)
	// 	}
	// 	print("1.5")
	// 	// orderItem.Item = item
	// 	tentlocation.Tenantusers = append(tentlocation.Tenantusers, profile)
	// 	print("1.6")
	// }
	// log.Print(pretty.Sprint(tentlocation))

	// print("st1")
	// rows, err := DB.Table("tenantlocation").Select("locationid,locationname,tenantid,email,contactno,address,city,state,postcode,latitude,longitude,countrycode,opentime,closetime,createdby,status").Where("tenantid", id).Rows()
	// print("st222")
	// for rows.Next() {
	// 	var data Tenantlocation
	// 	// DB.ScanRows(rows, &data)
	// 	err = rows.Scan(&data.LocationId, &data.LocationName, &data.TenantID, &data.Email, &data.Mobile, &data.Address, &data.Suburb, &data.State, &data.Zip, &data.Latitude, &data.Longitude, &data.Countrycode, &data.OpeningTime, &data.ClosingTime, &data.Createdby, &data.Status)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	loc = append(loc, data)
	// }

	return orders

}
func userget(ids []int) ([]*users.User, []error) {
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i := 0; i < len(ids); i++ {
		placeholders[i] = "?"
		args[i] = i
	}

	res, err := database.Db.Prepare(getAllLocations)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	rows, err := res.Query(ids)
	if err != nil {
		log.Fatal(err)
	}
	userById := map[int]*users.User{}
	for rows.Next() {
		user := users.User{}
		err := rows.Scan(&user.ID, &user.FirstName)
		if err != nil {
			panic(err)
		}
		userById[user.ID] = &user
	}

	users := make([]*users.User, len(ids))
	for i, id := range ids {
		users[i] = userById[id]
	}

	return users, nil
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
			&data.Locationid, &data.Userid, &data.Created, &data.Mobile, &data.Email,
			&data.Status, &data.Locationname, &data.Referenceid)
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
	_, err = statement.Exec(user.Email, user.Mobile, user.FirstName, user.LastName, user.Email, user.Mobile, user.Locationid, user.Userid)
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
	_, err = statement.Exec(user.Brandname, user.TenantaccId, user.About, user.Paymode1, user.Paymode2, &user.TenantID)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Row updated in business user profile!")
	return true
}
func (info *BusinessUpdate) InsertTenantSocial(soc []Social, id int) error {

	var inserts []string
	var params []interface{}
	for _, v := range soc {
		inserts = append(inserts, "(?, ?, ?, ?)")
		params = append(params, id, v.SociaProfile, v.SocialLink, v.SocialIcon)
	}
	queryVals := strings.Join(inserts, ",")
	query := insertSocialInfo + queryVals
	log.Println("query is", query)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := database.Db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, params...)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created simulatneously", rows)
	return nil

}
func (info *Social) UpdateTenantSocial(tenantid int) bool {

	statement, err := database.Db.Prepare(updatesocialinfo)
	print(statement)

	if err != nil {
		log.Fatal(err)
		return false
	}
	defer statement.Close()
	_, err = statement.Exec(&info.SociaProfile,&info.SocialLink,&info.SocialIcon,tenantid,&info.Socialid,)
	if err != nil {
		log.Fatal(err)
		return false
	}

	log.Print("Row updated in tenantsocial!")
	return true

}
func (info *AuthUser) UpdateAuthUser(userid int) bool {
	statement, err := database.Db.Prepare(updateauthuser)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	_, err1 := statement.Exec(&info.TenantID, &info.LocationId, userid)

	if err1 != nil {
		log.Fatal(err1)

	}

	log.Print("Row updated in auth user profile!")
	return true

}
func (business *BusinessUpdate) GetBusinessInfo(id int) (*BusinessUpdate, bool) {
	var data BusinessUpdate
	stmt, err := database.Db.Prepare(getBusinessbyid)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	// print(row)
	err = row.Scan(&data.TenantID, &data.Brandname, &data.About, &data.Paymode1, &data.Paymode2, &data.TenantaccId,&data.Address,&data.Email,&data.Phone,&data.Tenanttoken)
	print(err)
	fmt.Println("2")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
			return nil, false
		} else {
			log.Fatal(err)
			return nil, false
		}
	}
	// fmt.Println(user)

	fmt.Println("completed")
	return &data, true
}
func GetAllSocial(id int) []Social {
	stmt, err := database.Db.Prepare(getAllSocial)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var sociallist []Social
	for rows.Next() {
		var data Social
		err = rows.Scan(
			&data.Socialid,
			&data.SociaProfile,
			&data.SocialLink,
			&data.SocialIcon,
		)
		if err != nil {
			log.Fatal(err)
		}

		sociallist = append(sociallist, data)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return sociallist
}

func UserAuthentication(id int64) (*users.User, bool, error) {
	fmt.Println("enrty in sparkleauth")
	print(id)
	var data users.User
	stmt, err := database.Db.Prepare(userAuthentication)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(&data.ID, &data.FirstName, &data.LastName, &data.Mobile, &data.Email, &data.Status, &data.CreatedDate)
	print(err)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
			data1 := Errors.RestError{}
			data1.Error = err
			return &data, false, err
		} else {
			log.Fatal(err)
			fmt.Println("nodata")
			var data1 *Errors.RestError
			data1.Error = err
			return &data, false, err
		}

	}
	// user.Check=true
	return &data, true, err
}
func GetAllPromotions(tenantid int ) []Promotion {
	print("st1")
	stmt, err := database.Db.Prepare(Getpromotions)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(tenantid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var promolist []Promotion

	for rows.Next() {
		var promo Promotion
		err := rows.Scan(&promo.Promotionid,&promo.Promotiontypeid,&promo.Tenantid,&promo.Promoname,&promo.Promocode,&promo.Promoterms,&promo.Promovalue,
		&promo.Startdate,&promo.Enddate,&promo.Status,&promo.Promotype,&promo.Promotag,&promo.Tenantname)
		if err != nil {
			log.Fatal(err)
		}
		promolist = append(promolist, promo)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return promolist

}
func (p *Promotion) Createpromotion(created int ) int64 {
	statement, err := database.Db.Prepare(createpromotion)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	res, err := statement.Exec(&p.Promotiontypeid,&p.Tenantid,&p.Promoname,&p.Promocode,&p.Promoterms,
	&p.Promovalue,&p.Startdate,&p.Enddate,created)
	if err != nil {
		log.Fatal(err)
	}
	id, err1 := res.LastInsertId()
	if err1 != nil {
		log.Fatal("Error:", err1.Error())
	}
	log.Print("Row inserted in promotion!")
	return id
}
func Getpromotypes() []*model.Typedata {

	DB, err := gorm.Open(mysql.New(mysql.Config{Conn: dbconfig.Db}), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")

	} else {
		log.Println("Connection Established")
	}

	var data []*model.Typedata

	DB.Table("promotiontypes").Find(&data)

	fmt.Println(data)

	return data

}
func Getchargetypes() []*model.Chargetype {

	DB, err := gorm.Open(mysql.New(mysql.Config{Conn: dbconfig.Db}), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")

	} else {
		log.Println("Connection Established")
	}

	var data []*model.Chargetype

	DB.Table("chargetypes").Find(&data)

	fmt.Println(data)

	return data

}
func (o *Ordersequence) Insertsequence() (int64, error) {

	fmt.Println("0")

	statement, err := database.Db.Prepare(insertsequence)
	print(statement)
	fmt.Println("1")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer statement.Close()

	fmt.Println("2")
	res, err := statement.Exec(&o.Tenantid,&o.Tablename,&o.Seqno,&o.Prefix,&o.Subprefix)
	if err != nil {
		log.Fatal(err)

	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())

	}
	log.Print("Row inserted in sequence!")
	return id, nil
}
func (info *Charge) Insertothercharges(soc []Charge) error {

	var inserts []string
	var params []interface{}
	for _, v := range soc {
		inserts = append(inserts, "(?, ?, ?, ?,?,?)")
		params = append(params,v.Tenantid,v.Locationid,v.Chargeid,v.Chargetype,v.Chargevalue,v.Createdby )
	}
	queryVals := strings.Join(inserts, ",")
	query := insertcharge + queryVals
	log.Println("query is", query)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := database.Db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, params...)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d charges created simulatneously", rows)
	return nil

}
func (info *Delivery) Insertdeliverycharges(soc []Delivery) error {

	var inserts []string
	var params []interface{}
	for _, v := range soc {
		inserts = append(inserts, "(?, ?, ?, ?,?,?,?)")
		params = append(params,v.Tenantid,v.Locationid,v.Slabtype,v.Slab,v.Slablimit,v.Slabcharge,v.Createdby )
	}
	queryVals := strings.Join(inserts, ",")
	query := insertdelivery + queryVals
	log.Println("query is", query)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := database.Db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, params...)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d delivery created simulatneously", rows)
	return nil

}
func (o *Charge) Updateothercharge() bool {
	fmt.Println("0")
	statement, err := database.Db.Prepare(updatecharge)
	fmt.Println("1")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer statement.Close()
	_, err = statement.Exec(&o.Locationid,&o.Chargeid,&o.Chargetype,&o.Chargevalue,&o.Tenantchargeid,&o.Tenantid)
	if err != nil {

		log.Fatal(err)
		return false
	}

	log.Print("Row updated in charges!")
	return true
}
func (o *Delivery) Updatedeliverycharge() bool {
	fmt.Println("0")
	statement, err := database.Db.Prepare(updatedelivery)
	fmt.Println("1")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer statement.Close()
	_, err = statement.Exec(&o.Locationid,&o.Slabtype,&o.Slab,&o.Slablimit,&o.Slabcharge,&o.Settingsid,&o.Tenantid)
	if err != nil {

		log.Fatal(err)
		return false
	}

	log.Print("Row updated in charges!")
	return true
}
func (c *Charge) Deleteothercharge() bool {

	statement, err := database.Db.Prepare(deletecharge)
	print(statement)

	if err != nil {
		log.Fatal(err)
		return false
	}

	_, err = statement.Exec(&c.Tenantchargeid)
	if err != nil {
		log.Fatal(err)
		return false
	}

	log.Print("Row deleted in charges!")
	return true

}
func (c *Delivery) Deletedeliverycharge() bool {

	statement, err := database.Db.Prepare(deletedelivery)
	print(statement)

	if err != nil {
		log.Fatal(err)
		return false
	}

	_, err = statement.Exec(&c.Settingsid)
	if err != nil {
		log.Fatal(err)
		return false
	}

	log.Print("Row deleted in delivery!")
	return true

}