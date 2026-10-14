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
		local ui = Prefs.withDefaults({ tooltip = false, minimap = { angle = 45 } })
		assert.is_false(ui.tooltip)
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
		local ui = Prefs.withDefaults({ tooltip = "yes", chip = 7 })
		assert.is_true(ui.tooltip)
		assert.are.equal(Prefs.DEFAULTS.chip.point, ui.chip.point)
	end)

	it("fills a missing nested key without dropping its siblings", function()
		local ui = Prefs.withDefaults({ chip = { x = 120 } })
		assert.are.equal(120, ui.chip.x)
		assert.are.equal(Prefs.DEFAULTS.chip.y, ui.chip.y)
		assert.are.equal(Prefs.DEFAULTS.chip.combatHide, ui.chip.combatHide)
	end)

	it("drops a key the defaults do not carry", function()
		local ui = Prefs.withDefaults({ leftover = "from an older addon" })
		assert.is_nil(ui.leftover)
	end)

	it("does not mutate the saved table it was handed", function()
		local saved = { chip = { x = 120 } }
		Prefs.withDefaults(saved)
		assert.are.same({ chip = { x = 120 } }, saved)
	end)

	it("does not hand out the DEFAULTS table itself", function()
		local ui = Prefs.withDefaults(nil)
		assert.are_not.equal(Prefs.DEFAULTS, ui)
		assert.are_not.equal(Prefs.DEFAULTS.chip, ui.chip)
	end)

	it("creates ForeverSixtyDB.ui on the first read", function()
		assert.is_nil(_G.ForeverSixtyDB)
		local ui = Prefs.current()
		assert.are.equal(ui, _G.ForeverSixtyDB.ui)
		assert.are.equal(ui, Prefs.current())
	end)

	it("round-trips a flag, including false", function()
		Prefs.setFlag("autoSave", false)
		assert.is_false(Prefs.flag("autoSave"))
		Prefs.setFlag("autoSave", true)
		assert.is_true(Prefs.flag("autoSave"))
	end)

	it("round-trips a field inside a section", function()
		Prefs.set("chip", "combatHide", false)
		assert.is_false(Prefs.get("chip", "combatHide"))
		assert.is_false(_G.ForeverSixtyDB.ui.chip.combatHide)
	end)

	it("resets placements and leaves the player's other choices alone", function()
		Prefs.set("chip", "point", "TOPRIGHT")
		Prefs.set("chip", "x", 300)
		Prefs.set("chip", "y", -40)
		Prefs.set("chip", "combatHide", false)
		Prefs.set("chip", "hiddenUntilLevel", 31)
		Prefs.set("chip", "pathChosen", "warrior-fury")
		Prefs.set("minimap", "angle", 12)
		Prefs.set("minimap", "shown", false)

		Prefs.resetPositions()

		-- An empty point is the chip's minimap-relative default anchor.
		assert.are.equal("", Prefs.get("chip", "point"))
		assert.are.equal(Prefs.DEFAULTS.chip.x, Prefs.get("chip", "x"))
		assert.are.equal(Prefs.DEFAULTS.chip.y, Prefs.get("chip", "y"))
		assert.are.equal(Prefs.DEFAULTS.minimap.angle, Prefs.get("minimap", "angle"))
		-- Not placements: a reset must not change how the chip behaves in
		-- combat, bring back a hidden chip, forget a path, or re-show the
		-- minimap button.
		assert.is_false(Prefs.get("chip", "combatHide"))
		assert.are.equal(31, Prefs.get("chip", "hiddenUntilLevel"))
		assert.are.equal("warrior-fury", Prefs.get("chip", "pathChosen"))
		assert.is_false(Prefs.get("minimap", "shown"))
	end)

	it("preserves nested table identity across an unrelated accessor call", function()
		-- A caller holding a reference to a nested section (e.g., a drag
		-- handler holding `local chip = Prefs.get("chip")`) must still
		-- write into the live table after any other accessor call.
		local chip = Prefs.get("chip")
		Prefs.setFlag("autoSave", false)
		assert.are.equal(chip, Prefs.get("chip"))
		chip.x = 42
		assert.are.equal(42, Prefs.get("chip", "x"))
	end)

	it("normalises a saved table with junk and wrong types in place", function()
		-- A reload that finds stale data (old keys, wrongly-typed sections)
		-- must normalise in place: the table object the code holds is still
		-- the same table, and now has the right shape.
		_G.ForeverSixtyDB = {
			ui = {
				chip = { x = 100, y = 50 },
				minimap = 7,
				oldkey = "junk",
			}
		}
		local ui = Prefs.current()
		assert.are.equal(100, ui.chip.x)
		assert.are.equal(Prefs.DEFAULTS.chip.point, ui.chip.point)
		assert.are.equal("table", type(ui.minimap))
		assert.is_nil(ui.oldkey)
		-- The same table object, not a new one.
		assert.are.equal(ui, _G.ForeverSixtyDB.ui)
	end)

	it("shows gear tips by default", function()
		assert.is_true(Prefs.DEFAULTS.tooltip)
		assert.is_true(Prefs.flag("tooltip"))
	end)

	it("round-trips the tooltip flag, including false", function()
		Prefs.setFlag("tooltip", false)
		assert.is_false(Prefs.flag("tooltip"))
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
		-- An empty point is the minimap-relative default anchor; 0 and "" are
		-- "not hidden" and "no path chosen" (a saved nil would be dropped).
		assert.are.same({
			shown = true, combatHide = true, point = "", x = 0, y = 0,
			hiddenUntilLevel = 0, hiddenFor = "", pathChosen = "", pathFor = "",
			byCharacter = {},
		}, ui.chip)
	end)

	it("puts the chip back beside the minimap on Reset positions", function()
		assert.are.same({ "chip", "minimap", "dialog" }, Prefs.POSITION_SECTIONS)
	end)

	it("opens the build dialog in the middle of the screen until it is dragged, and again on Reset", function()
		-- L3.4: an empty point is the centre of the screen; a drag saves CENTER and the offset.
		assert.are.same({ point = "", x = 0, y = 0 }, Prefs.withDefaults(nil).dialog)
		_G.ForeverSixtyDB = { ui = { dialog = { point = "CENTER", x = 40, y = -12 } } }
		Prefs.resetPositions()
		assert.are.same({ point = "", x = 0, y = 0 }, _G.ForeverSixtyDB.ui.dialog)
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
		for _, key in ipairs({ "chip", "minimap", "tooltip", "autoSave", "advancedDetail" }) do
			assert.is_not_nil(ui[key], key)
		end
	end)

	it("drops the tracker and the toast now the chip replaces them", function()
		local ui = Prefs.withDefaults({ tracker = { x = 120 }, toast = false })
		assert.is_nil(ui.tracker)
		assert.is_nil(ui.toast)
	end)

	it("drops the window's placement now the window is gone", function()
		assert.is_nil(Prefs.withDefaults({ window = { x = 120 } }).window)
	end)
end)
