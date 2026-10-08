import pytest

from pipeline import savedvars


def test_reads_a_wow_style_file():
    text = """
ForeverSixtyDB = {
	["savedAt"] = "2026-10-08 12:00",
	["recorder"] = {
		["start"] = 1, -- comment
		["events"] = {
			{
				["k"] = "pw",
				["v"] = -1.5e1,
			}, -- [1]
			{
				["k"] = "sw",
				["ok"] = true,
				["gone"] = false,
			}, -- [2]
		},
		["empty"] = {},
	},
}
ForeverSixtyInbox = nil
"""
    parsed = savedvars.parse(text)
    db = parsed["ForeverSixtyDB"]
    assert db["savedAt"] == "2026-10-08 12:00"
    assert db["recorder"]["events"] == [
        {"k": "pw", "v": -15.0},
        {"k": "sw", "ok": True, "gone": False},
    ]
    assert db["recorder"]["empty"] == []
    assert parsed["ForeverSixtyInbox"] is None


def test_reads_strings_numbers_and_mixed_keys():
    parsed = savedvars.parse(
        r"""X = { "a\"b\\c\n\65", 0x10, [3] = 7, name = 'q', [1.5] = 2 } --[[ block ]] Y = 1"""
    )
    assert parsed["Y"] == 1
    table = parsed["X"]
    assert table[1] == 'a"b\\c\nA'
    assert table[2] == 16
    assert table[3] == 7
    assert table["name"] == "q"
    assert table[1.5] == 2


def test_dense_integer_keys_become_a_list_and_gaps_a_dict():
    assert savedvars.parse("T = { [1] = 5, [2] = 6 }")["T"] == [5, 6]
    assert savedvars.parse("T = { [1] = 5, [3] = 6 }")["T"] == {1: 5, 3: 6}


@pytest.mark.parametrize(
    "text",
    ["X = ", "X = { 1, ", "X = $", "= 1", "X = { [1] 2 }", "X"],
)
def test_rejects_what_is_not_saved_variables(text):
    with pytest.raises(savedvars.SavedVariablesError):
        savedvars.parse(text)


def test_read_opens_a_file(tmp_path):
    path = tmp_path / "ForeverSixty.lua"
    path.write_text("A = 1\n", encoding="utf-8")
    assert savedvars.read(path) == {"A": 1}
