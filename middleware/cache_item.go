package middleware

import "time"

// A cache item storing any data and a TTL (Time-to-Live) duration.
type CacheItem struct {
	data any
	ttl  time.Time
}

// Constructor.
func NewCacheItem(data any, ttl time.Time) CacheItem {
	return CacheItem{
		data: data,
		ttl:  ttl,
	}
}

// Returns true if the item has Expired.
func (i CacheItem) Expired() bool {
	return time.Now().After(i.ttl)
}
