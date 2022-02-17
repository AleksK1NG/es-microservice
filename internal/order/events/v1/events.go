package v1

import (
	"time"

	"github.com/AleksK1NG/es-microservice/internal/order/models"
	"github.com/AleksK1NG/es-microservice/pkg/es"
)

const (
	OrderCreated                = "V1_ORDER_CREATED"
	OrderPaid                   = "V1_ORDER_PAID"
	OrderSubmitted              = "V1_ORDER_SUBMITTED"
	OrderDelivered              = "V1_ORDER_DELIVERED"
	OrderCanceled               = "V1_ORDER_CANCELED"
	OrderUpdated                = "V1_ORDER_UPDATED"
	OrderDeliveryAddressUpdated = "V1_ORDER_DELIVERY_ADDRESS_UPDATED"
)

type OrderCreatedEvent struct {
	ShopItems       []*models.ShopItem `json:"shopItems" bson:"shopItems,omitempty"`
	AccountEmail    string             `json:"accountEmail" bson:"accountEmail,omitempty"`
	DeliveryAddress string             `json:"deliveryAddress" bson:"deliveryAddress,omitempty"`
}

func NewOrderCreatedEvent(aggregate es.Aggregate, shopItems []*models.ShopItem, accountEmail, deliveryAddress string) (es.Event, error) {
	eventData := OrderCreatedEvent{
		ShopItems:       shopItems,
		AccountEmail:    accountEmail,
		DeliveryAddress: deliveryAddress,
	}
	event := es.NewBaseEvent(aggregate, OrderCreated)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}

func NewOrderPaidEvent(aggregate es.Aggregate, payment *models.Payment) (es.Event, error) {
	event := es.NewBaseEvent(aggregate, OrderPaid)
	if err := event.SetJsonData(&payment); err != nil {
		return es.Event{}, err
	}
	return event, nil
}

func NewSubmitOrderEvent(aggregate es.Aggregate) (es.Event, error) {
	return es.NewBaseEvent(aggregate, OrderSubmitted), nil
}

type OrderUpdatedEvent struct {
	ShopItems []*models.ShopItem `json:"shopItems" bson:"shopItems,omitempty"`
}

func NewOrderUpdatedEvent(aggregate es.Aggregate, shopItems []*models.ShopItem) (es.Event, error) {
	eventData := OrderUpdatedEvent{ShopItems: shopItems}
	event := es.NewBaseEvent(aggregate, OrderUpdated)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}

type OrderDeliveryAddressChangedEvent struct {
	DeliveryAddress string `json:"deliveryAddress" bson:"deliveryAddress,omitempty"`
}

func NewOrderDeliveryAddressChangedEvent(aggregate es.Aggregate, deliveryAddress string) (es.Event, error) {
	eventData := OrderDeliveryAddressChangedEvent{DeliveryAddress: deliveryAddress}
	event := es.NewBaseEvent(aggregate, OrderDeliveryAddressUpdated)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}

type OrderCanceledEvent struct {
	CancelReason string `json:"cancelReason"`
}

func NewOrderCanceledEvent(aggregate es.Aggregate, cancelReason string) (es.Event, error) {
	eventData := OrderCanceledEvent{CancelReason: cancelReason}
	event := es.NewBaseEvent(aggregate, OrderCanceled)
	err := event.SetJsonData(&eventData)
	if err != nil {
		return es.Event{}, err
	}
	return event, nil
}

type OrderDeliveredEvent struct {
	DeliveryTimestamp time.Time `json:"deliveryTimestamp"`
}

func NewOrderDeliveredEvent(aggregate es.Aggregate, deliveryTimestamp time.Time) (es.Event, error) {
	eventData := OrderDeliveredEvent{DeliveryTimestamp: deliveryTimestamp}
	event := es.NewBaseEvent(aggregate, OrderDelivered)
	err := event.SetJsonData(&eventData)
	if err != nil {
		return es.Event{}, err
	}
	return event, nil
}