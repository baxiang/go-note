package profilesvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/url"
	"strings"
)
type Endpoints struct {
	PostProfileEndpoint endpoint.Endpoint
	GetProfileEndpoint    endpoint.Endpoint
	PutProfileEndpoint    endpoint.Endpoint
	PatchProfileEndpoint  endpoint.Endpoint
	DeleteProfileEndpoint endpoint.Endpoint
	GetAddressesEndpoint  endpoint.Endpoint
	GetAddressEndpoint    endpoint.Endpoint
	PostAddressEndpoint   endpoint.Endpoint
	DeleteAddressEndpoint endpoint.Endpoint

}

func MakeServerEndpoints(s Service)Endpoints{
	return Endpoints{
		PostProfileEndpoint:   MakePostProfileEndpoint(s),
		GetProfileEndpoint:    MakeGetProfileEndpoint(s),
		PutProfileEndpoint:    MakePutProfileEndpoint(s),
		PatchProfileEndpoint:  MakePatchProfileEndpoint(s),
		DeleteProfileEndpoint: MakeDeleteProfileEndpoint(s),
		GetAddressesEndpoint:  MakeGetAddressesEndpoint(s),
		GetAddressEndpoint:    MakeGetAddressEndpoint(s),
		PostAddressEndpoint:   MakePostAddressEndpoint(s),
		DeleteAddressEndpoint: MakeDeleteAddressEndpoint(s),
	}
}

func MakeClientEndpoints(instance string)(Endpoints,error){
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	tgt, err := url.Parse(instance)
	if err != nil {
		return Endpoints{}, err
	}
	tgt.Path = ""
	options := []httptransport.ClientOption{}
	return Endpoints{
		PostProfileEndpoint:   httptransport.NewClient("POST", tgt, encodePostProfileRequest, decodePostProfileResponse, options...).Endpoint(),
		GetProfileEndpoint:    httptransport.NewClient("GET", tgt, encodeGetProfileRequest, decodeGetProfileResponse, options...).Endpoint(),
		PutProfileEndpoint:    httptransport.NewClient("PUT", tgt, encodePutProfileRequest, decodePutProfileResponse, options...).Endpoint(),
		PatchProfileEndpoint:  httptransport.NewClient("PATCH", tgt, encodePatchProfileRequest, decodePatchProfileResponse, options...).Endpoint(),
		DeleteProfileEndpoint: httptransport.NewClient("DELETE", tgt, encodeDeleteProfileRequest, decodeDeleteProfileResponse, options...).Endpoint(),
		GetAddressesEndpoint:  httptransport.NewClient("GET", tgt, encodeGetAddressesRequest, decodeGetAddressesResponse, options...).Endpoint(),
		GetAddressEndpoint:    httptransport.NewClient("GET", tgt, encodeGetAddressRequest, decodeGetAddressResponse, options...).Endpoint(),
		PostAddressEndpoint:   httptransport.NewClient("POST", tgt, encodePostAddressRequest, decodePostAddressResponse, options...).Endpoint(),
		DeleteAddressEndpoint: httptransport.NewClient("DELETE", tgt, encodeDeleteAddressRequest, decodeDeleteAddressResponse, options...).Endpoint(),
	}, nil
}

type postProfileRequest struct {
	Profile Profile
}

type postProfileResponse struct {
	Err error `json:"err,omitempty"`
}
func MakePostProfileEndpoint(s Service)endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		  req :=request.(postProfileRequest)
		  err = s.PostProfile(ctx, req.Profile)
		  return postProfileResponse{Err:err},nil
	}
}

type getProfileRequest struct {
	ID string
}
type getProfileResponse struct {
	Profile Profile `json:"profile,omitempty"`
	Err     error   `json:"err,omitempty"`
}

func MakeGetProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getProfileRequest)
		p, e := s.GetProfile(ctx, req.ID)
		return getProfileResponse{Profile: p, Err: e}, nil
	}
}
func (r getProfileResponse) error() error { return r.Err }

type putProfileRequest struct {
	ID      string
	Profile Profile
}

type putProfileResponse struct {
	Err error `json:"err,omitempty"`
}
func (r putProfileResponse) error() error { return nil }

func MakePutProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(putProfileRequest)
		e := s.PutProfile(ctx, req.ID, req.Profile)
		return putProfileResponse{Err: e}, nil
	}
}


type patchProfileRequest struct {
	ID      string
	Profile Profile
}

type patchProfileResponse struct {
	Err error `json:"err,omitempty"`
}

func (r patchProfileResponse) error() error { return r.Err }


func MakePatchProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(patchProfileRequest)
		e := s.PatchProfile(ctx, req.ID, req.Profile)
		return patchProfileResponse{Err: e}, nil
	}
}
type deleteProfileRequest struct {
	ID string
}

type deleteProfileResponse struct {
	Err error `json:"err,omitempty"`
}

func (r deleteProfileResponse) error() error { return r.Err }

func MakeDeleteProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteProfileRequest)
		e := s.DeleteProfile(ctx, req.ID)
		return deleteProfileResponse{Err: e}, nil
	}
}

type getAddressesRequest struct {
	ProfileID string
}

type getAddressesResponse struct {
	Addresses []Address `json:"addresses,omitempty"`
	Err       error     `json:"err,omitempty"`
}

func (r getAddressesResponse) error() error { return r.Err }
func MakeGetAddressesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getAddressesRequest)
		a, e := s.GetAddresses(ctx, req.ProfileID)
		return getAddressesResponse{Addresses: a, Err: e}, nil
	}
}

type getAddressRequest struct {
	ProfileID string
	AddressID string
}

type getAddressResponse struct {
	Address Address `json:"address,omitempty"`
	Err     error   `json:"err,omitempty"`
}

func (r getAddressResponse) error() error { return r.Err }

func MakeGetAddressEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getAddressRequest)
		a, e := s.GetAddress(ctx, req.ProfileID, req.AddressID)
		return getAddressResponse{Address: a, Err: e}, nil
	}
}

type postAddressRequest struct {
	ProfileID string
	Address   Address
}

type postAddressResponse struct {
	Err error `json:"err,omitempty"`
}

func (r postAddressResponse) error() error { return r.Err }

func MakePostAddressEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postAddressRequest)
		e := s.PostAddress(ctx, req.ProfileID, req.Address)
		return postAddressResponse{Err: e}, nil
	}
}
type deleteAddressRequest struct {
	ProfileID string
	AddressID string
}

type deleteAddressResponse struct {
	Err error `json:"err,omitempty"`
}

func (r deleteAddressResponse) error() error { return r.Err }

func MakeDeleteAddressEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteAddressRequest)
		e := s.DeleteAddress(ctx, req.ProfileID, req.AddressID)
		return deleteAddressResponse{Err: e}, nil
	}
}
