// Reads a simdb.bin through the engine's own proto and item loader, and prints
// what it found as JSON. The Python pipeline writes that file; this is the
// proof that the Go engine reads back exactly what it wrote.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	goproto "google.golang.org/protobuf/proto"
)

type report struct {
	Bytes       int                `json:"bytes"`
	Items       int                `json:"items"`
	Enchants    int                `json:"enchants"`
	Weapons     int                `json:"weapons"`
	InASet      int                `json:"in_a_set"`
	WithStats   int                `json:"with_stats"`
	EnchantHits int                `json:"enchants_with_stats"`
	Spot        map[string]float64 `json:"spot"`
}

func main() {
	blob, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	db := &proto.SimDatabase{}
	if err := goproto.Unmarshal(blob, db); err != nil {
		panic(err)
	}
	out := report{Bytes: len(blob), Items: len(db.Items), Enchants: len(db.Enchants),
		Spot: map[string]float64{}}
	for _, row := range db.Items {
		item := core.ItemFromProto(row)
		if item.SwingSpeed > 0 {
			out.Weapons++
		}
		if item.SetID != 0 {
			out.InASet++
		}
		if len(row.Stats) > 0 {
			out.WithStats++
		}
		switch item.ID {
		case 12784: // Arcanite Reaper: curve-resolved stats plus curve-resolved damage
			out.Spot["reaper_attack_power"] = item.Stats[stats.AttackPower]
			out.Spot["reaper_stamina"] = item.Stats[stats.Stamina]
			out.Spot["reaper_speed"] = item.SwingSpeed
			out.Spot["reaper_damage_min"] = item.WeaponDamageMin
			out.Spot["reaper_damage_max"] = item.WeaponDamageMax
		case 22416: // Dreadnaught Breastplate: plate, armour, set name
			out.Spot["dreadnaught_armor"] = item.Stats[stats.Armor]
			out.Spot["dreadnaught_is_plate"] = boolAsFloat(
				item.ArmorType == proto.ArmorType_ArmorTypePlate)
			out.Spot["dreadnaught_in_a_set"] = boolAsFloat(
				item.SetName == "Dreadnaught's Battlegear")
		case 19120: // Rune of the Guard Captain: +42/+42 from two on-equip spells
			out.Spot["rune_attack_power"] = item.Stats[stats.AttackPower]
			out.Spot["rune_ranged_attack_power"] = item.Stats[stats.RangedAttackPower]
			out.Spot["rune_hit"] = item.Stats[stats.Hit]
		}
	}
	for _, row := range db.Enchants {
		if len(row.Stats) > 0 {
			out.EnchantHits++
		}
	}
	// The path the lanes actually use: the database rides on the player.
	player := &proto.Player{Database: db}
	out.Spot["player_database_items"] = float64(len(player.Database.Items))
	encoded, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encoded))
}

func boolAsFloat(value bool) float64 {
	if value {
		return 1
	}
	return 0
}
