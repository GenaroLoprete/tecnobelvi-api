package info

type Build struct {
	Name       string   `json:"Name"`
	Components []string `json:"Components"`
	Price      float32  `json:"Price"`
}

func GetBuildsInfo() []Build {
	return []Build{
		{
			Name: "Pc Intel Gamer",
			Components: []string{
				"I7 10700F",
				"RTX 3080",
				"64GB RAM DDR4",
			},
			Price: 1400,
		},
		{
			Name: "Pc Amd Gamer",
			Components: []string{
				"Ryzen 9 5950x",
				"RTX 3090TI",
				"64GB RAM DDR4",
			},
			Price: 1350,
		},
		{
			Name: "Pc Qualcomm Gamer",
			Components: []string{
				"Snapdragon 898",
				"RTX 3060",
				"64GB RAM DDR4",
			},
			Price: 1100,
		},
	}
}
