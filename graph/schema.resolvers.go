package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"net/http"
	"github.com/engajerest/auth/datacontext"
	"github.com/engajerest/sparkle/Models/subscription"
	"github.com/engajerest/sparkle/graph/generated"
	"github.com/engajerest/sparkle/graph/model"

)

func (r *mutationResolver) Subscribe(ctx context.Context, input model.Data) (*model.SubscribedData, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("userid=")
	print(id.ID)
	var auth subscription.AuthUser
	var d subscription.Initialsubscriptiondata
	var slist []subscription.TenantSubscription
	var s subscription.SubscribedData
	var list []*model.TenantData
	Subscribelist := *&input.Subscriptiondetails
	d.Userid = id.ID
	d.Categoryid = input.Tenantinfo.CategoryID
	d.Regno = input.Tenantinfo.Regno
	d.Name = input.Tenantinfo.Name
	d.Email = input.Tenantinfo.Email
	d.Mobile = input.Tenantinfo.Mobile
	d.SubCategoryid = input.Tenantinfo.SubCategoryID
	d.Tenanttoken = input.Tenantinfo.Tenanttoken
	d.Address = input.Tenantlocation.Address
	d.State = input.Tenantlocation.State
	d.Suburb = input.Tenantlocation.Suburb
	d.Zip = input.Tenantlocation.Zip
	d.Countrycode = input.Tenantlocation.Countrycode
	d.Latitude = input.Tenantlocation.Latitude
	d.Longitude = input.Tenantlocation.Longitude
	d.TimeZone = input.Tenantlocation.TimeZone
	d.OpenTime = input.Tenantlocation.Opentime
	d.CloseTime = input.Tenantlocation.Closetime
	d.Subcategoryname = input.Tenantinfo.Subcategoryname
	d.Partnerid = input.Subscriptiondetails[0].Partnerid
	if len(Subscribelist) != 0 {
		for _, k := range Subscribelist {
			slist = append(slist, subscription.TenantSubscription{
				Date: k.TransactionDate, Packageid: k.Packageid, Partnerid: k.Partnerid, Moduleid: k.Moduleid,
				Currencyid: k.Currencyid, Categoryid: input.Tenantinfo.CategoryID, SubCategoryid: input.Tenantinfo.SubCategoryID,
				Price: k.Price, TaxId: k.TaxID, TaxAmount: k.TaxAmount, TotalAmount: k.TotalAmount, PaymentStatus: k.PaymentStatus,
				PaymentId: *k.Paymentid, Quantity: k.Quantity, Promoid: k.Promoid, Promovalue: k.Promovalue, Validitydate: k.Validitydate, Promostatus: true,
			})
		}
	}

	d.Tenantsubscribe = slist

	status, tenantdata, err := d.Subscriptioninitial()
	if err != nil {
		return nil, err
	}
	if status == false {
		return nil, err
	}
	if tenantdata.TenantID != 0 {
		var seq subscription.Ordersequence
		seq.Tenantid = int(tenantdata.TenantID)
		seq.Tablename = "order"
		seq.Seqno = 0
		seq.Prefix = "ORD"
		seq.Subprefix = 2021
		seqid, err := seq.Insertsequence()
		if err != nil {
			print(err)
			print("seqid==", seqid)
		}
		seqid1, err := seq.Insertpaysequence()
		if err != nil {
			print(err)
			print("seqid==", seqid1)
		}
		seqid2, err := seq.Insertcustomersequence()
		if err != nil {
			print(err)
			print("seqid2==", seqid2)
		}
		auth.LocationId = tenantdata.Locationid
		auth.TenantID = tenantdata.TenantID
		status := auth.UpdateAuthUser(id.ID)
		print(status)
		if status != true {
			return nil, errors.New("tenant not subscribed")
		}

	} else {
		return nil, errors.New("Tenant not Created")
	}

	response := s.GetSubscribedData(int64(tenantdata.TenantID))
	if len(response) != 0 {
		for _, k := range response {
			list = append(list, &model.TenantData{TenantID: k.TenantID, TenantName: k.TenantName, Moduleid: k.ModuleID,
				Categoryid: k.Categoryid, Subcategoryid: k.Subcategoryid, Locationid: k.Locationid, Locationname: k.Locationname, Modulename: k.ModuleName, Subscriptionid: k.Subscriptionid})
		}
	}
	return &model.SubscribedData{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		Info:    list,
	}, nil
}

func (r *mutationResolver) Createtenantuser(ctx context.Context, create *model.Tenantuser) (*model.Tenantuserdata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("tenantuser")
	print(id.ID)
	var user subscription.TenantUser
	user.FirstName = create.Firstname
	user.LastName = create.Lastname
	user.Password = create.Password
	user.Email = create.Email
	user.Mobile = create.Mobile
	user.RoleId = create.Roleid
	user.Locationid = create.Locationid
	user.TenantID = create.TenantID
	tenantuserid := user.CreateTenantUser()
	if tenantuserid != 0 {
		tenantprofileid := user.InsertTenantUserintoProfile(tenantuserid)
		print(tenantprofileid)
	}

	return &model.Tenantuserdata{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		Tenantuser: &model.User{
			Userid: int(tenantuserid),
		},
	}, nil
}

func (r *mutationResolver) Updatetenantuser(ctx context.Context, update *model.Updatetenant) (*model.Tenantupdatedata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("userid=")
	print(id.ID)
	var data subscription.TenantUser
	data.Userid = update.Userid
	data.TenantID = update.Tenantid
	data.FirstName = update.Firstname
	data.LastName = update.Lastname
	data.Email = update.Email
	data.Mobile = update.Mobile
	data.Locationid = update.Locationid
	data1 := data.UpdateTenantUser()
	if data1 != false {
		return &model.Tenantupdatedata{
			Status:  true,
			Code:    http.StatusOK,
			Message: "Success",
			Updated: 1,
		}, nil
	}

	return &model.Tenantupdatedata{
		Status:  false,
		Code:    http.StatusBadRequest,
		Message: "failure",
		Updated: 0,
	}, nil
}

func (r *mutationResolver) Updatetenantbusiness(ctx context.Context, businessinfo *model.Business) (*model.Businessdata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("userid=")
	print(id.ID)

	var data subscription.BusinessUpdate
	data.TenantID = businessinfo.Businessupdate.Tenantid

	data.Brandname = *businessinfo.Businessupdate.Brandname
	data.About = *businessinfo.Businessupdate.About
	data.Paymode1 = *businessinfo.Businessupdate.Cod
	data.Paymode2 = *businessinfo.Businessupdate.Digital
	data.Tenantimage = businessinfo.Businessupdate.Tenantimage
	var check []subscription.Social
	var updatedata []subscription.Social
	schemasocialadd := *&businessinfo.Socialadd
	schemasocialupdate := *&businessinfo.Socialupdate
	socialdelete := *&businessinfo.Socialdelete
	data1 := data.UpdateTenantBusiness()
	for _, v := range schemasocialadd {
		check = append(check, subscription.Social{SociaProfile: *v.Socialprofile, SocialLink: *v.Sociallink, SocialIcon: *v.Socialicon})
	}
	if len(check) != 0 {
		social := data.InsertTenantSocial(check, data.TenantID)
		print(social)
	}
	for _, k := range schemasocialupdate {
		updatedata = append(updatedata, subscription.Social{Socialid: *k.Socialid, SociaProfile: *k.Socialprofile, SocialLink: *k.Sociallink, SocialIcon: *k.Socialicon})
	}
	if len(updatedata) != 0 {
		var s subscription.Social
		for i := 0; i < len(updatedata); i++ {
			s.Socialid = updatedata[i].Socialid
			s.SociaProfile = updatedata[i].SociaProfile
			s.SocialIcon = updatedata[i].SocialIcon
			s.SocialLink = updatedata[i].SocialLink
			status := s.UpdateTenantSocial(data.TenantID)
			if status == false {
				return nil, errors.New("error in updating socialinfo")
			}
		}

	}
	if len(socialdelete) != 0 {
		for i := 0; i < len(socialdelete); i++ {
			var d subscription.Social
			d.Socialid = *socialdelete[i]
			stat := d.Deletesocial()
			print(stat)
		}
	}
	if data1 != false {
		return &model.Businessdata{
			Status:  true,
			Code:    http.StatusOK,
			Message: "Success",
			Updated: 1,
		}, nil
	}

	return &model.Businessdata{
		Status:  false,
		Code:    http.StatusBadRequest,
		Message: "failure",
		Updated: 0,
	}, nil
}

func (r *mutationResolver) Createlocation(ctx context.Context, input *model.Location) (*model.Locationdata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr :=datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("userid=")
	print(id.ID)
	var loco subscription.Location
	loco.TenantID = input.TenantID
	loco.LocationName = input.LocationName
	loco.Email = input.Email
	loco.Mobile = input.Contact
	loco.Address = input.Address
	loco.Suburb = input.Suburb
	loco.State = input.State
	loco.Zip = input.Zip
	loco.Countrycode = input.Countrycode
	loco.Latitude = input.Latitude
	loco.Longitude = input.Longitude
	loco.OpeningTime = input.Openingtime
	loco.ClosingTime = input.Closingtime
	loco.Delivery = input.Delivery
	loco.Deliverytype = input.Deliverytype
	loco.Deliverymins = input.Deliverymins
	locationid, er := loco.CreateLocation(int64(id.ID))
	if er != nil {
		return nil, errors.New("location not created")
	}

	location, errr := loco.GetLocationById(locationid)
	if errr != nil {
		return nil, errors.New("location not found")
	}
	return &model.Locationdata{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		Locationinfo: &model.LocationInfo{
			Locationid:   location.LocationId,
			LocationName: location.LocationName,
			Status:       location.Status,
			Createdby:    location.Createdby,
		},
	}, nil
}

func (r *mutationResolver) Createpromotion(ctx context.Context, input *model.Promoinput) (*model.Promotioncreateddata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("userid=")
	print(id.ID)
	var p subscription.Promotion
	p.Promotiontypeid = input.Promotiontypeid
	p.Tenantid = input.Tenantid
	p.Promoname = *input.Promotionname
	p.Promoterms = *input.Promoterms
	p.Promocode = *input.Promocode
	p.Promovalue = *input.Promovalue
	p.Startdate = *input.Startdate
	p.Enddate = *input.Enddate

	promoid := p.Createpromotion(id.ID)
	if promoid == 0 {
		return nil, errors.New("error occurs in promotion")
	}
	return &model.Promotioncreateddata{Status: true, Code: http.StatusCreated, Message: "Promotion created Successfully"}, nil
}

func (r *mutationResolver) Createcharges(ctx context.Context, input *model.Chargecreate) (*model.Promotioncreateddata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr :=datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("raju")
	print(id.ID)
	var c []subscription.Tenantcharge
	var d []subscription.Tenantsetting
	var other subscription.Tenantcharge
	var delivery subscription.Tenantsetting
	othercharge := *&input.Othercharges
	deliverycharges := *&input.Deliverycharges
	for _, k := range othercharge {
		c = append(c, subscription.Tenantcharge{Tenantid: k.Tenantid, Locationid: k.Locationid, Chargeid: k.Chargeid, Chargename: k.Chargename, Chargetype: k.Chargetype, Chargevalue: k.Chargevalue, Createdby: id.ID})
	}
	if len(c) != 0 {
		er := other.Insertothercharges(c)
		if er != nil {
			return nil, er
		}
	}
	for _, j := range deliverycharges {
		d = append(d, subscription.Tenantsetting{Tenantid: j.Tenantid, Locationid: j.Locationid, Slabtype: j.Slabtype, Slab: j.Slab, Slablimit: j.Slablimit, Slabcharge: j.Slabcharge, Createdby: id.ID})
	}
	if len(d) != 0 {
		erd := delivery.Insertdeliverycharges(d)
		if erd != nil {
			return nil, erd
		}
	}

	return &model.Promotioncreateddata{Status: true, Code: http.StatusCreated, Message: "charges created successfully"}, nil
}

func (r *mutationResolver) Updatecharges(ctx context.Context, input *model.Chargeupdate) (*model.Promotioncreateddata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("raju")
	print(id.ID)
	othercreate := input.Updateothercharges.Create
	otherupdate := input.Updateothercharges.Update
	otherdelete := input.Updateothercharges.Delete
	deliverycreate := input.Updatedeliverycharges.Create
	deliveryupdate := input.Updatedeliverycharges.Update
	deliverydelete := input.Updatedeliverycharges.Delete
	var other []subscription.Tenantcharge
	var del []subscription.Tenantsetting
	var o subscription.Tenantcharge
	var d subscription.Tenantsetting
	if len(otherdelete) != 0 {
		for i := 0; i < len(otherdelete); i++ {
			o.Tenantchargeid = *otherdelete[i]
			stat := o.Deleteothercharge()
			print(stat)

		}
	}
	if len(otherupdate) != 0 {
		for i := 0; i < len(otherupdate); i++ {
			o.Chargeid = otherupdate[i].Chargeid
			o.Chargename = otherupdate[i].Chargename
			o.Chargetype = otherupdate[i].Chargetype
			o.Chargevalue = otherupdate[i].Chargevalue
			o.Locationid = otherupdate[i].Locationid
			o.Tenantchargeid = otherupdate[i].Tenantchargeid
			o.Tenantid = otherupdate[i].Tenantid
			st := o.Updateothercharge()
			print(st)
		}
	}
	if len(othercreate) != 0 {
		for _, k := range othercreate {
			other = append(other, subscription.Tenantcharge{Tenantid: k.Tenantid, Locationid: k.Locationid, Chargeid: k.Chargeid, Chargename: k.Chargename, Chargetype: k.Chargetype, Chargevalue: k.Chargevalue, Createdby: id.ID})
		}
		er := o.Insertothercharges(other)
		if er != nil {
			return nil, er
		}

	}

	if len(deliverydelete) != 0 {
		for i := 0; i < len(deliverydelete); i++ {
			d.Settingsid = *deliverydelete[i]
			stat := d.Deletedeliverycharge()
			print(stat)

		}
	}
	if len(deliveryupdate) != 0 {
		for i := 0; i < len(deliveryupdate); i++ {
			d.Locationid = deliveryupdate[i].Locationid
			d.Settingsid = deliveryupdate[i].Settingsid
			d.Slab = deliveryupdate[i].Slab
			d.Slabcharge = deliveryupdate[i].Slabcharge
			d.Slablimit = deliveryupdate[i].Slablimit
			d.Slabtype = deliveryupdate[i].Slabtype
			d.Tenantid = deliveryupdate[i].Tenantid
			s := d.Updatedeliverycharge()
			print(s)

		}
	}
	if len(deliverycreate) != 0 {
		for _, j := range deliverycreate {
			del = append(del, subscription.Tenantsetting{Tenantid: j.Tenantid, Locationid: j.Locationid, Slabtype: j.Slabtype, Slab: j.Slab, Slablimit: j.Slablimit, Slabcharge: j.Slabcharge, Createdby: id.ID})
		}

		erd := d.Insertdeliverycharges(del)
		if erd != nil {
			return nil, erd
		}

	}
	return &model.Promotioncreateddata{Status: true, Code: http.StatusCreated, Message: "Charges updated"}, nil
}

func (r *mutationResolver) Updatelocationstatus(ctx context.Context, input *model.Locationstatusinput) (*model.Promotioncreateddata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("raju")
	print(id.ID)
	loc := *&input.Locationstatus
	del := input.Deliverystatus
	var s subscription.Updatestatus
	if len(loc) != 0 {

		for i := 0; i < len(loc); i++ {

			s.Locationstatus = *&loc[i].Status
			s.Locationid = *&loc[i].Locationid
			s.Tenantid = *&loc[i].Tenantid
			stat := s.Updatelocationstatus()
			if stat == false {
				return nil, errors.New("location status not updated")
			}
		}
	}
	if len(del) != 0 {
		for i := 0; i < len(del); i++ {
			s.Deliverystatus = *&del[i].Delivery
			s.Locationid = *&del[i].Locationid
			s.Tenantid = *&del[i].Tenantid
			stat1 := s.Updatedeliverystatus()
			if stat1 == false {
				return nil, errors.New("delivery not updated")
			}

		}
	}
	return &model.Promotioncreateddata{Status: true, Code: http.StatusCreated, Message: "status updated"}, nil
}

func (r *mutationResolver) Updatelocation(ctx context.Context, input *model.Locationupdate) (*model.Promotioncreateddata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("update loc")
	print(id.ID)
	var loco subscription.Location
	loco.LocationId = input.Locationid
	loco.TenantID = input.TenantID
	loco.LocationName = input.LocationName
	loco.Email = input.Email
	loco.Mobile = input.Contact
	loco.Address = input.Address
	loco.Suburb = input.Suburb
	loco.State = input.State
	loco.Zip = input.Zip
	loco.Countrycode = input.Countrycode
	loco.Latitude = input.Latitude
	loco.Longitude = input.Longitude
	loco.OpeningTime = input.Openingtime
	loco.ClosingTime = input.Closingtime
	loco.Delivery = input.Delivery
	loco.Deliverytype = input.Deliverytype
	loco.Deliverymins = input.Deliverymins
	status, er := loco.UpdateLocation()
	if er != nil || status == false {
		return nil, errors.New("location not created")
	}

	return &model.Promotioncreateddata{Status: true, Code: http.StatusCreated, Message: "Location updated"}, nil
}

func (r *mutationResolver) Subscription(ctx context.Context, input []*model.Subscriptionnew) (*model.SubscribedData, error) {
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("update loc")
	print(id.ID)
	var data2 subscription.SubscribedData
	var data1 subscription.TenantSubscription
	var list []*model.TenantData
	var list1 []subscription.SubscribedData
	intlist := input
	if len(intlist) != 0 {
		for i := 0; i < len(intlist); i++ {

			data1.Currencyid = intlist[i].Currencyid
			data1.Date = intlist[i].TransactionDate
			data1.Partnerid = intlist[i].Partnerid
			data1.Categoryid = intlist[i].CategoryID
			data1.SubCategoryid = intlist[i].SubCategoryID

			data1.Moduleid = intlist[i].Moduleid
			data1.PaymentId = *intlist[i].Paymentid
			data1.PaymentStatus = intlist[i].PaymentStatus
			data1.Price = intlist[i].Price
			data1.Quantity = intlist[i].Quantity
			data1.TaxId = intlist[i].TaxID
			data1.TaxAmount = intlist[i].TaxAmount
			data1.TotalAmount = intlist[i].TotalAmount
			data1.Packageid = intlist[i].Packageid
			data1.Promoid = intlist[i].Promoid
			data1.Promovalue = intlist[i].Promovalue
			data1.Validitydate = intlist[i].Validitydate
			data1.Promostatus = true
			subscribedid := data1.InsertSubscription(int64(intlist[i].Tenantid))
			print("subs-id===")
			print(subscribedid)
			print(intlist[i].Tenantid)
		}
	}
	var d subscription.TenantSubscription
	d.Categoryid = intlist[0].CategoryID
	d.SubCategoryid = intlist[0].SubCategoryID
	d.Moduleid = intlist[0].Moduleid
	d.Tenantid = intlist[0].Tenantid
	d.Subcategoryname = intlist[0].Subcategoryname
	_, er := d.Insertsubcategory()
	if er != nil {
		print(er)
	}

	list1 = data2.GetSubscribedData(int64(intlist[0].Tenantid))
	if len(list1) != 0 {
		for _, k := range list1 {
			list = append(list, &model.TenantData{TenantID: k.TenantID, TenantName: k.TenantName, Moduleid: k.ModuleID,
				Categoryid: k.Categoryid, Subcategoryid: k.Subcategoryid, Locationid: k.Locationid, Locationname: k.Locationname, Modulename: k.ModuleName, Subscriptionid: k.Subscriptionid})
		}
	}

	return &model.SubscribedData{Status: true, Code: http.StatusCreated, Message: "Success",
		Info: list,
	}, nil
}

func (r *mutationResolver) Initialupdate(ctx context.Context, input *model.Updateinfo) (*model.Promotioncreateddata, error) {
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("raju")
	print(id.ID)
	var d subscription.Initial
	d.Tenantid = input.Tenantid
	d.Locationid = input.Locationid
	d.Brandname = input.Brandname
	d.Tenantimage = input.Tenantimage
	d.About = input.About
	d.Opentime = input.Openingtime
	d.Closetime = input.Closingtime
	stat, err := d.Initialupdate()
	if err != nil {
		return nil, err
	}
	if stat != true {
		print(stat)
	}

	return &model.Promotioncreateddata{Status: true, Code: http.StatusCreated, Message: "BusinessInfo Updated"}, nil
}

func (r *mutationResolver) Insertsubcategory(ctx context.Context, input []*model.Subcatinsertdata) (*model.Promotioncreateddata, error) {
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("raju")
	print(id.ID)
	var d subscription.TenantSubscription
	sublist := input
	if len(sublist) != 0 {
		for i := 0; i < len(sublist); i++ {
			d.Categoryid = sublist[i].Categoryid
			d.SubCategoryid = sublist[i].Subcategoryid
			d.Moduleid = sublist[i].Moduleid
			d.Tenantid = sublist[i].Tenantid
			d.Subcategoryname = sublist[i].Subcategoryname
			_, err := d.Insertsubcategory()
			if err != nil {
				print(err)
			}

		}
	}

	return &model.Promotioncreateddata{
		Status: true, Code: http.StatusCreated, Message: "Subcategories Added to Tenants",
	}, nil
}

func (r *queryResolver) Sparkle(ctx context.Context) (*model.Sparkle, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("raju")
	print(id.ID)

	var cat []*model.Category
	var sub []*model.SubCategory
	var pack []*model.Package
	var categoryGetAll []subscription.Category
	var subcategoryGetAll []subscription.SubCategory
	var packageGetAll []subscription.Packages

	categoryGetAll = subscription.GetAllCategory()
	for _, category := range categoryGetAll {
		cat = append(cat, &model.Category{Categoryid: category.CategoryID, Name: category.Name, Type: category.Typeid, SortOrder: category.SortOrder, Status: category.Status})

	}
	subcategoryGetAll = subscription.GetAllSubCategory()
	for _, subcategory := range subcategoryGetAll {
		sub = append(sub, &model.SubCategory{CategoryID: subcategory.CategoryID, SubCategoryID: subcategory.SubCategoryID, Name: subcategory.Name, Type: subcategory.Typeid, SortOrder: subcategory.SortOrder, Status: subcategory.Status, Icon: subcategory.Icon})
	}
	packageGetAll = subscription.GetAllPackages()
	for _, packdata := range packageGetAll {
		pack = append(pack, &model.Package{ModuleID: packdata.ModuleID, Modulename: packdata.ModuleName, Name: packdata.Name, PackageID: packdata.PackageID, Status: packdata.Status, PackageAmount: packdata.PackageAmount, PaymentMode: packdata.PaymentMode, PackageContent: packdata.PackageContent, PackageIcon: packdata.PackageIcon,
			Promocodeid: packdata.Promocodeid, Promonname: packdata.Promoname, Promodescription: packdata.Promodescription, Promotype: packdata.Promotype, Promovalue: packdata.Promovalue, Validitydate: packdata.Promovaliditydate, Validity: packdata.Validity, Packageexpiry: packdata.Packageexpiry})

	}

	return &model.Sparkle{
		Category:    cat,
		Subcategory: sub,
		Package:     pack,
	}, nil
}

func (r *queryResolver) Location(ctx context.Context, tenantid int) (*model.Getalllocations, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("getalloc")
	print(id.ID)

	var Result []*model.Locationgetall
	var userresult []*model.Usertenant
	var otherchargeresult []*model.Othercharge
	var deliverychargeresult []*model.Deliverycharge
	var locationGetAll []subscription.Tenantlocation

	locationGetAll = subscription.LocationTest(tenantid)

	for _, loco := range locationGetAll {
		userresult = make([]*model.Usertenant, len(loco.Appuserprofiles))
		for i, key := range loco.Appuserprofiles {
			userresult[i] = &model.Usertenant{
				Userid:         key.Userid,
				Userlocationid: key.Userlocationid,
				Firstname:      key.Firstname,
				Lastname:       key.Lastname,
				Mobile:         key.Contactno,
				Email:          key.Email,
			}
		}
		otherchargeresult = make([]*model.Othercharge, len(loco.Tenantcharges))
		for j, k := range loco.Tenantcharges {
			otherchargeresult[j] = &model.Othercharge{Tenantchargeid: k.Tenantchargeid, Tenantid: k.Tenantid, Locationid: k.Locationid,
				Chargeid: k.Chargeid, Chargename: k.Chargename, Chargetype: k.Chargetype, Chargevalue: k.Chargevalue}
		}
		deliverychargeresult = make([]*model.Deliverycharge, len(loco.Tenantsettings))
		for l, m := range loco.Tenantsettings {
			deliverychargeresult[l] = &model.Deliverycharge{Settingsid: m.Settingsid, Tenantid: m.Tenantid, Locationid: m.Locationid,
				Slabtype: m.Slabtype, Slab: m.Slab, Slablimit: m.Slablimit, Slabcharge: m.Slabcharge}
		}
		Result = append(Result, &model.Locationgetall{
			Locationid:      loco.Locationid,
			LocationName:    loco.Locationname,
			Tenantid:        loco.Tenantid,
			Email:           loco.Email,
			Contact:         loco.Contactno,
			Address:         loco.Address,
			Suburb:          loco.City,
			State:           loco.State,
			Countycode:      loco.Countrycode,
			Postcode:        loco.Postcode,
			Latitude:        loco.Latitude,
			Longitude:       loco.Longitude,
			Openingtime:     loco.Opentime,
			Closingtime:     loco.Closetime,
			Status:          loco.Status,
			Createdby:       loco.Createdby,
			Delivery:        loco.Delivery,
			Deliverytype:    loco.Deliverytype,
			Deliverymins:    loco.Deliverymins,
			Tenantusers:     userresult,
			Othercharges:    otherchargeresult,
			Deliverycharges: deliverychargeresult,
		})

	}

	return &model.Getalllocations{
		Status:    true,
		Code:      http.StatusOK,
		Message:   "Success",
		Locations: Result,
	}, nil
}

func (r *queryResolver) Tenantusers(ctx context.Context, tenantid int) (*model.Usersdata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("gettenantusers")
	print(id.ID)
	tId := tenantid
	if tId == 0 {
		return nil, errors.New("tenantid must not be 0")
	}
	var Result []*model.Userfromtenant
	var tenantusersGetAll []subscription.TenantUser
	tenantusersGetAll = subscription.GetAllTenantUsers(tId)
	for _, user := range tenantusersGetAll {
		Result = append(Result, &model.Userfromtenant{
			UserID:       user.Userid,
			Locationid:   user.Locationid,
			Tenantid:     user.Referenceid,
			LocationName: user.Locationname,
			Firstname:    user.FirstName,
			Lastname:     user.LastName,
			Mobile:       user.Mobile,
			Email:        user.Email,
			Status:       user.Status,
			Created:      user.Created,
		})
	}
	return &model.Usersdata{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		Users:   Result,
	}, nil
}

func (r *queryResolver) GetBusiness(ctx context.Context, tenantid int, categoryid int) (*model.GetBusinessdata, error) {
	// id, usererr := helper.ForSparkleContext(ctx)
	id, usererr :=datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("getbusin")
	print(id.ID)
	var Result []*model.Socialinfo
	var businessinfo *subscription.BusinessUpdate
	var stat bool
	var socialgetall []subscription.Social
	socialgetall = subscription.GetAllSocial(tenantid)
	for _, user := range socialgetall {
		Result = append(Result, &model.Socialinfo{
			Socialid:      user.Socialid,
			Socialprofile: user.SociaProfile,
			Sociallink:    user.SocialLink,
			Socialicon:    user.SocialIcon,
		})
	}
	if categoryid == 0 {
		print("cat0")
		businessinfo, stat = businessinfo.GetBusinessInfo(tenantid)
		print(stat)

	} else {
		print("cat!=0")
		businessinfo, stat = businessinfo.GetBusinessforassist(tenantid, categoryid)
	}

	return &model.GetBusinessdata{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		Businessinfo: &model.Info{
			Tenantid:    businessinfo.TenantID,
			Brandname:   &businessinfo.Brandname,
			About:       &businessinfo.About,
			Cod:         &businessinfo.Paymode1,
			Digital:     &businessinfo.Paymode2,
			Tenantaccid: &businessinfo.TenantaccId,
			Email:       &businessinfo.Email,
			Phone:       &businessinfo.Phone,
			Address:     &businessinfo.Address,
			Moduleid:    businessinfo.Moduleid,
			Modulename:  businessinfo.Modulename,
			Tenanttoken: &businessinfo.Tenanttoken,
			Tenantimage: &businessinfo.Tenantimage,
			Social:      Result,
		},
	}, nil
}

func (r *queryResolver) Getpromotions(ctx context.Context, tenantid int) (*model.Getpromotiondata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr :=datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("promo")
	print(id.ID)
	var promo []*model.Promotion
	var promotionGetAll []subscription.Promotion
	promotionGetAll = subscription.GetAllPromotions(tenantid)
	for _, p := range promotionGetAll {
		promo = append(promo, &model.Promotion{
			Promotionid: p.Promotionid, Promotiontypeid: p.Promotiontypeid, Promotionname: p.Promoname, Tenantid: p.Tenantid, Tenantame: p.Tenantname, Promocode: p.Promocode,
			Promoterms: p.Promoterms, Promovalue: p.Promovalue, Promotag: p.Promotag, Promotype: p.Promotype, Startdate: p.Startdate, Enddate: p.Enddate, Status: &p.Status,
		})
	}
	return &model.Getpromotiondata{
		Status: true, Code: http.StatusOK, Message: "Success", Promotions: promo,
	}, nil
}

func (r *queryResolver) Getpromotypes(ctx context.Context) (*model.Promotypesdata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("promotypes")
	print(id.ID)
	var data []*model.Typedata
	data = subscription.Getpromotypes()
	return &model.Promotypesdata{
		Status: true, Code: http.StatusOK, Message: "Success", Types: data,
	}, nil
}

func (r *queryResolver) Getchargetypes(ctx context.Context) (*model.Chargetypedata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("userid==")
	print(id.ID)
	var data []*model.Chargetype
	data = subscription.Getchargetypes()

	return &model.Chargetypedata{Status: true, Code: http.StatusOK, Message: "Success", Types: data}, nil
}

func (r *queryResolver) Getlocationbyid(ctx context.Context, tenantid int, locationid int) (*model.Locationbyiddata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("userid==")
	print(id.ID)

	var userresult []*model.Usertenant
	var otherchargeresult []*model.Othercharge
	var deliverychargeresult []*model.Deliverycharge
	loco := subscription.Locationbyid(tenantid, locationid)
	if loco.Locationid == 0 {
		return &model.Locationbyiddata{Status: false, Code: http.StatusBadRequest, Message: "Unsuccess", Locationdata: nil}, nil
	}
	if len(loco.Appuserprofiles) != 0 {
		userresult = make([]*model.Usertenant, len(loco.Appuserprofiles))
		for i, key := range loco.Appuserprofiles {
			userresult[i] = &model.Usertenant{
				Userid:         key.Userid,
				Userlocationid: key.Userlocationid,
				Firstname:      key.Firstname,
				Lastname:       key.Lastname,
				Mobile:         key.Contactno,
				Email:          key.Email,
			}
		}
	}
	if len(loco.Tenantcharges) != 0 {
		otherchargeresult = make([]*model.Othercharge, len(loco.Tenantcharges))
		for j, k := range loco.Tenantcharges {
			otherchargeresult[j] = &model.Othercharge{Tenantchargeid: k.Tenantchargeid, Tenantid: k.Tenantid, Locationid: k.Locationid,
				Chargeid: k.Chargeid, Chargename: k.Chargename, Chargetype: k.Chargetype, Chargevalue: k.Chargevalue}
		}
	}
	if len(loco.Tenantsettings) != 0 {
		deliverychargeresult = make([]*model.Deliverycharge, len(loco.Tenantsettings))
		for l, m := range loco.Tenantsettings {
			deliverychargeresult[l] = &model.Deliverycharge{Settingsid: m.Settingsid, Tenantid: m.Tenantid, Locationid: m.Locationid,
				Slabtype: m.Slabtype, Slab: m.Slab, Slablimit: m.Slablimit, Slabcharge: m.Slabcharge}
		}

	}
	return &model.Locationbyiddata{
		Status: true,
		Code:   http.StatusOK, Message: "Success", Locationdata: &model.Locationgetall{
			Locationid: loco.Locationid, LocationName: loco.Locationname, Tenantid: loco.Tenantid, Email: loco.Email, Contact: loco.Contactno,
			Address: loco.Address, Suburb: loco.City, State: loco.State, Postcode: loco.Postcode, Countycode: loco.Countrycode, Latitude: loco.Latitude,
			Delivery: loco.Delivery, Deliverytype: loco.Deliverytype, Deliverymins: loco.Deliverymins,
			Longitude: loco.Longitude, Openingtime: loco.Opentime, Closingtime: loco.Closetime, Status: loco.Status, Tenantusers: userresult, Createdby: loco.Createdby,
			Othercharges: otherchargeresult, Deliverycharges: deliverychargeresult,
		},
	}, nil
}

func (r *queryResolver) Getpayments(ctx context.Context, tenantid int, typeid int) (*model.Getpaymentdata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("userid==")
	print(id.ID)

	var data []*model.Paymentdata
	var d []subscription.Payment
	d = subscription.Payments(tenantid, typeid)
	if len(d) != 0 {
		for _, k := range d {
			data = append(data, &model.Paymentdata{Paymentid: k.Paymentid, Packageid: k.Packageid, Tenantid: k.Tenantid,
				Paymenttypeid: k.Paymenttypeid, Customerid: k.Customerid, Transactiondate: k.Transactiondate, Orderid: &k.Orderid,
				Chargeid: k.Chargeid, Amount: k.Amount, Refundamt: &k.Refundamt, Paymentstatus: &k.Paymentstatus,
				Created: &k.Created, Packagename: k.Packagename, Firstname: k.Firstname, Lastname: k.Lastname,
				Email: k.Email, Contact: k.Contactno, Paymentref: k.Paymentref})
		}
	}
	return &model.Getpaymentdata{Status: true, Code: http.StatusOK, Message: "Success", Payments: data}, nil
}

func (r *queryResolver) Getsubscriptions(ctx context.Context, tenantid int) (*model.Getsubscriptionsdata, error) {
	// id, usererr := controller.ForContext(ctx)
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("userid==")
	print(id.ID)
	var data []*model.Subscriptionsdata
	var d []subscription.Subscribe
	d = subscription.GetAllSubscription(tenantid)
	if len(d) != 0 {
		for _, k := range d {
			data = append(data, &model.Subscriptionsdata{Packageid: &k.Packageid, Moduleid: &k.Moduleid, Tenantid: &k.Tenantid, Modulename: &k.Modulename, Packagename: &k.Packagename,
				LogoURL: &k.Logourl, PackageIcon: &k.PackageIcon, PackageAmount: &k.PackageAmount, TotalAmount: &k.Totalamount, Customercount: &k.Customercount, Locationcount: &k.Locationcount})
		}
	}

	return &model.Getsubscriptionsdata{Status: true, Code: http.StatusOK, Message: "Success", Subscribed: data}, nil
}

func (r *queryResolver) Getnonsubscribed(ctx context.Context, tenantid int) (*model.Getnonsubscribeddata, error) {
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("userid==")
	print(id.ID)
	var pack []*model.Package
	var packageGetAll []subscription.Packages
	packageGetAll = subscription.Getallnonsubscribedpackages(tenantid)
	for _, packdata := range packageGetAll {
		pack = append(pack, &model.Package{ModuleID: packdata.ModuleID, Modulename: packdata.ModuleName, Name: packdata.Name, PackageID: packdata.PackageID, Status: packdata.Status, PackageAmount: packdata.PackageAmount, PaymentMode: packdata.PaymentMode, PackageContent: packdata.PackageContent, PackageIcon: packdata.PackageIcon,
			Packageexpiry: packdata.Packageexpiry, Promocodeid: packdata.Promocodeid, Promonname: packdata.Promoname, Promodescription: packdata.Promodescription, Promotype: packdata.Promotype, Promovalue: packdata.Promovalue, Validitydate: packdata.Promovaliditydate, Validity: packdata.Validity})
	}
	return &model.Getnonsubscribeddata{Status: true, Code: http.StatusCreated, Message: "Success", Nonsubscribed: pack}, nil
}

func (r *queryResolver) Getallmodule(ctx context.Context, categoryid int, tenantid int) (*model.Getallmoduledata, error) {
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("userid==")
	print(id.ID)
	var mods []*model.Mod

	mods = subscription.Getmodules(categoryid, tenantid)
	return &model.Getallmoduledata{Status: true, Code: http.StatusOK, Message: "Success", Modules: mods}, nil
}

func (r *queryResolver) Getallpromos(ctx context.Context, moduleid int) (*model.Getallpromodata, error) {
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("userid==")
	print(id.ID)
	var data []*model.Promo
	data = subscription.Getpromos(moduleid)

	return &model.Getallpromodata{Status: true, Code: http.StatusOK, Message: "Success", Promos: data}, nil
}

func (r *queryResolver) Getsubcategorybyid(ctx context.Context, categoryid int) (*model.Getsubcategorydata, error) {
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("userid==")
	print(id.ID)
	var subcat []*model.Subcat
	subcat = subscription.Getsubcatbyid(categoryid)
	return &model.Getsubcategorydata{Status: true,Code: http.StatusOK,Message: "Success",Subcategories: subcat},nil
}

func (r *queryResolver) Gettenantsubcategory(ctx context.Context, tenantid int, categoryid int, moduleid int) (*model.Gettenantsubcategorydata, error) {
	id, usererr :=datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("userid==")
	print(id.ID)
	var tenantsubcat []*model.Tenantsubcat
	tenantsubcat = subscription.Gettenantsubcat(moduleid,tenantid,categoryid)
	return &model.Gettenantsubcategorydata{Status: true,Code: http.StatusOK,Message: "Success",Tenantsubcategories: tenantsubcat},nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
