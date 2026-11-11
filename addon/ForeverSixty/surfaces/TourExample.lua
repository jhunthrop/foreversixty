-- addon/ForeverSixty/surfaces/TourExample.lua
-- The first-run tour's page 2 example (review 2026-11-10): the lines the item tooltip adds, on a
-- real item. The character's own first upgrade from the panel when it has one, else the data's
-- best-in-slot pick for the spec it follows (or its first path spec). Pure model plus one reader;
-- Onboarding draws it in the client's tooltip look.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Compat = ns.Compat or require("Compat")
local Follow = ns.Follow or require("Follow")
local Gear = ns.Gear or require("Gear")
local ItemTooltip = ns.ItemTooltip or require("ItemTooltip")

local TourExample = {}

--- The slots page 2's example reads from the data, in order, until one has a pick.
TourExample.EXAMPLE_SLOTS = { "main_hand", "chest", "legs" }

--- Page 2's tooltip example (review 2026-11-10): the lines the item tooltip adds, on a real item.
--- `input.upgrade` (the panel's first upgrade row: name, hex, gain, unit, detail) gives the arrow
--- and its figure over "<slot> · <source>"; else `input.pick` (a data pick: name, hex, band,
--- upcoming, and the decoded pick) gives "Best in slot" ("Best in slot at level 20" below the first
--- band) over its source. { name, hex, lines = { { text, colour }, ... } }, or nil. Pure.
function TourExample.example(input)
	local upgrade, pick = input.upgrade, input.pick
	if type(upgrade) == "table" and type(upgrade.gain) == "number" and upgrade.name ~= nil then
		return { name = upgrade.name, hex = upgrade.hex or Theme.HEX.strong, lines = {
			{ Theme.arrowMarkup() .. ItemTooltip.figure(upgrade.gain, upgrade.unit), "tooltipGreen" },
			{ upgrade.detail ~= "" and upgrade.detail or nil, "tooltipMuted" },
		} }
	end
	if type(pick) == "table" and pick.name ~= nil then
		local verdict = pick.upcoming and string.format(L.tooltipYourBisAt, pick.band) or L.tooltipYourBis
		return { name = pick.name, hex = pick.hex or Theme.HEX.strong, lines = {
			{ verdict, "clientGold" }, { ItemTooltip.sourceLine(pick.pick), "tooltipMuted" },
		} }
	end
	return nil
end

--- The data's pick for `spec` at `level` in the first of EXAMPLE_SLOTS that has one, from the band
--- the level is in, else the first band ahead of it (upcoming). Pure on `data`.
function TourExample.dataPick(data, spec, level, faction)
	local specBis = type(data) == "table" and type(data.bis) == "table" and data.bis[spec] or nil
	if specBis == nil or faction == nil then
		return nil
	end
	local band = Gear.bandFor(specBis, level)
	local upcoming = band == nil
	band = band or Gear.firstBand(specBis, level)
	if band == nil then
		return nil
	end
	for _, slot in ipairs(TourExample.EXAMPLE_SLOTS) do
		local pick = Gear.pick(spec, band, faction, slot, data)
		if pick ~= nil then
			local names = type(data.bisNames) == "table" and data.bisNames or {}
			return { pick = pick, itemId = pick.item_id, band = band, upcoming = upcoming,
				name = names[pick.item_id] }
		end
	end
	return nil
end

--- Page 2's example from the client: the panel's first scored upgrade, else the data's pick for the
--- spec followed (or the first path spec), named and coloured as the client has the item.
function TourExample.inputs(data, input)
	local Panel = ns.Panel
	if Panel ~= nil then
		local ok, region = pcall(function()
			return Panel.upgrades(Panel.inputs(data))
		end)
		for _, row in ipairs(ok and type(region) == "table" and region.rows or {}) do
			if type(row.gain) == "number" and not row.bag then
				return row, nil
			end
		end
	end
	local spec = Follow.currentSpec(data) or (input.specs[1] ~= nil and input.specs[1].slug or nil)
	local pick = spec ~= nil and TourExample.dataPick(data, spec, input.level, input.faction) or nil
	if pick ~= nil then
		pick.name = Compat.itemName(pick.itemId) or pick.name
		local quality = select(3, Compat.itemInfo(pick.itemId))
		pick.hex = quality ~= nil and Theme.hexOf(Theme.qualityColor(quality)) or nil
	end
	return nil, pick
end

ns.TourExample = TourExample
return TourExample
