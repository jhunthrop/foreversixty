-- addon/ForeverSixty/surfaces/RewardMarks.lua
-- The quest reward and roll marks (spec section 4.4, task L4.2): which one is for me?
--
-- On the quest frame's reward choices, "Best for Arms" in gold on the choice the weights
-- prefer, and "Vendor", muted, on each choice when none is of any use to the character now
-- (every one scores nothing over what is worn, or cannot be worn yet). On a group loot roll
-- frame, the item tooltip's own verdict line beside the item ("Upgrade +4.2 DPS over Worn
-- Chest", "Sidegrade", "Downgrade ...", "Not for you until 40: plate armor", "Not with your
-- two-hander").
--
-- One verdict source (ruling 1): every item is judged by ItemTooltip.judge, the part of the
-- tooltip's own section that scores it (Gear.currentWeights through Progress.itemVerdict,
-- the Rules.canEquip gate, the two-hander pair rule, the empty-slot fill and the 1.0 floor),
-- so a reward's label, a roll's line and the item's tooltip can never disagree.
--
-- The L4.1 taint rule: nothing is ever written into a client frame. Each label and line is
-- drawn on an addon-owned overlay frame parented to the client's reward button or roll
-- frame (no name, no scripts), so it shows and hides with it; the overlays are made the
-- first time one is needed, never at load. The only client hook is a hooksecurefunc on
-- QuestInfo_Display (the reward list redrawn, for the quest frame or the quest log).
-- Nothing polls. The marks follow the "Gear tips on item tooltips" switch (ruling 6).
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

--- One item judged as the tooltip judges it: { itemId, gear, pending, wearable, gain,
--- upgrade, primary, waits }. `gear`: it goes in a slot the plan covers; `pending`: the
--- client has not sent it yet (nothing else is said); `wearable`: the character can put it
--- on now (Rules allow it and no worn two-hander holds the hand); `gain`: its gain over
--- what is worn, rounded as shown (ItemTooltip.verdictLine's own figure; nil where it
--- cannot be scored); `upgrade`: its tooltip calls it one (fix round 1: from the 1.0 floor
--- over a worn item, at any gain that shows into an empty slot: the line in the success
--- colour); `primary`: the tooltip's own verdict line ({ text, color, kind?, ... });
--- `waits`: the item ids that line still waits on (a worn item's name not sent yet).
function RewardMarks.judge(link, base)
	local ctx = ItemTooltip.withItem(base, link)
	local itemId = ctx.item.itemId
	-- The item itself first (fix round 1): gear the client has not sent is not judged at all.
	if itemId ~= nil and #Gear.slotsFor(ctx.item.equipLocation) > 0 and not Compat.isItemCached(itemId) then
		return { itemId = itemId, gear = true, pending = true, waits = {} }
	end
	local waits = {}
	local judged = ItemTooltip.judge(ctx, waits)
	local result = { itemId = itemId, gear = judged.gear, primary = judged.primary, waits = waits }
	if not judged.gear then
		return result
	end
	result.wearable = #judged.usable > 0
	local verdict = judged.verdict
	if verdict ~= nil and verdict.delta ~= nil then
		local gain = Progress.roundedGain(-verdict.delta, judged.unit == "score" and "score" or "dps")
		result.gain = gain
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
--- choice index): { [index] = { text, color, kind } }.
---
--- - "Best for <specWord>" (gold) on the choice with the biggest gain that the character
---   can wear now and that its tooltip calls an upgrade (the first on a tie): a plate choice
---   at 21 never is, nor a +0.3 Sidegrade (fix round 1: that one gets no label at all);
--- - "Vendor" (muted) on every choice only when none is of any use: each is gear that
---   scores nothing over what is worn (a gain of 0 or less as shown) or cannot be worn
---   now. A choice that is not gear, or gear the weights cannot price, may be of use: no
---   Vendor then;
--- - nothing at all while a choice is still on its way from the client.
function RewardMarks.rewardLabels(judgements, count, specWord)
	local labels = {}
	local best, bestGain
	local useless = count > 0
	for index = 1, count do
		local judged = judgements[index]
		if type(judged) ~= "table" or judged.pending then
			return labels
		end
		if not judged.gear or (judged.wearable and judged.gain == nil) then
			useless = false
		elseif judged.wearable and judged.gain > 0 then
			useless = false
			if judged.upgrade and (bestGain == nil or judged.gain > bestGain) then
				best, bestGain = index, judged.gain
			end
		end
	end
	if best ~= nil then
		if specWord ~= nil then
			labels[best] = { text = string.format(L.rewardBest, specWord), color = "gold", kind = "best" }
		end
	elseif useless then
		for index = 1, count do
			labels[index] = { text = L.rewardVendor, color = "muted", kind = "vendor" }
		end
	end
	return labels
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

--- A dark plate and a small font string on a new overlay of `owner`.
local function labelParts(owner)
	local overlay = overlayFor(owner)
	local plate = Theme.paint(overlay:CreateTexture(nil, "ARTWORK"), "background", Theme.ALPHA.plate)
	local text = Theme.fontString(overlay, "OVERLAY", "small")
	return { overlay = overlay, plate = plate, text = text }
end

local function paintText(text, color)
	text:SetTextColor(Theme.rgb(Theme.HEX[color] or Theme.HEX.body))
end

--- The label at the reward button's bottom right, on its overlay: made the first time it
--- shows, hidden with `line` nil.
local function setLabel(button, line)
	local label = labels[button]
	if line == nil then
		if label ~= nil then
			label.text:Hide()
			label.plate:Hide()
		end
		return
	end
	if label == nil then
		label = labelParts(button)
		local pad = Theme.SIZES.labelPad
		label.text:SetPoint("BOTTOMRIGHT", label.overlay, "BOTTOMRIGHT", -pad * 2, pad * 2)
		label.plate:SetPoint("TOPLEFT", label.text, "TOPLEFT", -pad, pad)
		label.plate:SetPoint("BOTTOMRIGHT", label.text, "BOTTOMRIGHT", pad, -pad)
		labels[button] = label
	end
	label.text:SetText(line.text)
	paintText(label.text, line.color)
	label.text:Show()
	label.plate:Show()
end

--- The verdict `line` beside the roll frame's right edge, on its overlay, fitted to
--- Theme.SIZES.rollLineWidth as the tooltip fits it (the compared name on a second line
--- before it is shortened); hidden with `line` nil.
local function setLine(frame, line)
	local drawn = lines[frame]
	if line == nil then
		if drawn ~= nil then
			drawn.text:Hide()
			drawn.plate:Hide()
		end
		return
	end
	local S = Theme.SIZES
	if drawn == nil then
		drawn = labelParts(frame)
		drawn.plate:SetPoint("LEFT", drawn.overlay, "RIGHT", S.rollGap, 0)
		drawn.text:SetPoint("LEFT", drawn.plate, "LEFT", S.labelPad * 2, 0)
		if type(drawn.text.SetJustifyH) == "function" then
			drawn.text:SetJustifyH("LEFT")
		end
		lines[frame] = drawn
	end
	local text = drawn.text
	local function measure(value)
		text:SetText(value)
		return text:GetStringWidth()
	end
	local fitted
	if line.kind == "verdict" then
		fitted = ItemTooltip.fitVerdict(line.verdict, line.unit, line.wornName, measure, S.rollLineWidth)
	end
	fitted = fitted or { ItemTooltip.fitPlain(line.text, measure, S.rollLineWidth) }
	text:SetText(table.concat(fitted, "\n"))
	paintText(text, line.color)
	local width, height = text:GetStringWidth(), text:GetStringHeight()
	width = type(width) == "number" and math.min(S.rollLineWidth, math.ceil(width)) or S.rollLineWidth
	height = type(height) == "number" and math.ceil(height) or 12 * #fitted
	drawn.plate:SetSize(width + S.labelPad * 4, height + S.labelPad * 4)
	text:Show()
	drawn.plate:Show()
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
	local base, data = baseContext()
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
		shown = RewardMarks.rewardLabels(judgements, count, ItemTooltip.specName(data or {}, base.spec))
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
		line = judged.primary
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
