package webhooks

import (
	"testing"

	"github.com/processout/connect-sdk-go/defaultimpl"
)

func TestCreateHelper(t *testing.T) {
	store, err := NewInMemorySecretKeyStore()
	if err != nil {
		t.Fatal(err)
	}

	helper, err := CreateHelper(store)
	if err != nil {
		t.Fatal(err)
	}

	marshaller, err := defaultimpl.NewDefaultMarshaller()
	if err != nil {
		t.Fatal(err)
	}

	if helper.marshaller != marshaller {
		t.Fatalf("marshaller mismatch %v %v", helper.marshaller, marshaller)
	}
	if helper.secretKeyStore != store {
		t.Fatalf("secretKeyStore mismatch %v %v", helper.secretKeyStore, store)
	}
}
