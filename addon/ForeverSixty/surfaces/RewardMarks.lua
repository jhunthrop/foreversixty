-- addon/ForeverSixty/surfaces/RewardMarks.lua
-- The quest reward and roll marks (design/specs/2026-11-08-addon-ux.md section 2.4): which one
-- is for me?
--
-- On the quest frame's reward choices: the green arrow on the icon of the reward with the best
-- gain from the 1.0 floor (only one), "Best for 20 to 29" in the client gold under the name of the
-- reward that is the band's pick (it may be the same one; the item tooltip's own line in its own gold,
-- ItemTooltip.bestText, so the three never word the pick two ways: design/specs/2026-11-13-addon-hub.md
-- 5), and, when no reward helps, "Sells for most",
-- muted, under the reward with the highest sell price (the client's own GetItemInfo price, so
-- it is exact); nothing on the others. A reward the character cannot use gets no label: its
-- hover says why. On a group loot roll frame, the green arrow on the item's icon when it is an
-- upgrade and one short line under the item's name: "+5.8 DPS", "+5.8 DPS · Best at level 60",
-- "Best at level 60", "Sidegrade", "Downgrade", "Not for you". No source: the item is dropping in
-- front of the player.
--
-- One verdict source: every item is judged by ItemTooltip.judge, the part of the tooltip's own
-- section that scores it, so a reward's label, a roll's line and the item's tooltip can never
-- disagree.
--
-- The taint rule: nothing is ever written into a client frame. Each label, line and arrow is
-- drawn on an addon-owned overlay frame parented to the client's reward button or roll frame
-- (no name, no scripts), so it shows and hides with it; the overlays are made the first time
-- one is needed, never at load. The only client hook is a hooksecurefunc on QuestInfo_Display.
-- Nothing polls. The marks follow the "Gear tips on item tooltips" switch.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Prefs = ns.Prefs or require("Prefs")
local Compat = ns.Compat or require("Compat")
local Progress = ns.Progress or require("Progress")
local ItemTooltip = ns.ItemTooltip or require("ItemTooltip")
local Gear = ns.Gear or require("Gear")

local RewardMarks = {}

--- The generated data table, set by Options.register().
RewardMarks.data = nil

--- The event registered quietly (Options names the rest): the reward list's own refresh,
--- which an older client may not have.
RewardMarks.QUIET_EVENTS = { "QUEST_ITEM_UPDATE" }

--- The client's reward buttons, by name prefix: the QuestInfo template's
--- (QuestInfoRewardsFrameQuestInfoItem<n>, the modern quest frame's) and the classic
--- quest frame's own (QuestRewardItem<n>). The first prefix that names any button wins.
RewardMarks.REWARD_BUTTONS = { "QuestInfoRewardsFrameQuestInfoItem", "QuestRewardItem" }
RewardMarks.MAX_REWARD_BUTTONS = 10

--- The client's group loot roll frames: GroupLootFrame1..NUM_GROUP_LOOT_FRAMES (4 where the
--- client does not say), and the modern GroupLootContainer's pooled `rollFrames`.
RewardMarks.ROLL_FRAME_PREFIX = "GroupLootFrame"

--- Whether the quest frame's reward panel is open (QUEST_COMPLETE until QUEST_FINISHED).
RewardMarks.questOpen = false

--- Item ids a label or a line waits on (the client has not sent the item yet).
RewardMarks.waiting = {}

--- The rolls drawn, { [rollId] = frame }: redrawn by RewardMarks.refresh while their frame
--- still shows that roll.
RewardMarks.rolls = {}

-- The pure model -------------------------------------------------------------

--- Whether `base` (ItemTooltip.baseContext's) has weights to score by.
local function scorable(base)
	return type(base) == "table" and type(base.current) == "table" and type(base.current.weights) == "table"
end

--- The tooltip's own gold line when `ctx`'s item is the band's pick at the character's level
--- (ItemTooltip.plannedBand at the current band): "Best for 20 to 29", "Best at level 60"
--- (ItemTooltip.bestText, up to the band whose pick changes); nil for any other item.
local function bandPickWords(ctx)
	local data = ctx.data or {}
	local specBis = ctx.spec ~= nil and type(data.bis) == "table" and data.bis[ctx.spec] or nil
	local band = ctx.faction and Gear.bandFor(specBis, ctx.level) or nil
	local itemId = ctx.item.itemId
	local slots = Gear.slotsFor(ctx.item.equipLocation)
	if band == nil or ItemTooltip.plannedBand(data, ctx.spec, ctx.faction, slots, itemId, ctx.level) ~= band then
		return nil
	end
	for _, slot in ipairs(slots) do
		local pick = Gear.pick(ctx.spec, band, ctx.faction, slot, data)
		if pick ~= nil and pick.item_id == itemId then
			return ItemTooltip.bestText(band, ItemTooltip.bandAfter(data, ctx.spec, ctx.faction, slot, band, itemId))
		end
	end
	return ItemTooltip.bestText(band, nil)
end

--- One item judged as the tooltip judges it: { itemId, gear, pending, wearable, gain, upgrade,
--- isPick, best, sellPrice, primary, waits }. `gear`: it goes in a slot the plan covers; `pending`:
--- the client has not sent it yet (nothing else is said); `wearable`: the character can put it
--- on now; `gain`: its gain over what is worn, rounded as shown (nil where it cannot be
--- scored); `upgrade`: its tooltip calls it one; `isPick`: it is the band's pick, `best` the
--- tooltip's line naming it so ("Best for 20 to 29"); `sellPrice`:
--- the client's vendor price in copper (0 or nil: none); `primary`: the tooltip's own verdict
--- line; `waits`: the item ids that line still waits on.
function RewardMarks.judge(link, base)
	local ctx = ItemTooltip.withItem(base, link)
	local itemId = ctx.item.itemId
	-- The item itself first: gear the client has not sent is not judged at all.
	if itemId ~= nil and #Gear.slotsFor(ctx.item.equipLocation) > 0 and not Compat.isItemCached(itemId) then
		return { itemId = itemId, gear = true, pending = true, waits = {} }
	end
	local waits = {}
	local judged = ItemTooltip.judge(ctx, waits)
	local result = { itemId = itemId, gear = judged.gear, primary = judged.primary, waits = waits,
		sellPrice = Compat.itemSellPrice(link) }
	if not judged.gear then
		return result
	end
	result.wearable = #judged.usable > 0
	result.best = result.wearable and bandPickWords(ctx) or nil
	result.isPick = result.best ~= nil
	local verdict = judged.verdict
	if verdict ~= nil and verdict.delta ~= nil then
		local gain = Progress.roundedGain(-verdict.delta, judged.unit == "score" and "score" or "dps")
		result.gain, result.unit = gain, judged.unit
		result.upgrade = (verdict.state == "empty" and gain > 0) or gain >= Progress.MIN_GAIN
	elseif judged.wornSlot ~= nil then
		result.gain = 0
	end
	return result
end
--- Every id a judgement waits on goes to RewardMarks.waiting: its GET_ITEM_INFO_RECEIVED
--- redraws the label or the line (fix round 1: a worn item's name too).
local function noteWaits(judged)
	if judged.pending and judged.itemId ~= nil then
		RewardMarks.waiting[judged.itemId] = true
	end
	for id in pairs(judged.waits or {}) do
		RewardMarks.waiting[id] = true
	end
end

--- The labels for `count` reward choices from their judgements (RewardMarks.judge's, by
--- choice index): { [index] = { text?, color?, kind, arrow? } }.
---
--- - the arrow on the choice with the biggest gain the character can wear now that its
---   tooltip calls an upgrade (the first on a tie): never a plate choice at 21, never a +0.3
---   Sidegrade;
--- - the tooltip's "Best for 20 to 29" (the client gold) under the choice that is the band's pick, wearable now;
--- - with neither, and only when no choice is of any use (each is gear that scores nothing over
---   what is worn, or cannot be worn now; a choice that is not gear, or that the weights cannot
---   price, may be of use): "Sells for most" (muted) under the choice with the highest sell
---   price, none when no choice sells;
--- - nothing at all while a choice is still on its way from the client.
function RewardMarks.rewardLabels(judgements, count)
	local labels = {}
	local best, bestGain
	local useless = count > 0
	for index = 1, count do
		local judged = judgements[index]
		if type(judged) ~= "table" or judged.pending then
			return {}
		end
		if not judged.gear or (judged.wearable and judged.gain == nil) then
			useless = false
		elseif judged.wearable and (judged.gain > 0 or judged.isPick) then
			useless = false
			if judged.upgrade and (bestGain == nil or judged.gain > bestGain) then
				best, bestGain = index, judged.gain
			end
		end
		if judged.isPick then
			labels[index] = { kind = "best", text = judged.best, color = "clientGold" }
		end
	end
	if best ~= nil then
		labels[best] = labels[best] or { kind = "upgrade" }
		labels[best].arrow = true
	end
	if next(labels) == nil and useless then
		local richest, price
		for index = 1, count do
			-- A reward the character cannot use gets no label (its hover says why).
			local sell = judgements[index].wearable and tonumber(judgements[index].sellPrice) or 0
			if sell > 0 and (price == nil or sell > price) then
				richest, price = index, sell
			end
		end
		if richest ~= nil then
			labels[richest] = { kind = "vendor", text = L.rewardSellsForMost, color = "tooltipMuted" }
		end
	end
	return labels
end

--- The roll frame's short verdict for a judgement (spec 2.4): { text, arrow? } with the colours
--- inline: "+5.8 DPS" green with the arrow, "+5.8 DPS · Best at level 60" (green, the client gold),
--- "Best at level 60" (the client gold), "Sidegrade", "Downgrade", "Not for you" (grey). nil for an
--- item that is not gear, is still on its way, or cannot be priced. Pure.
function RewardMarks.rollLine(judged)
	if type(judged) ~= "table" or not judged.gear or judged.pending then
		return nil
	end
	local function paint(key, text)
		return Theme.inlineColor(key, text)
	end
	local best = judged.best ~= nil and paint("clientGold", judged.best) or nil
	if judged.wearable == false or (judged.primary ~= nil and judged.primary.kind == "refusal") then
		return { text = paint("tooltipMuted", L.rollNotForYou) }
	end
	if judged.upgrade then
		local text = paint("tooltipGreen", ItemTooltip.figure(judged.gain, judged.unit))
		if best ~= nil then
			text = text .. paint("tooltipMuted", L.rollBestJoin) .. best
		end
		return { text = text, arrow = true }
	end
	if best ~= nil then
		return { text = best }
	end
	if judged.gain == nil then
		return nil
	end
	if judged.gain <= -Progress.MIN_GAIN then
		return { text = paint("tooltipMuted", string.format(L.rollDowngrade,
			ItemTooltip.figure(judged.gain, judged.unit))) }
	end
	return { text = paint("tooltipMuted", L.tooltipSide) }
end

-- Reading the client ---------------------------------------------------------

--- A client function's answers, or nothing where this client has none.
local function call(name, ...)
	local fn = Compat.clientGlobal(name, "function")
	if fn == nil then
		return nil
	end
	return fn(...)
end

--- A method's answer on `object`, or nil where it has none or it fails.
local function ask(object, method)
	local fn = type(object) == "table" and object[method] or nil
	if type(fn) ~= "function" then
		return nil
	end
	local ok, value = pcall(fn, object)
	return ok and value or nil
end

local function baseContext()
	local data = RewardMarks.data or ns.Data
	return ItemTooltip.baseContext(data), data
end

--- The client's reward buttons for the `count` choices, { [choice index] = button }: a
--- QuestInfo button the client marked a choice (its `type`) answers its own id; a button
--- the client marks nothing on (the classic frame) is the choice of its own number.
function RewardMarks.choiceButtons(count)
	for _, prefix in ipairs(RewardMarks.REWARD_BUTTONS) do
		local found, any = {}, false
		for index = 1, RewardMarks.MAX_REWARD_BUTTONS do
			local button = Compat.clientGlobal(prefix .. index, "table")
			if button ~= nil then
				any = true
				local kind = button.type
				local id = ask(button, "GetID")
				if kind == "choice" and type(id) == "number" and found[id] == nil then
					found[id] = button
				elseif kind == nil and index <= count and found[index] == nil then
					found[index] = button
				end
			end
		end
		if any then
			return found
		end
	end
	return {}
end

--- Every group loot roll frame the client has.
function RewardMarks.rollFrames()
	local frames = {}
	local count = tonumber(Compat.clientGlobal("NUM_GROUP_LOOT_FRAMES", "number")) or 4
	for index = 1, count do
		local frame = Compat.clientGlobal(RewardMarks.ROLL_FRAME_PREFIX .. index, "table")
		if frame ~= nil then
			frames[#frames + 1] = frame
		end
	end
	local container = Compat.clientGlobal("GroupLootContainer", "table")
	local pooled = container ~= nil and container.rollFrames or nil
	if type(pooled) == "table" then
		for _, frame in pairs(pooled) do
			if type(frame) == "table" then
				frames[#frames + 1] = frame
			end
		end
	end
	return frames
end

--- The roll frame showing roll `rollId` (its `rollID`, the client's own field), or nil.
local function rollFrameFor(rollId)
	for _, frame in ipairs(RewardMarks.rollFrames()) do
		if frame.rollID == rollId then
			return frame
		end
	end
	return nil
end

-- The frames -----------------------------------------------------------------

--- Each reward button's label and each roll frame's line ({ overlay, plate, text }),
--- weak-keyed: the client's frames are never written to.
local labels = setmetatable({}, { __mode = "k" })
local lines = setmetatable({}, { __mode = "k" })

--- A reward button's label, for a spec to look at; nil before one was ever drawn there.
function RewardMarks.labelOn(button)
	return labels[button]
end

--- A roll frame's line, for a spec to look at; nil before one was ever drawn there.
function RewardMarks.lineOn(frame)
	return lines[frame]
end

--- One addon-owned frame parented to the client's `owner` and covering it, a level above
--- it, with no name and no scripts: it shows and hides with its owner.
local function overlayFor(owner)
	local overlay = CreateFrame("Frame", nil, owner)
	Theme.noteFrame("Reward mark", "a quest reward or loot roll shown")
	overlay:SetAllPoints(owner)
	local level = type(owner.GetFrameLevel) == "function" and owner:GetFrameLevel() or nil
	if type(level) == "number" then
		overlay:SetFrameLevel(level + 1)
	end
	return overlay
end

local function paintText(text, color)
	text:SetTextColor(Theme.rgb(Theme.HEX[color] or Theme.HEX.body))
end

--- A small font string on a new overlay of `owner`, and the green arrow (hidden) on it. The text casts
--- the client's own 1 px shadow, as its gold text does wherever it sits on the quest frame.
local function labelParts(owner)
	local overlay = overlayFor(owner)
	local text = Theme.fontString(overlay, "OVERLAY", "small")
	if type(text.SetShadowColor) == "function" then
		text:SetShadowColor(Theme.rgb(Theme.HEX.shadow))
		text:SetShadowOffset(1, -1)
	end
	local arrow = overlay:CreateTexture(nil, "OVERLAY")
	arrow:SetTexture(Theme.MEDIA.arrowInline)
	arrow:SetSize(Theme.SIZES.mark, Theme.SIZES.mark)
	arrow:Hide()
	return { overlay = overlay, text = text, arrow = arrow }
end

--- A child region of a client frame by key or by its global name's suffix, or nil.
local function partOf(frame, key, suffix)
	if type(frame[key]) == "table" then
		return frame[key]
	end
	local name = ask(frame, "GetName")
	return type(name) == "string" and Compat.clientGlobal(name .. suffix, "table") or nil
end

--- The reward button's icon and name: the QuestInfo template's `Icon` and `Name`, else the
--- classic frame's <name>IconTexture and <name>Name.
local function rewardParts(button)
	return partOf(button, "Icon", "IconTexture"), partOf(button, "Name", "Name")
end

--- The roll frame's icon and item name: the modern frame's `IconFrame.Icon` and `Name`, else the
--- classic <name>IconFrameIcon and <name>Name.
local function rollParts(frame)
	local iconFrame = partOf(frame, "IconFrame", "IconFrame")
	local icon = iconFrame ~= nil and partOf(iconFrame, "Icon", "Icon") or partOf(frame, "Icon", "IconFrameIcon")
	return icon or iconFrame, partOf(frame, "Name", "Name")
end

--- Places a label's text under `name` (the client's name text) and its arrow at `icon`'s top
--- right; with no name found, the text sits at the owner's bottom right instead.
local function place(drawn, icon, name)
	drawn.text:ClearAllPoints()
	if name ~= nil then
		drawn.text:SetPoint("TOPLEFT", name, "BOTTOMLEFT", 0, -1)
	else
		local pad = Theme.SIZES.labelPad
		drawn.text:SetPoint("BOTTOMRIGHT", drawn.overlay, "BOTTOMRIGHT", -pad * 2, pad * 2)
	end
	if type(drawn.text.SetJustifyH) == "function" then
		drawn.text:SetJustifyH("LEFT")
	end
	drawn.arrow:ClearAllPoints()
	drawn.arrow:SetPoint("TOPRIGHT", icon or drawn.overlay, "TOPRIGHT", 2, 2)
end

--- Shows `line` ({ text?, color?, arrow? }) on `drawn`, or hides it all with nil.
local function show(drawn, line)
	if line == nil then
		drawn.text:Hide()
		drawn.arrow:Hide()
		return
	end
	if line.text ~= nil then
		drawn.text:SetText(line.text)
		if line.color ~= nil then
			paintText(drawn.text, line.color)
		end
		drawn.text:Show()
	else
		drawn.text:Hide()
	end
	if line.arrow then
		drawn.arrow:Show()
	else
		drawn.arrow:Hide()
	end
end

--- The reward button's label and arrow, on its overlay: made the first time one shows, hidden
--- with `line` nil.
local function setLabel(button, line)
	local label = labels[button]
	if line == nil then
		if label ~= nil then
			show(label, nil)
		end
		return
	end
	if label == nil then
		label = labelParts(button)
		labels[button] = label
	end
	place(label, rewardParts(button))
	show(label, line)
end

--- The roll frame's short verdict under the item's name and the arrow on its icon, on its
--- overlay; hidden with `line` nil.
local function setLine(frame, line)
	local drawn = lines[frame]
	if line == nil then
		if drawn ~= nil then
			show(drawn, nil)
		end
		return
	end
	if drawn == nil then
		drawn = labelParts(frame)
		lines[frame] = drawn
	end
	place(drawn, rollParts(frame))
	show(drawn, line)
end
local function enabled()
	return not RewardMarks.disabled and Prefs.flag("tooltip") ~= false
end

--- Hides every reward label already drawn.
local function hideRewards()
	for button in pairs(labels) do
		setLabel(button, nil)
	end
end

--- Hides every label and line already drawn (the switch went off, or the marks failed).
function RewardMarks.hideAll()
	hideRewards()
	for frame in pairs(lines) do
		setLine(frame, nil)
	end
end

--- A failure here must never become a Lua error on every quest or roll: the marks turn
--- themselves off until /reload, hide what they drew and say so once in /fs diag.
local function guarded(action, ...)
	if RewardMarks.disabled then
		return
	end
	local ok, err = pcall(action, ...)
	if not ok then
		RewardMarks.disabled = true
		pcall(RewardMarks.hideAll)
		Theme.note(string.format(L.diagRewardMarksFailed, tostring(err)))
	end
end

--- The reward panel's labels, from the choices the client offers now.
local function drawRewards()
	if not enabled() or not RewardMarks.questOpen then
		hideRewards()
		return
	end
	local count = tonumber((call("GetNumQuestChoices"))) or 0
	local base = baseContext()
	local shown = {}
	if count > 0 and scorable(base) then
		local judgements = {}
		for index = 1, count do
			local link = call("GetQuestItemLink", "choice", index)
			if type(link) == "string" then
				local judged = RewardMarks.judge(link, base)
				noteWaits(judged)
				judgements[index] = judged
			end
		end
		shown = RewardMarks.rewardLabels(judgements, count)
	end
	local buttons = RewardMarks.choiceButtons(count)
	for button in pairs(labels) do
		local labelled = false
		for index, choice in pairs(buttons) do
			labelled = labelled or (choice == button and shown[index] ~= nil)
		end
		if not labelled then
			setLabel(button, nil)
		end
	end
	for index, button in pairs(buttons) do
		setLabel(button, shown[index])
	end
end

function RewardMarks.drawRewards()
	guarded(drawRewards)
end

--- Every line drawn on a frame that no longer shows the roll it was drawn for (the client
--- reused the frame for another roll), hidden (fix round 1).
local function clearStaleLines()
	for frame, drawn in pairs(lines) do
		if drawn.rollId ~= nil and frame.rollID ~= drawn.rollId then
			setLine(frame, nil)
			drawn.rollId = nil
		end
	end
end

--- One roll's line: the item's verdict beside the frame showing it. A frame the client has
--- not opened yet is looked for once more on the next frame (`retried`). A frame reused
--- for another roll loses its old line first, whether the new roll's frame is found or not.
local function drawRoll(rollId, retried)
	if not enabled() then
		return
	end
	clearStaleLines()
	local frame = rollFrameFor(rollId)
	if frame == nil then
		if not retried then
			Theme.after(0, function()
				guarded(drawRoll, rollId, true)
			end)
		end
		return
	end
	RewardMarks.rolls[rollId] = frame
	local link = call("GetLootRollItemLink", rollId)
	local line
	if type(link) == "string" then
		local base = baseContext()
		local judged = RewardMarks.judge(link, base)
		noteWaits(judged)
		line = RewardMarks.rollLine(judged)
	end

	setLine(frame, line)
	if lines[frame] ~= nil then
		lines[frame].rollId = line ~= nil and rollId or nil
	end
end

function RewardMarks.drawRoll(rollId)
	guarded(drawRoll, rollId)
end

--- Every roll drawn whose frame still shows it, drawn again; the rest forgotten.
local function redrawRolls()
	for rollId, frame in pairs(RewardMarks.rolls) do
		if frame.rollID == rollId then
			drawRoll(rollId, true)
		else
			RewardMarks.rolls[rollId] = nil
		end
	end
end

-- Events -----------------------------------------------------------------------

--- Draws the reward labels on the next frame, once the client has laid the reward
--- buttons out for this quest (its own handler of the same event may run after this one);
--- at once on a client with no timer.
local function scheduleRewards()
	if not Theme.after(0, RewardMarks.drawRewards) then
		RewardMarks.drawRewards()
	end
end

--- The quest frame's reward panel opened (QUEST_COMPLETE) or its items arrived
--- (QUEST_ITEM_UPDATE), it closed (QUEST_FINISHED), or a roll began (START_LOOT_ROLL).
function RewardMarks.onEvent(event, ...)
	if event == "QUEST_COMPLETE" then
		RewardMarks.questOpen = true
		scheduleRewards()
	elseif event == "QUEST_ITEM_UPDATE" then
		if RewardMarks.questOpen then
			scheduleRewards()
		end
	elseif event == "QUEST_FINISHED" then
		RewardMarks.questOpen = false
		guarded(hideRewards)
	elseif event == "START_LOOT_ROLL" then
		RewardMarks.drawRoll((...))
	end
end

--- The client redrew the reward list (QuestInfo_Display): for the quest log, which shows
--- the same buttons for another quest, no label; for the reward panel, the labels again.
function RewardMarks.onQuestDisplay()
	guarded(function()
		local info = Compat.clientGlobal("QuestInfoFrame", "table")
		if info ~= nil and info.questLog then
			hideRewards()
		elseif RewardMarks.questOpen then
			drawRewards()
		end
	end)
end

--- Whether a label or a line waits on `itemId`.
function RewardMarks.isWaiting(itemId)
	return RewardMarks.waiting[tonumber(itemId)] == true
end

--- GET_ITEM_INFO_RECEIVED: redraws when a label or a line waits on this item, and does
--- nothing at all otherwise.
function RewardMarks.onItemInfo(itemId)
	if not RewardMarks.isWaiting(itemId) then
		return
	end
	RewardMarks.waiting[tonumber(itemId)] = nil
	RewardMarks.refresh()
end

--- Something the verdicts read changed (a Settings switch, the bags, the level, the
--- spec): every label and line shown is drawn again, or hidden with the switch off.
function RewardMarks.refresh(data)
	if data ~= nil then
		RewardMarks.data = data
	end
	if not enabled() then
		guarded(RewardMarks.hideAll)
		return
	end
	if RewardMarks.questOpen then
		guarded(drawRewards)
	end
	guarded(redrawRolls)
end

--- Hooks the client's reward list redraw once, after the original (hooksecurefunc), where
--- the client has it. Idempotent; hooks nothing else.
function RewardMarks.register()
	if RewardMarks.registered then
		return
	end
	RewardMarks.registered = true
	local hook = Compat.clientGlobal("hooksecurefunc", "function")
	if hook ~= nil and Compat.clientGlobal("QuestInfo_Display", "function") ~= nil then
		hook("QuestInfo_Display", RewardMarks.onQuestDisplay)
	end
end

ns.RewardMarks = RewardMarks
return RewardMarks
