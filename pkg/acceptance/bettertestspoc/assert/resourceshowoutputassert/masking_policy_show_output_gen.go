// Code generated by assertions generator; DO NOT EDIT.

package resourceshowoutputassert

import (
	"testing"
	"time"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
)

// to ensure sdk package is used
var _ = sdk.Object{}

type MaskingPolicyShowOutputAssert struct {
	*assert.ResourceAssert
}

func MaskingPolicyShowOutput(t *testing.T, name string) *MaskingPolicyShowOutputAssert {
	t.Helper()

	maskingPolicyAssert := MaskingPolicyShowOutputAssert{
		ResourceAssert: assert.NewResourceAssert(name, "show_output"),
	}
	maskingPolicyAssert.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &maskingPolicyAssert
}

func ImportedMaskingPolicyShowOutput(t *testing.T, id string) *MaskingPolicyShowOutputAssert {
	t.Helper()

	maskingPolicyAssert := MaskingPolicyShowOutputAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "show_output"),
	}
	maskingPolicyAssert.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &maskingPolicyAssert
}

////////////////////////////
// Attribute value checks //
////////////////////////////

func (m *MaskingPolicyShowOutputAssert) HasCreatedOn(expected time.Time) *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueSet("created_on", expected.String()))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasName(expected string) *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueSet("name", expected))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasDatabaseName(expected string) *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueSet("database_name", expected))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasSchemaName(expected string) *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueSet("schema_name", expected))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasKind(expected string) *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueSet("kind", expected))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasOwner(expected string) *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueSet("owner", expected))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasComment(expected string) *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueSet("comment", expected))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasExemptOtherPolicies(expected bool) *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputBoolValueSet("exempt_other_policies", expected))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasOwnerRoleType(expected string) *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueSet("owner_role_type", expected))
	return m
}

///////////////////////////////
// Attribute no value checks //
///////////////////////////////

func (m *MaskingPolicyShowOutputAssert) HasNoCreatedOn() *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueNotSet("created_on"))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasNoName() *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueNotSet("name"))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasNoDatabaseName() *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueNotSet("database_name"))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasNoSchemaName() *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueNotSet("schema_name"))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasNoKind() *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueNotSet("kind"))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasNoOwner() *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueNotSet("owner"))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasNoComment() *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueNotSet("comment"))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasNoExemptOtherPolicies() *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputBoolValueNotSet("exempt_other_policies"))
	return m
}

func (m *MaskingPolicyShowOutputAssert) HasNoOwnerRoleType() *MaskingPolicyShowOutputAssert {
	m.AddAssertion(assert.ResourceShowOutputValueNotSet("owner_role_type"))
	return m
}
