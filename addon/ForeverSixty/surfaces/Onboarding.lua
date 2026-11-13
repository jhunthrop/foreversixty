-- addon/ForeverSixty/surfaces/Onboarding.lua
-- The first-run tour (design/specs/2026-11-08-addon-ux.md section 3.1): once per account, a
-- 640 x 450 window in the panel's chrome (umber ground, title band with the class medallion, the
-- client's red buttons) that shows how the site, the addon, the companion and the apps work as one
-- loop, then sets the character up. Five pages in a guild, four without (page 4 is the guild's).
-- It replaces the three chat lines of the first login.
--
-- Pure model, thin frame: Onboarding.pages and Onboarding.setup are plain tables in, plain
-- tables out; Onboarding.draw lays a page out. The frame is built on the first show, never at
-- load. It never shows in combat or inside an instance; combat hides it without ending it, and
-- it comes back once after combat on the page it was on. Skip tour, Tooltips only, the X,
-- Escape and Done end it for good (ForeverSixtyDB.tourSeen, this tour's version; a changed tour
-- shows once more); /fs tour and /fs help open it again.
--
-- Truth rule: every claim is what is built today. What is designed but not live carries a gold
-- "coming" tag (answering in game, the phone app); the app QR is small and labelled "Join the
-- app beta" until the apps ship (Onboarding.QR_SIZE makes it large then).
--
-- Spec section 6 (2026-11-11): every page opens with one lead saying what the player gets; the
-- last page asks the spec question when the points do not name a tree (Follow.specQuestion);
-- and the tour steps aside while the Builds dialog is open, coming back on the same page.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Compat = ns.Compat or require("Compat")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Prefs = ns.Prefs or require("Prefs")
local Progress = ns.Progress or require("Progress")
local Talents = ns.Talents or require("Talents")
local Chip = ns.Chip or require("Chip")
local TourExample = ns.TourExample or require("TourExample")
local SiteStatus = ns.SiteStatus or require("SiteStatus")
local SettingsPage = ns.SettingsPage or require("Settings")
local BuildDialog = ns.BuildDialog or require("BuildDialog")
local Hub = ns.Hub or require("Hub")

local Onboarding = {}

Onboarding.FRAME_NAME = "ForeverSixtyOnboarding"
Onboarding.WIDTH = 640
--- The tallest a page may grow; each page is sized to its content (Onboarding.frameHeight), never
--- shorter than MIN_HEIGHT.
Onboarding.HEIGHT = 450
Onboarding.MIN_HEIGHT = 240
--- The content's side inset inside the umber edge.
Onboarding.PAD = 26
--- The app QR's size in UI units. Small and labelled "Join the app beta" while the apps are in
--- testing (the coordinator's ruling, 2026-11-08); 112 once they are in the stores.
Onboarding.QR_SIZE = 48
Onboarding.QR_TEXTURE = "Interface\\AddOns\\ForeverSixty\\media\\app-qr"
--- The footer's and the set-up page's buttons: the client's red button, a step taller than the
--- panel's small ones.
Onboarding.BUTTON_HEIGHT = 22
--- The gap between a page's lead and its content, and between the set-up page's tiles.
Onboarding.LEAD_GAP = 12
Onboarding.TILE_GAP = 8
--- The addresses the set-up page offers to copy.
Onboarding.SITE = "foreversixty.gg"
Onboarding.SETUP = "foreversixty.gg/setup"
Onboarding.APP = "foreversixty.gg/app"
--- The faction banners for the guild node, by Compat.playerFaction's answer.
Onboarding.BANNERS = {
	horde = "Interface\\Icons\\INV_BannerPVP_01", alliance = "Interface\\Icons\\INV_BannerPVP_02",
}
Onboarding.ICONS = {
	site = "Interface\\AddOns\\ForeverSixty\\icon",
	game = "Interface\\Icons\\INV_Misc_Bag_08",
	fights = "Interface\\Icons\\Ability_Warrior_BattleShout",
	back = "Interface\\Icons\\INV_Misc_Book_09",
}

-- The pure model ----------------------------------------------------------------------------------

--- The page keys in order for a character `inGuild` or not: the loop, the site in game, the game
--- back to the site, the guild (only in one), set it up. Pure.
function Onboarding.pageKeys(inGuild)
	if inGuild then
		return { "loop", "site", "back", "guild", "setup" }
	end
	return { "loop", "site", "back", "setup" }
end

--- Page 1's nodes, in reading order (top row left to right, then bottom row right to left), each
--- { icon, title, line }; without a guild the guild node goes. Pure.
function Onboarding.loopNodes(inGuild, faction, classToken)
	local nodes = {
		{ icon = Theme.crestPath(classToken) or Theme.UNKNOWN_ICON, title = L.tourCharacter, line = L.tourCharacterLine },
		{ icon = Onboarding.ICONS.site, title = L.tourSite, line = L.tourSiteLine },
		{ icon = Onboarding.ICONS.game, title = L.tourGame, line = L.tourGameLine },
		{ icon = Onboarding.ICONS.fights, title = L.tourFights, line = L.tourFightsLine },
		{ icon = Onboarding.ICONS.back, title = L.tourBack, line = L.tourBackLine },
	}
	if inGuild then
		nodes[#nodes + 1] = { icon = Onboarding.BANNERS[faction] or Onboarding.BANNERS.alliance, title = L.tourGuild,
			line = L.tourGuildLine }
	end
	return nodes
end

--- The set-up page (spec 3.1, last page): the path (the spec tiles, or the build followed), the two
--- connect rows or the companion's last-seen line, and the QR. `input`: { level, build, buildName,
--- progress, pointName, specs, siteAge (seconds since the site last saw the character, nil for
--- never), name }. Pure.
function Onboarding.setup(input)
	local page = { kind = "setup" }
	if input.build ~= nil then
		page.path = { name = string.format(L.tourYourPath, input.buildName or ""), icon = input.pathIcon }
		local progress = input.progress
		if type(progress) == "table" then
			local allowed = Progress.pointsAllowed(input.build, input.level)
			if allowed > 0 then
				page.path.count = string.format(L.panelOnPlan, math.min(progress.matched or 0, allowed), allowed)
			end
			if progress.nextLevel ~= nil and input.pointName ~= nil then
				page.path.next = string.format(L.tourNextPoint, progress.nextLevel, input.pointName)
			end
		end
	else
		page.tiles = input.specs or {}
		-- The spec question when the points do not name a tree (spec 6.2), else the plain heading, with
		-- the points split beside it ("2 Holy · 2 Retribution").
		page.ask = input.question or L.panelPickPath
		page.split = input.split
	end
	local age = input.siteAge
	if type(age) == "number" and age <= SiteStatus.FRESH_SECONDS then
		page.seen = { ok = true, text = L.tourCompanionRan,
			detail = string.format(L.tourSiteSaw, input.name or "", Onboarding.ago(age)) }
	elseif type(age) == "number" and age > SiteStatus.SILENT_SECONDS then
		page.seen = { ok = false, text = string.format(L.tourSiteSilent, input.name or "", math.floor(age / 86400)) }
		page.companion = true
	else
		page.signIn, page.companion = true, true
	end
	if age == nil then
		page.signIn, page.companion = true, true
	end
	return page
end

--- A spec tile's width with `count` tiles across the content (spec 6.2): 188 for three, 141 for
--- the druid's four. Pure.
function Onboarding.tileWidth(count)
	count = math.max(1, count or 1)
	return math.floor((Onboarding.WIDTH - 2 * Onboarding.PAD - Onboarding.TILE_GAP * (count - 1)) / count)
end

--- A tile's line under its button: the spec's playstyle, then the talent the character would take
--- next on its build ("Next: Divine Strength"). Pure.
function Onboarding.tileLine(spec)
	local parts = {}
	if type(spec.playstyle) == "string" and spec.playstyle ~= "" then
		parts[#parts + 1] = spec.playstyle
	end
	if type(spec.next) == "string" then
		parts[#parts + 1] = string.format(L.pathNext, spec.next)
	end
	return table.concat(parts, L.exportTreeSeparator)
end

--- The lead every page opens with (spec 6.3): what the player gets, in one or two lines. Pure.
Onboarding.LEADS = { loop = "tourLoopLead", site = "tourSiteLead", back = "tourBackLead", guild = "tourGuildLead",
	setup = "tourSetupLead" }
function Onboarding.lead(key)
	return L[Onboarding.LEADS[key]]
end

--- "2 hours ago", "25 minutes ago", "just now". Pure.
function Onboarding.ago(seconds)
	if seconds >= 3600 then
		local hours = math.floor(seconds / 3600)
		return string.format(hours == 1 and L.tourHourAgo or L.tourHoursAgo, hours)
	elseif seconds >= 60 then
		return string.format(L.tourMinutesAgo, math.floor(seconds / 60))
	end
	return L.panelJustNow
end

-- When it shows --------------------------------------------------------------------------------------

--- The tour's version. Bump it when the tour changes and every account sees the new tour once:
--- ForeverSixtyDB.tourSeen holds the version the player ended (owner, 2026-11-13: deleting the
--- addon folder keeps SavedVariables, so a stamp alone never showed a reinstall the tour). 1 was
--- 2.0.38's tour, stamped as welcomedAt alone; 2 is the tour with its leads and the spec question.
Onboarding.TOUR_VERSION = 2

--- Whether the tour is due this login: this version not ended on this account (a first install, a
--- reinstall that wiped SavedVariables, one that kept an older tour's stamp, or a tour that changed
--- since) and not quieted by "Tooltips only". Pure on `db` and `tooltipsOnly`.
function Onboarding.due(db, tooltipsOnly)
	if tooltipsOnly then
		return false
	end
	return type(db) ~= "table" or db.tourSeen ~= Onboarding.TOUR_VERSION
end

--- Stamps the account: this tour version ended, and the first welcome's time (the chip's first
--- logins read it) kept from the first time it was set.
function Onboarding.stamp()
	ForeverSixtyDB = ForeverSixtyDB or {}
	ForeverSixtyDB.tourSeen = Onboarding.TOUR_VERSION
	ForeverSixtyDB.welcomedAt = ForeverSixtyDB.welcomedAt or (type(time) == "function" and time() or 0)
end

--- Ends the tour for good: the account is stamped and the frame closes.
function Onboarding.finish()
	Onboarding.stamp()
	Onboarding.pendingAfterCombat = nil
	if Onboarding.frame ~= nil then
		Onboarding.frame:Hide()
	end
end

--- The Builds dialog opens over the tour (Change, /fs build, the panel): the tour steps aside,
--- hidden without ending, and remembers its page (spec 6.1). False when it was not showing.
function Onboarding.stepAside()
	if Onboarding.frame == nil or not Onboarding.frame:IsShown() then
		return false
	end
	Onboarding.asidePage = Onboarding.page or 1
	Onboarding.steppingAside = true
	Onboarding.frame:Hide()
	Onboarding.steppingAside = nil
	return true
end

--- The dialog closed: the tour comes back on the page it left, read afresh, so a build picked in
--- the dialog shows as the new path. False when the tour had not stepped aside.
function Onboarding.stepBack()
	local page = Onboarding.asidePage
	if page == nil then
		return false
	end
	Onboarding.asidePage = nil
	return Onboarding.open(page)
end

--- "Tooltips only" (page 1): the setting on (the chip, the talent marks and the slot marks off), and
--- the tour ends.
function Onboarding.justTooltips()
	SettingsPage.onToggle("tooltipsOnly", true)
	Onboarding.finish()
end

--- "Open my panel": the tour ends and the panel opens.
function Onboarding.openPanel()
	Onboarding.finish()
	local Panel = ns.Panel
	if Panel ~= nil then
		Panel.open("key")
	end
end

-- Reading the client -----------------------------------------------------------------------------------

--- Everything the pages read, from the client.
function Onboarding.inputs(data)
	data = data or Onboarding.data or ns.Data
	local guild = Export.guildInfo()
	local build = Follow.build
	local level = Chip.level()
	local ranks = data ~= nil and Talents.readRanks(data) or {}
	local progress = build ~= nil and Progress.build(build, ranks, level) or nil
	local input = {
		inGuild = guild ~= nil and type(guild.name) == "string" and guild.name ~= "",
		faction = Compat.playerFaction(), classToken = Chip.classToken(), level = level, name = Compat.playerName(),
		build = build, buildName = build ~= nil and Follow.buildName(data, build) or nil, progress = progress,
		specs = Chip.pathSpecs(data, Talents.playerClassSlug(), ranks, level),
		question = Follow.specQuestion(data, Talents.playerClassSlug(), ranks),
		split = Follow.specSplit(data, Talents.playerClassSlug(), ranks),
	}
	local ok, facts = pcall(SiteStatus.facts)
	input.siteAge = ok and type(facts) == "table" and facts.age or nil
	if progress ~= nil and progress.point ~= nil then
		local point = progress.point
		input.pointName = Chip.talentName(Talents.cellOf(data, build.classSlug, point.tab, point.tier, point.column))
	end
	input.upgrade, input.pick = TourExample.inputs(data, input)
	return input
end

-- The frame ---------------------------------------------------------------------------------------------

local pools = { text = {}, icon = {}, button = {}, box = {}, rule = {}, card = {} }
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

local function text(parent, value, hexKey, fontKey, width)
	local region = take("text", function()
		return Widgets.label(parent, "", "body", "small")
	end)
	Theme.applyFont(region, fontKey or "small")
	region:SetText(value or "")
	region:SetTextColor(Theme.rgb(Theme.HEX[hexKey or "body"]))
	region:SetJustifyH("LEFT")
	region:SetWidth(width or 0)
	if type(region.SetWordWrap) == "function" then
		region:SetWordWrap(width ~= nil)
	end
	return region
end

local function height(region, floor)
	local measured = region:GetStringHeight()
	return math.max(floor or 12, type(measured) == "number" and math.ceil(measured) or 12)
end

local function icon(parent, path, size)
	local texture = take("icon", function()
		return Theme.icon(parent, "ARTWORK", nil, size)
	end)
	texture:SetSize(size, size)
	texture:SetTexture(path or Theme.UNKNOWN_ICON)
	-- The client's icons lose their baked bevel (Theme.ICON_CROP); the addon's own art is drawn whole.
	local own = type(path) == "string" and path:find("Interface\\AddOns\\", 1, true) == 1
	local crop = own and { 0, 1, 0, 1 } or Theme.ICON_CROP
	texture:SetTexCoord(crop[1], crop[2], crop[3], crop[4])
	return texture
end

--- The loop's flow arrow, drawn from flat brass rectangles (the slot marks' way: no font glyph,
--- no client art): a shaft and a three-step head, { x, y, width, height } in a 16 x 10 box
--- pointing right. Left mirrors it; down turns it.
Onboarding.ARROW_PARTS = { { 0, 4, 10, 2 }, { 10, 0, 2, 10 }, { 12, 2, 2, 6 }, { 14, 4, 2, 2 } }
Onboarding.ARROW_BOX = { 16, 10 }

--- The parts of an arrow pointing `direction` ("right", "left" or "down"). Pure.
function Onboarding.arrowParts(direction)
	local parts = {}
	for index, part in ipairs(Onboarding.ARROW_PARTS) do
		local x, y, width, h = part[1], part[2], part[3], part[4]
		if direction == "left" then
			x = Onboarding.ARROW_BOX[1] - x - width
		elseif direction == "down" then
			x, y, width, h = y, x, h, width
		end
		parts[index] = { x, y, width, h }
	end
	return parts
end

local function arrow(parent, direction, left, top)
	for _, part in ipairs(Onboarding.arrowParts(direction)) do
		local piece = take("rule", function()
			return Theme.texture(parent, "ARTWORK", "brass")
		end)
		piece:SetSize(part[3], part[4])
		piece:SetPoint("TOPLEFT", parent, "TOPLEFT", left + part[1], -top - part[2])
	end
end

local function button(parent, label, onClick, hint, primary)
	local made = take("button", function()
		-- The panel's button in the site's recipes (Widgets.clientButton), a little taller here.
		local built = Widgets.clientButton(parent, "", function(self)
			if self.foreverSixtyAction ~= nil then
				self.foreverSixtyAction()
			end
		end)
		built:SetHeight(Onboarding.BUTTON_HEIGHT)
		-- The hover a use gives it (Tooltips only's), none for the rest.
		built:SetScript("OnEnter", function(self)
			if self.foreverSixtyHint ~= nil then
				Theme.showLines(self, self.foreverSixtyHint, "ANCHOR_TOP")
			end
		end)
		built:SetScript("OnLeave", function()
			Theme.hideTooltip()
		end)
		return built
	end)
	made:SetText(label)
	made.foreverSixtyAction = onClick
	made.foreverSixtyHint = hint
	-- The page's one forward action (Next, Done) is the site's primary; the rest its secondary.
	Widgets.setSelected(made, primary == true)
	return Widgets.fitButton(made, 64, 24)
end

--- Page 2's example in the client's tooltip look (ground and edge as GameTooltip draws them): the
--- item's name in its quality colour, then the lines Forever Sixty adds. Returns its height.
local function exampleCard(parent, example, left, top)
	local card = take("card", function()
		local frame = CreateFrame("Frame", nil, parent)
		Theme.texture(frame, "BACKGROUND", "tooltipGround", 0.94):SetAllPoints(frame)
		Theme.outline(frame, Theme.SIZES.border, "tooltipEdge")
		frame.foreverSixtyLines = {}
		for index = 1, 3 do
			frame.foreverSixtyLines[index] = Widgets.label(frame, "", "tooltipBody", index == 1 and "normal" or "small")
		end
		return frame
	end)
	local lines = { { "|cff" .. example.hex .. example.name .. "|r", "tooltipBody" } }
	for _, line in ipairs(example.lines) do
		if line[1] ~= nil then
			lines[#lines + 1] = line
		end
	end
	local y, width = 8, 0
	for index, region in ipairs(card.foreverSixtyLines) do
		local line = lines[index]
		region:SetShown(line ~= nil)
		if line ~= nil then
			region:SetText(line[1])
			region:SetTextColor(Theme.rgb(Theme.HEX[line[2]]))
			region:ClearAllPoints()
			region:SetPoint("TOPLEFT", card, "TOPLEFT", 10, -y)
			local measured = region:GetStringWidth()
			width = math.max(width, type(measured) == "number" and measured or 160)
			y = y + (index == 1 and 18 or 15)
		end
	end
	card:SetSize(math.ceil(width) + 20, y + 7)
	card:SetPoint("TOPLEFT", parent, "TOPLEFT", left, -top)
	return y + 7
end

--- A read-only address box the player clicks to select, then Ctrl+C.
local function addressBox(parent, value)
	local field = take("box", function()
		local box = Widgets.editBox(parent, 200, 22, true, true)
		-- A click selects the whole address for Ctrl+C; nothing takes the keyboard on its own.
		box:SetScript("OnEditFocusGained", function(self)
			self:HighlightText()
		end)
		box.foreverSixtyField.foreverSixtyBox = box
		return box.foreverSixtyField
	end)
	Widgets.setText(field.foreverSixtyBox, value)
	return field
end

local function coming(parent, region)
	-- Muted, not gold (round 3): gold on page 4 is the job's; a status tag is quiet. The dot in
	-- L.tourComing parts it from the line it tags (2026-11-13: a bare word read as the line's own).
	local tag = text(parent, L.tourComing, "umberMuted")
	tag:SetPoint("LEFT", region, "RIGHT", 4, 0)
	return tag
end

local function drawLoop(content, input)
	local nodes = Onboarding.loopNodes(input.inGuild, input.faction, input.classToken)
	local columnX = { 0, 196, 392 }
	local rows = { { nodes[1], nodes[2], nodes[3] }, { nodes[6], nodes[5], nodes[4] } }
	local top = { 0, 140 }
	local bottom = 0
	for rowIndex, row in ipairs(rows) do
		for column, node in ipairs(row) do
			if node ~= nil then
				local mark = icon(content, node.icon, 34)
				mark:SetPoint("TOPLEFT", content, "TOPLEFT", columnX[column], -top[rowIndex])
				local title = text(content, node.title, "ink", "normal")
				title:SetPoint("TOPLEFT", content, "TOPLEFT", columnX[column] + 40, -top[rowIndex])
				local line = text(content, node.line, "umberMuted", "small", 130)
				line:SetPoint("TOPLEFT", content, "TOPLEFT", columnX[column] + 40, -top[rowIndex] - 18)
				bottom = math.max(bottom, top[rowIndex] + 18 + height(line))
			end
			if column < 3 and (rowIndex == 1 or row[column] ~= nil) then
				arrow(content, rowIndex == 1 and "right" or "left", columnX[column] + 174, top[rowIndex] + 12)
			end
		end
	end
	arrow(content, "down", columnX[3] + 12, top[1] + 96)
	return bottom
end

local function drawRows(content, rows)
	local y = 0
	for _, row in ipairs(rows) do
		local x = 0
		if row.source ~= nil then
			local source = text(content, row.source, "umberMuted", "small", 160)
			source:SetJustifyH("RIGHT")
			source:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
			x = 176
		end
		local mark = icon(content, row.icon, 30)
		mark:SetPoint("TOPLEFT", content, "TOPLEFT", x, -y)
		local title = text(content, row.title, "ink", "normal")
		title:SetPoint("TOPLEFT", content, "TOPLEFT", x + 40, -y)
		local line = text(content, row.line, "umberMuted", "small", Onboarding.WIDTH - 100 - x)
		line:SetPoint("TOPLEFT", content, "TOPLEFT", x + 40, -y - 18)
		local below = 22 + height(line)
		if row.example ~= nil then
			below = below + 6 + exampleCard(content, row.example, x + 40, y + below + 6)
		end
		y = y + math.max(44, below) + 6
	end
	return y - 6
end

local function drawSite(content, input)
	return drawRows(content, {
		{ source = L.tourSiteBis, icon = Theme.MEDIA.arrowInline, title = L.tourArrow, line = L.tourArrowLine },
		{ source = L.tourSiteWeights, icon = "Interface\\Icons\\INV_Weapon_Bow_02", title = L.tourOneLine,
			line = L.tourOneLineText, example = TourExample.example(input) },
		{ source = L.tourSiteBuild, icon = "Interface\\Icons\\Ability_SearingArrow", title = L.tourRing,
			line = L.tourRingLine },
		{ source = L.tourSiteRotation, icon = "Interface\\Icons\\INV_Spear_07", title = L.tourRotation,
			line = L.tourRotationLine },
	})
end

local function drawBack(content)
	return drawRows(content, {
		{ icon = Onboarding.ICONS.fights, title = L.tourUpload, line = L.tourUploadLine },
		{ icon = Onboarding.ICONS.back, title = L.tourRated, line = L.tourRatedLine },
		{ icon = Onboarding.ICONS.site, title = L.tourScore, line = L.tourScoreLine .. " " .. L.tourScoreHide },
	})
end

local function drawGuild(content)
	local y = 0
	local chipLine = text(content, string.format(L.chipTonight, Compat.clockText(20, 0), L.tourGuildRaid), "ink",
		"normal")
	chipLine:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
	local chipNote = text(content, L.tourGuildChipNote, "umberMuted")
	chipNote:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y - 20)
	local job = text(content, string.format(L.panelJob, L.tourGuildJob), "ink")
	job:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y - 46)
	local jobNote = text(content, L.tourGuildJobNote, "umberMuted")
	jobNote:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y - 62)
	local answer = text(content, L.tourAnswer, "ink", "normal")
	answer:SetPoint("TOPLEFT", content, "TOPLEFT", 320, -y)
	local site = text(content, L.tourAnswerSite, "ink")
	site:SetPoint("TOPLEFT", content, "TOPLEFT", 320, -y - 20)
	local discord = text(content, L.tourAnswerDiscord, "ink")
	discord:SetPoint("TOPLEFT", content, "TOPLEFT", 320, -y - 36)
	local later = text(content, L.tourAnswerLater, "umberMuted")
	later:SetPoint("TOPLEFT", content, "TOPLEFT", 320, -y - 52)
	coming(content, later)
	return y + 62 + height(jobNote)
end

local function drawSetup(content, input)
	local page = Onboarding.setup(input)
	local y = 0
	if page.path ~= nil then
		local name = text(content, page.path.name, "ink", "normal")
		name:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
		if page.path.count ~= nil then
			local count = text(content, page.path.count, "umberMuted")
			count:SetPoint("LEFT", name, "RIGHT", 10, 0)
		end
		if page.path.next ~= nil then
			local nextLine = text(content, page.path.next, "umberMuted")
			nextLine:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y - 18)
		end
		-- The Builds dialog (its open steps the tour aside; its close brings it back here).
		local change = button(content, L.panelChange, function()
			BuildDialog.open("paste")
		end)
		change:SetPoint("TOPRIGHT", content, "TOPRIGHT", 0, -y)
		y = y + 44
	else
		local title = text(content, page.ask, "ink", "normal")
		title:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
		if page.split ~= nil then
			local note = text(content, page.split, "umberMuted")
			note:SetPoint("LEFT", title, "RIGHT", 8, 0)
		end
		y = y + 20
		local tileWidth = Onboarding.tileWidth(#page.tiles)
		local tallest = 0
		for index, spec in ipairs(page.tiles) do
			local x = (index - 1) * (tileWidth + Onboarding.TILE_GAP)
			local pick = button(content, spec.name, function()
				Chip.choosePath(spec.slug)
				Onboarding.draw()
			end)
			pick:SetPoint("TOPLEFT", content, "TOPLEFT", x, -y)
			pick:SetWidth(tileWidth)
			local line = text(content, Onboarding.tileLine(spec), "umberMuted", "small", tileWidth)
			line:SetPoint("TOPLEFT", content, "TOPLEFT", x, -y - 28)
			tallest = math.max(tallest, height(line))
		end
		y = y + 28 + tallest + 2
	end
	local left = Onboarding.WIDTH - 60 - Onboarding.QR_SIZE - 120
	if page.seen ~= nil then
		local seen = text(content, page.seen.text, page.seen.ok and "tooltipGreen" or "warning", "normal", left)
		seen:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
		y = y + height(seen) + 2
		if page.seen.detail ~= nil then
			local detail = text(content, page.seen.detail, "umberMuted", "small", left)
			detail:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
			y = y + height(detail) + 8
		end
	end
	local rows = {}
	if page.signIn then
		rows[#rows + 1] = { L.tourSignIn, L.tourSignInLine, Onboarding.SITE }
	end
	if page.companion then
		rows[#rows + 1] = { L.tourCompanion, L.tourCompanionLine, Onboarding.SETUP }
	end
	for _, row in ipairs(rows) do
		local title = text(content, row[1], "ink", "normal")
		title:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
		local line = text(content, row[2], "umberMuted", "small", left)
		line:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y - 18)
		local box = addressBox(content, row[3])
		box:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y - 22 - height(line))
		-- 46: the lead above the page (spec 6.3) took the slack the rows had at 50.
		y = y + 46 + height(line)
	end
	if #rows > 0 then
		local links = text(content, L.tourLinks, "umberMuted", "small", left)
		links:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
		y = y + height(links)
	end
	-- "/fs help opens this again." sits in the footer's empty left on this page (spec 6.3).
	local qr = icon(content, Onboarding.QR_TEXTURE, Onboarding.QR_SIZE)
	qr:SetPoint("TOPRIGHT", content, "TOPRIGHT", -40, -(page.path ~= nil and 54 or 92))
	local label = text(content, L.tourAppBeta, "ink", "small")
	label:SetPoint("TOP", qr, "BOTTOM", 0, -4)
	local address = text(content, Onboarding.APP, "umberMuted", "small")
	address:SetPoint("TOP", label, "BOTTOM", 0, -2)
	return math.max(y, (page.path ~= nil and 54 or 92) + Onboarding.QR_SIZE + 36)
end

local DRAW = { loop = drawLoop, site = drawSite, back = drawBack, guild = drawGuild, setup = drawSetup }
local HEADINGS = { loop = "tourLoopHeading", site = "tourSiteHeading", back = "tourBackHeading",
	guild = "tourGuildHeading", setup = "tourSetupHeading" }

--- The window, in the panel's chrome (Widgets.umberFrame and Widgets.titleBand): the umber ground,
--- the title band with the class medallion, the page's heading in bone over "Welcome to Forever
--- Sixty", and "2 of 5" on the right; the client's red close X; the footer. The band drags it.
--- Built once.
function Onboarding.ensure()
	if Onboarding.frame ~= nil then
		return Onboarding.frame
	end
	local frame = Widgets.umberFrame(UIParent, Onboarding.WIDTH, Onboarding.HEIGHT, Onboarding.FRAME_NAME)
	frame:SetFrameStrata("DIALOG")
	frame:SetPoint("TOP", UIParent, "CENTER", 0, Onboarding.HEIGHT / 2)
	frame:SetClampedToScreen(true)
	frame:SetMovable(true)
	frame:EnableMouse(true)
	frame:Hide()
	Onboarding.band = Widgets.titleBand(frame)
	local close = Widgets.clientCloseButton(frame, Onboarding.finish)
	close:SetPoint("TOPRIGHT", frame, "TOPRIGHT", -2, -2)
	local content = CreateFrame("Frame", nil, frame)
	content:SetPoint("TOPLEFT", frame, "TOPLEFT", Onboarding.PAD, -(Theme.SIZES.titleBand + 16))
	content:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", -Onboarding.PAD, 50)
	Onboarding.content = content
	-- The page's lead (spec 6.3) at the content's top; the page's own parts draw in the body under it.
	local lead = Widgets.label(content, "", "ink", "normal")
	lead:SetPoint("TOPLEFT", content, "TOPLEFT", 0, 0)
	lead:SetWidth(Onboarding.WIDTH - 2 * Onboarding.PAD)
	if type(lead.SetWordWrap) == "function" then
		lead:SetWordWrap(true)
	end
	Onboarding.leadText = lead
	local body = CreateFrame("Frame", nil, content)
	body:SetPoint("BOTTOMRIGHT", content, "BOTTOMRIGHT", 0, 0)
	Onboarding.body = body
	local footer = CreateFrame("Frame", nil, frame)
	footer:SetPoint("BOTTOMLEFT", frame, "BOTTOMLEFT", 16, 14)
	footer:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", -16, 14)
	footer:SetHeight(Onboarding.BUTTON_HEIGHT)
	Onboarding.footer = footer
	frame:SetScript("OnKeyDown", function(self, key)
		if key == "ENTER" then
			self:SetPropagateKeyboardInput(false)
			Onboarding.next()
		else
			self:SetPropagateKeyboardInput(true)
		end
	end)
	if type(frame.EnableKeyboard) == "function" then
		frame:EnableKeyboard(true)
	end
	if type(UISpecialFrames) == "table" then
		table.insert(UISpecialFrames, Onboarding.FRAME_NAME)
	end
	-- Escape (UISpecialFrames) hides the frame: that ends the tour for good, unless combat hid it or
	-- it stepped aside for the Builds dialog.
	frame:SetScript("OnHide", function()
		if not Onboarding.hidingForCombat and not Onboarding.steppingAside then
			Onboarding.stamp()
		end
	end)
	Onboarding.frame = frame
	return frame
end

--- The window's height for a page whose content is `contentHeight` tall: the title band, the gap
--- under it, the content, the gap over the footer and the footer, within MIN_HEIGHT..HEIGHT. Pure.
function Onboarding.frameHeight(contentHeight)
	local chrome = Theme.SIZES.titleBand + 16 + 24 + Onboarding.BUTTON_HEIGHT + 14
	local wanted = chrome + math.ceil(type(contentHeight) == "number" and contentHeight or 0)
	return math.max(Onboarding.MIN_HEIGHT, math.min(Onboarding.HEIGHT, wanted))
end

--- The title band for page `index` of `count` on `key`, Widgets.setTitleBand's shape. Pure.
function Onboarding.head(key, index, count, input)
	-- "Welcome to Forever Sixty" under the heading on page 1 only (round 2: the same line on every
	-- page carried nothing).
	return { title = L[HEADINGS[key]], meta = index == 1 and L.tourTitle or nil, classToken = input.classToken,
		faction = input.faction, number = index, unit = string.format(L.tourPageOf, count) }
end

--- The footer for page `index` of `count`: Skip tour (and Tooltips only on page 1) left, Back and
--- Next right; the last page Back, Open my panel and Done.
local function drawFooter(index, count)
	local footer = Onboarding.footer
	local last = index == count
	if not last then
		local skip = button(footer, L.tourSkip, Onboarding.finish)
		skip:SetPoint("BOTTOMLEFT", footer, "BOTTOMLEFT", 0, 0)
		if index == 1 then
			local quiet = button(footer, L.tourJustTooltips, Onboarding.justTooltips, L.tourJustTooltipsHint)
			quiet:SetPoint("LEFT", skip, "RIGHT", 8, 0)
		end
	else
		-- The last page has no Skip tour: "/fs help opens this again." takes its place (spec 6.3).
		local again = text(footer, L.tourAgain, "umberMuted", "small")
		again:SetPoint("LEFT", footer, "LEFT", 0, 0)
	end
	local right
	if last then
		right = button(footer, L.tourDone, Onboarding.finish, nil, true)
		right:SetPoint("BOTTOMRIGHT", footer, "BOTTOMRIGHT", 0, 0)
		local open = button(footer, L.tourOpenPanel, Onboarding.openPanel)
		open:SetPoint("RIGHT", right, "LEFT", -8, 0)
		right = open
	else
		right = button(footer, L.tourNext, Onboarding.next, nil, true)
		right:SetPoint("BOTTOMRIGHT", footer, "BOTTOMRIGHT", 0, 0)
	end
	if index > 1 then
		local back = button(footer, L.tourBackButton, Onboarding.back)
		back:SetPoint("RIGHT", right, "LEFT", -8, 0)
	end
end

--- Draws the current page.
function Onboarding.draw()
	local frame = Onboarding.ensure()
	reset()
	local input = Onboarding.inputs()
	local keys = Onboarding.pageKeys(input.inGuild)
	Onboarding.page = math.max(1, math.min(Onboarding.page or 1, #keys))
	local key = keys[Onboarding.page]
	Widgets.setTitleBand(Onboarding.band, Onboarding.head(key, Onboarding.page, #keys, input))
	Onboarding.leadText:SetText(Onboarding.lead(key) or "")
	local leadHeight = height(Onboarding.leadText, 17)
	Onboarding.body:SetPoint("TOPLEFT", Onboarding.content, "TOPLEFT", 0, -(leadHeight + Onboarding.LEAD_GAP))
	local bodyHeight = DRAW[key](Onboarding.body, input)
	frame:SetHeight(Onboarding.frameHeight(leadHeight + Onboarding.LEAD_GAP + bodyHeight))
	drawFooter(Onboarding.page, #keys)
	frame:Show()
	return key
end

function Onboarding.next()
	Onboarding.page = (Onboarding.page or 1) + 1
	local count = #Onboarding.pageKeys(Onboarding.inputs().inGuild)
	if Onboarding.page > count then
		Onboarding.finish()
		return
	end
	Onboarding.draw()
end

function Onboarding.back()
	Onboarding.page = math.max(1, (Onboarding.page or 1) - 1)
	Onboarding.draw()
end

--- Opens the tour on page 1 (/fs help, and the first login). Never in combat or inside an instance:
--- then it waits (Onboarding.pendingAfterCombat, the zone change).
function Onboarding.open(page)
	if type(CreateFrame) ~= "function" then
		return false
	end
	if Theme.inCombat() or Chip.combat then
		Onboarding.pendingAfterCombat = page or 1
		return false
	end
	if Theme.instanceName ~= nil and Theme.instanceName() ~= nil then
		Onboarding.pendingOutside = page or 1
		return false
	end
	Onboarding.page = page or 1
	Onboarding.draw()
	return true
end

--- PLAYER_REGEN_DISABLED: hides an open tour without ending it, to come back once after combat on
--- the page it was on. PLAYER_REGEN_ENABLED: brings it back. ZONE_CHANGED_NEW_AREA: a tour waiting
--- to be outside an instance shows.
function Onboarding.onEvent(event)
	if event == "PLAYER_REGEN_DISABLED" and Onboarding.frame ~= nil and Onboarding.frame:IsShown() then
		Onboarding.pendingAfterCombat = Onboarding.page or 1
		Onboarding.hiddenByCombat = true
		Onboarding.hidingForCombat = true
		Onboarding.frame:Hide()
		Onboarding.hidingForCombat = nil
	elseif event == "PLAYER_REGEN_ENABLED" and Onboarding.pendingAfterCombat ~= nil then
		local page = Onboarding.pendingAfterCombat
		Onboarding.pendingAfterCombat = nil
		-- A tour combat hid comes back once a login; one that waited for combat to end to open first opens.
		if Onboarding.hiddenByCombat then
			Onboarding.hiddenByCombat = nil
			if (Onboarding.returnedAfterCombat or 0) >= 1 then
				return
			end
			Onboarding.returnedAfterCombat = (Onboarding.returnedAfterCombat or 0) + 1
		end
		Onboarding.open(page)
	elseif event == "ZONE_CHANGED_NEW_AREA" and Onboarding.pendingOutside ~= nil then
		local page = Onboarding.pendingOutside
		Onboarding.pendingOutside = nil
		Onboarding.open(page)
	end
end

--- The events the tour listens on. Idempotent.
function Onboarding.register()
	if Onboarding.events ~= nil then
		return Onboarding.events
	end
	local frame = CreateFrame("Frame")
	Theme.noteFrame("Onboarding events", "login")
	for _, event in ipairs({ "PLAYER_REGEN_DISABLED", "PLAYER_REGEN_ENABLED", "ZONE_CHANGED_NEW_AREA" }) do
		Theme.registerEvent(frame, event)
	end
	frame:SetScript("OnEvent", function(_, event)
		local ok, err = pcall(Onboarding.onEvent, event)
		if not ok then
			Theme.note(string.format(L.diagOnboardingFailed, tostring(err)))
		end
	end)
	Onboarding.events = frame
	-- The login's welcome (SiteStatus, 5 s after entering the world) is the tour, when due: the
	-- tour's own rule decides (this version not ended on the account), not the welcome stamp alone.
	SiteStatus.tourDue = function()
		return Onboarding.due(ForeverSixtyDB, Prefs.flag("tooltipsOnly"))
	end
	SiteStatus.openTour = function()
		if SiteStatus.tourDue() then
			Onboarding.open(1)
		end
	end
	return frame
end

-- The Builds dialog's open and close drive the step aside (spec 6.1): any opener, the tour's own
-- Change, the panel's, /fs build, behaves the same.
Hub.beforeOpen = Onboarding.stepAside
Hub.afterHide = Onboarding.stepBack

ns.Onboarding = Onboarding
return Onboarding
