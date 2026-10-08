"""The small statistics the recorder report needs: no numpy, no scipy."""

from __future__ import annotations

import math
import statistics
from collections.abc import Sequence
from dataclasses import dataclass

WILSON_Z = 1.96


class SingularFitError(ValueError):
    """The columns of a regression are not independent (a regressor never varied)."""


@dataclass(frozen=True)
class Interval:
    low: float
    high: float


@dataclass(frozen=True)
class Fit:
    coefficients: tuple[float, ...]
    errors: tuple[float, ...]
    samples: int


def wilson(successes: int, total: int, z: float = WILSON_Z) -> Interval:
    """The Wilson score interval for a proportion (95% by default)."""
    if total <= 0:
        raise ValueError("wilson needs at least one trial")
    p = successes / total
    denominator = 1 + z * z / total
    centre = (p + z * z / (2 * total)) / denominator
    spread = z * math.sqrt(p * (1 - p) / total + z * z / (4 * total * total)) / denominator
    return Interval(max(0.0, centre - spread), min(1.0, centre + spread))


def median(values: Sequence[float]) -> float:
    return statistics.median(values)


def _invert(matrix: list[list[float]]) -> list[list[float]]:
    """Gauss-Jordan inverse with partial pivoting."""
    size = len(matrix)
    work = [row[:] + [1.0 if i == j else 0.0 for j in range(size)] for i, row in enumerate(matrix)]
    scale = max(abs(value) for row in matrix for value in row) or 1.0
    for column in range(size):
        pivot = max(range(column, size), key=lambda r: abs(work[r][column]))
        if abs(work[pivot][column]) < 1e-9 * scale:
            raise SingularFitError("the regressors are linearly dependent")
        work[column], work[pivot] = work[pivot], work[column]
        divisor = work[column][column]
        work[column] = [value / divisor for value in work[column]]
        for row in range(size):
            if row != column:
                factor = work[row][column]
                work[row] = [a - factor * b for a, b in zip(work[row], work[column], strict=True)]
    return [row[size:] for row in work]


def least_squares(rows: Sequence[Sequence[float]], y: Sequence[float]) -> Fit:
    """Ordinary least squares of y on the columns of `rows`, with standard errors.

    Include a constant column yourself for an intercept. Errors are NaN when
    there are no spare degrees of freedom.
    """
    width = len(rows[0])
    if len(rows) < width:
        raise SingularFitError(f"{len(rows)} samples cannot fit {width} coefficients")
    normal = [[sum(r[i] * r[j] for r in rows) for j in range(width)] for i in range(width)]
    inverse = _invert(normal)
    moment = [sum(r[i] * value for r, value in zip(rows, y, strict=True)) for i in range(width)]
    beta = [sum(inverse[i][j] * moment[j] for j in range(width)) for i in range(width)]
    residual = sum(
        (value - sum(b * x for b, x in zip(beta, r, strict=True))) ** 2
        for r, value in zip(rows, y, strict=True)
    )
    spare = len(rows) - width
    sigma2 = residual / spare if spare > 0 else math.nan
    errors = tuple(
        math.sqrt(sigma2 * inverse[i][i]) if spare > 0 else math.nan for i in range(width)
    )
    return Fit(tuple(beta), errors, len(rows))
