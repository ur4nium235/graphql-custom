// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package genjson

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

func easyjson6ff3ac1dDecodeBaseInternalGenjson(in *jlexer.Lexer, out *ResponseGift) {
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
		case "result":
			if in.IsNull() {
				in.Skip()
				out.Result = nil
			} else {
				in.Delim('[')
				if out.Result == nil {
					if !in.IsDelim(']') {
						out.Result = make([]Gift, 0, 1)
					} else {
						out.Result = []Gift{}
					}
				} else {
					out.Result = (out.Result)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Gift
					(v1).UnmarshalEasyJSON(in)
					out.Result = append(out.Result, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "code":
			out.Code = int(in.Int())
		case "message":
			out.Message = string(in.String())
		case "status":
			out.Status = int(in.Int())
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
func easyjson6ff3ac1dEncodeBaseInternalGenjson(out *jwriter.Writer, in ResponseGift) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"result\":"
		out.RawString(prefix[1:])
		if in.Result == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Result {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix)
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.Int(int(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResponseGift) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeBaseInternalGenjson(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResponseGift) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeBaseInternalGenjson(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResponseGift) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeBaseInternalGenjson(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResponseGift) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeBaseInternalGenjson(l, v)
}
func easyjson6ff3ac1dDecodeBaseInternalGenjson1(in *jlexer.Lexer, out *ResponseClient) {
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
		case "status":
			out.Status = int(in.Int())
		case "message":
			out.Message = string(in.String())
		case "code":
			out.Code = int(in.Int())
		case "result":
			if in.IsNull() {
				in.Skip()
				out.Result = nil
			} else {
				in.Delim('[')
				if out.Result == nil {
					if !in.IsDelim(']') {
						out.Result = make([]Gift, 0, 1)
					} else {
						out.Result = []Gift{}
					}
				} else {
					out.Result = (out.Result)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Gift
					(v4).UnmarshalEasyJSON(in)
					out.Result = append(out.Result, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson6ff3ac1dEncodeBaseInternalGenjson1(out *jwriter.Writer, in ResponseClient) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Status))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix)
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"result\":"
		out.RawString(prefix)
		if in.Result == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Result {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResponseClient) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeBaseInternalGenjson1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResponseClient) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeBaseInternalGenjson1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResponseClient) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeBaseInternalGenjson1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResponseClient) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeBaseInternalGenjson1(l, v)
}
func easyjson6ff3ac1dDecodeBaseInternalGenjson2(in *jlexer.Lexer, out *Gift) {
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
		case "postID":
			out.PostID = string(in.String())
		case "version":
			out.Version = string(in.String())
		case "isSys":
			out.IsSys = int(in.Int())
		case "status":
			out.Status = int(in.Int())
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
func easyjson6ff3ac1dEncodeBaseInternalGenjson2(out *jwriter.Writer, in Gift) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"postID\":"
		out.RawString(prefix[1:])
		out.String(string(in.PostID))
	}
	{
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"isSys\":"
		out.RawString(prefix)
		out.Int(int(in.IsSys))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.Int(int(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Gift) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeBaseInternalGenjson2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Gift) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeBaseInternalGenjson2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Gift) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeBaseInternalGenjson2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Gift) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeBaseInternalGenjson2(l, v)
}
