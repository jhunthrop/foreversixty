package main

// setNativeClass maps a set_id (data/builds/<build>/items/<class>.json's
// own set_id field) to the class slug wowsims-forever's engine source
// wrote that set's ApplyEffect functions for - the class every
// `agent.(XAgent)` cast inside that set's Bonuses map assumes the
// wearer is.
//
// Generated once, from wowsims-forever @ 8d39f2b29 (sim/enginever.
// Version as of this lane's run), NOT from any of this site's own
// data: for every class package (sim/warrior, sim/paladin, sim/hunter,
// sim/rogue, sim/priest, sim/shaman, sim/mage, sim/warlock, sim/druid -
// sim/common/item_sets/*.go is deliberately excluded, because its sets
// use agent.GetCharacter(), not a class-specific Agent cast, so they
// are safe for any class and need no entry here), every
//
//	core.NewItemSet(core.ItemSet{ ... Name: "<set name>" ... })
//
// in that package's item_sets_pve.go/item_sets_pvp.go (rogue:
// items_sets_pve.go/items_sets_pvp.go) was matched by name against
// assets/database/db.json's own items[].setName/setId (the wowhead-
// scraped fact table sim/cmd/leveling-bis's own candidate data does
// NOT carry - see band.go's crossClassSetItem doc for why that data
// gap is what makes this table necessary rather than reading the
// field straight off a candidate). No two classes claimed the same set
// name in this engine build; 91 of 96 set names in these 9 packages'
// source resolved to a set id (the other 5, all sim/warrior, name sets
// with no released items in this build's db.json - AQ40/Naxx warrior
// tier the client does not carry yet - so no candidate can ever carry
// their set_id and they need no entry).
//
// Regenerate by re-running this lane's report's exact two-pass script
// (grep sim/<class>/item_sets*.go for Name: "..." per class package,
// then look each name up in wowsims-forever's assets/database/db.json
// for its setId) against the engine sha sim/enginever.Version names
// next, whenever a set item this command has never ranked turns up in
// a new engine pin - a set id absent from this map is treated as
// "no known cross-class restriction" (see crossClassSetItem), which is
// only safe for sets this map has not fallen behind on.
var setNativeClass = map[int]string{
	185: "druid",
	201: "mage",
	202: "priest",
	203: "warlock",
	204: "rogue",
	205: "druid",
	206: "hunter",
	207: "shaman",
	208: "paladin",
	209: "warrior",
	210: "mage",
	211: "priest",
	212: "warlock",
	213: "rogue",
	214: "druid",
	215: "hunter",
	216: "shaman",
	217: "paladin",
	218: "warrior",
	301: "shaman",
	342: "priest",
	344: "priest",
	361: "hunter",
	362: "hunter",
	383: "warrior",
	384: "warrior",
	386: "shaman",
	387: "mage",
	388: "mage",
	389: "priest",
	390: "priest",
	391: "warlock",
	392: "warlock",
	393: "rogue",
	394: "rogue",
	395: "hunter",
	396: "hunter",
	397: "druid",
	398: "druid",
	402: "paladin",
	474: "warrior",
	475: "paladin",
	477: "hunter",
	478: "rogue",
	480: "priest",
	481: "warlock",
	482: "mage",
	494: "druid",
	495: "warrior",
	496: "warrior",
	497: "rogue",
	498: "rogue",
	499: "warlock",
	501: "shaman",
	502: "shaman",
	503: "mage",
	505: "paladin",
	506: "paladin",
	507: "priest",
	509: "hunter",
	511: "warrior",
	512: "rogue",
	513: "druid",
	515: "hunter",
	517: "mage",
	518: "warlock",
	519: "shaman",
	522: "rogue",
	523: "warrior",
	524: "rogue",
	525: "priest",
	526: "mage",
	527: "shaman",
	528: "paladin",
	529: "warlock",
	530: "hunter",
	537: "warrior",
	538: "shaman",
	539: "druid",
	540: "priest",
	541: "warlock",
	542: "mage",
	543: "hunter",
	544: "paladin",
	545: "warrior",
	546: "mage",
	547: "warlock",
	548: "rogue",
	549: "priest",
	550: "hunter",
	551: "druid",
}
