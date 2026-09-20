import json
from pathlib import Path

import pytest

from pipeline.specs import load_specs
from pipeline.weights import WeightsError, load_weights, write_weights

CURATED = Path("curated")


def test_every_spec_has_weights():
    weights = load_weights(CURATED)
    assert sorted(weights) == sorted(spec.spec for spec in load_specs(CURATED))


def test_the_reference_stat_is_the_unit():
    """Weights are relative, and the sim's own reference stat is the 1.0.

    It is the sanity check that catches a pasted table from another spec:
    a caster weight list handed to a rogue has no attack_power at all.
    """
    weights = load_weights(CURATED)
    for spec in load_specs(CURATED):
        entry = weights[spec.spec]
        assert entry.weights[spec.reference_stat] == 1.0, spec.spec


def test_a_weight_key_outside_the_contract_vocabulary_is_refused(tmp_path):
    (tmp_path / "specs.json").write_text(
        (CURATED / "specs.json").read_text(encoding="utf-8"), encoding="utf-8"
    )
    (tmp_path / "stat-weights.json").write_text(
        json.dumps(
            [
                {
                    "spec": "warrior-fury",
                    "weights": {"attack_power": 1.0, "melee_crit": 2.0},
                    "sources": [{"label": "x", "url": "https://x", "kind": "community"}],
                }
            ]
        ),
        encoding="utf-8",
    )
    with pytest.raises(WeightsError, match="melee_crit"):
        load_weights(tmp_path)


def test_an_unsourced_entry_is_refused(tmp_path):
    (tmp_path / "specs.json").write_text(
        (CURATED / "specs.json").read_text(encoding="utf-8"), encoding="utf-8"
    )
    (tmp_path / "stat-weights.json").write_text(
        json.dumps([{"spec": "warrior-fury", "weights": {"attack_power": 1.0}, "sources": []}]),
        encoding="utf-8",
    )
    with pytest.raises(SystemExit, match="at least one source"):
        load_weights(tmp_path)


def test_write_weights_emits_the_build_copy(tmp_path):
    path = write_weights("1.60.1.69893", root=tmp_path, curated_dir=CURATED)
    rows = json.loads(path.read_text(encoding="utf-8"))
    assert [row["spec"] for row in rows] == sorted(row["spec"] for row in rows)
    assert rows[0]["sources"][0]["url"].startswith("http")
