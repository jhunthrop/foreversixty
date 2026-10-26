package guilds

import (
	"fmt"
	"reflect"
	"testing"
)

func encounter(id int) *int { return &id }

func TestKillOrderOf(t *testing.T) {
	ms := func(v int64) *int64 { return &v }
	cases := []struct {
		name        string
		fights      []reportFight
		want        []KillOrderView
		wantUnnamed *UnnamedView
	}{
		{
			name: "pulls after the first kill are ignored",
			fights: []reportFight{
				{index: 0, encounterID: encounter(1), name: "Boss", durationMS: 100, deaths: 2},
				{index: 1, encounterID: encounter(1), name: "Boss", kill: true, durationMS: 300, deaths: 1},
				{index: 2, encounterID: encounter(1), name: "Boss", kill: true, durationMS: 999, deaths: 9},
				{index: 3, encounterID: encounter(1), name: "Boss", durationMS: 50, deaths: 7},
			},
			want: []KillOrderView{{EncounterID: 1, Name: "Boss", FightIndex: 1, Killed: true, KillMS: ms(300), WipesBefore: 1, Deaths: 1}},
		},
		{
			name: "a never-killed boss: last pull, every pull a wipe, deaths summed",
			fights: []reportFight{
				{index: 0, encounterID: encounter(2), name: "Wall", deaths: 3},
				{index: 1, encounterID: encounter(2), name: "Wall", deaths: 4},
				{index: 2, encounterID: encounter(2), name: "Wall", deaths: 5},
			},
			want: []KillOrderView{{EncounterID: 2, Name: "Wall", FightIndex: 2, Killed: false, KillMS: nil, WipesBefore: 3, Deaths: 12}},
		},
		{
			name: "sorted by the deciding fight; unnamed pulls counted apart",
			fights: []reportFight{
				{index: 0, encounterID: encounter(2), name: "Wall"},
				{index: 1, name: "Trash", kill: true},
				{index: 2, encounterID: encounter(1), name: "Boss", kill: true, durationMS: 10},
				{index: 3, name: "Trash"},
				{index: 4, encounterID: encounter(2), name: "Wall", kill: true, durationMS: 20},
			},
			want: []KillOrderView{
				{EncounterID: 1, Name: "Boss", FightIndex: 2, Killed: true, KillMS: ms(10)},
				{EncounterID: 2, Name: "Wall", FightIndex: 4, Killed: true, KillMS: ms(20), WipesBefore: 1},
			},
			wantUnnamed: &UnnamedView{Pulls: 2, Kills: 1, Wipes: 1},
		},
		{name: "no fights", want: []KillOrderView{}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, unnamed := killOrderOf(c.fights)
			if !reflect.DeepEqual(got, c.want) {
				t.Fatalf("kill order = %+v, want %+v", got, c.want)
			}
			if !reflect.DeepEqual(unnamed, c.wantUnnamed) {
				t.Fatalf("unnamed = %+v, want %+v", unnamed, c.wantUnnamed)
			}
		})
	}
}

func TestDeathsTieBreakIsClassOrderThenName(t *testing.T) {
	rows := []ReportDeathView{
		{Name: "Zed", Class: "mage", OnWipes: 2},
		{Name: "Bob", Class: "warrior", OnWipes: 1, OnKills: 1},
		{Name: "Amy", Class: "mage", OnKills: 2},
		{Name: "Cal", Class: "warrior", OnWipes: 5},
	}
	sortDeaths(rows)
	var got []string
	for _, r := range rows {
		got = append(got, r.Name)
	}
	if want := []string{"Cal", "Bob", "Amy", "Zed"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("order = %v, want %v", got, want)
	}
}

func TestMoveOf(t *testing.T) {
	cases := []struct {
		name     string
		cur, old killStat
		want     bool
		metric   string
		from, to int
	}{
		{"dps change", killStat{dps: f64(150)}, killStat{dps: f64(100)}, true, "dps", 100, 150},
		{"healer change uses hps", killStat{healer: true, hps: f64(80)}, killStat{healer: true, hps: f64(50)}, true, "hps", 50, 80},
		{"no change", killStat{dps: f64(100)}, killStat{dps: f64(100)}, false, "", 0, 0},
		{"a role switch is not a change", killStat{healer: true, hps: f64(80)}, killStat{dps: f64(100)}, false, "", 0, 0},
		{"a missing metric now is not a change", killStat{dps: nil}, killStat{dps: f64(100)}, false, "", 0, 0},
		{"a missing metric before is not a change", killStat{dps: f64(100)}, killStat{dps: nil}, false, "", 0, 0},
		{"a healer without hps is not a change", killStat{healer: true}, killStat{healer: true, hps: f64(50)}, false, "", 0, 0},
		{"a change that rounds away is dropped", killStat{dps: f64(150.3)}, killStat{dps: f64(149.8)}, false, "", 0, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			row, moved := moveOf(c.cur, c.old, "Boss")
			if moved != c.want {
				t.Fatalf("moved = %v, want %v (%+v)", moved, c.want, row)
			}
			if moved && (row.Metric != c.metric || row.Before != c.from || row.After != c.to) {
				t.Fatalf("row = %+v, want %s %d -> %d", row, c.metric, c.from, c.to)
			}
		})
	}
}

func TestTopMovesCapsAtFiveLargestFirst(t *testing.T) {
	var rows []MoveRow
	for i := 1; i <= 8; i++ {
		d := float64(i * 10)
		if i%2 == 0 {
			d = -d // drops count by size, not sign
		}
		rows = append(rows, MoveRow{Name: fmt.Sprintf("P%d", i), delta: d})
	}
	got := topMoves(rows)
	if len(got) != movesCap {
		t.Fatalf("len = %d, want %d", len(got), movesCap)
	}
	for i, want := range []string{"P8", "P7", "P6", "P5", "P4"} {
		if got[i].Name != want {
			t.Fatalf("row %d = %s, want %s (%+v)", i, got[i].Name, want, got)
		}
	}
}
