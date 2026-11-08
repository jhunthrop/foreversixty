-- addon/ForeverSixty/surfaces/Panel.lua
-- The addon panel, the in-game home (design/specs/2026-11-08-addon-ux.md section 4): one frame
-- that answers "what should I do next on this character", opened from the chip, the character
-- pane's button, the keybind and the addon compartment. It replaces the character pane's flyout.
--
-- Head: the class crest, the name in its class colour, "Forever Score 1240" (the number in its
-- tier colour, absent with no record) and "Level 24 Hunter · <guild>". Then, in priority order,
-- each region only when it has something to say:
--   1. Tonight (a raid within 18 hours or running), or Last raid the day after one;
--   2. Upgrades (at most three, the item tooltip's own Next lines);
--   3. Talents (the next point), or Pick your path under 10 with no build;
--   4. Rotation at <level> (collapsed at 60).
--
-- Pure model, thin frame: Panel.tonight, upgrades, talents, path, rotation, lastRaid, model and
-- fit are plain tables in, plain tables out; Panel.inputs reads the client; Panel.draw lays the
-- frame out. The frame is built on first open, never at load; it never opens in combat and
-- closes when combat starts. Nothing is written into a client frame: docked to the character
-- pane it is anchored to it, never parented.
--
-- The addon never touches the network: a Tonight answer is queued in the companion outbox
-- (contract 4 `raid_answer`) and offered only when the companion says it uploads one (contract 3
-- `accepts`); without that the region says "On the roster" from the raid assignment alone.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Compat = ns.Compat or require("Compat")
local Talents = ns.Talents or require("Talents")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Progress = ns.Progress or require("Progress")
local Gear = ns.Gear or require("Gear")
local Codec = ns.Codec or require("Codec")
local Score = ns.Score or require("Score")
local RaidAssign = ns.RaidAssign or require("RaidAssign")
local ItemTooltip = ns.ItemTooltip or require("ItemTooltip")
local Chip = ns.Chip or require("Chip")
local BuildDialog = ns.BuildDialog or require("BuildDialog")
local PanelTab = ns.PanelTab or require("PanelTab")
local Prefs = ns.Prefs or require("Prefs")

local Panel = {}

Panel.FRAME_NAME = "ForeverSixtyPanel"
Panel.WIDTH = 340
Panel.MAX_HEIGHT = 520
Panel.MAX_UPGRADES = 3
Panel.MAX_ROTATION = 6
--- The height budget's shedding (spec 4.4): the rotation drops to this many rows first, then the
--- upgrades to this many, then the next-ability row goes.
Panel.SHED_ROTATION = 4
Panel.SHED_UPGRADES = 2
--- Tonight shows from this long before a raid's start (and while it runs, RaidAssign.SHOWN_FOR_SECONDS).
Panel.TONIGHT_SECONDS = 18 * 60 * 60
--- Last raid shows while the newest last_raid message is under this old.
Panel.LAST_RAID_SECONDS = 2 * 24 * 60 * 60
--- The four answers, in the apps' and Discord's order.
Panel.ANSWERS = { "in", "late", "bench", "out" }
--- The inbox message the answers read.
Panel.SIGNUP_KIND = "raid_signup"
Codec.MESSAGE_TYPES[Panel.SIGNUP_KIND] = true

--- Nominal heights the budget is counted in (the frame measures its own lines when it draws).
Panel.H = { head = 54, heading = 22, rule = 9, line = 14, row = 34, rotationRow = 20, button = 22, gap = 6 }

-- The pure model --------------------------------------------------------------------------------

--- The newest `raid_signup` for `raidId` in the inbox for this character, or nil.
function Panel.signupFor(inbox, key, raidId)
	for _, message in ipairs(Codec.inboxMessages(inbox, key, Panel.SIGNUP_KIND)) do
		if message.raid_id == raidId then
			return message
		end
	end
	return nil
end

--- The soonest raid the inbox's `raid_signup` messages name within the Tonight window, as
--- RaidAssign.tonight's shape, or nil. Pure.
function Panel.signupTonight(inbox, key, now)
	local soonest
	for _, message in ipairs(Codec.inboxMessages(inbox, key, Panel.SIGNUP_KIND)) do
		local start = RaidAssign.epochOf(message.starts_at)
		if start ~= nil and (now == nil or now <= start + RaidAssign.SHOWN_FOR_SECONDS)
			and (soonest == nil or start < soonest.start) then
			soonest = { raid_id = message.raid_id, raid = message.raid, title = message.title, start = start, jobs = {} }
		end
	end
	return soonest
end

--- Whether the companion uploads `raid_answer` entries (the inbox's `accepts`, contract 3): the
--- feature check for the answer buttons. Pure.
function Panel.acceptsAnswers(inbox)
	if type(inbox) ~= "table" or type(inbox.accepts) ~= "table" then
		return false
	end
	for _, name in ipairs(inbox.accepts) do
		if name == "raid_answer" then
			return true
		end
	end
	return false
end

--- "in 3 h 12 min", "in 25 min", or "now" once the raid has started. Pure.
function Panel.countdown(start, now)
	local left = math.floor((start - now) / 60)
	if left <= 0 then
		return L.panelTonightNow
	end
	if left < 60 then
		return string.format(L.panelInMinutes, left)
	end
	return string.format(L.panelInHours, math.floor(left / 60), left % 60)
end

--- The role word: "Tank", "Healer", "Damage".
local ROLE_KEYS = { tank = "panelRoleTank", healer = "panelRoleHealer", dps = "panelRoleDamage" }
local ANSWER_KEYS = { ["in"] = "panelAnswerIn", late = "panelAnswerLate", bench = "panelAnswerBench",
	out = "panelAnswerOut" }
local ANSWER_COLORS = { ["in"] = "tooltipGreen", late = "clientGold", bench = "tooltipBody", out = "tooltipMuted" }
local BUTTON_KEYS = { ["in"] = "panelButtonIn", late = "panelButtonLate", bench = "panelButtonBench",
	out = "panelButtonOut" }

--- The Tonight region (spec 4.3 item 1), or nil. `input`: { tonight (RaidAssign.tonight's), signup
--- (the raid's raid_signup, or nil), accepts (Panel.acceptsAnswers), pending (the queued outbox
--- entry for this raid, or nil), clock (epoch now), who (the character's name), clockText (the
--- start on the client's clock) }.
---
--- The answer reads, in order: a queued answer the inbox has not confirmed ("In, not sent yet",
--- white, with "Sends when the companion app runs"); the signup's answer ("You're In" green,
--- "Late" gold, "Bench" white, "Out" grey); "Not answered" (gold) with the four buttons when the
--- companion takes answers; with no signup at all, "On the roster" from the assignment alone.
function Panel.tonight(input)
	local tonight = input.tonight
	if type(tonight) ~= "table" or type(tonight.start) ~= "number" or type(input.clock) ~= "number" then
		return nil
	end
	if input.clock < tonight.start - Panel.TONIGHT_SECONDS
		or input.clock > tonight.start + RaidAssign.SHOWN_FOR_SECONDS then
		return nil
	end
	local region = {
		kind = "tonight", raidId = tonight.raid_id, raid = tonight.raid or "",
		note = string.format(L.panelTonightNote, input.clockText or "", Panel.countdown(tonight.start, input.clock)),
	}
	local signup, pending = input.signup, input.pending
	local role = signup ~= nil and ROLE_KEYS[signup.role] or nil
	region.who = role ~= nil and string.format(L.panelAsRole, input.who or "", L[role])
		or string.format(L.panelAs, input.who or "")
	local confirmedAt = signup ~= nil and RaidAssign.epochOf(signup.answered_at) or nil
	if pending ~= nil and ANSWER_KEYS[pending.answer] ~= nil
		and (confirmedAt == nil or confirmedAt < (pending.answered_at or 0)) then
		region.answer = string.format(L.panelAnswerPending, L[BUTTON_KEYS[pending.answer]])
		region.answerColor = "tooltipBody"
		region.sending = L.panelSendsLater
	elseif signup ~= nil and ANSWER_KEYS[signup.answer] ~= nil then
		region.answer, region.answerColor = L[ANSWER_KEYS[signup.answer]], ANSWER_COLORS[signup.answer]
	elseif signup ~= nil and input.accepts then
		region.answer, region.answerColor = L.panelAnswerNone, "clientGold"
		region.buttons = Panel.ANSWERS
	else
		region.answer, region.answerColor = L.panelOnRoster, "tooltipBody"
	end
	local job = (tonight.jobs or {})[1]
	if type(job) == "table" and type(job.label) == "string" then
		region.job = job.label
		region.purpose = job.purpose
		local names = {}
		for name in tostring(job.with or ""):gmatch("[^,]+") do
			names[#names + 1] = (name:gsub("^%s+", ""):gsub("%s+$", ""))
		end
		if #names >= 1 and #names <= 2 then
			region.with = string.format(L.panelWith, table.concat(names, L.panelWithJoin))
		end
	end
	return region
end

--- The Last raid region (spec 4.3, the day after a raid): the date, each boss with its
--- percentile, and the worst boss's biggest loss. nil without a message under two days old.
function Panel.lastRaid(message, clock)
	if type(message) ~= "table" or type(message.bosses) ~= "table" then
		return nil
	end
	local day = Score.dateSeconds(message.date)
	if day == nil or type(clock) ~= "number" or clock - day > Panel.LAST_RAID_SECONDS then
		return nil
	end
	local bosses, worst = {}, nil
	for _, boss in ipairs(message.bosses) do
		if type(boss) == "table" and type(boss.name) == "string" and type(boss.percentile) == "number" then
			bosses[#bosses + 1] = { name = boss.name, percentile = boss.percentile,
				text = string.format(L.homeRaidBoss, boss.name, Chip.ordinal(boss.percentile)) }
			if worst == nil or boss.percentile < worst.percentile then
				worst = boss
			end
		end
	end
	if #bosses == 0 then
		return nil
	end
	local region = { kind = "lastRaid", note = Compat.shortDate(message.date), bosses = bosses }
	if worst ~= nil and type(worst.biggest_loss) == "string" and worst.biggest_loss ~= "" then
		region.loss = string.format(L.panelWorstLoss, worst.name, worst.biggest_loss)
	end
	return region
end

--- The Upgrades region (spec 4.3 item 2): at most MAX_UPGRADES rows from PanelTab.targets (the
--- tooltip's own Next lines, or the bag item a slot's hover leads with), each { itemId, name, hex,
--- gain, unit, detail, bag, link }, `detail` "<slot> · <source>" plus " · level 27" when the item
--- asks for a higher level. `input`: { targets, level, sim (the note's parts: { age, queued } or
--- nil), picksStart }. The rest line under fewer rows. Never nil: the region always shows.
function Panel.upgrades(input)
	local rows = {}
	for _, target in ipairs(input.targets or {}) do
		if #rows >= Panel.MAX_UPGRADES then
			break
		end
		if type(target.gain) == "number" or target.bag then
			local detail = target.source ~= nil and string.format(L.panelTargetDetail, target.slotWord or "", target.source)
				or (target.slotWord or "")
			if type(target.minLevel) == "number" and type(input.level) == "number" and target.minLevel > input.level then
				detail = detail .. string.format(L.panelAtLevel, target.minLevel)
			end
			rows[#rows + 1] = { itemId = target.itemId, name = target.name, hex = target.hex, gain = target.gain,
				unit = target.unit, detail = detail, bag = target.bag == true, link = target.link }
		end
	end
	local region = { kind = "upgrades", rows = rows }
	if input.picksStart ~= nil and #rows == 0 then
		region.rest = string.format(L.homePicksStart, input.picksStart)
	elseif #rows == 0 then
		region.rest = L.panelAllBest
	elseif #rows < Panel.MAX_UPGRADES then
		region.rest = L.panelOthersBest
	end
	if type(input.sim) == "table" then
		region.note = input.sim.age ~= nil and string.format(L.panelSimNote, input.sim.age) or L.panelSimNoteBare
		region.simAgain = input.sim.queued and L.homeSimQueued or L.panelSimAgain
		region.simQueued = input.sim.queued == true
	end
	return region
end

--- The Talents region (spec 4.3 item 3), or nil for a finished build on plan. `input`: { build,
--- buildName, progress (Progress.build's), level, unspent, pointName, pointIcon, pointRank,
--- offPlanTalents }.
function Panel.talents(input)
	local build, progress = input.build, input.progress
	if build == nil or type(progress) ~= "table" then
		return nil
	end
	local offPlan = tonumber(progress.offPlan) or 0
	local allowed = Progress.pointsAllowed(build, input.level)
	local region = { kind = "talents", offPlanTalents = input.offPlanTalents }
	local notes = {}
	if allowed > 0 then
		notes[#notes + 1] = string.format(L.panelOnPlan, math.min(progress.matched or 0, allowed), allowed)
	end
	if offPlan > 0 then
		region.offPlan = string.format(L.panelOffPlanCount, offPlan)
	end
	region.note = #notes > 0 and notes[1] or nil
	if progress.point == nil then
		if offPlan == 0 then
			return nil
		end
		region.finished = true
		return region
	end
	region.icon = input.pointIcon
	region.name = input.pointName or L.panelNextTalentUnnamed
	if (input.unspent or 0) > 0 then
		region.title = region.name
		region.spend = string.format(L.panelSpendNow, input.unspent)
	elseif type(progress.nextLevel) == "number" then
		region.title = string.format(L.panelTalentAt, progress.nextLevel, region.name)
	else
		region.title = region.name
	end
	local rank = input.pointRank
	if type(rank) == "table" and type(rank.want) == "number" and rank.want > 1 then
		region.rank = string.format(L.panelRankOf, (rank.have or 0) + 1, rank.want)
	end
	region.buildName = input.buildName
	return region
end

--- Pick your path (spec 4.3 item 3): under 10 with no build, the class's specs with their
--- one-line playstyle (Chip.pathSpecs). nil otherwise.
function Panel.path(input)
	if input.build ~= nil or type(input.level) ~= "number" or input.level >= 10 then
		return nil
	end
	local specs = input.specs or {}
	if #specs == 0 then
		return nil
	end
	return { kind = "path", specs = specs, foot = L.panelPathFoot }
end

--- The first clause of a plain condition: up to the first ";" or ", and" or ", unless", so the
--- row reads one short line ("Whenever the cooldown it shares with Multi-Shot is up"). Pure.
function Panel.firstClause(text)
	if type(text) ~= "string" then
		return nil
	end
	local cut = text:find(";", 1, true) or text:find(", and ", 1, true) or text:find(", unless ", 1, true)
	local clause = cut ~= nil and text:sub(1, cut - 1) or text
	return (clause:gsub("%.$", ""))
end

--- The Rotation region (spec 4.3 item 4), or nil with no rows. `input`: { rows
--- (PanelTab.rotationRows), level, expanded (the player opened it at 60), newRank (a rank was
--- learned since the panel was last opened) }. Under 60: the learned rows (at most MAX_ROTATION)
--- and one dimmed row for the next ability. At 60: collapsed unless expanded or a new rank.
function Panel.rotation(input)
	local learned, nextRow = {}, nil
	for _, row in ipairs(input.rows or {}) do
		if not row.muted and #learned < Panel.MAX_ROTATION then
			learned[#learned + 1] = { spellId = row.spellId, name = row.name,
				clause = Panel.firstClause(row.fullCondition or row.condition), full = row.fullCondition }
		elseif row.muted and nextRow == nil and type(row.atLevel) == "number" then
			nextRow = { spellId = row.spellId, name = row.name, clause = string.format(L.panelAtLevelRow, row.atLevel),
				full = row.fullCondition }
		end
	end
	if #learned == 0 then
		return nil
	end
	local level = input.level or 0
	local region = { kind = "rotation", level = level, rows = learned, nextRow = nextRow }
	if level >= Progress.MAX_LEVEL then
		region.nextRow = nil
		region.collapsed = not (input.expanded or input.newRank)
	end
	return region
end

--- The regions' nominal heights (Panel.H): what the budget counts in. Pure.
function Panel.regionHeight(region)
	local H = Panel.H
	local height = H.heading + H.rule
	if region.kind == "tonight" then
		height = height + H.row + H.line * (region.sending and 1 or 0) + H.line * (region.with and 2 or 1)
			+ (region.buttons and H.button + H.gap or 0)
	elseif region.kind == "lastRaid" then
		height = height + H.line * (#region.bosses + (region.loss and 1 or 0))
	elseif region.kind == "upgrades" then
		height = height + H.row * #region.rows + (region.rest and H.line * 2 or 0)
	elseif region.kind == "talents" then
		height = height + (region.finished and 0 or H.row + H.button)
	elseif region.kind == "path" then
		height = height + H.row * #region.specs + H.line * 2
	elseif region.kind == "rotation" then
		if region.collapsed then
			return H.heading + H.rule
		end
		height = height + H.rotationRow * (#region.rows + (region.nextRow and 1 or 0))
	end
	return height
end

--- The whole panel's nominal height. Pure.
function Panel.height(model)
	local height = Panel.H.head
	for _, region in ipairs(model.regions) do
		height = height + Panel.regionHeight(region)
	end
	return height
end

local function regionOf(model, kind)
	for _, region in ipairs(model.regions) do
		if region.kind == kind then
			return region
		end
	end
	return nil
end

--- The model within `limit` (spec 4.4): past it the rotation drops to SHED_ROTATION rows, then the
--- upgrades to SHED_UPGRADES, then the next-ability row goes. Tonight and the next talent never
--- drop. Returns the model (shed in place on a copy of the lists).
function Panel.fit(model, limit)
	limit = limit or Panel.MAX_HEIGHT
	local rotation, upgrades = regionOf(model, "rotation"), regionOf(model, "upgrades")
	local steps = {
		function()
			if rotation ~= nil and #rotation.rows > Panel.SHED_ROTATION then
				local kept = {}
				for index = 1, Panel.SHED_ROTATION do
					kept[index] = rotation.rows[index]
				end
				rotation.rows = kept
			end
		end,
		function()
			if upgrades ~= nil and #upgrades.rows > Panel.SHED_UPGRADES then
				local kept = {}
				for index = 1, Panel.SHED_UPGRADES do
					kept[index] = upgrades.rows[index]
				end
				upgrades.rows = kept
			end
		end,
		function()
			if rotation ~= nil then
				rotation.nextRow = nil
			end
		end,
	}
	for _, step in ipairs(steps) do
		if Panel.height(model) <= limit then
			break
		end
		step()
	end
	model.height = math.min(limit, Panel.height(model))
	return model
end

--- The panel's model from its inputs: the head and the regions in order, fitted. Pure.
function Panel.model(input)
	local regions = {}
	local tonight = Panel.tonight(input)
	if tonight ~= nil then
		regions[#regions + 1] = tonight
	else
		local last = Panel.lastRaid(input.lastRaid, input.clock)
		if last ~= nil then
			regions[#regions + 1] = last
		end
	end
	regions[#regions + 1] = Panel.upgrades(input)
	local path = Panel.path(input)
	if path ~= nil then
		regions[#regions + 1] = path
	else
		local talents = Panel.talents(input)
		if talents ~= nil then
			regions[#regions + 1] = talents
		end
	end
	local rotation = Panel.rotation(input)
	if rotation ~= nil then
		regions[#regions + 1] = rotation
	end
	local head = { name = input.who or "", classToken = input.classToken, fact = input.fact }
	if type(input.score) == "table" and type(input.score.score) == "number" then
		head.score = input.score.score
		head.scoreColor = input.score.tierColor
	end
	return Panel.fit({ head = head, regions = regions })
end

-- The raid answer, through the companion outbox ----------------------------------------------

--- This character's queued answer for `raidId`, or nil.
function Panel.pendingAnswer(raidId)
	local character = Export.outboxCharacter()
	local outbox = type(ForeverSixtyDB) == "table" and ForeverSixtyDB.outbox or nil
	local list = Export.mergeRaidAnswers(type(outbox) == "table" and outbox.raid_answer or nil, Export.raidAnswers)
	for _, entry in ipairs(list) do
		if entry.raid_id == raidId and entry.character == character then
			return entry
		end
	end
	return nil
end

--- A Tonight button: the answer queued for the companion at once (Export.saveOutbox writes the
--- outbox; the file is written at the next reload or logout), only when the companion takes
--- answers. False when it did not queue.
function Panel.answer(raidId, answer)
	if raidId == nil or ANSWER_KEYS[answer] == nil or not Panel.acceptsAnswers(ForeverSixtyInbox) then
		return false
	end
	Export.raidAnswers = Export.mergeRaidAnswers(Export.raidAnswers, { {
		character = Export.outboxCharacter(), raid_id = raidId, answer = answer,
		answered_at = type(time) == "function" and time() or 0,
	} })
	Export.saveOutbox(Panel.data or ns.Data)
	Panel.refresh()
	return true
end

-- Reading the client ----------------------------------------------------------------------------

--- The localized class name and the guild for the fact line.
local function factLine(level, classToken)
	local className = type(UnitClass) == "function" and (UnitClass("player"))
		or Talents.titleCase(tostring(classToken or ""))
	local guild = Export.guildInfo()
	local fact = string.format(L.panelFact, level or 0, className or "")
	if guild ~= nil and type(guild.name) == "string" and guild.name ~= "" then
		fact = fact .. L.exportTreeSeparator .. guild.name
	end
	return fact
end

--- The upgrade rows' extras the pure model wants: the slot's word, the item's required level and,
--- for a bag row, its link (the Equip button).
local function upgradeTargets(base)
	local targets = PanelTab.targets(base, function(link)
		return ItemTooltip.sectionFor(link, base)
	end, function(slot)
		return ItemTooltip.bagLead(slot)
	end)
	for _, target in ipairs(targets) do
		target.slotWord = BuildDialog.slotWord(target.slot, _G)
		target.minLevel = Compat.itemMinLevel(target.itemId)
		if target.bag then
			local lead = ItemTooltip.bagLead(target.slot)
			target.link = lead ~= nil and lead.target ~= nil and lead.target.link or nil
			target.source = nil
		end
	end
	return targets
end

--- The newest inbox message of `kind` for this character, the first `keep` accepts.
local function newestMessage(kind, keep)
	for _, message in ipairs(Follow.messages(ForeverSixtyInbox, Export.characterKey(), kind)) do
		if keep == nil or keep(message) then
			return message
		end
	end
	return nil
end

--- The model's inputs from the live modules and the client.
function Panel.inputs(data)
	data = data or Panel.data or ns.Data
	local base = ItemTooltip.baseContext(data)
	local level = Chip.level()
	local classToken = Chip.classToken()
	local ranks = Talents.readRanks(data)
	local readOk, stagedRanks = pcall(Talents.readStagedRanks, data)
	if not readOk then
		stagedRanks = ranks
	end
	local build = Follow.build
	local key = Export.characterKey()
	local clock = Score.clock()
	local name = Compat.playerName()
	local tonight = RaidAssign.tonight(ForeverSixtyInbox, key, clock)
	local accepts = Panel.acceptsAnswers(ForeverSixtyInbox)
	if tonight == nil and accepts then
		tonight = Panel.signupTonight(ForeverSixtyInbox, key, clock)
	end
	local input = {
		level = level, classToken = classToken, who = name, fact = factLine(level, classToken),
		score = name ~= nil and Score.forCharacter(name) or nil,
		clock = clock, tonight = tonight, accepts = accepts,
		lastRaid = newestMessage("last_raid"),
		targets = upgradeTargets(base),
		build = build, buildName = build ~= nil and Follow.buildName(data, build) or nil,
		progress = Progress.build(build, stagedRanks, level),
		unspent = Progress.unspentPoints(level, ranks),
		offPlanTalents = PanelTab.offPlanTalents(build, stagedRanks, data),
		rows = PanelTab.rotationRows(data, build, ranks, level),
		expanded = Panel.rotationExpanded(),
		specs = Chip.pathSpecs(data, Talents.playerClassSlug()),
	}
	if tonight ~= nil then
		input.signup = Panel.signupFor(ForeverSixtyInbox, key, tonight.raid_id)
		input.pending = Panel.pendingAnswer(tonight.raid_id)
		local stamp = Compat.localTime(tonight.start)
		input.clockText = stamp ~= nil and Compat.clockText(stamp.hour, stamp.min) or ""
	end
	local specBis = base.spec ~= nil and type(data.bis) == "table" and data.bis[base.spec] or nil
	input.picksStart = base.faction and Gear.firstBand(specBis, level) or nil
	local current = base.current
	if current ~= nil and current.tag == "sim" then
		local result = newestMessage("sim_result", function(message)
			return message.spec == current.spec
		end)
		local made = result ~= nil and RaidAssign.epochOf(result.generated_at) or nil
		input.sim = { queued = Export.simRequest ~= nil,
			age = made ~= nil and Panel.age(made, clock) or nil }
	end
	local point = input.progress.point
	if point ~= nil and build ~= nil then
		local talent = Talents.cellOf(data, build.classSlug, point.tab, point.tier, point.column)
		input.pointName = Chip.talentName(talent)
		input.pointIcon = talent ~= nil and Talents.iconFor(talent) or nil
		input.pointRank = Chip.pointRankAt(build, stagedRanks, data, point)
	end
	return input
end

--- "2 days ago", "5 h ago", "just now": how old a sim is. Pure.
function Panel.age(made, now)
	local seconds = math.max(0, now - made)
	if seconds >= 2 * 86400 then
		return string.format(L.panelDaysAgo, math.floor(seconds / 86400))
	elseif seconds >= 86400 then
		return L.panelDayAgo
	elseif seconds >= 3600 then
		return string.format(L.panelHoursAgo, math.floor(seconds / 3600))
	end
	return L.panelJustNow
end

--- Whether the player opened the collapsed rotation at 60, remembered per character.
function Panel.rotationExpanded()
	local panel = type(ForeverSixtyDB) == "table" and ForeverSixtyDB.panelState or nil
	local mine = type(panel) == "table" and panel[Export.characterKey() or ""] or nil
	return type(mine) == "table" and mine.rotationOpen == true
end

function Panel.setRotationExpanded(open)
	ForeverSixtyDB = ForeverSixtyDB or {}
	ForeverSixtyDB.panelState = ForeverSixtyDB.panelState or {}
	local key = Export.characterKey() or ""
	ForeverSixtyDB.panelState[key] = ForeverSixtyDB.panelState[key] or {}
	ForeverSixtyDB.panelState[key].rotationOpen = open and true or nil
end

-- What the buttons do -----------------------------------------------------------------------------

--- "Sim again": the sim request queued in the outbox (contract 4), as the flyout's "Sim me" did.
function Panel.simAgain()
	local current = ItemTooltip.baseContext(Panel.data or ns.Data).current
	return PanelTab.simMe(current ~= nil and current.spec or nil)
end

--- "Equip": the bag item, out of combat only.
function Panel.equip(link)
	if link == nil or Theme.inCombat() then
		return false
	end
	return Theme.equip(link)
end

--- "Change": the build dialog on its paste tab.
function Panel.change()
	BuildDialog.data = BuildDialog.data or Panel.data
	return BuildDialog.open("paste")
end

-- The frame ---------------------------------------------------------------------------------------

--- Pools of the frame's parts, reused on every draw: nothing is created after the first draws.
local pools = { text = {}, icon = {}, button = {}, rule = {}, hover = {} }
local used = {}

local function reset()
	for kind, list in pairs(pools) do
		for _, part in ipairs(list) do
			part:Hide()
		end
		used[kind] = 0
	end
end

local function take(kind, build)
	used[kind] = (used[kind] or 0) + 1
	local part = pools[kind][used[kind]]
	if part == nil then
		part = build()
		pools[kind][used[kind]] = part
	end
	part:ClearAllPoints()
	part:Show()
	return part
end

local function text(content, value, hexKey, fontKey)
	local region = take("text", function()
		return Widgets.label(content, "", "body", "small")
	end)
	Theme.applyFont(region, fontKey or "small")
	region:SetText(value or "")
	region:SetTextColor(Theme.rgb(Theme.HEX[hexKey or "body"]))
	region:SetJustifyH("LEFT")
	region:SetWidth(0)
	if type(region.SetWordWrap) == "function" then
		region:SetWordWrap(false)
	end
	return region
end

local function icon(content, path, size)
	local texture = take("icon", function()
		return Theme.icon(content, "ARTWORK", nil, size)
	end)
	texture:SetSize(size, size)
	texture:SetTexture(path or Theme.UNKNOWN_ICON)
	return texture
end

local function button(content, label, onClick)
	local made = take("button", function()
		local built = Widgets.button(content, "", function(self)
			if self.foreverSixtyAction ~= nil then
				self.foreverSixtyAction()
			end
		end)
		built:SetHeight(Panel.H.button - 4)
		return built
	end)
	made:SetText(label)
	made.foreverSixtyAction = onClick
	local width = made.foreverSixtyLabel:GetStringWidth()
	made:SetWidth(math.max(36, math.ceil(type(width) == "number" and width or 30) + 14))
	Widgets.setEnabled(made, true)
	return made
end

local function rule(content)
	local texture = take("rule", function()
		return Theme.texture(content, "ARTWORK", "border")
	end)
	texture:SetHeight(1)
	return texture
end

--- An invisible frame over a row that shows `lines` (or an item) on hover.
local function hover(content, provider)
	local frame = take("hover", function()
		local made = CreateFrame("Frame", nil, content)
		made:EnableMouse(true)
		made:SetScript("OnEnter", function(self)
			local shown = self.foreverSixtyProvider ~= nil and self.foreverSixtyProvider(self) or nil
			if type(shown) == "table" then
				Theme.showLines(self, shown, "ANCHOR_RIGHT")
			end
		end)
		made:SetScript("OnLeave", function()
			Theme.hideTooltip()
		end)
		return made
	end)
	frame.foreverSixtyProvider = provider
	-- A pooled frame keeps no click from its last use (the collapsed rotation's "show").
	frame:SetScript("OnMouseUp", nil)
	return frame
end

local function itemHover(itemId)
	return function(self)
		if type(GameTooltip) == "table" and type(GameTooltip.SetOwner) == "function" and itemId ~= nil then
			GameTooltip:SetOwner(self, "ANCHOR_RIGHT")
			if type(GameTooltip.SetItemByID) == "function" then
				GameTooltip:SetItemByID(itemId)
			else
				GameTooltip:SetHyperlink(Compat.displayLink(itemId))
			end
			GameTooltip:Show()
		end
		return nil
	end
end

local function spellIcon(spellId)
	return Compat.spellTexture(spellId) or Theme.UNKNOWN_ICON
end

--- A raid's icon by its zone name (the client's own icons), a skull for one this table lacks.
Panel.RAID_ICONS = {
	["Onyxia's Lair"] = "Interface\\Icons\\INV_Misc_Head_Dragon_Black",
	["Molten Core"] = "Interface\\Icons\\Spell_Fire_Incinerate",
	["Blackwing Lair"] = "Interface\\Icons\\INV_Misc_Head_Dragon_01",
	["Zul'Gurub"] = "Interface\\Icons\\INV_Misc_Idol_03",
}
Panel.RAID_ICON_DEFAULT = "Interface\\Icons\\INV_Misc_Bone_HumanSkull_01"

--- Percentile colours, the parse colours every log site uses: grey, green, blue, purple, orange,
--- pink at 99 and gold at 100.
Panel.PARSE_COLORS = {
	{ 100, "e5cc80" }, { 99, "e268a8" }, { 95, "ff8000" }, { 75, "a335ee" }, { 50, "0070ff" }, { 25, "1eff00" },
	{ 0, "9d9d9d" },
}

--- A percentile's colour as a hex. Pure.
function Panel.parseHex(percentile)
	for _, step in ipairs(Panel.PARSE_COLORS) do
		if (percentile or 0) >= step[1] then
			return step[2]
		end
	end
	return Panel.PARSE_COLORS[#Panel.PARSE_COLORS][2]
end

--- The frame: the dark panel, the gold title, the crest over the top-left corner, the close X.
--- Built on first open; movable by its title; Escape closes it (UISpecialFrames).
function Panel.ensure()
	if Panel.frame ~= nil then
		return Panel.frame
	end
	local frame = Widgets.panel(UIParent, Panel.WIDTH, 200, Panel.FRAME_NAME)
	frame:SetFrameStrata("MEDIUM")
	frame:SetClampedToScreen(true)
	frame:SetMovable(true)
	frame:EnableMouse(true)
	frame:Hide()
	local title = Widgets.label(frame, L.addonName, "clientGold", "normal")
	title:SetPoint("TOP", frame, "TOP", 0, -6)
	local drag = CreateFrame("Frame", nil, frame)
	drag:SetPoint("TOPLEFT", frame, "TOPLEFT", 0, 0)
	drag:SetPoint("TOPRIGHT", frame, "TOPRIGHT", -24, 0)
	drag:SetHeight(24)
	drag:EnableMouse(true)
	drag:RegisterForDrag("LeftButton")
	drag:SetScript("OnDragStart", function()
		frame:StartMoving()
	end)
	drag:SetScript("OnDragStop", function()
		frame:StopMovingOrSizing()
		Panel.savePosition()
	end)
	local close = Widgets.closeButton(frame, function()
		Panel.close()
	end)
	close:SetPoint("TOPRIGHT", frame, "TOPRIGHT", -4, -4)
	local crest, ring = Theme.buildCrest(frame, 44)
	crest:SetPoint("TOPLEFT", frame, "TOPLEFT", -10, 10)
	Panel.crest, Panel.ring = crest, ring
	local content = CreateFrame("Frame", nil, frame)
	content:SetPoint("TOPLEFT", frame, "TOPLEFT", 14, -30)
	content:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", -14, 10)
	Panel.content = content
	if type(UISpecialFrames) == "table" then
		table.insert(UISpecialFrames, Panel.FRAME_NAME)
	end
	frame:SetScript("OnHide", function()
		Theme.hideTooltip()
	end)
	Panel.frame = frame
	return frame
end

--- The panel's saved spot ({ point, x, y } against UIParent), or nil.
function Panel.savedSpot()
	local spot = type(ForeverSixtyDB) == "table" and ForeverSixtyDB.panel or nil
	if type(spot) == "table" and type(spot.point) == "string" and type(spot.x) == "number"
		and type(spot.y) == "number" then
		return spot
	end
	return nil
end

function Panel.savePosition()
	local frame = Panel.frame
	if frame == nil then
		return
	end
	local point, _, _, x, y = frame:GetPoint()
	if type(point) == "string" then
		ForeverSixtyDB = ForeverSixtyDB or {}
		ForeverSixtyDB.panel = { point = point, x = x, y = y }
	end
end

--- Reset positions (Settings): the saved spot is forgotten.
function Panel.resetPosition()
	if type(ForeverSixtyDB) == "table" then
		ForeverSixtyDB.panel = nil
	end
	if Panel.frame ~= nil and Panel.frame:IsShown() then
		Panel.place(Panel.openedFrom)
	end
end

--- The anchor (spec 4.2): the saved spot first; from the chip TOPRIGHT to its BOTTOMLEFT; from the
--- character pane TOPLEFT to its TOPRIGHT (the mode-tab column's, where it shows); else centred.
function Panel.place(from)
	local frame = Panel.frame
	frame:ClearAllPoints()
	local spot = Panel.savedSpot()
	if spot ~= nil then
		frame:SetPoint(spot.point, UIParent, spot.point, spot.x, spot.y)
		return "saved"
	end
	if from == "chip" and Chip.frame ~= nil then
		frame:SetPoint("TOPRIGHT", Chip.frame, "BOTTOMLEFT", 0, 0)
		return "chip"
	end
	local pane = Compat.clientGlobal("CharacterFrame", "table")
	if from == "pane" and pane ~= nil then
		local column = Widgets.clientFrame(PanelTab.MODE_TABS)
		local anchor = column ~= nil and Widgets.placed(column) and column or pane
		frame:SetPoint("TOPLEFT", anchor, "TOPRIGHT", 0, 0)
		return "pane"
	end
	frame:SetPoint("CENTER", UIParent, "CENTER", 0, 0)
	return "center"
end

-- Drawing the regions ---------------------------------------------------------------------------

local function heading(content, y, label, note, noteKey, width)
	local title = text(content, label, "clientGold", "normal")
	title:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
	if note ~= nil then
		local right = text(content, note, noteKey or "tooltipMuted")
		right:SetJustifyH("RIGHT")
		right:SetPoint("TOPRIGHT", content, "TOPRIGHT", 0, -y - 2)
		right:SetWidth(width - (title:GetStringWidth() or 0) - 12)
	end
	return y + Panel.H.heading
end

local function ruleAt(content, y)
	local line = rule(content)
	line:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y - 4)
	line:SetPoint("TOPRIGHT", content, "TOPRIGHT", 0, -y - 4)
	return y + Panel.H.rule
end

--- A wrapping detail line `width` wide; its measured height.
local function wrapped(content, value, hexKey, x, y, width)
	local region = text(content, value, hexKey)
	region:SetWidth(width)
	if type(region.SetWordWrap) == "function" then
		region:SetWordWrap(true)
	end
	region:SetPoint("TOPLEFT", content, "TOPLEFT", x, -y)
	local height = region:GetStringHeight()
	return type(height) == "number" and height > 0 and math.ceil(height) or Panel.H.line
end

local function drawTonight(content, region, y, width)
	y = heading(content, y, L.panelTonight, region.note, nil, width)
	local mark = icon(content, Panel.RAID_ICONS[region.raid] or Panel.RAID_ICON_DEFAULT, 28)
	mark:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
	local name = text(content, region.raid, "tooltipBody", "normal")
	name:SetPoint("TOPLEFT", content, "TOPLEFT", 36, -y)
	local answer = text(content, region.answer, region.answerColor)
	answer:SetJustifyH("RIGHT")
	answer:SetPoint("TOPRIGHT", content, "TOPRIGHT", 0, -y)
	local who = text(content, region.who, "tooltipMuted")
	who:SetPoint("TOPLEFT", content, "TOPLEFT", 36, -y - Panel.H.line - 2)
	y = y + Panel.H.row
	if region.sending ~= nil then
		local note = text(content, region.sending, "tooltipMuted")
		note:SetPoint("TOPLEFT", content, "TOPLEFT", 36, -y + 2)
		y = y + Panel.H.line
	end
	if region.buttons ~= nil then
		local x = 36
		for _, answerKey in ipairs(region.buttons) do
			local made = button(content, L[BUTTON_KEYS[answerKey]], function()
				Panel.answer(region.raidId, answerKey)
			end)
			made:SetPoint("TOPLEFT", content, "TOPLEFT", x, -y)
			x = x + made:GetWidth() + 6
		end
		y = y + Panel.H.button + Panel.H.gap
	end
	if region.job ~= nil then
		local job = text(content, string.format(L.panelJob, region.job), "tooltipBody")
		job:SetPoint("TOPLEFT", content, "TOPLEFT", 36, -y)
		local over = hover(content, function()
			return region.purpose ~= nil and { region.purpose } or nil
		end)
		over:SetPoint("TOPLEFT", job, "TOPLEFT")
		over:SetPoint("BOTTOMRIGHT", job, "BOTTOMRIGHT")
		y = y + Panel.H.line
		if region.with ~= nil then
			local with = text(content, region.with, "tooltipMuted")
			with:SetPoint("TOPLEFT", content, "TOPLEFT", 36, -y)
			y = y + Panel.H.line
		end
	end
	return y + Panel.H.gap
end

local function drawLastRaid(content, region, y, width)
	y = heading(content, y, L.panelLastRaidTitle, region.note, nil, width)
	for _, boss in ipairs(region.bosses) do
		local line = text(content, boss.name .. " |cff" .. Panel.parseHex(boss.percentile)
			.. Chip.ordinal(boss.percentile) .. "|r", "tooltipBody")
		line:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
		y = y + Panel.H.line
	end
	if region.loss ~= nil then
		y = y + wrapped(content, region.loss, "tooltipMuted", 0, y, width)
	end
	return y + Panel.H.gap
end

local function drawUpgrades(content, region, y, width)
	local note = region.note
	y = heading(content, y, L.panelUpgrades, note, nil, width - (region.simAgain and 70 or 0))
	if region.simAgain ~= nil then
		local again = button(content, region.simAgain, function()
			Panel.simAgain()
		end)
		again:SetPoint("TOPRIGHT", content, "TOPRIGHT", 0, -y + Panel.H.heading - 2)
		Widgets.setEnabled(again, not region.simQueued)
	end
	for _, row in ipairs(region.rows) do
		local mark = icon(content, Compat.itemIcon(row.itemId), 28)
		mark:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
		local lead = row.bag and (Theme.arrowMarkup() .. Theme.inlineColor("tooltipGreen", L.panelInBags) .. " ") or ""
		local name = text(content, lead .. "|cff" .. (row.hex or "ffffff") .. (row.name or "") .. "|r", "body", "normal")
		name:SetPoint("TOPLEFT", content, "TOPLEFT", 36, -y)
		if type(row.gain) == "number" then
			local gain = text(content, ItemTooltip.figure(row.gain, row.unit), "tooltipGreen")
			gain:SetJustifyH("RIGHT")
			gain:SetPoint("TOPRIGHT", content, "TOPRIGHT", 0, -y)
		end
		local detailHeight = Panel.H.line
		if row.bag and row.link ~= nil then
			local equip = button(content, L.panelEquip, function()
				Panel.equip(row.link)
			end)
			equip:SetPoint("TOPRIGHT", content, "TOPRIGHT", 0, -y - Panel.H.line)
			Widgets.setEnabled(equip, not Theme.inCombat())
		else
			detailHeight = wrapped(content, row.detail, "tooltipMuted", 36, y + Panel.H.line, width - 36)
		end
		local over = hover(content, itemHover(row.itemId))
		over:SetPoint("TOPLEFT", mark, "TOPLEFT")
		over:SetPoint("BOTTOMRIGHT", mark, "BOTTOMRIGHT")
		y = y + math.max(Panel.H.row, Panel.H.line + detailHeight + 4)
	end
	if region.rest ~= nil then
		y = y + wrapped(content, region.rest, "tooltipMuted", 0, y, width) + 2
	end
	return y + Panel.H.gap
end

local function drawTalents(content, region, y, width)
	y = heading(content, y, L.panelTalents, region.note, nil, width)
	if region.offPlan ~= nil then
		local off = text(content, region.offPlan, "warning")
		off:SetJustifyH("RIGHT")
		off:SetPoint("TOPRIGHT", content, "TOPRIGHT", 0, -y)
		local over = hover(content, function()
			return region.offPlanTalents
		end)
		over:SetPoint("TOPLEFT", off, "TOPLEFT")
		over:SetPoint("BOTTOMRIGHT", off, "BOTTOMRIGHT")
		if region.finished then
			return y + Panel.H.line + Panel.H.gap
		end
		y = y + Panel.H.line
	end
	local mark = icon(content, region.icon, 28)
	mark:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
	local title = text(content, region.title, "tooltipBody", "normal")
	title:SetPoint("TOPLEFT", content, "TOPLEFT", 36, -y)
	local right = region.spend or region.rank
	if right ~= nil then
		local count = text(content, right, region.spend and "clientGold" or "tooltipMuted")
		count:SetJustifyH("RIGHT")
		count:SetPoint("TOPRIGHT", content, "TOPRIGHT", 0, -y)
	end
	local sub = text(content, (region.buildName or "") .. L.exportTreeSeparator, "tooltipMuted")
	sub:SetPoint("TOPLEFT", content, "TOPLEFT", 36, -y - Panel.H.line - 2)
	local change = button(content, L.panelChange, Panel.change)
	change:SetPoint("LEFT", sub, "RIGHT", 2, 0)
	local open = button(content, L.panelTalentsButton, Theme.toggleTalents)
	open:SetPoint("TOPRIGHT", content, "TOPRIGHT", 0, -y - Panel.H.line)
	return y + Panel.H.row + Panel.H.button + Panel.H.gap - 14
end

local function drawPath(content, region, y, width)
	y = heading(content, y, L.panelPickPath, L.panelPathNote, nil, width)
	for _, spec in ipairs(region.specs) do
		local name = text(content, spec.name, "tooltipBody", "normal")
		name:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
		local pick = button(content, L.panelPick, function()
			Chip.choosePath(spec.slug)
			Panel.refresh()
		end)
		pick:SetPoint("TOPRIGHT", content, "TOPRIGHT", 0, -y)
		local lineHeight = Panel.H.line
		if spec.playstyle ~= nil then
			lineHeight = wrapped(content, spec.playstyle, "tooltipMuted", 0, y + Panel.H.line, width - 60)
		end
		y = y + Panel.H.line + lineHeight + 6
	end
	return y + wrapped(content, region.foot, "tooltipMuted", 0, y, width) + Panel.H.gap
end

local function drawRotation(content, region, y, width)
	if region.collapsed then
		y = heading(content, y, string.format(L.homeRotationTitle, region.level), L.panelShow, "clientGold", width)
		local over = hover(content, nil)
		over:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y + Panel.H.heading)
		over:SetPoint("TOPRIGHT", content, "TOPRIGHT", 0, -y + Panel.H.heading)
		over:SetHeight(Panel.H.heading)
		over:SetScript("OnMouseUp", function()
			Panel.setRotationExpanded(true)
			Panel.refresh()
		end)
		return y
	end
	y = heading(content, y, string.format(L.homeRotationTitle, region.level), L.panelRotationNote, nil, width)
	local rows = {}
	for _, row in ipairs(region.rows) do
		rows[#rows + 1] = { row = row, dim = false }
	end
	if region.nextRow ~= nil then
		rows[#rows + 1] = { row = region.nextRow, dim = true }
	end
	for _, entry in ipairs(rows) do
		local row = entry.row
		local mark = icon(content, spellIcon(row.spellId), 16)
		mark:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
		mark:SetAlpha(entry.dim and Theme.ALPHA.dim or 1)
		local label = text(content, row.name, entry.dim and "tooltipMuted" or "tooltipBody")
		label:SetPoint("TOPLEFT", content, "TOPLEFT", 22, -y)
		local clause = text(content, row.clause, "tooltipMuted")
		clause:SetPoint("LEFT", label, "RIGHT", 6, 0)
		clause:SetWidth(math.max(20, width - 28 - (label:GetStringWidth() or 0)))
		local over = hover(content, function()
			return row.full ~= nil and { row.name, row.full } or nil
		end)
		over:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
		over:SetPoint("TOPRIGHT", content, "TOPRIGHT", 0, -y)
		over:SetHeight(Panel.H.rotationRow)
		y = y + Panel.H.rotationRow
	end
	return y + Panel.H.gap
end

local DRAW = {
	tonight = drawTonight, lastRaid = drawLastRaid, upgrades = drawUpgrades, talents = drawTalents,
	path = drawPath, rotation = drawRotation,
}

--- Lays the model out on the frame and sizes it to fit, never past MAX_HEIGHT.
function Panel.draw(model)
	local frame = Panel.ensure()
	local content = Panel.content
	reset()
	local width = Panel.WIDTH - 28
	Theme.paintCrest(Panel.crest, Panel.ring, model.head.classToken)
	local name = text(content, model.head.name, "tooltipBody", "large")
	if model.head.classToken ~= nil then
		name:SetTextColor(Theme.classColor(model.head.classToken))
	end
	name:SetPoint("TOPLEFT", content, "TOPLEFT", 26, 0)
	if model.head.score ~= nil then
		local color = model.head.scoreColor
		local number = color ~= nil and ("|cff" .. Theme.hexOf(color.r, color.g, color.b) .. model.head.score .. "|r")
			or tostring(model.head.score)
		local score = text(content, string.format(L.panelScore, number), "tooltipMuted")
		score:SetJustifyH("RIGHT")
		score:SetPoint("TOPRIGHT", content, "TOPRIGHT", 0, -2)
	end
	local fact = text(content, model.head.fact, "tooltipMuted")
	fact:SetPoint("TOPLEFT", content, "TOPLEFT", 26, -18)
	local y = Panel.H.head - 20
	for index, region in ipairs(model.regions) do
		if index > 1 then
			y = ruleAt(content, y)
		end
		y = DRAW[region.kind](content, region, y, width)
	end
	frame:SetHeight(math.min(Panel.MAX_HEIGHT, y + 44))
	Panel.current = model
	return model
end

-- Opening and closing -----------------------------------------------------------------------------

--- Opens the panel (from "chip", "pane", "key" or "compartment"), or closes it when it is open.
--- Never in combat. True when it is open afterwards.
function Panel.toggle(from)
	if Panel.frame ~= nil and Panel.frame:IsShown() then
		Panel.close()
		return false
	end
	return Panel.open(from)
end

function Panel.open(from)
	if Theme.inCombat() or Chip.combat then
		return false
	end
	local frame = Panel.ensure()
	Panel.openedFrom = from
	Panel.place(from)
	Panel.refresh()
	frame:Show()
	-- The raid has been seen today: the chip drops it to normal priority until an hour before.
	if Panel.current ~= nil and Panel.current.regions[1] ~= nil and Panel.current.regions[1].kind == "tonight" then
		Chip.raidSeenAt = Score.clock()
		Chip.refresh()
	end
	return true
end

function Panel.close()
	if Panel.frame ~= nil then
		Panel.frame:Hide()
	end
end

function Panel.isOpen()
	return Panel.frame ~= nil and Panel.frame:IsShown() == true
end

--- Redraws the open panel with what the client says now (the chip's events call it).
function Panel.refresh(data)
	if data ~= nil then
		Panel.data = data
	end
	if Panel.frame == nil or Panel.disabled then
		return nil
	end
	local ok, model = pcall(function()
		return Panel.draw(Panel.model(Panel.inputs(Panel.data or ns.Data)))
	end)
	if not ok then
		Panel.disabled = true
		Panel.close()
		Theme.note(string.format(L.diagPanelFailed, tostring(model)))
		return nil
	end
	return model
end

--- PLAYER_REGEN_DISABLED: the panel closes when combat starts (spec 4.2).
function Panel.onCombatStart()
	Panel.close()
end

--- Wires the doors: the chip's click, the pane button, the keybind and the compartment call
--- Panel.toggle through the injection points they leave. Idempotent.
function Panel.register()
	if Panel.registered then
		return
	end
	Panel.registered = true
	Chip.openPanel = function()
		return Panel.toggle("chip")
	end
	PanelTab.openPanel = function()
		return Panel.toggle("pane")
	end
	local frame = CreateFrame("Frame")
	Theme.noteFrame("Panel events", "login")
	Theme.registerEvent(frame, "PLAYER_REGEN_DISABLED")
	frame:SetScript("OnEvent", function()
		Panel.onCombatStart()
	end)
end

Prefs.onReset[#Prefs.onReset + 1] = function()
	Panel.resetPosition()
end

ns.Panel = Panel
return Panel
