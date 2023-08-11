package config

import "fmt"

const (
	ProfilesCQRSResolverName      = "profilescqrs"
	GeoCQRSResolverName           = "geocqrs"
	ShopsCQRSResolverName         = "shopscqrs"
	CourierCQRSResolverName       = "couriercqrs"
	DeliveryCQRSResolverName      = "deliverycqrs"
	OffersCQRSResolverName        = "offerscqrs"
	BizCQRSResolverName           = "bizcqrs"
	NotificationsCQRSResolverName = "notificationscqrs"
)

type RequestResolver string

func (rr RequestResolver) String() string {
	fmt.Println("resolver string: ", string(rr))
	return string(rr)
}

const (
	ProfilesCQRSResolverKey      = RequestResolver(ProfilesCQRSResolverName)
	GeoCQRSResolverKey           = RequestResolver(GeoCQRSResolverName)
	ShopsCQRSResolverKey         = RequestResolver(ShopsCQRSResolverName)
	CourierCQRSResolverKey       = RequestResolver(CourierCQRSResolverName)
	DeliveryCQRSResolverKey      = RequestResolver(DeliveryCQRSResolverName)
	OffersCQRSResolverKey        = RequestResolver(OffersCQRSResolverName)
	BizCQRSResolverKey           = RequestResolver(BizCQRSResolverName)
	NotificationsCQRSResolverKey = RequestResolver(NotificationsCQRSResolverName)
)
