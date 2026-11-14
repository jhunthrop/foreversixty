local helper = require("spec_helper")
local mock = require("wow_mock")

local ALL_TEMPLATES = { "PanelTabButtonTemplate", "UIPanelButtonTemplate", "InputBoxTemplate" }

describe("Widgets", function()
	local Theme, Widgets, state

	local function start(install)
		state = mock.install(install or {})
		Theme = helper.load("Theme")
		Theme.reset()
		package.loaded["Widgets"] = nil
		Widgets = require("Widgets")
		return state
	end

	after_each(function()
		mock.uninstall()
	end)

	it("builds no frame merely by being loaded", function()
		start({ templates = {} })
		assert.are.equal(0, #helper.builtFrames(state))
	end)

	it("draws a panel from a background and four border edges, with no template", function()
		start({ templates = {} })
		local panel = Widgets.panel(_G.UIParent, 200, 100)
		assert.are.equal(1 + #Theme.EDGES, #panel.regions)
		assert.are.equal(#Theme.EDGES, #panel.foreverSixtyBorder)
	end)

	it("gives the panel the global name a caller asked for", function()
		start({ templates = {} })
		local panel = Widgets.panel(_G.UIParent, 200, 100, "ForeverSixtyTestPanel")
		assert.are.equal(panel, state.widgets.ForeverSixtyTestPanel)
	end)

	it("colours a label with the palette entry it was asked for", function()
		start()
		local label = Widgets.label(_G.UIParent, "Forever Sixty", "gold")
		assert.are.equal("Forever Sixty", label:GetText())
		local call = mock.firstCall(label, "SetTextColor")
		assert.are.same({ Theme.rgb(Theme.HEX.gold) }, { call[1], call[2], call[3], call[4] })
	end)

	-- Found in game: the stock templates gave a red button, gold tabs and a
	-- one-line input, none of it the site's design.
	it("draws its own button, tab and field even on a client with the stock templates", function()
		start({ templates = ALL_TEMPLATES })
		local button = Widgets.clientButton(_G.UIParent, "Copy", function() end)
		local tab = Widgets.tab(_G.UIParent, "Export", function() end)
		local box = Widgets.editBox(_G.UIParent, 300, 72, true)
		assert.is_nil(button.template)
		assert.is_nil(tab.template)
		assert.is_nil(box.template)
		assert.are.equal("Copy", button:GetText())
	end)

	it("sets an edit box inside a field of exactly the size asked for", function()
		start()
		local box = Widgets.editBox(_G.UIParent, 300, 72, true)
		local field = Widgets.field(box)
		assert.are.equal(300, field.width)
		assert.are.equal(72, field.height)
		assert.is_true(box.width < field.width)
		assert.is_true(box.height < field.height)
	end)

	it("underlines the active tab only", function()
		start()
		local tab = Widgets.tab(_G.UIParent, "Export", function() end)
		assert.is_false(tab.foreverSixtyUnderline:IsShown())
		Widgets.setTabActive(tab, true)
		assert.is_true(tab.foreverSixtyUnderline:IsShown())
		Widgets.setTabActive(tab, false)
		assert.is_false(tab.foreverSixtyUnderline:IsShown())
	end)

	it("closes through the title bar's close button", function()
		start()
		local closed = 0
		local button = Widgets.closeButton(_G.UIParent, function() closed = closed + 1 end)
		button:GetScript("OnClick")(button)
		assert.are.equal(1, closed)
	end)

	it("ignores a click on a disabled button rather than acting on it", function()
		start({ templates = {} })
		local clicks = 0
		local button = Widgets.clientButton(_G.UIParent, "Equip", function() clicks = clicks + 1 end)
		Widgets.setEnabled(button, false)
		button.OnClick(button)
		assert.are.equal(0, clicks)
		assert.is_false(button:IsEnabled())
	end)

	it("dims a widget it turns off and brightens it again", function()
		start({ templates = {} })
		local button = Widgets.clientButton(_G.UIParent, "Equip", function() end)
		Widgets.setEnabled(button, false)
		assert.are.equal(Theme.ALPHA.disabled, mock.lastCall(button, "SetAlpha")[1])
		Widgets.setEnabled(button, true)
		assert.are.equal(1, mock.lastCall(button, "SetAlpha")[1])
		assert.is_true(button:IsEnabled())
	end)

	it("marks the active tab gold and the others muted", function()
		start({ templates = {} })
		local tab = Widgets.tab(_G.UIParent, "Follow", function() end)
		Widgets.setTabActive(tab, true)
		local gold = mock.lastCall(tab.foreverSixtyLabel, "SetTextColor")
		assert.are.same({ Theme.rgb(Theme.HEX.gold) }, { gold[1], gold[2], gold[3], gold[4] })
		assert.is_true(tab.foreverSixtyActive)
		Widgets.setTabActive(tab, false)
		local muted = mock.lastCall(tab.foreverSixtyLabel, "SetTextColor")
		assert.are.same({ Theme.rgb(Theme.HEX.muted) }, { muted[1], muted[2], muted[3], muted[4] })
		assert.is_false(tab.foreverSixtyActive)
	end)

	it("runs a tab's handler when it is clicked", function()
		start({ templates = {} })
		local clicked = false
		local tab = Widgets.tab(_G.UIParent, "Follow", function() clicked = true end)
		tab:GetScript("OnClick")(tab)
		assert.is_true(clicked)
	end)

	it("selects and focuses the text it puts in a read-only box", function()
		start({ templates = {} })
		local box = Widgets.editBox(_G.UIParent, 400, 40, true)
		Widgets.selectText(box, "FS1:1.60.1.69893:paladin")
		assert.are.equal("FS1:1.60.1.69893:paladin", box:GetText())
		assert.is_not_nil(mock.firstCall(box, "HighlightText"))
		assert.is_true(box.focused)
	end)

	it("puts a typed-over read-only box back the way it was", function()
		start({ templates = {} })
		local box = Widgets.editBox(_G.UIParent, 400, 40, true)
		Widgets.selectText(box, "FS1:code")
		box:SetText("the player typed this")
		box:GetScript("OnTextChanged")(box, true)
		assert.are.equal("FS1:code", box:GetText())
	end)

	it("lets the addon itself change a read-only box's text", function()
		start({ templates = {} })
		local box = Widgets.editBox(_G.UIParent, 400, 40, true)
		Widgets.selectText(box, "FS1:code")
		box:SetText("FS1:regenerated")
		box:GetScript("OnTextChanged")(box, false)
		assert.are.equal("FS1:regenerated", box:GetText())
	end)

	it("leaves a writable box alone", function()
		start({ templates = {} })
		local box = Widgets.editBox(_G.UIParent, 400, 40, false)
		assert.is_nil(box:GetScript("OnTextChanged"))
	end)

	it("reports a toggle's new value to its handler", function()
		start({ templates = {} })
		local seen
		local toggle = Widgets.toggle(_G.UIParent, "Show the tracker", false, function(value)
			seen = value
		end)
		assert.is_false(toggle.tick:IsShown())
		toggle.frame:GetScript("OnClick")(toggle.frame)
		assert.is_true(seen)
		assert.is_true(toggle.checked)
		assert.is_true(toggle.tick:IsShown())
	end)

	it("starts a toggle ticked when it was given a checked value", function()
		start({ templates = {} })
		local toggle = Widgets.toggle(_G.UIParent, "Show the tracker", true, function() end)
		assert.is_true(toggle.checked)
		assert.is_true(toggle.tick:IsShown())
	end)

	it("runs a text button's handler when clicked and sizes itself to the text", function()
		start({ templates = {} })
		local clicks = 0
		local button = Widgets.textButton(_G.UIParent, "+3 more", function() clicks = clicks + 1 end)
		assert.are.equal("+3 more", button:GetText())
		button:GetScript("OnClick")(button)
		assert.are.equal(1, clicks)
		button:SetText("Show fewer")
		assert.are.equal("Show fewer", button:GetText())
	end)

	it("ignores a click on a disabled text button", function()
		start({ templates = {} })
		local clicks = 0
		local button = Widgets.textButton(_G.UIParent, "+3 more", function() clicks = clicks + 1 end)
		button:Disable()
		button:GetScript("OnClick")(button)
		assert.are.equal(0, clicks)
	end)

	describe("keyboard focus", function()
		-- An edit box that holds focus swallows every keybind. It may take
		-- focus only when the player asks to copy, and must give it back.
		local function box()
			start()
			return Widgets.editBox(_G.CreateFrame("Frame"), 200, 80, true)
		end

		it("does not hold the keyboard merely for having been created", function()
			local b = box()
			assert.is_false(b.focused)
			assert.is_false(b.autoFocus)
		end)

		it("sets text without focus", function()
			local b = box()
			Widgets.setText(b, "FS1:code")
			assert.are.equal("FS1:code", b:GetText())
			assert.is_false(b.focused)
		end)

		it("gives focus back when the box is hidden", function()
			local b = box()
			Widgets.selectText(b, "FS1:code")
			assert.is_true(b.focused)
			b:GetScript("OnHide")(b)
			assert.is_false(b.focused)
		end)

		it("gives focus back on Escape and on Enter", function()
			local b = box()
			Widgets.selectText(b, "FS1:code")
			b:GetScript("OnEscapePressed")(b)
			assert.is_false(b.focused)
			Widgets.selectText(b, "FS1:code")
			b:GetScript("OnEnterPressed")(b)
			assert.is_false(b.focused)
		end)

		it("gives focus back shortly after the player copies with Ctrl+C", function()
			local b = box()
			Widgets.selectText(b, "FS1:code")
			_G.IsControlKeyDown = function()
				return true
			end
			b:GetScript("OnKeyDown")(b, "C")
			assert.are.equal(1, mock.runTimers(state))
			assert.is_false(b.focused)
			_G.IsControlKeyDown = nil
		end)

		it("keeps focus for any other key", function()
			local b = box()
			Widgets.selectText(b, "FS1:code")
			_G.IsControlKeyDown = function()
				return false
			end
			b:GetScript("OnKeyDown")(b, "C")
			mock.runTimers(state)
			assert.is_true(b.focused)
			_G.IsControlKeyDown = nil
		end)
	end)
	describe("a scrollable page", function()
		local function holderOf(height)
			start()
			local holder = _G.CreateFrame("Frame")
			holder:SetSize(400, height)
			return holder
		end

		it("does not scroll when the content fits", function()
			local scroll = Widgets.scrollable(holderOf(300))
			scroll:SetContentHeight(200)
			assert.are.equal(0, scroll:Range())
			assert.are.equal(0, scroll:ScrollTo(500))
		end)

		it("scrolls exactly as far as the content overflows, and no further", function()
			local scroll = Widgets.scrollable(holderOf(300))
			scroll:SetContentHeight(500)
			assert.are.equal(200, scroll:Range())
			assert.are.equal(200, scroll:ScrollTo(999))
			assert.are.equal(0, scroll:ScrollTo(-50))
		end)

		it("moves one step per wheel notch and clamps", function()
			local holder = holderOf(300)
			local scroll = Widgets.scrollable(holder)
			scroll:SetContentHeight(360)
			holder:GetScript("OnMouseWheel")(holder, -1)
			assert.are.equal(math.min(Widgets.SCROLL_STEP, 60), scroll.offset)
			holder:GetScript("OnMouseWheel")(holder, -1)
			assert.are.equal(60, scroll.offset)
			holder:GetScript("OnMouseWheel")(holder, 1)
			assert.are.equal(60 - Widgets.SCROLL_STEP, scroll.offset)
		end)

		it("clips the holder so scrolled-away rows do not draw over the page", function()
			local holder = holderOf(300)
			Widgets.scrollable(holder)
			local clipped = false
			for _, call in ipairs(holder.calls) do
				if call.method == "SetClipsChildren" and call[1] == true then
					clipped = true
				end
			end
			assert.is_true(clipped)
		end)
	end)

	-- The site's buttons on the Forever Sixty window.
	describe("the site's buttons", function()
		--- A texture's last colour, as six hex digits.
		local function paintOf(texture)
			local call = mock.lastCall(texture, "SetColorTexture")
			return call and Theme.hexOf(call[1], call[2], call[3]) or nil
		end

		it("draws the site's button, never the client's red template, and clicks only while enabled", function()
			start({ templates = ALL_TEMPLATES })
			local clicked = 0
			local made = Widgets.clientButton(_G.UIParent, "In", function()
				clicked = clicked + 1
			end)
			assert.is_nil(made.template)
			assert.are.equal("In", made:GetText())
			assert.are.equal(Theme.SIZES.clientButtonHeight, made:GetHeight())
			made.OnClick(made)
			made:Disable()
			made.OnClick(made)
			assert.are.equal(1, clicked)
		end)

		it("is the secondary recipe at rest and the primary while selected, a white label on both", function()
			start()
			local made = Widgets.clientButton(_G.UIParent, "Late", function() end)
			assert.are.equal(Theme.HEX.blue, paintOf(made.foreverSixtyGround))
			assert.are.equal(Theme.HEX.blueEdge, paintOf(made.foreverSixtyEdges[1]))
			assert.is_false(made.foreverSixtyTeal[1].shown)
			assert.is_true(Widgets.setSelected(made, true))
			for _, part in ipairs(made.foreverSixtyTeal) do
				assert.is_true(part.shown)
			end
			assert.are.equal(Theme.HEX.tealTop, paintOf(made.foreverSixtyTeal[1]))
			assert.are.equal(Theme.HEX.tealEdge, paintOf(made.foreverSixtyTeal[2]))
			local label = mock.lastCall(made.foreverSixtyLabel, "SetTextColor")
			assert.are.equal(Theme.HEX.strong, Theme.hexOf(label[1], label[2], label[3]))
			assert.is_false(Widgets.setSelected(made, false))
			assert.is_false(made.foreverSixtyTeal[1].shown)
		end)

		it("closes with the flat X, never the client's red one", function()
			start({ templates = ALL_TEMPLATES })
			local closed = 0
			local close = Widgets.closeButton(_G.UIParent, function()
				closed = closed + 1
			end)
			assert.is_nil(close.template)
			close.OnClick(close)
			assert.are.equal(1, closed)
		end)

	end)

	-- Spec 2026-11-13-addon-hub 2: the Forever Sixty window is the addon's one framed window. No other file builds a
	-- framed panel on the screen itself (the chip and the marks in the world are not panels).
	it("leaves the window the only framed panel on the screen", function()
		local makers = {}
		for line in io.lines("ForeverSixty/ForeverSixty.toc") do
			local path = line:match("^([%w\\_]+%.lua)$")
			if path ~= nil and path ~= "Data.lua" then
				local file = assert(io.open("ForeverSixty/" .. path:gsub("\\", "/"), "r"))
				local text = file:read("a")
				file:close()
				if text:find("Widgets%.panel%(%s*_?G?%.?UIParent") ~= nil then
					makers[#makers + 1] = path
				end
			end
		end
		assert.are.same({ "surfaces\\Hub.lua" }, makers)
	end)
end)
