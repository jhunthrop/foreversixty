-- addon/ForeverSixty/surfaces/Onboarding.lua
-- The first open (design/specs/2026-11-13-addon-hub.md 4.1, board hub-me-setup.png): the first login of a
-- new tour version opens the Forever Sixty window on Me with Set up in place of Right now. Three numbered
-- steps, each an icon, the step and its object, and a link to the tab that owns it: Pick your spec (the
-- specs, Builds), Hover an item you wear (the item, Gear, with the client's own tooltip of it drawn under
-- the steps, TourExample), Get the companion app (the gold foreversixty.gg/app). A step done is ticked and
-- muted. Never in combat or inside an instance: then it waits for combat to end or for the zone to change.
-- /fs tour brings it back. It replaces the 640 tour window and its explanatory pages.
--
-- When it is due: this tour version not seen on the account (ForeverSixtyDB.tourSeen), so a first install,
-- a reinstall that wiped SavedVariables, and one that kept an older version's stamp all see it (owner,
-- 2026-11-13: deleting the addon folder keeps SavedVariables). "Tooltips only" quiets it. The hub runs the
-- login timer (Hub.afterEnteringWorld); this file decides and draws.
--
-- Pure model, thin frame: Onboarding.steps is plain tables in, plain tables out; Onboarding.draw lays the
-- section out in Me's right column with MeFrame's own pieces (MeFrame.kit), so it reads as one tab.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Prefs = ns.Prefs or require("Prefs")
local Follow = ns.Follow or require("Follow")
local Talents = ns.Talents or require("Talents")
local Chip = ns.Chip or require("Chip")
local Hub = ns.Hub or require("Hub")
local BuildsPane = ns.BuildsPane or require("BuildsPane")
local MeTab = ns.MeTab or require("MeTab")
local MeFrame = ns.MeFrame or require("MeFrame")
local ItemTooltip = ns.ItemTooltip or require("ItemTooltip")
local TourExample = ns.TourExample or require("TourExample")

local Onboarding = {}

--- The tour's version. Bump it when Set up changes and every account sees it once more. 1 was 2.0.38's
--- tour (stamped as welcomedAt alone), 2 the tour window with its leads, 3 the window's Set up.
Onboarding.TOUR_VERSION = 3
--- The step icons Set up draws where the step has no item or spec of its own.
Onboarding.ICONS = { app = "Interface\\Icons\\INV_Scroll_08" }
--- The client's own green tick (the ready check's), drawn over a done step's number.
Onboarding.TICK = "Interface\\RaidFrame\\ReadyCheck-Ready"
--- Pixel layout: the number's column, the tick's side, the row's two lines, the gap above the tooltip.
Onboarding.LAYOUT = { number = 16, tick = 12, row = 38, tooltipGap = 6 }

-- When it shows -------------------------------------------------------------------------------------

--- Whether Set up is due this login: this version not seen on this account and "Tooltips only" off. Pure.
function Onboarding.due(db, tooltipsOnly)
	if tooltipsOnly then
		return false
	end
	return type(db) ~= "table" or db.tourSeen ~= Onboarding.TOUR_VERSION
end

--- Stamps the account: this tour version seen, and the first welcome's time (the chip's first logins read
--- it) kept from the first time it was set.
function Onboarding.stamp()
	ForeverSixtyDB = ForeverSixtyDB or {}
	ForeverSixtyDB.tourSeen = Onboarding.TOUR_VERSION
	ForeverSixtyDB.welcomedAt = ForeverSixtyDB.welcomedAt or (type(time) == "function" and time() or 0)
end

--- Opens the window on Me with Set up (the first login, /fs tour). In combat it waits for combat to end,
--- inside an instance for the zone to change (Onboarding.pending). The account is stamped when it shows.
--- True when it showed.
function Onboarding.open()
	if Theme.inCombat() or Chip.combat then
		Onboarding.pending = "combat"
		return false
	end
	if Theme.instanceName() ~= nil then
		Onboarding.pending = "outside"
		return false
	end
	Onboarding.pending = nil
	Onboarding.showing = true
	if Hub.open("me") == nil then
		Onboarding.showing = false
		return false
	end
	Onboarding.stamp()
	return true
end

--- The hub's login timer (Hub.firstOpen): Set up when it is due.
function Onboarding.firstOpen()
	if Onboarding.due(ForeverSixtyDB, Prefs.flag("tooltipsOnly")) then
		return Onboarding.open()
	end
	return false
end

--- The window hid: Me goes back to Right now on its next open.
function Onboarding.onWindowHide()
	Onboarding.showing = false
end

--- An item tooltip the player hovered (ItemTooltip.onHover): step 2 is done, once for the account; an
--- open Set up redraws with its tick.
function Onboarding.noteHover()
	if type(ForeverSixtyDB) == "table" and ForeverSixtyDB.tourHovered then
		return false
	end
	ForeverSixtyDB = ForeverSixtyDB or {}
	ForeverSixtyDB.tourHovered = true
	if Onboarding.showing then
		MeFrame.refresh()
	end
	return true
end

--- PLAYER_REGEN_ENABLED: Set up waiting on combat shows. ZONE_CHANGED_NEW_AREA: one waiting to be outside
--- an instance shows (or waits again, still inside).
function Onboarding.onEvent(event)
	local waited = (event == "PLAYER_REGEN_ENABLED" and Onboarding.pending == "combat")
		or (event == "ZONE_CHANGED_NEW_AREA" and Onboarding.pending == "outside")
	if waited then
		return Onboarding.open()
	end
	return false
end

-- The pure model ------------------------------------------------------------------------------------

local function specNames(specs)
	local names = {}
	for _, spec in ipairs(specs or {}) do
		names[#names + 1] = spec.name
	end
	return #names > 0 and table.concat(names, L.dialogSeparator) or nil
end

--- Set up (spec 4.1): { note, rows, example }. Each row { number, done, line, sub, link, icon, itemId,
--- crest, tab }: `crest` the class token whose ringed crest stands for the step (the spec not picked yet),
--- `tab` the tab its gold link opens, `link` a gold address under the line. `input`: { build, buildName,
--- specIcon, classToken, specs, example ({ itemId, name, slotId }), gearOpen, companion, hovered }. Pure.
function Onboarding.steps(input)
	local picked = type(input.build) == "table"
	local example = input.example
	local rows = {
		{ done = picked, line = L.hubPickSpec, tab = "builds",
			sub = picked and input.buildName or specNames(input.specs),
			icon = picked and input.specIcon or nil, crest = not picked and input.classToken or nil },
		{ done = input.hovered == true, line = L.meStepHover, tab = input.gearOpen and "gear" or nil,
			sub = example ~= nil and example.name ~= nil and string.format(L.meStepBelow, example.name) or nil,
			itemId = example ~= nil and example.itemId or nil },
		{ done = input.companion == true, line = L.meStepApp, icon = Onboarding.ICONS.app,
			link = not input.companion and L.meAppLink or nil },
	}
	local done = 0
	for index, row in ipairs(rows) do
		row.number = index
		done = done + (row.done and 1 or 0)
	end
	-- The numbered rows say how many there are; the note says only how far along the player is.
	local note = done > 0 and string.format(L.meStepsDone, done, #rows) or nil
	return { note = note, rows = rows, example = example }
end

-- Reading the client --------------------------------------------------------------------------------

--- Set up's inputs from the live modules and the client.
function Onboarding.inputs(data)
	data = data or MeFrame.data or ns.Data
	local build = Follow.build
	local classSlug = Talents.playerClassSlug()
	return {
		build = build,
		buildName = type(build) == "table" and build.name or nil,
		specIcon = build ~= nil and BuildsPane.specIcon(data, Follow.buildSpec(data, build)) or nil,
		classToken = Chip.classToken(),
		specs = Chip.pathSpecs(data, classSlug, nil, Chip.level()),
		example = TourExample.read(data),
		gearOpen = Hub.reason("gear") == nil,
		companion = MeTab.hasCompanion(),
		hovered = type(ForeverSixtyDB) == "table" and ForeverSixtyDB.tourHovered == true,
	}
end

--- MeFrame's hook: Set up's model while it shows, else nil (Right now then).
function Onboarding.model(data)
	if not Onboarding.showing then
		return nil
	end
	return Onboarding.steps(Onboarding.inputs(data))
end

-- The frame -------------------------------------------------------------------------------------------

--- The section, once, into Me's right column: the heading, a row per step, the tooltip.
local function build(parent, width)
	local kit = MeFrame.kit
	local G = MeFrame.LAYOUT
	local section = kit.section(parent, width, L.meSetUp)
	section.rows = {}
	local room = width - Onboarding.LAYOUT.number - G.nowIcon - G.nowGap - G.nowLink
	for index = 1, 3 do
		local row = {}
		row.number = kit.text(parent, "muted", Onboarding.LAYOUT.number)
		-- The tick whole, not cropped as an ability icon is.
		row.tick = Theme.newTexture(parent, "OVERLAY")
		row.tick:SetSize(Onboarding.LAYOUT.tick, Onboarding.LAYOUT.tick)
		row.tick:SetTexture(Onboarding.TICK)
		row.cell = kit.itemCell(parent, G.nowIcon)
		row.crest, row.ring = Theme.buildCrest(parent, G.nowIcon)
		row.line = kit.text(parent, "body", room)
		-- The second line runs under the link too: the link sits on the first.
		row.sub = kit.text(parent, "muted", room + G.nowLink)
		row.address = kit.link(parent, "", function()
			MeFrame.openApp()
		end)
		row.open = kit.link(parent, "", function()
			if row.tab ~= nil then
				Hub.select(row.tab)
			end
		end)
		section.rows[index] = row
	end
	section.tooltip = TourExample.ensureTooltip(parent)
	return section
end

local function hideRow(row)
	MeFrame.kit.hideAll({ row.number, row.tick, row.cell, row.crest, row.ring, row.line, row.sub, row.address,
		row.open })
end

--- One step at `top`.
local function drawRow(row, parent, step, top, width)
	local kit, G, O = MeFrame.kit, MeFrame.LAYOUT, Onboarding.LAYOUT
	row.tab = step.tab
	local mid = top + (G.nowIcon + 4 - G.line) / 2
	row.number:SetText(tostring(step.number))
	-- place shows a region; the number and the tick are then left to the step's state.
	kit.place(row.number, parent, 0, mid)
	kit.place(row.tick, parent, 0, top + (G.nowIcon + 4 - O.tick) / 2)
	Theme.setShown({ row.number }, not step.done)
	Theme.setShown({ row.tick }, step.done)
	-- The class's ringed crest while the spec is not picked (one crest language), else the item or spec icon.
	if step.crest ~= nil then
		row.cell:Hide()
		Theme.paintCrest(row.crest, row.ring, step.crest)
		kit.place(row.crest, parent, O.number, top + 2)
		row.ring:Show()
	else
		row.crest:Hide()
		row.ring:Hide()
		kit.paintCell(row.cell, step.itemId, step.icon)
		if type(row.cell.icon.SetDesaturated) == "function" then
			row.cell.icon:SetDesaturated(step.done)
		end
		kit.place(row.cell, parent, O.number, top + 2)
	end
	local x = O.number + G.nowIcon + G.nowGap
	local room = width - x - G.nowLink
	kit.fit(row.line, step.line, room)
	row.line:SetTextColor(Theme.rgb(Theme.HEX[step.done and "muted" or "body"]))
	kit.place(row.line, parent, x, top)
	row.sub:Hide()
	row.address:Hide()
	if step.link ~= nil then
		row.address:SetLabel(step.link)
		kit.place(row.address, parent, x, top + G.line)
	elseif step.sub ~= nil then
		kit.fit(row.sub, step.sub, width - x)
		kit.place(row.sub, parent, x, top + G.line)
	end
	local tab = step.tab ~= nil and Hub.tab(step.tab) or nil
	if tab ~= nil and Hub.isAvailable(tab) then
		row.open:SetLabel(tab.label)
		row.open:ClearAllPoints()
		row.open:SetPoint("TOPRIGHT", parent, "TOPLEFT", width, -top)
		row.open:Show()
	else
		row.open:Hide()
	end
end

--- Set up at `y` in `parent` (Me's right column, `width` wide): the heading with its note, the steps, the
--- client's tooltip of the item step 2 names. The y under it. Built on the first draw.
function Onboarding.draw(parent, width, model, y)
	local kit, O = MeFrame.kit, Onboarding.LAYOUT
	Onboarding.section = Onboarding.section or build(parent, width)
	local section = Onboarding.section
	y = kit.heading(section, parent, width, y, model.note)
	for index, row in ipairs(section.rows) do
		local step = model.rows[index]
		if step == nil then
			hideRow(row)
		else
			drawRow(row, parent, step, y + (index - 1) * O.row, width)
		end
	end
	y = y + #model.rows * O.row
	local example = model.example
	local height = example ~= nil and TourExample.showTooltip(section.tooltip, parent, example.slotId, 0,
		y + O.tooltipGap, width) or 0
	if height > 0 then
		y = y + O.tooltipGap + height
	end
	return y + MeFrame.LAYOUT.sectionGap
end

--- Set up leaves Me's column (Right now draws instead).
function Onboarding.hide()
	local section = Onboarding.section
	if section == nil then
		return
	end
	MeFrame.kit.hideSection(section)
	for _, row in ipairs(section.rows) do
		hideRow(row)
	end
	TourExample.hideTooltip(section.tooltip)
end

--- The events Set up waits on, and its hooks into the hub, Me and the item tooltip. Idempotent.
function Onboarding.register()
	if Onboarding.events ~= nil then
		return Onboarding.events
	end
	local frame = CreateFrame("Frame")
	Theme.noteFrame("Onboarding events", "login")
	for _, event in ipairs({ "PLAYER_REGEN_ENABLED", "ZONE_CHANGED_NEW_AREA" }) do
		Theme.registerEvent(frame, event)
	end
	frame:SetScript("OnEvent", function(_, event)
		local ok, err = pcall(Onboarding.onEvent, event)
		if not ok then
			Theme.note(string.format(L.diagOnboardingFailed, tostring(err)))
		end
	end)
	Onboarding.events = frame
	Hub.firstOpen = Onboarding.firstOpen
	ItemTooltip.onHover = Onboarding.noteHover
	return frame
end

Hub.afterHide = Onboarding.onWindowHide
MeFrame.setup = Onboarding

ns.Onboarding = Onboarding
return Onboarding
