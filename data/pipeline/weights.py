"""Curated stat weights per spec, and the per-build copy the site reads.

Design section "Stat weights": every entry has at least one source, the
weights are opinions labelled as such, and a sim-derived weight list
replaces a curated one spec at a time once the engine has validated it.

The vocabulary is parity contract 10.8's -- the engine's `Stat` enum in
snake case, one `hit` and one `crit`, `healing_power` rather than the
planner's own `healing` -- because these numbers are multiplied against
engine stats by `/sim/weights` and against `GetItemStats` output by the
addon, and a key neither side knows scores as zero without saying so.
"""

from __future__ import annotations

import json
from pathlib import Path

from pipeline.curated import parse_sources
from pipeline.models import StatWeights
from pipeline.normalize import write_records
from pipeline.simdb.statmap import STAT_IDS
from pipeline.specs import load_specs


class WeightsError(SystemExit):
    """curated/stat-weights.json says something the pipeline will not publish."""


def load_weights(curated_dir: Path = Path("curated")) -> dict[str, StatWeights]:
    path = curated_dir / "stat-weights.json"
    if not path.exists():
        raise WeightsError(f"missing {path}")
    by_spec: dict[str, StatWeights] = {}
    known = {spec.spec: spec for spec in load_specs(curated_dir)}
    for entry in json.loads(path.read_text(encoding="utf-8")):
        record = StatWeights(**entry)
        where = f"stat-weights.json entry {record.spec}"
        if record.spec not in known:
            raise WeightsError(f"{where} names a spec curated/specs.json does not have")
        if record.spec in by_spec:
            raise WeightsError(f"stat-weights.json names {record.spec} twice")
        for key, value in record.weights.items():
            if key not in STAT_IDS:
                raise WeightsError(
                    f"{where} weights {key!r}; contract 10.8's vocabulary is the engine's "
                    f"Stat enum in snake case -- use one of {sorted(STAT_IDS)}"
                )
            if value < 0:
                raise WeightsError(f"{where} weights {key!r} at {value}; a weight is not negative")
        reference = known[record.spec].reference_stat
        if record.weights.get(reference) != 1.0:
            raise WeightsError(
                f"{where} must weight its reference stat {reference!r} at exactly 1.0 -- "
                f"weights are relative to it and /sim/weights normalises the same way"
            )
        parse_sources([source.model_dump() for source in record.sources], where)
        by_spec[record.spec] = record
    missing = sorted(set(known) - set(by_spec))
    if missing:
        raise WeightsError(f"stat-weights.json has no entry for {', '.join(missing)}")
    return by_spec


def write_weights(
    build: str,
    root: Path = Path("builds"),
    curated_dir: Path = Path("curated"),
) -> Path:
    """Emit the build's copy: one array, ordered by spec key."""
    by_spec = load_weights(curated_dir)
    path = root / build / "stat-weights.json"
    path.parent.mkdir(parents=True, exist_ok=True)
    write_records([by_spec[key] for key in sorted(by_spec)], path)
    return path


def check_weights(
    build: str,
    root: Path = Path("builds"),
    curated_dir: Path = Path("curated"),
) -> bool:
    """True when the emitted copy has drifted from the curated list."""
    import tempfile

    with tempfile.TemporaryDirectory() as directory:
        expected = write_weights(build, root=Path(directory), curated_dir=curated_dir)
        wanted = expected.read_text(encoding="utf-8")
    path = root / build / "stat-weights.json"
    return not path.exists() or path.read_text(encoding="utf-8") != wanted
