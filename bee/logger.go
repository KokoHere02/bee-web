package bee

import (
	"log"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Log the request details
		log.Printf("Request: %s %s", c.Method, c.Path)

		c.Next() // Call the next handler in the chain
		// Log the response status code
		log.Printf("Response: %d", c.StateCode)
	}
}
