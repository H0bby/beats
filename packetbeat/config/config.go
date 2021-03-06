package config

import (
	"github.com/elastic/beats/libbeat/common/droppriv"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	"github.com/elastic/beats/packetbeat/procs"
)

type Config struct {
	Interfaces InterfacesConfig
	Flows      *Flows
	Protocols  Protocols
	Shipper    publisher.ShipperConfig
	Procs      procs.ProcsConfig
	RunOptions droppriv.RunOptions
	Logging    logp.Logging
	Filter     map[string]interface{}
}

type InterfacesConfig struct {
	Device         string
	Type           string
	File           string
	With_vlans     bool
	Bpf_filter     string
	Snaplen        int
	Buffer_size_mb int
	TopSpeed       bool
	Dumpfile       string
	OneAtATime     bool
	Loop           int
}

type Flows struct {
	Timeout string
	Period  string
}

type Protocols struct {
	Icmp     Icmp
	Amqp     Amqp
	Dns      Dns
	Http     Http
	Memcache Memcache
	Mysql    Mysql
	Mongodb  Mongodb
	Pgsql    Pgsql
	Redis    Redis
	Thrift   Thrift
}

type ProtocolCommon struct {
	Ports              []int `config:"ports"`
	SendRequest        *bool `config:"send_request"`
	SendResponse       *bool `config:"send_response"`
	TransactionTimeout *int  `config:"transaction_timeout"`
}

type Icmp struct {
	Enabled            bool
	SendRequest        *bool `config:"send_request"`
	SendResponse       *bool `config:"send_response"`
	TransactionTimeout *int  `config:"transaction_timeout"`
}

type Amqp struct {
	ProtocolCommon            `config:",inline"`
	ParseHeaders              *bool `config:"parse_headers"`
	ParseArguments            *bool `config:"parse_arguments"`
	MaxBodyLength             *int  `config:"max_body_length"`
	HideConnectionInformation *bool `config:"hide_connection_information"`
}

type Dns struct {
	ProtocolCommon      `config:",inline"`
	Include_authorities *bool
	Include_additionals *bool
}

type Http struct {
	ProtocolCommon       `config:",inline"`
	Send_all_headers     *bool
	Send_headers         []string
	Split_cookie         *bool
	Real_ip_header       *string
	Include_body_for     []string
	Hide_keywords        []string
	Redact_authorization *bool
}

type Memcache struct {
	ProtocolCommon        `config:",inline"`
	MaxValues             int
	MaxBytesPerValue      int
	UdpTransactionTimeout *int
	ParseUnknown          bool
}

type Mysql struct {
	ProtocolCommon `config:",inline"`
	Max_row_length *int
	Max_rows       *int
}

type Mongodb struct {
	ProtocolCommon `config:",inline"`
	Max_doc_length *int
	Max_docs       *int
}

type Pgsql struct {
	ProtocolCommon `config:",inline"`
	Max_row_length *int
	Max_rows       *int
}

type Thrift struct {
	ProtocolCommon             `config:",inline"`
	String_max_size            *int
	Collection_max_size        *int
	Drop_after_n_struct_fields *int
	Transport_type             *string
	Protocol_type              *string
	Capture_reply              *bool
	Obfuscate_strings          *bool
	Idl_files                  []string
}

type Redis struct {
	ProtocolCommon `config:",inline"`
}

// Config Singleton
var ConfigSingleton Config
