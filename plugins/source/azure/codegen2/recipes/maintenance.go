// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"

func init() {
	tables := []Table{
		{
			Service:        "armmaintenance",
			Name:           "configurations",
			Struct:         &armmaintenance.Configuration{},
			ResponseStruct: &armmaintenance.ConfigurationsClientListResponse{},
			Client:         &armmaintenance.ConfigurationsClient{},
			ListFunc:       (&armmaintenance.ConfigurationsClient{}).NewListPager,
			NewFunc:        armmaintenance.NewConfigurationsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/maintenanceConfigurations",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Maintenance)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armmaintenance",
			Name:           "public_maintenance_configurations",
			Struct:         &armmaintenance.Configuration{},
			ResponseStruct: &armmaintenance.PublicMaintenanceConfigurationsClientListResponse{},
			Client:         &armmaintenance.PublicMaintenanceConfigurationsClient{},
			ListFunc:       (&armmaintenance.PublicMaintenanceConfigurationsClient{}).NewListPager,
			NewFunc:        armmaintenance.NewPublicMaintenanceConfigurationsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/publicMaintenanceConfigurations",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Maintenance)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}