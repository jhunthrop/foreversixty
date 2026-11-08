package events

// The Activity icons the approved mocks draw (design/mocks/ws_app_flows.py), each a name under
// /data/{build}/icons. Producers set Event.Icon; a row with none takes its group's icon, so
// every row has one.
const (
	// IconSim is sim.finished's pocketwatch.
	IconSim = "inv_misc_pocketwatch_01"
	// IconBook is bis.changed's and tier.moved's book.
	IconBook = "inv_misc_book_09"
)

// groupIcons is each group's icon for a row its producer gave none.
var groupIcons = map[Group]string{
	GroupRaids:   IconBook,
	GroupGear:    "inv_misc_bag_10",
	GroupLogs:    "inv_misc_book_11",
	GroupLadder:  IconSim,
	GroupTiers:   IconBook,
	GroupOfficer: "inv_scroll_08",
	GroupAccount: "inv_misc_gear_01",
}

// kindIcons is the icon of the kinds that always draw the same one, so a row stored without it
// (written before icons) still reads it.
var kindIcons = map[Kind]string{SimFinished: IconSim, BisChanged: IconBook, TierMoved: IconBook}

// IconOf is a row's icon: its own, else its kind's, else its kind's group's.
func IconOf(kind Kind, icon string) string {
	if icon != "" {
		return icon
	}
	if k, ok := kindIcons[kind]; ok {
		return k
	}
	if g, ok := groupIcons[GroupOf(kind)]; ok {
		return g
	}
	return groupIcons[GroupAccount]
}
