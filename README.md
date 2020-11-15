# Installation
- Clone this repository
- Adjust config file with your environment in `.env` file
- Run `go run main.go`
- Test `go test ./... -v`

# Or with Docker Compose
- Run `docker-compose up -d`
- Test `docker-compose exec app go test ./... -v`
- Goto `http://localhost:8001/api/sendsms`

# API Documentation
- Adjust vendor config available at `vendor.toml`
  - `name` : Vendor Name
  - `url` : Vendor endpoint url
  - `enabled` : (true or false) Enable or disable vendor
- You can enable more than one of third-party service, and will be randomized
- [GET] `/api/sendsms?phone=+62845323&text=test+kirim+sms`
  - Response 
```json
{
	"message": "Success",
	"payload": {
		"phone": "+62845323",
		"text": "test+kirim+sms"
	},
	"result": {
		"info": "vendor 2",
		"message": "SMS was delivered to this phone number!",
		"status": true
	},
	"status": true,
	"vendor": {
		"name": "Vendor 2",
		"url": "https://run.mocky.io/v3/28a891af-d94a-431c-a26f-1e68d0b12dc3?username=vendor2user&password=vendor2pass&hp=+845323&txt=test+kirim+sms",
		"enabled": true
	}
}
```