-- addon/ForeverSixty/surfaces/GearEnchants.lua
-- The Gear tab's enchant rows under the doll (GearTab.enchantRows, EnchantPlan's verdicts): one row per
-- slot the spec's plan enchants and the player's item lacks it, or wears a weaker one. Each row is the
-- enchant's icon and its words ("No enchant: Enchant Chest - Greater Stats", "Enchant Boots - Greater
-- Agility beats +5 Agility") with the slot after them, muted; the hover is the client's own tooltip for
-- the enchant (the kit's or scope's item, else the enchanter's spell).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Compat = ns.Compat or require("Compat")
local HubFrame = ns.HubFrame or require("HubFrame")

local GearEnchants = {}

--- Rows: the most a doll can flag (every enchantable slot), the row's pitch, its icon.
GearEnchants.LAYOUT = { rows = 12, row = 18, icon = 16, gap = 6 }

local ICON_PATH = "Interface\\Icons\\"

--- The client's tooltip for an enchant record: its item (a kit, a scope) by id, else its spell.
local function showEnchantTooltip(owner, plan)
	if plan.itemId ~= nil then
		return Theme.showItemTooltip(owner, plan.itemId)
	end
	if plan.spellId == nil or type(GameTooltip) ~= "table" or type(GameTooltip.SetSpellByID) ~= "function" then
		return Theme.showLines(owner, { plan.name }, "ANCHOR_RIGHT")
	end
	GameTooltip:SetOwner(owner, "ANCHOR_RIGHT")
	GameTooltip:SetSpellByID(plan.spellId)
	GameTooltip:Show()
	return true
end

local function buildRow(parent, width)
	local G = GearEnchants.LAYOUT
	local row = {}
	row.frame = Theme.createFrame("Button", nil, parent)
	row.frame:SetSize(width, G.row)
	row.icon = Theme.icon(row.frame, "ARTWORK", nil, G.icon)
	row.icon:SetPoint("LEFT", row.frame, "LEFT", 0, 0)
	row.text = HubFrame.line(row.frame, "warning", width - G.icon - G.gap)
	row.text:SetPoint("LEFT", row.icon, "RIGHT", G.gap, 0)
	row.frame:SetScript("OnEnter", function(self)
		if self.plan ~= nil then
			showEnchantTooltip(self, self.plan)
		end
	end)
	row.frame:SetScript("OnLeave", function()
		Theme.hideTooltip()
	end)
	return row
end

--- The rows' frames, made once under `parent`, `width` wide: `count` of them (the Me tab's two), else as
--- many as a doll can flag.
function GearEnchants.build(parent, width, count)
	local section = { rows = {} }
	for index = 1, count or GearEnchants.LAYOUT.rows do
		section.rows[index] = buildRow(parent, width)
	end
	return section
end

--- One row's words: the verdict, then the slot, muted. Pure.
function GearEnchants.words(row)
	local lead = row.missing and string.format(L.gearEnchantNone, row.plan.name)
		or string.format(L.gearEnchantBeats, row.plan.name, row.worn or "")
	local slot = L.gearSlotLabels[row.slot] or row.slot
	return lead .. Theme.inlineColor("muted", L.dialogSeparator .. slot)
end

--- A row's words where room is short (the Me tab): the slot, then the enchant without its "Enchant Chest - "
--- lead ("Chest: Minor Stats"), or that over the weaker one worn. Pure.
function GearEnchants.shortWords(row)
	local name = row.plan.name:gsub("^Enchant [^%-]+ %- ", "")
	local text = row.missing and name or string.format(L.gearEnchantBeats, name, row.worn or "")
	return string.format(L.gearEnchantShort, L.gearSlotLabels[row.slot] or row.slot, text)
end

--- Draw `rows` (GearTab.enchantRows) at `y` under `parent`; the y after them. A row that stays a row is
--- moved and repainted in place, never hidden and shown again (it may be under the cursor). `words` (a
--- row's text) is GearEnchants.words unless given.
function GearEnchants.draw(section, parent, rows, y, words)
	local G = GearEnchants.LAYOUT
	for index, row in ipairs(section.rows) do
		local shown = rows[index]
		row.frame.plan = shown ~= nil and shown.plan or nil
		if shown == nil then
			row.frame:Hide()
		else
			local icon = shown.plan.icon ~= nil and ICON_PATH .. shown.plan.icon
				or Compat.spellTexture(shown.plan.spellId) or Theme.UNKNOWN_ICON
			row.icon:SetTexture(icon)
			row.text:SetText((words or GearEnchants.words)(shown))
			row.frame:ClearAllPoints()
			row.frame:SetPoint("TOPLEFT", parent, "TOPLEFT", 0, -(y + (index - 1) * G.row))
			row.frame:Show()
		end
	end
	return y + math.min(#rows, #section.rows) * G.row
end

ns.GearEnchants = GearEnchants
return GearEnchants
