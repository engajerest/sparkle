package subscription

type Category struct {
	CategoryID int    `json:"categoryId"`
	Name       string `json:"name"`
	Typeid     int    `json:"typeId"`
	SortOrder  int    `json:"sortOrder"`
	Status     string `json:"status"`
}
type SubCategory struct {
	CategoryID    int    `json:"categoryId"`
	SubCategoryID int    `json:"subcategoryId"`
	Name          string `json:"name"`
	Typeid        int    `json:"typeId"`
	SortOrder     int    `json:"sortOrder"`
	Status        string `json:"status"`
	Icon          string `json:"icon"`
}

// type Appsubcategory struct{
// 	Categoryid   int    `json:"categoryid"`
// 	SubCategoryid int    `json:"subcategoryid"`
// 	Subcategoryname string `json:"subcategoryname"`
// 	Status string

// }
type Modules struct {
	CategoryID int    `json:"categoryId"`
	ModuleID   int    `json:"moduleId"`
	Name       string `json:"name"`
	Content    string `json:"content"`
	Imageurl   string `json:"imageurl"`
	Logourl    string `json:"logourl"`
}
type Packages struct {
	ModuleID          int     `json:"moduleId"`
	ModuleName        string  `json:"modulename"`
	PackageID         int     `json:"packageId"`
	Name              string  `json:"name"`
	Status            string  `json:"status"`
	PackageAmount     string  `json:"packageamount"`
	PaymentMode       string  `json:"paymentmode"`
	PackageContent    string  `json:"packagecontent"`
	PackageIcon       string  `json:"packageicon"`
	Promocodeid       int     `json:"promocodeid"`
	Promoname         string  `json:"promoname"`
	Promodescription  string  `json:"promodescription"`
	Promotype         string  `json:"promotype"`
	Promovalue        float64 `json:"promovalue"`
	Promovaliditydate string  `json:"promovaliditydate"`
	Validity          bool    `json:"validity"`
	Packageexpiry     int     `json:"packageexpiry"`
}
type Tenantinfo struct {
	Name          string `json:"name"`
	Regno         string `json:"regno"`
	Email         string `json:"email"`
	Mobile        string `json:"mobile"`
	CategoryId    int    `json:"categoryId"`
	Typeid        int    `json:"typeId"`
	SubCategoryID int    `json:"subcategoryId"`
	Tenanttoken   string `json:"tenanttoken"`
}
type TenantLocation struct {
	Address      string `json:"address"`
	Suburb       string `json:"suburb"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
	Countrycode  string `json:"countrycode"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	TimeZone     string `json:"timezone"`
	CurrencyCode string `json:"currencycode"`
	OpenTime     string `json:"opentime"`
	CloseTime    string `json:"closetime"`
}

type Initialsubscriptiondata struct {
	Name            string `json:"name"`
	Regno           string `json:"regno"`
	Email           string `json:"email"`
	Mobile          string `json:"mobile"`
	Categoryid      int    `json:"categoryid"`
	Typeid          int    `json:"typeId"`
	SubCategoryid   int    `json:"subcategoryid"`
	Subcategoryname string `json:"subcategoryname"`
	Tenanttoken     string `json:"tenanttoken"`
	Address         string `json:"address"`
	Suburb          string `json:"suburb"`
	City            string `json:"city"`
	State           string `json:"state"`
	Zip             string `json:"zip"`
	Countrycode     string `json:"countrycode"`
	Currencyid      int    `json:"currencyid"`

	Currencysymbol  string               `json:"currencysymbol"`
	Latitude        string               `json:"latitude"`
	Longitude       string               `json:"longitude"`
	TimeZone        string               `json:"timezone"`
	CurrencyCode    string               `json:"currencycode"`
	OpenTime        string               `json:"opentime"`
	CloseTime       string               `json:"closetime"`
	Userid          int                  `json:"userid"`
	Partnerid       int                  `json:"partnerid"`
	Tenantsubscribe []TenantSubscription `json:"tenantsubscribe"`
}

type TenantSubscription struct {
	Date       string `json:"date"`
	Packageid  int    `json:"packageid"`
	Partnerid  int    `json:"partnerid"`
	Moduleid   int    `json:"moduleId"`
	Currencyid int    `json:"currencyid"`

	Categoryid      int    `json:"categoryid"`
	SubCategoryid   int    `json:"subcategoryid"`
	Subcategoryname string `json:"subcategoryname"`
	Tenantid        int    `json:"tenantid"`
	Price           string `json:"price"`
	TaxId           int    `json:"taxId"`
	TaxAmount       string `json:"taxamount"`
	TotalAmount     string `json:"totalamount"`
	PaymentStatus   int    `json:"paymentstatus"`
	PaymentId       int    `json:"paymentId"`
	Quantity        int    `json:"quantity"`
	Promoid         int    `json:"promoid"`
	Promovalue      string `json:"promovalue"`
	Promostatus     bool   `json:"promostatus"`
	Validitydate    string `json:"validitydate"`
	Subscriptionid int `json:"subscriptionid"`
}

type SubscriptionData struct {
	Info    Tenantinfo     `json:"info"`
	Address TenantLocation `json:"address"`
}

type SubscribedData struct {
	TenantID       int     `json:"tenantId"`
	TenantName     string  `json:"tenantname"`
	ModuleName     string  `json:"modulename"`
	ModuleID       int     `json:"moduleId"`
	Subscriptionid int     `json:"subscriptionid"`
	Locationid     int     `json:"locationid"`
	Locationname   string  `json:"locationname"`
	Subcategoryid  int     `json:"subcategoryid"`
	Categoryid     int     `json:"categoryid"`
	Tenantaccid    string  `json:"tenantaccid"`
	Taxamount      float64 `json:"taxamount"`
	Totalamount    float64 `json:"totalamount"`
}

type TenantUser struct {
	Tenantid      int    `json:"tenantid"`
	Userid        int    `json:"userid"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	Mobile        string `json:"mobile"`
	Profileimage  string `json:"profileimage"`
	Locationid    int    `json:"locationid"`
	Locationname  string `json:"locationname"`
	Roleid        int    `json:"roleid"`
	Configid      int    `json:"configid"`
	Referenceid   int    `json:"referenceid"`
	Status        string `json:"status"`
	Created       string `json:"created`
	Moduleid      int    `json:"moduleid"`
	Tenantstaffid int    `json:"tenantstaffid"`
	Staffdetailid int    `json:"staffdetailid"`
}
type Location struct {
	LocationId   int    `json:"locationid"`
	TenantID     int    `json:"tenantid"`
	LocationName string `json:"locationname"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	Address      string `json:"address"`
	Suburb       string `json:"suburb"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
	Countrycode  string `json:"countrycode"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	OpeningTime  string `json:""openingtime`
	ClosingTime  string `json:"closingtime"`
	Delivery     bool   `json:"delivery"`
	Deliverytype string `json:"deliverytype"`
	Deliverymins int    `json:"deliverymins"`
	Createdby    int    `json:"createdby"`
	Status       string `json:"status"`
}
type BusinessUpdate struct {
	TenantID       int      `json:"tenantid"`
	Moduleid       int      `json:"moduleid"`
	Modulename     string   `json:"modulename"`
	Brandname      string   `json:"brandname"`
	TenantaccId    string   `json:"tenantaccid"`
	About          string   `json:"about"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Address        string   `json:"address"`
	Paymode1       int      `json:"paymode1"`
	Paymode2       int      `json:"paymode2"`
	Tenanttoken    string   `json:"tenanttoken"`
	Tenantimage    string   `json:"tenantimage"`
	Countrycode    string   `json:"countrycode"`
	Currencycode   string   `json:"currencycode"`
	Currencysymbol string   `json:"currencysymbol"`
	SocialData     []Social `json:"social"`
	// SociaProfile string `json:"socialprofile"`
	// SocialLink string `json:"sociallink"`
	// SocialIcon string `json:"socialicon"`

}
type Tenant struct {
	Tenantid    int    `json:"tenantid" gorm:"primary_key"`
	Brandname   string `json:"brandname"`
	TenantaccId int    `json:"tenantaccid"`
	About       string `json:"about"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Paymode1    int    `json:"paymode1"`
	Paymode2    int    `json:"paymode2"`
	Tenanttoken string `json:"tenanttoken"`
}
type Initial struct {
	Tenantid    int    `json:"tenantid"`
	Locationid  int    `json:"locationid"`
	Brandname   string `json:"brandname"`
	About       string `json:"about"`
	Tenantimage string `json:"tenantimage"`
	Opentime    string `json:""opentime`
	Closetime   string `json:"closetime"`
}
type AuthUser struct {
	TenantID   int `json:"tenantid"`
	LocationId int `json:"locationid"`
}
type Social struct {
	Socialid     int    `json:"socialid" `
	SociaProfile string `json:"socialprofile"`
	Dailcode     string `json:"dailcode"`
	SocialLink   string `json:"sociallink"`
	SocialIcon   string `json:"socialicon"`
}
type Tenantsocial struct {
	Socialid     int    `json:"socialid" gorm:"primary_key"`
	Tenantid     int    `json:"tenantid"`
	Sociaprofile string `json:"socialprofile"`
	Sociallink   string `json:"sociallink"`
	Socialicon   string `json:"socialicon"`
}
type Tenantlocation struct {
	// gorm.Model
	Locationid   int    `json:"locationid" gorm:"primary_key"`
	Tenantid     int    `json:"tenantid" gorm:"ForeignKey"`
	Locationname string `json:"locationname"`
	Email        string `json:"email"`
	Contactno    string `json:"contactno"`
	Address      string `json:"address"`
	Suburb       string `json:"suburb"`
	City         string `json:"city"`
	State        string `json:"state"`
	Postcode     string `json:"postcode"`
	Countrycode  string `json:"countrycode"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Opentime     string `json:""opentime`
	Closetime    string `json:"closetime"`
	Delivery     bool   `json:"delivery"`
	Deliverytype string `json:"deliverytype"`
	Deliverymins int    `json:"deliverymins"`
	Createdby    int    `json:"createdby"`
	Status       string `json:"status"`
	// Tenantstaffdetails []Tenantstaffdetails `json:"tenantstaffdetails" gorm:"ForeignKey:locationid"`
	Appuserprofiles []App_userprofiles `json:"appuserprofile" gorm:"ForeignKey:userlocationid"`
	Tenantcharges   []Tenantcharge     `json:"tenantcharge" gorm:"ForeignKey:locationid"`
	Tenantsettings  []Tenantsetting    `json:"tenantsetting" gorm:"ForeignKey:locationid"`
}
type App_userprofiles struct {
	// gorm.Model
	Profileid      int    `gorm:"primary_key"`
	Userid         int    `json:"userid"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Email          string `json:"email"`
	Contactno      string `json:"contactno"`
	Profileimage   string `json:"profileimage"`
	Userlocationid int    `json:"userlocationid"`
}
type Promotion struct {
	Promotionid     int    `json:"promotionid"`
	Promotiontypeid int    `json:"promotionid"`
	Tenantid        int    `json:"tenantid"`
	Tenantname      string `json:"tenantname"`
	Promoname       string `json:"promoname"`
	Promocode       string `json:"promocode"`
	Promovalue      string `json:"promovalue"`
	Promoterms      string `json:"promoterms"`
	Promotype       string `json:"promotype"`
	Promotag        string `json:"promotag"`
	Startdate       string `json:"startdate"`
	Enddate         string `json:"enddate"`
	Broadcaststatus bool   `json:"broadcaststatus"`
	Success         int    `json:"success"`
	Failure         int    `json:"failure"`
	Status          string `json:"status"`
}
type Ordersequence struct {
	Sequenceid int    `json:"sequenceid"`
	Tenantid   int    `json:"tenantid"`
	Tablename  string `json:"tablename"`
	Seqno      int    `json:"seqno"`
	Prefix     string `json:"prefix"`
	Subprefix  int    `json:"subprefix"`
}
type Tenantcharge struct {
	Tenantchargeid int    `json:"tenantchargeid" gorm:"primary_key"`
	Tenantid       int    `json:"tenantid"`
	Locationid     int    `json:"locationid"`
	Chargeid       int    `json:"chargeid"`
	Chargename     string `json:"chargename"`
	Chargetype     string `json:"chargetype"`
	Chargevalue    string `json:"chargevalue"`
	Createdby      int    `json:"createdby"`
	// Chargetypes *Chargetype `gorm:"ForeignKey:chargeid" `

}
type Chargetype struct {
	Chargeid   int    `gorm:"primary_key"`
	Chargename string `json:"chargename"`
	Status     string `json:"status"`
}

type Tenantsetting struct {
	Settingsid int    `json:"settinsid"`
	Tenantid   int    `json:"tenantid"`
	Locationid int    `json:"locationid"`
	Slabtype   string `json:"slabtype"`
	Slab       string `json:"slab"`
	Slablimit  int    `json:"slablimit"`
	Slabcharge string `json:"slabcharge"`
	Createdby  int    `json:"createdby"`
}
type Updatestatus struct {
	Tenantid       int    `json:"tenantid"`
	Locationid     int    `json:"locationid"`
	Locationstatus string `json:"locationstatus"`
	Deliverystatus bool   `json:"deliverystatus"`
}
type Payment struct {
	Paymentid       int             `json:"paymentid" gorm:"primary_key"`
	Moduleid        int             `json:"moduleid"`
	Locationid      int             `json:"locationid"`
	Paymentref      string          `json:"paymentref"`
	Paymenttypeid   int             `json:"paymenttypeid"`
	Tenantid        int             `json:"tenantid"`
	Customerid      int             `gorm:"ForeignKey"`
	Transactiondate string          `json:"transactiondate"`
	Orderid         int             `json:"orderid"`
	Chargeid        string          `json:"chargeid"`
	Refundid        string          `json:"chargeid"`
	Amount          float64         `json:"amount"`
	Refundamt       float64         `json:"refundamount"`
	Paymentstatus   string          `json:"paymentstatus"`
	Created string `json:"created"`
	Paymentdetails  []Paymentdetail `json:"paymentdetails" gorm:"ForeignKey:paymentid"`
}
type Paymentdetail struct {
	Paymentdetailid int      `json:"paymentdetailid"`
	Paymentid       int      `json:"paymentid"`
	Tenantid        int      `json:"tenantid"`
	Subscriptionid  int      `json:"subscriptionid"`
	Moduleid        int      `json:"moduleid"`
	Locationid      int      `json:"locationid"`
	Orderid         int      `json:"orderid"`
	Customerid      int      `json:"customerid" `
	Amount          float64  `json:"amount"`
	Taxpercent      int      `json:"taxpercent"`
	Taxamount       float64  `json"taxamount"`
	Payamount       float64  `json:"payamount"`
	Customers       Customer `json:"customers" gorm:"ForeignKey:customerid;references:customerid"  `
}

type Customer struct {
	Customerid int    `gorm:"primary_key"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Contactno  string `json:"contactno"`
	Email      string `json:"email"`
	Address    string `json:"address"`
}
type Subscribe struct {
	Packageid            int     `json:"packageid"`
	Subscriptionid       int     `json:"subscriptionid"`
	Subscriptionaccid    string  `json:"subscriptionaccid"`
	Subscriptionmethodid string  `json:"subscriptionmethodid"`
	Tenantid             int     `json:"tenantid"`
	Moduleid             int     `json:"moduleId"`
	Modulename           string  `json:"modulename"`
	Packagename          string  `json:"packagename"`
	PackageAmount        float64 `json:"packageamount"`
	Totalamount          float64 `json:"totalamount"`
	Taxamount            float64 `json:"taxamount"`
	Logourl              string  `json:"logourl"`
	Iconurl              string  `json:"iconurl"`
	PackageIcon          string  `json:"packageicon"`
	Tenantaccid          string  `json:"tenantaccid"`
	Customercount        int     `json:"customercount"`
	Locationcount        int     `json:"locationcount"`
	Subcategoryid        int     `json:"subcategoryid"`
	Categoryid           int     `json:"categoryid"`
	Paymentstatus        bool    `json:"paymentstatus"`
	Validity bool `json:"validity"`
	Validitydate string `json:"validitydate"`
}
type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	CreatedDate string `json:"created"`
	Status      string `json:"status"`
	Roleid      int    `json:"roleid"`
	Configid    int    `json:"configid"`
	Referenceid int    `json:"referenceid"`
	LocationId  int    `json:"locationid"`
	Moduleid    int    `json:"moduleid"`
	Packageid   int    `json:"packageid"`
	Modulename  string `json:"modulename"`
	Tenantname  string `json:"tenantname"`
	From        string `json:"from"`
}
type Module struct {
	Moduleid   int    `json:"moduleid"`
	Categoryid int    `json:"cateoryid"`
	Modulename string `json:"modulename"`
	Content    string `json:"content"`
	Logourl    string `json:"logourl"`
	Iconurl    string `json:"iconurl"`
}
type Tenants struct {
	Tenantid            int                  `json:"tenantid" gorm:"primary_key"`
	Partnerid           int                  `json:"partnerid"`
	Registrationno      string               `json:"registrationno"`
	Tenanttoken         string               `json:"tenanttoken"`
	Tenantname          string               `json:"tenantname"`
	Primaryemail        string               `json:"primaryemail"`
	Primarycontact      string               `json:"primarycontact"`
	Brandname           string               `json:"brandname"`
	TenantaccId         int                  `json:"tenantaccid"`
	About               string               `json:"about"`
	Email               string               `json:"email"`
	Phone               string               `json:"phone"`
	Address             string               `json:"address"`
	Paymode1            int                  `json:"paymode1"`
	Paymode2            int                  `json:"paymode2"`
	Tenantsubscriptions []Tenantsubscription `json:"tenantsubscriptions" gorm:"ForeignKey:tenantid;references:tenantid"`
	Tenantlocations     []Tenantlocation     `json:"tenantlocations" gorm:"ForeignKey:tenantid;references:tenantid"`
	Tenantsubcategories []Tenantsubcategory  `json:"tenantsubcategories" gorm:"ForeignKey:tenantid;references:tenantid"`
}
type Tenantsubcategory struct {
	Tenantid        int    `json:"tenantid"`
	Moduleid        int    `json:"moduleid"`
	Categoryid      int    `json:"categoryid"`
	Subcategoryid   int    `json:"subcategoryid"`
	Subcategoryname string `json:"subcategoryname"`
}
type Tenantsubscription struct {
	Date            string `json:"date"`
	Packageid       int    `json:"packageid"`
	Partnerid       int    `json:"partnerid"`
	Moduleid        int    `json:"moduleId"`
	Currencyid      int    `json:"currencyid"`
	Categoryid      int    `json:"categoryid"`
	SubCategoryid   int    `json:"subcategoryid"`
	Subcategoryname string `json:"subcategoryname"`
	Tenantid        int    `json:"tenantid"`
	Price           string `json:"price"`
	TaxId           int    `json:"taxId"`
	TaxAmount       string `json:"taxamount"`
	TotalAmount     string `json:"totalamount"`
	PaymentStatus   int    `json:"paymentstatus"`
	PaymentId       int    `json:"paymentId"`
	Quantity        int    `json:"quantity"`
	Promoid         int    `json:"promoid"`
	Promovalue      string `json:"promovalue"`
	Promostatus     bool   `json:"promostatus"`
	Validitydate    string `json:"validitydate"`
}
type Tenantstaff struct {
	Tenantstaffid      int                 `json:"tenantstaffid" gorm:"primary_key"`
	Tenantid           int                 `json:"tenantid"`
	Moduleid           int                 `json:"moduleid"`
	Userid             int                 `json:"userid"`
	Firstname          string              `json:"firstname"`
	Lastname           string              `json:"lastname"`
	Email              string              `json:"email"`
	Contactno          string              `json:"contactno"`
	Profileimage       string              `json:"profileimage"`
	Tenantstaffdetails []Tenantstaffdetail `json:"tenantstaffdetails" gorm:"ForeignKey:tenantstaffid;references:tenantstaffid"`
}
type Tenantstaffdetail struct {
	Staffdetailid   int            `json:"staffdetailid"`
	Tenantstaffid   int            `json:"tenantstaffid"`
	Tenantid        int            `json:"tenantid"`
	Locationid      int            `json:"locationid"`
	Tenantlocations Tenantlocation `json:"tenantlocation" gorm:"ForeignKey:locationid;references:locationid"`
}

//seperation process
type Tenantstaffdetails struct {
	Staffdetailid int            `json:"staffdetailid"`
	Tenantstaffid int            `json:"tenantstaffid"`
	Tenantid      int            `json:"tenantid"`
	Locationid    int            `json:"locationid"`
	Tenantstaffs  []Tenantstaffs `json:"tenantstaffs" gorm:"ForeignKey:tenantstaffid;references:tenantstaffid"`
}

type Tenantstaffs struct {
	Tenantstaffid   int              `json:"tenantstaffid" gorm:"primary_key"`
	Tenantid        int              `json:"tenantid"`
	Moduleid        int              `json:"moduleid"`
	Userid          int              `json:"userid"`
	Appuserprofiles App_userprofiles `json:"appuserprofiles" gorm:"ForeignKey:userid;references:userid"`
}
type TenantUsers struct {
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Profileimage   string `json:"profileimage"`
	Userlocationid int    `json:"userlocationid"`
	Userid         int    `json:"userid"`
	Created        string `json:"created"`
	Contactno      string `json:"contactno"`
	Email          string `json:"email"`
	Locationname   string `json:"locationname"`
	Referenceid    int    `json:"referenceid"`
}
