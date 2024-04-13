package cache

import (
	"banners_service/internal/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

type allCache struct {
    Data *cache.Cache
}

const (
    defaultExpiration = 5 * time.Minute
    purgeTime         = 5 * time.Minute
)

func newCache() *allCache {
    Cache := cache.New(defaultExpiration, purgeTime)
    return &allCache{
        Data: Cache,
    }
}

func (cach *allCache) Read(c *fiber.Ctx) error {
	useLastRevision := c.QueryBool("use_last_revision")
	if useLastRevision {
		return c.Next()
	}
	tagID := c.QueryInt("tag_id")
	featureID := c.QueryInt("feature_id")

    banner, ok :=  CacheData.Data.Get(strconv.Itoa(tagID) + " " + strconv.Itoa(featureID))
    if ok {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"example": banner.Content,
	})
    }
    return c.Next()
}

func (c *allCache) Update(id string, data models.Banner) {
    c.Data.Set(id, data, cache.DefaultExpiration)
}

var CacheData = newCache()