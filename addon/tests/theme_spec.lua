local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local ALL_TEMPLATES = { "PanelTabButtonTemplate", "UIPanelButtonTemplate", "InputBoxTemplate" }

describe("Theme", function()
	local Theme

	local state

	--- Returns the mock state, which several examples below assert on.
	local function start(install)
		state = mock.install(install or {})
		Theme = helper.load("Theme")
		Theme.reset()
		return state
	end

	after_each(function()
		mock.uninstall()
	end)

	it("reads a hex colour as the client's 0-to-1 floats", function()
		start()
		local r, g, b, a = Theme.rgb("e5b955")
		assert.is_true(math.abs(r - 229 / 255) < 1e-9)
		assert.is_true(math.abs(g - 185 / 255) < 1e-9)
		assert.is_true(math.abs(b - 85 / 255) < 1e-9)
		assert.are.equal(1, a)
	end)

	it("carries an alpha through and defaults it to opaque", function()
		start()
		assert.are.equal(1, select(4, Theme.rgb("0d111a")))
		assert.are.equal(0.96, select(4, Theme.rgb("0d111a", 0.96)))
	end)

	it("carries a six-digit hex for every palette entry", function()
		start()
		for name, hex in pairs(Theme.HEX) do
			assert.are.equal(6, #hex, name .. " is not six hex digits")
			assert.is_not_nil(tonumber(hex, 16), name .. " is not hexadecimal")
		end
	end)

	it("says a template this client has is present", function()
		start({ templates = ALL_TEMPLATES })
		assert.is_true(Theme.hasTemplate("Button", "UIPanelButtonTemplate"))
	end)

	it("says a template this client lacks is absent rather than erroring", function()
		start({ templates = {} })
		assert.is_false(Theme.hasTemplate("Button", "UIPanelButtonTemplate"))
	end)

	it("records the missing template where /fs diag can read it", function()
		start({ templates = {} })
		Theme.hasTemplate("Button", "UIPanelButtonTemplate")
		assert.are.same(
			{ string.format(L.diagNoTemplate, "UIPanelButtonTemplate") },
			Theme.diagnostics())
	end)

	it("asks the client about a template only once", function()
		start({ templates = ALL_TEMPLATES })
		Theme.hasTemplate("Button", "UIPanelButtonTemplate")
		local after = #state.frames
		Theme.hasTemplate("Button", "UIPanelButtonTemplate")
		assert.are.equal(after, #state.frames)
	end)

	it("builds with the template when the client has it", function()
		start({ templates = ALL_TEMPLATES })
		local frame, used = Theme.createFrame("Button", nil, _G.UIParent, "button")
		assert.is_true(used)
		assert.are.equal("UIPanelButtonTemplate", frame.template)
	end)

	it("builds a bare frame when the client does not", function()
		start({ templates = {} })
		local frame, used = Theme.createFrame("Button", nil, _G.UIParent, "button")
		assert.is_false(used)
		assert.is_nil(frame.template)
	end)

	it("keeps the client's own font when the inherited font object exists", function()
		start({ fonts = { "GameFontNormal" } })
		local frame = _G.CreateFrame("Frame")
		local region = Theme.fontString(frame, "ARTWORK", "normal")
		assert.are.equal("GameFontNormal", region:GetFont())
		assert.is_nil(mock.firstCall(region, "SetFont"))
	end)

	it("falls back to the shipped font when the inherited font object is missing", function()
		start()
		local frame = _G.CreateFrame("Frame")
		local region = Theme.fontString(frame, "ARTWORK", "normal")
		local call = mock.firstCall(region, "SetFont")
		assert.is_not_nil(call)
		assert.are.equal(Theme.FALLBACK_FONT.path, call[1])
		assert.are.equal(Theme.FALLBACK_FONT.size, call[2])
		assert.are.same(
			{ string.format(L.diagNoTemplate, "GameFontNormal") },
			Theme.diagnostics())
	end)

	it("keeps the client's own font on a region that already exists", function()
		start({ fonts = { "GameFontNormalSmall" } })
		local box = _G.CreateFrame("EditBox")
		Theme.applyFont(box, "small")
		assert.are.equal("GameFontNormalSmall", box:GetFont())
		assert.is_nil(mock.firstCall(box, "SetFont"))
	end)

	it("falls back to the shipped font when the region's font object is missing", function()
		start()
		local box = _G.CreateFrame("EditBox")
		Theme.applyFont(box, "small")
		local call = mock.firstCall(box, "SetFont")
		assert.is_not_nil(call)
		assert.are.equal(Theme.FALLBACK_FONT.path, call[1])
		assert.are.equal(Theme.FALLBACK_FONT.size, call[2])
		assert.are.same(
			{ string.format(L.diagNoTemplate, "GameFontNormalSmall") },
			Theme.diagnostics())
	end)

	it("takes the class's own colour for the header", function()
		start({ globals = { RAID_CLASS_COLORS = { WARRIOR = { r = 0.78, g = 0.61, b = 0.43 } } } })
		local r, g, b = Theme.classColor("WARRIOR")
		assert.are.equal(0.78, r)
		assert.are.equal(0.61, g)
		assert.are.equal(0.43, b)
	end)

	it("falls back to gold when the client has no class colour table", function()
		start()
		assert.are.same({ Theme.rgb(Theme.HEX.gold) }, { Theme.classColor("WARRIOR") })
	end)

	it("falls back to gold for a class token the table does not carry", function()
		start({ globals = { RAID_CLASS_COLORS = { WARRIOR = { r = 1, g = 1, b = 1 } } } })
		assert.are.same({ Theme.rgb(Theme.HEX.gold) }, { Theme.classColor("SKYBORNE") })
	end)

	it("sees the modern Settings API when both halves are there", function()
		start({ globals = { Settings = {
			RegisterCanvasLayoutCategory = function() end,
			RegisterAddOnCategory = function() end,
		} } })
		assert.is_true(Theme.hasSettingsApi())
	end)

	it("does not claim the Settings API on a half-present table", function()
		start({ globals = { Settings = { RegisterCanvasLayoutCategory = function() end } } })
		assert.is_false(Theme.hasSettingsApi())
	end)

	it("registers an event the client knows", function()
		start()
		local frame = _G.CreateFrame("Frame")
		assert.is_true(Theme.registerEvent(frame, "PLAYER_LOGIN"))
		assert.is_true(frame.events.PLAYER_LOGIN)
	end)

	it("records and skips an event the client refuses", function()
		start({ refusedEvents = { TRAIT_CONFIG_UPDATED = true } })
		local frame = _G.CreateFrame("Frame")
		assert.is_false(Theme.registerEvent(frame, "TRAIT_CONFIG_UPDATED"))
		assert.are.same(
			{ string.format(L.diagNoEvent, "TRAIT_CONFIG_UPDATED") },
			Theme.diagnostics())
	end)

	it("is out of combat on a client with no InCombatLockdown", function()
		start()
		assert.is_false(Theme.inCombat())
	end)

	it("asks the client when it has InCombatLockdown", function()
		start({ globals = { InCombatLockdown = function() return true end } })
		assert.is_true(Theme.inCombat())
	end)

	it("equips through C_Item when the client has it", function()
		local equipped = {}
		start({ globals = { C_Item = { EquipItemByName = function(link)
			equipped[#equipped + 1] = link
		end } } })
		assert.is_true(Theme.equip("|Hitem:1234|h"))
		assert.are.same({ "|Hitem:1234|h" }, equipped)
	end)

	it("equips through the flat function when that is the one present", function()
		local equipped = {}
		start({ globals = { EquipItemByName = function(link)
			equipped[#equipped + 1] = link
		end } })
		assert.is_true(Theme.equip("|Hitem:1234|h"))
		assert.are.same({ "|Hitem:1234|h" }, equipped)
	end)

	it("records and refuses to equip when the client has neither", function()
		start()
		assert.is_false(Theme.equip("|Hitem:1234|h"))
		assert.are.same({ L.diagNoEquipApi }, Theme.diagnostics())
	end)

	it("delays through C_Timer.After", function()
		start()
		local ran = false
		assert.is_true(Theme.after(2, function() ran = true end))
		assert.is_false(ran)
		assert.are.equal(1, mock.runTimers(state))
		assert.is_true(ran)
	end)

	it("says so rather than running now when the client has no C_Timer", function()
		start()
		_G.C_Timer = nil
		local ran = false
		assert.is_false(Theme.after(2, function() ran = true end))
		assert.is_false(ran)
	end)

	it("glows through the client's overlay when it has one", function()
		local glowed, ungl = {}, {}
		start({ globals = {
			ActionButton_ShowOverlayGlow = function(button)
				glowed[#glowed + 1] = button
			end,
			ActionButton_HideOverlayGlow = function(button)
				ungl[#ungl + 1] = button
			end,
		} })
		local button = _G.CreateFrame("Button")
		assert.are.equal("overlay", Theme.showGlow(button))
		assert.are.same({ button }, glowed)
		assert.are.equal("overlay", Theme.hideGlow(button))
		assert.are.same({ button }, ungl)
		assert.are.same({}, Theme.diagnostics())
	end)

	it("will not use an overlay glow it cannot take off again", function()
		local glowed = {}
		start({ globals = { ActionButton_ShowOverlayGlow = function(button)
			glowed[#glowed + 1] = button
		end } })
		local button = _G.CreateFrame("Button")
		-- Half a pair is no pair: a client that can put the overlay on but
		-- never take it off would strand a glow on a button forever, which
		-- is worse than no glow, because the player acts on it. The addon's
		-- own outline is used instead -- it can always be removed.
		assert.are.equal("texture", Theme.showGlow(button))
		assert.are.same({}, glowed)
		assert.are.equal("texture", Theme.hideGlow(button))
		-- Both calls consulted the pair; the miss is recorded once.
		assert.are.same({ string.format(L.diagNoTemplate, "ActionButton_HideOverlayGlow") },
			Theme.diagnostics())
	end)

	it("will not use an overlay glow it cannot put on", function()
		local ungl = {}
		start({ globals = { ActionButton_HideOverlayGlow = function(button)
			ungl[#ungl + 1] = button
		end } })
		local button = _G.CreateFrame("Button")
		assert.are.equal("texture", Theme.showGlow(button))
		assert.are.equal("texture", Theme.hideGlow(button))
		assert.are.same({}, ungl)
		assert.are.same({ string.format(L.diagNoTemplate, "ActionButton_ShowOverlayGlow") },
			Theme.diagnostics())
	end)

	it("draws its own gold outline when the client has no overlay glow", function()
		start()
		local button = _G.CreateFrame("Button")
		assert.are.equal("texture", Theme.showGlow(button))
		assert.are.equal(#Theme.EDGES, #button.foreverSixtyGlow)
		for _, edge in ipairs(button.foreverSixtyGlow) do
			-- A texture is already shown at creation, so asserting
			-- edge.shown alone would pass even if setShown(edges, true)
			-- were never called; assert the recorded call instead.
			assert.is_true(mock.countCalls(edge, "Show") > 0)
		end
		assert.are.equal("texture", Theme.hideGlow(button))
		for _, edge in ipairs(button.foreverSixtyGlow) do
			assert.is_false(edge.shown)
		end
		-- Neither half present is the ordinary case, not a half-pair, so
		-- it stays silent rather than filling /fs diag with noise.
		assert.are.same({}, Theme.diagnostics())
	end)

	it("paints through SetTexture on a client with no SetColorTexture", function()
		start({ missingMethods = { SetColorTexture = true } })
		local frame = _G.CreateFrame("Frame")
		local texture = frame:CreateTexture()
		Theme.paint(texture, "gold")
		assert.is_nil(mock.firstCall(texture, "SetColorTexture"))
		assert.is_not_nil(mock.firstCall(texture, "SetTexture"))
	end)

	it("does nothing on a client with no GameTooltip at all", function()
		start()
		_G.GameTooltip = nil
		assert.is_false(Theme.showItemTooltip(_G.UIParent, 123, nil))
		assert.is_false(Theme.hideTooltip())
	end)

	it("shows a seen item's own tooltip via SetHyperlink when there is a link", function()
		start()
		local owner = _G.UIParent
		assert.is_true(Theme.showItemTooltip(owner, nil, "|Hitem:1234|h"))
		assert.are.same({ method = "SetHyperlink", n = 1, "|Hitem:1234|h" },
			mock.firstCall(_G.GameTooltip, "SetHyperlink"))
		assert.are.equal(1, mock.countCalls(_G.GameTooltip, "Show"))
	end)

	it("shows a planned item with no link yet via SetItemByID", function()
		start()
		local owner = _G.UIParent
		assert.is_true(Theme.showItemTooltip(owner, 1234, nil))
		assert.are.same({ method = "SetItemByID", n = 1, 1234 },
			mock.firstCall(_G.GameTooltip, "SetItemByID"))
	end)

	it("records and refuses a planned item when the client has no SetItemByID", function()
		start({ missingMethods = { SetItemByID = true } })
		local owner = _G.UIParent
		assert.is_false(Theme.showItemTooltip(owner, 1234, nil))
		assert.are.same(
			{ string.format(L.diagNoTooltipApi, "SetItemByID") },
			Theme.diagnostics())
	end)

	it("hides the tooltip through the client's own GameTooltip", function()
		start()
		assert.is_true(Theme.hideTooltip())
		assert.are.equal(1, mock.countCalls(_G.GameTooltip, "Hide"))
	end)

	it("shows a plain multi-line tooltip, one AddLine per line", function()
		start()
		local owner = _G.UIParent
		local lines = { "ForeverSixty", "Deep Holy", "Left-click", "Right-click" }
		assert.is_true(Theme.showLines(owner, lines))
		assert.are.same({ method = "SetOwner", n = 2, owner, "ANCHOR_LEFT" },
			mock.firstCall(_G.GameTooltip, "SetOwner"))
		assert.are.equal(1, mock.countCalls(_G.GameTooltip, "ClearLines"))
		assert.are.equal(#lines, mock.countCalls(_G.GameTooltip, "AddLine"))
		assert.are.equal(1, mock.countCalls(_G.GameTooltip, "Show"))
	end)

	it("does nothing for showLines on a client with no GameTooltip at all", function()
		start()
		_G.GameTooltip = nil
		assert.is_false(Theme.showLines(_G.UIParent, { "ForeverSixty" }))
	end)

	it("forgets what it learned about the client on reset", function()
		start({ templates = {} })
		Theme.hasTemplate("Button", "UIPanelButtonTemplate")
		assert.are.equal(1, #Theme.diagnostics())
		Theme.reset()
		assert.are.same({}, Theme.diagnostics())
		assert.are.same({}, Theme.templates)
	end)
end)
