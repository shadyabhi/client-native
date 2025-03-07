// Code generated with struct_equal_generator; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// Equal checks if two structs of type ServerParams are equal
//
// By default empty maps and slices are equal to nil:
//
//	var a, b ServerParams
//	equal := a.Equal(b)
//
// For more advanced use case you can configure these options (default values are shown):
//
//	var a, b ServerParams
//	equal := a.Equal(b,Options{
//		NilSameAsEmpty: true,
//	})
func (s ServerParams) Equal(t ServerParams, opts ...Options) bool {
	opt := getOptions(opts...)

	if s.AgentAddr != t.AgentAddr {
		return false
	}

	if s.AgentCheck != t.AgentCheck {
		return false
	}

	if !equalPointers(s.AgentInter, t.AgentInter) {
		return false
	}

	if !equalPointers(s.AgentPort, t.AgentPort) {
		return false
	}

	if s.AgentSend != t.AgentSend {
		return false
	}

	if s.Allow0rtt != t.Allow0rtt {
		return false
	}

	if s.Alpn != t.Alpn {
		return false
	}

	if s.Backup != t.Backup {
		return false
	}

	if s.Check != t.Check {
		return false
	}

	if s.CheckSendProxy != t.CheckSendProxy {
		return false
	}

	if s.CheckSni != t.CheckSni {
		return false
	}

	if s.CheckSsl != t.CheckSsl {
		return false
	}

	if s.CheckAlpn != t.CheckAlpn {
		return false
	}

	if s.CheckProto != t.CheckProto {
		return false
	}

	if s.CheckViaSocks4 != t.CheckViaSocks4 {
		return false
	}

	if s.Ciphers != t.Ciphers {
		return false
	}

	if s.Ciphersuites != t.Ciphersuites {
		return false
	}

	if s.ClientSigalgs != t.ClientSigalgs {
		return false
	}

	if s.Cookie != t.Cookie {
		return false
	}

	if s.CrlFile != t.CrlFile {
		return false
	}

	if s.Curves != t.Curves {
		return false
	}

	if !equalPointers(s.Downinter, t.Downinter) {
		return false
	}

	if s.ErrorLimit != t.ErrorLimit {
		return false
	}

	if !equalPointers(s.Fall, t.Fall) {
		return false
	}

	if !equalPointers(s.Fastinter, t.Fastinter) {
		return false
	}

	if s.ForceSslv3 != t.ForceSslv3 {
		return false
	}

	if s.ForceTlsv10 != t.ForceTlsv10 {
		return false
	}

	if s.ForceTlsv11 != t.ForceTlsv11 {
		return false
	}

	if s.ForceTlsv12 != t.ForceTlsv12 {
		return false
	}

	if s.ForceTlsv13 != t.ForceTlsv13 {
		return false
	}

	if s.HealthCheckAddress != t.HealthCheckAddress {
		return false
	}

	if !equalPointers(s.HealthCheckPort, t.HealthCheckPort) {
		return false
	}

	if !equalPointers(s.InitAddr, t.InitAddr) {
		return false
	}

	if !equalPointers(s.Inter, t.Inter) {
		return false
	}

	if s.LogProto != t.LogProto {
		return false
	}

	if s.Maintenance != t.Maintenance {
		return false
	}

	if !equalPointers(s.MaxReuse, t.MaxReuse) {
		return false
	}

	if !equalPointers(s.Maxconn, t.Maxconn) {
		return false
	}

	if !equalPointers(s.Maxqueue, t.Maxqueue) {
		return false
	}

	if !equalPointers(s.Minconn, t.Minconn) {
		return false
	}

	if s.Namespace != t.Namespace {
		return false
	}

	if s.NoSslv3 != t.NoSslv3 {
		return false
	}

	if s.NoTlsv10 != t.NoTlsv10 {
		return false
	}

	if s.NoTlsv11 != t.NoTlsv11 {
		return false
	}

	if s.NoTlsv12 != t.NoTlsv12 {
		return false
	}

	if s.NoTlsv13 != t.NoTlsv13 {
		return false
	}

	if s.NoVerifyhost != t.NoVerifyhost {
		return false
	}

	if s.Npn != t.Npn {
		return false
	}

	if s.Observe != t.Observe {
		return false
	}

	if s.OnError != t.OnError {
		return false
	}

	if s.OnMarkedDown != t.OnMarkedDown {
		return false
	}

	if s.OnMarkedUp != t.OnMarkedUp {
		return false
	}

	if !equalPointers(s.PoolLowConn, t.PoolLowConn) {
		return false
	}

	if !equalPointers(s.PoolMaxConn, t.PoolMaxConn) {
		return false
	}

	if !equalPointers(s.PoolPurgeDelay, t.PoolPurgeDelay) {
		return false
	}

	if s.Proto != t.Proto {
		return false
	}

	if !equalComparableSlice(s.ProxyV2Options, t.ProxyV2Options, opt) {
		return false
	}

	if s.Redir != t.Redir {
		return false
	}

	if s.ResolveNet != t.ResolveNet {
		return false
	}

	if s.ResolvePrefer != t.ResolvePrefer {
		return false
	}

	if s.ResolveOpts != t.ResolveOpts {
		return false
	}

	if s.Resolvers != t.Resolvers {
		return false
	}

	if !equalPointers(s.Rise, t.Rise) {
		return false
	}

	if s.SendProxy != t.SendProxy {
		return false
	}

	if s.SendProxyV2 != t.SendProxyV2 {
		return false
	}

	if s.SendProxyV2Ssl != t.SendProxyV2Ssl {
		return false
	}

	if s.SendProxyV2SslCn != t.SendProxyV2SslCn {
		return false
	}

	if s.Shard != t.Shard {
		return false
	}

	if s.Sigalgs != t.Sigalgs {
		return false
	}

	if !equalPointers(s.Slowstart, t.Slowstart) {
		return false
	}

	if s.Sni != t.Sni {
		return false
	}

	if s.Socks4 != t.Socks4 {
		return false
	}

	if s.Source != t.Source {
		return false
	}

	if s.Ssl != t.Ssl {
		return false
	}

	if s.SslCafile != t.SslCafile {
		return false
	}

	if s.SslCertificate != t.SslCertificate {
		return false
	}

	if s.SslMaxVer != t.SslMaxVer {
		return false
	}

	if s.SslMinVer != t.SslMinVer {
		return false
	}

	if s.SslReuse != t.SslReuse {
		return false
	}

	if s.Stick != t.Stick {
		return false
	}

	if !equalPointers(s.TCPUt, t.TCPUt) {
		return false
	}

	if s.Tfo != t.Tfo {
		return false
	}

	if s.TLSTickets != t.TLSTickets {
		return false
	}

	if s.Track != t.Track {
		return false
	}

	if s.Verify != t.Verify {
		return false
	}

	if s.Verifyhost != t.Verifyhost {
		return false
	}

	if !equalPointers(s.Weight, t.Weight) {
		return false
	}

	if s.Ws != t.Ws {
		return false
	}

	return true
}

// Diff checks if two structs of type ServerParams are equal
//
// By default empty maps and slices are equal to nil:
//
//	var a, b ServerParams
//	diff := a.Diff(b)
//
// For more advanced use case you can configure these options (default values are shown):
//
//	var a, b ServerParams
//	diff := a.Diff(b,Options{
//		NilSameAsEmpty: true,
//	})
func (s ServerParams) Diff(t ServerParams, opts ...Options) map[string][]interface{} {
	opt := getOptions(opts...)

	diff := make(map[string][]interface{})
	if s.AgentAddr != t.AgentAddr {
		diff["AgentAddr"] = []interface{}{s.AgentAddr, t.AgentAddr}
	}

	if s.AgentCheck != t.AgentCheck {
		diff["AgentCheck"] = []interface{}{s.AgentCheck, t.AgentCheck}
	}

	if !equalPointers(s.AgentInter, t.AgentInter) {
		diff["AgentInter"] = []interface{}{ValueOrNil(s.AgentInter), ValueOrNil(t.AgentInter)}
	}

	if !equalPointers(s.AgentPort, t.AgentPort) {
		diff["AgentPort"] = []interface{}{ValueOrNil(s.AgentPort), ValueOrNil(t.AgentPort)}
	}

	if s.AgentSend != t.AgentSend {
		diff["AgentSend"] = []interface{}{s.AgentSend, t.AgentSend}
	}

	if s.Allow0rtt != t.Allow0rtt {
		diff["Allow0rtt"] = []interface{}{s.Allow0rtt, t.Allow0rtt}
	}

	if s.Alpn != t.Alpn {
		diff["Alpn"] = []interface{}{s.Alpn, t.Alpn}
	}

	if s.Backup != t.Backup {
		diff["Backup"] = []interface{}{s.Backup, t.Backup}
	}

	if s.Check != t.Check {
		diff["Check"] = []interface{}{s.Check, t.Check}
	}

	if s.CheckSendProxy != t.CheckSendProxy {
		diff["CheckSendProxy"] = []interface{}{s.CheckSendProxy, t.CheckSendProxy}
	}

	if s.CheckSni != t.CheckSni {
		diff["CheckSni"] = []interface{}{s.CheckSni, t.CheckSni}
	}

	if s.CheckSsl != t.CheckSsl {
		diff["CheckSsl"] = []interface{}{s.CheckSsl, t.CheckSsl}
	}

	if s.CheckAlpn != t.CheckAlpn {
		diff["CheckAlpn"] = []interface{}{s.CheckAlpn, t.CheckAlpn}
	}

	if s.CheckProto != t.CheckProto {
		diff["CheckProto"] = []interface{}{s.CheckProto, t.CheckProto}
	}

	if s.CheckViaSocks4 != t.CheckViaSocks4 {
		diff["CheckViaSocks4"] = []interface{}{s.CheckViaSocks4, t.CheckViaSocks4}
	}

	if s.Ciphers != t.Ciphers {
		diff["Ciphers"] = []interface{}{s.Ciphers, t.Ciphers}
	}

	if s.Ciphersuites != t.Ciphersuites {
		diff["Ciphersuites"] = []interface{}{s.Ciphersuites, t.Ciphersuites}
	}

	if s.ClientSigalgs != t.ClientSigalgs {
		diff["ClientSigalgs"] = []interface{}{s.ClientSigalgs, t.ClientSigalgs}
	}

	if s.Cookie != t.Cookie {
		diff["Cookie"] = []interface{}{s.Cookie, t.Cookie}
	}

	if s.CrlFile != t.CrlFile {
		diff["CrlFile"] = []interface{}{s.CrlFile, t.CrlFile}
	}

	if s.Curves != t.Curves {
		diff["Curves"] = []interface{}{s.Curves, t.Curves}
	}

	if !equalPointers(s.Downinter, t.Downinter) {
		diff["Downinter"] = []interface{}{ValueOrNil(s.Downinter), ValueOrNil(t.Downinter)}
	}

	if s.ErrorLimit != t.ErrorLimit {
		diff["ErrorLimit"] = []interface{}{s.ErrorLimit, t.ErrorLimit}
	}

	if !equalPointers(s.Fall, t.Fall) {
		diff["Fall"] = []interface{}{ValueOrNil(s.Fall), ValueOrNil(t.Fall)}
	}

	if !equalPointers(s.Fastinter, t.Fastinter) {
		diff["Fastinter"] = []interface{}{ValueOrNil(s.Fastinter), ValueOrNil(t.Fastinter)}
	}

	if s.ForceSslv3 != t.ForceSslv3 {
		diff["ForceSslv3"] = []interface{}{s.ForceSslv3, t.ForceSslv3}
	}

	if s.ForceTlsv10 != t.ForceTlsv10 {
		diff["ForceTlsv10"] = []interface{}{s.ForceTlsv10, t.ForceTlsv10}
	}

	if s.ForceTlsv11 != t.ForceTlsv11 {
		diff["ForceTlsv11"] = []interface{}{s.ForceTlsv11, t.ForceTlsv11}
	}

	if s.ForceTlsv12 != t.ForceTlsv12 {
		diff["ForceTlsv12"] = []interface{}{s.ForceTlsv12, t.ForceTlsv12}
	}

	if s.ForceTlsv13 != t.ForceTlsv13 {
		diff["ForceTlsv13"] = []interface{}{s.ForceTlsv13, t.ForceTlsv13}
	}

	if s.HealthCheckAddress != t.HealthCheckAddress {
		diff["HealthCheckAddress"] = []interface{}{s.HealthCheckAddress, t.HealthCheckAddress}
	}

	if !equalPointers(s.HealthCheckPort, t.HealthCheckPort) {
		diff["HealthCheckPort"] = []interface{}{ValueOrNil(s.HealthCheckPort), ValueOrNil(t.HealthCheckPort)}
	}

	if !equalPointers(s.InitAddr, t.InitAddr) {
		diff["InitAddr"] = []interface{}{ValueOrNil(s.InitAddr), ValueOrNil(t.InitAddr)}
	}

	if !equalPointers(s.Inter, t.Inter) {
		diff["Inter"] = []interface{}{ValueOrNil(s.Inter), ValueOrNil(t.Inter)}
	}

	if s.LogProto != t.LogProto {
		diff["LogProto"] = []interface{}{s.LogProto, t.LogProto}
	}

	if s.Maintenance != t.Maintenance {
		diff["Maintenance"] = []interface{}{s.Maintenance, t.Maintenance}
	}

	if !equalPointers(s.MaxReuse, t.MaxReuse) {
		diff["MaxReuse"] = []interface{}{ValueOrNil(s.MaxReuse), ValueOrNil(t.MaxReuse)}
	}

	if !equalPointers(s.Maxconn, t.Maxconn) {
		diff["Maxconn"] = []interface{}{ValueOrNil(s.Maxconn), ValueOrNil(t.Maxconn)}
	}

	if !equalPointers(s.Maxqueue, t.Maxqueue) {
		diff["Maxqueue"] = []interface{}{ValueOrNil(s.Maxqueue), ValueOrNil(t.Maxqueue)}
	}

	if !equalPointers(s.Minconn, t.Minconn) {
		diff["Minconn"] = []interface{}{ValueOrNil(s.Minconn), ValueOrNil(t.Minconn)}
	}

	if s.Namespace != t.Namespace {
		diff["Namespace"] = []interface{}{s.Namespace, t.Namespace}
	}

	if s.NoSslv3 != t.NoSslv3 {
		diff["NoSslv3"] = []interface{}{s.NoSslv3, t.NoSslv3}
	}

	if s.NoTlsv10 != t.NoTlsv10 {
		diff["NoTlsv10"] = []interface{}{s.NoTlsv10, t.NoTlsv10}
	}

	if s.NoTlsv11 != t.NoTlsv11 {
		diff["NoTlsv11"] = []interface{}{s.NoTlsv11, t.NoTlsv11}
	}

	if s.NoTlsv12 != t.NoTlsv12 {
		diff["NoTlsv12"] = []interface{}{s.NoTlsv12, t.NoTlsv12}
	}

	if s.NoTlsv13 != t.NoTlsv13 {
		diff["NoTlsv13"] = []interface{}{s.NoTlsv13, t.NoTlsv13}
	}

	if s.NoVerifyhost != t.NoVerifyhost {
		diff["NoVerifyhost"] = []interface{}{s.NoVerifyhost, t.NoVerifyhost}
	}

	if s.Npn != t.Npn {
		diff["Npn"] = []interface{}{s.Npn, t.Npn}
	}

	if s.Observe != t.Observe {
		diff["Observe"] = []interface{}{s.Observe, t.Observe}
	}

	if s.OnError != t.OnError {
		diff["OnError"] = []interface{}{s.OnError, t.OnError}
	}

	if s.OnMarkedDown != t.OnMarkedDown {
		diff["OnMarkedDown"] = []interface{}{s.OnMarkedDown, t.OnMarkedDown}
	}

	if s.OnMarkedUp != t.OnMarkedUp {
		diff["OnMarkedUp"] = []interface{}{s.OnMarkedUp, t.OnMarkedUp}
	}

	if !equalPointers(s.PoolLowConn, t.PoolLowConn) {
		diff["PoolLowConn"] = []interface{}{ValueOrNil(s.PoolLowConn), ValueOrNil(t.PoolLowConn)}
	}

	if !equalPointers(s.PoolMaxConn, t.PoolMaxConn) {
		diff["PoolMaxConn"] = []interface{}{ValueOrNil(s.PoolMaxConn), ValueOrNil(t.PoolMaxConn)}
	}

	if !equalPointers(s.PoolPurgeDelay, t.PoolPurgeDelay) {
		diff["PoolPurgeDelay"] = []interface{}{ValueOrNil(s.PoolPurgeDelay), ValueOrNil(t.PoolPurgeDelay)}
	}

	if s.Proto != t.Proto {
		diff["Proto"] = []interface{}{s.Proto, t.Proto}
	}

	if !equalComparableSlice(s.ProxyV2Options, t.ProxyV2Options, opt) {
		diff["ProxyV2Options"] = []interface{}{s.ProxyV2Options, t.ProxyV2Options}
	}

	if s.Redir != t.Redir {
		diff["Redir"] = []interface{}{s.Redir, t.Redir}
	}

	if s.ResolveNet != t.ResolveNet {
		diff["ResolveNet"] = []interface{}{s.ResolveNet, t.ResolveNet}
	}

	if s.ResolvePrefer != t.ResolvePrefer {
		diff["ResolvePrefer"] = []interface{}{s.ResolvePrefer, t.ResolvePrefer}
	}

	if s.ResolveOpts != t.ResolveOpts {
		diff["ResolveOpts"] = []interface{}{s.ResolveOpts, t.ResolveOpts}
	}

	if s.Resolvers != t.Resolvers {
		diff["Resolvers"] = []interface{}{s.Resolvers, t.Resolvers}
	}

	if !equalPointers(s.Rise, t.Rise) {
		diff["Rise"] = []interface{}{ValueOrNil(s.Rise), ValueOrNil(t.Rise)}
	}

	if s.SendProxy != t.SendProxy {
		diff["SendProxy"] = []interface{}{s.SendProxy, t.SendProxy}
	}

	if s.SendProxyV2 != t.SendProxyV2 {
		diff["SendProxyV2"] = []interface{}{s.SendProxyV2, t.SendProxyV2}
	}

	if s.SendProxyV2Ssl != t.SendProxyV2Ssl {
		diff["SendProxyV2Ssl"] = []interface{}{s.SendProxyV2Ssl, t.SendProxyV2Ssl}
	}

	if s.SendProxyV2SslCn != t.SendProxyV2SslCn {
		diff["SendProxyV2SslCn"] = []interface{}{s.SendProxyV2SslCn, t.SendProxyV2SslCn}
	}

	if s.Shard != t.Shard {
		diff["Shard"] = []interface{}{s.Shard, t.Shard}
	}

	if s.Sigalgs != t.Sigalgs {
		diff["Sigalgs"] = []interface{}{s.Sigalgs, t.Sigalgs}
	}

	if !equalPointers(s.Slowstart, t.Slowstart) {
		diff["Slowstart"] = []interface{}{ValueOrNil(s.Slowstart), ValueOrNil(t.Slowstart)}
	}

	if s.Sni != t.Sni {
		diff["Sni"] = []interface{}{s.Sni, t.Sni}
	}

	if s.Socks4 != t.Socks4 {
		diff["Socks4"] = []interface{}{s.Socks4, t.Socks4}
	}

	if s.Source != t.Source {
		diff["Source"] = []interface{}{s.Source, t.Source}
	}

	if s.Ssl != t.Ssl {
		diff["Ssl"] = []interface{}{s.Ssl, t.Ssl}
	}

	if s.SslCafile != t.SslCafile {
		diff["SslCafile"] = []interface{}{s.SslCafile, t.SslCafile}
	}

	if s.SslCertificate != t.SslCertificate {
		diff["SslCertificate"] = []interface{}{s.SslCertificate, t.SslCertificate}
	}

	if s.SslMaxVer != t.SslMaxVer {
		diff["SslMaxVer"] = []interface{}{s.SslMaxVer, t.SslMaxVer}
	}

	if s.SslMinVer != t.SslMinVer {
		diff["SslMinVer"] = []interface{}{s.SslMinVer, t.SslMinVer}
	}

	if s.SslReuse != t.SslReuse {
		diff["SslReuse"] = []interface{}{s.SslReuse, t.SslReuse}
	}

	if s.Stick != t.Stick {
		diff["Stick"] = []interface{}{s.Stick, t.Stick}
	}

	if !equalPointers(s.TCPUt, t.TCPUt) {
		diff["TCPUt"] = []interface{}{ValueOrNil(s.TCPUt), ValueOrNil(t.TCPUt)}
	}

	if s.Tfo != t.Tfo {
		diff["Tfo"] = []interface{}{s.Tfo, t.Tfo}
	}

	if s.TLSTickets != t.TLSTickets {
		diff["TLSTickets"] = []interface{}{s.TLSTickets, t.TLSTickets}
	}

	if s.Track != t.Track {
		diff["Track"] = []interface{}{s.Track, t.Track}
	}

	if s.Verify != t.Verify {
		diff["Verify"] = []interface{}{s.Verify, t.Verify}
	}

	if s.Verifyhost != t.Verifyhost {
		diff["Verifyhost"] = []interface{}{s.Verifyhost, t.Verifyhost}
	}

	if !equalPointers(s.Weight, t.Weight) {
		diff["Weight"] = []interface{}{ValueOrNil(s.Weight), ValueOrNil(t.Weight)}
	}

	if s.Ws != t.Ws {
		diff["Ws"] = []interface{}{s.Ws, t.Ws}
	}

	return diff
}
