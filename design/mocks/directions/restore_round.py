"""Copies one earlier round's page renders out of git for the gallery's side-by-side:
    python3 design/mocks/directions/restore_round.py <commit> <label>   -> png/<label>/g-*.png
"""
import subprocess
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
RENDERS = ['g-home-1440', 'g-home-390', 'g-home-2000', 'g-bis-1440', 'g-bis-390', 'g-tiers-1440', 'g-tiers-390']


def restore(commit: str, label: str) -> None:
    out = HERE / 'png' / label
    out.mkdir(parents=True, exist_ok=True)
    for name in RENDERS:
        blob = subprocess.run(['git', 'show', f'{commit}:design/mocks/directions/png/{name}.png'], cwd=HERE,
                              capture_output=True, check=True).stdout
        (out / f'{name}.png').write_bytes(blob)
        print(out / f'{name}.png')


if __name__ == '__main__':
    restore(sys.argv[1], sys.argv[2])
