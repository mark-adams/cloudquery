// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"

func Armpeering() []*Table {
	tables := []*Table{
		{
			NewFunc:        armpeering.NewServiceCountriesClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringServiceCountries",
			Namespace:      "Microsoft.Peering",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Peering)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ServiceCountriesClientListResponse",
		},
		{
			NewFunc:        armpeering.NewServiceLocationsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringServiceLocations",
			Namespace:      "Microsoft.Peering",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Peering)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ServiceLocationsClientListResponse",
		},
		{
			NewFunc:        armpeering.NewServiceProvidersClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringServiceProviders",
			Namespace:      "Microsoft.Peering",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Peering)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ServiceProvidersClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armpeering())
}