"""The client's reputation factions, named and sided (sim-accuracy round 9,
items lane): Darkspear Islands' two reward sets are Darkspear Raiders 2798
(Horde) and Theramore Expeditionary Force 2799 (Alliance) reputation
rewards, and the Faction table is where the client says which side earns
each one."""

from pipeline.normalize.factions import FactionInfo, load_factions

# Faction rows as wago exports them: the race mask is a 64-bit value split
# into two signed 32-bit halves. Mask 0 is the side that earns the
# reputation; mask 1 starts at Hated, -42000, and never moves.
ALLIANCE_LOW, ALLIANCE_HIGH = "-1321907123", "1427461461"
HORDE_LOW, HORDE_HIGH = "1309324210", "-1440044374"


def faction(faction_id, name, earners=None, hated=None, base0="0"):
    row = {"ID": str(faction_id), "Name_lang": name}
    for index in range(4):
        row[f"ReputationRaceMasks{index}_0"] = "0"
        row[f"ReputationRaceMasks{index}_1"] = "0"
        row[f"ReputationBase_{index}"] = "0"
    row["ReputationBase_0"] = base0
    if earners:
        row["ReputationRaceMasks0_0"], row["ReputationRaceMasks0_1"] = earners
    if hated:
        row["ReputationRaceMasks1_0"], row["ReputationRaceMasks1_1"] = hated
        row["ReputationBase_1"] = "-42000"
    return row


# ChrRaces rows: a race's mask bit is its PlayableRaceBit, not its id.
# Dark Iron Dwarf 34 is an Alliance race on bit 11; Fel Orc 12 has no bit.
RACES = [
    {"ID": str(race_id), "PlayableRaceBit": str(race_id - 1), "Alliance": str(side)}
    for race_id, side in ((1, 0), (2, 1), (3, 0), (4, 0), (5, 1), (6, 1), (7, 0), (8, 1))
] + [
    {"ID": "34", "PlayableRaceBit": "11", "Alliance": "0"},
    {"ID": "12", "PlayableRaceBit": "-1", "Alliance": "2"},
]


def test_a_side_reputation_takes_the_side_of_the_races_that_earn_it():
    factions = load_factions(
        [
            faction(
                2798, "Darkspear Raiders", (HORDE_LOW, HORDE_HIGH), (ALLIANCE_LOW, ALLIANCE_HIGH)
            ),
            faction(
                2799,
                "Theramore Expeditionary Force",
                (ALLIANCE_LOW, ALLIANCE_HIGH),
                (HORDE_LOW, HORDE_HIGH),
            ),
        ],
        RACES,
    )
    assert factions[2798] == FactionInfo(name="Darkspear Raiders", side="horde")
    assert factions[2799] == FactionInfo(name="Theramore Expeditionary Force", side="alliance")


def test_a_reputation_both_sides_earn_has_no_side():
    factions = load_factions([faction(576, "Timbermaw Hold", base0="-3500")], RACES)
    assert factions[576] == FactionInfo(name="Timbermaw Hold", side=None)


def test_no_faction_table_is_no_factions():
    assert load_factions(None, RACES) == {}


def test_a_race_counts_by_its_playable_bit_not_its_id():
    """Bit 11 is Dark Iron Dwarf (race 34, Alliance), not race 12."""
    factions = load_factions([faction(9001, "Ironforge Delvers", (str(1 << 11), "0"))], RACES)
    assert factions[9001].side == "alliance"
