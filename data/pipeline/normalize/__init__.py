import json
from collections.abc import Sequence
from pathlib import Path

from pydantic import BaseModel


def write_json(records: Sequence[BaseModel], path: Path) -> None:
    ordered = sorted(records, key=lambda r: r.id)  # type: ignore[attr-defined]
    payload = [r.model_dump() for r in ordered]
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(payload, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")


def normalize_build(build: str) -> None:  # completed in Task 6
    raise NotImplementedError(build)
