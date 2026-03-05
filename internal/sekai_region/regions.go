package sekairegion

import "time"

type SekaiRegion struct {
	id       string
	name     string
	location *time.Location
	enable   bool
}

var Regions []*SekaiRegion

// TODO: 这里要改成从config中获取
func init() {
	Regions = []*SekaiRegion{
		{
			"cn",
			"国服",
			loadLocation("Asia/Shanghai"),
			true,
		}, {
			"jp",
			"日服",
			loadLocation("Asia/Tokyo"),
			true,
		}, {
			"tw",
			"台服",
			loadLocation("Asia/Taipei"),
			true,
		}, {
			"kr",
			"韩服",
			loadLocation("Asia/Seoul"),
			true,
		}, {
			"en",
			"国际服",
			loadLocation("America/Los_Angeles"),
			true,
		},
	}
}

func GetRegionById(id string) *SekaiRegion {
	for _, r := range Regions {
		if r.id == id {
			return r
		}
	}
	return nil
}

func loadLocation(name string) *time.Location {
	location, err := time.LoadLocation(name)
	if err != nil {
		location = time.Local
	}
	return location
}

func (r *SekaiRegion) Id() string {
	return r.id
}

func (r *SekaiRegion) Name() string {
	return r.name
}

func (r *SekaiRegion) Location() *time.Location {
	return r.location
}

func (r *SekaiRegion) Enabled() bool {
	return r.enable
}

// const (
// 	CN = SekaiRegion("cn")
// 	JP = SekaiRegion("jp")
// 	TW = SekaiRegion("tw")
// 	KR = SekaiRegion("kr")
// 	EN = SekaiRegion("en")
// )
