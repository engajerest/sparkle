package subscription

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
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
	getAllPackageQuery             = "SELECT a.packageid,a.moduleid,a.packagename,a.packageamount,a.paymentmode,a.packagecontent,a.packageicon,b.modulename,IFNULL(c.promocodeid,0) AS promocodeid,IFNULL(c.promoname ,'') AS promoname,IFNULL(c.promodescription,'') AS promodescription,IFNULL(c.packageexpiry,0) AS packageexpiry,IFNULL(c.promotype ,'') AS promotype,IFNULL(c.promovalue,0) AS promovalue,IFNULL(c.validity,'') AS validity,IF(c.validity>=DATE(NOW()), true, false) AS validity FROM app_package a Inner JOIN app_module b ON a.moduleid=b.moduleid LEFT OUTER JOIN  app_promocodes c ON a.packageid=c.packageid WHERE a.`status`='Active' "
	insertTenantInfoQuery          = "INSERT INTO tenants (createdby,partnerid,registrationno,tenantname,primaryemail,primarycontact,Address,state,city,latitude,longitude,postcode,countrycode,timezone,currencycode,tenanttoken) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	insertTenantLocationQuery      = "INSERT INTO tenantlocations (tenantid,locationname,email,contactno,address,state,city,latitude,longitude,postcode,countrycode,opentime,closetime,createdby) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	insertTenantSubscription       = "INSERT INTO tenantsubscription (tenantid,transactiondate,packageid,partnerid,moduleid,categoryid,subcategoryid,currencyid,subscriptionprice,quantity,taxid,taxamount,totalamount,paymentstatus,paymentid,promoid,promoprice,promostatus,validitydate) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	getSubscribedDataQuery         = "SELECT a.tenantid, a.tenantname ,b.moduleid,b.subscriptionid,b.categoryid,b.subcategoryid, c.modulename ,d.locationid,d.locationname FROM tenants a,tenantsubscription b,app_module c ,tenantlocations d WHERE a.tenantid=b.tenantid AND b.moduleid=c.moduleid AND a.tenantid = d.tenantid AND  a.tenantid=?"
	createLocationQuery            = "INSERT INTO tenantlocations (tenantid,locationname,email,contactno,address,state,city,latitude,longitude,postcode,countrycode,opentime,closetime,createdby,delivery,deliverytype,deliverymins) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	updatelocation                 = "UPDATE tenantlocations SET locationname=?,email=?,contactno=?,address=?,state=?,city=?,latitude=?,longitude=?,postcode=?,countrycode=?,opentime=?,closetime=?,delivery=?,deliverytype=?,deliverymins=? WHERE tenantid=? AND locationid=?"
	getLocationbyid                = "SELECT  locationid,locationname,address,city,state,postcode,latitude,longitude,countrycode,opentime,closetime,createdby,status,IFNULL(delivery,false) AS delivery,IFNULL(deliverytype,'') AS deliverytype,IFNULL(deliverymins,0) AS deliverymins FROM tenantlocations WHERE status='Active' AND locationid=? "
	getAllLocations                = "SELECT  locationid,locationname,tenantid,email,contactno,address,city,state,postcode,latitude,longitude,countrycode,opentime,closetime,createdby,status FROM tenantlocations WHERE status='Active' AND tenantid=? "
	createTenantUserQuery          = "INSERT INTO app_users (authname,password,hashsalt,contactno,roleid,referenceid) VALUES(?,?,?,?,?,?)"
	insertTenantUsertoProfileQuery = "INSERT INTO app_userprofiles (userid,firstname,lastname,email,contactno,userlocationid) VALUES(?,?,?,?,?,?)"
	getAllTenantUsers              = "SELECT a.firstname,a.lastname,a.userlocationid,a.userid,a.created,a.contactno,a.email,a.status,b.locationname,c.referenceid FROM app_userprofiles a, tenantlocations b, app_users c WHERE  a.userid=c.userid AND a.userlocationid=b.locationid AND c.referenceid=b.tenantid AND b.tenantid=?"
	updateTenantUser               = "UPDATE app_users a, app_userprofiles b  SET  a.authname=?,a.contactno=?,b.firstname=?,b.lastname=?,b.email=?,b.contactno=?,b.userlocationid=? WHERE a.userid=b.userid AND a.userid=?"
	getAllTenantUserByLocationId   = ""
	updateTenantBusiness           = "UPDATE tenants SET brandname=?,tenantinfo=?,paymode1=?,paymode2=?,tenantimage=? WHERE tenantid=?"
	insertSocialInfo               = "INSERT INTO tenantsocial (tenantid,socialprofile,sociallink,socialicon) VALUES"
	updatesocialinfo               = "UPDATE tenantsocial SET socialprofile=?, sociallink=?,socialicon=? WHERE tenantid=? AND socialid=? "
	updateauthuser                 = "UPDATE  app_users a,app_userprofiles b SET a.referenceid=?,b.userlocationid=? WHERE a.userid=b.userid AND a.userid=?"
	getBusinessbyid                = "SELECT tenantid,IFNULL(brandname,'') AS brandname,IFNULL(tenantinfo,'') AS tenantinfo,IFNULL(paymode1,0) AS paymode1,IFNULL(paymode2,0) AS paymode2,IFNULL(tenantaccid,0) AS tenantaccid,IFNULL(address,'') AS address,IFNULL(primaryemail,'') AS primaryemail,IFNULL(primarycontact,'') AS  primarycontact,IFNULL(tenanttoken,'') AS tenanttoken,IFNULL(tenantimage,'') AS tenantimage FROM tenants WHERE tenantid=?"
	getAllSocial                   = "SELECT socialid, IFNULL(socialprofile,'') AS socialprofile , IFNULL(sociallink,'') AS sociallink, IFNULL(socialicon,'') AS socialicon FROM tenantsocial WHERE tenantid= ?"
	userAuthentication             = "SELECT a.userid,b.firstname,b.lastname,b.email,b.contactno,b.status,b.created FROM app_users a, app_userprofiles b WHERE a.userid=b.userid AND a.status ='Active' AND a.userid=?"
	Getpromotions                  = "SELECT a.promotionid,a.promotiontypeid,a.tenantid,IFNULL(a.promoname,'') AS promoname,IFNULL(a.promocode,'') AS promocode,IFNULL(a.promoterms,'') AS promoterms,a.promovalue,a.startdate,a.enddate,a.status,b.typename,b.tag, c.tenantname FROM promotions a, promotiontypes b,tenants c WHERE a.promotiontypeid=b.promotiontypeid AND a.tenantid=c.tenantid AND a.`status`='Active' AND a.tenantid=?"
	createpromotion                = "INSERT INTO promotions (promotiontypeid,tenantid,promoname,promocode,promoterms,promovalue,startdate,enddate,createdby) VALUES(?,?,?,?,?,?,?,?,?)"
	insertsequence                 = "INSERT INTO ordersequence (tenantid,sequencename,seqno,prefix,subprefix) VALUES(?,?,?,?,?)"
	insertcharge                   = "INSERT INTO tenantcharges (tenantid,locationid,chargeid,chargename,chargetype,chargevalue,createdby) VALUES"
	insertdelivery                 = "INSERT INTO tenantsettings (tenantid,locationid,slabtype,slab,slablimit,slabcharge,createdby) VALUES"
	updatecharge                   = "UPDATE tenantcharges SET locationid=?,chargeid=?,chargename=?,chargetype=?, chargevalue=? WHERE tenantchargeid=? AND tenantid=?"
	updatedelivery                 = "UPDATE  tenantsettings SET locationid=?,slabtype=?,slab=?,slablimit=?,slabcharge=? WHERE settingsid=? AND tenantid=?"
	deletecharge                   = "DELETE FROM tenantcharges WHERE tenantchargeid=?"
	deletedelivery                 = "DELETE FROM  tenantsettings WHERE settingsid=?"
	updatelocationstatus           = "UPDATE tenantlocations SET status=? WHERE tenantid= ? AND locationid=?"
	updatedeliverystatus           = "UPDATE tenantlocations SET delivery=? WHERE tenantid= ? AND locationid=?"
	getCustomerByid                = "SELECT customerid,firstname,lastname,contactno,email,IFNULL(configid,0) AS configid  FROM customers WHERE customerid=?"
	getsubscription                = "SELECT a.packageid,a.moduleid,a.tenantid,a.totalamount,b.modulename,b.logourl,c.packagename,c.packageamount,c.packageicon, (SELECT COUNT(locationid)  FROM tenantlocations where  tenantid =?) AS location, (SELECT COUNT(tenantcustomerid)  FROM tenantcustomers WHERE tenantid =?) AS customer FROM tenantsubscription a , app_module b,app_package c   WHERE a.moduleid=b.moduleid AND a.packageid=c.packageid AND  a.tenantid=?"
	nonsubscribed                  = "SELECT a.packageid,a.moduleid,a.packagename,a.packageamount,a.paymentmode,a.packagecontent,a.packageicon,b.modulename,IFNULL(d.promocodeid,0) AS promocodeid,IFNULL(d.promoname ,'') AS promoname,IFNULL(d.promodescription,'') AS promodescription,IFNULL(d.promotype ,'') AS promotype,IFNULL(d.promovalue,0) AS promovalue,IFNULL(d.packageexpiry,0) AS packageexpiry,IFNULL(d.validity,'') AS validity,IF(d.validity>=DATE(NOW()), true, false) AS validity FROM app_package a Inner JOIN app_module b ON a.moduleid=b.moduleid INNER JOIN  app_promocodes d ON a.packageid=d.packageid WHERE a.`status`='Active' AND  a.packageid  NOT IN (SELECT packageid FROM tenantsubscription WHERE tenantid= ? )"
	getpayments                    = "SELECT a.paymentid,a.packageid,IFNULL(a.paymentref,'') AS paymentref,IFNULL(a.locationid,0) AS locationid,a.paymenttypeid,a.tenantid,IFNULL(a.customerid,0) AS customerid,a.transactiondate,IFNULL(a.orderid,0) AS orderid,a.chargeid,a.amount,a.refundamt,a.paymentstatus,a.created,IFNULL(b.packagename,'') AS  packagename,IFNULL(c.firstname,'') AS firstname,IFNULL(c.lastname,'')AS lastname,IFNULL(c.contactno,'')AS contactno,IFNULL(c.email,'')AS email FROM payments a LEFT OUTER JOIN  app_package b ON a.packageid=b.packageid LEFT OUTER JOIN customers c ON  a.customerid=c.customerid WHERE tenantid=? AND paymenttypeid=?"
	getbusinessforassist           = "SELECT a.tenantid,IFNULL(a.brandname,'') AS brandname,IFNULL(a.tenantinfo,'') AS tenantinfo,IFNULL(a.paymode1,0) AS paymode1,IFNULL(a.paymode2,0) AS paymode2,IFNULL(a.tenantaccid,0) AS tenantaccid,IFNULL(a.address,'') AS address,IFNULL(a.primaryemail,'') AS primaryemail,IFNULL(a.primarycontact,'') AS  primarycontact,IFNULL(a.tenanttoken,'') AS tenanttoken,IFNULL(tenantimage,'') AS tenantimage,IFNULL(b.moduleid,0) AS moduleid,IFNULL(d.modulename,'') AS modulename FROM tenants a, tenantsubscription b , app_category c, app_module d WHERE a.tenantid=b.tenantid AND b.moduleid=d.moduleid AND c.categoryid=d.categoryid AND a.tenantid=? AND  c.categoryid=?"
	updateinitial1                 = "UPDATE tenants SET brandname=?,tenantinfo=?,tenantimage=? WHERE tenantid=?"
	updateinitial2                 = "UPDATE tenantlocations SET opentime=?,closetime=? WHERE tenantid=? AND locationid=?"
	deletesocial                   = "DELETE FROM tenantsocial WHERE socialid=?"
	getmodules                     = "SELECT moduleid,categoryid,modulename,content,IFNULL(logourl,'') AS logourl,IFNULL(iconurl,'') AS iconurl FROM app_module WHERE STATUS='Active' AND categoryid=?"
	getpromo                       = "SELECT IFNULL(promocodeid,0) AS promocodeid,moduleid,partnerid,packageid, IFNULL(promoname,'') AS promoname, IFNULL(promodescription,'') AS promodescription, IFNULL(packageexpiry,0) AS packageexpiry, IFNULL(promotype,'') AS promotype, IFNULL(promovalue,0) AS promovalue, IFNULL(validity,'') AS validity,IF(validity>= DATE(NOW()), TRUE, FALSE) AS validitystatus FROM app_promocodes WHERE STATUS='Active' AND moduleid=?"
	insertsubcategory              = "INSERT INTO tenantsubcategories (tenantid,moduleid,categoryid,subcategoryid,subcategoryname) VALUES(?,?,?,?,?)"
)

func (s *Initialsubscriptiondata) Subscriptioninitial() (bool, *SubscribedData, error) {
	var tenantid int64
	var locationid int64
	var subcatid int64
	var err error
	var data SubscribedData
	tx, err := database.Db.Begin()
	if err != nil {
		return false, nil, err
	}

	{
		print("entry in tenants")
		stmt, err := tx.Prepare(insertTenantInfoQuery)
		if err != nil {
			tx.Rollback()
			return false, nil, err
		}
		defer stmt.Close()

		res, err := stmt.Exec(&s.Userid, &s.Partnerid, &s.Regno, &s.Name, &s.Email, &s.Mobile,
			&s.Address, &s.State, &s.Suburb, &s.Latitude, &s.Longitude, &s.Zip, &s.Countrycode,
			&s.TimeZone, &s.CurrencyCode, &s.Tenanttoken)
		if err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			return false, nil, err
		}
		tenantid, err = res.LastInsertId()
		print("tenantid=", tenantid)
	}

	{
		print("entry in location")
		stmt, err := tx.Prepare(insertTenantLocationQuery)
		if err != nil {
			tx.Rollback()
			return false, nil, err
		}
		defer stmt.Close()

		res, err := stmt.Exec(tenantid, &s.Name, &s.Email,
			&s.Mobile, &s.Address, &s.State,
			&s.Suburb, &s.Latitude, &s.Longitude,
			&s.Zip, &s.Countrycode, &s.OpenTime,
			&s.CloseTime, &s.Userid)
		if err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			return false, nil, err
		}
		locationid, err = res.LastInsertId()
		print("locationid=", locationid)
	}
	datalist := s.Tenantsubscribe
	if len(datalist) != 0 {
		var a TenantSubscription
		for i := 0; i < len(datalist); i++ {
			a.Currencyid = datalist[i].Currencyid
			a.Partnerid = datalist[i].Partnerid
			a.Date = datalist[i].Date
			a.Moduleid = datalist[i].Moduleid
			a.PaymentId = datalist[i].PaymentId
			a.PaymentStatus = datalist[i].PaymentStatus
			a.Price = datalist[i].Price
			a.Quantity = datalist[i].Quantity
			a.TaxId = datalist[i].TaxId
			a.TaxAmount = datalist[i].TaxAmount
			a.TotalAmount = datalist[i].TotalAmount
			a.Packageid = datalist[i].Packageid
			a.Promoid = datalist[i].Promoid
			a.Promovalue = datalist[i].Promovalue
			a.Validitydate = datalist[i].Validitydate
			a.Promostatus = true
			{
				print("entry in subscription")
				stmt, err := tx.Prepare(insertTenantSubscription)
				if err != nil {
					tx.Rollback()
					return false, nil, err
				}
				defer stmt.Close()
				if _, err := stmt.Exec(tenantid, &a.Date, &a.Packageid, &a.Partnerid, &a.Moduleid,
					&s.Categoryid, &s.SubCategoryid,
					&a.Currencyid, &a.Price, &a.Quantity, &a.TaxId, &a.TaxAmount,
					&a.TotalAmount, &a.PaymentStatus, &a.PaymentId, &a.Promoid,
					&a.Promovalue, &a.Promostatus, &a.Validitydate); err != nil {
					tx.Rollback() // return an error too, we may want to wrap them
					return false, nil, err
				}
			}
		}

	}
	{
		print("entry in subcat")
		stmt, err := tx.Prepare(insertsubcategory)
		if err != nil {
			tx.Rollback()
			return false, nil, err
		}
		defer stmt.Close()

		res, err := stmt.Exec(tenantid, &s.Tenantsubscribe[0].Moduleid, &s.Categoryid, &s.SubCategoryid, &s.Subcategoryname)
		if err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			return false, nil, err
		}
		subcatid, err = res.LastInsertId()
		print("subcat=", subcatid)
	}

	data.TenantID = int(tenantid)
	data.Locationid = int(locationid)

	tx.Commit()
	return true, &data, nil
}

func GetAllCategory() []Category {
	print("st1c")
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
	print("st1s")

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
	print("st1p")
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
		err := rows.Scan(&pack.PackageID, &pack.ModuleID, &pack.Name, &pack.PackageAmount, &pack.PaymentMode, &pack.PackageContent, &pack.PackageIcon, &pack.ModuleName,
			&pack.Promocodeid, &pack.Promoname, &pack.Promodescription, &pack.Packageexpiry, &pack.Promotype, &pack.Promovalue, &pack.Promovaliditydate, &pack.Validity)
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

func Getallnonsubscribedpackages(tenantid int) []Packages {
	print("st1p")
	stmt, err := database.Db.Prepare(nonsubscribed)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(tenantid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var packageList []Packages

	for rows.Next() {
		var pack Packages
		err := rows.Scan(&pack.PackageID, &pack.ModuleID, &pack.Name, &pack.PackageAmount, &pack.PaymentMode, &pack.PackageContent, &pack.PackageIcon, &pack.ModuleName,
			&pack.Promocodeid, &pack.Promoname, &pack.Promodescription, &pack.Promotype, &pack.Promovalue, &pack.Packageexpiry, &pack.Promovaliditydate, &pack.Validity)
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
		&info.Address.Address, &info.Address.State, &info.Address.Suburb, &info.Address.Latitude, &info.Address.Longitude, &info.Address.Zip, &info.Address.Countrycode, &info.Address.TimeZone, &info.Address.CurrencyCode, &info.Info.Tenanttoken)
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
	res, err := statement.Exec(tenantid, &info.Date, &info.Packageid, &info.Partnerid, &info.Moduleid, &info.Categoryid, &info.SubCategoryid, &info.Currencyid, &info.Price, &info.Quantity, &info.TaxId, &info.TaxAmount,
		&info.TotalAmount, &info.PaymentStatus, &info.PaymentId, &info.Promoid, &info.Promovalue, &info.Promostatus, &info.Validitydate)
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
func (info *SubscribedData) GetSubscribedData(tenantid int64) []SubscribedData {
	print("st1c")
	stmt, err := database.Db.Prepare(getSubscribedDataQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(tenantid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var paylist []SubscribedData

	for rows.Next() {
		var p SubscribedData
		err := rows.Scan(&p.TenantID, &p.TenantName, &p.ModuleID, &p.Subscriptionid, &p.Categoryid, &p.Subcategoryid, &p.ModuleName, &p.Locationid, &p.Locationname)
		if err != nil {
			log.Fatal(err)
		}
		paylist = append(paylist, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return paylist
}
func (info *SubscribedData) GetInitialSubscribedData(tenantid int64) *SubscribedData {
	print("st1c")
	stmt, err := database.Db.Prepare(getSubscribedDataQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(tenantid)
	// print(row)
	var p SubscribedData
	err = row.Scan(&p.TenantID, &p.TenantName, &p.ModuleID, &p.Subscriptionid, &p.ModuleName, &p.Locationid, &p.Locationname)
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

	return &p
}

func Payments(tenantid, typeid int) []Payment {
	print("st1c")
	stmt, err := database.Db.Prepare(getpayments)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(tenantid, typeid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var paylist []Payment

	for rows.Next() {
		var p Payment
		err := rows.Scan(&p.Paymentid, &p.Packageid, &p.Paymentref, &p.Locationid, &p.Paymenttypeid, &p.Tenantid, &p.Customerid,
			&p.Transactiondate, &p.Orderid, &p.Chargeid, &p.Amount, &p.Refundamt, &p.Paymentstatus, &p.Created, &p.Packagename, &p.Firstname,
			&p.Lastname, &p.Contactno, &p.Email)
		if err != nil {
			log.Fatal(err)
		}
		paylist = append(paylist, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return paylist

}
func (loco *Location) CreateLocation(id int64) (int64, error) {
	statement, err := database.Db.Prepare(createLocationQuery)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	res, err := statement.Exec(&loco.TenantID, &loco.LocationName, &loco.Email, &loco.Mobile, &loco.Address, &loco.State, &loco.Suburb, &loco.Latitude, &loco.Longitude, &loco.Zip, &loco.Countrycode, &loco.OpeningTime, &loco.ClosingTime, id, &loco.Delivery, &loco.Deliverytype, &loco.Deliverymins)
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
func (loco *Location) UpdateLocation() (bool, error) {
	statement, err := database.Db.Prepare(updatelocation)
	print(statement)

	if err != nil {

		log.Fatal(err)
		return false, err
	}
	defer statement.Close()
	_, err = statement.Exec(&loco.LocationName, &loco.Email, &loco.Mobile, &loco.Address, &loco.State, &loco.Suburb, &loco.Latitude, &loco.Longitude, &loco.Zip, &loco.Countrycode, &loco.OpeningTime, &loco.ClosingTime, &loco.Delivery, &loco.Deliverytype, &loco.Deliverymins, &loco.TenantID, &loco.LocationId)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	log.Print("Row updated in location!")
	return true, nil
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
	err = row.Scan(&data.LocationId, &data.LocationName, &data.Address, &data.Suburb, &data.State, &data.Zip, &data.Latitude, &data.Longitude, &data.Countrycode, &data.OpeningTime, &data.ClosingTime, &data.Createdby, &data.Status, &data.Delivery, &data.Deliverytype, &data.Deliverymins)
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

	var orders []Tenantlocation

	DB.Table("tenantlocations").Preload("Appuserprofiles").Preload("Tenantcharges").Preload("Tenantsettings").Where("tenantid=?", id).Find(&orders)

	// fmt.Println(orders)
	// for index, value := range orders {
	// 	fmt.Println(index, " = ", value)
	// log.Print(pretty.Sprint(tentlocation))
	return orders

}

func Locationbyid(tenantid, locationid int) *Tenantlocation {

	DB, err := gorm.Open(mysql.New(mysql.Config{Conn: dbconfig.Db}), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")

	} else {
		log.Println("Connection Established")
	}

	var data Tenantlocation

	DB.Table("tenantlocations").Preload("Appuserprofiles").Preload("Tenantcharges").Preload("Tenantsettings").Where("tenantid=? AND locationid=?", tenantid, locationid).Find(&data)

	fmt.Println(data)

	return &data

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
	_, err = statement.Exec(user.Brandname, user.About, user.Paymode1, user.Paymode2, user.Tenantimage, user.TenantID)
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
	_, err = statement.Exec(&info.SociaProfile, &info.SocialLink, &info.SocialIcon, tenantid, &info.Socialid)
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
	err = row.Scan(&data.TenantID, &data.Brandname, &data.About, &data.Paymode1, &data.Paymode2, &data.TenantaccId, &data.Address, &data.Email, &data.Phone, &data.Tenanttoken, &data.Tenantimage)
	print(err)
	fmt.Println("2")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
			return &data, false
		} else {
			// log.Fatal(err)
			return &data, false
		}
	}
	// fmt.Println(user)

	fmt.Println("completed")
	return &data, true
}
func (business *BusinessUpdate) GetBusinessforassist(id, catid int) (*BusinessUpdate, bool) {
	var data BusinessUpdate
	stmt, err := database.Db.Prepare(getbusinessforassist)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id, catid)
	// print(row)
	err = row.Scan(&data.TenantID, &data.Brandname, &data.About, &data.Paymode1,
		&data.Paymode2, &data.TenantaccId, &data.Address, &data.Email, &data.Phone,
		&data.Tenanttoken, &data.Tenantimage, &data.Moduleid, &data.Modulename)
	print(err)
	fmt.Println("2")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
			return &data, false
		} else {
			log.Fatal(err)
			return &data, false
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

func UserAuthentication(id int64) (*User, bool, error) {
	fmt.Println("enrty in sparkleauth")
	print(id)
	var data User
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
func Customerauthenticate(id int64) (*User, bool, error) {

	fmt.Println("enrty in customergetbyid")
	print(id)
	var data User
	stmt, err := database.Db.Prepare(getCustomerByid)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	// print(row)
	err = row.Scan(&data.ID, &data.FirstName, &data.LastName, &data.Mobile, &data.Email, &data.Configid)
	print(err)
	fmt.Println("2")
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
	data.From = "CUSTOMER"
	fmt.Println("completed")
	return &data, true, nil
}
func GetAllPromotions(tenantid int) []Promotion {
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
		err := rows.Scan(&promo.Promotionid, &promo.Promotiontypeid, &promo.Tenantid, &promo.Promoname, &promo.Promocode, &promo.Promoterms, &promo.Promovalue,
			&promo.Startdate, &promo.Enddate, &promo.Status, &promo.Promotype, &promo.Promotag, &promo.Tenantname)
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
func (p *Promotion) Createpromotion(created int) int64 {
	statement, err := database.Db.Prepare(createpromotion)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	res, err := statement.Exec(&p.Promotiontypeid, &p.Tenantid, &p.Promoname, &p.Promocode, &p.Promoterms,
		&p.Promovalue, &p.Startdate, &p.Enddate, created)
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

func Getmodules(catid, tenantid int) []*model.Mod {

	DB, err := gorm.Open(mysql.New(mysql.Config{Conn: dbconfig.Db}), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")

	} else {
		log.Println("Connection Established")
	}

	var data []*model.Mod
	if catid != 0 && tenantid == 0 {
		print("con1")
		DB.Table("app_module").Where("status='Active' and categoryid=?", catid).Find(&data)
	} else if catid == 0 && tenantid != 0 {
		print("con2")
		var q1 string
		n1 := int64(tenantid)
		tenant := strconv.FormatInt(n1, 10)
		q1 = " SELECT moduleid,categoryid,subcategoryid,subcategoryname,modulename,baseprice,taxpercent,taxamount,amount,IFNULL(content,'') AS content,IFNULL(logourl,'') AS logourl,IFNULL(iconurl,'') AS iconurl FROM app_module WHERE STATUS ='Active' and categoryid NOT IN (SELECT categoryid FROM tenantsubcategories WHERE tenantid= " + tenant + " )"
		DB.Raw(q1).Find(&data)
	} else {
		print("con3")

		DB.Table("app_module").Where("status='Active'").Find(&data)
	}

	fmt.Println(data)

	return data

}

func Getpromos(moduleid int) []*model.Promo {

	DB, err := gorm.Open(mysql.New(mysql.Config{Conn: dbconfig.Db}), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")

	} else {
		log.Println("Connection Established")
	}
	var q1 string
	n1 := int64(moduleid)
	mod := strconv.FormatInt(n1, 10)
	if moduleid != 0 {
		q1 = "SELECT IFNULL(a.promocodeid,0) AS promocodeid,a.moduleid,a.partnerid,a.packageid, IFNULL(a.promoname,'') AS promoname, IFNULL(a.promodescription,'') AS promodescription, IFNULL(a.packageexpiry,0) AS packageexpiry, IFNULL(a.promotype,'') AS promotype, IFNULL(a.promovalue,0) AS promovalue, IFNULL(a.validity,'') AS validity,IF(a.validity>= DATE(NOW()), TRUE, FALSE) AS validitystatus, b.companyname,b.address,b.city,b.postcode FROM app_promocodes a, partnerinfo b WHERE a.STATUS='Active' AND a.partnerid=b.partnerid AND a.moduleid=" + mod
	} else {
		q1 = "SELECT IFNULL(a.promocodeid,0) AS promocodeid,a.moduleid,a.partnerid,a.packageid, IFNULL(a.promoname,'') AS promoname, IFNULL(a.promodescription,'') AS promodescription, IFNULL(a.packageexpiry,0) AS packageexpiry, IFNULL(a.promotype,'') AS promotype, IFNULL(a.promovalue,0) AS promovalue, IFNULL(a.validity,'') AS validity,IF(a.validity>= DATE(NOW()), TRUE, FALSE) AS validitystatus, b.companyname,b.address,b.city,b.postcode FROM app_promocodes a, partnerinfo b WHERE a.STATUS='Active' AND a.partnerid=b.partnerid"
	}

	var data []*model.Promo

	DB.Raw(q1).Find(&data)

	fmt.Println(data)

	return data

}

func Getunsubscribecategory(tenantid int) []*model.Cat {
	DB, err := gorm.Open(mysql.New(mysql.Config{Conn: dbconfig.Db}), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")

	} else {
		log.Println("Connection Established")
	}
	var q1 string
	n1 := int64(tenantid)
	tent := strconv.FormatInt(n1, 10)
q1="SELECT categoryid,categoryname,categorytype,sortorder,`status` FROM app_category WHERE categoryid NOT IN(SELECT categoryid FROM tenantsubscription WHERE tenantid= "+tent+  " )"
	var data []*model.Cat

	DB.Raw(q1).Find(&data)

	fmt.Println(data)

	return data

}
func Getsubcatbyid(categoryid int) []*model.Subcat {

	DB, err := gorm.Open(mysql.New(mysql.Config{Conn: dbconfig.Db}), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")

	} else {
		log.Println("Connection Established")
	}

	var data []*model.Subcat

	DB.Table("app_subcategory").Where("status='Active' AND categoryid=?", categoryid).Find(&data)

	fmt.Println(data)

	return data

}
func Gettenantsubcat(moduleid, tenantid, categoryid int) []*model.Tenantsubcat {

	DB, err := gorm.Open(mysql.New(mysql.Config{Conn: dbconfig.Db}), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")

	} else {
		log.Println("Connection Established")
	}
	var q1 string
	var q2 string
	var q3 string
	n1 := int64(moduleid)
	n2 := int64(tenantid)
	n3 := int64(categoryid)

	mod := strconv.FormatInt(n1, 10)
	tent := strconv.FormatInt(n2, 10)
	cat := strconv.FormatInt(n3, 10)

	q1 = "SELECT categoryid,subcategoryid,subcategoryname,icon,0 AS selected FROM app_subcategory WHERE categoryid=" + cat + "  AND STATUS='Active' AND subcategoryid NOT IN"
	q2 = "(SELECT subcategoryid FROM tenantsubcategories WHERE tenantid=" + tent + "  AND moduleid="+ mod +" ) UNION"
	q3 = "  SELECT a.categoryid,a.subcategoryid,a.subcategoryname,IFNULL(b.icon,'') AS icon,1 AS selected FROM tenantsubcategories a, app_subcategory b WHERE a.subcategoryid=b.subcategoryid AND a.tenantid=" + tent + "  AND a.moduleid=" + mod + "  AND a.STATUS='Active'"
	var data []*model.Tenantsubcat

	DB.Raw(q1 + q2 + q3).Find(&data)

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
	res, err := statement.Exec(&o.Tenantid, &o.Tablename, &o.Seqno, &o.Prefix, &o.Subprefix)
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
func (o *Ordersequence) Insertpaysequence() (int64, error) {

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
	res, err := statement.Exec(&o.Tenantid, "payment", &o.Seqno, "REC", &o.Subprefix)
	if err != nil {
		log.Fatal(err)

	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())

	}
	log.Print("Row inserted in paysequence!")
	return id, nil
}
func (o *Ordersequence) Insertcustomersequence() (int64, error) {

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
	res, err := statement.Exec(&o.Tenantid, "customer", &o.Seqno, "CUS", &o.Subprefix)
	if err != nil {
		log.Fatal(err)

	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())

	}
	log.Print("Row inserted in CUSTOMERsequence!")
	return id, nil
}
func (info *Tenantcharge) Insertothercharges(soc []Tenantcharge) error {

	var inserts []string
	var params []interface{}
	for _, v := range soc {
		inserts = append(inserts, "(?, ?, ?, ?,?,?,?)")
		params = append(params, v.Tenantid, v.Locationid, v.Chargeid, v.Chargename, v.Chargetype, v.Chargevalue, v.Createdby)
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
func (info *Tenantsetting) Insertdeliverycharges(soc []Tenantsetting) error {

	var inserts []string
	var params []interface{}
	for _, v := range soc {
		inserts = append(inserts, "(?, ?, ?, ?,?,?,?)")
		params = append(params, v.Tenantid, v.Locationid, v.Slabtype, v.Slab, v.Slablimit, v.Slabcharge, v.Createdby)
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
func (o *Tenantcharge) Updateothercharge() bool {
	fmt.Println("0")
	statement, err := database.Db.Prepare(updatecharge)
	fmt.Println("1")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer statement.Close()
	_, err = statement.Exec(&o.Locationid, &o.Chargeid, &o.Chargename, &o.Chargetype, &o.Chargevalue, &o.Tenantchargeid, &o.Tenantid)
	if err != nil {

		log.Fatal(err)
		return false
	}

	log.Print("Row updated in charges!")
	return true
}
func (o *Tenantsetting) Updatedeliverycharge() bool {
	fmt.Println("0")
	statement, err := database.Db.Prepare(updatedelivery)
	fmt.Println("1")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer statement.Close()
	_, err = statement.Exec(&o.Locationid, &o.Slabtype, &o.Slab, &o.Slablimit, &o.Slabcharge, &o.Settingsid, &o.Tenantid)
	if err != nil {

		log.Fatal(err)
		return false
	}

	log.Print("Row updated in charges!")
	return true
}
func (c *Tenantcharge) Deleteothercharge() bool {

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
func (c *Tenantsetting) Deletedeliverycharge() bool {

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
func (u *Updatestatus) Updatelocationstatus() bool {
	statement, err := database.Db.Prepare(updatelocationstatus)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	_, err = statement.Exec(&u.Locationstatus, &u.Tenantid, &u.Locationid)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Row updated in tenant location!")
	return true
}
func (u *Updatestatus) Updatedeliverystatus() bool {
	statement, err := database.Db.Prepare(updatedeliverystatus)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	_, err = statement.Exec(&u.Deliverystatus, &u.Tenantid, &u.Locationid)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Row updated in tenant deliverystatus!")
	return true
}

func GetAllSubscription(tenantid int) []Subscribe {
	stmt, err := database.Db.Prepare(getsubscription)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(tenantid, tenantid, tenantid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var Subscribelist []Subscribe

	for rows.Next() {
		var s Subscribe
		err := rows.Scan(&s.Packageid, &s.Moduleid, &s.Tenantid, &s.Totalamount, &s.Modulename, &s.Logourl, &s.Packagename,
			&s.PackageAmount, &s.PackageIcon, &s.Locationcount, &s.Customercount)
		if err != nil {
			log.Fatal(err)
		}
		Subscribelist = append(Subscribelist, s)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return Subscribelist
}
func (c *Initial) Initialupdate() (bool, error) {

	tx, err := database.Db.Begin()
	if err != nil {
		return false, err
	}

	{
		stmt, err := tx.Prepare(updateinitial1)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(&c.Brandname, &c.About, &c.Tenantimage, &c.Tenantid); err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			return false, err
		}
	}

	{
		stmt, err := tx.Prepare(updateinitial2)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(&c.Opentime, &c.Closetime, &c.Tenantid, &c.Locationid); err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			return false, err
		}
	}
	tx.Commit()
	return true, nil
}

func (c *Social) Deletesocial() bool {

	statement, err := database.Db.Prepare(deletesocial)
	print(statement)

	if err != nil {
		log.Fatal(err)
		return false
	}

	_, err = statement.Exec(&c.Socialid)
	if err != nil {
		log.Fatal(err)
		return false
	}

	log.Print("Row deleted in social!")
	return true

}

func (d *TenantSubscription) Insertsubcategory() (int64, error) {

	fmt.Println("0")

	statement, err := database.Db.Prepare(insertsubcategory)
	print(statement)
	fmt.Println("1")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer statement.Close()

	fmt.Println("2")
	res, err := statement.Exec(&d.Tenantid, &d.Moduleid, &d.Categoryid, &d.SubCategoryid, &d.Subcategoryname)
	if err != nil {
		log.Fatal(err)

	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())

	}
	log.Print("Row inserted in subcat!")
	return id, nil
}
