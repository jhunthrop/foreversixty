-- addon/ForeverSixty/views/GuildView.lua
-- The Guild page: who you are in your guild, from the client; what your
-- guild has done and how its roster rates, from the Forever Sixty Data
-- addon; and, from the companion's private inbox, this account's own claim
-- state and (for an officer or leader) the approval queue's size (Wave C,
-- section 3/4). Without the data addon the page still names the guild and
-- says what to install; it never shows an empty table as if the guild had
-- no data, and it never shows the private state block for a guild the
-- inbox message is not actually about (guildMessage below).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Ratings = ns.Ratings or require("Ratings")
local Codec = ns.Codec or require("Codec")
local Export = ns.Export or require("Export")

local GuildView = {}

--- Section 3/4's claim-state line, keyed by the "guild" message's own
--- `claim_state` (guilds.ClaimStateView.State on the API side: unclaimed,
--- pending, claimed, contested). A state this addon build does not
--- recognise -- a future phase the site added -- reads as no line at all,
--- the same forward-compatibility rule Codec.inboxMessages already
--- applies to a message type it does not recognise.
local CLAIM_LINES = {
	unclaimed = L.guildClaimUnclaimed,
	pending = L.guildClaimPending,
	claimed = L.guildClaimClaimed,
	contested = L.guildClaimContested,
}

--- The "guild" inbox message for the logged-in character, if the
--- companion has one AND it is about the guild the client itself
--- reports (`guildName`) -- a stale or mismatched message (an alt that
--- transferred guilds since the companion's last ten-minute sync) is
--- dropped rather than shown under the wrong guild's tab, same as this
--- file already refuses to show an empty roster table as if it were a
--- guild with no data.
local function guildMessage(guildName)
	local key = Export.characterKey and Export.characterKey() or nil
	for _, message in ipairs(Codec.inboxMessages(_G.ForeverSixtyInbox, key, "guild")) do
		if message.guild_name == guildName then
			return message
		end
	end
	return nil
end

--- The state block's one line, joining whichever of claim state and the
--- approval count apply -- pending_approvals is 0 (and so this half is
--- silent) for anyone who is not an officer or leader, per the "guild"
--- message's own contract. nil, not "", when there is nothing to say,
--- so the caller can tell "no message" from "a message with nothing to
--- show" and hide the line either way.
local function stateLine(message)
	if message == nil then
		return nil
	end
	local parts = {}
	local claimText = CLAIM_LINES[message.claim_state]
	if claimText ~= nil then
		parts[#parts + 1] = claimText
	end
	if (message.pending_approvals or 0) > 0 then
		parts[#parts + 1] = string.format(L.guildPendingApprovals, message.pending_approvals)
	end
	if #parts == 0 then
		return nil
	end
	return table.concat(parts, " · ")
end

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
		stateLine = stateLine(guildMessage(info.name)),
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
	-- The private state block (claim state, and for an officer or
	-- leader, the approval queue's size): section 3/4's addition, from
	-- the companion's inbox rather than the public data addon, so it
	-- has its own line below the public standing note.
	view.state = Widgets.label(parent, "", "gold", "small")
	view.state:SetPoint("TOPLEFT", view.note, "BOTTOMLEFT", 0, -S.gap)
	view.state:SetWidth(ctx.contentWidth)
	view.rosterTitle = Widgets.label(parent, L.guildRosterTitle, "muted", "small")
	view.rosterTitle:SetPoint("TOPLEFT", view.state, "BOTTOMLEFT", 0, -S.padding)
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
	view.state:SetText(model.stateLine or "")
	showAs(view.state, model.stateLine ~= nil)
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
