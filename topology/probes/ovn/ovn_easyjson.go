// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package ovn

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson36bd8e45DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvn(in *jlexer.Lexer, out *Metadata) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ExtID":
			(out.ExtID).UnmarshalEasyJSON(in)
		case "Options":
			(out.Options).UnmarshalEasyJSON(in)
		case "Action":
			out.Action = string(in.String())
		case "Direction":
			out.Direction = string(in.String())
		case "Log":
			out.Log = bool(in.Bool())
		case "Match":
			out.Match = string(in.String())
		case "Priority":
			out.Priority = int64(in.Int64())
		case "GatewayChassis":
			if in.IsNull() {
				in.Skip()
				out.GatewayChassis = nil
			} else {
				in.Delim('[')
				if out.GatewayChassis == nil {
					if !in.IsDelim(']') {
						out.GatewayChassis = make([]string, 0, 4)
					} else {
						out.GatewayChassis = []string{}
					}
				} else {
					out.GatewayChassis = (out.GatewayChassis)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.GatewayChassis = append(out.GatewayChassis, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "IPv6RAConfigs":
			(out.IPv6RAConfigs).UnmarshalEasyJSON(in)
		case "Networks":
			if in.IsNull() {
				in.Skip()
				out.Networks = nil
			} else {
				in.Delim('[')
				if out.Networks == nil {
					if !in.IsDelim(']') {
						out.Networks = make([]string, 0, 4)
					} else {
						out.Networks = []string{}
					}
				} else {
					out.Networks = (out.Networks)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Networks = append(out.Networks, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Peer":
			out.Peer = string(in.String())
		case "Addresses":
			if in.IsNull() {
				in.Skip()
				out.Addresses = nil
			} else {
				in.Delim('[')
				if out.Addresses == nil {
					if !in.IsDelim(']') {
						out.Addresses = make([]string, 0, 4)
					} else {
						out.Addresses = []string{}
					}
				} else {
					out.Addresses = (out.Addresses)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.Addresses = append(out.Addresses, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "PortSecurity":
			if in.IsNull() {
				in.Skip()
				out.PortSecurity = nil
			} else {
				in.Delim('[')
				if out.PortSecurity == nil {
					if !in.IsDelim(']') {
						out.PortSecurity = make([]string, 0, 4)
					} else {
						out.PortSecurity = []string{}
					}
				} else {
					out.PortSecurity = (out.PortSecurity)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.PortSecurity = append(out.PortSecurity, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "DHCPv4Options":
			out.DHCPv4Options = string(in.String())
		case "DHCPv6Options":
			out.DHCPv6Options = string(in.String())
		case "Type":
			out.Type = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson36bd8e45EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvn(out *jwriter.Writer, in Metadata) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.ExtID) != 0 {
		const prefix string = ",\"ExtID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.ExtID).MarshalEasyJSON(out)
	}
	if len(in.Options) != 0 {
		const prefix string = ",\"Options\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Options).MarshalEasyJSON(out)
	}
	if in.Action != "" {
		const prefix string = ",\"Action\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Action))
	}
	if in.Direction != "" {
		const prefix string = ",\"Direction\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Direction))
	}
	if in.Log {
		const prefix string = ",\"Log\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Log))
	}
	if in.Match != "" {
		const prefix string = ",\"Match\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Match))
	}
	if in.Priority != 0 {
		const prefix string = ",\"Priority\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Priority))
	}
	if len(in.GatewayChassis) != 0 {
		const prefix string = ",\"GatewayChassis\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.GatewayChassis {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	if len(in.IPv6RAConfigs) != 0 {
		const prefix string = ",\"IPv6RAConfigs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.IPv6RAConfigs).MarshalEasyJSON(out)
	}
	if len(in.Networks) != 0 {
		const prefix string = ",\"Networks\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v7, v8 := range in.Networks {
				if v7 > 0 {
					out.RawByte(',')
				}
				out.String(string(v8))
			}
			out.RawByte(']')
		}
	}
	if in.Peer != "" {
		const prefix string = ",\"Peer\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Peer))
	}
	if len(in.Addresses) != 0 {
		const prefix string = ",\"Addresses\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v9, v10 := range in.Addresses {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.String(string(v10))
			}
			out.RawByte(']')
		}
	}
	if len(in.PortSecurity) != 0 {
		const prefix string = ",\"PortSecurity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.PortSecurity {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	if in.DHCPv4Options != "" {
		const prefix string = ",\"DHCPv4Options\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DHCPv4Options))
	}
	if in.DHCPv6Options != "" {
		const prefix string = ",\"DHCPv6Options\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DHCPv6Options))
	}
	if in.Type != "" {
		const prefix string = ",\"Type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Metadata) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson36bd8e45EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvn(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Metadata) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson36bd8e45EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvn(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Metadata) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson36bd8e45DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvn(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Metadata) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson36bd8e45DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvn(l, v)
}
func easyjson36bd8e45DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvn1(in *jlexer.Lexer, out *LSPMetadata) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Addresses":
			if in.IsNull() {
				in.Skip()
				out.Addresses = nil
			} else {
				in.Delim('[')
				if out.Addresses == nil {
					if !in.IsDelim(']') {
						out.Addresses = make([]string, 0, 4)
					} else {
						out.Addresses = []string{}
					}
				} else {
					out.Addresses = (out.Addresses)[:0]
				}
				for !in.IsDelim(']') {
					var v13 string
					v13 = string(in.String())
					out.Addresses = append(out.Addresses, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "PortSecurity":
			if in.IsNull() {
				in.Skip()
				out.PortSecurity = nil
			} else {
				in.Delim('[')
				if out.PortSecurity == nil {
					if !in.IsDelim(']') {
						out.PortSecurity = make([]string, 0, 4)
					} else {
						out.PortSecurity = []string{}
					}
				} else {
					out.PortSecurity = (out.PortSecurity)[:0]
				}
				for !in.IsDelim(']') {
					var v14 string
					v14 = string(in.String())
					out.PortSecurity = append(out.PortSecurity, v14)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "DHCPv4Options":
			out.DHCPv4Options = string(in.String())
		case "DHCPv6Options":
			out.DHCPv6Options = string(in.String())
		case "Type":
			out.Type = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson36bd8e45EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvn1(out *jwriter.Writer, in LSPMetadata) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Addresses) != 0 {
		const prefix string = ",\"Addresses\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v15, v16 := range in.Addresses {
				if v15 > 0 {
					out.RawByte(',')
				}
				out.String(string(v16))
			}
			out.RawByte(']')
		}
	}
	if len(in.PortSecurity) != 0 {
		const prefix string = ",\"PortSecurity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v17, v18 := range in.PortSecurity {
				if v17 > 0 {
					out.RawByte(',')
				}
				out.String(string(v18))
			}
			out.RawByte(']')
		}
	}
	if in.DHCPv4Options != "" {
		const prefix string = ",\"DHCPv4Options\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DHCPv4Options))
	}
	if in.DHCPv6Options != "" {
		const prefix string = ",\"DHCPv6Options\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DHCPv6Options))
	}
	if in.Type != "" {
		const prefix string = ",\"Type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LSPMetadata) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson36bd8e45EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvn1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LSPMetadata) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson36bd8e45EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvn1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LSPMetadata) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson36bd8e45DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvn1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LSPMetadata) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson36bd8e45DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvn1(l, v)
}
func easyjson36bd8e45DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvn2(in *jlexer.Lexer, out *LRPMetadata) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "GatewayChassis":
			if in.IsNull() {
				in.Skip()
				out.GatewayChassis = nil
			} else {
				in.Delim('[')
				if out.GatewayChassis == nil {
					if !in.IsDelim(']') {
						out.GatewayChassis = make([]string, 0, 4)
					} else {
						out.GatewayChassis = []string{}
					}
				} else {
					out.GatewayChassis = (out.GatewayChassis)[:0]
				}
				for !in.IsDelim(']') {
					var v19 string
					v19 = string(in.String())
					out.GatewayChassis = append(out.GatewayChassis, v19)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "IPv6RAConfigs":
			(out.IPv6RAConfigs).UnmarshalEasyJSON(in)
		case "Networks":
			if in.IsNull() {
				in.Skip()
				out.Networks = nil
			} else {
				in.Delim('[')
				if out.Networks == nil {
					if !in.IsDelim(']') {
						out.Networks = make([]string, 0, 4)
					} else {
						out.Networks = []string{}
					}
				} else {
					out.Networks = (out.Networks)[:0]
				}
				for !in.IsDelim(']') {
					var v20 string
					v20 = string(in.String())
					out.Networks = append(out.Networks, v20)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Peer":
			out.Peer = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson36bd8e45EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvn2(out *jwriter.Writer, in LRPMetadata) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.GatewayChassis) != 0 {
		const prefix string = ",\"GatewayChassis\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v21, v22 := range in.GatewayChassis {
				if v21 > 0 {
					out.RawByte(',')
				}
				out.String(string(v22))
			}
			out.RawByte(']')
		}
	}
	if len(in.IPv6RAConfigs) != 0 {
		const prefix string = ",\"IPv6RAConfigs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.IPv6RAConfigs).MarshalEasyJSON(out)
	}
	if len(in.Networks) != 0 {
		const prefix string = ",\"Networks\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v23, v24 := range in.Networks {
				if v23 > 0 {
					out.RawByte(',')
				}
				out.String(string(v24))
			}
			out.RawByte(']')
		}
	}
	if in.Peer != "" {
		const prefix string = ",\"Peer\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Peer))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LRPMetadata) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson36bd8e45EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvn2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LRPMetadata) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson36bd8e45EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvn2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LRPMetadata) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson36bd8e45DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvn2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LRPMetadata) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson36bd8e45DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvn2(l, v)
}
func easyjson36bd8e45DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvn3(in *jlexer.Lexer, out *ACLMetadata) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Action":
			out.Action = string(in.String())
		case "Direction":
			out.Direction = string(in.String())
		case "Log":
			out.Log = bool(in.Bool())
		case "Match":
			out.Match = string(in.String())
		case "Priority":
			out.Priority = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson36bd8e45EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvn3(out *jwriter.Writer, in ACLMetadata) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Action != "" {
		const prefix string = ",\"Action\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Action))
	}
	if in.Direction != "" {
		const prefix string = ",\"Direction\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Direction))
	}
	if in.Log {
		const prefix string = ",\"Log\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Log))
	}
	if in.Match != "" {
		const prefix string = ",\"Match\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Match))
	}
	if in.Priority != 0 {
		const prefix string = ",\"Priority\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Priority))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ACLMetadata) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson36bd8e45EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvn3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ACLMetadata) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson36bd8e45EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvn3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ACLMetadata) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson36bd8e45DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvn3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ACLMetadata) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson36bd8e45DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvn3(l, v)
}
