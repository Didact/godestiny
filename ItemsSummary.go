package destiny

import (
	"encoding/json"
	"log"
	"time"
)

type ItemsSummary struct {
	Response struct {
		Data struct {
			Items []struct {
				ItemHash    float64
				ItemId      string
				Quantity    float64
				DamageType  float64
				PrimaryStat struct {
					StatHash     float64
					Value        float64
					MaximumValue float64
				}
				IsGridComplete bool
				TransferStatus float64
				State          float64
				CharacterIndex float64
				BucketHash     float64
			}
			Characters []struct {
				CharacterBase struct {
					MembershipId             string
					MembershipType           float64
					CharacterId              string
					DateLastPlayed           time.Time
					MinutesPlayedThisSession string
					MinutesPlayedTotal       string
					PowerLevel               float64
					RaceHash                 float64
					GenderHash               float64
					ClassHash                float64
					CurrentActivityHash      float64
					LastCompletedStoryHash   float64
					Stats                    struct {
						Stat_Defense           StatData
						Stat_Intellect         StatData
						Stat_Discipline        StatData
						Stat_Strength          StatData
						Stat_Light             StatData
						Stat_Armor             StatData
						Stat_Agility           StatData
						Stat_Recovery          StatData
						Stat_Optics            StatData
						Stat_Attack_Speed      StatData
						Stat_Damage_Reduction  StatData
						Stat_Attack_Efficiency StatData
						Stat_Attack_Energy     StatData
					}
					Customization struct {
						Personality  float64
						Face         float64
						SkinColor    float64
						LipColor     float64
						EyeColor     float64
						HairColor    float64
						FeatureColor float64
						DecalColor   float64
						WearHelmet   bool
						HairIndex    float64
						FeatureIndex float64
						DecalIndex   float64
					}
					GrimoireScore float64
					PeerView      struct {
						Equipment []struct {
							ItemHash float64
							Dyes     []struct {
								ChannelHash float64
								dyeHash     float64
							}
						}
					}
					GenderType         float64
					ClassType          float64
					BuildStatGroupHash float64
				}
				LevelProgression struct {
					DailyProgress       float64
					WeeklyProgress      float64
					CurrentProgress     float64
					Level               float64
					Step                float64
					ProgressToNextLevel float64
					NextLevelAt         float64
					Progressionhash     float64
				}
				EmblemPath         string
				BackgroundPath     string
				EmblemHash         float64
				CharacterLevel     float64
				BaseCharacterLevel float64
				IsPrestigeLevel    bool
				PercentToNextLevel float64
			}
		}
	}
	ErrorCode       float64
	ThrottleSeconds float64
	ErrorStatus     string
	Message         string
	MessageData     interface{}
}

type StatData struct {
	StatHash     float64
	Value        float64
	MaximumValue float64
}

func (is *ItemsSummary) GetAllItems() []string {
	var res []string
	for _, item := range is.Response.Data.Items {
		rows, err := db.Query("SELECT json FROM DestinyInventoryItemDefinition where id=?", item.ItemHash)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		for rows.Next() {
			var row []byte
			err = rows.Scan(&row)
			if err != nil {
				log.Fatal(err)
			}
			id := ItemData{}
			err = json.Unmarshal(row, &id)
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Println(id.ItemName)
			res = append(res, id.ItemName)
		}
	}
	return res
}
