import json
from pathlib import Path

import pytest

from pipeline.clientwatch import (
    Build,
    ClientConfig,
    ClientWatchError,
    check,
    interface_number,
    is_forever,
    load_config,
    newer_than,
    newest_blizzard,
    newest_forever,
    parse_blizzard_versions,
    read_active_build,
    report,
)

CONFIG = ClientConfig(
    products=frozenset({"wowf", "wow_cn_beta"}),
    ignore_products=frozenset({"wow", "wowt", "wow_classic_era", "wow_classic_titan"}),
    min_version=(1, 60, 0, 0),
)

#: The shape wago.tools' /api/builds answered on 2026-10-16, trimmed to what matters.
BUILDS = {
    "wow": [{"version": "12.1.5.70391", "created_at": "2026-10-13 21:16:02"}],
    "wowt": [{"version": "12.1.7.70424", "created_at": "2026-10-15 19:01:02"}],
    "wow_classic_era": [{"version": "1.15.9.70003", "created_at": "2026-09-29 14:06:05"}],
    "wow_classic_titan": [{"version": "3.80.2.70315", "created_at": "2026-10-12 19:35:08"}],
    "wow_cn_beta": [
        {"version": "1.60.1.70374", "created_at": "2026-10-13 23:59:08"},
        {"version": "1.60.1.70338", "created_at": "2026-10-10 04:54:10"},
    ],
    "wowf": [{"version": "1.60.1.70245", "created_at": "2026-10-13 21:18:10"}],
}


def test_the_newest_forever_build_is_the_highest_version_across_products():
    newest = newest_forever(
        [Build(p, r["version"], r["created_at"]) for p, rows in BUILDS.items() for r in rows],
        CONFIG,
    )
    assert newest == Build("wow_cn_beta", "1.60.1.70374", "2026-10-13 23:59:08")


def test_a_listed_product_is_forever_whatever_its_version():
    assert is_forever(Build("wowf", "2.0.0.80000", ""), CONFIG)


def test_an_unlisted_product_counts_by_version_family_alone():
    """The key moved once already; a renamed product with a 1.6x build is still found."""
    assert is_forever(Build("wow_forever_live", "1.60.2.71000", ""), CONFIG)
    assert is_forever(Build("wow_forever_live", "1.61.0.71000", ""), CONFIG)
    assert not is_forever(Build("wow_forever_live", "1.15.9.70003", ""), CONFIG)
    assert not is_forever(Build("wow_forever_live", "2.5.6.70006", ""), CONFIG)


def test_an_ignored_product_is_never_forever():
    assert not is_forever(Build("wow_classic_titan", "1.99.0.1", ""), CONFIG)


def test_check_says_newer_only_when_the_candidate_beats_the_active_build():
    assert not check(BUILDS, "1.60.1.70374", CONFIG).newer
    assert check(BUILDS, "1.60.1.70338", CONFIG).newer
    newer_elsewhere = {
        **BUILDS,
        "wow_live_forever": [{"version": "1.60.2.70500", "created_at": "2026-11-04 00:00:00"}],
    }
    result = check(newer_elsewhere, "1.60.1.70374", CONFIG)
    assert result.newer and result.newest == Build(
        "wow_live_forever", "1.60.2.70500", "2026-11-04 00:00:00"
    )


def test_check_lists_newer_builds_no_rule_claimed_so_a_renumbering_is_warned_about():
    odd = {
        **BUILDS,
        "wow_mystery": [{"version": "3.0.0.70999", "created_at": "2026-11-04 00:00:00"}],
    }
    result = check(odd, "1.60.1.70374", CONFIG)
    assert [b.product for b in result.unmatched_newer] == ["wow_mystery"]
    # The ignored retail products are newer by version too, but they are known and never listed.
    assert all(b.product not in CONFIG.ignore_products for b in result.unmatched_newer)


def test_the_interface_number_follows_blizzards_toc_convention():
    assert interface_number("1.60.1.70374") == 16001
    assert interface_number("1.15.9.70003") == 11509
    assert interface_number("1.60.2.71000") == 16002


def test_version_comparison_is_numeric_not_lexical():
    assert newer_than("1.60.1.70374", "1.60.1.9999")
    assert not newer_than("1.60.1.70374", "1.60.1.70374")
    with pytest.raises(ClientWatchError):
        newer_than("1.60", "1.60.1.70374")


def test_report_is_flat_enough_for_step_outputs():
    out = report(check(BUILDS, "1.60.1.70338", CONFIG))
    assert out["newer"] is True
    assert out["newest_build"] == "1.60.1.70374"
    assert out["newest_product"] == "wow_cn_beta"
    assert out["newest_interface"] == 16001
    assert out["active_interface"] == 16001
    assert out["unmatched_newer"] == []


def test_the_committed_config_and_active_build_load():
    config = load_config()
    assert "wow_cn_beta" in config.products
    assert config.min_version == (1, 60, 0, 0)
    assert read_active_build().startswith("1.")


def test_read_active_build_fails_fast_on_an_empty_file(tmp_path: Path):
    path = tmp_path / "active-build.json"
    path.write_text(json.dumps({"build": ""}))
    with pytest.raises(ClientWatchError):
        read_active_build(path)


#: us.version.battle.net/wow_classic_beta/versions on 2026-10-16, two regions kept.
BLIZZARD_VERSIONS = (
    "Region!STRING:0|BuildConfig!HEX:16|CDNConfig!HEX:16|KeyRing!HEX:16|BuildId!DEC:4"
    "|VersionsName!String:0|ProductConfig!HEX:16\n"
    "## seqn = 4081170\n"
    "us|ea9c525eef74f1d682c00a8ef39a5d2b|fcab7fec2ccc6b8158b56b6938ec4551||70425|1.60.1.70425|fc\n"
    "eu|ea9c525eef74f1d682c00a8ef39a5d2b|fcab7fec2ccc6b8158b56b6938ec4551||70425|1.60.1.70425|fc\n"
)


def test_blizzards_versions_file_parses_to_the_us_rows_build():
    assert parse_blizzard_versions("wow_classic_beta", BLIZZARD_VERSIONS) == Build(
        "wow_classic_beta", "1.60.1.70425", ""
    )
    assert parse_blizzard_versions("wowf", "Region!STRING:0\n## seqn = 1\n") is None
    assert parse_blizzard_versions("wowf", "") is None


def test_blizzard_ahead_of_wago_is_reported_as_wago_behind_and_never_activated():
    """2026-10-16: Blizzard served 70425 while wago.tools' newest export was 70374."""
    result = check(
        BUILDS, "1.60.1.70374", CONFIG, blizzard_versions={"wow_classic_beta": BLIZZARD_VERSIONS}
    )
    assert result.blizzard == Build("wow_classic_beta", "1.60.1.70425", "")
    assert result.wago_behind
    assert not result.newer
    out = report(result)
    assert out["blizzard_build"] == "1.60.1.70425" and out["wago_behind"] is True


def test_newest_blizzard_takes_the_highest_version_across_products():
    older = BLIZZARD_VERSIONS.replace("1.60.1.70425", "1.60.1.70245").replace("70425|", "70245|")
    assert (
        newest_blizzard({"wowf": older, "wow_classic_beta": BLIZZARD_VERSIONS}).version
        == "1.60.1.70425"
    )


def test_the_committed_config_names_blizzard_products():
    assert "wow_classic_beta" in load_config().blizzard_products


# Launch-day shapes (rehearsal 2026-10-30). Forever goes live on 4 November 2026 and the
# live client may arrive under a key the config has never seen and with a new number.

#: wago.tools on 2026-10-30 still lists wow_classic_beta's Cataclysm and Mists betas
#: (3.4/4.4/5.5, 2022-2025): numerically newer than any 1.60 build, so before the
#: ingest-time rule every nightly warned about ~130 of them and buried a real rename.
HISTORIC_BETAS = {
    "wow_classic_beta": [
        {"version": "5.5.0.62071", "created_at": "2025-07-18 01:19:04"},
        {"version": "4.4.0.54737", "created_at": "2024-05-17 03:34:04"},
        {"version": "1.60.1.70532", "created_at": "2026-10-22 19:09:04"},
    ],
}


def test_unclaimed_builds_ingested_before_the_active_one_are_not_warned_about():
    result = check({**BUILDS, **HISTORIC_BETAS}, "1.60.1.70532", CONFIG)
    assert result.unmatched_newer == ()


def test_a_launch_renumbering_to_2_0_under_a_new_key_is_warned_about_not_activated():
    launch = {
        **BUILDS,
        **HISTORIC_BETAS,
        "wow_forever": [{"version": "2.0.0.71000", "created_at": "2026-11-04 15:00:00"}],
    }
    result = check(launch, "1.60.1.70532", CONFIG)
    assert [(b.product, b.version) for b in result.unmatched_newer] == [
        ("wow_forever", "2.0.0.71000")
    ]
    assert not result.newer


def test_a_launch_renumbering_under_a_listed_key_is_activated_with_its_interface():
    launch = {**BUILDS, "wowf": [{"version": "2.0.0.71000", "created_at": "2026-11-04 15:00:00"}]}
    out = report(check(launch, "1.60.1.70374", CONFIG))
    assert out["newer"] is True
    assert (out["newest_build"], out["newest_interface"]) == ("2.0.0.71000", 20000)


def test_a_minor_bump_under_a_renamed_key_is_activated():
    launch = {
        **BUILDS,
        "wow_forever": [{"version": "1.60.2.71000", "created_at": "2026-11-04 15:00:00"}],
    }
    result = check(launch, "1.60.1.70374", CONFIG)
    assert result.newer and result.newest == Build(
        "wow_forever", "1.60.2.71000", "2026-11-04 15:00:00"
    )
    assert report(result)["newest_interface"] == 16002


def test_a_forever_version_under_an_ignored_key_is_warned_about_never_silently_dropped():
    """Should the live client land on a key the config ignores (Era's own, say), a 1.60+
    build there is not Forever by rule, but it is not Era either: say so."""
    moved = {
        **BUILDS,
        "wow_classic_era": [
            {"version": "1.15.9.70003", "created_at": "2026-09-29 14:06:05"},
            {"version": "1.60.2.71000", "created_at": "2026-11-04 15:00:00"},
        ],
    }
    result = check(moved, "1.60.1.70374", CONFIG)
    assert not result.newer
    assert [(b.product, b.version) for b in result.unmatched_newer] == [
        ("wow_classic_era", "1.60.2.71000")
    ]


def test_a_forever_build_ingested_after_the_active_one_with_a_lower_number_is_reported():
    """A renumbering DOWN (live as 1.16.0, say) under a listed key is never `newer`; it must
    not pass as "nothing to do"."""
    down = {
        **BUILDS,
        "wowf": [
            {"version": "1.60.1.70245", "created_at": "2026-10-13 21:18:10"},
            {"version": "1.16.0.71000", "created_at": "2026-11-04 15:00:00"},
        ],
    }
    result = check(down, "1.60.1.70374", CONFIG)
    assert not result.newer
    assert [(b.product, b.version) for b in result.out_of_order] == [("wowf", "1.16.0.71000")]
    assert report(result)["out_of_order"] == [
        {"product": "wowf", "version": "1.16.0.71000", "created_at": "2026-11-04 15:00:00"}
    ]
    # The same build re-listed under a second key later is not out of order.
    relisted = {
        **BUILDS,
        "wowf": [{"version": "1.60.1.70374", "created_at": "2026-10-20 00:00:00"}],
    }
    assert check(relisted, "1.60.1.70374", CONFIG).out_of_order == ()


PTR_CONFIG = ClientConfig(
    products=CONFIG.products,
    ignore_products=CONFIG.ignore_products,
    min_version=CONFIG.min_version,
    test_product_pattern=r"ptr|test|^wow[a-z]?t$|^wowz$",
)


@pytest.mark.parametrize(
    "product", ["wowft", "wowf_ptr", "wow_forever_ptr", "wowflivetest", "wowz"]
)
def test_a_forever_ptr_or_test_key_carrying_1_6x_is_never_forever(product: str):
    """A Forever PTR is numbered like live (1.6x) and is newer than it; it must never move
    the site, so a test-shaped key loses to the version rule."""
    ptr = {**BUILDS, product: [{"version": "1.61.0.71500", "created_at": "2026-11-05 15:00:00"}]}
    result = check(ptr, "1.60.1.70374", PTR_CONFIG)
    assert not is_forever(Build(product, "1.61.0.71500", ""), PTR_CONFIG)
    assert not result.newer
    assert result.newest == Build("wow_cn_beta", "1.60.1.70374", "2026-10-13 23:59:08")
    # Nor is it warned about every night as an unclaimed launch client.
    assert result.unmatched_newer == ()


def test_the_ptr_rule_leaves_a_live_shaped_key_alone():
    assert is_forever(Build("wow_forever", "1.60.2.71000", ""), PTR_CONFIG)
    assert is_forever(Build("wowf", "1.60.2.71000", ""), PTR_CONFIG)


def test_a_listed_product_that_looks_like_a_ptr_is_refused_at_load(tmp_path: Path):
    path = tmp_path / "client.json"
    path.write_text(
        json.dumps(
            {
                "products": ["wowf", "wowft"],
                "min_version": "1.60.0.0",
                "test_product_pattern": "^wow[a-z]?t$",
            }
        ),
        encoding="utf-8",
    )
    with pytest.raises(ClientWatchError, match="wowft"):
        load_config(path)


def test_the_committed_config_never_activates_a_ptr_key():
    config = load_config()
    assert config.test_product_pattern, "curated/client.json names no test_product_pattern"
    for product in ("wowt", "wowxptr", "wow_classic_ptr", "wowft", "wowf_ptr", "wowlivetest"):
        assert not is_forever(Build(product, "1.60.2.71000", ""), config), product
    for product in config.products:
        assert is_forever(Build(product, "1.60.2.71000", ""), config), product
