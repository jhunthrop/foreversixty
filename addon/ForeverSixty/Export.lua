-- addon/ForeverSixty/Export.lua
-- The character, as the string the planner reads.
--
-- Version 2 (parity contract 7, corrected by 10.5): the version-1 head plus
-- bags, bank and professions, and the raid-prep sections RaidPrep.lua reads
-- (D14). `sets` and `loadouts` are left out -- nothing
-- in the API this addon targets states either, and inventing one would be a
-- lie the planner would then display. Codec.encodeFS1 encodes both sections
-- already, so the day the beta shows an API for them this file grows two
-- lines and the format does not move.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Codec = ns.Codec or require("Codec")
local Talents = ns.Talents or require("Talents")
local Compat = ns.Compat or require("Compat")
local Theme = ns.Theme or require("Theme")
local RaidPrep = ns.RaidPrep or require("RaidPrep")
local RaidDrops = ns.RaidDrops or require("RaidDrops")

local Export = {}

--- The site's slot name -> the client's inventory slot id.
--- finger1/finger2 and trinket1/trinket2 are 11/12 and 13/14; the site
--- numbers them from 1 and the client does too, so the pairing is direct.
Export.INVENTORY_SLOTS = {
	{ slot = "head", id = 1 },
	{ slot = "neck", id = 2 },
	{ slot = "shoulder", id = 3 },
	{ slot = "back", id = 15 },
	{ slot = "chest", id = 5 },
	{ slot = "wrist", id = 9 },
	{ slot = "hands", id = 10 },
	{ slot = "waist", id = 6 },
	{ slot = "legs", id = 7 },
	{ slot = "feet", id = 8 },
	{ slot = "finger1", id = 11 },
	{ slot = "finger2", id = 12 },
	{ slot = "trinket1", id = 13 },
	{ slot = "trinket2", id = 14 },
	{ slot = "main_hand", id = 16 },
	{ slot = "off_hand", id = 17 },
	{ slot = "ranged", id = 18 },
}

--- Bags the player carries. 0 is the backpack; 1..4 are the equipped bags.
Export.CARRIED_BAGS = { 0, 1, 2, 3, 4 }
--- The bank. -1 is the bank window itself; 5..11 are the purchased bag slots.
Export.BANK_BAGS = { -1, 5, 6, 7, 8, 9, 10, 11 }

--- GetCurrentRegion's numbering. Spike check 13 confirms it.
Export.REGION_NAMES = { [1] = "US", [2] = "KR", [3] = "EU", [4] = "TW", [5] = "CN" }

--- How ForeverSixtyDB.savedAt reads. The Export tab shows it verbatim.
Export.SAVED_AT_FORMAT = "%Y-%m-%d %H:%M"

--- The container API moved into C_Container; Classic Era still carries the
--- flat function on some builds. Spike check 9 says which the beta has; this
--- takes whichever exists rather than betting on one.
function Export.containerLink(bag, slot)
	if C_Container and C_Container.GetContainerItemLink then
		return C_Container.GetContainerItemLink(bag, slot)
	end
	return GetContainerItemLink and GetContainerItemLink(bag, slot) or nil
end

function Export.containerSize(bag)
	if C_Container and C_Container.GetContainerNumSlots then
		return C_Container.GetContainerNumSlots(bag) or 0
	end
	return (GetContainerNumSlots and GetContainerNumSlots(bag)) or 0
end

--- An item link's id, or nil for a link this client will not parse.
--- The lookup goes through Compat: the 1.60 client keeps
--- GetItemInfoInstant on C_Item and has no global of that name, which is
--- what the first in-game export died on. Where the client cannot answer,
--- the id is read out of the link text itself.
local function itemIdOf(link)
	if link == nil then
		return nil
	end
	local id = Compat.itemInfoInstant(link)
	if id ~= nil then
		return tonumber(id)
	end
	return tonumber(link:match("item:(%d+)"))
end

--- True when the item can be worn. An export that carried reagents, quest
--- items and bags would be several times longer for nothing: the planner
--- has no slot to put any of them in, `INVTYPE_BAG` included.
local function isEquippable(link)
	local _, _, _, equipSlot = Compat.itemInfoInstant(link)
	return equipSlot ~= nil and equipSlot ~= "" and equipSlot ~= "INVTYPE_NON_EQUIP"
		and equipSlot ~= "INVTYPE_BAG"
end

--- Public because the Export tab counts what is carried and banked
--- without rebuilding the whole string to do it.
function Export.itemsInBags(bags)
	local items = {}
	for _, bag in ipairs(bags) do
		for slot = 1, Export.containerSize(bag) do
			local link = Export.containerLink(bag, slot)
			local id = itemIdOf(link)
			if id ~= nil and isEquippable(link) then
				items[#items + 1] = { itemId = id }
			end
		end
	end
	return items
end

function Export.equippedSlots()
	local slots = {}
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		local id = itemIdOf(GetInventoryItemLink("player", entry.id))
		if id ~= nil then
			slots[#slots + 1] = { slot = entry.slot, itemId = id }
		end
	end
	return slots
end

--- Shared with Gear.lua, which requires Export for the container helpers
--- already and has no other slug vocabulary of its own to keep in.
function Export.slugify(name)
	if name == nil or name == "" then
		return nil
	end
	return (name:lower():gsub("[^%w]+", "-"):gsub("^%-+", ""):gsub("%-+$", ""))
end

--- `table.pack` does not exist on the client's Lua 5.1; this is its shape
--- (positional values plus `n`, the true count including trailing nils)
--- from a single call to `...`.
local function packn(...)
	return { n = select("#", ...), ... }
end

--- Every profession slot the client reports, primaries and secondaries
--- alike. `GetProfessions()` returns five positions and any of them may be
--- nil -- an unlearned primary, most commonly -- so this reads all five by
--- count rather than iterating the return values directly: `ipairs` on a
--- plain table literal stops at the first nil and silently drops every
--- profession after it.
function Export.professionSlugs()
	if GetProfessions == nil then
		return {}
	end
	local slugs = {}
	local slots = packn(GetProfessions())
	for position = 1, slots.n do
		local index = slots[position]
		if index ~= nil then
			local name = GetProfessionInfo and GetProfessionInfo(index)
			local slug = Export.slugify(name)
			if slug ~= nil then
				slugs[#slugs + 1] = slug
			end
		end
	end
	return slugs
end

--- The character's current guild, or nil when unguilded. GetGuildInfo returns nil for
--- an unguilded character (also the shape a login before guild data has loaded would
--- produce, per the WoW API's own documented behaviour -- both read the same way here:
--- no guild= section this export). Rank name (the second return) is player-chosen free
--- text and is never read; only the index is trustworthy (spike check 23, README.md).
function Export.guildInfo()
	local name, _, rankIndex = GetGuildInfo("player")
	if name == nil then
		return nil
	end
	-- The last guild read in the world (X.2): the 1.60 client answers nothing at logout.
	local guild = { name = name, rankIndex = rankIndex, rankNames = Export.rankNames() }
	Export.lastGuild = guild
	Export.guildGone = false
	return guild
end

--- The guild's rank names by rank index (0, the guild master, first), as the client names them
--- (GuildControlGetNumRanks, GuildControlGetRankName, which counts ranks from 1), for the site's raider
--- ranks control; nil where the client cannot say. Display text only: nothing trusts a rank by its name.
function Export.rankNames()
	local count, rankName = _G.GuildControlGetNumRanks, _G.GuildControlGetRankName
	if type(count) ~= "function" or type(rankName) ~= "function" then
		return nil
	end
	local ok, ranks = pcall(count)
	if not ok or type(ranks) ~= "number" or ranks < 1 then
		return nil
	end
	local names = {}
	for index = 1, ranks do
		local read, text = pcall(rankName, index)
		names[index] = read and type(text) == "string" and text or ""
	end
	return names
end

--- The guild read in the world, kept for the logout.
Export.lastGuild = nil
--- Whether the player left their guild this session (Export.onGuildUpdate): no guild is then
--- carried over from the last read or from the record on file.
Export.guildGone = false

--- Who this is: the site shows the name on the character it loads from the
--- string, instead of the class alone. The name is first and last where the
--- client has both (Compat.playerName). nil in a test double without UnitName.
function Export.characterInfo()
	local name = Compat.playerName()
	if name == nil then
		return nil
	end
	return { name = name, realm = GetRealmName() or "" }
end

--- The key's region segment: the client's own region name (Compat.regionName,
--- "US"), else its region number through REGION_NAMES, lowercased ("us");
--- the site's first region ("us") when the client gives neither, the same
--- fallback Compat.siteRegion uses for links. Seen on the beta (16 Oct 2026):
--- GetCurrentRegion answers 110, GetCurrentRegionName "", and the portal
--- cvar "test"; the beta realms are one world, hosted in the US, and the
--- site refuses a record with no region at all. Keys saved before this
--- carried "" there; Export.migrateRegions moves them, and an empty region
--- still matches any region in Follow.sameCharacter.
function Export.regionSegment(regionName, regionIndex)
	if type(regionName) == "string" and regionName ~= "" then
		return regionName:lower()
	end
	local named = Export.REGION_NAMES[regionIndex]
	return named ~= nil and named:lower() or Compat.SITE_REGIONS[1]
end

--- Re-key every saved character record whose region segment is empty under
--- `region` ("/realm/name" becomes "us/realm/name"), stamping the record's
--- region field too, so the companion never posts a record the site refuses.
--- A record already saved under the full key wins; the stale one is dropped.
--- Returns the number of records moved or dropped. Pure.
function Export.migrateRegions(characters, region)
	if type(characters) ~= "table" or type(region) ~= "string" or region == "" then
		return 0
	end
	local moved = 0
	local stale = {}
	for key, record in pairs(characters) do
		if type(key) == "string" and key:sub(1, 1) == "/" and type(record) == "table" then
			stale[#stale + 1] = key
		end
	end
	for _, key in ipairs(stale) do
		local record = characters[key]
		local target = region .. key
		characters[key] = nil
		if characters[target] == nil then
			record.region = region
			characters[target] = record
		end
		moved = moved + 1
	end
	return moved
end

--- Drop the one-word copies an older addon build saved of a two-word Forever
--- character ("Booty" beside "Booty Twerk"): a record whose name has no space
--- goes when another record's name starts with that name plus a space, with
--- the same class (case-insensitive) and either the same realm or an empty
--- realm on the one-word side. A one-word record with no such twin, or whose
--- twin is another class, stays. Returns the number dropped. Pure.
function Export.dropAliases(characters)
	if type(characters) ~= "table" then
		return 0
	end
	local function lower(v)
		return type(v) == "string" and v:lower() or ""
	end
	local doomed = {}
	for key, short in pairs(characters) do
		if type(short) == "table" and type(short.name) == "string" and short.name ~= ""
			and not short.name:find(" ", 1, true) then
			local prefix = short.name .. " "
			for otherKey, long in pairs(characters) do
				if otherKey ~= key and type(long) == "table" and type(long.name) == "string"
					and long.name:sub(1, #prefix) == prefix
					and lower(long.class) == lower(short.class)
					and (short.realm == nil or short.realm == "" or short.realm == long.realm) then
					doomed[#doomed + 1] = key
					break
				end
			end
		end
	end
	for _, key in ipairs(doomed) do
		characters[key] = nil
	end
	return #doomed
end

--- name, realm, region as both Export.characterKey and Export.save need
--- them -- one path, so the key ForeverSixtyDB.characters and
--- ForeverSixtyDB.follows both index by can never drift from the record
--- Export.save writes. Guarded rather than called straight (unlike
--- characterInfo above): Follow reaches this on every load/restore/forget,
--- including from specs that never install the client mock at all.
local function characterParts()
	local name = Compat.playerName() or "player"
	local realm = type(GetRealmName) == "function" and GetRealmName() or ""
	local regionIndex = type(GetCurrentRegion) == "function" and GetCurrentRegion() or nil
	return name, realm, Export.regionSegment(Compat.regionName(), regionIndex)
end

--- "<region>/<realm>/<name>": the key every saved record and inbox message is
--- addressed by. Pure.
function Export.keyFor(name, realm, region)
	return (region or "") .. "/" .. (realm or "") .. "/" .. name
end

local characterKeyOf = Export.keyFor

--- The account-wide key this client's saved data indexes by: the same
--- region/realm/name Export.save has always written into
--- ForeverSixtyDB.characters (realm stands in for ruleset, spike check
--- 12). Follow's per-character state and inbox filter use this so a
--- second character logging in never reads the first one's build.
function Export.characterKey()
	return characterKeyOf(characterParts())
end

--- Move a character's records saved under a key Follow.sameCharacter
--- matches but that is not `key` itself (the old empty-region key) to `key`,
--- in ForeverSixtyDB.characters and .follows, at login before anything
--- reads them. A record already under `key` is kept; another character's
--- is never touched. Pure over `db`.
function Export.adoptKey(db, key)
	if type(db) ~= "table" or key == nil then
		return db
	end
	local Follow = ns.Follow or require("Follow")
	for _, field in ipairs({ "characters", "follows" }) do
		local records = db[field]
		if type(records) == "table" and records[key] == nil then
			for oldKey, record in pairs(records) do
				if oldKey ~= key and Follow.sameCharacter(oldKey, key) then
					records[key] = record
					records[oldKey] = nil
					break
				end
			end
		end
	end
	return db
end

--- The export string, or nil and the reason. `fallbackGuild` ({ name, rankIndex }) stands in
--- when the client answers no guild (X.2: it does not at logout).
function Export.string(data, fallbackGuild)
	-- The class slug comes from the locale-neutral class token (see
	-- Talents.playerClassSlug); slugifying the display name on a non-English
	-- client would produce a slug data.classes does not carry. Race gets the
	-- same treatment via Compat.playerRaceSlug.
	local classSlug = Talents.playerClassSlug()
	local treeRanks = Talents.treeRanks(data, classSlug)
	if treeRanks == nil then
		return nil, string.format(L.codecUnknownClass, tostring(classSlug))
	end
	-- Resistances, consumables, keyring and attunement quests (D14): RaidPrep reads the client.
	local prep = RaidPrep.read(Export.CARRIED_BAGS, { size = Export.containerSize, link = Export.containerLink })
	return Codec.encodeFS1({
		resistances = prep.resistances,
		supplies = prep.supplies,
		keyring = prep.keyring,
		quests = prep.quests,
		owned = prep.owned,
		-- What dropped in this character's raids (RaidDrops.lua), newest corpses within budget.
		drops = RaidDrops.forExport(),
		dataBuild = data.build,
		classSlug = classSlug,
		raceSlug = Compat.playerRaceSlug(),
		-- The level-aware sim design (2026-09-27): the site no longer has to guess this
		-- from talent points. UnitLevel always answers in the real client; wow_mock's own
		-- default (state.level or 60) keeps every existing test double answering too.
		level = UnitLevel("player"),
		treeRanks = treeRanks,
		gearSlots = Export.equippedSlots(),
		bags = Export.itemsInBags(Export.CARRIED_BAGS),
		bank = Export.itemsInBags(Export.BANK_BAGS),
		professions = Export.professionSlugs(),
		guild = Export.guildInfo() or fallbackGuild,
		character = Export.characterInfo(),
	})
end

--- Whether an export code carries at least one gear slot: its head's sixth field (what follows
--- the talents, up to the first section) names a slot. Pure.
function Export.hasGear(code)
	if type(code) ~= "string" then
		return false
	end
	local head = code:match("^[^|]*")
	local gear = head:match("^[^:]*:[^:]*:[^:]*:[^:]*:[^:]*:(.*)$")
	return gear ~= nil and gear:find("=", 1, true) ~= nil
end

--- How many gear slots an export code carries and its talent ranks ("a/b/c"), for /fs diag;
--- nil for anything that is not a code. Pure.
function Export.describe(code)
	if type(code) ~= "string" then
		return nil
	end
	local head = code:match("^[^|]*")
	local talents, gear = head:match("^[^:]*:[^:]*:[^:]*:[^:]*:([^:]*):?(.*)$")
	if talents == nil then
		return nil
	end
	local slots = 0
	for _ in gear:gmatch("[^,]+=") do
		slots = slots + 1
	end
	return slots, talents
end

--- The guild to carry when the client answers none: the one read earlier in the world, else the
--- one on the record on file, unless the player left the guild this session.
local function keptGuild(record)
	if Export.guildGone then
		return nil
	end
	if Export.lastGuild ~= nil then
		return Export.lastGuild
	end
	if type(record) == "table" and type(record.export) == "string" then
		local build = Codec.decodeFS1(record.export)
		return build and build.guild or nil
	end
	return nil
end

--- Write the record the companion reads out of SavedVariables.
---
--- The shape is companion/internal/addon/addon.go's: any table carrying a
--- string `export` with `name`, `ruleset` and `region` beside it. `ruleset`
--- is the realm name until the beta shows an API that states Forever's own
--- ruleset (spike check 12); the companion falls back to `realm` for an
--- older addon, and this writes both so neither side has to guess.
function Export.save(data)
	return Export.write(data)
end

--- Build the export and write the character record: the one writer behind Export.save (logout)
--- and Export.refresh (in the world). A read with no gear never replaces a record that has
--- gear (the 1.60 client answers empty at PLAYER_LOGOUT); that record's export is kept (the
--- DB's `savedAt` stamped) and returned.
function Export.write(data)
	-- UnitName("player") always succeeds in the real client; the "player"
	-- fallback only fires in a test double (wow_mock.lua does not stub
	-- UnitName) or a future client that drops the API, so the record is
	-- still written rather than erroring.
	local name, realm, region = characterParts()
	local key = characterKeyOf(name, realm, region)
	local before = ForeverSixtyDB and ForeverSixtyDB.characters and ForeverSixtyDB.characters[key]
	local code, message = Export.string(data, keptGuild(before))
	if code == nil then
		return nil, message
	end
	ForeverSixtyDB = ForeverSixtyDB or {}
	ForeverSixtyDB.characters = ForeverSixtyDB.characters or {}
	-- Stamped on the DB, not on the character record: the Export tab shows
	-- one "last saved" line, and a player with four characters wants the
	-- last time anything was written, not the last time this one was.
	ForeverSixtyDB.savedAt = date(Export.SAVED_AT_FORMAT)
	Export.migrateRegions(ForeverSixtyDB.characters, region)
	Export.dropAliases(ForeverSixtyDB.characters)
	before = ForeverSixtyDB.characters[key] or before
	if not Export.hasGear(code) and before ~= nil and Export.hasGear(before.export) then
		return before.export
	end
	ForeverSixtyDB.characters[key] = {
		name = name,
		class = UnitClass("player"),
		realm = realm,
		ruleset = realm,
		region = region,
		build = data.build,
		export = code,
		builtAt = date(Export.SAVED_AT_FORMAT),
		-- When this record first carried gear (X.3): the login line tells a player whose last
		-- upload the site saw before that, that their gear is on its way. Kept while the record
		-- keeps its gear; nil on a record that predates the field.
		gearAt = Export.hasGear(code) and not (before ~= nil and Export.hasGear(before.export))
			and date(Export.SAVED_AT_FORMAT) or (before ~= nil and before.gearAt or nil),
	}
	return code
end

-- The companion outbox (contract 4, docs/contracts/addon-v2.md) ---------------------------

--- The outbox format this addon writes; the companion ignores a format it does not know.
Export.OUTBOX_FORMAT = 1

--- The sim request L4.3's "Sim me" leaves for the next logout, or nil: { character, spec,
--- requested_at, gear_code }, all as of the request.
Export.simRequest = nil

--- The raid answers queued in game this session (the addon panel's Tonight buttons), each
--- { character, raid_id, answer, late_minutes?, answered_at }: contract 4's `raid_answer`.
Export.raidAnswers = nil

--- A queued answer older than this is dropped from the outbox (its raid is long over).
Export.RAID_ANSWER_KEEP_SECONDS = 7 * 24 * 60 * 60

--- The queued `raid_answer` entries (contract 4), from the outbox the companion has not cleared
--- yet and this session's, newest per raid, none older than ANSWER_KEEP_SECONDS. Pure.
function Export.mergeRaidAnswers(saved, added, now)
	local byRaid, order = {}, {}
	local function take(entry)
		if type(entry) ~= "table" or entry.raid_id == nil or type(entry.answered_at) ~= "number" then
			return
		end
		if type(now) == "number" and now - entry.answered_at > Export.RAID_ANSWER_KEEP_SECONDS then
			return
		end
		local key = tostring(entry.character) .. "\n" .. tostring(entry.raid_id)
		local kept = byRaid[key]
		if kept == nil then
			order[#order + 1] = key
		end
		if kept == nil or entry.answered_at >= kept.answered_at then
			byRaid[key] = entry
		end
	end
	for _, entry in ipairs(saved or {}) do
		take(entry)
	end
	for _, entry in ipairs(added or {}) do
		take(entry)
	end
	local list = {}
	for _, key in ipairs(order) do
		list[#list + 1] = byRaid[key]
	end
	return list
end

--- The class tokens of the client's guild roster, by name without the realm. Empty where the
--- client has not sent the roster yet.
local function rosterClasses()
	local classes = {}
	for _, row in ipairs(Theme.guildRoster()) do
		classes[(row.name:gsub("%-.*$", ""))] = row.classToken
	end
	return classes
end

--- An opt-out code's letters: "FS-" and six of A-Z (no I or O) and 2-9, the site's own alphabet.
local OPT_OUT_CODE = "FS%-([A-HJ-NP-Z2-9][A-HJ-NP-Z2-9][A-HJ-NP-Z2-9][A-HJ-NP-Z2-9][A-HJ-NP-Z2-9][A-HJ-NP-Z2-9])"

--- The "Remove my character" codes in the guild roster's public notes: { name, code } for each
--- member whose note carries one, sorted by name. Only the code leaves the client, never the
--- rest of the note. `roster` is Theme.guildRoster(). Pure.
function Export.optOutCodes(roster)
	local codes = {}
	for _, row in ipairs(roster or {}) do
		local letters = type(row.note) == "string" and row.note:match(OPT_OUT_CODE) or nil
		if letters ~= nil then
			codes[#codes + 1] = { name = row.name, code = "FS-" .. letters }
		end
	end
	table.sort(codes, function(left, right)
		return left.name < right.name
	end)
	return codes
end

--- The roster snapshot of `guild` from its members and opt-out codes; nil when it has neither.
local function rosterSnapshot(guild, members, codes)
	if #members == 0 and #codes == 0 then
		return nil
	end
	return { guild = guild, members = members, optout_codes = #codes > 0 and codes or nil }
end

--- The members of `guild` this session heard on GUILD, by name: the seven contract fields,
--- `seen_at` as local wall time. `seen` is Comms.seen (its `time` is the game clock), `now`
--- the game clock and `wallNow` the wall clock at the same moment; `classes` is the class
--- token by name, and a name it lacks leaves `class` out.
local function rosterMembers(seen, now, wallNow, classes)
	local members = {}
	for name, entry in pairs(seen or {}) do
		if entry.channel == "GUILD" then
			members[#members + 1] = {
				name = name,
				level = entry.level or 0,
				spec = entry.spec or "-",
				gear_gap = entry.gearGap or 0,
				build = entry.buildName or "",
				seen_at = date(Export.SAVED_AT_FORMAT, wallNow - (now - (entry.time or now))),
				class = (classes or {})[(name:gsub("%-.*$", ""))],
			}
		end
	end
	table.sort(members, function(left, right)
		return left.name < right.name
	end)
	return members
end

--- Contract 4's `export.score_visible`: false when the player hid the score, absent (nil) otherwise.
function Export.scoreVisibleField(visible)
	if visible == false then
		return false
	end
	return nil
end

--- The outbox, built from plain values: { character, code, wallNow, simRequest, raidAnswers, guild (a name,
--- or nil/false outside a guild), seen, now, classes, optOutCodes }. Pure.
function Export.outbox(parts)
	local roster
	if parts.guild then
		local members = rosterMembers(parts.seen, parts.now or 0, parts.wallNow, parts.classes)
		roster = rosterSnapshot(parts.guild, members, parts.optOutCodes or {})
	end
	return {
		format = Export.OUTBOX_FORMAT,
		export = {
			character = parts.character,
			code = parts.code,
			saved_at = date(Export.SAVED_AT_FORMAT, parts.wallNow),
			-- Contract 4: false when the player hid their score from others; absent (visible) otherwise.
			score_visible = Export.scoreVisibleField(parts.scoreVisible),
		},
		sim_request = parts.simRequest,
		roster_snapshot = roster,
		raid_answer = parts.raidAnswers ~= nil and #parts.raidAnswers > 0 and parts.raidAnswers or nil,
	}
end

--- The outbox's `character`: the full name exactly as the export's who= field
--- writes it (Compat.playerName, the same source as Export.characterKey),
--- then the realm with its spaces and hyphens gone: "Obnoxious Yell-
--- ClassicBetaPvP". The API matches it against the synced characters by
--- that name. The name alone with no realm. Pure.
function Export.outboxName(name, realm)
	local compact = (realm or ""):gsub("[%s%-]", "")
	if compact == "" then
		return name
	end
	return name .. "-" .. compact
end

--- The words a realm's name spells a ruleset with, and the ruleset segment of a site path for a
--- realm's name: Codec's (the one copy, which the key matcher reads too). Pure.
Export.RULESET_WORDS = Codec.RULESET_WORDS
Export.rulesetFromRealm = Codec.rulesetFromRealm

--- A character's slug as the site and the API write it (characters.ts characterSlug,
--- api/internal/character Slug): the name alone, trimmed, lower-cased, each run of spaces
--- one hyphen. Pure.
function Export.characterSlug(name)
	local trimmed = (name or ""):match("^%s*(.-)%s*$")
	-- The common accented capitals (L3.7 item 33): U+00C0 to U+00DE except the multiplication sign
	-- U+00D7, two bytes C3 80..9E, lower by the 0x20 the site's lower-casing adds.
	local lowered = trimmed:lower():gsub("\195([\128-\158])", function(second)
		if second == "\151" then
			return nil
		end
		return "\195" .. string.char(second:byte() + 32)
	end)
	return (lowered:gsub("%s+", "-"))
end

--- A character's page on the site: SITE/character/<region>/<ruleset>/<slug>, the one form the
--- site's parsePath serves: the region ("us" where none is given; Compat.siteRegion is the
--- client's), the ruleset the realm's words name (Export.rulesetFromRealm) and the name's slug,
--- no realm in it. Pure.
function Export.siteLink(name, realm, region)
	if type(region) ~= "string" or region == "" then
		region = Compat.SITE_REGIONS[1]
	end
	return L.siteName .. "/character/" .. region:lower() .. "/" .. Export.rulesetFromRealm(realm) .. "/"
		.. Export.characterSlug(name)
end

--- This character's name as the outbox writes it (Export.outboxName), so the build
--- dialog can tell whether the outbox it reads is this character's.
function Export.outboxCharacter()
	local name, realm = characterParts()
	return Export.outboxName(name, realm)
end

--- Write ForeverSixtyDB.outbox at logout, beside Export.save: the export (`code` when the
--- caller already built it), any pending sim request, and the guild's members this session
--- heard. nil and the reason when there is no export to carry.
function Export.saveOutbox(data, code)
	local message
	if code == nil then
		code, message = Export.string(data)
	end
	if code == nil then
		return nil, message
	end
	local name, realm = characterParts()
	local live = Export.guildInfo()
	local key = characterKeyOf(characterParts())
	local record = ForeverSixtyDB and ForeverSixtyDB.characters and ForeverSixtyDB.characters[key]
	local guild = live or keptGuild(record)
	-- Comms loads after this file; by logout it is there.
	local Comms = ns.Comms or require("Comms")
	local now = Theme.now()
	Comms.prune(now)
	ForeverSixtyDB = ForeverSixtyDB or {}
	local outbox = Export.outbox({
		character = Export.outboxName(name, realm),
		code = code,
		wallNow = time(),
		simRequest = Export.simRequest,
		scoreVisible = (ns.Prefs or require("Prefs")).flag("showScore") ~= false,
		-- Answers the companion has not cleared from the file yet, with this session's (contract 4).
		raidAnswers = Export.mergeRaidAnswers(type(ForeverSixtyDB.outbox) == "table"
			and ForeverSixtyDB.outbox.raid_answer or nil, Export.raidAnswers, time()),
		guild = guild ~= nil and guild.name or nil,
		seen = Comms.seen,
		now = now,
		classes = rosterClasses(),
		optOutCodes = Export.optOutCodes(Theme.guildRoster()),
	})
	-- The roster heard in the world stands when logout finds none to read (X.2).
	if outbox.roster_snapshot == nil and guild ~= nil and Export.lastRoster ~= nil
		and Export.lastRoster.guild == guild.name then
		outbox.roster_snapshot = Export.lastRoster
	end
	ForeverSixtyDB.outbox = outbox
	return outbox
end

-- The export kept while in the world (X.2) ---------------------------------------------------

--- The least seconds between two builds on equipment, bag and guild events.
Export.REFRESH_INTERVAL = 10
--- Seconds after PLAYER_ENTERING_WORLD of the first build, and of the one retry.
Export.WORLD_BUILD_DELAY = 5
Export.WORLD_RETRY_DELAY = 20

--- The roster the outbox would carry now, kept for the logout.
Export.lastRoster = nil
local lastBuild, pending

local function snapshotRoster()
	local guild = Export.lastGuild
	if guild == nil then
		return
	end
	local Comms = ns.Comms or require("Comms")
	local members = rosterMembers(Comms.seen, Theme.now(), time(), rosterClasses())
	local snapshot = rosterSnapshot(guild.name, members, Export.optOutCodes(Theme.guildRoster()))
	if snapshot ~= nil then
		Export.lastRoster = snapshot
	end
end

--- Build the export and write the character record now (no outbox); the code, or nil and the reason.
function Export.refresh(data)
	lastBuild = Theme.now()
	local code, message = Export.write(data)
	snapshotRoster()
	return code, message
end

--- Refresh now, or at most once per Export.REFRESH_INTERVAL: a build inside the window is
--- deferred to its end (one at a time, however many events ask), never dropped.
function Export.refreshSoon(data)
	local now = Theme.now()
	if lastBuild == nil or now - lastBuild >= Export.REFRESH_INTERVAL then
		return Export.refresh(data)
	end
	if not pending then
		pending = Theme.after(Export.REFRESH_INTERVAL - (now - lastBuild), function()
			pending = false
			Export.refresh(data)
		end)
	end
end

--- PLAYER_ENTERING_WORLD: build five seconds in, once the item cache is warm; retry once at
--- twenty seconds if a slot still read nothing.
function Export.afterEnteringWorld(data)
	Theme.after(Export.WORLD_BUILD_DELAY, function()
		Export.refresh(data)
		if #Export.equippedSlots() < #Export.INVENTORY_SLOTS then
			Theme.after(Export.WORLD_RETRY_DELAY - Export.WORLD_BUILD_DELAY, function()
				Export.refresh(data)
			end)
		end
	end)
end

--- PLAYER_GUILD_UPDATE: the guild is read again; a client that says the player is in none (and
--- answers no guild) means they left, so the last guild is dropped rather than carried.
function Export.onGuildUpdate(data)
	if Export.guildInfo() == nil and Theme.inGuild() == false then
		Export.lastGuild = nil
		Export.lastRoster = nil
		Export.guildGone = true
	end
	Export.refreshSoon(data)
end

--- /fs diag's line: when the record on file was built, how many slots it carries, the talents.
function Export.diagLine()
	local record = ForeverSixtyDB and ForeverSixtyDB.characters and ForeverSixtyDB.characters[Export.characterKey()]
	if type(record) ~= "table" or type(record.export) ~= "string" then
		return L.diagExportNone
	end
	local slots, talents = Export.describe(record.export)
	local stamp = type(record.builtAt) == "string" and record.builtAt:match(" (%d+:%d+)$") or nil
	-- In the chat's own clock (Compat.clockText), so it never sits beside a differing one (X.5).
	stamp = stamp ~= nil and Compat.clockOf(stamp) or "?"
	return string.format(L.diagExport, stamp, slots or 0, talents or "?")
end

ns.Export = Export
return Export
