package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/engajerest/auth/Models/users"
	"github.com/engajerest/auth/controller"
	"github.com/engajerest/sparkle/Models/subscription"
	"github.com/engajerest/sparkle/graph/generated"
	"github.com/engajerest/sparkle/graph/model"
)

func (r *mutationResolver) Subscribe(ctx context.Context, input model.Data) (*model.SubscribedData, error) {
	id, usererr := controller.ForContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("rajjj")
	print(id.ID)
	var user users.User
	user.ID = id.ID
	print("check")
	print(user.ID)
	var data subscription.SubscriptionData
	data.Info.CategoryId = input.Tenantinfo.CategoryID
	data.Info.Regno = input.Tenantinfo.Regno
	data.Info.Name = input.Tenantinfo.Name
	data.Info.Email = input.Tenantinfo.Email
	data.Info.Mobile = input.Tenantinfo.Mobile
	data.Info.SubCategoryID = input.Tenantinfo.SubCategoryID
	data.Address.Address = input.Tenantlocation.Address
	data.Address.State = input.Tenantlocation.State
	data.Address.Suburb = input.Tenantlocation.Suburb
	data.Address.Zip = input.Tenantlocation.Zip
	data.Address.Countrycode = input.Tenantlocation.Countrycode
	data.Address.Latitude = input.Tenantlocation.Latitude
	data.Address.Longitude = input.Tenantlocation.Longitude
	data.Address.TimeZone = input.Tenantlocation.TimeZone
	data.Address.OpenTime = input.Tenantlocation.Opentime
	data.Address.CloseTime = input.Tenantlocation.Closetime
	data.Address.CurrencyCode = input.Subscriptiondetails.CurrencyCode
	var data1 subscription.TenantSubscription
	data1.CurrencyId = input.Subscriptiondetails.CurrencyID

	data1.Date = input.Subscriptiondetails.TransactionDate
	data1.CurrencyId = input.Subscriptiondetails.CurrencyID
	data1.ModuleId = input.Subscriptiondetails.ModuleID
	data1.PackageId = input.Subscriptiondetails.PackageID
	data1.PaymentId = *input.Subscriptiondetails.PaymentID
	data1.PaymentStatus = input.Subscriptiondetails.PaymentStatus
	data1.Price = input.Subscriptiondetails.Price
	data1.Quantity = input.Subscriptiondetails.Quantity
	data1.TaxId = input.Subscriptiondetails.TaxID
	data1.TaxAmount = input.Subscriptiondetails.TaxAmount
	data1.TotalAmount = input.Subscriptiondetails.TotalAmount
	var data2 subscription.SubscribedData
	var auth subscription.AuthUser
	print("check456")
	tenantId, err := data.CreateTenant(user.ID)
	if err != nil {
		return nil, err
	}
	print(tenantId)
	if tenantId != 0 {
		tenantlocationid := data.InsertTenantLocation(tenantId, id.ID)
		print("loc-id")
		print(tenantlocationid)
		auth.TenantID = int(tenantId)
		auth.LocationId = int(tenantlocationid)
		subscribedid := data1.InsertSubscription(tenantId)
		print("subs-id")
		print(subscribedid)
		print(tenantId)

	}
	subscribed, Error := data2.GetSubscribedData(tenantId)
	if Error != nil {
		fmt.Println("rows were not found")
		return nil, Error
	}

	status := auth.UpdateAuthUser(id.ID)
	print(status)
	if status != true {
		return nil, errors.New("tenant not subscribed")
	}
	return &model.SubscribedData{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		Info: &model.TenantData{
			TenantID:   subscribed.TenantID,
			TenantName: subscribed.TenantName,
			ModuleID:   subscribed.ModuleID,
			ModuleName: subscribed.ModuleName,
		},
	}, nil
}

func (r *mutationResolver) Createtenantuser(ctx context.Context, create *model.Tenantuser) (*model.Tenantuserdata, error) {
	id, usererr := controller.ForContext(ctx)
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
	id, usererr := controller.ForContext(ctx)
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
	id, usererr := controller.ForContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("userid=")
	print(id.ID)

	var data subscription.BusinessUpdate
	data.TenantID = businessinfo.Businessupdate.Tenantid
	data.TenantaccId = *businessinfo.Businessupdate.Tenantaccid
	data.Brandname = *businessinfo.Businessupdate.Brandname
	data.About = *businessinfo.Businessupdate.About
	data.Paymode1 = *businessinfo.Businessupdate.Cod
	data.Paymode2 = *businessinfo.Businessupdate.Digital
	var check []subscription.Social
	schemasocial := *&businessinfo.Socialupdate
	for _, v := range schemasocial {
		check = append(check, subscription.Social{SociaProfile: *v.Socialprofile, SocialLink: *v.Sociallink, SocialIcon: *v.Socialicon})
	}
	data1 := data.UpdateTenantBusiness()
	social := data.InsertTenantSocial(check, data.TenantID)
	print(social)
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
	id, usererr := controller.ForContext(ctx)
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

func (r *queryResolver) Sparkle(ctx context.Context) (*model.Sparkle, error) {
	id, usererr := controller.ForContext(ctx)
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
		cat = append(cat, &model.Category{CategoryID: category.CategoryID, Name: category.Name, Type: category.Typeid, SortOrder: category.SortOrder, Status: category.Status})

	}
	subcategoryGetAll = subscription.GetAllSubCategory()
	for _, subcategory := range subcategoryGetAll {
		sub = append(sub, &model.SubCategory{CategoryID: subcategory.CategoryID, SubCategoryID: subcategory.SubCategoryID, Name: subcategory.Name, Type: subcategory.Typeid, SortOrder: subcategory.SortOrder, Status: subcategory.Status, Icon: subcategory.Icon})
	}
	packageGetAll = subscription.GetAllPackages()
	for _, packdata := range packageGetAll {
		pack = append(pack, &model.Package{ModuleID: packdata.ModuleID, Modulename: packdata.ModuleName, Name: packdata.Name, PackageID: packdata.PackageID, Status: packdata.Status, PackageAmount: packdata.PackageAmount, PaymentMode: packdata.PaymentMode, PackageContent: packdata.PackageContent, PackageIcon: packdata.PackageIcon})

	}

	return &model.Sparkle{
		Category:    cat,
		Subcategory: sub,
		Package:     pack,
	}, nil
}

func (r *queryResolver) Location(ctx context.Context, tenantid int) (*model.Getalllocations, error) {
	id, usererr := controller.ForContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("getalloc")
	print(id.ID)

	var Result []*model.Locationgetall
	var userresult []*model.Usertenant
	var locationGetAll []subscription.Tenantlocation

	locationGetAll = subscription.LocationTest(tenantid)

	for _, loco := range locationGetAll {
		userresult = make([]*model.Usertenant, len(loco.Appuserprofiles))
		for i, key := range loco.Appuserprofiles {
			userresult[i] = &model.Usertenant{
				Userid: key.Userid,
				 Userlocationid: key.Userlocationid,
				 Firstname: key.Firstname,
				 Lastname: key.Lastname,
				 Mobile: key.Contactno,
				 Email: key.Email,
				}
		}
		Result = append(Result, &model.Locationgetall{
			Locationid:   loco.Locationid,
			LocationName: loco.Locationname,
			Tenantid:     loco.Tenantid,
			Email:        &loco.Email,
			Contact:      &loco.Contactno,
			Address:      loco.Address,
			Suburb:       loco.City,
			State:        loco.State,
			Countycode:   loco.Countrycode,
			Postcode:     loco.Postcode,
			Latitude:     loco.Latitude,
			Longitude:    loco.Longitude,
			Openingtime:  loco.Opentime,
			Closingtime:  loco.Closetime,
			Status:       loco.Status,
			Createdby:    loco.Createdby,
			Tenantusers:  userresult,
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
	id, usererr := controller.ForContext(ctx)
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

func (r *queryResolver) GetBusiness(ctx context.Context, tenantid int) (*model.GetBusinessdata, error) {
	id, usererr := controller.ForContext(ctx)
	if usererr != nil {
		return nil, errors.New("user not detected")
	}
	print("getbusin")
	print(id.ID)
	var Result []*model.Socialinfo
	var business subscription.BusinessUpdate
	businessinfo, value := business.GetBusinessInfo(tenantid)
	if value != true {
		print(value)
		return &model.GetBusinessdata{
			Status:       true,
			Code:         http.StatusOK,
			Message:      "Success",
			Businessinfo: nil,
		}, nil
	}

	var socialgetall []subscription.Social
	socialgetall = subscription.GetAllSocial(tenantid)
	for _, user := range socialgetall {
		Result = append(Result, &model.Socialinfo{
			Socialprofile: &user.SociaProfile,
			Sociallink:    &user.SocialLink,
			Socialicon:    &user.SocialIcon,
		})
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
			Social:      Result,
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
