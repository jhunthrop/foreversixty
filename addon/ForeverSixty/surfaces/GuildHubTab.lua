-- addon/ForeverSixty/surfaces/GuildHubTab.lua
-- The Forever Sixty window's Guild tab (design/specs/2026-11-13-addon-hub.md 4.5, board hub-guild.png):
-- are we raiding tonight, am I in, what is my job, and am I ready?
--
-- Left: Tonight (the raid, its time, your sign-up with the four answers, your jobs), Your prep (what the
-- raid asks of you, RaidPrep.summary), Sign-ups (the guild's answers and roles against the raid's comp
-- target, from the companion's raid_roster). Right: Next raids (each with your answer; a click puts it
-- in Tonight's place), the guild (its standing, then the top five raiders by Forever Score and your
-- place, from the ratings pack), Officers sent (the last build an officer pushed, and Follow).
--
-- Every section reads only what the companion and the ratings pack carry; a section with nothing to
-- say is left out (spec 8). An answer is queued in the companion outbox (contract 4 `raid_answer`) and
-- offered only when the companion says it uploads one (contract 3 `accepts`); the addon never touches
-- the network. Pure model (GuildHubTab.model and its parts), thin frame (GuildHubTab.draw).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Compat = ns.Compat or require("Compat")
local Codec = ns.Codec or require("Codec")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local GuildRaid = ns.GuildRaid or require("GuildRaid")
local RaidPrep = ns.RaidPrep or require("RaidPrep")
local BuildsPane = ns.BuildsPane or require("BuildsPane")
local HubDraw = ns.HubDraw or require("HubDraw")
local Hub = ns.Hub or require("Hub")

local GuildHubTab = {}

--- Pixel layout: the gap between sections, a section heading's line, a text line, the raid icon and
--- the gap after it, the answer buttons (their least width and the gap between them), the gap above
--- the buttons, the standings' columns (rank, the spec icon and its gap, score, attendance), the
--- Next raids' date and raid columns, the build icon of Officers sent, Your prep's label column, the
--- least Follow width, the raid's name size.
GuildHubTab.LAYOUT = {
	sectionGap = 20, head = 18, line = 15, raidIcon = 32, iconGap = 8, answerWidth = 52, answerGap = 4,
	buttonsGap = 6, rankWidth = 18, specIcon = 14, specGap = 4, scoreWidth = 44, attendanceWidth = 40,
	dateWidth = 112, raidWidth = 98, buildIcon = 20, prepLabel = 132, followWidth = 64, raidName = 14, ownGap = 4,
}

local SEP = L.dialogSeparator

-- The model -------------------------------------------------------------------------------------

--- Officers sent (spec 4.5 right 3): the newest pushed build { code, name, from, day, following }, or nil.
--- `following(code)` answers whether the player follows that build now. Pure over its inputs.
function GuildHubTab.officersSent(pushed, following)
	local newest = type(pushed) == "table" and pushed[1] or nil
	if type(newest) ~= "table" or type(newest.code) ~= "string" then
		return nil
	end
	return { code = newest.code, name = newest.name, from = newest.from, day = GuildRaid.weekday(newest.at),
		following = following(newest.code) == true }
end

--- The shipped data (ns.Data); a spec sets it.
GuildHubTab.data = nil

local function dataOf()
	return GuildHubTab.data or GuildRaid.data or ns.Data
end

--- Whether the player follows the build `code` now: the same talent order as the build followed.
local function followsCode(code)
	local build = Follow.build
	if build == nil then
		return false
	end
	local pushed = Codec.loadBuild(code, dataOf())
	return pushed ~= nil and Follow.orderDiffCount(build, pushed) == 0
end

--- The whole tab's model, read now. `selectedId` names the raid on top (a Next raids click), else the
--- soonest. { tonight, empty, prep, prepFor, signups, nextRaids, guild, sent }.
function GuildHubTab.model(selectedId)
	local ctx = GuildRaid.context()
	local raids = GuildRaid.liveRaids()
	local shown = raids[1]
	for _, raid in ipairs(raids) do
		if raid.raidId == selectedId then
			shown = raid
		end
	end
	local guildName = Theme.playerGuild()
	local model = {}
	if shown ~= nil then
		model.tonight = GuildRaid.tonight(shown, ctx)
		local container = { size = Export.containerSize, link = Export.containerLink }
		local prep = RaidPrep.summary(shown.raid, dataOf(), Export.CARRIED_BAGS, container)
		if #prep > 0 then
			model.prep, model.prepFor = prep, shown.title ~= nil and shown.title ~= "" and shown.title or shown.raid
		end
		model.signups = GuildRaid.signups(shown.roster)
		model.nextRaids = GuildRaid.nextRaids(raids, shown.raidId, ctx.pending)
	elseif type(ForeverSixtyInbox) == "table" and guildName ~= nil then
		model.empty = string.format(L.guildNoRaids, guildName)
	else
		model.empty = L.guildNoRaid
	end
	model.guild = GuildRaid.guildModel(guildName)
	model.sent = GuildHubTab.officersSent(Follow.pushedBuilds(Export.characterKey()), followsCode)
	return model
end

--- The header's figures (spec 3.2): the guild's name and the player's rank in it.
function GuildHubTab.figures()
	if type(GetGuildInfo) ~= "function" then
		return {}
	end
	local ok, name, rankName = pcall(GetGuildInfo, "player")
	if not ok then
		return {}
	end
	return { name, rankName }
end

--- The line Me's Right now shows for this tab (Hub's `lead`).
function GuildHubTab.lead()
	return GuildRaid.leadOf(GuildRaid.liveRaids(), GuildRaid.context())
end

--- What waits on this tab now (Hub's dot and the chip's line).
function GuildHubTab.waiting()
	return GuildRaid.waitingOf(GuildRaid.liveRaids(), GuildRaid.context())
end

-- Acting -------------------------------------------------------------------------------------------

--- A Next raids row: that raid takes Tonight's place.
function GuildHubTab.select(raidId)
	GuildHubTab.selected = raidId
	GuildHubTab.redraw()
end

--- Follow the build the officers sent (the Builds tab's own use, so every surface redraws).
function GuildHubTab.follow(sent)
	local BuildDialog = ns.BuildDialog or package.loaded.BuildDialog
	if BuildDialog ~= nil and sent ~= nil then
		BuildDialog.use(sent.code, sent.name)
		GuildHubTab.redraw()
	end
end

-- The frame ------------------------------------------------------------------------------------------

local textHeight, place, pooled, hit, cover = HubDraw.textHeight, HubDraw.place, HubDraw.pooled, HubDraw.hit,
	HubDraw.cover
local nameText, classesByName, nameClicks, itemClicks = HubDraw.nameText, HubDraw.classesByName, HubDraw.nameClicks,
	HubDraw.itemClicks
local heading, textLine = HubDraw.heading, HubDraw.textLine

--- An item's name in its quality colour, or nil until the client has it (it redraws the tab then).
local function itemText(itemId)
	return HubDraw.itemText(itemId, GuildHubTab.redraw)
end

local function drawAnswerButtons(view, parent, section, y)
	local G = GuildHubTab.LAYOUT
	local x = 0
	for index, answer in ipairs(section.buttons) do
		local button = view.buttons[index]
		if button == nil then
			button = Widgets.clientButton(parent, L.guildAnswers[answer], function(self)
				GuildRaid.answer(self.foreverSixtyRaid, self.foreverSixtyAnswer)
			end)
			view.buttons[index] = button
		end
		Widgets.resetButtonLabel(button, L.guildAnswers[answer])
		Widgets.fitButton(button, G.answerWidth)
		button.foreverSixtyRaid, button.foreverSixtyAnswer = section.raidId, answer
		Widgets.setSelected(button, answer == section.selected)
		button:ClearAllPoints()
		button:SetPoint("TOPLEFT", parent, "TOPLEFT", x, -y)
		button:Show()
		x = x + button:GetWidth() + G.answerGap
	end
	return y + Theme.SIZES.clientButtonHeight
end

local function drawJobs(view, parent, width, jobs, y, classes)
	for index, job in ipairs(jobs) do
		local key = "job" .. index
		local label = L.guildYourJob .. " " .. Theme.inlineColor("clientGold", job.label)
		if job.encounter ~= nil then
			label = label .. Theme.inlineColor("muted", SEP .. job.encounter)
		end
		y = textLine(view, parent, key, 0, y, width, label, "body")
		local purpose = job.purpose ~= nil and job.purpose ~= "" and job.purpose or nil
		if #job.with > 0 then
			local names = {}
			for _, name in ipairs(job.with) do
				names[#names + 1] = nameText(name, classes[name])
			end
			local with = string.format(L.guildWith, table.concat(names, ", "))
			purpose = purpose ~= nil and (purpose .. " " .. with) or with
		end
		if purpose ~= nil then
			y = textLine(view, parent, key .. ".p", 0, y, width, purpose, "muted", true)
		end
	end
	return y
end

--- Tonight (spec 4.5 left 1): the heading, the raid's icon, name and time, your sign-up and the
--- answers, your jobs. The y under it.
local function drawTonight(view, parent, width, section, y)
	local G = GuildHubTab.LAYOUT
	y = heading(view, parent, "tonight", width, y, section.heading, nil)
	local icon = view.raidIcon
	if icon == nil then
		icon = Theme.icon(parent, "ARTWORK", section.icon, G.raidIcon)
		view.raidIcon = icon
	end
	icon:SetTexture(section.icon)
	place(icon, parent, 0, y)
	local x = G.raidIcon + G.iconGap
	local name = view.texts.raid
	if name == nil then
		name = pooled(view, parent, "raid")
		Theme.setFace(name, Theme.FONT_ROLES.highlight, G.raidName)
	end
	view.used.raid = true
	name:SetTextColor(Theme.rgb(Theme.HEX.strong))
	name:SetWidth(width - x)
	name:SetText(section.raid)
	place(name, parent, x, y)
	if section.when ~= nil then
		textLine(view, parent, "when", x, y + G.line + 1, width - x, section.when, "muted")
	end
	y = y + G.raidIcon + G.buttonsGap
	local answer = section.answer
	local line = L.guildYourSignup .. " " .. Theme.inlineColor(answer.hex, answer.word)
	if answer.detail ~= nil then
		line = line .. " " .. Theme.inlineColor("muted", answer.detail)
	end
	y = textLine(view, parent, "signup", 0, y, width, line, "body")
	if section.buttons ~= nil then
		y = drawAnswerButtons(view, parent, section, y + 2) + G.buttonsGap
	elseif section.onSite ~= nil then
		y = textLine(view, parent, "onsite", 0, y, width, section.onSite, "muted")
	end
	return drawJobs(view, parent, width, section.jobs, y + 2, view.classes)
end

--- Your prep (spec 4.5 left 2): one row per need, its label muted and its value. The y under it.
local function drawPrep(view, parent, width, rows, prepFor, y)
	local G = GuildHubTab.LAYOUT
	y = heading(view, parent, "prep", width, y, L.guildPrep, string.format(L.guildFor, prepFor))
	local valueWidth = width - G.prepLabel
	for index, row in ipairs(rows) do
		local key = "prep" .. index
		local label, value, hexKey
		if row.kind == "door" or row.kind == "item" then
			label = row.kind == "door" and L.guildPrepDoor or RaidPrep.BOSS_NAMES[row.npc] or ""
			local named = itemText(row.itemId) or ""
			local where = row.state == "worn" and L.guildWorn or row.state == "bags" and L.guildInBags or L.guildNotInBags
			value = named .. " " .. Theme.inlineColor(row.state == "missing" and "warning" or "success", where)
		elseif row.kind == "resist" then
			label, value = L.guildResist[row.school], tostring(row.value)
		elseif row.kind == "carried" then
			local parts = {}
			for _, carried in ipairs(row.counts) do
				parts[#parts + 1] = string.format(L.guildCarriedCount, L.guildCarriedKinds[carried.kind] or carried.kind,
					carried.count)
			end
			label, value = L.guildCarried, table.concat(parts, SEP)
		elseif row.kind == "missing" then
			local names = {}
			for _, need in ipairs(row.needs) do
				names[#names + 1] = L.guildNeeds[need.kinds[1]] or need.kinds[1]
			end
			label, value, hexKey = L.guildMissing, table.concat(names, ", "), "warning"
		else
			label, value, hexKey = L.guildPackedLabel, L.guildPacked, "success"
		end
		textLine(view, parent, key .. ".l", 0, y, G.prepLabel, label or "", "muted")
		local x = G.prepLabel
		if row.itemId ~= nil then
			local icon = HubDraw.icon(view, parent, key, G.specIcon)
			icon:SetTexture(Compat.itemIcon(row.itemId) or Theme.UNKNOWN_ICON)
			place(icon, parent, x, y + (G.line - G.specIcon) / 2)
			x = x + G.specIcon + G.specGap
		end
		local next, region = textLine(view, parent, key .. ".v", x, y, width - x, value, hexKey or "body", true)
		if row.itemId ~= nil then
			local frame = hit(view, parent, key)
			itemClicks(frame, row.itemId)
			cover(frame, region, math.min(width - x, region:GetStringWidth() or valueWidth))
		end
		y = next
	end
	return y
end

--- Sign-ups (spec 4.5 left 3): the answers by kind in their colours, then the roles against the target.
local function drawSignups(view, parent, width, signups, y)
	y = heading(view, parent, "signups", width, y, L.guildSignups, signups.note)
	local parts = {}
	for _, count in ipairs(signups.counts) do
		parts[#parts + 1] = Theme.inlineColor(count.hex, count.text)
	end
	y = textLine(view, parent, "counts", 0, y, width, table.concat(parts, Theme.inlineColor("muted", SEP)), "body", true)
	if signups.roles ~= nil then
		local roles = {}
		for _, role in ipairs(signups.roles) do
			roles[#roles + 1] = Theme.inlineColor(role.hex, role.text)
		end
		y = textLine(view, parent, "roles", 0, y, width, table.concat(roles, Theme.inlineColor("muted", SEP)), "body", true)
	end
	return y
end

--- Next raids (spec 4.5 right 1): date, raid, your answer; a click puts the row's raid on top.
local function drawNextRaids(view, parent, width, rows, y)
	local G = GuildHubTab.LAYOUT
	y = heading(view, parent, "next", width, y, L.guildNext, nil)
	for index, row in ipairs(rows) do
		local key = "next" .. index
		textLine(view, parent, key .. ".d", 0, y, G.dateWidth, row.when, row.shown and "body" or "muted")
		textLine(view, parent, key .. ".r", G.dateWidth, y, G.raidWidth, row.raid, "strong")
		local answerWidth = width - G.dateWidth - G.raidWidth
		local _, region = textLine(view, parent, key .. ".a", G.dateWidth + G.raidWidth, y, answerWidth,
			row.answer.word, row.answer.hex)
		local frame = hit(view, parent, key)
		frame:ClearAllPoints()
		frame:SetPoint("TOPLEFT", parent, "TOPLEFT", 0, -y)
		frame:SetSize(width, G.line)
		local coming = row.coming
		frame:SetScript("OnEnter", coming ~= nil and function(self)
			Theme.showLines(self, { row.raid, Theme.inlineColor("muted", coming) }, "ANCHOR_RIGHT")
		end or nil)
		local raidId = row.raidId
		frame:SetScript("OnClick", function()
			GuildHubTab.select(raidId)
		end)
		y = y + math.max(G.line, textHeight(region))
	end
	return y
end

--- One standings row at `y`: the place, the spec icon and the class-coloured name, score, attendance.
local function drawStanding(view, parent, width, row, key, y)
	local G = GuildHubTab.LAYOUT
	local right = width - G.attendanceWidth - G.scoreWidth
	textLine(view, parent, key .. ".n", 0, y, G.rankWidth, tostring(row.place), "muted")
	-- The ringed class crest, every character's one image (as the site's rows).
	HubDraw.crest(view, parent, key, G.specIcon, row.classToken, G.rankWidth, y + (G.line - G.specIcon) / 2)
	local x = G.rankWidth + G.specIcon + G.specGap
	local _, name = textLine(view, parent, key .. ".m", x, y, right - x, nameText(row.short, row.classToken), "body")
	nameClicks(hit(view, parent, key), row.name)
	cover(view.hits[key], name, math.min(right - x, name:GetStringWidth() or 0))
	local score = pooled(view, parent, key .. ".s", G.scoreWidth)
	score:SetJustifyH("RIGHT")
	score:SetText(tostring(row.score))
	score:SetTextColor(Theme.rgb(Theme.panelTierHex(row.tier, Theme.HEX.body) or Theme.HEX.body))
	place(score, parent, right, y)
	local attendance = pooled(view, parent, key .. ".a", G.attendanceWidth)
	attendance:SetJustifyH("RIGHT")
	attendance:SetTextColor(Theme.rgb(Theme.HEX.muted))
	attendance:SetText(row.attendance ~= nil and string.format(L.guildPercent, row.attendance) or "")
	place(attendance, parent, right + G.scoreWidth, y)
	return y + G.line
end

--- The guild block (spec 4.5 right 2): its name, your place, its line, the top five and your row.
local function drawGuild(view, parent, width, guild, y)
	local G = GuildHubTab.LAYOUT
	y = heading(view, parent, "guild", width, y, guild.name, guild.note)
	if guild.line ~= nil then
		y = textLine(view, parent, "guildline", 0, y, width, guild.line, "muted", true) + 2
	end
	if #guild.top == 0 then
		return y
	end
	local right = width - G.attendanceWidth - G.scoreWidth
	textLine(view, parent, "col.n", 0, y, G.rankWidth, L.guildColumns.rank, "muted")
	textLine(view, parent, "col.m", G.rankWidth, y, right - G.rankWidth, L.guildColumns.name, "muted")
	for _, column in ipairs({ { "col.s", L.guildColumns.score, right, G.scoreWidth },
		{ "col.a", L.guildColumns.attendance, right + G.scoreWidth, G.attendanceWidth } }) do
		local _, region = textLine(view, parent, column[1], column[3], y, column[4], column[2], "muted")
		region:SetJustifyH("RIGHT")
	end
	y = y + G.line
	for index, row in ipairs(guild.top) do
		y = drawStanding(view, parent, width, row, "top" .. index, y)
	end
	if guild.own ~= nil then
		y = drawStanding(view, parent, width, guild.own, "own", y + G.ownGap)
	end
	return y
end

--- Officers sent (spec 4.5 right 3): the build's name, who sent it, and Follow or "you follow it".
local function drawSent(view, parent, width, sent, y)
	local G = GuildHubTab.LAYOUT
	y = heading(view, parent, "sent", width, y, L.guildOfficersSent, nil)
	local x = 0
	local spec = Follow.buildSpec(dataOf(), Codec.loadBuild(sent.code, dataOf()))
	local path = spec ~= nil and BuildsPane.specIcon(dataOf(), spec) or nil
	if path ~= nil then
		local icon = view.sentIcon
		if icon == nil then
			icon = Theme.icon(parent, "ARTWORK", path, G.buildIcon)
			view.sentIcon = icon
		end
		icon:SetTexture(path)
		place(icon, parent, 0, y + 2)
		x = G.buildIcon + G.iconGap
	elseif view.sentIcon ~= nil then
		view.sentIcon:Hide()
	end
	local follow = view.follow
	local room = width - x - (sent.following and 0 or G.followWidth + G.iconGap)
	local top = y
	y = textLine(view, parent, "sent.name", x, y, room, sent.name or "", "body")
	local from = {}
	if sent.from ~= nil and sent.from ~= "" then
		local short = sent.from:match("^([^%-]+)") or sent.from
		local who = nameText(short, view.classes[short])
		from[#from + 1] = sent.day ~= nil and string.format(L.guildFromOn, who, sent.day) or string.format(L.guildFrom, who)
	end
	if sent.following then
		from[#from + 1] = L.guildYouFollow
	end
	if #from > 0 then
		y = textLine(view, parent, "sent.from", x, y, room, table.concat(from, SEP), "muted")
	end
	if not sent.following then
		if follow == nil then
			follow = Widgets.clientButton(parent, L.guildFollow, function()
				GuildHubTab.follow(GuildHubTab.current ~= nil and GuildHubTab.current.sent or nil)
			end)
			view.follow = follow
		end
		Widgets.resetButtonLabel(follow, L.guildFollow)
		Widgets.fitButton(follow, G.followWidth)
		follow:ClearAllPoints()
		follow:SetPoint("TOPRIGHT", parent, "TOPRIGHT", 0, -top)
		follow:Show()
	elseif follow ~= nil then
		follow:Hide()
	end
	return math.max(y, top + Theme.SIZES.clientButtonHeight)
end

--- The left column: Tonight (or the empty line), Your prep, Sign-ups. Its height.
local function drawLeft(view, model)
	local parent, width = view.left, Theme.SIZES.windowLeft
	local G = GuildHubTab.LAYOUT
	local y = 0
	for _, button in ipairs(view.buttons) do
		button:Hide()
	end
	if model.tonight == nil then
		if view.raidIcon ~= nil then
			view.raidIcon:Hide()
		end
		y = heading(view, parent, "tonight", width, y, L.guildTonight, nil)
		return textLine(view, parent, "empty", 0, y, width, model.empty or L.guildNoRaid, "muted", true)
	end
	y = drawTonight(view, parent, width, model.tonight, y)
	if model.prep ~= nil then
		y = drawPrep(view, parent, width, model.prep, model.prepFor, y + G.sectionGap)
	end
	if model.signups ~= nil then
		y = drawSignups(view, parent, width, model.signups, y + G.sectionGap)
	end
	return y
end

--- The right column: Next raids, the guild, Officers sent. Its height.
local function drawRight(view, model)
	local parent, width = view.right, Theme.SIZES.windowRight
	local G = GuildHubTab.LAYOUT
	local y, any = 0, false
	local function gap()
		local spaced = any and y + G.sectionGap or y
		any = true
		return spaced
	end
	if model.nextRaids ~= nil then
		y = drawNextRaids(view, parent, width, model.nextRaids, gap())
	end
	if model.guild ~= nil then
		y = drawGuild(view, parent, width, model.guild, gap())
	end
	if model.sent ~= nil then
		y = drawSent(view, parent, width, model.sent, gap())
	elseif view.follow ~= nil then
		view.follow:Hide()
	end
	if view.sentIcon ~= nil and model.sent == nil then
		view.sentIcon:Hide()
	end
	Theme.setShown({ view.rule }, any)
	return y
end

--- Redraw the tab from a fresh model: both columns, the footer, the window's height.
function GuildHubTab.draw()
	local view = GuildHubTab.view
	if view == nil then
		return nil
	end
	local model = GuildHubTab.model(GuildHubTab.selected)
	GuildHubTab.current = model
	HubDraw.begin(view)
	view.classes = classesByName()
	local left = drawLeft(view, model)
	local right = drawRight(view, model)
	HubDraw.sweep(view)
	view.rule:SetHeight(math.max(left, right, 1))
	local pending = model.tonight ~= nil and model.tonight.pending
	Hub.setHints(nil, pending and L.guildSendsLater or nil)
	Hub.fitHeight(math.max(left, right))
	Hub.refresh()
	return model
end

--- The tab redrawn while it shows (an answer queued, a raid picked, the roster or the bags changed).
function GuildHubTab.redraw()
	if GuildHubTab.view ~= nil and GuildHubTab.view.shown then
		GuildHubTab.draw()
	end
end

--- The tab's body, once, into the hub's content frame: the two columns and the rule between them.
function GuildHubTab.build(content)
	GuildHubTab.view = HubDraw.view(content)
	return GuildHubTab.view
end

function GuildHubTab.show()
	local view = GuildHubTab.view
	if view == nil then
		return nil
	end
	view.shown = true
	Theme.requestGuildRoster()
	return GuildHubTab.draw()
end

function GuildHubTab.hide()
	if GuildHubTab.view ~= nil then
		GuildHubTab.view.shown = false
	end
	GuildHubTab.selected = nil
	Theme.hideTooltip()
end

--- The events that change what the tab shows while it is open: the roster, the bags, the gear.
GuildHubTab.EVENTS = { "GUILD_ROSTER_UPDATE", "BAG_UPDATE_DELAYED", "PLAYER_EQUIPMENT_CHANGED", "UNIT_RESISTANCES" }

--- Runs at file load: one hidden event frame, where the client has frames.
function GuildHubTab.register()
	if type(CreateFrame) ~= "function" then
		return nil
	end
	local frame = CreateFrame("Frame", nil, UIParent)
	frame:SetScript("OnEvent", function()
		GuildHubTab.redraw()
	end)
	for _, event in ipairs(GuildHubTab.EVENTS) do
		Theme.registerEvent(frame, event)
	end
	return frame
end

GuildHubTab.frame = GuildHubTab.register()
GuildRaid.onAnswer = function()
	GuildHubTab.redraw()
end

Hub.register({
	id = "guild",
	label = L.hubTabGuild,
	build = GuildHubTab.build,
	show = GuildHubTab.show,
	hide = GuildHubTab.hide,
	figures = GuildHubTab.figures,
	waiting = GuildHubTab.waiting,
	lead = GuildHubTab.lead,
})

ns.GuildHubTab = GuildHubTab
return GuildHubTab
