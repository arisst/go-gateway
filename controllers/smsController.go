package controllers

import (
	"go-gateway/models"
	u "go-gateway/utils"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/BurntSushi/toml"
)

func getVendor(config_file string) *models.Vendor {
	var vendors models.Vendors
	if _, err := toml.DecodeFile(config_file, &vendors); err != nil {
		log.Fatal(err)
	}

	enabledVendor := []models.Vendor{}

	for _, s := range vendors.Vendor {
		if s.Enabled {
			enabledVendor = append(enabledVendor, s)
		}
	}

	if len(enabledVendor) == 0 {
		return nil
	}

	return models.GetRandom(enabledVendor)
}

func SendSMS(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	phone := url.QueryEscape(query.Get("phone"))
	text := url.QueryEscape(query.Get("text"))

	vendor := getVendor("vendor.toml")

	if vendor == nil {
		u.Respond(w, u.Message(false, "Vendor not available!"))
		return
	}

	replacer := strings.NewReplacer("{phone}", phone, "{text}", text)

	vendor.Url = replacer.Replace(vendor.Url)

	result := u.Request(vendor.Url)

	resp := u.Message(true, "Success")

	resp["payload"] = map[string]interface{}{"phone": phone, "text": text}
	resp["vendor"] = vendor
	resp["result"] = result

	u.Respond(w, resp)
}
