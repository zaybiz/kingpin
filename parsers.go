package kingpin

import (
	"net"
	"net/url"
	"os"
	"time"
)

type Settings interface {
	SetValue(value Value)
}

type parserMixin struct {
	value    Value
	required bool
}

func (p *parserMixin) SetValue(value Value) {
	p.value = value
}

// String sets the parser to a string parser.
func (p *parserMixin) String() (target *string) {
	target = new(string)
	p.StringVar(target)
	return
}

// Strings appends multiple occurrences to a string slice.
func (p *parserMixin) Strings() (target *[]string) {
	target = new([]string)
	p.StringsVar(target)
	return
}

// StringMap provides key=value parsing into a map.
func (p *parserMixin) StringMap() (target *map[string]string) {
	target = &(map[string]string{})
	p.StringMapVar(target)
	return
}

// Bool sets the parser to a boolean parser. Supports --no-<X> to disable the flag.
func (p *parserMixin) Bool() (target *bool) {
	target = new(bool)
	p.BoolVar(target)
	return
}

// Int sets the parser to an int parser.
func (p *parserMixin) Int() (target *int) {
	target = new(int)
	p.IntVar(target)
	return
}

// Int64 parses an int64
func (p *parserMixin) Int64() (target *int64) {
	target = new(int64)
	p.Int64Var(target)
	return
}

// Uint64 parses a uint64
func (p *parserMixin) Uint64() (target *uint64) {
	target = new(uint64)
	p.Uint64Var(target)
	return
}

// Float sets the parser to a float64 parser.
func (p *parserMixin) Float() (target *float64) {
	target = new(float64)
	p.FloatVar(target)
	return
}

// Duration sets the parser to a time.Duration parser.
func (p *parserMixin) Duration() (target *time.Duration) {
	target = new(time.Duration)
	p.DurationVar(target)
	return
}

// IP sets the parser to a net.IP parser.
func (p *parserMixin) IP() (target *net.IP) {
	target = new(net.IP)
	p.IPVar(target)
	return
}

// TCP (host:port) address.
func (p *parserMixin) TCP() (target **net.TCPAddr) {
	target = new(*net.TCPAddr)
	p.TCPVar(target)
	return
}

// TCPVar (host:port) address.
func (p *parserMixin) TCPVar(target **net.TCPAddr) {
	p.SetValue(newTCPAddrValue(target))
}

// TCP (host:port) address list.
func (p *parserMixin) TCPList() (target *[]*net.TCPAddr) {
	target = new([]*net.TCPAddr)
	p.TCPListVar(target)
	return
}

// TCPVar (host:port) address list.
func (p *parserMixin) TCPListVar(target *[]*net.TCPAddr) {
	p.SetValue(newTCPAddrsValue(target))
}

// ExistingFile sets the parser to one that requires and returns an existing file.
func (p *parserMixin) ExistingFile() (target *string) {
	target = new(string)
	p.ExistingFileVar(target)
	return
}

// ExistingDir sets the parser to one that requires and returns an existing directory.
func (p *parserMixin) ExistingDir() (target *string) {
	target = new(string)
	p.ExistingDirVar(target)
	return
}

// File sets the parser to one that requires and opens a valid os.File.
func (p *parserMixin) File() (target **os.File) {
	target = new(*os.File)
	p.FileVar(target)
	return
}

// URL provides a valid, parsed url.URL.
func (p *parserMixin) URL() (target **url.URL) {
	target = new(*url.URL)
	p.URLVar(target)
	return
}

// String sets the parser to a string parser.
func (p *parserMixin) StringVar(target *string) {
	p.SetValue(newStringValue("", target))
}

// Strings appends multiple occurrences to a string slice.
func (p *parserMixin) StringsVar(target *[]string) {
	p.SetValue(newStringsValue(target))
}

// StringMap provides key=value parsing into a map.
func (p *parserMixin) StringMapVar(target *map[string]string) {
	p.SetValue(newStringMapValue(target))
}

// Bool sets the parser to a boolean parser. Supports --no-<X> to disable the flag.
func (p *parserMixin) BoolVar(target *bool) {
	p.SetValue(newBoolValue(false, target))
}

// Int sets the parser to an int parser.
func (p *parserMixin) IntVar(target *int) {
	p.SetValue(newIntValue(0, target))
}

// Int64 parses an int64
func (p *parserMixin) Int64Var(target *int64) {
	p.SetValue(newInt64Value(0, target))
}

// Uint64 parses a uint64
func (p *parserMixin) Uint64Var(target *uint64) {
	p.SetValue(newUint64Value(0, target))
}

// Float sets the parser to a float64 parser.
func (p *parserMixin) FloatVar(target *float64) {
	p.SetValue(newFloat64Value(0, target))
}

// Duration sets the parser to a time.Duration parser.
func (p *parserMixin) DurationVar(target *time.Duration) {
	p.SetValue(newDurationValue(time.Duration(0), target))
}

// IP sets the parser to a net.IP parser.
func (p *parserMixin) IPVar(target *net.IP) {
	p.SetValue(newIPValue(target))
}

// ExistingFile sets the parser to one that requires and returns an existing file.
func (p *parserMixin) ExistingFileVar(target *string) {
	p.SetValue(newFileStatValue(target, func(s os.FileInfo) bool { return !s.IsDir() }))
}

// ExistingDir sets the parser to one that requires and returns an existing directory.
func (p *parserMixin) ExistingDirVar(target *string) {
	p.SetValue(newFileStatValue(target, func(s os.FileInfo) bool { return s.IsDir() }))
}

// File sets the parser to one that requires and opens a valid os.File.
func (p *parserMixin) FileVar(target **os.File) {
	p.SetValue(newFileValue(target))
}

// URL provides a valid, parsed url.URL.
func (p *parserMixin) URLVar(target **url.URL) {
	p.SetValue(newURLValue(target))
}
