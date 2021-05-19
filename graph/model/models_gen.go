// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Cat struct {
	Categoryid   int    `json:"Categoryid"`
	Categoryname string `json:"Categoryname"`
	Categorytype int    `json:"Categorytype"`
	Sortorder    int    `json:"Sortorder"`
	Status       string `json:"Status"`
}

type Category struct {
	Categoryid int    `json:"Categoryid"`
	Name       string `json:"Name"`
	Type       int    `json:"Type"`
	SortOrder  int    `json:"SortOrder"`
	Status     string `json:"Status"`
}

type Custinfo struct {
	Customerid int    `json:"Customerid"`
	Firstname  string `json:"Firstname"`
	Lastname   string `json:"Lastname"`
	Email      string `json:"Email"`
	Contact    string `json:"Contact"`
	Address    string `json:"Address"`
}

type LocationInfo struct {
	Locationid   int    `json:"Locationid"`
	LocationName string `json:"LocationName"`
	Status       string `json:"status"`
	Createdby    int    `json:"createdby"`
}

type Module struct {
	ModuleID   int    `json:"ModuleId"`
	CategoryID int    `json:"CategoryId"`
	Name       string `json:"Name"`
	Content    string `json:"Content"`
	ImageURL   string `json:"ImageUrl"`
	LogoURL    string `json:"LogoUrl"`
}

type Package struct {
	ModuleID         int     `json:"ModuleId"`
	Modulename       string  `json:"Modulename"`
	Name             string  `json:"Name"`
	PackageID        int     `json:"PackageId"`
	Status           string  `json:"Status"`
	PackageAmount    string  `json:"PackageAmount"`
	PaymentMode      string  `json:"PaymentMode"`
	PackageContent   string  `json:"PackageContent"`
	PackageIcon      string  `json:"PackageIcon"`
	Promocodeid      int     `json:"Promocodeid"`
	Promonname       string  `json:"Promonname"`
	Promodescription string  `json:"Promodescription"`
	Promotype        string  `json:"Promotype"`
	Promovalue       float64 `json:"Promovalue"`
	Packageexpiry    int     `json:"Packageexpiry"`
	Validitydate     string  `json:"Validitydate"`
	Validity         bool    `json:"Validity"`
}

type Promotion struct {
	Promotionid     int     `json:"Promotionid"`
	Promotiontypeid int     `json:"Promotiontypeid"`
	Promotionname   string  `json:"Promotionname"`
	Tenantid        int     `json:"Tenantid"`
	Tenantame       string  `json:"Tenantame"`
	Promocode       string  `json:"Promocode"`
	Promoterms      string  `json:"Promoterms"`
	Promovalue      string  `json:"Promovalue"`
	Promotag        string  `json:"Promotag"`
	Promotype       string  `json:"Promotype"`
	Startdate       string  `json:"Startdate"`
	Enddate         string  `json:"Enddate"`
	Broadstatus     bool    `json:"Broadstatus"`
	Success         int     `json:"Success"`
	Failure         int     `json:"Failure"`
	Status          *string `json:"Status"`
}

type Sparkle struct {
	Category    []*Cat     `json:"category"`
	Subcategory []*Subcat  `json:"subcategory"`
	Package     []*Package `json:"package"`
}

type SubCategory struct {
	CategoryID    int    `json:"CategoryId"`
	SubCategoryID int    `json:"SubCategoryId"`
	Name          string `json:"Name"`
	Type          int    `json:"Type"`
	SortOrder     int    `json:"SortOrder"`
	Status        string `json:"Status"`
	Icon          string `json:"Icon"`
}

type TenantAddress struct {
	Address        string `json:"Address"`
	Suburb         string `json:"Suburb"`
	City           string `json:"City"`
	State          string `json:"State"`
	Zip            string `json:"Zip"`
	Countrycode    string `json:"Countrycode"`
	Currencyid     int    `json:"Currencyid"`
	Currencycode   string `json:"Currencycode"`
	Currencysymbol string `json:"Currencysymbol"`
	Latitude       string `json:"Latitude"`
	Longitude      string `json:"Longitude"`
	TimeZone       string `json:"TimeZone"`
	Opentime       string `json:"Opentime"`
	Closetime      string `json:"Closetime"`
}

type TenantData struct {
	Tenantid       int     `json:"Tenantid"`
	Tenantname     string  `json:"Tenantname"`
	Moduleid       int     `json:"Moduleid"`
	Modulename     string  `json:"Modulename"`
	Subscriptionid int     `json:"Subscriptionid"`
	Tenantaccid    string  `json:"Tenantaccid"`
	Locationid     int     `json:"Locationid"`
	Locationname   string  `json:"Locationname"`
	Categoryid     int     `json:"Categoryid"`
	Subcategoryid  int     `json:"Subcategoryid"`
	Taxamount      float64 `json:"Taxamount"`
	Totalamount    float64 `json:"Totalamount"`
}

type TenantDetails struct {
	Name        string `json:"Name"`
	Regno       string `json:"Regno"`
	Email       string `json:"Email"`
	Mobile      string `json:"Mobile"`
	Type        int    `json:"Type"`
	Tenanttoken string `json:"Tenanttoken"`
}

type Tenantschema struct {
	Tenantid    int     `json:"Tenantid"`
	Moduleid    int     `json:"Moduleid"`
	Modulename  string  `json:"Modulename"`
	Brandname   *string `json:"brandname"`
	About       *string `json:"about"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	Address     *string `json:"address"`
	Cod         *int    `json:"cod"`
	Digital     *int    `json:"digital"`
	Tenantaccid *string `json:"tenantaccid"`
	Tenanttoken *string `json:"tenanttoken"`
	Tenantimage *string `json:"tenantimage"`
}

type Business struct {
	Businessupdate *Businessupdatedata `json:"businessupdate"`
	Socialadd      []*Socialadddata    `json:"socialadd"`
	Socialupdate   []*Socialupdatedata `json:"socialupdate"`
	Socialdelete   []*int              `json:"socialdelete"`
}

type Businessdata struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Updated int    `json:"updated"`
}

type Businessupdatedata struct {
	Tenantid    int     `json:"tenantid"`
	Brandname   *string `json:"brandname"`
	About       *string `json:"about"`
	Cod         *int    `json:"cod"`
	Digital     *int    `json:"digital"`
	Tenantimage string  `json:"tenantimage"`
}

type Chargecreate struct {
	Deliverycharges []*Deliverychargeinput `json:"deliverycharges"`
	Othercharges    []*Chargecreateinput   `json:"othercharges"`
}

type Chargecreateinput struct {
	Tenantid    int    `json:"Tenantid"`
	Locationid  int    `json:"Locationid"`
	Chargeid    int    `json:"Chargeid"`
	Chargename  string `json:"Chargename"`
	Chargetype  string `json:"Chargetype"`
	Chargevalue string `json:"Chargevalue"`
}

type Chargetype struct {
	Chargeid   int     `json:"Chargeid"`
	Chargename string  `json:"Chargename"`
	Status     *string `json:"Status"`
}

type Chargetypedata struct {
	Status  bool          `json:"status"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Types   []*Chargetype `json:"types"`
}

type Chargeupdate struct {
	Updatedeliverycharges *Updatedelivery `json:"updatedeliverycharges"`
	Updateothercharges    *Updateother    `json:"updateothercharges"`
}

type Chargeupdateinput struct {
	Tenantchargeid int    `json:"Tenantchargeid"`
	Tenantid       int    `json:"Tenantid"`
	Locationid     int    `json:"Locationid"`
	Chargeid       int    `json:"Chargeid"`
	Chargename     string `json:"Chargename"`
	Chargetype     string `json:"Chargetype"`
	Chargevalue    string `json:"Chargevalue"`
}

type Data struct {
	Tenantinfo          *TenantDetails      `json:"tenantinfo"`
	Tenantlocation      *TenantAddress      `json:"tenantlocation"`
	Subscriptiondetails []*Initialsubscribe `json:"subscriptiondetails"`
}

type Deliverycharge struct {
	Settingsid int    `json:"Settingsid"`
	Tenantid   int    `json:"Tenantid"`
	Locationid int    `json:"Locationid"`
	Slabtype   string `json:"Slabtype"`
	Slab       string `json:"Slab"`
	Slablimit  int    `json:"Slablimit"`
	Slabcharge string `json:"Slabcharge"`
}

type Deliverychargeinput struct {
	Tenantid   int    `json:"Tenantid"`
	Locationid int    `json:"Locationid"`
	Slabtype   string `json:"Slabtype"`
	Slab       string `json:"Slab"`
	Slablimit  int    `json:"Slablimit"`
	Slabcharge string `json:"Slabcharge"`
}

type Delstatus struct {
	Tenantid   int  `json:"tenantid"`
	Locationid int  `json:"locationid"`
	Delivery   bool `json:"delivery"`
}

type GetBusinessdata struct {
	Status       bool   `json:"status"`
	Code         int    `json:"code"`
	Message      string `json:"message"`
	Businessinfo *Info  `json:"businessinfo"`
}

type Getalllocations struct {
	Status    bool              `json:"status"`
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Locations []*Locationgetall `json:"locations"`
}

type Getallmoduledata struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Modules []*Mod `json:"modules"`
}

type Getallpromodata struct {
	Status  bool     `json:"status"`
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Promos  []*Promo `json:"promos"`
}

type Getnonsubscribedcategorydata struct {
	Status   bool   `json:"status"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Category []*Cat `json:"category"`
}

type Getnonsubscribeddata struct {
	Status        bool       `json:"status"`
	Code          int        `json:"code"`
	Message       string     `json:"message"`
	Nonsubscribed []*Package `json:"nonsubscribed"`
}

type Getpaymentdata struct {
	Status   bool           `json:"status"`
	Code     int            `json:"code"`
	Message  string         `json:"message"`
	Payments []*Paymentdata `json:"payments"`
}

type Getpromotiondata struct {
	Status     bool         `json:"status"`
	Code       int          `json:"code"`
	Message    string       `json:"message"`
	Promotions []*Promotion `json:"promotions"`
}

type Getsubcategorydata struct {
	Status        bool      `json:"status"`
	Code          int       `json:"code"`
	Message       string    `json:"message"`
	Subcategories []*Subcat `json:"subcategories"`
}

type Getsubscriptionsdata struct {
	Status     bool                 `json:"status"`
	Code       int                  `json:"code"`
	Message    string               `json:"message"`
	Subscribed []*Subscriptionsdata `json:"subscribed"`
}

type Gettenantsubcategorydata struct {
	Status              bool            `json:"status"`
	Code                int             `json:"code"`
	Message             string          `json:"message"`
	Tenantsubcategories []*Tenantsubcat `json:"tenantsubcategories"`
}

type Info struct {
	Tenantid       int           `json:"tenantid"`
	Moduleid       int           `json:"moduleid"`
	Modulename     string        `json:"modulename"`
	Brandname      *string       `json:"brandname"`
	About          *string       `json:"about"`
	Email          *string       `json:"email"`
	Phone          *string       `json:"phone"`
	Address        *string       `json:"address"`
	Cod            *int          `json:"cod"`
	Digital        *int          `json:"digital"`
	Tenantaccid    *string       `json:"tenantaccid"`
	Tenanttoken    *string       `json:"tenanttoken"`
	Tenantimage    *string       `json:"tenantimage"`
	Countrycode    string        `json:"countrycode"`
	Currencycode   string        `json:"currencycode"`
	Currencysymbol string        `json:"currencysymbol"`
	Social         []*Socialinfo `json:"social"`
}

type Initialsubscribe struct {
	TransactionDate string `json:"TransactionDate"`
	Packageid       int    `json:"Packageid"`
	Partnerid       int    `json:"Partnerid"`
	Moduleid        int    `json:"Moduleid"`
	Categoryid      int    `json:"Categoryid"`
	SubCategoryid   int    `json:"SubCategoryid"`
	Subcategoryname string `json:"Subcategoryname"`
	Currencyid      int    `json:"Currencyid"`
	CurrencyCode    string `json:"CurrencyCode"`
	Price           string `json:"Price"`
	TaxID           int    `json:"TaxId"`
	Quantity        int    `json:"Quantity"`
	Promoid         int    `json:"Promoid"`
	Promovalue      string `json:"Promovalue"`
	TaxAmount       string `json:"TaxAmount"`
	TotalAmount     string `json:"TotalAmount"`
	PaymentStatus   int    `json:"PaymentStatus"`
	Paymentid       *int   `json:"Paymentid"`
	Validitydate    string `json:"Validitydate"`
}

type Location struct {
	TenantID     int    `json:"TenantId"`
	LocationName string `json:"LocationName"`
	Email        string `json:"Email"`
	Contact      string `json:"Contact"`
	Address      string `json:"Address"`
	Suburb       string `json:"Suburb"`
	City         string `json:"City"`
	State        string `json:"State"`
	Zip          string `json:"Zip"`
	Countrycode  string `json:"Countrycode"`
	Latitude     string `json:"Latitude"`
	Longitude    string `json:"Longitude"`
	Openingtime  string `json:"Openingtime"`
	Closingtime  string `json:"Closingtime"`
	Delivery     bool   `json:"Delivery"`
	Deliverytype string `json:"Deliverytype"`
	Deliverymins int    `json:"Deliverymins"`
}

type Locationbyiddata struct {
	Status       bool            `json:"status"`
	Code         int             `json:"code"`
	Message      string          `json:"message"`
	Locationdata *Locationgetall `json:"locationdata"`
}

type Locationdata struct {
	Status       bool          `json:"status"`
	Code         int           `json:"code"`
	Message      string        `json:"message"`
	Locationinfo *LocationInfo `json:"locationinfo"`
}

type Locationgetall struct {
	Locationid      int               `json:"locationid"`
	LocationName    string            `json:"locationName"`
	Tenantid        int               `json:"tenantid"`
	Email           string            `json:"email"`
	Contact         string            `json:"contact"`
	Address         string            `json:"address"`
	Suburb          string            `json:"suburb"`
	City            string            `json:"city"`
	State           string            `json:"state"`
	Postcode        string            `json:"postcode"`
	Countycode      string            `json:"countycode"`
	Latitude        string            `json:"latitude"`
	Longitude       string            `json:"longitude"`
	Openingtime     string            `json:"openingtime"`
	Closingtime     string            `json:"closingtime"`
	Delivery        bool              `json:"delivery"`
	Deliverytype    string            `json:"deliverytype"`
	Deliverymins    int               `json:"deliverymins"`
	Status          string            `json:"status"`
	Createdby       int               `json:"createdby"`
	Tenantusers     []*Userinfodata   `json:"tenantusers"`
	Othercharges    []*Othercharge    `json:"othercharges"`
	Deliverycharges []*Deliverycharge `json:"deliverycharges"`
}

type Locationstatusinput struct {
	Locationstatus []*Locstatus `json:"locationstatus"`
	Deliverystatus []*Delstatus `json:"deliverystatus"`
}

type Locationupdate struct {
	Locationid   int    `json:"Locationid"`
	TenantID     int    `json:"TenantId"`
	LocationName string `json:"LocationName"`
	Email        string `json:"Email"`
	Contact      string `json:"Contact"`
	Address      string `json:"Address"`
	Suburb       string `json:"Suburb"`
	City         string `json:"City"`
	State        string `json:"State"`
	Zip          string `json:"Zip"`
	Countrycode  string `json:"Countrycode"`
	Latitude     string `json:"Latitude"`
	Longitude    string `json:"Longitude"`
	Openingtime  string `json:"Openingtime"`
	Closingtime  string `json:"Closingtime"`
	Delivery     bool   `json:"Delivery"`
	Deliverytype string `json:"Deliverytype"`
	Deliverymins int    `json:"Deliverymins"`
}

type Locstatus struct {
	Tenantid   int    `json:"tenantid"`
	Locationid int    `json:"locationid"`
	Status     string `json:"status"`
}

type Mod struct {
	Moduleid        int    `json:"Moduleid"`
	Categoryid      int    `json:"Categoryid"`
	Subcategoryid   int    `json:"Subcategoryid"`
	Subcategoryname string `json:"Subcategoryname"`
	Modulename      string `json:"Modulename"`
	Baseprice       string `json:"Baseprice"`
	Taxpercent      int    `json:"Taxpercent"`
	Taxamount       string `json:"Taxamount"`
	Amount          string `json:"Amount"`
	Content         string `json:"Content"`
	Logourl         string `json:"Logourl"`
	Iconurl         string `json:"Iconurl"`
}

type Othercharge struct {
	Tenantchargeid int    `json:"Tenantchargeid"`
	Tenantid       int    `json:"Tenantid"`
	Locationid     int    `json:"Locationid"`
	Chargeid       int    `json:"Chargeid"`
	Chargename     string `json:"Chargename"`
	Chargetype     string `json:"Chargetype"`
	Chargevalue    string `json:"Chargevalue"`
}

type Paymentdata struct {
	Paymentid       int                  `json:"Paymentid"`
	Moduleid        int                  `json:"Moduleid"`
	Locationid      int                  `json:"Locationid"`
	Tenantid        int                  `json:"Tenantid"`
	Paymentref      string               `json:"Paymentref"`
	Paymenttypeid   int                  `json:"Paymenttypeid"`
	Customerid      int                  `json:"Customerid"`
	Transactiondate string               `json:"Transactiondate"`
	Orderid         int                  `json:"Orderid"`
	Chargeid        string               `json:"Chargeid"`
	Amount          float64              `json:"Amount"`
	Refundamt       float64              `json:"Refundamt"`
	Paymentstatus   string               `json:"Paymentstatus"`
	Created         string               `json:"Created"`
	Paymentdetails  []*Paymentdetaildata `json:"Paymentdetails"`
}

type Paymentdetaildata struct {
	Paymentdetailid int       `json:"Paymentdetailid"`
	Paymentid       int       `json:"Paymentid"`
	Moduleid        int       `json:"Moduleid"`
	Locationid      int       `json:"Locationid"`
	Tenantid        int       `json:"Tenantid"`
	Orderid         int       `json:"Orderid"`
	Subscriptionid  int       `json:"Subscriptionid"`
	Amount          float64   `json:"Amount"`
	Taxpercent      int       `json:"Taxpercent"`
	Taxamount       float64   `json:"Taxamount"`
	Payamount       float64   `json:"Payamount"`
	Customerinfo    *Custinfo `json:"Customerinfo"`
}

type Promo struct {
	Promocodeid      int    `json:"Promocodeid"`
	Moduleid         int    `json:"Moduleid"`
	Partnerid        int    `json:"Partnerid"`
	Packageid        int    `json:"Packageid"`
	Promoname        string `json:"Promoname"`
	Promodescription string `json:"Promodescription"`
	Packageexpiry    string `json:"Packageexpiry"`
	Promotype        string `json:"Promotype"`
	Promovalue       string `json:"Promovalue"`
	Validity         string `json:"Validity"`
	Validitystatus   bool   `json:"Validitystatus"`
	Companyname      string `json:"Companyname"`
	Address          string `json:"Address"`
	City             string `json:"City"`
	Postcode         string `json:"Postcode"`
}

type Promoinput struct {
	Promotiontypeid int     `json:"Promotiontypeid"`
	Promotionname   *string `json:"Promotionname"`
	Tenantid        int     `json:"Tenantid"`
	Promocode       *string `json:"Promocode"`
	Promoterms      *string `json:"Promoterms"`
	Promovalue      *string `json:"Promovalue"`
	Startdate       *string `json:"Startdate"`
	Enddate         *string `json:"Enddate"`
}

type Promotioncreateddata struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Promotypesdata struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Types   []*Typedata `json:"types"`
}

type Result struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Socialadddata struct {
	Socialprofile *string `json:"socialprofile"`
	Dailcode      *string `json:"dailcode"`
	Sociallink    *string `json:"sociallink"`
	Socialicon    *string `json:"socialicon"`
}

type Socialinfo struct {
	Socialid      int    `json:"socialid"`
	Socialprofile string `json:"socialprofile"`
	Dailcode      string `json:"dailcode"`
	Sociallink    string `json:"sociallink"`
	Socialicon    string `json:"socialicon"`
}

type Socialupdatedata struct {
	Socialid      *int    `json:"socialid"`
	Socialprofile *string `json:"socialprofile"`
	Dailcode      *string `json:"dailcode"`
	Sociallink    *string `json:"sociallink"`
	Socialicon    *string `json:"socialicon"`
}

type Staffdetail struct {
	Staffdetailid   int            `json:"Staffdetailid"`
	Tenanatstaffid  int            `json:"Tenanatstaffid"`
	Tenantid        int            `json:"Tenantid"`
	Locationid      int            `json:"Locationid"`
	Locationdetails *Stafflocation `json:"Locationdetails"`
}

type Stafflocation struct {
	Locationid   int    `json:"Locationid"`
	Locationname string `json:"Locationname"`
	Email        string `json:"Email"`
	Contact      string `json:"Contact"`
	Address      string `json:"Address"`
	City         string `json:"City"`
	Postcode     string `json:"Postcode"`
}

type Subcat struct {
	Subcategoryname string `json:"Subcategoryname"`
	Subcategoryid   int    `json:"Subcategoryid"`
	Categoryid      int    `json:"Categoryid"`
	Status          string `json:"Status"`
	Icon            string `json:"Icon"`
}

type Subcatinsertdata struct {
	Tenantid        int    `json:"Tenantid"`
	Moduleid        int    `json:"Moduleid"`
	Categoryid      int    `json:"Categoryid"`
	Subcategoryid   int    `json:"Subcategoryid"`
	Subcategoryname string `json:"Subcategoryname"`
}

type SubscribedData struct {
	Status  bool          `json:"status"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Info    []*TenantData `json:"info"`
}

type SubscribedDataResponse struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Info    *TenantData `json:"info"`
}

type Subscribemoreinput struct {
	Subscriptionid  int    `json:"Subscriptionid"`
	Tenantid        int    `json:"Tenantid"`
	TransactionDate string `json:"TransactionDate"`
	Partnerid       int    `json:"Partnerid"`
	Currencyid      int    `json:"Currencyid"`
	Price           string `json:"Price"`
	TaxID           int    `json:"TaxId"`
	Quantity        int    `json:"Quantity"`
	Promoid         int    `json:"Promoid"`
	Promovalue      string `json:"Promovalue"`
	TaxAmount       string `json:"TaxAmount"`
	TotalAmount     string `json:"TotalAmount"`
	PaymentStatus   int    `json:"PaymentStatus"`
	Paymentid       *int   `json:"Paymentid"`
	Validitydate    string `json:"Validitydate"`
}

type Subscription struct {
	TransactionDate string `json:"TransactionDate"`
	PackageID       int    `json:"PackageId"`
	ModuleID        int    `json:"ModuleId"`
	CurrencyID      int    `json:"CurrencyId"`
	CurrencyCode    string `json:"CurrencyCode"`
	Price           string `json:"Price"`
	TaxID           int    `json:"TaxId"`
	Quantity        int    `json:"Quantity"`
	Promoid         int    `json:"Promoid"`
	Promovalue      string `json:"Promovalue"`
	TaxAmount       string `json:"TaxAmount"`
	TotalAmount     string `json:"TotalAmount"`
	PaymentStatus   int    `json:"PaymentStatus"`
	PaymentID       *int   `json:"PaymentId"`
	Validitydate    string `json:"Validitydate"`
}

type Subscriptionnew struct {
	Tenantid        int    `json:"Tenantid"`
	TransactionDate string `json:"TransactionDate"`
	Packageid       int    `json:"Packageid"`
	Partnerid       int    `json:"Partnerid"`
	Moduleid        int    `json:"Moduleid"`
	CategoryID      int    `json:"CategoryId"`
	SubCategoryID   int    `json:"SubCategoryId"`
	Subcategoryname string `json:"Subcategoryname"`
	Currencyid      int    `json:"Currencyid"`
	CurrencyCode    string `json:"CurrencyCode"`
	Price           string `json:"Price"`
	TaxID           int    `json:"TaxId"`
	Quantity        int    `json:"Quantity"`
	Promoid         int    `json:"Promoid"`
	Promovalue      string `json:"Promovalue"`
	TaxAmount       string `json:"TaxAmount"`
	TotalAmount     string `json:"TotalAmount"`
	PaymentStatus   int    `json:"PaymentStatus"`
	Paymentid       *int   `json:"Paymentid"`
	Validitydate    string `json:"Validitydate"`
}

type Subscriptionsdata struct {
	Subscriptionid       int      `json:"Subscriptionid"`
	Packageid            *int     `json:"Packageid"`
	Moduleid             int      `json:"Moduleid"`
	Tenantid             int      `json:"Tenantid"`
	Categoryid           int      `json:"Categoryid"`
	Subcategoryid        int      `json:"Subcategoryid"`
	Validitydate         string   `json:"Validitydate"`
	Validity             bool     `json:"Validity"`
	Modulename           string   `json:"Modulename"`
	Subscriptionaccid    string   `json:"Subscriptionaccid"`
	Subscriptionmethodid string   `json:"Subscriptionmethodid"`
	Paymentstatus        bool     `json:"Paymentstatus"`
	Packagename          *string  `json:"Packagename"`
	LogoURL              string   `json:"LogoUrl"`
	Iconurl              string   `json:"Iconurl"`
	PackageIcon          *string  `json:"PackageIcon"`
	PackageAmount        *float64 `json:"PackageAmount"`
	TotalAmount          float64  `json:"TotalAmount"`
	Taxamount            float64  `json:"Taxamount"`
	Tenantaccid          string   `json:"Tenantaccid"`
	Customercount        *int     `json:"Customercount"`
	Locationcount        *int     `json:"Locationcount"`
}

type Tenantsubcat struct {
	Categoryid      int    `json:"Categoryid"`
	Subcategoryid   int    `json:"Subcategoryid"`
	Subcategoryname string `json:"Subcategoryname"`
	Icon            string `json:"Icon"`
	Selected        int    `json:"Selected"`
	Categoryname    string `json:"Categoryname"`
}

type Tenantupdatedata struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Updated int    `json:"updated"`
}

type Tenantuser struct {
	Tenantid     int    `json:"Tenantid"`
	Moduleid     int    `json:"Moduleid"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Mobile       string `json:"mobile"`
	Email        string `json:"email"`
	Profileimage string `json:"profileimage"`
	Locationid   int    `json:"locationid"`
	Roleid       int    `json:"roleid"`
	Configid     int    `json:"configid"`
}

type Tenantuserdata struct {
	Status     bool   `json:"status"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Tenantuser *User  `json:"tenantuser"`
}

type Typedata struct {
	Promotiontypeid int     `json:"Promotiontypeid"`
	Typename        *string `json:"Typename"`
	Tag             *string `json:"Tag"`
}

type Updatedelivery struct {
	Create []*Deliverychargeinput       `json:"create"`
	Update []*Updatedeliverychargeinput `json:"update"`
	Delete []*int                       `json:"delete"`
}

type Updatedeliverychargeinput struct {
	Settingsid int    `json:"Settingsid"`
	Tenantid   int    `json:"Tenantid"`
	Locationid int    `json:"Locationid"`
	Slabtype   string `json:"Slabtype"`
	Slab       string `json:"Slab"`
	Slablimit  int    `json:"Slablimit"`
	Slabcharge string `json:"Slabcharge"`
}

type Updateinfo struct {
	Tenantid    int    `json:"Tenantid"`
	Locationid  int    `json:"Locationid"`
	Brandname   string `json:"Brandname"`
	About       string `json:"About"`
	Tenantimage string `json:"Tenantimage"`
	Openingtime string `json:"Openingtime"`
	Closingtime string `json:"Closingtime"`
}

type Updateother struct {
	Create []*Chargecreateinput `json:"create"`
	Update []*Chargeupdateinput `json:"update"`
	Delete []*int               `json:"delete"`
}

type Updatetenant struct {
	Userid       int    `json:"userid"`
	Tenantid     int    `json:"tenantid"`
	Moduleid     int    `json:"moduleid"`
	Locationid   int    `json:"locationid"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Mobile       string `json:"mobile"`
	Email        string `json:"email"`
	Profileimage string `json:"profileimage"`
}

type User struct {
	Userid int `json:"userid"`
}

type Userfromtenant struct {
	Tenantid     int    `json:"Tenantid"`
	Userid       int    `json:"Userid"`
	Firstname    string `json:"Firstname"`
	Lastname     string `json:"Lastname"`
	Email        string `json:"Email"`
	Contact      string `json:"Contact"`
	Profileimage string `json:"Profileimage"`
	Locationid   int    `json:"Locationid"`
	Locationname string `json:"Locationname"`
}

type Userinfodata struct {
	Profileid    int    `json:"Profileid"`
	Userid       int    `json:"Userid"`
	Locationid   int    `json:"Locationid"`
	Firstname    string `json:"Firstname"`
	Lastname     string `json:"Lastname"`
	Email        string `json:"Email"`
	Contact      string `json:"Contact"`
	Profileimage string `json:"Profileimage"`
}

type Userlist struct {
	Tenantstaffid int           `json:"Tenantstaffid"`
	Tenantid      int           `json:"Tenantid"`
	Moduleid      int           `json:"Moduleid"`
	Userid        int           `json:"Userid"`
	Userinfo      *Userinfodata `json:"Userinfo"`
}

type Usersdata struct {
	Status  bool              `json:"status"`
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Users   []*Userfromtenant `json:"users"`
}

type Usertenant struct {
	Staffdetailid  int         `json:"Staffdetailid"`
	Tenanatstaffid int         `json:"Tenanatstaffid"`
	Tenantid       int         `json:"Tenantid"`
	Locationid     int         `json:"Locationid"`
	Tenantusers    []*Userlist `json:"Tenantusers"`
}
