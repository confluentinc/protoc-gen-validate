package goshared

const bytesTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
		{{ if $r.Len }}
			if m.maskHas(mask, "{{ $f.Name }}") && len({{ accessor . }}) != {{ $r.GetLen }} {
				return {{ err . "value length must be " $r.GetLen " bytes" }}
			}
		{{ else }}
			if m.maskHas(mask, "{{ $f.Name }}") && len({{ accessor . }}) != {{ $r.GetMinLen }} {
				return {{ err . "value length must be " $r.GetMinLen " bytes" }}
			}
		{{ end }}
	{{ else if $r.MinLen }}
		{{ if $r.MaxLen }}
			l := len({{ accessor . }})
			if m.maskHas(mask, "{{ $f.Name }}") && l < {{ $r.GetMinLen }} || l > {{ $r.GetMaxLen }} {
				return {{ err . "value length must be between " $r.GetMinLen " and " $r.GetMaxLen " bytes, inclusive" }}
			}
		{{ else }}
			if m.maskHas(mask, "{{ $f.Name }}") && len({{ accessor . }}) < {{ $r.GetMinLen }} {
				return {{ err . "value length must be at least " $r.GetMinLen " bytes" }}
			}
		{{ end }}
	{{ else if $r.MaxLen }}
		if m.maskHas(mask, "{{ $f.Name }}") && len({{ accessor . }}) > {{ $r.GetMaxLen }} {
			return {{ err . "value length must be at most " $r.GetMaxLen " bytes" }}
		}
	{{ end }}

	{{ if $r.Prefix }}
		if m.maskHas(mask, "{{ $f.Name }}") && !bytes.HasPrefix({{ accessor . }}, {{ lit $r.GetPrefix }}) {
			return {{ err . "value does not have prefix " (byteStr $r.GetPrefix) }}
		}
	{{ end }}

	{{ if $r.Suffix }}
		if m.maskHas(mask, "{{ $f.Name }}") && !bytes.HasSuffix({{ accessor . }}, {{ lit $r.GetSuffix }}) {
			return {{ err . "value does not have suffix " (byteStr $r.GetSuffix) }}
		}
	{{ end }}

	{{ if $r.Contains }}
		if m.maskHas(mask, "{{ $f.Name }}") && !bytes.Contains({{ accessor . }}, {{ lit $r.GetContains }}) {
			return {{ err . "value does not contain " (byteStr $r.GetContains) }}
		}
	{{ end }}

	{{ if $r.In }}
		if _, ok := {{ lookup $f "InLookup" }}[string({{ accessor . }})]; m.maskHas(mask, "{{ $f.Name }}") && !ok {
			return {{ err . "value must be in list " $r.In }}
		}
	{{ else if $r.NotIn }}
		if _, ok := {{ lookup $f "NotInLookup" }}[string({{ accessor . }})]; m.maskHas(mask, "{{ $f.Name }}") && ok {
			return {{ err . "value must not be in list " $r.NotIn }}
		}
	{{ end }}

	{{ if $r.Const }}
		if m.maskHas(mask, "{{ $f.Name }}") && !bytes.Equal({{ accessor . }}, {{ lit $r.Const }}) {
			return {{ err . "value must equal " $r.Const }}
		}
	{{ end }}

	{{ if $r.GetIp }}
		if ip := net.IP({{ accessor . }}); m.maskHas(mask, "{{ $f.Name }}") && ip.To16() == nil {
			return {{ err . "value must be a valid IP address" }}
		}
	{{ else if $r.GetIpv4 }}
		if ip := net.IP({{ accessor . }}); m.maskHas(mask, "{{ $f.Name }}") && ip.To4() == nil {
			return {{ err . "value must be a valid IPv4 address" }}
		}
	{{ else if $r.GetIpv6 }}
		if ip := net.IP({{ accessor . }}); m.maskHas(mask, "{{ $f.Name }}") && (ip.To16() == nil || ip.To4() != nil) {
			return {{ err . "value must be a valid IPv6 address" }}
		}
	{{ end }}

	{{ if $r.Pattern }}
	if m.maskHas(mask, "{{ $f.Name }}") && !{{ lookup $f "Pattern" }}.Match({{ accessor . }}) {
		return {{ err . "value does not match regex pattern " (lit $r.GetPattern) }}
	}
	{{ end }}
`
