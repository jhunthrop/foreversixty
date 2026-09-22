-- addon/ForeverSixty/views/GuildView.lua
-- The Guild page: who you are in your guild, from the client; what your
-- guild has done and how its roster rates, from the Forever Sixty Data
-- addon. Without that addon the page still names the guild and says what
-- to install; it never shows an empty table as if the guild had no data.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Ratings = ns.Ratings or require("Ratings")

local GuildView = {}

GuildView.ROSTER_ROWS = 10

local function guildInfo()
	if type(GetGuildInfo) ~= "function" then
		return nil
	end
	local name, rankName, rankIndex = GetGuildInfo("player")
	if name == nil then
		return nil
	end
	return { name = name, rankName = rankName, rankIndex = rankIndex }
end

local function rosterOf(guild)
	local roster = {}
	for _, member in ipairs(guild.members or {}) do
		local card = Ratings.forCharacter(member)
		roster[#roster + 1] = {
			name = member,
			rating = card and card.rating or nil,
			fights = card and card.fights or 0,
			ratingLine = card and tostring(card.rating) or L.guildUnrated,
		}
	end
	table.sort(roster, function(left, right)
		if (left.rating ~= nil) ~= (right.rating ~= nil) then
			return left.rating ~= nil
		end
		if left.rating ~= right.rating then
			return (left.rating or 0) > (right.rating or 0)
		end
		return left.name < right.name
	end)
	return roster
end

--- Everything the page shows, as plain tables and finished strings.
function GuildView.summary()
	local info = guildInfo()
	local data = Ratings.status()
	if info == nil then
		return { inGuild = false, title = L.guildNone, rankLine = L.guildNoneHint, data = data, roster = {} }
	end
	local model = {
		inGuild = true,
		title = info.name,
		rankLine = string.format(L.guildRank, info.rankName or ""),
		data = data,
		roster = {},
	}
	if not data.available then
		model.standingLine = data.reason
		return model
	end
	local guild = Ratings.forGuild(info.name)
	if guild == nil then
		model.standingLine = L.guildNotOnSite
		return model
	end
	model.standing = {
		progress = guild.progress or "",
		nightsLine = string.format(L.guildNights, guild.nights or 0),
		rosterLine = string.format(L.guildRoster, guild.roster or 0),
	}
	model.standingLine = data.staleLine
	model.roster = rosterOf(guild)
	return model
end

local function rosterRow(parent, width)
	local S = Theme.SIZES
	local frame = CreateFrame("Frame", nil, parent)
	frame:SetSize(width, S.rowHeight)
	local text = Widgets.label(frame, "", "body", "small")
	text:SetPoint("LEFT", frame, "LEFT", S.gap * 2, 0)
	local right = Widgets.label(frame, "", "gold", "small")
	right:SetJustifyH("RIGHT")
	right:SetPoint("RIGHT", frame, "RIGHT", -S.gap * 2, 0)
	return { frame = frame, text = text, right = right }
end

local function renderRow(row, item)
	row.text:SetText(item.name)
	row.right:SetText(item.ratingLine)
	row.right:SetTextColor(Theme.rgb(Theme.HEX[item.rating ~= nil and "gold" or "muted"]))
end

local function layout(parent, ctx)
	local S = Theme.SIZES
	local view = { frame = parent, ctx = ctx }
	view.name = Widgets.label(parent, "", "gold", "large")
	view.name:SetPoint("TOPLEFT", parent, "TOPLEFT", S.padding, -S.padding)
	view.rank = Widgets.label(parent, "", "muted", "small")
	view.rank:SetPoint("TOPLEFT", view.name, "BOTTOMLEFT", 0, -S.gap)
	view.progress = Widgets.label(parent, "", "body")
	view.progress:SetPoint("TOPRIGHT", parent, "TOPRIGHT", -S.padding, -S.padding)
	view.nights = Widgets.label(parent, "", "muted", "small")
	view.nights:SetPoint("TOPRIGHT", view.progress, "BOTTOMRIGHT", 0, -S.gap)
	view.note = Widgets.label(parent, "", "muted", "small")
	view.note:SetPoint("TOPLEFT", view.rank, "BOTTOMLEFT", 0, -S.padding)
	view.note:SetWidth(ctx.contentWidth)
	view.rosterTitle = Widgets.label(parent, L.guildRosterTitle, "muted", "small")
	view.rosterTitle:SetPoint("TOPLEFT", view.note, "BOTTOMLEFT", 0, -S.padding)
	view.list = Widgets.list(parent, ctx.contentWidth, GuildView.ROSTER_ROWS, rosterRow)
	view.list.frame:SetPoint("TOPLEFT", view.rosterTitle, "BOTTOMLEFT", 0, -S.gap * 2)
	view.list:SetRenderer(renderRow)
	view.generated = Widgets.label(parent, "", "muted", "small")
	view.generated:SetPoint("BOTTOMLEFT", parent, "BOTTOMLEFT", S.padding, S.padding)
	return view
end

local function showAs(region, shown)
	if shown then
		region:Show()
	else
		region:Hide()
	end
end

function GuildView.apply(view, model)
	view.model = model
	view.name:SetText(model.title)
	view.rank:SetText(model.rankLine or "")
	view.progress:SetText(model.standing and model.standing.progress or "")
	view.nights:SetText(model.standing and model.standing.nightsLine or "")
	view.note:SetText(model.standingLine or "")
	view.note:SetTextColor(Theme.rgb(Theme.HEX[model.data.available and "muted" or "warning"]))
	showAs(view.rosterTitle, #model.roster > 0)
	view.list:SetItems(model.roster)
	view.generated:SetText(model.data.available and model.data.generatedLine or "")
	return view
end

function GuildView.mount(parent, ctx)
	local view = layout(parent, ctx)
	function view.refresh()
		return GuildView.apply(view, GuildView.summary())
	end
	view.refresh()
	return view
end

ns.GuildView = GuildView
return GuildView
