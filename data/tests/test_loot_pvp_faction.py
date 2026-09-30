# data/tests/test_loot_pvp_faction.py
"""pvp-faction lane, 2026-09-29: every `pvp:rank-N` `LootSource` splits
into its alliance/horde copies -- see `pipeline.loot.pvp_faction`'s own
doc for the two-tier (classic-db vendor, then title) resolution this
closes the third wow-player sweep's cross-wired-faction defect with."""

import re
from pathlib import Path

from pipeline.classic_sources import ClassicDbSourceRecord
from pipeline.loot.pvp_faction import (
    ALLIANCE_TITLES,
    HORDE_TITLES,
    TITLE_FACTION,
    PvpFactionStats,
    rank_title,
    split_pvp_sources_by_faction,
    title_faction,
    vendor_npc_factions,
)
from pipeline.models import LootFile, LootSource


def test_title_faction_reads_an_unambiguous_rank_title():
    assert title_faction("Knight-Lieutenant's Pauldrons") == "alliance"
    assert title_faction("Blood Guard's Pauldrons") == "horde"
    assert title_faction("Sergeant Major's Cape") == "alliance"
    assert title_faction("First Sergeant's Silk Cuffs") == "horde"


def test_title_faction_reads_the_literal_faction_word_when_no_title_matches():
    assert title_faction("Insignia of the Alliance") == "alliance"
    assert title_faction("Insignia of the Horde") == "horde"


def test_title_faction_refuses_the_rank_both_sides_share():
    """Rank 3 is called "Sergeant" on both ladders -- `TITLE_FACTION`
    deliberately excludes it, so a bare "Sergeant's Cloak" resolves to
    neither side by title alone (the source's own vendor must settle it
    instead, see the "barracks-less npc" vendor test below)."""
    assert "Sergeant" not in TITLE_FACTION
    assert title_faction("Sergeant's Cloak") is None


def test_title_faction_matches_a_compound_title_before_its_shorter_prefix():
    assert title_faction("Master Sergeant's Insignia") == "alliance"
    assert title_faction("Senior Sergeant's Insignia") == "horde"


def test_title_faction_rejects_an_unrelated_name():
    assert title_faction("Worn Shortsword") is None


ZONE_ROWS = [
    {"id": 2918, "name": "Champions' Hall", "map_id": 449, "map_name": "Alliance PVP Barracks"},
    {"id": 2917, "name": "Hall of Legends", "map_id": 450, "map_name": "Horde PVP Barracks"},
    {"id": 1, "name": "Stormwind City", "map_id": 0, "map_name": "Eastern Kingdoms"},
    {"id": 2, "name": "Orgrimmar", "map_id": 1, "map_name": "Kalimdor"},
]


def _vendor(
    item_id: int, npc_id: int, name: str, map_id: int | None
) -> tuple[int, ClassicDbSourceRecord]:
    return item_id, ClassicDbSourceRecord(kind="vendor", npc_id=npc_id, name=name, map_id=map_id)


def test_vendor_npc_factions_resolves_a_barracks_npc_from_its_own_spawn_map():
    """Rank 18's own quartermasters both stand in their side's PvP hall
    (map 449 Alliance, map 450 Horde) -- classic-db's own npc_vendor
    row plus this build's own zones.json settle both, no title needed."""
    classic_sources: dict[int, list[ClassicDbSourceRecord]] = {}
    for item_id, record in [
        _vendor(18873, 12782, "Captain O'Neal", 449),
        _vendor(18852, 14581, "Sergeant Thunderhorn", 450),
    ]:
        classic_sources[item_id] = [record]
    factions = vendor_npc_factions(classic_sources, ZONE_ROWS, item_names={})
    assert factions == {12782: "alliance", 14581: "horde"}


def test_vendor_npc_factions_infers_a_barracks_less_npc_from_its_own_titled_catalog():
    """Officer Areyn and Sergeant Ba'sha both stand on the open
    continent (map 0/1, no "PvP Barracks" name either side), so their
    own map settles nothing -- but Areyn's OTHER item is unambiguously
    Alliance-titled ("Private's Tabard") and Ba'sha's is Horde-titled
    ("Scout's Tabard"), which resolves the vendor and, through it, the
    ambiguous "Sergeant's Cloak"/"Sergeant's Cape" pair neither vendor's
    map nor either item's own bare title could settle alone."""
    classic_sources: dict[int, list[ClassicDbSourceRecord]] = {}
    for item_id, record in [
        _vendor(15196, 12805, "Officer Areyn", 0),
        _vendor(16342, 12805, "Officer Areyn", 0),  # "Sergeant's Cape"
        _vendor(15197, 12799, "Sergeant Ba'sha", 1),
        _vendor(16341, 12799, "Sergeant Ba'sha", 1),  # "Sergeant's Cloak"
    ]:
        classic_sources[item_id] = [record]
    item_names = {
        15196: "Private's Tabard",
        16342: "Sergeant's Cape",
        15197: "Scout's Tabard",
        16341: "Sergeant's Cloak",
    }
    factions = vendor_npc_factions(classic_sources, ZONE_ROWS, item_names)
    assert factions == {12805: "alliance", 12799: "horde"}


def test_vendor_npc_factions_resolves_nothing_for_a_vendor_whose_titles_disagree():
    """A vendor whose own titled catalog contradicts itself (a fixture
    error, or two unrelated npcs sharing an id) resolves to neither
    side, logged rather than guessed at."""
    classic_sources: dict[int, list[ClassicDbSourceRecord]] = {
        1: [ClassicDbSourceRecord(kind="vendor", npc_id=999, name="Confused Vendor", map_id=0)],
        2: [ClassicDbSourceRecord(kind="vendor", npc_id=999, name="Confused Vendor", map_id=0)],
    }
    item_names = {1: "Private's Tabard", 2: "Scout's Tabard"}
    assert vendor_npc_factions(classic_sources, ZONE_ROWS, item_names) == {}


def _pvp_document(items: list[int], rank: int = 11) -> LootFile:
    return LootFile(
        sources=[
            LootSource(
                id=f"pvp:rank-{rank}", kind="pvp", name=f"Rank {rank}", rank=rank, items=items
            ),
            LootSource(id="quest", kind="quest", name="Quest", items=[900]),
        ]
    )


def test_split_pvp_sources_by_faction_splits_by_title_when_no_vendor_resolves():
    document = _pvp_document([16338, 16391])
    item_names = {16338: "Knight-Lieutenant's Pauldrons", 16391: "Blood Guard's Pauldrons"}
    result, stats = split_pvp_sources_by_faction(
        document, item_names, classic_sources=None, zone_rows=None
    )
    by_id = {s.id: s for s in result.sources}
    assert set(by_id) == {"pvp:rank-11:alliance", "pvp:rank-11:horde", "quest"}
    alliance = by_id["pvp:rank-11:alliance"]
    assert (alliance.name, alliance.rank, alliance.items) == ("Rank 11 (Alliance)", 11, [16338])
    assert (alliance.faction, alliance.faction_source) == ("alliance", "title")
    horde = by_id["pvp:rank-11:horde"]
    assert (horde.name, horde.rank, horde.items) == ("Rank 11 (Horde)", 11, [16391])
    assert (horde.faction, horde.faction_source) == ("horde", "title")
    assert stats == PvpFactionStats(resolved_vendor=0, resolved_title=2, unresolved=())


def test_split_pvp_sources_by_faction_prefers_the_vendor_over_the_title():
    """`16342` ("Sergeant's Cape") carries no title of its own
    (`test_title_faction_refuses_the_rank_both_sides_share`), so only
    the classic-db vendor tier -- Officer Areyn, resolved via his own
    other Alliance-titled sale -- can place it."""
    document = _pvp_document([16342, 15196], rank=7)
    item_names = {16342: "Sergeant's Cape", 15196: "Private's Tabard"}
    classic_sources = {
        16342: [ClassicDbSourceRecord(kind="vendor", npc_id=12805, name="Officer Areyn", map_id=0)],
        15196: [ClassicDbSourceRecord(kind="vendor", npc_id=12805, name="Officer Areyn", map_id=0)],
    }
    result, stats = split_pvp_sources_by_faction(document, item_names, classic_sources, ZONE_ROWS)
    by_id = {s.id: s for s in result.sources}
    assert "pvp:rank-7:horde" not in by_id
    alliance = by_id["pvp:rank-7:alliance"]
    assert sorted(alliance.items) == [15196, 16342]
    assert alliance.faction_source == "classic-db"
    assert stats.resolved_vendor == 2
    assert stats.resolved_title == 0


def test_split_pvp_sources_by_faction_drops_an_item_neither_tier_resolves():
    """A name with no rank title and no vendor record at all -- tenet
    8's "unverifiable is labelled or left out, never shown as fact" --
    is dropped from both split buckets and counted, never guessed onto
    either side."""
    document = _pvp_document([1, 16338])
    item_names = {1: "A Perfectly Ordinary Blade", 16338: "Knight-Lieutenant's Pauldrons"}
    result, stats = split_pvp_sources_by_faction(
        document, item_names, classic_sources=None, zone_rows=None
    )
    by_id = {s.id: s for s in result.sources}
    assert "pvp:rank-11:horde" not in by_id
    assert by_id["pvp:rank-11:alliance"].items == [16338]
    assert stats.unresolved == (1,)


def test_split_pvp_sources_by_faction_leaves_every_other_kind_untouched():
    document = _pvp_document([16338])
    item_names = {16338: "Knight-Lieutenant's Pauldrons"}
    result, _ = split_pvp_sources_by_faction(document, item_names, None, None)
    quest = next(s for s in result.sources if s.id == "quest")
    assert quest.items == [900]
    assert quest.faction is None


def test_rank_title_reads_the_right_ladder_position():
    # Blizzard's client RequiredPVPRank is the ladder position + 4:
    # ladder 1 (Private/Scout) is rank 5, ladder 14 (Grand Marshal/High
    # Warlord) is rank 18.
    assert rank_title("alliance", 5) == "Private"
    assert rank_title("alliance", 18) == "Grand Marshal"
    assert rank_title("horde", 5) == "Scout"
    assert rank_title("horde", 18) == "High Warlord"
    assert rank_title("alliance", 8) == "Master Sergeant"


def test_rank_title_is_none_outside_the_known_ladder():
    assert rank_title("alliance", 4) is None
    assert rank_title("horde", 19) is None


def test_split_pvp_sources_by_faction_writes_the_rank_title_onto_each_split_source():
    document = _pvp_document([16338, 16391], rank=11)
    item_names = {16338: "Knight-Lieutenant's Pauldrons", 16391: "Blood Guard's Pauldrons"}
    result, _ = split_pvp_sources_by_faction(
        document, item_names, classic_sources=None, zone_rows=None
    )
    by_id = {s.id: s for s in result.sources}
    assert by_id["pvp:rank-11:alliance"].title == "Knight-Lieutenant"
    assert by_id["pvp:rank-11:horde"].title == "Blood Guard"


def test_split_pvp_sources_by_faction_exposes_faction_on_a_matching_vendor_row():
    """The fourth wow-player sweep's own item 1: Captain O'Neal's vendor
    row sells the exact same rank-18 items pvp:rank-18 does -- once his
    own npc_id resolves via the classic-db barracks map, his row must
    carry that same faction so sim/cmd/leveling-bis's
    vendorInheritsPvpRankGate can tell which of (possibly several) pvp
    sources for a shared item id is really his own side's."""
    document = LootFile(
        sources=[
            LootSource(id="pvp:rank-18", kind="pvp", name="Rank 18", rank=18, items=[18873]),
            LootSource(
                id="vendor:12782", kind="vendor", name="Captain O'Neal", npc_id=12782, items=[18873]
            ),
            LootSource(id="vendor:1", kind="vendor", name="Ordinary Vendor", npc_id=1, items=[2]),
        ]
    )
    item_names = {18873: "Grand Marshal's Stave"}
    classic_sources = {
        18873: [
            ClassicDbSourceRecord(kind="vendor", npc_id=12782, name="Captain O'Neal", map_id=449)
        ]
    }
    result, _ = split_pvp_sources_by_faction(document, item_names, classic_sources, ZONE_ROWS)
    by_id = {s.id: s for s in result.sources}
    assert by_id["vendor:12782"].faction == "alliance"
    # An ordinary vendor never resolved as a rank quartermaster keeps no
    # faction at all -- this must not blanket-tag every vendor row.
    assert by_id["vendor:1"].faction is None


_WEB_COPY_PATH = (
    Path(__file__).resolve().parents[2] / "web" / "src" / "lib" / "bis" / "copy.ts"
)


def _read_web_title_tuple(name: str) -> tuple[str, ...]:
    text = _WEB_COPY_PATH.read_text()
    match = re.search(rf"{name}[^=]*=\s*\[(.*?)\]", text, re.DOTALL)
    assert match is not None, f"{name} not found in {_WEB_COPY_PATH}"
    titles = re.findall(r"'([^']+)'", match.group(1))
    return tuple(titles)


def test_split_pvp_sources_by_faction_backfills_a_vendor_item_missing_from_the_pvp_source():
    """The fourth wow-player sweep's own regen-check finding: Captain
    O'Neal sells Grand Marshal's Polearm (a Forever-exclusive rank-18
    weapon type vanilla Classic's own pvp:rank-18 item list never
    carried), which left it with no pvp source at all for Go's
    vendorInheritsPvpRankGate to inherit from -- an Alliance-only weapon
    an Alliance vendor was already correctly gated to, but with no Rank
    to cap it out of the default pick and no Side reaching Go at all.
    The backfill must add it to pvp:rank-18:alliance using the SAME
    unanimous-rank rule vendor_npc_factions already uses for faction."""
    document = LootFile(
        sources=[
            LootSource(id="pvp:rank-18", kind="pvp", name="Rank 18", rank=18, items=[18873]),
            LootSource(
                id="vendor:12782",
                kind="vendor",
                name="Captain O'Neal",
                npc_id=12782,
                items=[18873, 234570],
            ),
        ]
    )
    item_names = {18873: "Grand Marshal's Stave"}
    classic_sources = {
        18873: [
            ClassicDbSourceRecord(kind="vendor", npc_id=12782, name="Captain O'Neal", map_id=449)
        ]
    }
    result, stats = split_pvp_sources_by_faction(document, item_names, classic_sources, ZONE_ROWS)
    by_id = {s.id: s for s in result.sources}
    assert sorted(by_id["pvp:rank-18:alliance"].items) == [18873, 234570]
    assert by_id["vendor:12782"].faction == "alliance"
    assert stats.backfilled_from_vendor == 1


def test_split_pvp_sources_by_faction_never_backfills_when_the_vendors_own_items_span_ranks():
    """A vendor whose OWN classic items resolve to more than one rank
    (a fixture error, or an npc sharing an id across two real
    quartermasters) settles nothing for its unresolved item, same
    "left out, never guessed at" rule as everywhere else in this
    module."""
    document = LootFile(
        sources=[
            LootSource(id="pvp:rank-18", kind="pvp", name="Rank 18", rank=18, items=[18873]),
            LootSource(id="pvp:rank-11", kind="pvp", name="Rank 11", rank=11, items=[16338]),
            LootSource(
                id="vendor:12782",
                kind="vendor",
                name="Confused Quartermaster",
                npc_id=12782,
                items=[18873, 16338, 999999],
            ),
        ]
    )
    item_names = {18873: "Grand Marshal's Stave", 16338: "Knight-Lieutenant's Pauldrons"}
    confused_vendor_record = ClassicDbSourceRecord(
        kind="vendor", npc_id=12782, name="Confused Quartermaster", map_id=449
    )
    classic_sources = {18873: [confused_vendor_record], 16338: [confused_vendor_record]}
    result, stats = split_pvp_sources_by_faction(document, item_names, classic_sources, ZONE_ROWS)
    assert stats.backfilled_from_vendor == 0
    for source in result.sources:
        if source.kind == "pvp":
            assert 999999 not in (source.items or [])


def test_rank_title_tables_stay_in_sync_with_the_web_copy():
    """`pipeline.loot.pvp_faction`'s own `ALLIANCE_TITLES`/`HORDE_TITLES`
    is the primary source (this module's doc); `web/src/lib/bis/copy.ts`
    mirrors it for the sources this pipeline has not reached yet
    (module doc, "keep both in sync with a test that compares them").
    A change to one ladder without the other fails here instead of
    silently drifting the two apart."""
    assert _read_web_title_tuple("ALLIANCE_PVP_TITLES") == ALLIANCE_TITLES
    assert _read_web_title_tuple("HORDE_PVP_TITLES") == HORDE_TITLES
