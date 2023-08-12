package config

import "fmt"

const (
	ProfilesCQRSResolverName      = "profiles"
	GeoCQRSResolverName           = "geo"
	ShopsCQRSResolverName         = "shops"
	CourierCQRSResolverName       = "courier"
	DeliveryCQRSResolverName      = "delivery"
	OffersCQRSResolverName        = "offers"
	BizCQRSResolverName           = "biz"
	NotificationsCQRSResolverName = "notifications"
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
