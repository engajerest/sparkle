package subscription

// import "gorm.io/gorm"

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
}
type Tenantinfo struct {
	Name          string `json:"name"`
	Regno         string `json:"regno"`
	Email         string `json:"email"`
	Mobile        string `json:"mobile"`
	CategoryId    int    `json:"categoryId"`
	Typeid        int    `json:"typeId"`
	SubCategoryID int    `json:"subcategoryId"`
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
	Createdby    int    `json:"createdby"`
	Status       string `json:"status"`
}
type BusinessUpdate struct {
	TenantID    int      `json:"tenantid"`
	Brandname   string   `json:"brandname"`
	TenantaccId int      `json:"tenantaccid"`
	About       string   `json:"about"`
	Paymode1    int      `json:"paymode1"`
	Paymode2    int      `json:"paymode2"`
	SocialData  []Social `json:"social"`
	// SociaProfile string `json:"socialprofile"`
	// SocialLink string `json:"sociallink"`
	// SocialIcon string `json:"socialicon"`

}
type AuthUser struct {
	TenantID   int `json:"tenantid"`
	LocationId int `json:"locationid"`
}
type Social struct {
	Socialid int `json:"socialid"`
	SociaProfile string `json:"socialprofile"`
	SocialLink   string `json:"sociallink"`
	SocialIcon   string `json:"socialicon"`
}
type Tenantlocation struct {
	// gorm.Model
	Locationid      int                `json:"locationid" gorm:"primary_key"`
	Tenantid        int                `json:"tenantid"`
	Locationname    string             `json:"locationname"`
	Email           string             `json:"email"`
	Contactno       string             `json:"contactno"`
	Address         string             `json:"address"`
	City            string             `json:"city"`
	State           string             `json:"state"`
	Postcode        string             `json:"postcode"`
	Countrycode     string             `json:"countrycode"`
	Latitude        string             `json:"latitude"`
	Longitude       string             `json:"longitude"`
	Opentime        string             `json:""opentime`
	Closetime       string             `json:"closetime"`
	Createdby       int                `json:"createdby"`
	Status          string             `json:"status"`
	Appuserprofiles []App_userprofiles `json:"appuserprofile" gorm:"ForeignKey:userlocationid"`
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
	Promotionid     int       `json:"promotionid"`
	Promotiontypeid int       `json:"promotionid"`
	Tenantid        int       `json:"tenantid"`
	Tenantname      string    `json:"tenantname"`
	Promoname       string    `json:"promoname"`
	Promocode       string    `json:"promocode"`
	Promovalue      string   `json:"promovalue"`
	Promoterms      string    `json:"promoterms"`
	Promotype       string    `json:"promotype"`
	Promotag        string    `json:"promotag"`
	Startdate       string `json:"startdate"`
	Enddate         string `json:"enddate"`
	Status          string    `json:"status"`
	
}
type Ordersequence struct {
	Sequenceid int    `json:"sequenceid"`
	Tenantid   int    `json:"tenantid"`
	Tablename  string `json:"tablename"`
	Seqno      int `json:"seqno"`
	Prefix     string `json:"prefix"`
	Subprefix  int    `json:"subprefix"`
}