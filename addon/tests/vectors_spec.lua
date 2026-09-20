describe("the shared codec fixture vectors", function()
	local function read(path)
		local file = assert(io.open(path, "r"), path .. " is missing")
		local text = file:read("a")
		file:close()
		return text
	end

	it("are byte-identical in both lanes", function()
		-- Two copies rather than one shared path because the web lane's
		-- bundler and the addon's zip each need the file inside their own
		-- tree. This is the gate that a regeneration reached both.
		assert.are.equal(
			read("tests/fixtures/codec-vectors.json"),
			read("../web/src/fixtures/addon/codec-vectors.json")
		)
	end)

	it("carry FS1 and FSB1 cases in both directions", function()
		local helper = require("spec_helper")
		local vectors = helper.vectors()
		assert.is_true(#vectors.fs1 >= 7)
		assert.is_true(#vectors.fs1Invalid >= 7)
		assert.is_true(#vectors.fsb1 >= 3)
		assert.is_true(#vectors.fsb1Invalid >= 6)
	end)
end)
