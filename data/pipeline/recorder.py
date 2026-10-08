"""`python -m pipeline recorder <SavedVariables lua>`: fit what the addon's recorder wrote.

The addon (addon/ForeverSixty/Recorder.lua, `/fs record on`) stores the raw
events of a play session under ForeverSixtyDB.recorder. This module reads
them back and prints the fits the simulator wants but cannot read from the
client tables: the energy tick, white-hit outcome rates, the Eviscerate
attack-power term, Windfury proc counts, the mana regeneration formula and
the Shadowfiend's cadence. The store's shape is documented at the top of the
Lua module; SCHEMA must match its `Recorder.SCHEMA`.
"""

from __future__ import annotations

import statistics
from collections import Counter, defaultdict
from collections.abc import Iterable, Sequence
from dataclasses import dataclass
from pathlib import Path
from typing import Any

from pipeline import savedvars
from pipeline.recorder_stats import (
    Fit,
    Interval,
    SingularFitError,
    least_squares,
    median,
    wilson,
)

SCHEMA = 1
SAVED_VARIABLE = "ForeverSixtyDB"
FIVE_SECOND_RULE = 5.0
#: A gap longer than this between two power gains is not a tick interval.
MAX_TICK_INTERVAL = 4.0
MIN_TICK_INTERVAL = 0.5
#: Shadowfiend returns further apart than this belong to different summons.
MAX_FIEND_GAP = 10.0
WHITE_OUTCOMES = ("hit", "crit", "glance", "miss", "dodge", "parry", "block")
HANDS = {"m": "main", "o": "off"}
PROC_KINDS = ("swing_main", "swing_off", "special")
EVISCERATE = "Eviscerate"
MIN_FIT_SAMPLES = 4


class RecorderError(ValueError):
    """The file holds no readable recorder record."""


@dataclass(frozen=True)
class Recording:
    """The record in time order, with the session snapshots by id."""

    sessions: dict[int, dict[str, Any]]
    events: list[dict[str, Any]]
    capacity: int

    def of_kind(self, kind: str) -> list[dict[str, Any]]:
        return [event for event in self.events if event.get("k") == kind]

    def level_of(self, event: dict[str, Any]) -> int | None:
        return self.sessions.get(event.get("n"), {}).get("level")


def _sequence(value: Any) -> list[Any]:
    """A Lua array as a list; an empty table or a keyed one as what it holds."""
    if isinstance(value, list):
        return value
    return [value[key] for key in sorted(value)] if isinstance(value, dict) else []


def load(path: Path) -> Recording:
    """Read and validate the record, oldest event first."""
    try:
        store = savedvars.read(path).get(SAVED_VARIABLE, {})
    except savedvars.SavedVariablesError as error:
        raise RecorderError(f"{path} is not a SavedVariables file: {error}") from error
    store = store.get("recorder") if isinstance(store, dict) else None
    if not isinstance(store, dict):
        raise RecorderError(f"{path} has no {SAVED_VARIABLE}.recorder; run /fs record on in game")
    if store.get("schema") != SCHEMA:
        raise RecorderError(f"record schema {store.get('schema')!r} is not the supported {SCHEMA}")
    physical = _sequence(store.get("events"))
    start = int(store.get("start", 1)) - 1
    sessions = {
        int(session["id"]): session
        for session in _sequence(store.get("sessions"))
        if isinstance(session, dict) and "id" in session
    }
    return Recording(sessions, physical[start:] + physical[:start], int(store.get("cap", 0)))


# Energy ---------------------------------------------------------------------


@dataclass(frozen=True)
class EnergyTick:
    amount: int
    amount_count: int
    gains: int
    period: float | None
    period_stdev: float | None
    period_count: int


def _gain_intervals(points: Sequence[tuple[float, int]], amount: int) -> list[float]:
    """Gaps between consecutive gains of `amount`, with no other gain between."""
    intervals: list[float] = []
    previous: float | None = None
    for time, gained in points:
        if gained != amount:
            previous = None
            continue
        if previous is not None and MIN_TICK_INTERVAL <= time - previous <= MAX_TICK_INTERVAL:
            intervals.append(time - previous)
        previous = time
    return intervals


def _power_gains(
    recording: Recording, token: str
) -> dict[int, list[tuple[float, int, dict[str, Any]]]]:
    """Per session, the (time, gain, event) of every rise in a power, in order."""
    last: dict[int, int] = {}
    gains: dict[int, list[tuple[float, int, dict[str, Any]]]] = defaultdict(list)
    for event in recording.of_kind("pw"):
        if event.get("p") != token:
            continue
        session, value = event["n"], event["v"]
        if session in last and value > last[session]:
            gains[session].append((event["t"], value - last[session], event))
        last[session] = value
    return gains


def energy_tick(recording: Recording) -> EnergyTick | None:
    """The energy tick amount (the commonest gain below the cap) and its period."""
    gains = _power_gains(recording, "ENERGY")
    uncapped = [
        gain
        for per in gains.values()
        for _, gain, event in per
        if event["v"] < event.get("mx", 1e9)
    ]
    if not uncapped:
        return None
    amount, amount_count = Counter(uncapped).most_common(1)[0]
    intervals = [
        interval
        for per in gains.values()
        for interval in _gain_intervals([(t, g) for t, g, _ in per], amount)
    ]
    return EnergyTick(
        amount,
        amount_count,
        len(uncapped),
        median(intervals) if intervals else None,
        statistics.pstdev(intervals) if len(intervals) > 1 else None,
        len(intervals),
    )


# White swings -----------------------------------------------------------------


@dataclass(frozen=True)
class OutcomeRate:
    count: int
    rate: float
    interval: Interval


@dataclass(frozen=True)
class WhiteGroup:
    hand: str
    level_gap: int
    swings: int
    rates: dict[str, OutcomeRate]
    glancing_ratio: float | None


def _mean(values: Iterable[float]) -> float | None:
    values = list(values)
    return statistics.fmean(values) if values else None


def _glancing_ratio(swings: list[dict[str, Any]]) -> float | None:
    glance = _mean(s["a"] for s in swings if s["o"] == "glance" and s.get("a"))
    hit = _mean(s["a"] for s in swings if s["o"] == "hit" and s.get("a"))
    return glance / hit if glance is not None and hit else None


def white_groups(recording: Recording) -> list[WhiteGroup]:
    """Outcome rates (Wilson 95%) by hand and target level minus player level."""
    grouped: dict[tuple[str, int], list[dict[str, Any]]] = defaultdict(list)
    for swing in recording.of_kind("sw"):
        level, target = recording.level_of(swing), swing.get("tl")
        if level is not None and target is not None and target > 0:
            grouped[(swing["h"], target - level)].append(swing)
    groups = []
    for (hand, gap), swings in sorted(grouped.items()):
        tally = Counter(s["o"] for s in swings)
        total = len(swings)
        rates = {
            outcome: OutcomeRate(
                tally[outcome], tally[outcome] / total, wilson(tally[outcome], total)
            )
            for outcome in WHITE_OUTCOMES
            if tally[outcome]
        }
        groups.append(WhiteGroup(HANDS.get(hand, hand), gap, total, rates, _glancing_ratio(swings)))
    return groups


# Yellow attacks ----------------------------------------------------------------


def eviscerate_fit(recording: Recording) -> Fit | None:
    """Plain-hit Eviscerate damage = a + b*points + c*(attack power * points).

    `c` is the per-point attack-power coefficient. Returns None with too few
    hits, and raises SingularFitError when points or attack power never varied.
    """
    hits = [
        e
        for e in recording.of_kind("ya")
        if e.get("s") == EVISCERATE
        and e.get("o") == "hit"
        and e.get("a")
        and e.get("cp")
        and e.get("ap")
    ]
    if len(hits) < MIN_FIT_SAMPLES:
        return None
    rows = [(1.0, float(e["cp"]), float(e["ap"] * e["cp"])) for e in hits]
    return least_squares(rows, [float(e["a"]) for e in hits])


# Windfury -----------------------------------------------------------------------


@dataclass(frozen=True)
class WindfuryRow:
    unit: str
    preceding: str
    following: bool
    procs: int
    eligible: int | None

    @property
    def rate(self) -> float | None:
        return self.procs / self.eligible if self.eligible else None


def windfury_rows(recording: Recording) -> tuple[list[WindfuryRow], Counter[str]]:
    """Weapon procs by unit and preceding event kind, and totem auras by unit."""
    eligible: Counter[str] = Counter()
    for session in recording.sessions.values():
        counts = session.get("counts")
        for key, value in counts.items() if isinstance(counts, dict) else []:
            eligible[key] += value
    procs: Counter[tuple[str, str, bool]] = Counter()
    auras: Counter[str] = Counter()
    for event in recording.of_kind("wf"):
        if event.get("e") == "aura":
            auras[event["u"]] += 1
        else:
            procs[(event["u"], event["p"], bool(event.get("after")))] += 1
    rows = [
        WindfuryRow(unit, kind, after, count, None if after else eligible.get(f"{unit}.{kind}"))
        for (unit, kind, after), count in sorted(procs.items())
    ]
    return rows, auras


# Mana -----------------------------------------------------------------------------


@dataclass(frozen=True)
class ManaTicks:
    inside_rule: bool
    spirit: float
    ticks: int
    amount: float
    period: float | None


@dataclass(frozen=True)
class ManaFit:
    inside_rule: bool
    ticks: list[ManaTicks]
    fit: Fit | None


def _mana_ticks(recording: Recording) -> list[ManaTicks]:
    buckets: dict[tuple[int, bool], list[tuple[float, int]]] = defaultdict(list)
    for session, per in _power_gains(recording, "MANA").items():
        for time, gain, event in per:
            since = event.get("since")
            buckets[(session, since is not None and since < FIVE_SECOND_RULE)].append((time, gain))
    ticks = []
    for (session, inside), points in sorted(buckets.items()):
        spirit = recording.sessions.get(session, {}).get("spirit")
        if spirit is None:
            continue
        amount = Counter(gain for _, gain in points).most_common(1)[0][0]
        intervals = _gain_intervals(points, amount)
        count = sum(1 for _, gain in points if gain == amount)
        ticks.append(
            ManaTicks(inside, spirit, count, amount, median(intervals) if intervals else None)
        )
    return ticks


def mana_fits(recording: Recording) -> list[ManaFit]:
    """Per tick: mana = a + b*spirit, outside then inside the five-second rule."""
    ticks = _mana_ticks(recording)
    fits = []
    for inside in (False, True):
        chosen = [t for t in ticks if t.inside_rule == inside]
        fit = None
        if len({t.spirit for t in chosen}) >= 2:
            try:
                fit = least_squares([(1.0, t.spirit) for t in chosen], [t.amount for t in chosen])
            except SingularFitError:
                fit = None
        fits.append(ManaFit(inside, chosen, fit))
    return fits


# Shadowfiend -------------------------------------------------------------------------


@dataclass(frozen=True)
class FiendCadence:
    spell: str
    returns: int
    amount: float
    interval: float | None
    interval_count: int


def fiend_cadence(recording: Recording) -> list[FiendCadence]:
    by_spell: dict[str, list[dict[str, Any]]] = defaultdict(list)
    for event in recording.of_kind("en"):
        by_spell[event["s"]].append(event)
    rows = []
    for spell, events in sorted(by_spell.items()):
        times = [e["t"] for e in events]
        gaps = [b - a for a, b in zip(times, times[1:], strict=False) if 0 < b - a <= MAX_FIEND_GAP]
        rows.append(
            FiendCadence(
                spell,
                len(events),
                statistics.fmean(e["a"] for e in events),
                median(gaps) if gaps else None,
                len(gaps),
            )
        )
    return rows


# Report ------------------------------------------------------------------------------------


def _percent(rate: OutcomeRate) -> str:
    low, high = rate.interval.low, rate.interval.high
    return f"{rate.rate:6.1%} [{low:.1%}-{high:.1%}] n={rate.count}"


def _fit_line(name: str, fit: Fit, labels: Sequence[str]) -> str:
    terms = ", ".join(
        f"{label} {value:.4g} +/- {error:.2g}"
        for label, value, error in zip(labels, fit.coefficients, fit.errors, strict=True)
    )
    return f"  {name}: {terms} (n={fit.samples})"


def _energy_section(recording: Recording) -> list[str]:
    tick = energy_tick(recording)
    if tick is None:
        return ["  no energy gains recorded"]
    lines = [f"  tick amount {tick.amount} ({tick.amount_count} of {tick.gains} gains)"]
    if tick.period is None:
        return [*lines, "  tick period: no two consecutive ticks recorded"]
    spread = f", s.d. {tick.period_stdev:.3f}" if tick.period_stdev is not None else ""
    return [
        *lines,
        f"  tick period {tick.period:.3f}s median{spread} over {tick.period_count} intervals",
    ]


def _white_section(recording: Recording) -> list[str]:
    groups = white_groups(recording)
    if not groups:
        return ["  no white swings with a target level recorded"]
    lines = []
    for group in groups:
        lines.append(
            f"  {group.hand} hand, target {group.level_gap:+d} levels: {group.swings} swings"
        )
        lines.extend(f"    {name:<7}{_percent(rate)}" for name, rate in group.rates.items())
        if group.glancing_ratio is not None:
            lines.append(f"    glancing damage / plain hit damage = {group.glancing_ratio:.3f}")
    return lines


def _yellow_section(recording: Recording) -> list[str]:
    try:
        fit = eviscerate_fit(recording)
    except SingularFitError as error:
        return [f"  Eviscerate fit impossible: {error}; vary points spent and attack power"]
    if fit is None:
        return [
            f"  fewer than {MIN_FIT_SAMPLES} plain Eviscerate hits with points and attack power"
        ]
    return [
        "  Eviscerate hit = base + points * per_point + attack_power * points * ap_term",
        _fit_line("fit", fit, ("base", "per_point", "ap_term")),
    ]


def _windfury_section(recording: Recording) -> list[str]:
    rows, auras = windfury_rows(recording)
    if not rows and not auras:
        return ["  no Windfury events recorded"]
    lines = []
    for row in rows:
        when = "after" if row.following else "before"
        rate = f"rate {row.rate:.1%} of {row.eligible}" if row.rate is not None else "rate n/a"
        lines.append(f"  {row.unit}: proc {when} {row.preceding}: {row.procs} ({rate})")
    lines.extend(
        f"  {unit}: totem aura applied {count} times" for unit, count in sorted(auras.items())
    )
    return lines


def _mana_section(recording: Recording) -> list[str]:
    lines = []
    for result in mana_fits(recording):
        state = "inside" if result.inside_rule else "outside"
        lines.append(
            f"  {state} the five-second rule: {len(result.ticks)} (session, spirit) points"
        )
        lines.extend(
            f"    spirit {t.spirit:g}: {t.amount} mana per tick x{t.ticks}"
            + (f", every {t.period:.2f}s" if t.period else "")
            for t in result.ticks
        )
        if result.fit is not None:
            lines.append(_fit_line("per tick = a + b*spirit", result.fit, ("a", "b")))
        elif result.ticks:
            lines.append("    need ticks at two different spirit values to fit a + b*spirit")
    return lines


def _fiend_section(recording: Recording) -> list[str]:
    rows = fiend_cadence(recording)
    if not rows:
        return ["  no Shadowfiend returns recorded"]
    return [
        f"  {row.spell}: {row.returns} returns, mean {row.amount:.1f} mana"
        + (
            f", every {row.interval:.2f}s median over {row.interval_count} gaps"
            if row.interval
            else ""
        )
        for row in rows
    ]


def report(recording: Recording) -> str:
    """The full text report."""
    sections = (
        ("Energy", _energy_section),
        ("White swings", _white_section),
        ("Yellow attacks", _yellow_section),
        ("Windfury", _windfury_section),
        ("Mana", _mana_section),
        ("Shadowfiend", _fiend_section),
    )
    lines = [
        f"recorder: {len(recording.events)} events (capacity {recording.capacity}), "
        f"{len(recording.sessions)} sessions"
    ]
    for title, build in sections:
        lines.extend(["", title, *build(recording)])
    return "\n".join(lines)


def run(path: Path) -> str:
    return report(load(path))
