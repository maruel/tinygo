// +build !wasm

package time

type zone struct {
	name   string // abbreviated name, "CET"
	offset int    // seconds east of UTC
	//isDST  bool   // is this zone Daylight Savings Time?
}

type Location struct {
	cacheStart int64
	cacheEnd   int64
	cacheZone  *zone
}

var utcLoc Location
var localLoc Location

var UTC *Location = &utcLoc
var Local *Location = &localLoc

func (l *Location) get() *Location {
	return &utcLoc
}

func (l *Location) lookup(sec int64) (name string, offset int, start, end int64) {
	name = "UTC"
	return
}

// FixedZone returns a Location that always uses
// the given zone name and offset (seconds east of UTC).
func FixedZone(name string, offset int) *Location {
	return &Location{}
}

func (l *Location) lookupName(name string, unix int64) (offset int, ok bool) {
	return 0, true
}
