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

func easyjson105c6686DecodeGithubComSkydiveProjectSkydiveTopology(in *jlexer.Lexer, out *NextHop) {
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
		case "Priority":
			out.Priority = int64(in.Int64())
		case "IP":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.IP).UnmarshalText(data))
			}
		case "MAC":
			out.MAC = string(in.String())
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
func easyjson105c6686EncodeGithubComSkydiveProjectSkydiveTopology(out *jwriter.Writer, in NextHop) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Priority\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Priority))
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
func (v NextHop) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson105c6686EncodeGithubComSkydiveProjectSkydiveTopology(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NextHop) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson105c6686EncodeGithubComSkydiveProjectSkydiveTopology(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NextHop) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson105c6686DecodeGithubComSkydiveProjectSkydiveTopology(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NextHop) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson105c6686DecodeGithubComSkydiveProjectSkydiveTopology(l, v)
}
