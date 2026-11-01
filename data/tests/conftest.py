"""Shared fixtures for the data suite."""

import json
import shutil
from dataclasses import dataclass
from pathlib import Path

import pytest

from pipeline.addondata import write_addon_data
from pipeline.addonlua import write_lua

#: The active build, read the way the Makefile and data.yml read it.
ACTIVE_BUILD = json.loads(
    (Path(__file__).resolve().parents[2] / "web" / "src" / "data" / "active-build.json").read_text()
)["build"]

#: Build-directory entries the addon generator never reads; left out of the copy to keep it small.
_NOT_ADDON_INPUTS = shutil.ignore_patterns("raw", "icons", "trees", "simdb.bin")


@dataclass(frozen=True)
class AddonBuild:
    """A scratch copy of the active build whose addon-data.json and Data.lua were
    regenerated from the committed inputs."""

    build: str
    root: Path
    lua_path: Path


@pytest.fixture(scope="session")
def regenerated_addon_build(tmp_path_factory: pytest.TempPathFactory) -> AddonBuild:
    """The committed generated files trail their inputs between a curated edit and
    the push-time regeneration (data.yml's `addon-data` job) or the nightly, so a
    test of the generator round trip reads a fresh regeneration, never the
    committed copies."""
    base = tmp_path_factory.mktemp("addon-build")
    root = base / "builds"
    shutil.copytree(Path("builds") / ACTIVE_BUILD, root / ACTIVE_BUILD, ignore=_NOT_ADDON_INPUTS)
    lua_path = base / "addon" / "ForeverSixty" / "Data.lua"
    write_addon_data(ACTIVE_BUILD, root=root)
    write_lua(ACTIVE_BUILD, root=root, lua_path=lua_path)
    return AddonBuild(build=ACTIVE_BUILD, root=root, lua_path=lua_path)
