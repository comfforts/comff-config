package config

type RequestResolver string

func (rr RequestResolver) String() string {
	return string(rr)
}

const (
	ProfilesCQRSResolverKey      = RequestResolver("profiles-cqrs")
	GeoCQRSResolverKey           = RequestResolver("geo-cqrs")
	ShopsCQRSResolverKey         = RequestResolver("shops-cqrs")
	CourierCQRSResolverKey       = RequestResolver("courier-cqrs")
	DeliveryCQRSResolverKey      = RequestResolver("delivery-cqrs")
	OffersCQRSResolverKey        = RequestResolver("offers-cqrs")
	BizCQRSResolverKey           = RequestResolver("biz-cqrs")
	NotificationsCQRSResolverKey = RequestResolver("notifications-cqrs")
)
