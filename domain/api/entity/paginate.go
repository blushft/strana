package entity

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber"
)

type Paginate struct {
	Offset int
	Limit  int
	Start  time.Time
	End    time.Time
}

func defaultParams() *Paginate {
	return &Paginate{
		Limit:  20,
		Offset: 0,
		Start:  time.Now().AddDate(0, 0, -7),
		End:    time.Now(),
	}
}

func apiParams(c *fiber.Ctx) {
	params := defaultParams()

	limit, err := strconv.Atoi(c.Query("limit"))
	if err == nil {
		params.Limit = limit
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err == nil {
		params.Offset = offset
	}

	start, err := strconv.ParseInt(c.Query("start"), 10, 64)
	if err == nil {
		params.Start = time.Unix(start, 0)
	}

	end, err := strconv.ParseInt(c.Query("end"), 10, 64)
	if err == nil {
		params.End = time.Unix(end, 0)
	}

	c.Locals("params", params)

	c.Next()
}
