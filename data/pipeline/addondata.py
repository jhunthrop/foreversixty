"""What the in-game addon needs, as JSON and then as Lua.

The addon cannot reach the network, so everything it knows about talent
trees and stat weights is generated here and committed as
`addon/ForeverSixty/Data.lua` (pipeline/addonlua.py renders it).

Two decisions are load-bearing:

* Talents travel in the **site's array order** inside each tab. The
  export string encodes "the ranks of that tree's talents in tab order",
  and the client's own `GetTalentInfo` index order is not guaranteed to
  match this file's; the addon therefore keys the client's talents by
  tier and column and emits ranks in the order here. If this order ever
  changed, every export string in flight would decode onto the wrong
  talents, which is why `test_addondata.py` pins it against the source.
* tier and column are **1-based**, because `GetTalentInfo` returns them
  that way. The talent files are 0-based; the conversion happens once,
  here.
* Every talent carries its **TraitNode id** (`node`). The 1.60 client has
  no `GetTalentInfo`; its trees live in the modern trait system and the
  addon reads each rank with `C_Traits.GetNodeInfo(configID, node)`. The
  tier/column cell stays the key between the addon's own modules.
"""

from __future__ import annotations

import json
from pathlib import Path

from pipeline.models import AddonClass, AddonData, AddonTab, AddonTalent
from pipeline.normalize import write_model
from pipeline.weights import load_weights

#: What is added to a tree's `position` to get the client's tab index.
#: `GetTalentInfo(tabIndex, …)` is 1-based; spike check 7 confirms the
#: tab order is the same order.
TAB_BASE = 1


class AddonDataError(SystemExit):
    """A build directory says something the addon generator will not publish."""


def _class_slugs(root: Path, build: str) -> list[str]:
    path = root / build / "classes.json"
    if not path.exists():
        raise AddonDataError(f"missing {path}")
    return sorted(entry["slug"] for entry in json.loads(path.read_text(encoding="utf-8")))


def _tabs(root: Path, build: str, class_slug: str) -> list[AddonTab]:
    path = root / build / "talents" / f"{class_slug}.json"
    if not path.exists():
        raise AddonDataError(f"missing {path}")
    file = json.loads(path.read_text(encoding="utf-8"))
    if file["build"] != build:
        raise AddonDataError(
            f"{path} says build {file['build']}, not {build}; re-run normalize for this build"
        )
    tabs: list[AddonTab] = []
    for tree in sorted(file["trees"], key=lambda tree: tree["position"]):
        talents = [
            AddonTalent(
                name=talent["name"],
                tier=talent["tier"] + TAB_BASE,
                column=talent["column"] + TAB_BASE,
                max_rank=talent["max_rank"],
                node=talent["id"],
            )
            for talent in tree["talents"]
        ]
        cells = {(talent.tier, talent.column) for talent in talents}
        if len(cells) != len(talents):
            raise AddonDataError(
                f"{path} tree {tree['name']!r} has two talents in one tier/column cell; "
                f"the addon keys the client's talents by that pair and cannot tell them apart"
            )
        tabs.append(AddonTab(name=tree["name"], talents=talents))
    return tabs


def build_addon_data(
    build: str,
    root: Path = Path("builds"),
    curated_dir: Path = Path("curated"),
) -> AddonData:
    classes = {
        slug: AddonClass(tabs=_tabs(root, build, slug)) for slug in _class_slugs(root, build)
    }
    weights = {key: entry.weights for key, entry in load_weights(curated_dir).items()}
    return AddonData(build=build, classes=classes, weights=weights)


def write_addon_data(
    build: str,
    root: Path = Path("builds"),
    out_root: Path | None = None,
    curated_dir: Path = Path("curated"),
) -> Path:
    path = (out_root or root) / build / "addon-data.json"
    path.parent.mkdir(parents=True, exist_ok=True)
    write_model(build_addon_data(build, root=root, curated_dir=curated_dir), path)
    return path


def check_addon_data(
    build: str,
    root: Path = Path("builds"),
    curated_dir: Path = Path("curated"),
) -> bool:
    """True when the committed copy has drifted from build_addon_data's output.

    Compares parsed JSON rather than raw text so formatting-only drift (key
    order, whitespace) is not mistaken for content drift.
    """
    import tempfile

    with tempfile.TemporaryDirectory() as directory:
        expected = write_addon_data(
            build, root=root, out_root=Path(directory), curated_dir=curated_dir
        )
        wanted = json.loads(expected.read_text(encoding="utf-8"))
    path = root / build / "addon-data.json"
    return not path.exists() or json.loads(path.read_text(encoding="utf-8")) != wanted
