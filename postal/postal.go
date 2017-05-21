package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type webConf struct {
	Host     string `yaml:"host,omitempty"`
	Protocol string `yaml:"protocol,omitempty"`
}
type fastServerConf struct {
	Enabled bool `yaml:"enabled,omitempty"`
}

type generalConf struct {
	UseIPPools bool `yaml:"use_ip_pools,omitempty"`
}

type mainDbConf struct {
	Host     string `yaml:"host,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	Database string `yaml:"database,omitempty"`
}

type messageDbConf struct {
	Host     string `yaml:"host,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	Prefix   string `yaml:"prefix,omitempty"`
}

type rabbitmqConf struct {
	Host     string `yaml:"host,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	Vhost    string `yaml:"vhost,omitempty"`
}

type dnsConf struct {
	MxRecords           []string `yaml:"mx_records,omitempty"`
	SMTPpServerHostName string   `yaml:"smtp_server_hostname,omitempty"`
	SpfInclude          string   `yaml:"spf_include,omitempty"`
	ReturnPath          string   `yaml:"return_path,omitempty"`
	RouteDomain         string   `yaml:"route_domain,omitempty"`
	TrackDomain         string   `yaml:"track_domain,omitempty"`
}

type smtpConf struct {
	Host        string `yaml:"host,omitempty"`
	Port        string `yaml:"port,omitempty"`
	Username    string `yaml:"username,omitempty"`
	Password    string `yaml:"password,omitempty"`
	FromName    string `yaml:"from_name,omitempty"`
	FromAddress string `yaml:"from_address,omitempty"`
}

type railsConf struct {
	SecretKey string `yaml:"secret_key,omitempty"`
}

type postalConfg struct {
	Web            webConf        `yaml:"web"`
	FastServerConf fastServerConf `yaml:"fast_server"`
	GeneralConf    generalConf    `yaml:"general"`
	MainDbConf     mainDbConf     `yaml:"main_db"`
	MessageDbConf  messageDbConf  `yaml:"message_db"`
	RabbitmqConf   rabbitmqConf   `yaml:"rabbitmq"`
	DNSConf        dnsConf        `yaml:"dns"`
	SMTPConf       smtpConf       `yaml:"smtp"`
}

func main() {

	if err != nil {
		log.Fatalf("error: %v", err)
		panic("Die!")
	}
}

func newSMTPConf(String t, path string) interface{} {

	if t == "" {

	}
	f, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("error: %v", err)
		panic("Unable to open file.")
	}

	err = yaml.Unmarshal(f, &conf)

	if err != nil {
		log.Fatalf("error: %v", err)
		panic("Unable to unmarshall file.")
	}

	return conf
}
