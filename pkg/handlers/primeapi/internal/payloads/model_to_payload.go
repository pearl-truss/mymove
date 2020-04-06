package payloads

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/handlers"

	"github.com/transcom/mymove/pkg/etag"
	"github.com/transcom/mymove/pkg/gen/primemessages"
	"github.com/transcom/mymove/pkg/models"
)

// MoveTaskOrder payload
func MoveTaskOrder(moveTaskOrder *models.MoveTaskOrder) *primemessages.MoveTaskOrder {
	if moveTaskOrder == nil {
		return nil
	}
	paymentRequests := PaymentRequests(&moveTaskOrder.PaymentRequests)
	mtoServiceItems := MTOServiceItems(&moveTaskOrder.MTOServiceItems)
	mtoShipments := MTOShipments(&moveTaskOrder.MTOShipments)
	payload := &primemessages.MoveTaskOrder{
		ID:                 strfmt.UUID(moveTaskOrder.ID.String()),
		CreatedAt:          strfmt.Date(moveTaskOrder.CreatedAt),
		IsAvailableToPrime: &moveTaskOrder.IsAvailableToPrime,
		IsCanceled:         &moveTaskOrder.IsCanceled,
		MoveOrderID:        strfmt.UUID(moveTaskOrder.MoveOrderID.String()),
		MoveOrder:          MoveOrder(&moveTaskOrder.MoveOrder),
		ReferenceID:        moveTaskOrder.ReferenceID,
		PaymentRequests:    *paymentRequests,
		MtoShipments:       *mtoShipments,
		UpdatedAt:          strfmt.Date(moveTaskOrder.UpdatedAt),
		ETag:               etag.GenerateEtag(moveTaskOrder.UpdatedAt),
	}

	if moveTaskOrder.PPMEstimatedWeight != nil {
		payload.PpmEstimatedWeight = int64(*moveTaskOrder.PPMEstimatedWeight)
	}

	if moveTaskOrder.PPMType != nil {
		payload.PpmType = *moveTaskOrder.PPMType
	}

	// mto service item references a polymorphic type which auto-generates an interface and getters and setters
	payload.SetMtoServiceItems(*mtoServiceItems)

	return payload
}

// MoveTaskOrders payload
func MoveTaskOrders(moveTaskOrders *models.MoveTaskOrders) []*primemessages.MoveTaskOrder {
	payload := make(primemessages.MoveTaskOrders, len(*moveTaskOrders))

	for i, m := range *moveTaskOrders {
		payload[i] = MoveTaskOrder(&m)
	}
	return payload
}

// Customer payload
func Customer(customer *models.Customer) *primemessages.Customer {
	if customer == nil {
		return nil
	}
	payload := primemessages.Customer{
		FirstName:          swag.StringValue(customer.FirstName),
		LastName:           swag.StringValue(customer.LastName),
		DodID:              swag.StringValue(customer.DODID),
		ID:                 strfmt.UUID(customer.ID.String()),
		UserID:             strfmt.UUID(customer.UserID.String()),
		CurrentAddress:     Address(&customer.CurrentAddress),
		DestinationAddress: Address(&customer.DestinationAddress),
		ETag:               etag.GenerateEtag(customer.UpdatedAt),
		Branch:             swag.StringValue(customer.Agency),
	}

	if customer.PhoneNumber != nil {
		payload.Phone = *customer.PhoneNumber
	}

	if customer.Email != nil {
		payload.Email = *customer.Email
	}
	return &payload
}

// MoveOrder payload
func MoveOrder(moveOrder *models.MoveOrder) *primemessages.MoveOrder {
	if moveOrder == nil {
		return nil
	}
	destinationDutyStation := DutyStation(moveOrder.DestinationDutyStation)
	originDutyStation := DutyStation(moveOrder.OriginDutyStation)
	if moveOrder.Grade != nil && moveOrder.Entitlement != nil {
		moveOrder.Entitlement.SetWeightAllotment(*moveOrder.Grade)
	}
	reportByDate := strfmt.Date(*moveOrder.ReportByDate)
	entitlements := Entitlement(moveOrder.Entitlement)
	payload := primemessages.MoveOrder{
		CustomerID:             strfmt.UUID(moveOrder.CustomerID.String()),
		Customer:               Customer(moveOrder.Customer),
		DestinationDutyStation: destinationDutyStation,
		Entitlement:            entitlements,
		ID:                     strfmt.UUID(moveOrder.ID.String()),
		OriginDutyStation:      originDutyStation,
		OrderNumber:            moveOrder.OrderNumber,
		LinesOfAccounting:      moveOrder.LinesOfAccounting,
		Rank:                   moveOrder.Grade,
		ConfirmationNumber:     moveOrder.ConfirmationNumber,
		ReportByDate:           reportByDate,
		ETag:                   etag.GenerateEtag(moveOrder.UpdatedAt),
	}
	return &payload
}

// Entitlement payload
func Entitlement(entitlement *models.Entitlement) *primemessages.Entitlements {
	if entitlement == nil {
		return nil
	}
	var proGearWeight, proGearWeightSpouse, totalWeight int64
	if entitlement.WeightAllotment() != nil {
		proGearWeight = int64(entitlement.WeightAllotment().ProGearWeight)
		proGearWeightSpouse = int64(entitlement.WeightAllotment().ProGearWeightSpouse)
		totalWeight = int64(entitlement.WeightAllotment().TotalWeightSelf)
	}
	var authorizedWeight *int64
	if entitlement.AuthorizedWeight() != nil {
		aw := int64(*entitlement.AuthorizedWeight())
		authorizedWeight = &aw
	}
	var sit int64
	if entitlement.StorageInTransit != nil {
		sit = int64(*entitlement.StorageInTransit)
	}
	var totalDependents int64
	if entitlement.TotalDependents != nil {
		totalDependents = int64(*entitlement.TotalDependents)
	}
	return &primemessages.Entitlements{
		ID:                    strfmt.UUID(entitlement.ID.String()),
		AuthorizedWeight:      authorizedWeight,
		DependentsAuthorized:  entitlement.DependentsAuthorized,
		NonTemporaryStorage:   entitlement.NonTemporaryStorage,
		PrivatelyOwnedVehicle: entitlement.PrivatelyOwnedVehicle,
		ProGearWeight:         proGearWeight,
		ProGearWeightSpouse:   proGearWeightSpouse,
		StorageInTransit:      sit,
		TotalDependents:       totalDependents,
		TotalWeight:           totalWeight,
		ETag:                  etag.GenerateEtag(entitlement.UpdatedAt),
	}
}

// DutyStation payload
func DutyStation(dutyStation *models.DutyStation) *primemessages.DutyStation {
	if dutyStation == nil {
		return nil
	}
	address := Address(&dutyStation.Address)
	payload := primemessages.DutyStation{
		Address:   address,
		AddressID: address.ID,
		ID:        strfmt.UUID(dutyStation.ID.String()),
		Name:      dutyStation.Name,
	}
	return &payload
}

// Address payload
func Address(address *models.Address) *primemessages.Address {
	if address == nil {
		return nil
	}
	return &primemessages.Address{
		ID:             strfmt.UUID(address.ID.String()),
		StreetAddress1: &address.StreetAddress1,
		StreetAddress2: address.StreetAddress2,
		StreetAddress3: address.StreetAddress3,
		City:           &address.City,
		State:          &address.State,
		PostalCode:     &address.PostalCode,
		Country:        address.Country,
		ETag:           etag.GenerateEtag(address.UpdatedAt),
	}
}

// MTOAgent payload
func MTOAgent(mtoAgent *models.MTOAgent) *primemessages.MTOAgent {
	if mtoAgent == nil {
		return nil
	}

	return &primemessages.MTOAgent{
		AgentType:     primemessages.MTOAgentType(mtoAgent.MTOAgentType),
		Email:         mtoAgent.Email,
		FirstName:     mtoAgent.FirstName,
		ID:            strfmt.UUID(mtoAgent.ID.String()),
		LastName:      mtoAgent.LastName,
		MtoShipmentID: strfmt.UUID(mtoAgent.MTOShipmentID.String()),
		Phone:         mtoAgent.Phone,
	}
}

// MTOAgents payload
func MTOAgents(mtoAgents *models.MTOAgents) *primemessages.MTOAgents {
	if mtoAgents == nil {
		return nil
	}

	agents := make(primemessages.MTOAgents, len(*mtoAgents))

	for i, m := range *mtoAgents {
		agents[i] = MTOAgent(&m)
	}

	return &agents
}

// PaymentRequest payload
func PaymentRequest(paymentRequest *models.PaymentRequest) *primemessages.PaymentRequest {
	return &primemessages.PaymentRequest{
		ID:                   strfmt.UUID(paymentRequest.ID.String()),
		IsFinal:              &paymentRequest.IsFinal,
		MoveTaskOrderID:      strfmt.UUID(paymentRequest.MoveTaskOrderID.String()),
		PaymentRequestNumber: paymentRequest.PaymentRequestNumber,
		RejectionReason:      paymentRequest.RejectionReason,
		Status:               primemessages.PaymentRequestStatus(paymentRequest.Status),
		ETag:                 etag.GenerateEtag(paymentRequest.UpdatedAt),
	}
}

// PaymentRequests payload
func PaymentRequests(paymentRequests *models.PaymentRequests) *primemessages.PaymentRequests {
	payload := make(primemessages.PaymentRequests, len(*paymentRequests))

	for i, p := range *paymentRequests {
		payload[i] = PaymentRequest(&p)
	}
	return &payload
}

// MTOShipment payload
func MTOShipment(mtoShipment *models.MTOShipment) *primemessages.MTOShipment {
	payload := &primemessages.MTOShipment{
		ID:                       strfmt.UUID(mtoShipment.ID.String()),
		Agents:                   *MTOAgents(&mtoShipment.MTOAgents),
		MoveTaskOrderID:          strfmt.UUID(mtoShipment.MoveTaskOrderID.String()),
		ShipmentType:             primemessages.MTOShipmentType(mtoShipment.ShipmentType),
		CustomerRemarks:          mtoShipment.CustomerRemarks,
		PickupAddress:            Address(mtoShipment.PickupAddress),
		Status:                   string(mtoShipment.Status),
		DestinationAddress:       Address(mtoShipment.DestinationAddress),
		SecondaryPickupAddress:   Address(mtoShipment.SecondaryPickupAddress),
		SecondaryDeliveryAddress: Address(mtoShipment.SecondaryDeliveryAddress),
		CreatedAt:                strfmt.DateTime(mtoShipment.CreatedAt),
		UpdatedAt:                strfmt.DateTime(mtoShipment.UpdatedAt),
		ETag:                     etag.GenerateEtag(mtoShipment.UpdatedAt),
	}

	if mtoShipment.ApprovedDate != nil {
		payload.ApprovedDate = strfmt.Date(*mtoShipment.ApprovedDate)
	}

	if mtoShipment.ScheduledPickupDate != nil {
		payload.ScheduledPickupDate = strfmt.Date(*mtoShipment.ScheduledPickupDate)
	}

	if mtoShipment.RequestedPickupDate != nil && !mtoShipment.RequestedPickupDate.IsZero() {
		payload.RequestedPickupDate = strfmt.Date(*mtoShipment.RequestedPickupDate)
	}

	if mtoShipment.ActualPickupDate != nil && !mtoShipment.ActualPickupDate.IsZero() {
		payload.ActualPickupDate = strfmt.Date(*mtoShipment.ActualPickupDate)
	}

	if mtoShipment.FirstAvailableDeliveryDate != nil && !mtoShipment.FirstAvailableDeliveryDate.IsZero() {
		payload.FirstAvailableDeliveryDate = strfmt.Date(*mtoShipment.FirstAvailableDeliveryDate)
	}

	if mtoShipment.PrimeEstimatedWeight != nil && mtoShipment.PrimeEstimatedWeightRecordedDate != nil {
		payload.PrimeEstimatedWeight = int64(*mtoShipment.PrimeEstimatedWeight)
		payload.PrimeEstimatedWeightRecordedDate = strfmt.Date(*mtoShipment.PrimeEstimatedWeightRecordedDate)
	}
	return payload
}

// MTOShipments payload
func MTOShipments(mtoShipments *models.MTOShipments) *primemessages.MTOShipments {
	payload := make(primemessages.MTOShipments, len(*mtoShipments))

	for i, m := range *mtoShipments {
		payload[i] = MTOShipment(&m)
	}
	return &payload
}

// MTOServiceItem payload
func MTOServiceItem(mtoServiceItem *models.MTOServiceItem) primemessages.MTOServiceItem {
	var payload primemessages.MTOServiceItem

	// here we determine which payload model to use based on the re service code
	switch mtoServiceItem.ReService.Code {
	case models.ReServiceCodeDOFSIT:
		payload = &primemessages.MTOServiceItemDOFSIT{
			ReServiceCode:    primemessages.ReServiceCode(mtoServiceItem.ReService.Code),
			PickupPostalCode: mtoServiceItem.PickupPostalCode,
			Reason:           mtoServiceItem.Reason,
		}
	case models.ReServiceCodeDCRT, models.ReServiceCodeDUCRT:
		item := mtoServiceItem.GetItemDimension()
		if item == nil {
			item = &models.MTOServiceItemDimension{}
		}
		crate := mtoServiceItem.GetCrateDimension()
		if crate == nil {
			crate = &models.MTOServiceItemDimension{}
		}
		payload = &primemessages.MTOServiceItemDomesticCrating{
			ReServiceCode: primemessages.ReServiceCode(mtoServiceItem.ReService.Code),
			Item: &primemessages.MTOServiceItemDimension{
				ID:     strfmt.UUID(item.ID.String()),
				Type:   primemessages.DimensionType(item.Type),
				Height: item.Height.Int32Ptr(),
				Length: item.Length.Int32Ptr(),
				Width:  item.Width.Int32Ptr(),
			},
			Crate: &primemessages.MTOServiceItemDimension{
				ID:     strfmt.UUID(crate.ID.String()),
				Type:   primemessages.DimensionType(crate.Type),
				Height: crate.Height.Int32Ptr(),
				Length: crate.Length.Int32Ptr(),
				Width:  crate.Width.Int32Ptr(),
			},
			Description: mtoServiceItem.Description,
		}
	case models.ReServiceCodeDDSHUT, models.ReServiceCodeDOSHUT:
		payload = &primemessages.MTOServiceItemShuttle{
			Description:   mtoServiceItem.Description,
			ReServiceCode: primemessages.ReServiceCode(mtoServiceItem.ReService.Code),
			Reason:        mtoServiceItem.Reason,
		}
	default:
		// otherwise, basic service item
		payload = &primemessages.MTOServiceItemBasic{
			ReServiceCode: primemessages.ReServiceCode(mtoServiceItem.ReService.Code),
		}
	}

	// set all relevant fields that apply to all service items
	payload.SetID(strfmt.UUID(mtoServiceItem.ID.String()))
	payload.SetMoveTaskOrderID(strfmt.UUID(mtoServiceItem.MoveTaskOrderID.String()))
	payload.SetReServiceID(strfmt.UUID(mtoServiceItem.ReServiceID.String()))
	payload.SetReServiceName(mtoServiceItem.ReService.Name)
	payload.SetETag(etag.GenerateEtag(mtoServiceItem.UpdatedAt))

	return payload
}

// MTOServiceItems payload
func MTOServiceItems(mtoServiceItems *models.MTOServiceItems) *[]primemessages.MTOServiceItem {
	var payload []primemessages.MTOServiceItem

	for _, p := range *mtoServiceItems {
		payload = append(payload, MTOServiceItem(&p))
	}
	return &payload
}

// ValidationError describes validation errors from the model or properties
func ValidationError(title string, detail string, instance uuid.UUID, validationErrors *validate.Errors) *primemessages.ValidationError {
	return &primemessages.ValidationError{
		InvalidFields: handlers.NewValidationErrorsResponse(validationErrors).Errors,
		ClientError:   *clientError(title, detail, instance),
	}
}

func clientError(title string, detail string, instance uuid.UUID) *primemessages.ClientError {
	return &primemessages.ClientError{
		Title:    handlers.FmtString(title),
		Detail:   handlers.FmtString(detail),
		Instance: handlers.FmtUUID(instance),
	}
}
