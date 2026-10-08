"""The recorder report, on a synthetic record written the way the addon writes it."""

from __future__ import annotations

import random
from pathlib import Path
from typing import Any

import pytest

from pipeline import recorder
from pipeline.__main__ import main
from pipeline.recorder_stats import SingularFitError


def lua(value: Any) -> str:
    if isinstance(value, bool):
        return "true" if value else "false"
    if isinstance(value, int | float):
        return repr(value)
    if isinstance(value, str):
        return '"' + value.replace("\\", "\\\\").replace('"', '\\"') + '"'
    if isinstance(value, list):
        return "{" + ",".join(lua(item) for item in value) + "}"
    return "{" + ",".join(f"[{lua(k)}]={lua(v)}" for k, v in value.items()) + "}"


def session(sid: int, **fields: Any) -> dict[str, Any]:
    return {"id": sid, "level": 60, "spirit": 100, "ap": 1000, "counts": {}, **fields}


def store(
    events: list[dict[str, Any]], sessions: list[dict[str, Any]], **extra: Any
) -> dict[str, Any]:
    return {
        "schema": 1,
        "cap": 20000,
        "start": 1,
        "count": len(events),
        "events": events,
        "sessions": sessions,
        "nextSession": len(sessions) + 1,
        **extra,
    }


def recording(
    events: list[dict[str, Any]], sessions: list[dict[str, Any]] | None = None
) -> recorder.Recording:
    sessions = sessions if sessions is not None else [session(1)]
    return recorder.Recording({s["id"]: s for s in sessions}, events, 20000)


def write(tmp_path: Path, recorder_store: dict[str, Any]) -> Path:
    path = tmp_path / "ForeverSixty.lua"
    path.write_text(f"ForeverSixtyDB = {lua({'recorder': recorder_store})}\n", encoding="utf-8")
    return path


# Loading -----------------------------------------------------------------------


def test_load_reorders_the_ring_buffer_oldest_first(tmp_path):
    physical = [{"k": "pw", "t": 4}, {"k": "pw", "t": 5}, {"k": "pw", "t": 3}]
    path = write(tmp_path, store(physical, [session(1)], start=3))
    loaded = recorder.load(path)
    assert [e["t"] for e in loaded.events] == [3, 4, 5]
    assert loaded.sessions[1]["spirit"] == 100
    assert loaded.capacity == 20000


def test_load_reads_an_empty_record(tmp_path):
    loaded = recorder.load(write(tmp_path, store([], [])))
    assert loaded.events == [] and loaded.sessions == {}
    assert "0 events" in recorder.report(loaded)


@pytest.mark.parametrize(
    ("text", "message"),
    [
        ("ForeverSixtyDB = { characters = {} }", "no ForeverSixtyDB.recorder"),
        ("Other = 1", "no ForeverSixtyDB.recorder"),
        ("ForeverSixtyDB = {recorder = {schema = 9}}", "schema 9"),
        ("ForeverSixtyDB = {", "not a SavedVariables file"),
    ],
)
def test_load_refuses_what_is_not_a_record(tmp_path, text, message):
    path = tmp_path / "x.lua"
    path.write_text(text, encoding="utf-8")
    with pytest.raises(recorder.RecorderError, match=message):
        recorder.load(path)


# Energy ------------------------------------------------------------------------


def energy_events(period: float, amount: int, ticks: int) -> list[dict[str, Any]]:
    events = []
    value = 20
    for index in range(ticks):
        value = min(100, value + amount)
        events.append(
            {"k": "pw", "n": 1, "p": "ENERGY", "t": 100 + index * period, "v": value, "mx": 100}
        )
        if value == 100:
            events.append(
                {
                    "k": "pw",
                    "n": 1,
                    "p": "ENERGY",
                    "t": 100 + index * period + 0.3,
                    "v": 55,
                    "mx": 100,
                }
            )
            value = 55
    return events


def test_energy_tick_reads_amount_and_period():
    tick = recorder.energy_tick(recording(energy_events(2.0, 20, 12)))
    assert tick is not None
    assert tick.amount == 20
    assert tick.period == pytest.approx(2.0)
    assert tick.period_count >= 5


def test_energy_gain_cut_off_at_the_cap_is_not_a_tick_amount():
    events = [
        {"k": "pw", "n": 1, "p": "ENERGY", "t": 1, "v": 90, "mx": 100},
        {"k": "pw", "n": 1, "p": "ENERGY", "t": 3, "v": 100, "mx": 100},
    ]
    assert recorder.energy_tick(recording(events)) is None


def test_energy_period_is_unknown_with_a_single_gain():
    events = [
        {"k": "pw", "n": 1, "p": "ENERGY", "t": 1, "v": 40, "mx": 100},
        {"k": "pw", "n": 1, "p": "ENERGY", "t": 3, "v": 60, "mx": 100},
    ]
    tick = recorder.energy_tick(recording(events))
    assert tick is not None and tick.period is None
    assert "no two consecutive ticks" in recorder.report(recording(events))


def test_energy_period_skips_a_tick_broken_by_another_gain():
    values = [(0, 20), (2, 40), (3, 45), (5, 65), (7, 85)]
    events = [{"k": "pw", "n": 1, "p": "ENERGY", "t": t, "v": v, "mx": 100} for t, v in values]
    tick = recorder.energy_tick(recording(events))
    assert tick is not None and tick.period_count == 1


# White swings ---------------------------------------------------------------------


def swings(hand: str, target: int, outcomes: dict[str, tuple[int, int]]) -> list[dict[str, Any]]:
    return [
        {"k": "sw", "n": 1, "t": 1, "h": hand, "o": outcome, "tl": target, "a": damage}
        for outcome, (count, damage) in outcomes.items()
        for _ in range(count)
    ]


def test_white_groups_rates_intervals_and_glancing_ratio():
    events = swings(
        "m", 63, {"hit": (50, 100), "crit": (10, 200), "glance": (20, 70), "miss": (20, 0)}
    )
    events += swings("o", 60, {"hit": (9, 40), "dodge": (1, 0)})
    groups = recorder.white_groups(recording(events))
    assert [(g.hand, g.level_gap, g.swings) for g in groups] == [("main", 3, 100), ("off", 0, 10)]
    main_group = groups[0]
    assert main_group.rates["glance"].rate == pytest.approx(0.2)
    low, high = main_group.rates["glance"].interval.low, main_group.rates["glance"].interval.high
    assert low < 0.2 < high
    assert main_group.glancing_ratio == pytest.approx(0.7)
    assert groups[1].glancing_ratio is None
    assert "dodge" in groups[1].rates and "glance" not in groups[1].rates


def test_white_groups_skip_swings_without_a_target_level_or_session():
    events = [
        {"k": "sw", "n": 1, "t": 1, "h": "m", "o": "hit", "a": 5},
        {"k": "sw", "n": 1, "t": 1, "h": "m", "o": "hit", "a": 5, "tl": -1},
        {"k": "sw", "n": 9, "t": 1, "h": "m", "o": "hit", "a": 5, "tl": 60},
    ]
    assert recorder.white_groups(recording(events)) == []
    assert "no white swings" in recorder.report(recording(events))


# Yellow attacks ----------------------------------------------------------------------


def eviscerates(rng: random.Random, base: float, per_point: float, ap_term: float, spread: float):
    events = []
    for index in range(40):
        points, ap = 1 + index % 5, 900 + 40 * (index % 7)
        damage = base + per_point * points + ap_term * ap * points + rng.uniform(-spread, spread)
        events.append(
            {
                "k": "ya",
                "n": 1,
                "t": index,
                "s": "Eviscerate",
                "o": "hit",
                "a": damage,
                "cp": points,
                "ap": ap,
            }
        )
    return events


def test_eviscerate_fit_recovers_the_attack_power_term():
    events = eviscerates(random.Random(7), 60.0, 120.0, 0.03, 4.0)
    events += [
        {"k": "ya", "n": 1, "t": 99, "s": "Eviscerate", "o": "crit", "a": 9999, "cp": 5, "ap": 1000}
    ]
    events += [
        {"k": "ya", "n": 1, "t": 99, "s": "Backstab", "o": "hit", "a": 9999, "cp": 0, "ap": 1000}
    ]
    fit = recorder.eviscerate_fit(recording(events))
    assert fit is not None
    base, per_point, ap_term = fit.coefficients
    assert ap_term == pytest.approx(0.03, abs=3 * fit.errors[2])
    assert per_point == pytest.approx(120.0, abs=3 * fit.errors[1])
    assert base == pytest.approx(60.0, abs=3 * fit.errors[0])
    assert fit.errors[2] < 0.01
    assert fit.samples == 40


def test_eviscerate_fit_needs_samples_and_variation():
    assert recorder.eviscerate_fit(recording([])) is None
    flat = [
        {
            "k": "ya",
            "n": 1,
            "t": i,
            "s": "Eviscerate",
            "o": "hit",
            "a": 300 + i,
            "cp": 5,
            "ap": 1000,
        }
        for i in range(8)
    ]
    with pytest.raises(SingularFitError):
        recorder.eviscerate_fit(recording(flat))
    assert "fit impossible" in recorder.report(recording(flat))


# Windfury --------------------------------------------------------------------------------


def test_windfury_counts_by_unit_and_preceding_event_with_rates():
    sessions = [
        session(1, counts={"player.swing_main": 200, "player.special": 50, "pet.swing_main": 100})
    ]
    events = (
        [
            {"k": "wf", "n": 1, "t": i, "e": "extra", "u": "player", "p": "swing_main"}
            for i in range(20)
        ]
        + [
            {"k": "wf", "n": 1, "t": i, "e": "extra", "u": "player", "p": "special"}
            for i in range(5)
        ]
        + [
            {"k": "wf", "n": 1, "t": i, "e": "extra", "u": "pet", "p": "swing_main"}
            for i in range(8)
        ]
        + [
            {
                "k": "wf",
                "n": 1,
                "t": 1,
                "e": "extra",
                "u": "player",
                "p": "swing_main",
                "after": True,
            }
        ]
        + [{"k": "wf", "n": 1, "t": i, "e": "aura", "u": "pet", "p": "none"} for i in range(3)]
    )
    rows, auras = recorder.windfury_rows(recording(events, sessions))
    by_key = {(r.unit, r.preceding, r.following): r for r in rows}
    assert by_key[("player", "swing_main", False)].rate == pytest.approx(0.10)
    assert by_key[("player", "special", False)].rate == pytest.approx(0.10)
    assert by_key[("pet", "swing_main", False)].rate == pytest.approx(0.08)
    assert by_key[("player", "swing_main", True)].rate is None
    assert auras == {"pet": 3}
    text = recorder.report(recording(events, sessions))
    assert "pet: totem aura applied 3 times" in text
    assert "proc after swing_main" in text


def test_windfury_with_nothing_recorded():
    assert "no Windfury events" in recorder.report(recording([]))


# Mana ----------------------------------------------------------------------------------


def mana_session_events(sid: int, per_tick: int, since: float | None, start: float):
    value, events = 100, []
    for index in range(8):
        value += per_tick
        event = {"k": "pw", "n": sid, "p": "MANA", "t": start + 2 * index, "v": value, "mx": 5000}
        if since is not None:
            event["since"] = since
        events.append(event)
    return events


def test_mana_fit_is_linear_in_spirit_per_rule_state():
    sessions = [session(1, spirit=100), session(2, spirit=200), session(3, spirit=300)]
    events = []
    for sid, spirit in ((1, 100), (2, 200), (3, 300)):
        events += mana_session_events(sid, 10 + spirit // 10, None, 1000 * sid)
        events += mana_session_events(sid, 3 + spirit // 50, 2.0, 1000 * sid + 100)
    outside, inside = recorder.mana_fits(recording(events, sessions))
    assert not outside.inside_rule and inside.inside_rule
    assert outside.fit is not None and inside.fit is not None
    assert outside.fit.coefficients == pytest.approx((10.0, 0.1))
    assert inside.fit.coefficients == pytest.approx((3.0, 0.02))
    assert outside.ticks[0].period == pytest.approx(2.0)
    text = recorder.report(recording(events, sessions))
    assert "inside the five-second rule" in text and "per tick = a + b*spirit" in text


def test_mana_fit_asks_for_a_second_spirit_value():
    events = mana_session_events(1, 12, None, 10)
    outside, inside = recorder.mana_fits(recording(events))
    assert outside.fit is None and inside.ticks == []
    assert "two different spirit values" in recorder.report(recording(events))


def test_mana_spending_is_not_a_tick():
    events = [
        {"k": "pw", "n": 1, "p": "MANA", "t": 1, "v": 500, "mx": 900},
        {"k": "pw", "n": 1, "p": "MANA", "t": 2, "v": 300, "mx": 900},
        {"k": "pw", "n": 1, "p": "MANA", "t": 4, "v": 312, "mx": 900},
    ]
    outside, _ = recorder.mana_fits(recording(events))
    assert [(t.amount, t.ticks) for t in outside.ticks] == [(12, 1)]


def test_mana_points_without_a_session_spirit_are_skipped():
    sessions = [{"id": 1, "level": 60}]
    outside, _ = recorder.mana_fits(recording(mana_session_events(1, 5, None, 1), sessions))
    assert outside.ticks == []


# Shadowfiend ------------------------------------------------------------------------------


def test_fiend_cadence_median_gap_and_mean_amount():
    events = [
        {"k": "en", "n": 1, "t": 10 + 1.5 * i, "s": "Shadowfiend", "a": 100 + i, "pt": 0}
        for i in range(6)
    ]
    events += [{"k": "en", "n": 1, "t": 200, "s": "Shadowfiend", "a": 100, "pt": 0}]
    events += [{"k": "en", "n": 1, "t": 5, "s": "Dark Sacrifice", "a": 50, "pt": 0}]
    rows = {r.spell: r for r in recorder.fiend_cadence(recording(events))}
    assert rows["Shadowfiend"].returns == 7
    assert rows["Shadowfiend"].interval == pytest.approx(1.5)
    assert rows["Shadowfiend"].interval_count == 5
    assert rows["Dark Sacrifice"].interval is None
    assert "Shadowfiend: 7 returns" in recorder.report(recording(events))


def test_fiend_with_nothing_recorded():
    assert "no Shadowfiend returns" in recorder.report(recording([]))


# Report and CLI ------------------------------------------------------------------------------


def test_report_has_every_section():
    text = recorder.report(recording(energy_events(2.0, 20, 6)))
    for title in ("Energy", "White swings", "Yellow attacks", "Windfury", "Mana", "Shadowfiend"):
        assert f"\n{title}\n" in text
    assert "tick amount 20" in text
    assert "tick period 2.000s" in text


def test_report_prints_the_white_table_and_fit_line():
    events = swings("m", 63, {"hit": (50, 100), "glance": (20, 70), "miss": (20, 0)})
    events += eviscerates(random.Random(1), 60.0, 120.0, 0.03, 2.0)
    text = recorder.report(recording(events))
    assert "main hand, target +3 levels: 90 swings" in text
    assert "glancing damage / plain hit damage = 0.700" in text
    assert "ap_term 0.0" in text


def test_cli_prints_the_report(tmp_path, capsys):
    path = write(tmp_path, store(energy_events(2.0, 20, 8), [session(1)]))
    assert main(["recorder", str(path)]) == 0
    assert "tick amount 20" in capsys.readouterr().out


def test_cli_reports_a_missing_or_wrong_file(tmp_path, caplog):
    assert main(["recorder", str(tmp_path / "missing.lua")]) == 1
    bad = tmp_path / "bad.lua"
    bad.write_text("Other = 1", encoding="utf-8")
    assert main(["recorder", str(bad)]) == 1
    assert "no ForeverSixtyDB.recorder" in caplog.text
