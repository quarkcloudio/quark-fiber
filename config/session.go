package config

import (
	"time"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/memory"
)

var Session = map[string]interface{}{

	// Allowed session duration
	// Optional. Default value 24 * time.Hour
	"expiration": 24 * time.Hour,

	// Storage interface to store the session data
	// Optional. Default value memory.New()
	"storage": memory.New(),

	// KeyLookup is a string in the form of "<source>:<name>" that is used
	// to extract session id from the request.
	// Possible values: "header:<name>", "query:<name>" or "cookie:<name>"
	// Optional. Default value "cookie:session_id".
	"key_lookup": "cookie:session_id",

	// Domain of the cookie.
	// Optional. Default value "".
	"cookie_domain": "",

	// Path of the cookie.
	// Optional. Default value "".
	"cookie_path": "",

	// Indicates if cookie is secure.
	// Optional. Default value false.
	"cookie_secure": false,

	// Indicates if cookie is HTTP only.
	// Optional. Default value false.
	"cookie_http_only": false,

	// Sets the cookie SameSite attribute.
	// Optional. Default value "Lax".
	"cookie_same_site": "Lax",

	// KeyGenerator generates the session key.
	// Optional. Default value utils.UUIDv4
	"key_generator": utils.UUIDv4,
}
