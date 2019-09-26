// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDNSManagedZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSManagedZoneCreate,
		Read:   resourceDNSManagedZoneRead,
		Update: resourceDNSManagedZoneUpdate,
		Delete: resourceDNSManagedZoneDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDNSManagedZoneImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dns_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Managed by Terraform",
			},
			"dnssec_config": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_key_specs": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"algorithm": {
										Type:         schema.TypeString,
										Optional:     true,
										ForceNew:     true,
										ValidateFunc: validation.StringInSlice([]string{"ecdsap256sha256", "ecdsap384sha384", "rsasha1", "rsasha256", "rsasha512", ""}, false),
									},
									"key_length": {
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
									},
									"key_type": {
										Type:         schema.TypeString,
										Optional:     true,
										ForceNew:     true,
										ValidateFunc: validation.StringInSlice([]string{"keySigning", "zoneSigning", ""}, false),
									},
									"kind": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Default:  "dns#dnsKeySpec",
									},
								},
							},
						},
						"kind": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Default:  "dns#managedZoneDnsSecConfig",
						},
						"non_existence": {
							Type:         schema.TypeString,
							Computed:     true,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"nsec", "nsec3", ""}, false),
						},
						"state": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"off", "on", "transfer", ""}, false),
						},
					},
				},
			},
			"forwarding_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_name_servers": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     dnsManagedZoneForwardingConfigTargetNameServersSchema(),
							Set: func(v interface{}) int {
								raw := v.(map[string]interface{})
								if address, ok := raw["ipv4_address"]; ok {
									hashcode.String(address.(string))
								}
								var buf bytes.Buffer
								schema.SerializeResourceForHash(&buf, raw, dnsManagedZoneForwardingConfigTargetNameServersSchema())
								return hashcode.String(buf.String())
							},
						},
					},
				},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"peering_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_network": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_url": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"private_visibility_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"networks": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     dnsManagedZonePrivateVisibilityConfigNetworksSchema(),
							Set: func(v interface{}) int {
								raw := v.(map[string]interface{})
								if url, ok := raw["network_url"]; ok {
									return selfLinkNameHash(url)
								}
								var buf bytes.Buffer
								schema.SerializeResourceForHash(&buf, raw, dnsManagedZonePrivateVisibilityConfigNetworksSchema())
								return hashcode.String(buf.String())
							},
						},
					},
				},
			},
			"visibility": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateFunc:     validation.StringInSlice([]string{"private", "public", ""}, false),
				DiffSuppressFunc: caseDiffSuppress,
				Default:          "public",
			},
			"name_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func dnsManagedZonePrivateVisibilityConfigNetworksSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_url": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
		},
	}
}

func dnsManagedZoneForwardingConfigTargetNameServersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ipv4_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDNSManagedZoneCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandDNSManagedZoneDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	dnsNameProp, err := expandDNSManagedZoneDnsName(d.Get("dns_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dns_name"); !isEmptyValue(reflect.ValueOf(dnsNameProp)) && (ok || !reflect.DeepEqual(v, dnsNameProp)) {
		obj["dnsName"] = dnsNameProp
	}
	dnssecConfigProp, err := expandDNSManagedZoneDnssecConfig(d.Get("dnssec_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dnssec_config"); !isEmptyValue(reflect.ValueOf(dnssecConfigProp)) && (ok || !reflect.DeepEqual(v, dnssecConfigProp)) {
		obj["dnssecConfig"] = dnssecConfigProp
	}
	nameProp, err := expandDNSManagedZoneName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	labelsProp, err := expandDNSManagedZoneLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	visibilityProp, err := expandDNSManagedZoneVisibility(d.Get("visibility"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("visibility"); !isEmptyValue(reflect.ValueOf(visibilityProp)) && (ok || !reflect.DeepEqual(v, visibilityProp)) {
		obj["visibility"] = visibilityProp
	}
	privateVisibilityConfigProp, err := expandDNSManagedZonePrivateVisibilityConfig(d.Get("private_visibility_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("private_visibility_config"); !isEmptyValue(reflect.ValueOf(privateVisibilityConfigProp)) && (ok || !reflect.DeepEqual(v, privateVisibilityConfigProp)) {
		obj["privateVisibilityConfig"] = privateVisibilityConfigProp
	}
	forwardingConfigProp, err := expandDNSManagedZoneForwardingConfig(d.Get("forwarding_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("forwarding_config"); !isEmptyValue(reflect.ValueOf(forwardingConfigProp)) && (ok || !reflect.DeepEqual(v, forwardingConfigProp)) {
		obj["forwardingConfig"] = forwardingConfigProp
	}
	peeringConfigProp, err := expandDNSManagedZonePeeringConfig(d.Get("peering_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peering_config"); !isEmptyValue(reflect.ValueOf(peeringConfigProp)) && (ok || !reflect.DeepEqual(v, peeringConfigProp)) {
		obj["peeringConfig"] = peeringConfigProp
	}

	url, err := replaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/managedZones")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ManagedZone: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ManagedZone: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ManagedZone %q: %#v", d.Id(), res)

	return resourceDNSManagedZoneRead(d, meta)
}

func resourceDNSManagedZoneRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/managedZones/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("DNSManagedZone %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}

	if err := d.Set("description", flattenDNSManagedZoneDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("dns_name", flattenDNSManagedZoneDnsName(res["dnsName"], d)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("dnssec_config", flattenDNSManagedZoneDnssecConfig(res["dnssecConfig"], d)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("name", flattenDNSManagedZoneName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("name_servers", flattenDNSManagedZoneNameServers(res["nameServers"], d)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("labels", flattenDNSManagedZoneLabels(res["labels"], d)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("visibility", flattenDNSManagedZoneVisibility(res["visibility"], d)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("private_visibility_config", flattenDNSManagedZonePrivateVisibilityConfig(res["privateVisibilityConfig"], d)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("forwarding_config", flattenDNSManagedZoneForwardingConfig(res["forwardingConfig"], d)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("peering_config", flattenDNSManagedZonePeeringConfig(res["peeringConfig"], d)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}

	return nil
}

func resourceDNSManagedZoneUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	d.Partial(true)

	if d.HasChange("description") || d.HasChange("labels") || d.HasChange("private_visibility_config") || d.HasChange("forwarding_config") || d.HasChange("peering_config") {
		obj := make(map[string]interface{})
		descriptionProp, err := expandDNSManagedZoneDescription(d.Get("description"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
			obj["description"] = descriptionProp
		}
		labelsProp, err := expandDNSManagedZoneLabels(d.Get("labels"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
			obj["labels"] = labelsProp
		}
		privateVisibilityConfigProp, err := expandDNSManagedZonePrivateVisibilityConfig(d.Get("private_visibility_config"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("private_visibility_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, privateVisibilityConfigProp)) {
			obj["privateVisibilityConfig"] = privateVisibilityConfigProp
		}
		forwardingConfigProp, err := expandDNSManagedZoneForwardingConfig(d.Get("forwarding_config"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("forwarding_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, forwardingConfigProp)) {
			obj["forwardingConfig"] = forwardingConfigProp
		}
		peeringConfigProp, err := expandDNSManagedZonePeeringConfig(d.Get("peering_config"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("peering_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, peeringConfigProp)) {
			obj["peeringConfig"] = peeringConfigProp
		}

		url, err := replaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/managedZones/{{name}}")
		if err != nil {
			return err
		}
		_, err = sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating ManagedZone %q: %s", d.Id(), err)
		}

		d.SetPartial("description")
		d.SetPartial("labels")
		d.SetPartial("private_visibility_config")
		d.SetPartial("forwarding_config")
		d.SetPartial("peering_config")
	}

	d.Partial(false)

	return resourceDNSManagedZoneRead(d, meta)
}

func resourceDNSManagedZoneDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/managedZones/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ManagedZone %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ManagedZone")
	}

	log.Printf("[DEBUG] Finished deleting ManagedZone %q: %#v", d.Id(), res)
	return nil
}

func resourceDNSManagedZoneImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/managedZones/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDNSManagedZoneDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSManagedZoneDnsName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSManagedZoneDnssecConfig(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kind"] =
		flattenDNSManagedZoneDnssecConfigKind(original["kind"], d)
	transformed["non_existence"] =
		flattenDNSManagedZoneDnssecConfigNonExistence(original["nonExistence"], d)
	transformed["state"] =
		flattenDNSManagedZoneDnssecConfigState(original["state"], d)
	transformed["default_key_specs"] =
		flattenDNSManagedZoneDnssecConfigDefaultKeySpecs(original["defaultKeySpecs"], d)
	return []interface{}{transformed}
}
func flattenDNSManagedZoneDnssecConfigKind(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSManagedZoneDnssecConfigNonExistence(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSManagedZoneDnssecConfigState(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSManagedZoneDnssecConfigDefaultKeySpecs(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"algorithm":  flattenDNSManagedZoneDnssecConfigDefaultKeySpecsAlgorithm(original["algorithm"], d),
			"key_length": flattenDNSManagedZoneDnssecConfigDefaultKeySpecsKeyLength(original["keyLength"], d),
			"key_type":   flattenDNSManagedZoneDnssecConfigDefaultKeySpecsKeyType(original["keyType"], d),
			"kind":       flattenDNSManagedZoneDnssecConfigDefaultKeySpecsKind(original["kind"], d),
		})
	}
	return transformed
}
func flattenDNSManagedZoneDnssecConfigDefaultKeySpecsAlgorithm(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSManagedZoneDnssecConfigDefaultKeySpecsKeyLength(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenDNSManagedZoneDnssecConfigDefaultKeySpecsKeyType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSManagedZoneDnssecConfigDefaultKeySpecsKind(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSManagedZoneName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSManagedZoneNameServers(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSManagedZoneLabels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSManagedZoneVisibility(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil || v.(string) == "" {
		return "public"
	}
	return v
}

func flattenDNSManagedZonePrivateVisibilityConfig(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["networks"] =
		flattenDNSManagedZonePrivateVisibilityConfigNetworks(original["networks"], d)
	return []interface{}{transformed}
}
func flattenDNSManagedZonePrivateVisibilityConfigNetworks(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(func(v interface{}) int {
		raw := v.(map[string]interface{})
		if url, ok := raw["network_url"]; ok {
			return selfLinkNameHash(url)
		}
		var buf bytes.Buffer
		schema.SerializeResourceForHash(&buf, raw, dnsManagedZonePrivateVisibilityConfigNetworksSchema())
		return hashcode.String(buf.String())
	}, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"network_url": flattenDNSManagedZonePrivateVisibilityConfigNetworksNetworkUrl(original["networkUrl"], d),
		})
	}
	return transformed
}
func flattenDNSManagedZonePrivateVisibilityConfigNetworksNetworkUrl(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSManagedZoneForwardingConfig(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["target_name_servers"] =
		flattenDNSManagedZoneForwardingConfigTargetNameServers(original["targetNameServers"], d)
	return []interface{}{transformed}
}
func flattenDNSManagedZoneForwardingConfigTargetNameServers(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(func(v interface{}) int {
		raw := v.(map[string]interface{})
		if address, ok := raw["ipv4_address"]; ok {
			hashcode.String(address.(string))
		}
		var buf bytes.Buffer
		schema.SerializeResourceForHash(&buf, raw, dnsManagedZoneForwardingConfigTargetNameServersSchema())
		return hashcode.String(buf.String())
	}, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"ipv4_address": flattenDNSManagedZoneForwardingConfigTargetNameServersIpv4Address(original["ipv4Address"], d),
		})
	}
	return transformed
}
func flattenDNSManagedZoneForwardingConfigTargetNameServersIpv4Address(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSManagedZonePeeringConfig(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["target_network"] =
		flattenDNSManagedZonePeeringConfigTargetNetwork(original["targetNetwork"], d)
	return []interface{}{transformed}
}
func flattenDNSManagedZonePeeringConfigTargetNetwork(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["network_url"] =
		flattenDNSManagedZonePeeringConfigTargetNetworkNetworkUrl(original["networkUrl"], d)
	return []interface{}{transformed}
}
func flattenDNSManagedZonePeeringConfigTargetNetworkNetworkUrl(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandDNSManagedZoneDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnsName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKind, err := expandDNSManagedZoneDnssecConfigKind(original["kind"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKind); val.IsValid() && !isEmptyValue(val) {
		transformed["kind"] = transformedKind
	}

	transformedNonExistence, err := expandDNSManagedZoneDnssecConfigNonExistence(original["non_existence"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNonExistence); val.IsValid() && !isEmptyValue(val) {
		transformed["nonExistence"] = transformedNonExistence
	}

	transformedState, err := expandDNSManagedZoneDnssecConfigState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !isEmptyValue(val) {
		transformed["state"] = transformedState
	}

	transformedDefaultKeySpecs, err := expandDNSManagedZoneDnssecConfigDefaultKeySpecs(original["default_key_specs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultKeySpecs); val.IsValid() && !isEmptyValue(val) {
		transformed["defaultKeySpecs"] = transformedDefaultKeySpecs
	}

	return transformed, nil
}

func expandDNSManagedZoneDnssecConfigKind(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigNonExistence(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigState(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigDefaultKeySpecs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAlgorithm, err := expandDNSManagedZoneDnssecConfigDefaultKeySpecsAlgorithm(original["algorithm"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAlgorithm); val.IsValid() && !isEmptyValue(val) {
			transformed["algorithm"] = transformedAlgorithm
		}

		transformedKeyLength, err := expandDNSManagedZoneDnssecConfigDefaultKeySpecsKeyLength(original["key_length"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKeyLength); val.IsValid() && !isEmptyValue(val) {
			transformed["keyLength"] = transformedKeyLength
		}

		transformedKeyType, err := expandDNSManagedZoneDnssecConfigDefaultKeySpecsKeyType(original["key_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKeyType); val.IsValid() && !isEmptyValue(val) {
			transformed["keyType"] = transformedKeyType
		}

		transformedKind, err := expandDNSManagedZoneDnssecConfigDefaultKeySpecsKind(original["kind"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKind); val.IsValid() && !isEmptyValue(val) {
			transformed["kind"] = transformedKind
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSManagedZoneDnssecConfigDefaultKeySpecsAlgorithm(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigDefaultKeySpecsKeyLength(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigDefaultKeySpecsKeyType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigDefaultKeySpecsKind(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDNSManagedZoneVisibility(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZonePrivateVisibilityConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNetworks, err := expandDNSManagedZonePrivateVisibilityConfigNetworks(original["networks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworks); val.IsValid() && !isEmptyValue(val) {
		transformed["networks"] = transformedNetworks
	}

	return transformed, nil
}

func expandDNSManagedZonePrivateVisibilityConfigNetworks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetworkUrl, err := expandDNSManagedZonePrivateVisibilityConfigNetworksNetworkUrl(original["network_url"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetworkUrl); val.IsValid() && !isEmptyValue(val) {
			transformed["networkUrl"] = transformedNetworkUrl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSManagedZonePrivateVisibilityConfigNetworksNetworkUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneForwardingConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargetNameServers, err := expandDNSManagedZoneForwardingConfigTargetNameServers(original["target_name_servers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetNameServers); val.IsValid() && !isEmptyValue(val) {
		transformed["targetNameServers"] = transformedTargetNameServers
	}

	return transformed, nil
}

func expandDNSManagedZoneForwardingConfigTargetNameServers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpv4Address, err := expandDNSManagedZoneForwardingConfigTargetNameServersIpv4Address(original["ipv4_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpv4Address); val.IsValid() && !isEmptyValue(val) {
			transformed["ipv4Address"] = transformedIpv4Address
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSManagedZoneForwardingConfigTargetNameServersIpv4Address(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZonePeeringConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargetNetwork, err := expandDNSManagedZonePeeringConfigTargetNetwork(original["target_network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetNetwork); val.IsValid() && !isEmptyValue(val) {
		transformed["targetNetwork"] = transformedTargetNetwork
	}

	return transformed, nil
}

func expandDNSManagedZonePeeringConfigTargetNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNetworkUrl, err := expandDNSManagedZonePeeringConfigTargetNetworkNetworkUrl(original["network_url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworkUrl); val.IsValid() && !isEmptyValue(val) {
		transformed["networkUrl"] = transformedNetworkUrl
	}

	return transformed, nil
}

func expandDNSManagedZonePeeringConfigTargetNetworkNetworkUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
