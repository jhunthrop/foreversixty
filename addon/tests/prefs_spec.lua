local helper = require("spec_helper")
local mock = require("wow_mock")

describe("Prefs", function()
	local Prefs

	before_each(function()
		mock.install({})
		Prefs = helper.load("Prefs")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("hands back every default for a saved table that is not there yet", function()
		local ui = Prefs.withDefaults(nil)
		assert.are.same(Prefs.DEFAULTS, ui)
	end)

	it("keeps a value the player already chose", function()
		local ui = Prefs.withDefaults({ chat = true, minimap = { angle = 45 } })
		assert.is_true(ui.chat)
		assert.are.equal(45, ui.minimap.angle)
	end)

	it("keeps a saved false rather than treating it as missing", function()
		-- `given ~= nil` and not `given or default`: autoSave defaults to
		-- true, so a player who turned it off must stay turned off.
		local ui = Prefs.withDefaults({ autoSave = false })
		assert.is_false(ui.autoSave)
	end)

	it("replaces a saved value of the wrong type with the default", function()
		-- An older or hand-edited saved table must never reach a caller as
		-- the wrong type; that is what turns a pref read into a Lua error.
		local ui = Prefs.withDefaults({ chat = "yes", tracker = 7 })
		assert.is_false(ui.chat)
		assert.are.equal(Prefs.DEFAULTS.tracker.point, ui.tracker.point)
	end)

	it("fills a missing nested key without dropping its siblings", function()
		local ui = Prefs.withDefaults({ window = { x = 120 } })
		assert.are.equal(120, ui.window.x)
		assert.are.equal(Prefs.DEFAULTS.window.y, ui.window.y)
		assert.are.equal(Prefs.DEFAULTS.window.tab, ui.window.tab)
	end)

	it("drops a key the defaults do not carry", function()
		local ui = Prefs.withDefaults({ leftover = "from an older addon" })
		assert.is_nil(ui.leftover)
	end)

	it("does not mutate the saved table it was handed", function()
		local saved = { window = { x = 120 } }
		Prefs.withDefaults(saved)
		assert.are.same({ window = { x = 120 } }, saved)
	end)

	it("does not hand out the DEFAULTS table itself", function()
		local ui = Prefs.withDefaults(nil)
		assert.are_not.equal(Prefs.DEFAULTS, ui)
		assert.are_not.equal(Prefs.DEFAULTS.window, ui.window)
	end)

	it("creates ForeverSixtyDB.ui on the first read", function()
		assert.is_nil(_G.ForeverSixtyDB)
		local ui = Prefs.current()
		assert.are.equal(ui, _G.ForeverSixtyDB.ui)
		assert.are.equal(ui, Prefs.current())
	end)

	it("round-trips a flag, including false", function()
		Prefs.setFlag("chat", true)
		assert.is_true(Prefs.flag("chat"))
		Prefs.setFlag("chat", false)
		assert.is_false(Prefs.flag("chat"))
	end)

	it("round-trips a field inside a section", function()
		Prefs.set("window", "tab", "gear")
		assert.are.equal("gear", Prefs.get("window", "tab"))
		assert.are.equal("gear", _G.ForeverSixtyDB.ui.window.tab)
	end)

	it("resets placements and leaves the player's other choices alone", function()
		Prefs.set("window", "x", 300)
		Prefs.set("window", "tab", "gear")
		Prefs.set("tracker", "y", -40)
		Prefs.set("tracker", "locked", true)
		Prefs.set("minimap", "angle", 12)
		Prefs.set("minimap", "shown", false)

		Prefs.resetPositions()

		assert.are.equal(Prefs.DEFAULTS.window.x, Prefs.get("window", "x"))
		assert.are.equal(Prefs.DEFAULTS.tracker.y, Prefs.get("tracker", "y"))
		assert.are.equal(Prefs.DEFAULTS.minimap.angle, Prefs.get("minimap", "angle"))
		-- Not placements: a reset must not un-lock the tracker, re-show the
		-- minimap button or send the player back to the Export tab.
		assert.are.equal("gear", Prefs.get("window", "tab"))
		assert.is_true(Prefs.get("tracker", "locked"))
		assert.is_false(Prefs.get("minimap", "shown"))
	end)
end)
