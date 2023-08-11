package config

import "fmt"

type RequestResolver string

func (rr RequestResolver) String() string {
	fmt.Println("resolver string: ", string(rr))
	return string(rr)
}

const (
	ProfilesCQRSResolverKey      = RequestResolver("profilescqrs")
	GeoCQRSResolverKey           = RequestResolver("geocqrs")
	ShopsCQRSResolverKey         = RequestResolver("shopscqrs")
	CourierCQRSResolverKey       = RequestResolver("couriercqrs")
	DeliveryCQRSResolverKey      = RequestResolver("deliverycqrs")
	OffersCQRSResolverKey        = RequestResolver("offerscqrs")
	BizCQRSResolverKey           = RequestResolver("bizcqrs")
	NotificationsCQRSResolverKey = RequestResolver("notificationscqrs")
)
