// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: settings.proto

package pb

import (
	json "encoding/json"
	url "net/url"
	strconv "strconv"
	strings "strings"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*GetSettingsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetSettingsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PutSettingsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PutSettingsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RegisterMonitorConfigRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*MonitorConfig)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RegisterMonitorConfigResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ConfigGroups)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ConfigGroup)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ConfigItem)(nil)

// GetSettingsRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetSettingsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "orgID":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.OrgID = val
			case "workspace":
				m.Workspace = vals[0]
			}
		}
	}
	return nil
}

// GetSettingsResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetSettingsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// PutSettingsRequest implement urlenc.URLValuesUnmarshaler.
func (m *PutSettingsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "orgID":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.OrgID = val
			}
		}
	}
	return nil
}

// PutSettingsResponse implement urlenc.URLValuesUnmarshaler.
func (m *PutSettingsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				m.Data = vals[0]
			}
		}
	}
	return nil
}

// RegisterMonitorConfigRequest implement urlenc.URLValuesUnmarshaler.
func (m *RegisterMonitorConfigRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "desc":
				m.Desc = vals[0]
			}
		}
	}
	return nil
}

// MonitorConfig implement urlenc.URLValuesUnmarshaler.
func (m *MonitorConfig) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "namespace":
				m.Namespace = vals[0]
			case "type":
				m.Type = vals[0]
			case "names":
				m.Names = vals[0]
			case "filters":
				m.Filters = vals[0]
			case "enable":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Enable = val
			case "desc":
				m.Desc = vals[0]
			}
		}
	}
	return nil
}

// RegisterMonitorConfigResponse implement urlenc.URLValuesUnmarshaler.
func (m *RegisterMonitorConfigResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				m.Data = vals[0]
			}
		}
	}
	return nil
}

// ConfigGroups implement urlenc.URLValuesUnmarshaler.
func (m *ConfigGroups) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ConfigGroup implement urlenc.URLValuesUnmarshaler.
func (m *ConfigGroup) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "name":
				m.Name = vals[0]
			}
		}
	}
	return nil
}

// ConfigItem implement urlenc.URLValuesUnmarshaler.
func (m *ConfigItem) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "name":
				m.Name = vals[0]
			case "value":
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Value = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Value = val
					} else {
						m.Value = structpb.NewStringValue(vals[0])
					}
				}
			case "type":
				m.Type = vals[0]
			case "unit":
				m.Unit = vals[0]
			}
		}
	}
	return nil
}
