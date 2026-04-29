package roundtrip

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/apitesting/roundtrip"
	"github.com/hashicorp/terraform-provider-azuread/xpprovider"
	"k8s.io/apimachinery/pkg/runtime"

	namespacedapis "github.com/upbound/provider-azuread/v2/apis/namespaced"

	clusterapis "github.com/upbound/provider-azuread/v2/apis/cluster"
	"github.com/upbound/provider-azuread/v2/config"
)

func TestAPIRoundTrip(t *testing.T) {
	schema, err := xpprovider.GetProviderSchema(t.Context())
	if err != nil {
		t.Fatalf("GetProviderSchema: %s", err)
	}
	provider, err := config.GetProvider(t.Context(), schema, false)
	if err != nil {
		t.Fatalf("GetProvider: %s", err)
	}
	providerNamespaced, err := config.GetNamespacedProvider(t.Context(), schema, false)
	if err != nil {
		t.Fatalf("GetNamespacedProvider: %s", err)
	}

	testScheme := runtime.NewScheme()
	if err := clusterapis.AddToScheme(testScheme); err != nil {
		t.Fatalf("cluster-scoped apis AddToScheme: %s", err)
	}
	if err := namespacedapis.AddToScheme(testScheme); err != nil {
		t.Fatalf("namespaced apis AddToScheme: %s", err)
	}

	rt, err := roundtrip.NewRoundTripTest(provider, providerNamespaced, testScheme,
		roundtrip.WithComparisonOptions(roundtrip.EquateEmptyAndSingleZeroSlice()),
	)
	if err != nil {
		t.Fatalf("NewRoundTripTest: %s", err)
	}

	t.Run("TestSerializationRoundtrip", func(t *testing.T) {
		rt.TestSerializationRoundtrip(t)
	})

	t.Run("TestConversionRoundtrip", func(t *testing.T) {
		rt.TestConversionRoundtrip(t)
	})
}
