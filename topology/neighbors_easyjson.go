// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package topology

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

func easyjsonB8b61a3fDecodeGithubComSkydiveProjectSkydiveTopology(in *jlexer.Lexer, out *Neighbors) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Neighbors, 0, 8)
			} else {
				*out = Neighbors{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *Neighbor
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(Neighbor)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB8b61a3fEncodeGithubComSkydiveProjectSkydiveTopology(out *jwriter.Writer, in Neighbors) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				(*v3).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Neighbors) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB8b61a3fEncodeGithubComSkydiveProjectSkydiveTopology(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Neighbors) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB8b61a3fEncodeGithubComSkydiveProjectSkydiveTopology(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Neighbors) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB8b61a3fDecodeGithubComSkydiveProjectSkydiveTopology(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Neighbors) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB8b61a3fDecodeGithubComSkydiveProjectSkydiveTopology(l, v)
}
func easyjsonB8b61a3fDecodeGithubComSkydiveProjectSkydiveTopology1(in *jlexer.Lexer, out *Neighbor) {
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
		case "Flags":
			if in.IsNull() {
				in.Skip()
				out.Flags = nil
			} else {
				in.Delim('[')
				if out.Flags == nil {
					if !in.IsDelim(']') {
						out.Flags = make([]string, 0, 4)
					} else {
						out.Flags = []string{}
					}
				} else {
					out.Flags = (out.Flags)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Flags = append(out.Flags, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "MAC":
			out.MAC = string(in.String())
		case "IP":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.IP).UnmarshalText(data))
			}
		case "State":
			if in.IsNull() {
				in.Skip()
				out.State = nil
			} else {
				in.Delim('[')
				if out.State == nil {
					if !in.IsDelim(']') {
						out.State = make([]string, 0, 4)
					} else {
						out.State = []string{}
					}
				} else {
					out.State = (out.State)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.State = append(out.State, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Vlan":
			out.Vlan = int64(in.Int64())
		case "VNI":
			out.VNI = int64(in.Int64())
		case "IfIndex":
			out.IfIndex = int64(in.Int64())
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
func easyjsonB8b61a3fEncodeGithubComSkydiveProjectSkydiveTopology1(out *jwriter.Writer, in Neighbor) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Flags) != 0 {
		const prefix string = ",\"Flags\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v6, v7 := range in.Flags {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	if in.MAC != "" {
		const prefix string = ",\"MAC\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MAC))
	}
	if len(in.IP) != 0 {
		const prefix string = ",\"IP\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.RawText((in.IP).MarshalText())
	}
	if len(in.State) != 0 {
		const prefix string = ",\"State\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.State {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	if in.Vlan != 0 {
		const prefix string = ",\"Vlan\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Vlan))
	}
	if in.VNI != 0 {
		const prefix string = ",\"VNI\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.VNI))
	}
	{
		const prefix string = ",\"IfIndex\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.IfIndex))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Neighbor) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB8b61a3fEncodeGithubComSkydiveProjectSkydiveTopology1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Neighbor) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB8b61a3fEncodeGithubComSkydiveProjectSkydiveTopology1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Neighbor) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB8b61a3fDecodeGithubComSkydiveProjectSkydiveTopology1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Neighbor) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB8b61a3fDecodeGithubComSkydiveProjectSkydiveTopology1(l, v)
}
