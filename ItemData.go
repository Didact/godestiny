package destiny

type ItemData struct {
	ItemHash            float64
	ItemName            string
	ItemDescription     string
	Icon                string
	SecondaryIcon       string
	DisplaySource       string
	ActionName          string
	HasAction           bool
	DeleteOnAction      bool
	TierTypeName        string
	TierType            float64
	ItemTypeName        string
	BucketTypehash      float64
	PrimaryBaseStatHash float64
	Stats               interface{}
	PerkHashes          interface{}
	SpecialItemType     float64
	TalentGridHash      float64
	HasGeometry         bool
	StatGroupHash       float64
	ItemLevels          []interface{}
	QualityLevel        float64
	Equippable          bool
	Instanced           bool
	RewardItemhash      float64
	Values              interface{}
	ItemType            float64
	ItemSubType         float64
	ClassType           float64
	Sources             []struct {
		ExpansionIndex   float64
		Level            float64
		MinQuality       float64
		MaxQuality       float64
		MinLevelRequired float64
		MaxLevelRequired float64
		Exclusivity      float64
		ComputedStats    interface{}
		SourceHashes     []float64
		SpawnIndexes     []float64
	}
	ItemCategoryHashes  []interface{}
	SourceHashes        []interface{}
	NonTransferrable    bool
	Exclusive           float64
	MaxStackSize        float64
	ItemIndex           float64
	SetItemHashes       []interface{}
	QuestlineItemhash   float64
	NeedsFullCompletion bool
	ObjectiveHashes     []interface{}
}
