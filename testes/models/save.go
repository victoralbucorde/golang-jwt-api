package models

type Save struct {
	Roach         int64 `json:"roach"`
	AllTimeRoachs int64 `json:"all_time_roachs"`
	TimesClicked  int64 `json:"times_clicked"`
	Upgrade1      bool  `json:"upgrade_1"`
	Upgrade2      bool  `json:"upgrade_2"`
	Upgrade3      bool  `json:"upgrade_3"`
	Upgrade4      bool  `json:"upgrade_4"`
	Upgrade5      bool  `json:"upgrade_5"`
	Upgrade6      bool  `json:"upgrade_6"`
	Achievement1  bool  `json:"achievement_1"`
	Achievement2  bool  `json:"achievement_2"`
	Achievement3  bool  `json:"achievement_3"`
	Achievement4  bool  `json:"achievement_4"`
	Achievement5  bool  `json:"achievement_5"`
	Achievement6  bool  `json:"achievement_6"`
}

func NewSave() *Save {
	return &Save{Roach: 0, AllTimeRoachs: 0, Achievement1: false, Achievement2: false, Achievement3: false, Achievement4: false, Achievement5: false, Achievement6: false, Upgrade1: false, Upgrade2: false, Upgrade3: false, Upgrade4: false, Upgrade5: false, Upgrade6: false}
}
