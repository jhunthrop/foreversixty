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

	it("heads only the broken things, then the chip line and the findings outside the heading", function()
		Theme.note("a broken thing")
		Theme.noteInfo("a finding")
		assert.are.same({ "Not working:", "a broken thing", "Chip: hidden in combat", "a finding" },
			Theme.diagReport(Theme.diagnostics(), "Chip: hidden in combat", Theme.infoNotes()))
	end)

	it("leaves the chip line out before the chip has drawn", function()
		assert.are.same({ "Everything is working." }, Theme.diagReport({}, nil, {}))
	end)

	it("forgets both lists on reset", function()
		Theme.note("x")
		Theme.noteInfo("y")
		Theme.reset()
		assert.are.same({}, Theme.diagnostics())
		assert.are.same({}, Theme.infoNotes())
	end)
end)
