package pingfederate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func resourcePingFederateOauthAccessTokenMappingsResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingFederateOauthAccessTokenMappingsResourceCreate,
		Read:   resourcePingFederateOauthAccessTokenMappingsResourceRead,
		Update: resourcePingFederateOauthAccessTokenMappingsResourceUpdate,
		Delete: resourcePingFederateOauthAccessTokenMappingsResourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"context": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
							//TODO ValidateFunc: 'DEFAULT' or 'PCV' or 'IDP_CONNECTION' or 'IDP_ADAPTER' or 'AUTHENTICATION_POLICY_CONTRACT' or 'CLIENT_CREDENTIALS' or 'TOKEN_EXCHANGE_PROCESSOR_POLICY']: The Access Token Mapping Context type.
						},
						"context_ref": resourceLinkSchema(),
					},
				},
			},
			"access_token_manager_ref": resourceRequiredLinkSchema(),
			"ldap_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceCustomAttributeSource(),
			},
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     resourceAttributeFulfillmentValue(),
			},
			"issuance_criteria": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     resourceIssuanceCriteria(),
			},
		},
	}
}

func resourcePingFederateOauthAccessTokenMappingsResourceCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).OauthAccessTokenMappings
	input := pf.CreateMappingInput{
		Body: *resourcePingFederateOauthAccessTokenMappingsResourceReadData(d),
	}
	result, _, err := svc.CreateMapping(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	d.SetId(*result.Id)
	return resourcePingFederateOauthAccessTokenMappingsResourceReadResult(d, result)
}

func resourcePingFederateOauthAccessTokenMappingsResourceRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).OauthAccessTokenMappings
	input := pf.GetMappingInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetMapping(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return resourcePingFederateOauthAccessTokenMappingsResourceReadResult(d, result)
}

func resourcePingFederateOauthAccessTokenMappingsResourceUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).OauthAccessTokenMappings
	input := pf.UpdateMappingInput{
		Id:   d.Id(),
		Body: *resourcePingFederateOauthAccessTokenMappingsResourceReadData(d),
	}
	result, _, err := svc.UpdateMapping(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return resourcePingFederateOauthAccessTokenMappingsResourceReadResult(d, result)
}

func resourcePingFederateOauthAccessTokenMappingsResourceDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).OauthAccessTokenMappings
	input := pf.DeleteMappingInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteMapping(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func resourcePingFederateOauthAccessTokenMappingsResourceReadResult(d *schema.ResourceData, rv *pf.AccessTokenMapping) (err error) {
	if rv.AccessTokenManagerRef != nil {
		if err = d.Set("access_token_manager_ref", flattenResourceLink(rv.AccessTokenManagerRef)); err != nil {
			return err
		}
	}
	if rv.AttributeContractFulfillment != nil {
		if err = d.Set("attribute_contract_fulfillment", flattenMapOfAttributeFulfillmentValue(rv.AttributeContractFulfillment)); err != nil {
			return err
		}
	}
	if rv.IssuanceCriteria != nil && (rv.IssuanceCriteria.ExpressionCriteria != nil && rv.IssuanceCriteria.ConditionalCriteria != nil) {
		if err = d.Set("issuance_criteria", flattenIssuanceCriteria(rv.IssuanceCriteria)); err != nil {
			return err
		}
	}
	if rv.Context != nil {
		if err = d.Set("context", flattenAccessTokenMappingContext(rv.Context)); err != nil {
			return err
		}
	}

	if err := flattenAttributeSources(d, rv.AttributeSources); err != nil {
		return err
	}
	return nil
}

func resourcePingFederateOauthAccessTokenMappingsResourceReadData(d *schema.ResourceData) *pf.AccessTokenMapping {
	mapping := &pf.AccessTokenMapping{
		Id:                           String(d.Id()),
		Context:                      expandAccessTokenMappingContext(d.Get("context").([]interface{})),
		AccessTokenManagerRef:        expandResourceLink(d.Get("access_token_manager_ref").([]interface{})),
		AttributeContractFulfillment: expandMapOfAttributeFulfillmentValue(d.Get("attribute_contract_fulfillment").(*schema.Set).List()),
	}
	if v, ok := d.GetOk("issuance_criteria"); ok {
		mapping.IssuanceCriteria = expandIssuanceCriteria(v.([]interface{}))
	}
	if v, ok := d.GetOk("ldap_attribute_source"); ok && len(v.([]interface{})) > 0 {
		*mapping.AttributeSources = append(*mapping.AttributeSources, *expandLdapAttributeSource(v.([]interface{}))...)
	}
	if v, ok := d.GetOk("jdbc_attribute_source"); ok && len(v.([]interface{})) > 0 {
		*mapping.AttributeSources = append(*mapping.AttributeSources, *expandJdbcAttributeSource(v.([]interface{}))...)
	}
	if v, ok := d.GetOk("custom_attribute_source"); ok && len(v.([]interface{})) > 0 {
		*mapping.AttributeSources = append(*mapping.AttributeSources, *expandCustomAttributeSource(v.([]interface{}))...)
	}
	return mapping
}