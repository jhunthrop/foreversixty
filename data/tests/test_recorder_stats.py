import math

import pytest

from pipeline.recorder_stats import SingularFitError, least_squares, median, wilson


def test_wilson_matches_known_values():
    interval = wilson(50, 100)
    assert interval.low == pytest.approx(0.4038, abs=1e-3)
    assert interval.high == pytest.approx(0.5962, abs=1e-3)


def test_wilson_stays_inside_zero_and_one_at_the_extremes():
    assert wilson(0, 20).low == 0.0
    assert wilson(20, 20).high == 1.0
    assert 0 < wilson(0, 20).high < 0.2


def test_wilson_needs_trials():
    with pytest.raises(ValueError):
        wilson(0, 0)


def test_least_squares_recovers_a_line_and_its_errors():
    xs = [1, 2, 3, 4, 5, 6]
    ys = [
        3 + 2 * x + noise for x, noise in zip(xs, [0.1, -0.1, 0.05, -0.05, 0.1, -0.1], strict=True)
    ]
    fit = least_squares([(1.0, float(x)) for x in xs], ys)
    assert fit.coefficients[0] == pytest.approx(3, abs=0.2)
    assert fit.coefficients[1] == pytest.approx(2, abs=0.05)
    assert all(error > 0 for error in fit.errors)
    assert fit.samples == 6


def test_least_squares_has_no_error_without_spare_degrees_of_freedom():
    fit = least_squares([(1.0, 1.0), (1.0, 2.0)], [1.0, 2.0])
    assert fit.coefficients == pytest.approx((0.0, 1.0))
    assert all(math.isnan(error) for error in fit.errors)


def test_least_squares_refuses_a_regressor_that_never_varies():
    with pytest.raises(SingularFitError):
        least_squares([(1.0, 5.0), (1.0, 5.0), (1.0, 5.0)], [1.0, 2.0, 3.0])


def test_least_squares_refuses_too_few_samples():
    with pytest.raises(SingularFitError):
        least_squares([(1.0, 2.0, 3.0)], [1.0])


def test_median():
    assert median([3, 1, 2]) == 2
