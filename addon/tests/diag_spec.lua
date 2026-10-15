-- addon/tests/diag_spec.lua
-- /fs diag (L3.1 fix round 3): "Not working:" heads only what is broken;
-- otherwise "Everything is working."; the chip's own line and the
-- informational findings print after it, outside the heading. Pure: Theme's
-- note lists and its report builder, no client mock.
local helper = require("spec_helper")

describe("Theme.diagReport", function()
	local Theme, L

	before_each(function()
		L = helper.load("Locale")
		Theme = helper.load("Theme")
		Theme.reset()
	end)

	it("says everything is working when nothing is broken", function()
		assert.are.same({ "Everything is working.", "Chip: shown, nothing cut off" },
			Theme.diagReport(Theme.diagnostics(), "Chip: shown, nothing cut off", Theme.infoNotes()))
		assert.are.equal("Everything is working.", L.diagAllWorking)
	end)

	-- L4.1 fix round 5: the status lines get their own "Working:" heading under the faults,
	-- never read as one of them.
	it("heads only the broken things, then the status lines under Working:, then the findings", function()
		Theme.note("a broken thing")
		Theme.noteInfo("a finding")
		assert.are.same({ "Not working:", "a broken thing", "Working:", "Chip: hidden in combat", "a finding" },
			Theme.diagReport(Theme.diagnostics(), "Chip: hidden in combat", Theme.infoNotes()))
		assert.are.same({ "Not working:", "a broken thing", "Working:", "Build: x", "Bag marks: y", "a finding" },
			Theme.diagReport(Theme.diagnostics(), { "Build: x", "Bag marks: y" }, Theme.infoNotes()))
		assert.are.same({ "Not working:", "a broken thing", "a finding" },
			Theme.diagReport(Theme.diagnostics(), {}, Theme.infoNotes()))
		assert.are.equal("Working:", L.diagWorkingHeading)
	end)

	it("leaves the chip line out before the chip has drawn", function()
		assert.are.same({ "Everything is working." }, Theme.diagReport({}, nil, {}))
	end)

	it("says everything is working above a finding such as a missing data addon", function()
		Theme.noteInfo(L.diagDataAddonMissingInfo)
		assert.are.same({ "Everything is working.", L.diagDataAddonMissingInfo },
			Theme.diagReport(Theme.diagnostics(), nil, Theme.infoNotes()))
	end)

	-- Fix round 4: the path card registers its click-outside once, and a client
	-- that refuses it gets one fault line saying to press Escape.
	it("tries an event without noting anything, so the caller words the refusal", function()
		local refusing = { RegisterEvent = function() error("unknown event") end }
		assert.is_false(Theme.tryRegisterEvent(refusing, "GLOBAL_MOUSE_DOWN"))
		local taking = { RegisterEvent = function() end }
		assert.is_true(Theme.tryRegisterEvent(taking, "GLOBAL_MOUSE_DOWN"))
		assert.are.same({}, Theme.diagnostics())
		assert.are.same({}, Theme.infoNotes())
		assert.are.equal("Clicking outside the path card does not close it on this client; press Escape.",
			L.diagNoClickOutside)
	end)

	it("forgets both lists on reset", function()
		Theme.note("x")
		Theme.noteInfo("y")
		Theme.reset()
		assert.are.same({}, Theme.diagnostics())
		assert.are.same({}, Theme.infoNotes())
	end)
end)
