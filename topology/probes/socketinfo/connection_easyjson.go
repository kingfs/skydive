// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package socketinfo

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

func easyjson28548e40DecodeGithubComSkydiveProjectSkydiveTopologyProbesSocketinfo(in *jlexer.Lexer, out *ConnectionInfo) {
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
		case "LocalAddress":
			out.LocalAddress = string(in.String())
		case "LocalPort":
			out.LocalPort = int64(in.Int64())
		case "RemoteAddress":
			out.RemoteAddress = string(in.String())
		case "RemotePort":
			out.RemotePort = int64(in.Int64())
		case "Protocol":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Protocol).UnmarshalJSON(data))
			}
		case "State":
			out.State = ConnectionState(in.String())
		case "Process":
			out.Process = string(in.String())
		case "Pid":
			out.Pid = int64(in.Int64())
		case "Name":
			out.Name = string(in.String())
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
func easyjson28548e40EncodeGithubComSkydiveProjectSkydiveTopologyProbesSocketinfo(out *jwriter.Writer, in ConnectionInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"LocalAddress\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LocalAddress))
	}
	{
		const prefix string = ",\"LocalPort\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.LocalPort))
	}
	{
		const prefix string = ",\"RemoteAddress\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RemoteAddress))
	}
	{
		const prefix string = ",\"RemotePort\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.RemotePort))
	}
	{
		const prefix string = ",\"Protocol\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Protocol).MarshalJSON())
	}
	{
		const prefix string = ",\"State\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.State))
	}
	{
		const prefix string = ",\"Process\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Process))
	}
	{
		const prefix string = ",\"Pid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Pid))
	}
	{
		const prefix string = ",\"Name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ConnectionInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson28548e40EncodeGithubComSkydiveProjectSkydiveTopologyProbesSocketinfo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ConnectionInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson28548e40EncodeGithubComSkydiveProjectSkydiveTopologyProbesSocketinfo(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ConnectionInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson28548e40DecodeGithubComSkydiveProjectSkydiveTopologyProbesSocketinfo(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ConnectionInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson28548e40DecodeGithubComSkydiveProjectSkydiveTopologyProbesSocketinfo(l, v)
}
