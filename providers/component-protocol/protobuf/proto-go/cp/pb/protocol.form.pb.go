// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: protocol.proto

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
var _ urlenc.URLValuesUnmarshaler = (*ComponentProtocol)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Hierarchy)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Component)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Scenario)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ComponentEvent)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DebugOptions)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ProtocolOptions)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RenderRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RenderResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*IdentityInfo)(nil)

// ComponentProtocol implement urlenc.URLValuesUnmarshaler.
func (m *ComponentProtocol) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "version":
				m.Version = vals[0]
			case "scenario":
				m.Scenario = vals[0]
			case "hierarchy":
				if m.Hierarchy == nil {
					m.Hierarchy = &Hierarchy{}
				}
			case "hierarchy.root":
				if m.Hierarchy == nil {
					m.Hierarchy = &Hierarchy{}
				}
				m.Hierarchy.Root = vals[0]
			case "options":
				if m.Options == nil {
					m.Options = &ProtocolOptions{}
				}
			case "options.syncIntervalSecond":
				if m.Options == nil {
					m.Options = &ProtocolOptions{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Options.SyncIntervalSecond = val
			}
		}
	}
	return nil
}

// Hierarchy implement urlenc.URLValuesUnmarshaler.
func (m *Hierarchy) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "root":
				m.Root = vals[0]
			}
		}
	}
	return nil
}

// Component implement urlenc.URLValuesUnmarshaler.
func (m *Component) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "type":
				m.Type = vals[0]
			case "name":
				m.Name = vals[0]
			case "props":
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
					m.Props = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Props = val
					} else {
						m.Props = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}

// Scenario implement urlenc.URLValuesUnmarshaler.
func (m *Scenario) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scenarioKey":
				m.ScenarioKey = vals[0]
			case "scenarioType":
				m.ScenarioType = vals[0]
			}
		}
	}
	return nil
}

// ComponentEvent implement urlenc.URLValuesUnmarshaler.
func (m *ComponentEvent) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "component":
				m.Component = vals[0]
			case "operation":
				m.Operation = vals[0]
			}
		}
	}
	return nil
}

// DebugOptions implement urlenc.URLValuesUnmarshaler.
func (m *DebugOptions) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "componentKey":
				m.ComponentKey = vals[0]
			}
		}
	}
	return nil
}

// ProtocolOptions implement urlenc.URLValuesUnmarshaler.
func (m *ProtocolOptions) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "syncIntervalSecond":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.SyncIntervalSecond = val
			}
		}
	}
	return nil
}

// RenderRequest implement urlenc.URLValuesUnmarshaler.
func (m *RenderRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scenario":
				if m.Scenario == nil {
					m.Scenario = &Scenario{}
				}
			case "scenario.scenarioKey":
				if m.Scenario == nil {
					m.Scenario = &Scenario{}
				}
				m.Scenario.ScenarioKey = vals[0]
			case "scenario.scenarioType":
				if m.Scenario == nil {
					m.Scenario = &Scenario{}
				}
				m.Scenario.ScenarioType = vals[0]
			case "event":
				if m.Event == nil {
					m.Event = &ComponentEvent{}
				}
			case "event.component":
				if m.Event == nil {
					m.Event = &ComponentEvent{}
				}
				m.Event.Component = vals[0]
			case "event.operation":
				if m.Event == nil {
					m.Event = &ComponentEvent{}
				}
				m.Event.Operation = vals[0]
			case "protocol":
				if m.Protocol == nil {
					m.Protocol = &ComponentProtocol{}
				}
			case "protocol.version":
				if m.Protocol == nil {
					m.Protocol = &ComponentProtocol{}
				}
				m.Protocol.Version = vals[0]
			case "protocol.scenario":
				if m.Protocol == nil {
					m.Protocol = &ComponentProtocol{}
				}
				m.Protocol.Scenario = vals[0]
			case "protocol.hierarchy":
				if m.Protocol == nil {
					m.Protocol = &ComponentProtocol{}
				}
				if m.Protocol.Hierarchy == nil {
					m.Protocol.Hierarchy = &Hierarchy{}
				}
			case "protocol.hierarchy.root":
				if m.Protocol == nil {
					m.Protocol = &ComponentProtocol{}
				}
				if m.Protocol.Hierarchy == nil {
					m.Protocol.Hierarchy = &Hierarchy{}
				}
				m.Protocol.Hierarchy.Root = vals[0]
			case "protocol.options":
				if m.Protocol == nil {
					m.Protocol = &ComponentProtocol{}
				}
				if m.Protocol.Options == nil {
					m.Protocol.Options = &ProtocolOptions{}
				}
			case "protocol.options.syncIntervalSecond":
				if m.Protocol == nil {
					m.Protocol = &ComponentProtocol{}
				}
				if m.Protocol.Options == nil {
					m.Protocol.Options = &ProtocolOptions{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Protocol.Options.SyncIntervalSecond = val
			case "debugOptions":
				if m.DebugOptions == nil {
					m.DebugOptions = &DebugOptions{}
				}
			case "debugOptions.componentKey":
				if m.DebugOptions == nil {
					m.DebugOptions = &DebugOptions{}
				}
				m.DebugOptions.ComponentKey = vals[0]
			}
		}
	}
	return nil
}

// RenderResponse implement urlenc.URLValuesUnmarshaler.
func (m *RenderResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scenario":
				if m.Scenario == nil {
					m.Scenario = &Scenario{}
				}
			case "scenario.scenarioKey":
				if m.Scenario == nil {
					m.Scenario = &Scenario{}
				}
				m.Scenario.ScenarioKey = vals[0]
			case "scenario.scenarioType":
				if m.Scenario == nil {
					m.Scenario = &Scenario{}
				}
				m.Scenario.ScenarioType = vals[0]
			case "protocol":
				if m.Protocol == nil {
					m.Protocol = &ComponentProtocol{}
				}
			case "protocol.version":
				if m.Protocol == nil {
					m.Protocol = &ComponentProtocol{}
				}
				m.Protocol.Version = vals[0]
			case "protocol.scenario":
				if m.Protocol == nil {
					m.Protocol = &ComponentProtocol{}
				}
				m.Protocol.Scenario = vals[0]
			case "protocol.hierarchy":
				if m.Protocol == nil {
					m.Protocol = &ComponentProtocol{}
				}
				if m.Protocol.Hierarchy == nil {
					m.Protocol.Hierarchy = &Hierarchy{}
				}
			case "protocol.hierarchy.root":
				if m.Protocol == nil {
					m.Protocol = &ComponentProtocol{}
				}
				if m.Protocol.Hierarchy == nil {
					m.Protocol.Hierarchy = &Hierarchy{}
				}
				m.Protocol.Hierarchy.Root = vals[0]
			case "protocol.options":
				if m.Protocol == nil {
					m.Protocol = &ComponentProtocol{}
				}
				if m.Protocol.Options == nil {
					m.Protocol.Options = &ProtocolOptions{}
				}
			case "protocol.options.syncIntervalSecond":
				if m.Protocol == nil {
					m.Protocol = &ComponentProtocol{}
				}
				if m.Protocol.Options == nil {
					m.Protocol.Options = &ProtocolOptions{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Protocol.Options.SyncIntervalSecond = val
			}
		}
	}
	return nil
}

// IdentityInfo implement urlenc.URLValuesUnmarshaler.
func (m *IdentityInfo) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "userID":
				m.UserID = vals[0]
			case "internalClient":
				m.InternalClient = vals[0]
			case "orgID":
				m.OrgID = vals[0]
			}
		}
	}
	return nil
}
