package subscription

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/engajerest/auth/Models/users"
	"github.com/engajerest/auth/utils/Errors"
	"github.com/engajerest/auth/utils/dbconfig"
	database "github.com/engajerest/auth/utils/dbconfig"
	"github.com/engajerest/sparkle/graph/model"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/option"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	getAllCategoryQuery            = "SELECT categoryid,categoryname,categorytype,sortorder,status FROM app_category WHERE STATUS='Active'"
	getAllSubCategoryQuery         = "SELECT  subcategoryid,categorytypeid,categoryid,subcategoryname,status,icon FROM app_subcategory WHERE statuscode=1"
	getAllPackageQuery             = "SELECT a.packageid,a.moduleid,a.packagename,a.packageamount,a.paymentmode,a.packagecontent,a.packageicon,b.modulename,IFNULL(c.promocodeid,0) AS promocodeid,IFNULL(c.promoname ,'') AS promoname,IFNULL(c.promodescription,'') AS promodescription,IFNULL(c.packageexpiry,0) AS packageexpiry,IFNULL(c.promotype ,'') AS promotype,IFNULL(c.promovalue,0) AS promovalue,IFNULL(c.validity,'') AS validity,IF(c.validity>=DATE(NOW()), true, false) AS validity FROM app_package a Inner JOIN app_module b ON a.moduleid=b.moduleid LEFT OUTER JOIN  app_promocodes c ON a.packageid=c.packageid WHERE a.`status`='Active' "
	insertTenantInfoQuery          = "INSERT INTO tenants (createdby,partnerid,registrationno,tenantname,primaryemail,primarycontact,Address,state,city,suburb,latitude,longitude,postcode,countrycode,timezone,currencyid,currencycode,currencysymbol,tenanttoken) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	insertTenantLocationQuery      = "INSERT INTO tenantlocations (tenantid,locationname,email,contactno,address,state,city,suburb,latitude,longitude,postcode,countrycode,opentime,closetime,deliverymins,createdby) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	insertTenantSubscription       = "INSERT INTO tenantsubscription (tenantid,transactiondate,packageid,partnerid,moduleid,featureid,categoryid,subcategoryid,currencyid,subscriptionprice,quantity,taxid,taxamount,totalamount,paymentstatus,paymentid,promoid,promoprice,promostatus,validitydate) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	Updatesubscriptionpay          = "UPDATE  tenantsubscription SET transactiondate=?,featureid=?,partnerid=?,currencyid=?,subscriptionprice=?,quantity=?,taxid=?,taxamount=?,totalamount=?,paymentstatus=?,paymentid=?,promoid=?,promoprice=?,promostatus=?,validitydate=?  WHERE subscriptionid=? "
	getSubscribedDataQuery         = "SELECT a.tenantid, a.tenantname,IFNULL(a.tenantaccid,'') AS tenantaccid, b.moduleid,b.featureid,b.subscriptionid,b.categoryid,b.subcategoryid, b.taxamount,b.totalamount,b.status, c.modulename ,d.locationid,d.locationname FROM tenants a,tenantsubscription b,app_module c ,tenantlocations d WHERE a.tenantid=b.tenantid AND b.moduleid=c.moduleid AND a.tenantid = d.tenantid AND b.status='Active' AND  a.tenantid=?"
	createLocationQuery            = "INSERT INTO tenantlocations (tenantid,locationname,email,contactno,address,state,city,suburb,latitude,longitude,postcode,countrycode,opentime,closetime,createdby,delivery,deliverytype,deliverymins) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	updatelocation                 = "UPDATE tenantlocations SET locationname=?,email=?,contactno=?,address=?,state=?,city=?,suburb=?,latitude=?,longitude=?,postcode=?,countrycode=?,opentime=?,closetime=?,delivery=?,deliverytype=?,deliverymins=? WHERE tenantid=? AND locationid=?"
	getLocationbyid                = "SELECT  locationid,locationname,address,IFNULL(suburb,'') AS suburb,city,state,postcode,latitude,longitude,countrycode,opentime,closetime,createdby,status,IFNULL(delivery,false) AS delivery,IFNULL(deliverytype,'') AS deliverytype,IFNULL(deliverymins,0) AS deliverymins FROM tenantlocations WHERE status='Active' AND locationid=? "
	getAllLocations                = "SELECT  locationid,locationname,tenantid,email,contactno,address,IFNULL(suburb,'') AS suburb,city,state,postcode,latitude,longitude,countrycode,opentime,closetime,createdby,status FROM tenantlocations WHERE status='Active' AND tenantid=? "
	createTenantUserQuery          = "INSERT INTO app_users (authname,password,hashsalt,contactno,roleid,referenceid) VALUES(?,?,?,?,?,?)"
	insertTenantUsertoProfileQuery = "INSERT INTO app_userprofiles (userid,firstname,lastname,email,contactno,profileimage,userlocationid) VALUES(?,?,?,?,?,?,?)"
	insertTenantstaff              = "INSERT INTO tenantstaffs (tenantid,moduleid,userid) VALUES(?,?,?)"
	insertTenantstaffdetails       = "INSERT INTO tenantstaffdetails (tenantstaffid,tenantid,locationid) VALUES(?,?,?)"
	checktenantstaffid             = "SELECT IFNULL(tenantstaffid,0) AS tenantstaffid FROM tenantstaffs WHERE tenantid=? AND moduleid=? AND  userid=?"
	checkfordeletestaff            = "SELECT IFNULL(staffdetailid,0) AS staffdetailid FROM tenantstaffdetails WHERE tenantstaffid=?"
	deletetenantstaff              = "DELETE FROM tenantstaffs WHERE tenantstaffid=?"
	deletetenantstaffdetails       = "DELETE FROM tenantstaffdetails WHERE staffdetailid=?"
	getAllTenantUsers              = "SELECT a.firstname,a.lastname,a.userlocationid,a.userid,a.created,a.contactno,a.email,a.status,b.locationname,c.referenceid FROM app_userprofiles a, tenantlocations b, app_users c WHERE  a.userid=c.userid AND a.userlocationid=b.locationid AND c.referenceid=b.tenantid AND b.tenantid=?"
	updateTenantUser               = "UPDATE app_users a, app_userprofiles b  SET  a.authname=?,a.contactno=?,b.firstname=?,b.lastname=?,b.email=?,b.contactno=?,b.profileimage=?,b.userlocationid=? WHERE a.userid=b.userid AND a.userid=?"
	getAllTenantUserByLocationId   = ""
	updateTenantBusiness           = "UPDATE tenants SET brandname=?,tenantinfo=?,paymode1=?,paymode2=?,tenantimage=? WHERE tenantid=?"
	insertSocialInfo               = "INSERT INTO tenantsocial (tenantid,socialprofile,dailcode,sociallink,socialicon) VALUES"
	updatesocialinfo               = "UPDATE tenantsocial SET socialprofile=?,dailcode=?, sociallink=?,socialicon=? WHERE tenantid=? AND socialid=? "
	updateauthuser                 = "UPDATE  app_users a,app_userprofiles b SET a.referenceid=?,b.userlocationid=? WHERE a.userid=b.userid AND a.userid=?"
	getBusinessbyid                = "SELECT tenantid,IFNULL(brandname,'') AS brandname,IFNULL(tenantinfo,'') AS tenantinfo,IFNULL(paymode1,0) AS paymode1,IFNULL(paymode2,0) AS paymode2,IFNULL(tenantaccid,0) AS tenantaccid,IFNULL(address,'') AS address,IFNULL(primaryemail,'') AS primaryemail,IFNULL(primarycontact,'') AS  primarycontact,IFNULL(tenanttoken,'') AS tenanttoken,IFNULL(tenantimage,'') AS tenantimage, IFNULL(countrycode,'') AS countrycode,IFNULL(currencycode,'') AS currencycode,IFNULL(currencysymbol,'') AS currencysymbol FROM tenants WHERE tenantid=?"
	getAllSocial                   = "SELECT socialid, IFNULL(socialprofile,'') AS socialprofile ,IFNULL(dailcode,'') AS dailcode, IFNULL(sociallink,'') AS sociallink, IFNULL(socialicon,'') AS socialicon FROM tenantsocial WHERE tenantid= ?"
	userAuthentication             = "SELECT a.userid,b.firstname,b.lastname,b.email,b.contactno,b.status,b.created FROM app_users a, app_userprofiles b WHERE a.userid=b.userid AND a.status ='Active' AND a.userid=?"
	Getpromotions                  = "SELECT a.promotionid,a.promotiontypeid,a.tenantid,IFNULL(a.promoname,'') AS promoname,IFNULL(a.promocode,'') AS promocode,IFNULL(a.promoterms,'') AS promoterms,a.promovalue,a.startdate,a.enddate,IFNULL(a.broadcaststatus,0) as broadcaststatus,IFNULL(a.success,0) as success,IFNULL(a.failure,0) as failure,a.status,b.typename,b.tag, c.tenantname FROM promotions a, promotiontypes b,tenants c WHERE a.promotiontypeid=b.promotiontypeid AND a.tenantid=c.tenantid AND a.`status`='Active' AND a.tenantid=?"
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
	getsubscription                = "SELECT a.subscriptionid,a.packageid,a.moduleid,a.featureid,a.tenantid,a.categoryid,a.subcategoryid,IFNULL(a.validitydate,'') AS validitydate,IF(a.validitydate>= DATE(NOW()), TRUE, FALSE) AS validity,a.totalamount,a.taxamount,ifnull(a.subscriptionaccid,'') as subscriptionaccid,ifnull(a.subscriptionmethodid,'') as subscriptionmethodid, a.paymentstatus,a.status, b.modulename,b.logourl,b.iconurl,'' as packagename,'0.0' AS packageamount,'' as packageicon,IFNULL(c.tenantaccid,'') AS tenantaccid, (SELECT COUNT(locationid)  FROM tenantlocations where  tenantid =?) AS location, (SELECT COUNT(tenantcustomerid)  FROM tenantcustomers WHERE tenantid =?) AS customer FROM tenantsubscription a , app_module b,tenants c WHERE a.moduleid=b.moduleid AND a.tenantid=c.tenantid AND a.status='Active' AND a.tenantid=?"
	nonsubscribed                  = "SELECT a.packageid,a.moduleid,a.packagename,a.packageamount,a.paymentmode,a.packagecontent,a.packageicon,b.modulename,IFNULL(d.promocodeid,0) AS promocodeid,IFNULL(d.promoname ,'') AS promoname,IFNULL(d.promodescription,'') AS promodescription,IFNULL(d.promotype ,'') AS promotype,IFNULL(d.promovalue,0) AS promovalue,IFNULL(d.packageexpiry,0) AS packageexpiry,IFNULL(d.validity,'') AS validity,IF(d.validity>=DATE(NOW()), true, false) AS validity FROM app_package a Inner JOIN app_module b ON a.moduleid=b.moduleid INNER JOIN  app_promocodes d ON a.packageid=d.packageid WHERE a.`status`='Active' AND  a.packageid  NOT IN (SELECT packageid FROM tenantsubscription WHERE tenantid= ? )"
	getpayments                    = "SELECT a.paymentid,a.packageid,IFNULL(a.paymentref,'') AS paymentref,IFNULL(a.locationid,0) AS locationid,a.paymenttypeid,a.tenantid,IFNULL(a.customerid,0) AS customerid,a.transactiondate,IFNULL(a.orderid,0) AS orderid,a.chargeid,a.amount,a.refundamt,a.paymentstatus,a.created,IFNULL(b.packagename,'') AS  packagename,IFNULL(c.firstname,'') AS firstname,IFNULL(c.lastname,'')AS lastname,IFNULL(c.contactno,'')AS contactno,IFNULL(c.email,'')AS email FROM payments a LEFT OUTER JOIN  app_package b ON a.packageid=b.packageid LEFT OUTER JOIN customers c ON  a.customerid=c.customerid WHERE tenantid=? AND paymenttypeid=?"
	getbusinessforassist           = "SELECT a.tenantid,IFNULL(a.brandname,'') AS brandname,IFNULL(a.tenantinfo,'') AS tenantinfo,IFNULL(a.paymode1,0) AS paymode1,IFNULL(a.paymode2,0) AS paymode2,IFNULL(a.tenantaccid,0) AS tenantaccid,IFNULL(a.address,'') AS address,IFNULL(a.primaryemail,'') AS primaryemail,IFNULL(a.primarycontact,'') AS  primarycontact,IFNULL(a.tenanttoken,'') AS tenanttoken,IFNULL(a.tenantimage,'') AS tenantimage, IFNULL(a.countrycode,'') AS countrycode,IFNULL(a.currencycode,'') AS currencycode,IFNULL(a.currencysymbol,'') AS currencysymbol,IFNULL(b.moduleid,0) AS moduleid,IFNULL(d.modulename,'') AS modulename FROM tenants a, tenantsubscription b , app_category c, app_module d WHERE a.tenantid=b.tenantid AND b.moduleid=d.moduleid AND c.categoryid=d.categoryid AND a.tenantid=? AND  c.categoryid=?"
	updateinitial1                 = "UPDATE tenants SET brandname=?,tenantinfo=?,tenantimage=? WHERE tenantid=?"
	updateinitial2                 = "UPDATE tenantlocations SET opentime=?,closetime=? WHERE tenantid=? AND locationid=?"
	deletesocial                   = "DELETE FROM tenantsocial WHERE socialid=?"
	getmodules                     = "SELECT moduleid,categoryid,modulename,content,IFNULL(logourl,'') AS logourl,IFNULL(iconurl,'') AS iconurl FROM app_module WHERE STATUS='Active' AND categoryid=?"
	getpromo                       = "SELECT IFNULL(promocodeid,0) AS promocodeid,moduleid,partnerid,packageid, IFNULL(promoname,'') AS promoname, IFNULL(promodescription,'') AS promodescription, IFNULL(packageexpiry,0) AS packageexpiry, IFNULL(promotype,'') AS promotype, IFNULL(promovalue,0) AS promovalue, IFNULL(validity,'') AS validity,IF(validity>= DATE(NOW()), TRUE, FALSE) AS validitystatus FROM app_promocodes WHERE STATUS='Active' AND moduleid=?"
	insertsubcategory              = "INSERT INTO tenantsubcategories (tenantid,moduleid,categoryid,subcategoryid,subcategoryname) VALUES(?,?,?,?,?)"
	createUsernopassword           = "INSERT INTO app_users (authname,contactno,roleid,configid,referenceid) VALUES(?,?,?,?,?)"
    unsubscribe = "UPDATE tenantsubscription SET categoryid=0, status='Inactive' WHERE subscriptionid=?"
	deletecategories = "DELETE  FROM tenantsubcategories WHERE tenantid=? AND moduleid=?"
	deletesubcatbyid = "DELETE  FROM tenantsubcategories WHERE tenantsubcatid=?"
	//firestore
	firestorejsonkey = "./engaje-2021-firebase-adminsdk-7sb61-42247472ad.json"
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
			&s.Address, &s.State, &s.City, &s.Suburb, &s.Latitude, &s.Longitude, &s.Zip, &s.Countrycode,
			&s.TimeZone, &s.Currencyid, &s.CurrencyCode, &s.Currencysymbol, &s.Tenanttoken)
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
			&s.City, s.Suburb, &s.Latitude, &s.Longitude,
			&s.Zip, &s.Countrycode, &s.OpenTime,
			&s.CloseTime, 30, &s.Userid)
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
			a.Categoryid = datalist[i].Categoryid
			a.SubCategoryid = datalist[i].SubCategoryid
			a.Subcategoryname = datalist[i].Subcategoryname
			a.Currencyid = datalist[i].Currencyid
			a.Partnerid = datalist[i].Partnerid
			a.Date = datalist[i].Date
			a.Moduleid = datalist[i].Moduleid
			a.Featureid = datalist[i].Featureid
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
				if _, err := stmt.Exec(tenantid, &a.Date, &a.Packageid, &a.Partnerid, &a.Moduleid, &a.Featureid,
					&a.Categoryid, &a.SubCategoryid,
					&a.Currencyid, &a.Price, &a.Quantity, &a.TaxId, &a.TaxAmount,
					&a.TotalAmount, &a.PaymentStatus, &a.PaymentId, &a.Promoid,
					&a.Promovalue, &a.Promostatus, &a.Validitydate); err != nil {
					tx.Rollback() // return an error too, we may want to wrap them
					return false, nil, err
				}
			}
		}

	}
	datalist1 := s.Tenantsubscribe

	if len(datalist1) != 0 {
		var d TenantSubscription
		for i := 0; i < len(datalist1); i++ {
			d.Categoryid = datalist1[i].Categoryid
			d.SubCategoryid = datalist1[i].SubCategoryid
			d.Subcategoryname = datalist1[i].Subcategoryname
			d.Moduleid = datalist1[i].Moduleid
			print("entry in subcat")
			stmt, err := tx.Prepare(insertsubcategory)
			if err != nil {
				tx.Rollback()
				return false, nil, err
			}
			defer stmt.Close()

			res, err := stmt.Exec(tenantid, &d.Moduleid, &d.Categoryid, &d.SubCategoryid, &d.Subcategoryname)
			if err != nil {
				tx.Rollback() // return an error too, we may want to wrap them
				return false, nil, err
			}
			subcatid, err = res.LastInsertId()
			print("subcat=", subcatid)

		}

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
	res, err := statement.Exec(tenantid, &info.Date, &info.Packageid, &info.Partnerid, &info.Moduleid, &info.Featureid, &info.Categoryid, &info.SubCategoryid, &info.Currencyid, &info.Price, &info.Quantity, &info.TaxId, &info.TaxAmount,
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

func (info *TenantSubscription) Updatesubscription() (bool, error) {
	statement, err := database.Db.Prepare(Updatesubscriptionpay)
	print(statement)

	if err != nil {

		log.Fatal(err)
		return false, err
	}
	defer statement.Close()
	_, err = statement.Exec(&info.Date, &info.Featureid, &info.Partnerid, &info.Currencyid, &info.Price, &info.Quantity, &info.TaxId, &info.TaxAmount,
		&info.TotalAmount, &info.PaymentStatus, &info.PaymentId, &info.Promoid, &info.Promovalue, &info.Promostatus, &info.Validitydate, &info.Subscriptionid)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	log.Print("Row updated in subscriptionpayment!")
	return true, nil
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
		err := rows.Scan(&p.TenantID, &p.TenantName, &p.Tenantaccid, &p.ModuleID, &p.Featureid, &p.Subscriptionid, &p.Categoryid, &p.Subcategoryid, &p.Taxamount, &p.Totalamount,&p.Status, &p.ModuleName, &p.Locationid, &p.Locationname)
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
	err = row.Scan(&p.TenantID, &p.TenantName, &p.Tenantaccid, &p.ModuleID, &p.Featureid, &p.Subscriptionid, &p.ModuleName, &p.Locationid, &p.Locationname)
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
	DB, err := gorm.Open(mysql.New(mysql.Config{Conn: dbconfig.Db}), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")

	} else {
		log.Println("Connection Established")
	}

	var data []Payment

	DB.Table("payments").
		Preload("Paymentdetails").Preload("Paymentdetails.Customers").Where("paymenttypeid=? AND tenantid=?", typeid, tenantid).Find(&data)
	for index, value := range data {
		fmt.Println(index, " = ", value)
	}

	return data

}
func (loco *Location) CreateLocation(id int64) (int64, error) {
	statement, err := database.Db.Prepare(createLocationQuery)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	res, err := statement.Exec(&loco.TenantID, &loco.LocationName, &loco.Email, &loco.Mobile, &loco.Address, &loco.State, &loco.City, &loco.Suburb, &loco.Latitude, &loco.Longitude, &loco.Zip, &loco.Countrycode, &loco.OpeningTime, &loco.ClosingTime, id, &loco.Delivery, &loco.Deliverytype, &loco.Deliverymins)
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
	_, err = statement.Exec(&loco.LocationName, &loco.Email, &loco.Mobile, &loco.Address, &loco.State, &loco.City, &loco.Suburb, &loco.Latitude, &loco.Longitude, &loco.Zip, &loco.Countrycode, &loco.OpeningTime, &loco.ClosingTime, &loco.Delivery, &loco.Deliverytype, &loco.Deliverymins, &loco.TenantID, &loco.LocationId)
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
	err = row.Scan(&data.LocationId, &data.LocationName, &data.Address, &data.Suburb, &data.City, &data.State, &data.Zip, &data.Latitude, &data.Longitude, &data.Countrycode, &data.OpeningTime, &data.ClosingTime, &data.Createdby, &data.Status, &data.Delivery, &data.Deliverytype, &data.Deliverymins)
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

	var data []Tenantlocation

	DB.Table("tenantlocations").
		Preload("Appuserprofiles").Preload("Tenantcharges").Preload("Tenantsettings").Where("tenantid=?", id).Find(&data)
	for index, value := range data {
		fmt.Println(index, " = ", value)
	}

	return data

}

func Locationbyid(tenantid, locationid int) *Tenantlocation {

	DB, err := gorm.Open(mysql.New(mysql.Config{Conn: dbconfig.Db}), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")

	} else {
		log.Println("Connection Established")
	}

	var data Tenantlocation

	DB.Table("tenantlocations").
		Preload("Appuserprofiles").Preload("Tenantcharges").Preload("Tenantsettings").Where("tenantid=? AND locationid=?", tenantid, locationid).Find(&data)

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
	res, err := statement.Exec(id, &user.FirstName, &user.LastName, &user.Email, &user.Mobile, &user.Profileimage, &user.Locationid)
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
func (user *TenantUser) InsertTenantstaffs() int64 {
	statement, err := database.Db.Prepare(insertTenantstaff)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	res, err := statement.Exec(&user.Tenantid, &user.Moduleid, &user.Locationid, &user.Userid)
	if err != nil {
		log.Fatal(err)
	}
	id, err1 := res.LastInsertId()
	if err1 != nil {
		log.Fatal("Error:", err1.Error())
	}
	log.Print("Row inserted in tenantstaffs!")
	return id
}
func (user *TenantUser) InsertTenantstaffdetails() int64 {
	statement, err := database.Db.Prepare(insertTenantstaffdetails)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	res, err := statement.Exec(&user.Tenantstaffid,
		&user.Tenantid, &user.Locationid)
	if err != nil {
		log.Fatal(err)
	}
	id, err1 := res.LastInsertId()
	if err1 != nil {
		log.Fatal("Error:", err1.Error())
	}
	log.Print("Row inserted in tenantstaffdetails!")
	return id
}
func (c *TenantUser) Deletetenantstaff() bool {

	statement, err := database.Db.Prepare(deletetenantstaff)
	print(statement)

	if err != nil {
		log.Fatal(err)
		return false
	}

	_, err = statement.Exec(&c.Tenantstaffid)
	if err != nil {
		log.Fatal(err)
		return false
	}

	log.Print("Row deleted in tenantstaffs!")
	return true

}
func (c *TenantUser) Deletetenantstaffdetails() bool {

	statement, err := database.Db.Prepare(deletetenantstaffdetails)
	print(statement)

	if err != nil {
		log.Fatal(err)
		return false
	}

	_, err = statement.Exec(&c.Staffdetailid)
	if err != nil {
		log.Fatal(err)
		return false
	}

	log.Print("Row deleted in tenantstaffs!")
	return true

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
func (user *TenantUser) CreateTenantUser() (int64, error) {
	fmt.Println("nopass")
	statement, err := database.Db.Prepare(createUsernopassword)
	fmt.Println("1")
	if err != nil {
		print(err)
		return 0, err

	}
	defer statement.Close()

	fmt.Println("2")

	res, err1 := statement.Exec(&user.Email, &user.Mobile, &user.Roleid, &user.Configid, &user.Tenantid)
	if err1 != nil {

		fmt.Println(err1)

		return 0, err1

	}
	id, err2 := res.LastInsertId()
	if err2 != nil {
		log.Fatal("Error:", err2.Error())
		return 0, err
	}
	log.Print("Row inserted in tenantuser!")
	return id, nil
}
func (user *TenantUser) UpdateTenantUser() (bool, error) {
	statement, err := database.Db.Prepare(updateTenantUser)
	print(statement)

	if err != nil {
		print(err)
		return false, err

	}
	defer statement.Close()
	_, err1 := statement.Exec(user.Email, user.Mobile, user.FirstName, user.LastName, user.Email, user.Mobile, user.Profileimage, user.Locationid, user.Userid)
	if err1 != nil {

		fmt.Println(err1)

		return false, err1

	}
	log.Print("Row updated in tenant user profile!")
	return true, nil
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
		inserts = append(inserts, "(?, ?, ?, ?,?)")
		params = append(params, id, v.SociaProfile, v.Dailcode, v.SocialLink, v.SocialIcon)
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
	_, err = statement.Exec(&info.SociaProfile, &info.Dailcode, &info.SocialLink, &info.SocialIcon, tenantid, &info.Socialid)
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
	err = row.Scan(&data.TenantID, &data.Brandname, &data.About, &data.Paymode1, &data.Paymode2, &data.TenantaccId, &data.Address, &data.Email, &data.Phone, &data.Tenanttoken, &data.Tenantimage,
		&data.Countrycode, &data.Currencycode, &data.Currencysymbol)
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
		&data.Tenanttoken, &data.Tenantimage, &data.Countrycode, &data.Currencycode, &data.Currencysymbol, &data.Moduleid, &data.Modulename)
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
			&data.SociaProfile, &data.Dailcode,
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
			&promo.Startdate, &promo.Enddate, &promo.Broadcaststatus, &promo.Success, &promo.Failure, &promo.Status, &promo.Promotype, &promo.Promotag, &promo.Tenantname)
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

func Getmodules(catid, tenantid int, mode bool) []*model.Mod {

	DB, err := gorm.Open(mysql.New(mysql.Config{Conn: dbconfig.Db}), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")

	} else {
		log.Println("Connection Established")
	}

	var data []*model.Mod
	if catid != 0 && tenantid == 0 && mode == false {
		print("con1")
		DB.Table("app_module").Where("status='Active' and categoryid=?", catid).Find(&data)
	} else if catid == 0 && tenantid != 0 && mode == false {
		print("con2")
		var q1 string
		n1 := int64(tenantid)
		tenant := strconv.FormatInt(n1, 10)
		q1 = " SELECT moduleid,categoryid,subcategoryid,subcategoryname,modulename,baseprice,taxpercent,taxamount,amount,IFNULL(content,'') AS content,IFNULL(logourl,'') AS logourl,IFNULL(iconurl,'') AS iconurl FROM app_module WHERE STATUS ='Active' and categoryid NOT IN (SELECT categoryid FROM tenantsubcategories WHERE tenantid= " + tenant + " )"
		DB.Raw(q1).Find(&data)
	} else if catid != 0 && tenantid == 0 && mode == true {
		print("con3")
		n1 := int64(catid)
		cat := strconv.FormatInt(n1, 10)
		DB.Table("app_module").Where("status='Active' and categoryid NOT IN (" + cat + ")").Find(&data)
	} else {
		print("con4")

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
	q1 = "SELECT categoryid,categoryname,categorytype,sortorder,`status` FROM app_category WHERE categoryid NOT IN(SELECT categoryid FROM tenantsubscription WHERE tenantid= " + tent + " )"
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

	q1 = "SELECT 0 AS tenantsubcatid,a.categoryid,a.subcategoryid,a.subcategoryname,a.icon,0 AS selected ,b.categoryname FROM app_subcategory a , app_category b WHERE a.categoryid=b.categoryid AND  a.categoryid= " + cat + "  AND a.STATUS='Active' AND a.subcategoryid NOT IN"
	q2 = "(SELECT subcategoryid FROM tenantsubcategories WHERE tenantid=" + tent + "  AND moduleid=" + mod + " ) UNION"
	q3 = "  SELECT a.tenantsubcatid,a.categoryid,a.subcategoryid,a.subcategoryname,IFNULL(b.icon,'') AS icon,1 AS selected,c.categoryname  FROM tenantsubcategories a, app_subcategory b,app_category c WHERE a.subcategoryid=b.subcategoryid AND a.categoryid=c.categoryid AND a.tenantid=" + tent + "  AND a.moduleid=" + mod + "  AND a.STATUS='Active'"
	var data []*model.Tenantsubcat

	DB.Raw(q1 + q2 + q3).Find(&data)
print(q1+q2+q3)


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
func (o *Ordersequence) Insertappointmentsequence() (int64, error) {

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
	res, err := statement.Exec(&o.Tenantid, "appointment", &o.Seqno, "APP", &o.Subprefix)
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
func  Unsubscribe(id int) (bool,error) {

	statement, err := database.Db.Prepare(unsubscribe)
	print(statement)

	if err != nil {
	
		return false,err
		
	}
	defer statement.Close()
	_, err1 := statement.Exec(id)
	if err1 != nil {
		
		return false,err1
		
	}

	log.Print("Row updated in tenant subscription!")
	return true,nil
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
		err := rows.Scan(&s.Subscriptionid, &s.Packageid, &s.Moduleid, &s.Featureid, &s.Tenantid, &s.Categoryid, &s.Subcategoryid, &s.Validitydate, &s.Validity, &s.Totalamount, &s.Taxamount, &s.Subscriptionaccid, &s.Subscriptionmethodid, &s.Paymentstatus,&s.Status, &s.Modulename, &s.Logourl, &s.Iconurl, &s.Packagename,
			&s.PackageAmount, &s.PackageIcon, &s.Tenantaccid, &s.Locationcount, &s.Customercount)
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
func  Deletesubcat(tenantid,moduleid int) (bool,error) {

	statement, err := database.Db.Prepare(deletecategories)
	

	if err != nil {
	
		return false,err
	}

	_, err1 := statement.Exec(tenantid,moduleid)
	if err1 != nil {
		
		return false,err1
	}

	log.Print("Row deleted in subcategories!")
	return true,nil

}
func  Deletesubcatbyid(tenantsubcatid int) (bool,error) {

	statement, err := database.Db.Prepare(deletesubcatbyid)
	

	if err != nil {
	
		return false,err
	}

	_, err1 := statement.Exec(tenantsubcatid)
	if err1 != nil {
		
		return false,err1
	}

	log.Print("Row deleted in subcategories by id!")
	return true,nil

}
func Gettenantinfo(tenantid int) Tenants {
	DB, err := gorm.Open(mysql.New(mysql.Config{Conn: dbconfig.Db}), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")

	} else {
		log.Println("Connection Established")
	}
	var data Tenants
	DB.Table("tenants").Preload("Tenantsubscriptions").Preload("Tenantlocations").Preload("Tenantlocations.Appuserprofiles").Preload("Tenantlocations.Tenantcharges").Preload("Tenantlocations.Tenantsettings").Preload("Tenantsubcategories").Where("tenantid=?", tenantid).Find(&data)
	fmt.Println(data)
	return data
}
func Gettenantusers(tenantid, userid int) []TenantUsers {

	n := int64(tenantid)
	tent := strconv.FormatInt(n, 10)
	var user string
	if userid != 0 {
		n1 := int64(userid)
		user = strconv.FormatInt(n1, 10)
	}
	var q1 string

	if userid != 0 {
		q1 = "SELECT a.firstname,a.lastname,IFNULL(a.profileimage,'') AS profileimage, a.userlocationid,a.userid,a.created,a.contactno,a.email,a.status,b.locationname,c.referenceid FROM app_userprofiles a, tenantlocations b, app_users c WHERE a.userid=c.userid AND a.userlocationid=b.locationid AND c.referenceid=b.tenantid AND b.tenantid= " + tent + "  AND a.userid=" + user
	} else {
		q1 = "SELECT a.firstname,a.lastname,IFNULL(a.profileimage,'') AS profileimage, a.userlocationid,a.userid,a.created,a.contactno,a.email,a.status,b.locationname,c.referenceid FROM app_userprofiles a, tenantlocations b, app_users c WHERE a.userid=c.userid AND a.userlocationid=b.locationid AND c.referenceid=b.tenantid AND b.tenantid=" + tent
	}

	DB, err := gorm.Open(mysql.New(mysql.Config{Conn: dbconfig.Db}), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")

	} else {
		log.Println("Connection Established")
	}

	var data []TenantUsers

	DB.Raw(q1).Find(&data)
	for index, value := range data {
		fmt.Println(index, " = ", value)
	}
	return data
}
func Checkstaffdata(tenantid, moduleid, userid int) (int, error) {
	var staffid int
	fmt.Println("enrty in staffs")

	stmt, err := database.Db.Prepare(checktenantstaffid)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(tenantid, moduleid, userid)
	// print(row)
	err = row.Scan(&staffid)
	print(err)
	fmt.Println("2")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
			data1 := Errors.RestError{}
			data1.Error = err
			return 0, nil
		} else {

			fmt.Println("nodata")
			var data1 *Errors.RestError
			data1.Error = err
			return 0, nil
		}

	}

	fmt.Println("completed")
	return staffid, nil
}
func Checkfordeletestaffdata(tenantstaffid int) (int, error) {
	var staffid int
	fmt.Println("enrty in staffs")

	stmt, err := database.Db.Prepare(checkfordeletestaff)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(tenantstaffid)
	// print(row)
	err = row.Scan(&staffid)
	print(err)
	fmt.Println("2")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
			data1 := Errors.RestError{}
			data1.Error = err
			return 0, nil
		} else {

			fmt.Println("nodata")
			var data1 *Errors.RestError
			data1.Error = err
			return 0, nil
		}

	}

	fmt.Println("completed")
	return staffid, nil
}

func (s *TenantUser) TenantstaffCreation(data []int) (bool, error) {
	var staffid int64
	var err error
	tx, err := database.Db.Begin()
	if err != nil {
		return false, err
	}
	{
		print("entry in tenanstaff")
		stmt, err := tx.Prepare(insertTenantstaff)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		defer stmt.Close()

		res, err := stmt.Exec(&s.Tenantid, &s.Moduleid, &s.Userid)
		if err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			return false, err
		}
		staffid, err = res.LastInsertId()
		print("tenantstaffid=", staffid)
	}
	datalist := data
	if len(datalist) != 0 {
		var a TenantUser
		for i := 0; i < len(datalist); i++ {
			a.Tenantstaffid = int(staffid)
			a.Tenantid = s.Tenantid
			a.Locationid = datalist[i]
			{
				print("entry in tenantstaffdetails")
				stmt, err := tx.Prepare(insertTenantstaffdetails)
				if err != nil {
					tx.Rollback()
					return false, err
				}
				defer stmt.Close()
				if _, err := stmt.Exec(&a.Tenantstaffid, &a.Tenantid, &a.Locationid); err != nil {
					tx.Rollback() // return an error too, we may want to wrap them
					return false, err
				}
			}
		}

	}
	tx.Commit()
	return true, nil
}

//firestore
func (t *Initialsubscriptiondata) Firestoreinsertenant(tenantid, locationid int64, moduleid,catid int) error {
	print("st1 firestore")
	n1 := int64(tenantid)
	id := strconv.FormatInt(n1, 10)
	n2 := int64(locationid)
	loc := strconv.FormatInt(n2, 10)
	ctx := context.Background()
	sa := option.WithCredentialsFile(firestorejsonkey)

	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		log.Fatal("failed to create 1 firestore %V", err)
		log.Fatalln(err)
		return err
	}

	client, err := app.Firestore(ctx)

	if err != nil {

		log.Fatal("failed to create  firestore %V", err)
		log.Fatalln(err)
		return err
	}

	defer client.Close()

	_, err1 := client.Collection("tenants").Doc(id).Set(ctx, map[string]interface{}{
		"tenantid":   tenantid,
		"moduleid":   moduleid,
		"locationid": locationid,
		"tenantname": &t.Name,
		"email":      &t.Email,
		"contactno":  &t.Mobile,
		"address":    &t.Address,
		"suburb":     &t.Suburb,
		"city":       &t.City,
		"state":      &t.State,
		"postcode":   &t.Zip,
		"latitude":   &t.Latitude,
		"longitude":  &t.Longitude,
		"status":     "Active",
		"categoryid":catid,
		"tenantimage":"",
	})
	if err1 != nil {

		log.Fatal("failed to insert in  firestore %v", err1)
		return err1
	}
	_, err2 := client.Collection("locations").Doc(loc).Set(ctx, map[string]interface{}{
		"locationid":   locationid,
		"tenantid":     tenantid,
		"locationname": &t.Name,
		"email":        &t.Email,
		"contactno":    &t.Mobile,
		"address":      &t.Address,
		"suburb":       &t.Suburb,
		"city":         &t.City,
		"state":        &t.State,
		"postcode":     &t.Zip,
		"latitude":     &t.Latitude,
		"longitude":    &t.Longitude,
		"status":       "Active",
		"opentime":&t.OpenTime,
		"closetime":&t.CloseTime,
		"delivery":false,
		
	})
	if err2 != nil {

		log.Fatal("failed to insert in  firestore %v", err2)
		return err2
	}
	return nil
}
func (l *Location) Firestorecreatelocation(locationid int64) error {
	print("st1 firestore")

	n2 := int64(locationid)
	loc := strconv.FormatInt(n2, 10)
	ctx := context.Background()
	sa := option.WithCredentialsFile(firestorejsonkey)

	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		log.Fatal("failed to create 1 firestore %V", err)
		log.Fatalln(err)
		return err
	}

	client, err := app.Firestore(ctx)

	if err != nil {

		log.Fatal("failed to create  firestore %V", err)
		log.Fatalln(err)
		return err
	}

	defer client.Close()
	_, err2 := client.Collection("locations").Doc(loc).Set(ctx, map[string]interface{}{
		"locationid":   locationid,
		"tenantid":     &l.TenantID,
		"locationname": &l.LocationName,
		"email":        &l.Email,
		"contactno":    &l.Mobile,
		"address":      &l.Address,
		"suburb":       &l.Suburb,
		"city":         &l.City,
		"state":        &l.State,
		"postcode":     &l.Zip,
		"latitude":     &l.Latitude,
		"longitude":    &l.Longitude,
		"status":       "Active",
		"opentime":&l.OpeningTime,
		"closetime":&l.ClosingTime,
		"delivery":&l.Delivery,
	})
	if err2 != nil {

		log.Fatal("failed to insert in  firestore %v", err2)
		return err2
	}

	return nil
}

func (p *Location) Firestorelocationupdate(locationid int) error {
	print("st1 firestore")
	n1 := int64(locationid)
	id := strconv.FormatInt(n1, 10)
	ctx := context.Background()
	sa := option.WithCredentialsFile(firestorejsonkey)

	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		log.Fatal("failed to create 1 firestore %V", err)
		log.Fatalln(err)
		return err
	}

	client, err := app.Firestore(ctx)

	if err != nil {

		log.Fatal("failed to create  firestore %V", err)
		log.Fatalln(err)
		return err
	}

	defer client.Close()

	ca := client.Collection("locations").Doc(id)

	_, err = ca.Update(context.Background(), []firestore.Update{
		{
			Path: "locationname", Value: &p.LocationName,
		},
		{
			Path: "email", Value: &p.Email,
		},
		{
			Path: "contactno", Value: &p.Mobile,
		},
		{
			Path: "address", Value: &p.Address,
		},
		{
			Path: "suburb", Value: &p.Suburb,
		},
		{
			Path: "city", Value: &p.City,
		},
		{
			Path: "state", Value: &p.State,
		},
		{
			Path: "postcode", Value: &p.Zip,
		},
		{
			Path: "latitude", Value: &p.Latitude,
		},
		{
			Path: "longitude", Value: &p.Longitude,
		},
		{
			Path: "opentime", Value: &p.OpeningTime,
		},
		{
			Path: "closetime", Value: &p.ClosingTime,
		},
		{
			Path: "delivery", Value: &p.Delivery,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
func (p *Updatestatus) Firestoreupdatelocationstatus(locationid int) error {
	print("st1 firestore")
	n1 := int64(locationid)
	id := strconv.FormatInt(n1, 10)
	ctx := context.Background()
	sa := option.WithCredentialsFile(firestorejsonkey)

	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		log.Fatal("failed to create 1 firestore %V", err)
		log.Fatalln(err)
		return err
	}

	client, err := app.Firestore(ctx)

	if err != nil {

		log.Fatal("failed to create  firestore %V", err)
		log.Fatalln(err)
		return err
	}

	defer client.Close()

	ca := client.Collection("locations").Doc(id)

	_, err = ca.Update(context.Background(), []firestore.Update{
		{
			Path: "status", Value: &p.Locationstatus,
		},

	})
	if err != nil {
		return err
	}

	return nil
}

func (p *BusinessUpdate) Firestoreupdatetenant(tenantid int) error {
	print("st1 firestore")
	n1 := int64(tenantid)
	id := strconv.FormatInt(n1, 10)
	ctx := context.Background()
	sa := option.WithCredentialsFile(firestorejsonkey)

	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		log.Fatal("failed to create 1 firestore %V", err)
		log.Fatalln(err)
		return err
	}

	client, err := app.Firestore(ctx)

	if err != nil {

		log.Fatal("failed to create  firestore %V", err)
		log.Fatalln(err)
		return err
	}

	defer client.Close()

	ca := client.Collection("tenants").Doc(id)

	_, err = ca.Update(context.Background(), []firestore.Update{
		{
			Path: "tenantimage", Value: &p.Tenantimage,
		},

	})
	if err != nil {
		return err
	}

	return nil
}