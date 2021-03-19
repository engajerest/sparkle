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
type Modules struct {
	CategoryID int    `json:"categoryId"`
	ModuleID   int    `json:"moduleId"`
	Name       string `json:"name"`
	Content    string `json:"content"`
	Imageurl   string `json:"imageurl"`
	Logourl    string `json:"logourl"`
}
type Packages struct {
	ModuleID       int    `json:"moduleId"`
	ModuleName     string `json:"modulename"`
	PackageID      int    `json:"packageId"`
	Name           string `json:"name"`
	Status         string `json:"status"`
	PackageAmount  string `json:"packageamount"`
	PaymentMode    string `json:"paymentmode"`
	PackageContent string `json:"packagecontent"`
	PackageIcon    string `json:"packageicon"`
	Promocodeid int `json:"promocodeid"`
	Promoname string `json:"promoname"`
	Promodescription string `json:"promodescription"`
	Promotype string `json:"promotype"`
	Promovalue float64 `json:"promovalue"`
	Promovaliditydate string `json:"promovaliditydate"`
	Validity bool `json:"validity"`
	Packageexpiry int `json:"packageexpiry"`


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

type TenantSubscription struct {
	
	Date          string `json:"date"`
	PackageId     int    `json:"packageId"`
	ModuleId      int    `json:"moduleId"`
	CurrencyId    int    `json:"currencyId"`
	CurrencyCode  string `json:"currencycode"`
	Price         string `json:"price"`
	TaxId         int    `json:"taxId"`
	TaxAmount     string `json:"taxamount"`
	TotalAmount   string `json:"totalamount"`
	PaymentStatus int    `json:"paymentstatus"`
	PaymentId     int    `json:"paymentId"`
	Quantity      int    `json:"quantity"`
	Promoid int `json:"promoid"`
	Promovalue string `json:"promovalue"`
	Promostatus bool `json:"promostatus"`
	Validitydate string `json:"validitydate"`
}

type SubscriptionData struct {
	Info    Tenantinfo     `json:"info"`
	Address TenantLocation `json:"address"`
}

type SubscribedData struct {
	TenantID   int    `json:"tenantId"`
	TenantName string `json:"tenantname"`
	ModuleName string `json:"modulename"`
	ModuleID   int    `json:"moduleId"`
	Subscriptionid int `json:"subscriptionid"`
	Locationid int `json:"locationid"`
	Locationname string `json:"locationname"`
}

type TenantUser struct {
	TenantID     int    `json:"tenantid"`
	Userid       int    `json:"userid"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	Locationid   int    `json:"locationid"`
	Locationname string `json:"locationname"`
	RoleId       int    `json:""roleid`
	Referenceid  int    `json:"referenceid"`
	Status       string `json:"status"`
	Created      string `json:created`
}
type Location struct {
	LocationId   int    `json:"locationid"`
	TenantID     int    `json:"tenantid"`
	LocationName string `json:"locationname"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	Address      string `json:"address"`
	Suburb       string `json:"suburb"`
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
	TenantID    int      `json:"tenantid"`
	Moduleid int `json:"moduleid"`
	Modulename string `json:"modulename"`
	Brandname   string   `json:"brandname"`
	TenantaccId string      `json:"tenantaccid"`
	About       string   `json:"about"`
	Email       string   `json:"email"`
	Phone       string   `json:"phone"`
	Address     string   `json:"address"`
	Paymode1    int      `json:"paymode1"`
	Paymode2    int      `json:"paymode2"`
	Tenanttoken string   `json:"tenanttoken"`
	Tenantimage string `json:"tenantimage"`
	SocialData  []Social `json:"social"`
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
type Initial struct{
	Tenantid    int    `json:"tenantid"`
	Locationid int `json:"locationid"`
	Brandname   string `json:"brandname"`
	About       string `json:"about"`
	Tenantimage string `json:"tenantimage"`
	Opentime     string `json:""opentime`
	Closetime    string `json:"closetime"`
}
type AuthUser struct {
	TenantID   int `json:"tenantid"`
	LocationId int `json:"locationid"`
}
type Social struct {
	Socialid     int    `json:"socialid" `
	SociaProfile string `json:"socialprofile"`
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
	Paymentid       int      `json:"paymentid" gorm:"primary_key"`
	Packageid       int      `json:"packageid"`
	Packagename    string `json:"packagename"`
	Locationid int `json:"locationid"`
	Paymenttypeid   int      `json:"paymenttypeid"`
	Paymentref string `json:"paymentref"`
	Tenantid        int      `json:"tenantid"`
	Customerid      int      `gorm:"ForeignKey"`
	Transactiondate string   `json:"transactiondate"`
	Chargeid        string   `json:"chargeid"`
	Amount          float64  `json:"amount"`
	Paymentstatus   string   `json:"paymentstatus"`
	Refundamt       float64  `json:"refundamount"`
	Orderid         int      `json:"orderid"`
	Created         string   `json:"created"`
	
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Contactno  string `json:"contactno"`
	Email      string `json:"email"`
	Address    string `json:"address"`


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
	Packageid  int    `json:"packageid"`
	Tenantid   int    `json:"tenantid"`
	Moduleid   int    `json:"moduleId"`
	Modulename string `json:"modulename"`
	Packagename string `json:"packagename"`
	PackageAmount float64 `json:"packageamount"`
	Totalamount   float64 `json:"totalamount"`
	Logourl       string  `json:"logourl"`
	PackageIcon   string  `json:"packageicon"`
	Customercount int     `json:"customercount"`
	Locationcount int     `json:"locationcount"`
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