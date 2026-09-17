"""Loading the synthetic trait build. A sibling module, not a test module, so
test_traits.py and test_normalize_trait_trees.py read the same fixture the same
way. pytest puts data/tests/ on sys.path for test files with no __init__.py, so
both import it as a top-level module."""

from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.normalize.traits import TraitRows

TRAITS = Path(__file__).parent / "fixtures/traits"

#: TraitRows field -> the CSV that fills it.
TABLES = {
    "skill_line": "SkillLine",
    "skill_line_x_trait_tree": "SkillLineXTraitTree",
    "talent_tab": "TalentTab",
    "chr_classes": "ChrClasses",
    "node": "TraitNode",
    "node_entry": "TraitNodeEntry",
    "node_x_entry": "TraitNodeXTraitNodeEntry",
    "definition": "TraitDefinition",
    "edge": "TraitEdge",
    "cond": "TraitCond",
    "currency": "TraitCurrency",
    "node_group": "TraitNodeGroup",
    "node_group_x_node": "TraitNodeGroupXTraitNode",
}


def trait_rows(**replace) -> TraitRows:
    """The fixture build, with any one table swapped out by keyword."""
    loaded = {field: read_csv(TRAITS / f"{name}.csv") for field, name in TABLES.items()}
    return TraitRows(**{**loaded, **replace})
