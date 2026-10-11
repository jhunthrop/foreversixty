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

	-- Design section 4: "off until the account has a level-60 export".
	-- Prefs itself does not know about exports; it only owns the default,
	-- which is off.
	it("defaults advancedDetail to off", function()
		assert.is_false(Prefs.DEFAULTS.advancedDetail)
		assert.is_false(Prefs.flag("advancedDetail"))
	end)

	it("remembers advancedDetail once the player turns it on", function()
		Prefs.setFlag("advancedDetail", true)
		assert.is_true(Prefs.flag("advancedDetail"))
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
		local ui = Prefs.withDefaults({ tracker = { x = 120 } })
		assert.are.equal(120, ui.tracker.x)
		assert.are.equal(Prefs.DEFAULTS.tracker.y, ui.tracker.y)
		assert.are.equal(Prefs.DEFAULTS.tracker.locked, ui.tracker.locked)
	end)

	it("drops a key the defaults do not carry", function()
		local ui = Prefs.withDefaults({ leftover = "from an older addon" })
		assert.is_nil(ui.leftover)
	end)

	it("does not mutate the saved table it was handed", function()
		local saved = { tracker = { x = 120 } }
		Prefs.withDefaults(saved)
		assert.are.same({ tracker = { x = 120 } }, saved)
	end)

	it("does not hand out the DEFAULTS table itself", function()
		local ui = Prefs.withDefaults(nil)
		assert.are_not.equal(Prefs.DEFAULTS, ui)
		assert.are_not.equal(Prefs.DEFAULTS.tracker, ui.tracker)
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
		Prefs.set("tracker", "locked", true)
		assert.is_true(Prefs.get("tracker", "locked"))
		assert.is_true(_G.ForeverSixtyDB.ui.tracker.locked)
	end)

	it("resets placements and leaves the player's other choices alone", function()
		Prefs.set("tracker", "x", 300)
		Prefs.set("tracker", "y", -40)
		Prefs.set("tracker", "locked", true)
		Prefs.set("minimap", "angle", 12)
		Prefs.set("minimap", "shown", false)

		Prefs.resetPositions()

		assert.are.equal(Prefs.DEFAULTS.tracker.x, Prefs.get("tracker", "x"))
		assert.are.equal(Prefs.DEFAULTS.tracker.y, Prefs.get("tracker", "y"))
		assert.are.equal(Prefs.DEFAULTS.minimap.angle, Prefs.get("minimap", "angle"))
		-- Not placements: a reset must not un-lock the tracker or re-show the
		-- minimap button.
		assert.is_true(Prefs.get("tracker", "locked"))
		assert.is_false(Prefs.get("minimap", "shown"))
	end)

	it("preserves nested table identity across an unrelated accessor call", function()
		-- A caller holding a reference to a nested section (e.g., a drag
		-- handler holding `local tracker = Prefs.get("tracker")`) must still
		-- write into the live table after any other accessor call.
		local tracker = Prefs.get("tracker")
		Prefs.setFlag("chat", true)
		assert.are.equal(tracker, Prefs.get("tracker"))
		tracker.x = 42
		assert.are.equal(42, Prefs.get("tracker", "x"))
	end)

	it("normalises a saved table with junk and wrong types in place", function()
		-- A reload that finds stale data (old keys, wrongly-typed sections)
		-- must normalise in place: the table object the code holds is still
		-- the same table, and now has the right shape.
		_G.ForeverSixtyDB = {
			ui = {
				tracker = { x = 100, y = 50 },
				minimap = 7,
				oldkey = "junk",
			}
		}
		local ui = Prefs.current()
		assert.are.equal(100, ui.tracker.x)
		assert.are.equal(Prefs.DEFAULTS.tracker.point, ui.tracker.point)
		assert.are.equal("table", type(ui.minimap))
		assert.is_nil(ui.oldkey)
		-- The same table object, not a new one.
		assert.are.equal(ui, _G.ForeverSixtyDB.ui)
	end)

	it("shows gear tips and the level-up toast by default", function()
		assert.is_true(Prefs.DEFAULTS.tooltip)
		assert.is_true(Prefs.DEFAULTS.toast)
		assert.is_true(Prefs.flag("tooltip"))
		assert.is_true(Prefs.flag("toast"))
	end)

	it("round-trips the tooltip and toast flags, including false", function()
		Prefs.setFlag("tooltip", false)
		assert.is_false(Prefs.flag("tooltip"))
		Prefs.setFlag("toast", false)
		assert.is_false(Prefs.flag("toast"))
	end)
end)

-- The v2 keys (L2.6): the build chip, the guild badges and the hello broadcast. Pure: these
-- read withDefaults only, never the client mock.
describe("Prefs, the v2 keys", function()
	local Prefs

	before_each(function()
		Prefs = helper.load("Prefs")
	end)

	it("shows the chip and hides it in combat by default", function()
		local ui = Prefs.withDefaults(nil)
		assert.are.same({ shown = true, combatHide = true }, ui.chip)
	end)

	it("turns the badges and the broadcast on by default", function()
		local ui = Prefs.withDefaults(nil)
		assert.is_true(ui.badges)
		assert.is_true(ui.broadcast)
	end)

	it("keeps the player's own chip, badge and broadcast choices", function()
		local ui = Prefs.withDefaults({ chip = { combatHide = false }, badges = false, broadcast = false })
		assert.is_true(ui.chip.shown)
		assert.is_false(ui.chip.combatHide)
		assert.is_false(ui.badges)
		assert.is_false(ui.broadcast)
	end)

	it("keeps the old keys the surfaces still read", function()
		local ui = Prefs.withDefaults(nil)
		for _, key in ipairs({ "tracker", "minimap", "tooltip", "toast", "autoSave", "chat", "advancedDetail" }) do
			assert.is_not_nil(ui[key], key)
		end
	end)

	it("drops the window's placement now the window is gone", function()
		assert.is_nil(Prefs.withDefaults({ window = { x = 120 } }).window)
	end)
end)
