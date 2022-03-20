package instana

import (
	"fmt"
	"github.com/gessnerfl/terraform-provider-instana/instana/restapi"
	"github.com/gessnerfl/terraform-provider-instana/instana/tagfilter"
	"github.com/gessnerfl/terraform-provider-instana/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

//ResourceInstanaApplicationAlertConfig the name of the terraform-provider-instana resource to manage application alert config
const ResourceInstanaApplicationAlertConfig = "instana_application_alert_config"

const (
	applicationAlertConfigFieldRuleMetricName                                 = "metric_name"
	applicationAlertConfigFieldRuleAggregation                                = "aggregation"
	applicationAlertConfigFieldRuleOperator                                   = "operator"
	applicationAlertConfigFieldAlertChannelIDs                                = "alerting_channel_ids"
	applicationAlertConfigFieldApplications                                   = "applications"
	applicationAlertConfigFieldApplicationsApplicationID                      = "application_id"
	applicationAlertConfigFieldApplicationsInclusive                          = "inclusive"
	applicationAlertConfigFieldApplicationsServices                           = "services"
	applicationAlertConfigFieldApplicationsServicesServiceId                  = "service_id"
	applicationAlertConfigFieldApplicationsServicesServiceEndpoints           = "endpoints"
	applicationAlertConfigFieldApplicationsServicesServiceEndpointsEndpointId = "endpoint_id"
	applicationAlertConfigFieldBoundaryScope                                  = "boundary_scope"
	applicationAlertConfigFieldCustomPayloadFields                            = "custom_payload_fields"
	applicationAlertConfigFieldCustomPayloadFieldsKey                         = "key"
	applicationAlertConfigFieldCustomPayloadFieldsValue                       = "value"
	applicationAlertConfigFieldDescription                                    = "description"
	applicationAlertConfigFieldEvaluationType                                 = "evaluation_type"
	applicationAlertConfigFieldGranularity                                    = "granularity"
	applicationAlertConfigFieldIncludeInternal                                = "include_internal"
	applicationAlertConfigFieldIncludeSynthetic                               = "include_synthetic"
	applicationAlertConfigFieldName                                           = "name"
	applicationAlertConfigFieldFullName                                       = "full_name"
	applicationAlertConfigFieldRule                                           = "rule"
	applicationAlertConfigFieldRuleErrorRate                                  = "error_rate"
	applicationAlertConfigFieldRuleLogs                                       = "logs"
	applicationAlertConfigFieldRuleLogsLevel                                  = "level"
	applicationAlertConfigFieldRuleLogsMessage                                = "message"
	applicationAlertConfigFieldRuleSlowness                                   = "slowness"
	applicationAlertConfigFieldRuleStatusCode                                 = "status_code"
	applicationAlertConfigFieldRuleStatusCodeStart                            = "status_code_start"
	applicationAlertConfigFieldRuleStatusCodeEnd                              = "status_code_end"
	applicationAlertConfigFieldRuleThroughput                                 = "throughput"
	applicationAlertConfigFieldSeverity                                       = "severity"
	applicationAlertConfigFieldTagFilter                                      = "tag_filter"
	applicationAlertConfigFieldThreshold                                      = "threshold"
	applicationAlertConfigFieldThresholdHistoricBaseline                      = "historic_baseline"
	applicationAlertConfigFieldThresholdHistoricBaselineBaseline              = "baseline"
	applicationAlertConfigFieldThresholdHistoricBaselineDeviationFactor       = "deviation_factor"
	applicationAlertConfigFieldThresholdHistoricBaselineSeasonality           = "seasonality"
	applicationAlertConfigFieldThresholdStatic                                = "static"
	applicationAlertConfigFieldThresholdStaticValue                           = "value"
	applicationAlertConfigFieldTimeThreshold                                  = "time_threshold"
	applicationAlertConfigFieldTimeThresholdRequestImpact                     = "request_impact"
	applicationAlertConfigFieldTimeThresholdRequestImpactRequests             = "requests"
	applicationAlertConfigFieldTimeThresholdViolationsInPeriod                = "violations_in_period"
	applicationAlertConfigFieldTimeThresholdViolationsInPeriodViolations      = "violations"
	applicationAlertConfigFieldTimeThresholdViolationsInSequence              = "violations_in_sequence"
	applicationAlertConfigFieldTriggering                                     = "triggering"

	applicationAlertConfigFieldThresholdOperator    = "operator"
	applicationAlertConfigFieldThresholdLastUpdated = "last_updated"

	applicationAlertConfigFieldTimeThresholdTimeWindow = "time_window"
)

var (
	applicationAlertRuleTypeKeys = []string{
		"rule.0.error_rate",
		"rule.0.logs",
		"rule.0.slowness",
		"rule.0.status_code",
		"rule.0.throughput",
	}

	applicationAlertSchemaRuleMetricName = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The metric name of the application alert rule",
	}

	applicationAlertSchemaRequiredRuleAggregation = &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(restapi.SupportedAggregations.ToStringSlice(), true),
		Description:  "The aggregation function of the application alert rule",
	}

	applicationAlertSchemaOptionalRuleAggregation = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringInSlice(restapi.SupportedAggregations.ToStringSlice(), true),
		Description:  "The aggregation function of the application alert rule",
	}

	applicationAlertSchemaRequiredRuleOperator = &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		Description:  "The operator which will be applied to evaluate this rule",
		ValidateFunc: validation.StringInSlice(restapi.SupportedExpressionOperators.ToStringSlice(), true),
	}

	applicationAlertThresholdTypeKeys = []string{
		"threshold.0.historic_baseline",
		"threshold.0.static",
	}

	applicationAlertSchemaRequiredThresholdOperator = &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		Description:  "The operator which will be applied to evaluate the threshold",
		ValidateFunc: validation.StringInSlice(restapi.SupportedThresholdOperators.ToStringSlice(), true),
	}

	applicationAlertSchemaOptionalThresholdLastUpdated = &schema.Schema{
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntAtLeast(0),
		Description:  "The last updated value of the threshold",
	}

	applicationAlertTimeThresholdTypeKeys = []string{
		"time_threshold.0.request_impact",
		"time_threshold.0.violations_in_period",
		"time_threshold.0.violations_in_sequence",
	}

	applicationAlertSchemaRequiredTimeThresholdTimeWindow = &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "The time window if the time threshold",
	}
)

//NewApplicationAlertConfigResourceHandle creates a new instance of the ResourceHandle for application alert configs
func NewApplicationAlertConfigResourceHandle() ResourceHandle {
	return &applicationAlertConfigResource{
		metaData: ResourceMetaData{
			ResourceName: ResourceInstanaApplicationAlertConfig,
			Schema: map[string]*schema.Schema{
				applicationAlertConfigFieldAlertChannelIDs: {
					Type:     schema.TypeSet,
					MinItems: 0,
					MaxItems: 1024,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "List of IDs of alert channels defined in Instana.",
				},
				applicationAlertConfigFieldApplications: {
					Type:     schema.TypeList,
					Required: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							applicationAlertConfigFieldApplicationsApplicationID: {
								Type:         schema.TypeString,
								Required:     true,
								Description:  "ID of the included application",
								ValidateFunc: validation.StringLenBetween(0, 64),
							},
							applicationAlertConfigFieldApplicationsInclusive: {
								Type:        schema.TypeBool,
								Optional:    true,
								Default:     false,
								Description: "Defines whether this node and his child nodes are included (true) or excluded (false)",
							},
							applicationAlertConfigFieldApplicationsServices: {
								Type:     schema.TypeList,
								Required: true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										applicationAlertConfigFieldApplicationsServicesServiceId: {
											Type:         schema.TypeString,
											Required:     true,
											Description:  "ID of the included service",
											ValidateFunc: validation.StringLenBetween(0, 64),
										},
										applicationAlertConfigFieldApplicationsInclusive: {
											Type:        schema.TypeBool,
											Optional:    true,
											Default:     false,
											Description: "Defines whether this node and his child nodes are included (true) or excluded (false)",
										},
										applicationAlertConfigFieldApplicationsServicesServiceEndpoints: {
											Type:     schema.TypeList,
											Required: true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													applicationAlertConfigFieldApplicationsServicesServiceEndpointsEndpointId: {
														Type:         schema.TypeString,
														Required:     true,
														Description:  "ID of the included endpoint",
														ValidateFunc: validation.StringLenBetween(0, 64),
													},
													applicationAlertConfigFieldApplicationsInclusive: {
														Type:        schema.TypeBool,
														Optional:    true,
														Default:     false,
														Description: "Defines whether this node and his child nodes are included (true) or excluded (false)",
													},
												},
											},
											Description: "Selection of endpoints in scope.",
										},
									},
								},
								Description: "Selection of services in scope.",
							},
						},
					},
					Description: "Selection of applications in scope.",
				},
				applicationAlertConfigFieldBoundaryScope: {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringInSlice(restapi.SupportedApplicationAlertConfigBoundaryScopes.ToStringSlice(), false),
					Description:  "The boundary scope of the application alert config",
				},
				applicationAlertConfigFieldCustomPayloadFields: {
					Type:     schema.TypeList,
					Optional: true,
					MinItems: 0,
					MaxItems: 20,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							applicationAlertConfigFieldCustomPayloadFieldsKey: {
								Type:        schema.TypeString,
								Required:    true,
								Description: "The key of the custom payload field",
							},
							applicationAlertConfigFieldCustomPayloadFieldsValue: {
								Type:        schema.TypeString,
								Required:    true,
								Description: "The value of the custom payload field",
							},
						},
					},
					Description: "An optional list of custom payload fields (static key/value pairs added to the event)",
				},
				applicationAlertConfigFieldDescription: {
					Type:         schema.TypeString,
					Required:     true,
					Description:  "The description text of the application alert config",
					ValidateFunc: validation.StringLenBetween(0, 65536),
				},
				applicationAlertConfigFieldEvaluationType: {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringInSlice(restapi.SupportedApplicationAlertEvaluationTypes.ToStringSlice(), false),
					Description:  "The evaluation type of the application alert config",
				},
				applicationAlertConfigFieldGranularity: {
					Type:         schema.TypeInt,
					Optional:     true,
					Default:      restapi.Granularity600000,
					ValidateFunc: validation.IntInSlice(restapi.SupportedGranularities.ToIntSlice()),
					Description:  "The evaluation granularity used for detection of violations of the defined threshold. In other words, it defines the size of the tumbling window used",
				},
				applicationAlertConfigFieldIncludeInternal: {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					Description: "Optional flag to indicate whether also internal calls are included in the scope or not. The default is false",
				},
				applicationAlertConfigFieldIncludeSynthetic: {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					Description: "Optional flag to indicate whether also synthetic calls are included in the scope or not. The default is false",
				},
				applicationAlertConfigFieldName: {
					Type:         schema.TypeString,
					Required:     true,
					Description:  "Name for the application alert configuration",
					ValidateFunc: validation.StringLenBetween(0, 256),
				},
				applicationAlertConfigFieldFullName: {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The full name field of the application alert config. The field is computed and contains the name which is sent to Instana. The computation depends on the configured default_name_prefix and default_name_suffix at provider level",
				},
				applicationAlertConfigFieldRule: {
					Type:        schema.TypeList,
					MinItems:    1,
					MaxItems:    1,
					Required:    true,
					Description: "Indicates the type of rule this alert configuration is about.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							applicationAlertConfigFieldRuleErrorRate: {
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Optional:    true,
								Description: "Rule based on the error rate of the configured alert configuration target",
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										applicationAlertConfigFieldRuleMetricName:  applicationAlertSchemaRuleMetricName,
										applicationAlertConfigFieldRuleAggregation: applicationAlertSchemaOptionalRuleAggregation,
									},
								},
								ExactlyOneOf: applicationAlertRuleTypeKeys,
							},
							applicationAlertConfigFieldRuleLogs: {
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Optional:    true,
								Description: "Rule based on logs of the configured alert configuration target",
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										applicationAlertConfigFieldRuleMetricName:  applicationAlertSchemaRuleMetricName,
										applicationAlertConfigFieldRuleAggregation: applicationAlertSchemaOptionalRuleAggregation,
										applicationAlertConfigFieldRuleLogsLevel: {
											Type:         schema.TypeString,
											Required:     true,
											Description:  "The log level for which this rule applies to",
											ValidateFunc: validation.StringInSlice(restapi.SupportedLogLevels.ToStringSlice(), true),
										},
										applicationAlertConfigFieldRuleLogsMessage: {
											Type:        schema.TypeString,
											Optional:    true,
											Description: "The log message for which this rule applies to",
										},
										applicationAlertConfigFieldRuleOperator: applicationAlertSchemaRequiredRuleOperator,
									},
								},
								ExactlyOneOf: applicationAlertRuleTypeKeys,
							},
							applicationAlertConfigFieldRuleSlowness: {
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Optional:    true,
								Description: "Rule based on the slowness of the configured alert configuration target",
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										applicationAlertConfigFieldRuleMetricName:  applicationAlertSchemaRuleMetricName,
										applicationAlertConfigFieldRuleAggregation: applicationAlertSchemaRequiredRuleAggregation,
									},
								},
								ExactlyOneOf: applicationAlertRuleTypeKeys,
							},
							applicationAlertConfigFieldRuleStatusCode: {
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Optional:    true,
								Description: "Rule based on the HTTP status code of the configured alert configuration target",
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										applicationAlertConfigFieldRuleMetricName:  applicationAlertSchemaRuleMetricName,
										applicationAlertConfigFieldRuleAggregation: applicationAlertSchemaOptionalRuleAggregation,
										applicationAlertConfigFieldRuleStatusCodeStart: {
											Type:         schema.TypeInt,
											Optional:     true,
											Description:  "minimal HTTP status code applied for this rule",
											ValidateFunc: validation.IntAtLeast(1),
										},
										applicationAlertConfigFieldRuleStatusCodeEnd: {
											Type:         schema.TypeInt,
											Optional:     true,
											Description:  "maximum HTTP status code applied for this rule",
											ValidateFunc: validation.IntAtLeast(1),
										},
									},
								},
								ExactlyOneOf: applicationAlertRuleTypeKeys,
							},
							applicationAlertConfigFieldRuleThroughput: {
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Optional:    true,
								Description: "Rule based on the throughput of the configured alert configuration target",
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										applicationAlertConfigFieldRuleMetricName:  applicationAlertSchemaRuleMetricName,
										applicationAlertConfigFieldRuleAggregation: applicationAlertSchemaOptionalRuleAggregation,
									},
								},
								ExactlyOneOf: applicationAlertRuleTypeKeys,
							},
						},
					},
				},
				applicationAlertConfigFieldSeverity: {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringInSlice(restapi.SupportedSeverities.TerraformRepresentations(), false),
					Description:  "The severity of the alert when triggered",
				},
				applicationAlertConfigFieldTagFilter: {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The tag filter of the application alert config",
					DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
						normalized, err := tagfilter.Normalize(new)
						if err == nil {
							return normalized == old
						}
						return old == new
					},
					StateFunc: func(val interface{}) string {
						normalized, err := tagfilter.Normalize(val.(string))
						if err == nil {
							return normalized
						}
						return val.(string)
					},
					ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
						v := val.(string)
						if _, err := tagfilter.NewParser().Parse(v); err != nil {
							errs = append(errs, fmt.Errorf("%q is not a valid tag filter; %s", key, err))
						}

						return
					},
				},
				applicationAlertConfigFieldThreshold: {
					Type:        schema.TypeList,
					MinItems:    1,
					MaxItems:    1,
					Required:    true,
					Description: "Indicates the type of threshold this alert rule is evaluated on.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							applicationAlertConfigFieldThresholdHistoricBaseline: {
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Optional:    true,
								Description: "Threshold based on a historic baseline.",
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										applicationAlertConfigFieldThresholdOperator:    applicationAlertSchemaRequiredThresholdOperator,
										applicationAlertConfigFieldThresholdLastUpdated: applicationAlertSchemaOptionalThresholdLastUpdated,
										applicationAlertConfigFieldThresholdHistoricBaselineBaseline: {
											Type:     schema.TypeSet,
											Optional: true,
											Elem: &schema.Schema{
												Type:     schema.TypeSet,
												Optional: false,
												Elem: &schema.Schema{
													Type: schema.TypeFloat,
												},
											},
											Description: "The baseline of the historic baseline threshold",
										},
										applicationAlertConfigFieldThresholdHistoricBaselineDeviationFactor: {
											Type:         schema.TypeFloat,
											Optional:     true,
											ValidateFunc: validation.FloatBetween(0.5, 16),
											Description:  "The baseline of the historic baseline threshold",
										},
										applicationAlertConfigFieldThresholdHistoricBaselineSeasonality: {
											Type:         schema.TypeString,
											Required:     true,
											ValidateFunc: validation.StringInSlice(restapi.SupportedThresholdSeasonalities.ToStringSlice(), true),
											Description:  "The seasonality of the historic baseline threshold",
										},
									},
								},
								ExactlyOneOf: applicationAlertThresholdTypeKeys,
							},
							applicationAlertConfigFieldThresholdStatic: {
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Optional:    true,
								Description: "Static threshold definition",
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										applicationAlertConfigFieldThresholdOperator:    applicationAlertSchemaRequiredThresholdOperator,
										applicationAlertConfigFieldThresholdLastUpdated: applicationAlertSchemaOptionalThresholdLastUpdated,
										applicationAlertConfigFieldThresholdStaticValue: {
											Type:        schema.TypeFloat,
											Optional:    true,
											Description: "The value of the static threshold",
										},
									},
								},
								ExactlyOneOf: applicationAlertThresholdTypeKeys,
							},
						},
					},
				},
				applicationAlertConfigFieldTimeThreshold: {
					Type:        schema.TypeList,
					MinItems:    1,
					MaxItems:    1,
					Required:    true,
					Description: "Indicates the type of violation of the defined threshold.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							applicationAlertConfigFieldTimeThresholdRequestImpact: {
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Optional:    true,
								Description: "Time threshold base on request impact",
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										applicationAlertConfigFieldTimeThresholdTimeWindow: applicationAlertSchemaRequiredTimeThresholdTimeWindow,
										applicationAlertConfigFieldTimeThresholdRequestImpactRequests: {
											Type:         schema.TypeInt,
											Optional:     true,
											ValidateFunc: validation.IntAtLeast(1),
											Description:  "The number of requests in the given window",
										},
									},
								},
								ExactlyOneOf: applicationAlertTimeThresholdTypeKeys,
							},
							applicationAlertConfigFieldTimeThresholdViolationsInPeriod: {
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Optional:    true,
								Description: "Time threshold base on violations in period",
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										applicationAlertConfigFieldTimeThresholdTimeWindow: applicationAlertSchemaRequiredTimeThresholdTimeWindow,
										applicationAlertConfigFieldTimeThresholdViolationsInPeriodViolations: {
											Type:         schema.TypeInt,
											Optional:     true,
											ValidateFunc: validation.IntBetween(1, 12),
											Description:  "The violations appeared in the period",
										},
									},
								},
								ExactlyOneOf: applicationAlertTimeThresholdTypeKeys,
							},
							applicationAlertConfigFieldTimeThresholdViolationsInSequence: {
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Optional:    true,
								Description: "Time threshold base on violations in sequence",
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										applicationAlertConfigFieldTimeThresholdTimeWindow: applicationAlertSchemaRequiredTimeThresholdTimeWindow,
									},
								},
								ExactlyOneOf: applicationAlertTimeThresholdTypeKeys,
							},
						},
					},
				},
				applicationAlertConfigFieldTriggering: {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					Description: "Optional flag to indicate whether also an Incident is triggered or not. The default is false",
				},
			},
		},
	}
}

type applicationAlertConfigResource struct {
	metaData ResourceMetaData
}

func (r *applicationAlertConfigResource) MetaData() *ResourceMetaData {
	return &r.metaData
}

func (r *applicationAlertConfigResource) StateUpgraders() []schema.StateUpgrader {
	return []schema.StateUpgrader{}
}

func (r *applicationAlertConfigResource) GetRestResource(api restapi.InstanaAPI) restapi.RestResource {
	return api.ApplicationAlertConfigs()
}

func (r *applicationAlertConfigResource) SetComputedFields(d *schema.ResourceData) {
	//No computed fields defined
}

func (r *applicationAlertConfigResource) UpdateState(d *schema.ResourceData, obj restapi.InstanaDataObject, formatter utils.ResourceNameFormatter) error {
	config := obj.(*restapi.ApplicationAlertConfig)

	name := formatter.UndoFormat(config.Name)
	severity, err := ConvertSeverityFromInstanaAPIToTerraformRepresentation(config.Severity)
	if err != nil {
		return err
	}
	var normalizedTagFilterString *string
	if config.TagFilterExpression != nil {
		normalizedTagFilterString, err = r.mapTagFilterExpressionToSchema(config.TagFilterExpression.(restapi.TagFilterExpressionElement))
		if err != nil {
			return err
		}
	}

	d.Set(applicationAlertConfigFieldAlertChannelIDs, config.AlertChannelIDs)
	d.Set(applicationAlertConfigFieldApplications, r.mapApplicationsToSchema(config))
	d.Set(applicationAlertConfigFieldBoundaryScope, config.BoundaryScope)
	d.Set(applicationAlertConfigFieldCustomPayloadFields, r.mapCustomPayloadFieldsToSchema(config))
	d.Set(applicationAlertConfigFieldDescription, config.Description)
	d.Set(applicationAlertConfigFieldEvaluationType, config.EvaluationType)
	d.Set(applicationAlertConfigFieldGranularity, config.Granularity)
	d.Set(applicationAlertConfigFieldIncludeInternal, config.IncludeInternal)
	d.Set(applicationAlertConfigFieldIncludeSynthetic, config.IncludeSynthetic)
	d.Set(applicationAlertConfigFieldName, name)
	d.Set(applicationAlertConfigFieldFullName, config.Name)
	d.Set(applicationAlertConfigFieldRule, r.mapRuleToSchema(config))
	d.Set(applicationAlertConfigFieldSeverity, severity)
	d.Set(applicationAlertConfigFieldTagFilter, normalizedTagFilterString)
	d.Set(applicationAlertConfigFieldThreshold, r.mapThresholdToSchema(config))
	d.Set(applicationAlertConfigFieldTimeThreshold, r.mapTimeThresholdToSchema(config))
	d.Set(applicationAlertConfigFieldTriggering, config.Triggering)
	return nil
}

func (r *applicationAlertConfigResource) mapApplicationsToSchema(config *restapi.ApplicationAlertConfig) []interface{} {
	result := make([]interface{}, len(config.Applications))
	i := 0
	for _, v := range config.Applications {
		result[i] = r.mapApplicationToSchema(&v)
		i++
	}
	return result
}

func (r *applicationAlertConfigResource) mapApplicationToSchema(app *restapi.IncludedApplication) map[string]interface{} {
	result := make(map[string]interface{})
	result[applicationAlertConfigFieldApplicationsApplicationID] = app.ApplicationID
	result[applicationAlertConfigFieldApplicationsInclusive] = app.Inclusive

	services := make([]interface{}, len(app.Services))
	i := 0
	for _, v := range app.Services {
		services[i] = r.mapServiceToSchema(&v)
		i++
	}
	result[applicationAlertConfigFieldApplicationsServices] = services
	return result
}

func (r *applicationAlertConfigResource) mapServiceToSchema(service *restapi.IncludedService) map[string]interface{} {
	result := make(map[string]interface{})
	result[applicationAlertConfigFieldApplicationsServicesServiceId] = service.ServiceID
	result[applicationAlertConfigFieldApplicationsInclusive] = service.Inclusive

	endpoints := make([]interface{}, len(service.Endpoints))
	i := 0
	for _, v := range service.Endpoints {
		endpoints[i] = r.mapEndpointToSchema(&v)
		i++
	}
	result[applicationAlertConfigFieldApplicationsServicesServiceEndpoints] = endpoints
	return result
}

func (r *applicationAlertConfigResource) mapEndpointToSchema(endpoint *restapi.IncludedEndpoint) map[string]interface{} {
	result := make(map[string]interface{})
	result[applicationAlertConfigFieldApplicationsServicesServiceEndpointsEndpointId] = endpoint.EndpointID
	result[applicationAlertConfigFieldApplicationsInclusive] = endpoint.Inclusive
	return result
}

func (r *applicationAlertConfigResource) mapCustomPayloadFieldsToSchema(config *restapi.ApplicationAlertConfig) []map[string]string {
	result := make([]map[string]string, len(config.CustomerPayloadFields))
	for i, v := range config.CustomerPayloadFields {
		field := make(map[string]string)
		field[applicationAlertConfigFieldCustomPayloadFieldsKey] = v.Key
		field[applicationAlertConfigFieldCustomPayloadFieldsValue] = v.Value
		result[i] = field
	}
	return result
}

func (r *applicationAlertConfigResource) mapRuleToSchema(config *restapi.ApplicationAlertConfig) []map[string]interface{} {
	ruleAttribute := make(map[string]interface{})
	ruleAttribute[applicationAlertConfigFieldRuleMetricName] = config.Rule.MetricName
	ruleAttribute[applicationAlertConfigFieldRuleAggregation] = config.Rule.Aggregation

	if config.Rule.StatusCodeStart != nil {
		ruleAttribute[applicationAlertConfigFieldRuleStatusCodeStart] = *config.Rule.StatusCodeStart
	}
	if config.Rule.StatusCodeEnd != nil {
		ruleAttribute[applicationAlertConfigFieldRuleStatusCodeEnd] = *config.Rule.StatusCodeEnd
	}
	if config.Rule.Level != nil {
		ruleAttribute[applicationAlertConfigFieldRuleLogsLevel] = *config.Rule.Level
	}
	if config.Rule.Message != nil {
		ruleAttribute[applicationAlertConfigFieldRuleLogsMessage] = *config.Rule.Message
	}

	alertType := r.mapAlertTypeToSchema(config.Rule.AlertType)
	rule := make(map[string]interface{})
	rule[alertType] = []interface{}{ruleAttribute}
	result := make([]map[string]interface{}, 1)
	result[0] = rule
	return result
}

func (r *applicationAlertConfigResource) mapAlertTypeToSchema(alertType string) string {
	if alertType == "errorRate" {
		return applicationAlertConfigFieldRuleErrorRate
	} else if alertType == "statusCode" {
		return applicationAlertConfigFieldRuleStatusCode
	}
	return alertType
}

func (r *applicationAlertConfigResource) mapTagFilterExpressionToSchema(input restapi.TagFilterExpressionElement) (*string, error) {
	mapper := tagfilter.NewMapper()
	expr, err := mapper.FromAPIModel(input)
	if err != nil {
		return nil, err
	}
	renderedExpression := expr.Render()
	return &renderedExpression, nil
}

func (r *applicationAlertConfigResource) mapThresholdToSchema(config *restapi.ApplicationAlertConfig) []map[string]interface{} {
	thresholdConfig := make(map[string]interface{})
	thresholdConfig[applicationAlertConfigFieldThresholdOperator] = config.Threshold.Operator
	thresholdConfig[applicationAlertConfigFieldThresholdLastUpdated] = config.Threshold.LastUpdated

	if config.Threshold.Value != nil {
		thresholdConfig[applicationAlertConfigFieldThresholdStaticValue] = *config.Threshold.Value
	}
	if config.Threshold.Baseline != nil {
		thresholdConfig[applicationAlertConfigFieldThresholdHistoricBaselineBaseline] = *config.Threshold.Baseline
	}
	if config.Threshold.DeviationFactor != nil {
		thresholdConfig[applicationAlertConfigFieldThresholdHistoricBaselineDeviationFactor] = *config.Threshold.DeviationFactor
	}
	if config.Threshold.Seasonality != nil {
		thresholdConfig[applicationAlertConfigFieldThresholdHistoricBaselineSeasonality] = *config.Threshold.Seasonality
	}

	thresholdType := r.mapThresholdTypeToSchema(config.Threshold.Type)
	threshold := make(map[string]interface{})
	threshold[thresholdType] = []interface{}{thresholdConfig}
	result := make([]map[string]interface{}, 1)
	result[0] = threshold
	return result
}

func (r *applicationAlertConfigResource) mapThresholdTypeToSchema(input string) string {
	if input == "historicBaseline" {
		return applicationAlertConfigFieldThresholdHistoricBaseline
	} else if input == "staticThreshold" {
		return applicationAlertConfigFieldThresholdStatic
	}
	return input
}

func (r *applicationAlertConfigResource) mapTimeThresholdToSchema(config *restapi.ApplicationAlertConfig) []map[string]interface{} {
	timeThresholdConfig := make(map[string]interface{})
	timeThresholdConfig[applicationAlertConfigFieldTimeThresholdTimeWindow] = config.TimeThreshold.TimeWindow

	if config.TimeThreshold.Violations != nil {
		timeThresholdConfig[applicationAlertConfigFieldTimeThresholdViolationsInPeriodViolations] = *config.TimeThreshold.Violations
	}
	if config.TimeThreshold.Requests != nil {
		timeThresholdConfig[applicationAlertConfigFieldTimeThresholdRequestImpactRequests] = *config.TimeThreshold.Requests
	}

	timeThresholdType := r.mapTimeThresholdTypeToSchema(config.TimeThreshold.Type)
	timeThreshold := make(map[string]interface{})
	timeThreshold[timeThresholdType] = []interface{}{timeThresholdConfig}
	result := make([]map[string]interface{}, 1)
	result[0] = timeThreshold
	return result
}

func (r *applicationAlertConfigResource) mapTimeThresholdTypeToSchema(input string) string {
	if input == "requestImpact" {
		return applicationAlertConfigFieldTimeThresholdRequestImpact
	} else if input == "violationsInPeriod" {
		return applicationAlertConfigFieldTimeThresholdViolationsInPeriod
	} else if input == "violationsInSequence" {
		return applicationAlertConfigFieldTimeThresholdViolationsInSequence
	}
	return input
}

func (r *applicationAlertConfigResource) MapStateToDataObject(d *schema.ResourceData, formatter utils.ResourceNameFormatter) (restapi.InstanaDataObject, error) {
	fullName := r.mapFullNameStringFromSchema(d, formatter)
	severity, err := ConvertSeverityFromTerraformToInstanaAPIRepresentation(d.Get(applicationAlertConfigFieldSeverity).(string))
	if err != nil {
		return nil, err
	}

	var tagFilter restapi.TagFilterExpressionElement
	tagFilterStr, ok := d.GetOk(applicationAlertConfigFieldTagFilter)
	if ok {
		tagFilter, err = r.mapTagFilterExpressionFromSchema(tagFilterStr.(string))
		if err != nil {
			return &restapi.ApplicationConfig{}, err
		}
	}

	return &restapi.ApplicationAlertConfig{
		ID:                    d.Id(),
		AlertChannelIDs:       ReadStringSetParameterFromResource(d, applicationAlertConfigFieldAlertChannelIDs),
		Applications:          r.mapApplicationsFromSchema(d),
		BoundaryScope:         restapi.BoundaryScope(d.Get(applicationAlertConfigFieldBoundaryScope).(string)),
		CustomerPayloadFields: r.mapCustomPayloadFieldsFromSchema(d),
		Description:           d.Get(applicationAlertConfigFieldDescription).(string),
		EvaluationType:        restapi.ApplicationAlertEvaluationType(d.Get(applicationAlertConfigFieldEvaluationType).(string)),
		Granularity:           restapi.Granularity(d.Get(applicationAlertConfigFieldGranularity).(int32)),
		IncludeInternal:       d.Get(applicationAlertConfigFieldIncludeInternal).(bool),
		IncludeSynthetic:      d.Get(applicationAlertConfigFieldIncludeSynthetic).(bool),
		Name:                  fullName,
		Rule:                  r.mapRuleFromSchema(d),
		Severity:              severity,
		TagFilterExpression:   tagFilter,
		Threshold:             r.mapThresholdFromSchema(d),
		TimeThreshold:         r.mapTimeThresholdFromSchema(d),
		Triggering:            d.Get(applicationAlertConfigFieldTriggering).(bool),
	}, nil
}

func (r *applicationAlertConfigResource) mapFullNameStringFromSchema(d *schema.ResourceData, formatter utils.ResourceNameFormatter) string {
	if d.HasChange(applicationAlertConfigFieldName) {
		return formatter.Format(d.Get(applicationAlertConfigFieldName).(string))
	}
	return d.Get(applicationAlertConfigFieldFullName).(string)
}

func (r *applicationAlertConfigResource) mapApplicationsFromSchema(d *schema.ResourceData) map[string]restapi.IncludedApplication {
	val := d.Get(applicationAlertConfigFieldApplications)
	result := make(map[string]restapi.IncludedApplication)
	if val != nil {
		for _, v := range val.([]interface{}) {
			app := r.mapApplicationFromSchema(v.(map[string]interface{}))
			result[app.ApplicationID] = app
		}
	}
	return result
}

func (r *applicationAlertConfigResource) mapApplicationFromSchema(appData map[string]interface{}) restapi.IncludedApplication {
	services := make(map[string]restapi.IncludedService)
	if appData[applicationAlertConfigFieldApplicationsServices] != nil {
		for _, v := range appData[applicationAlertConfigFieldApplicationsServices].([]interface{}) {
			service := r.mapServiceFromSchema(v.(map[string]interface{}))
			services[service.ServiceID] = service
		}
	}
	return restapi.IncludedApplication{
		ApplicationID: appData[applicationAlertConfigFieldApplicationsApplicationID].(string),
		Inclusive:     appData[applicationAlertConfigFieldApplicationsInclusive].(bool),
		Services:      services,
	}
}

func (r *applicationAlertConfigResource) mapServiceFromSchema(appData map[string]interface{}) restapi.IncludedService {
	endpoints := make(map[string]restapi.IncludedEndpoint)
	if appData[applicationAlertConfigFieldApplicationsServicesServiceEndpoints] != nil {
		for _, v := range appData[applicationAlertConfigFieldApplicationsServicesServiceEndpoints].([]interface{}) {
			endpoint := r.mapEndpointFromSchema(v.(map[string]interface{}))
			endpoints[endpoint.EndpointID] = endpoint
		}
	}
	return restapi.IncludedService{
		ServiceID: appData[applicationAlertConfigFieldApplicationsServicesServiceId].(string),
		Inclusive: appData[applicationAlertConfigFieldApplicationsInclusive].(bool),
		Endpoints: endpoints,
	}
}

func (r *applicationAlertConfigResource) mapEndpointFromSchema(appData map[string]interface{}) restapi.IncludedEndpoint {
	return restapi.IncludedEndpoint{
		EndpointID: appData[applicationAlertConfigFieldApplicationsServicesServiceEndpointsEndpointId].(string),
		Inclusive:  appData[applicationAlertConfigFieldApplicationsInclusive].(bool),
	}
}

func (r *applicationAlertConfigResource) mapCustomPayloadFieldsFromSchema(d *schema.ResourceData) []restapi.StaticStringField {
	val := d.Get(applicationAlertConfigFieldCustomPayloadFields)
	if val != nil {
		fields := val.([]map[string]string)
		result := make([]restapi.StaticStringField, len(fields))
		for i, v := range fields {
			result[i] = restapi.StaticStringField{
				Key:   v[applicationAlertConfigFieldCustomPayloadFieldsKey],
				Value: v[applicationAlertConfigFieldCustomPayloadFieldsValue],
			}
		}
		return result
	}
	return []restapi.StaticStringField{}
}

func (r *applicationAlertConfigResource) mapRuleFromSchema(d *schema.ResourceData) restapi.ApplicationAlertRule {
	rule := d.Get(applicationAlertConfigFieldRule).([]map[string]interface{})[0]
	for alertType, v := range rule {
		config := v.(map[string]interface{})
		var levelPtr *restapi.LogLevel
		levelString := GetStringPointerFromResourceData(d, applicationAlertConfigFieldRuleLogsLevel)
		if levelString != nil {
			level := restapi.LogLevel(*levelString)
			levelPtr = &level
		}
		return restapi.ApplicationAlertRule{
			AlertType:       r.mapAlertTypeFromSchema(alertType),
			MetricName:      config[applicationAlertConfigFieldRuleMetricName].(string),
			Aggregation:     restapi.Aggregation(config[applicationAlertConfigFieldRuleAggregation].(string)),
			StatusCodeStart: GetInt32PointerFromResourceData(d, applicationAlertConfigFieldRuleStatusCodeStart),
			StatusCodeEnd:   GetInt32PointerFromResourceData(d, applicationAlertConfigFieldRuleStatusCodeEnd),
			Level:           levelPtr,
			Message:         GetStringPointerFromResourceData(d, applicationAlertConfigFieldRuleLogsMessage),
		}
	}
	return restapi.ApplicationAlertRule{}
}

func (r *applicationAlertConfigResource) mapAlertTypeFromSchema(alertType string) string {
	if alertType == applicationAlertConfigFieldRuleErrorRate {
		return "errorRate"
	} else if alertType == applicationAlertConfigFieldRuleStatusCode {
		return "statusCode"
	}
	return alertType
}

func (r *applicationAlertConfigResource) mapTagFilterExpressionFromSchema(input string) (restapi.TagFilterExpressionElement, error) {
	parser := tagfilter.NewParser()
	expr, err := parser.Parse(input)
	if err != nil {
		return nil, err
	}

	mapper := tagfilter.NewMapper()
	return mapper.ToAPIModel(expr), nil

}

func (r *applicationAlertConfigResource) mapThresholdFromSchema(d *schema.ResourceData) restapi.Threshold {
	threshold := d.Get(applicationAlertConfigFieldThreshold).([]map[string]interface{})[0]
	for thresholdType, v := range threshold {
		config := v.(map[string]interface{})
		var seasonalityPtr *restapi.ThresholdSeasonality
		if v, ok := config[applicationAlertConfigFieldThresholdHistoricBaselineSeasonality]; ok {
			seasonality := restapi.ThresholdSeasonality(v.(string))
			seasonalityPtr = &seasonality
		}
		return restapi.Threshold{
			Type:            r.mapThresholdTypeToSchema(thresholdType),
			Operator:        restapi.ThresholdOperator(config[applicationAlertConfigFieldThresholdOperator].(string)),
			LastUpdated:     config[applicationAlertConfigFieldThresholdLastUpdated].(int64),
			Value:           GetFloat64PointerFromResourceData(d, applicationAlertConfigFieldThresholdStaticValue),
			DeviationFactor: GetFloat32PointerFromResourceData(d, applicationAlertConfigFieldThresholdHistoricBaselineDeviationFactor),
			Baseline:        config[applicationAlertConfigFieldThresholdHistoricBaselineBaseline].(*[][]float64),
			Seasonality:     seasonalityPtr,
		}
	}
	return restapi.Threshold{}
}

func (r *applicationAlertConfigResource) mapThresholdTypeFromSchema(input string) string {
	if input == applicationAlertConfigFieldThresholdHistoricBaseline {
		return "historicBaseline"
	} else if input == applicationAlertConfigFieldThresholdStatic {
		return "staticThreshold"
	}
	return input
}

func (r *applicationAlertConfigResource) mapTimeThresholdFromSchema(d *schema.ResourceData) restapi.TimeThreshold {
	timeThreshold := d.Get(applicationAlertConfigFieldThreshold).([]map[string]interface{})[0]
	for timeThresholdType, v := range timeThreshold {
		config := v.(map[string]interface{})
		return restapi.TimeThreshold{
			Type:       r.mapTimeThresholdTypeFromSchema(timeThresholdType),
			TimeWindow: config[applicationAlertConfigFieldTimeThresholdTimeWindow].(int64),
			Violations: GetInt32PointerFromResourceData(d, applicationAlertConfigFieldTimeThresholdViolationsInPeriodViolations),
			Requests:   GetInt32PointerFromResourceData(d, applicationAlertConfigFieldTimeThresholdRequestImpactRequests),
		}
	}
	return restapi.TimeThreshold{}
}

func (r *applicationAlertConfigResource) mapTimeThresholdTypeFromSchema(input string) string {
	if input == applicationAlertConfigFieldTimeThresholdRequestImpact {
		return "requestImpact"
	} else if input == applicationAlertConfigFieldTimeThresholdViolationsInPeriod {
		return "violationsInPeriod"
	} else if input == applicationAlertConfigFieldTimeThresholdViolationsInSequence {
		return "violationsInSequence"
	}
	return input
}
