import argparse
import logging
import sys


def build_parser() -> argparse.ArgumentParser:
    p = argparse.ArgumentParser(prog="pipeline", description="Forever Sixty data pipeline")
    sub = p.add_subparsers(dest="command", required=True)

    f = sub.add_parser("fetch", help="download DB2 CSV exports for a build")
    f.add_argument("--product", required=True, help="wago.tools product key, e.g. wow_classic_era")
    f.add_argument("--build", help="build string like 1.15.7.61582; default: latest for product")

    fw = sub.add_parser(
        "fetch-wowhead",
        help="download wowhead's Forever gear planner, the items the client's ItemSparse lacks",
    )
    fw.add_argument("--build", required=True, help="the build whose raw/ receives the payload")

    fwt = sub.add_parser(
        "fetch-wowhead-talents",
        help="download wowhead's live Forever talent payload, which carries the hotfixed "
        "talent changes the client's DB2 tables lack; normalize overlays it onto the trees",
    )
    fwt.add_argument("--build", required=True, help="the build whose raw/ receives the payload")

    fcq = sub.add_parser(
        "fetch-classic-quest-levels",
        help="ONE-TIME (or occasional, when SOURCE_COMMIT is repinned): download "
        "cmangos/classic-db's quest_template table and merge it into the committed "
        "raw/quests/quest-levels.json `loot` reads, tagged source: classic-db",
    )
    fcq.add_argument("--build", required=True, help="the build whose raw/ receives the file")

    fcs = sub.add_parser(
        "fetch-classic-sources",
        help="ONE-TIME (or occasional, same pinned commit as fetch-classic-quest-levels): "
        "download cmangos/classic-db's loot/vendor/quest tables and write the committed "
        "raw/classicdb/sources.json `loot`/`loot-merge` read, tagged source: classic-db",
    )
    fcs.add_argument("--build", required=True, help="the build whose raw/ receives the file")

    qlv = sub.add_parser(
        "quest-levels",
        help="THE NIGHTLY STEP (.github/workflows/bis.yml, before `make bis`): fetch "
        "wowhead pages ONLY for quest ids classic-db does not cover, politely and capped "
        "at --max-pages, and merge them into quest-levels.json tagged source: wowhead",
    )
    qlv.add_argument("--build", required=True, help="the build to fill in quest levels for")
    qlv.add_argument(
        "--engine",
        required=True,
        help="path to the wowsims-forever checkout, for the quest ids loot.json will need",
    )
    qlv.add_argument(
        "--max-pages",
        type=int,
        default=200,
        help="cap on live wowhead requests this run sends (default 200); remaining ids "
        "stay missing (item_level_proxy fallback) for a later run to pick up",
    )

    vwq = sub.add_parser(
        "verify-wowhead-quests",
        help="SPOT-CHECK ONLY, never run in CI: fetch a sample of wowhead's Forever quest "
        "pages (politely, with backoff -- wowhead throttles this address) and report how "
        "often their min_level/level agree with the committed quest-levels.json",
    )
    vwq.add_argument("--build", required=True, help="the build to verify against")
    vwq.add_argument(
        "--engine",
        required=True,
        help="path to the wowsims-forever checkout, for the quest ids to sample from",
    )
    vwq.add_argument(
        "--sample",
        type=int,
        default=20,
        help="how many quest ids to fetch and compare (default 20)",
    )

    ivs = sub.add_parser(
        "item-sources",
        help="THE NIGHTLY STEP (.github/workflows/bis.yml, before `fetch`/`loot`): fetch "
        "wowhead item pages ONLY for real (non-placeholder) item ids the COMMITTED "
        "loot.json names no source for at all, politely and capped at --max-pages, and "
        "merge them into the committed raw/items/item-sources.json `loot` reads. Needs "
        "no raw/ CSVs and no engine checkout -- items/<class>.json and loot.json are "
        "both already committed for the active build.",
    )
    ivs.add_argument("--build", required=True, help="the build to fill in item sources for")
    ivs.add_argument(
        "--max-pages",
        type=int,
        default=200,
        help="cap on live wowhead requests this run sends (default 200); remaining ids "
        "stay unsourced for a later run to pick up",
    )
    ivs.add_argument(
        "--only-new",
        action="store_true",
        help="fetch ONLY Forever-new item ids (id >= pipeline.loot.wowhead."
        "FOREVER_NEW_ID_THRESHOLD) -- for measuring or resuming just this lane's actual "
        "target without spending the page budget on Classic ids first",
    )

    hf = sub.add_parser(
        "hotfixes",
        help="decode a copied DBCache.bin ('XFTH' hotfix cache) into raw/hotfixes/"
        "{ItemSparse,Item}.csv for `normalize` to merge over the shipped tables -- "
        "never fetches the cache itself; the controller copies it from the client",
    )
    hf.add_argument("--build", required=True)
    hf.add_argument("--cache", required=True, help="path to the copied DBCache.bin")

    n = sub.add_parser("normalize", help="normalize raw CSVs into JSON")
    n.add_argument("--build", required=True)
    n.add_argument(
        "--allow-shrink",
        action="store_true",
        help="skip the ItemSparse-completeness and items/<class>.json-shrink gates "
        "(csvio.check_item_sparse_completeness / normalize._check_class_items_not_shrunk) "
        "for a deliberate re-baseline; logs loudly when used",
    )
    n.add_argument(
        "--engine",
        help="path to a wowsims-forever checkout; an item the client states no icon "
        "for falls back to the fork's own assets/database/db.json before the wowhead "
        "gear-planner payload's. Optional: without it, only the wowhead payload is tried",
    )

    inm = sub.add_parser(
        "itemnames", help="write itemnames.json from a build's committed items (no raw/ needed)"
    )
    inm.add_argument("--build", required=True)

    i = sub.add_parser("icons", help="download and convert the icons a build refers to")
    i.add_argument("--build", required=True)
    i.add_argument(
        "--engine",
        help="path to a wowsims-forever checkout; before downloading art, rewrites any "
        "items/<class>.json row still on the placeholder icon using the fork's own "
        "assets/database/db.json and the build's wowhead gear-planner payload (in that "
        "order). Optional: without it, only the wowhead payload is tried",
    )

    a = sub.add_parser("tree-art", help="download and process each talent tree's background")
    a.add_argument("--build", required=True)

    gt = sub.add_parser("gametables", help="fetch the client's GameTables for a build")
    gt.add_argument("--build", required=True)

    ft = sub.add_parser(
        "forever-talents",
        help="build Forever talent files from a Wowhead snapshot (pre-beta only)",
    )
    ft.add_argument(
        "--snapshot",
        default="data/raw-forever/wowhead-talents-2026-09-14.json",
        help="the saved Wowhead payload; see data/raw-forever/README.md",
    )
    ft.add_argument(
        "--from-build", required=True, help="the build whose class and tree tables to reuse"
    )
    ft.add_argument("--build", required=True, help="the build directory to write")

    d = sub.add_parser("diff", help="diff two normalized builds")
    d.add_argument("--from", dest="from_build", required=True)
    d.add_argument("--to", dest="to_build", required=True)

    wd = sub.add_parser(
        "wowhead-diff", help="diff a build's talent trees against the Wowhead snapshot"
    )
    wd.add_argument(
        "--snapshot",
        default="data/raw-forever/wowhead-talents-2026-09-14.json",
        help="the saved Wowhead payload; see data/raw-forever/README.md",
    )
    wd.add_argument("--build", required=True)

    g = sub.add_parser("simproto", help="re-vendor the engine protos and regenerate bindings")
    g.add_argument(
        "--engine",
        required=True,
        help="path to the wowsims-forever checkout, e.g. $FOREVER_ENGINE_PATH",
    )

    sc = sub.add_parser("simconst", help="write per-class spell constants for a build")
    sc.add_argument("--build", required=True)

    tr = sub.add_parser(
        "trainables",
        help="write per-class trainable abilities (builds/<build>/trainables/) from the "
        "client's SkillLineAbility; needs raw/ from `fetch`",
    )
    tr.add_argument("--build", required=True)

    lv = sub.add_parser(
        "levels",
        help="write levels.json (per-level base stats) and spellranks.json (rank "
        "chains) for a build; run after simconst",
    )
    lv.add_argument("--build", required=True)

    sd = sub.add_parser("simdb", help="build the engine's SimDatabase for a build")
    sd.add_argument("--build", required=True)

    lt = sub.add_parser("loot", help="build the Droptimizer and Top Gear data for a build")
    lt.add_argument("--build", required=True)
    lt.add_argument(
        "--engine",
        required=True,
        help="path to the wowsims-forever checkout, e.g. $FOREVER_ENGINE_PATH",
    )
    lt.add_argument(
        "--allow-shrink",
        action="store_true",
        help="skip the ItemSparse-completeness gate (csvio.check_item_sparse_completeness) "
        "for a deliberate re-baseline; logs loudly when used",
    )

    lm = sub.add_parser(
        "loot-merge",
        help="THE NIGHTLY FALLBACK when raw/ CSVs are not available (night-fetch-guard): "
        "re-derive loot.json ALONE from committed inputs -- engine fork, zones.json, "
        "items.json, quest-levels.json, item-sources.json, classicdb/sources.json -- "
        "no raw/ItemSparse.csv or raw/Map.csv needed. Run the full `loot` command instead "
        "whenever raw/ is healthy; this never writes enchants/suffixes/simbuffs/items "
        "columns, only loot.json.",
    )
    lm.add_argument("--build", required=True)
    lm.add_argument(
        "--engine",
        required=True,
        help="path to the wowsims-forever checkout, e.g. $FOREVER_ENGINE_PATH",
    )

    sp = sub.add_parser("specs", help="generate the Go and TypeScript spec lists")
    sp.add_argument("--go", default="../sim/specs/specs.go")
    sp.add_argument("--ts", default="../web/src/lib/sim/specs.ts")
    sp.add_argument(
        "--check",
        action="store_true",
        help="write nothing; exit non-zero if a generated file has drifted",
    )

    ph = sub.add_parser("phases", help="emit the phase calendar for the web")
    ph.add_argument("--web", default="../web/src/data/phases.json")
    ph.add_argument(
        "--check",
        action="store_true",
        help="write nothing; exit non-zero if the emitted file has drifted",
    )

    w = sub.add_parser("weights", help="emit a build's copy of the curated stat weights")
    w.add_argument("--build", required=True)
    w.add_argument(
        "--check",
        action="store_true",
        help="write nothing; exit non-zero if the emitted file has drifted",
    )

    ad = sub.add_parser("addon-data", help="emit addon-data.json and the addon's Data.lua")
    ad.add_argument("--build", required=True)
    ad.add_argument(
        "--check",
        action="store_true",
        help="write nothing; exit non-zero if either emitted file has drifted",
    )

    au = sub.add_parser(
        "audit",
        help="THE ACCURACY-AUDIT INSTRUMENT (tenet 8, docs/tenets.md): compare every "
        "published fact -- items, quests, drops, vendors, crafted recipes, leveling BiS, "
        "the addon and the curated specs -- against its primary source (client tables, "
        "classic-db, the engine). Read-only over data/builds; always exits 0, the caller "
        "reads audit.json/audit.md to decide.",
    )
    au.add_argument("--build", required=True)
    au.add_argument(
        "--engine", help="path to a wowsims-forever checkout, reserved for a future check"
    )
    au.add_argument(
        "--classicdb-dump",
        help="path to the pinned cmangos/classic-db mysqldump (.sql or .sql.gz) for the "
        "checks that need raw tables `raw/classicdb/sources.json` does not carry (boss spawn "
        "maps, vendor item_template rep columns, crafted-item recipe spells); omit to skip "
        "just those sub-checks",
    )
    au.add_argument(
        "--out",
        default=None,
        help="directory to write audit.json/audit.md into (default: a scratch dir under "
        "the system temp directory, never data/builds)",
    )
    return p


def main(argv: list[str] | None = None) -> int:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s %(name)s: %(message)s")
    args = build_parser().parse_args(argv)
    if args.command == "fetch":
        from pipeline.wago import fetch_build

        fetch_build(args.product, args.build)
    elif args.command == "fetch-wowhead":
        from pipeline.wowhead_items import fetch_wowhead

        fetch_wowhead(args.build)
    elif args.command == "fetch-wowhead-talents":
        from pipeline.wowhead_talents import fetch_wowhead_talents

        fetch_wowhead_talents(args.build)
    elif args.command == "fetch-classic-quest-levels":
        from pipeline.quest_levels import merge_classic_db

        stats = merge_classic_db(args.build)
        print(
            f"fetch-classic-quest-levels: merged {stats.added} classic-db entries "
            f"({stats.total} total)"
        )
    elif args.command == "fetch-classic-sources":
        from pipeline.classic_sources import fetch_and_write_classic_sources

        path = fetch_and_write_classic_sources(args.build)
        print(path)
    elif args.command == "quest-levels":
        import json
        from pathlib import Path

        from pipeline.forkdb import load_fork_database
        from pipeline.loot.sources import quest_ids_for_build
        from pipeline.quest_levels import fetch_missing_from_wowhead

        build_dir = Path("builds") / args.build
        item_rows = json.loads((build_dir / "items.json").read_text(encoding="utf-8"))
        build_items = {int(row["id"]) for row in item_rows}
        fork = load_fork_database(Path(args.engine))
        ids = quest_ids_for_build(fork, build_items)
        stats = fetch_missing_from_wowhead(args.build, ids, max_pages=args.max_pages)
        print(
            f"quest-levels: {stats.fetched}/{stats.needed} ids classic-db lacked were filled "
            f"from wowhead ({stats.still_missing} still missing; item_level_proxy fallback "
            "applies until a later run picks them up)"
        )
    elif args.command == "verify-wowhead-quests":
        import json
        from pathlib import Path

        from pipeline.forkdb import load_fork_database
        from pipeline.loot.sources import quest_ids_for_build
        from pipeline.quest_levels import load_quest_levels
        from pipeline.wowhead_quests import compare_with_classic_db, fetch_quest_levels

        build_dir = Path("builds") / args.build
        item_rows = json.loads((build_dir / "items.json").read_text(encoding="utf-8"))
        build_items = {int(row["id"]) for row in item_rows}
        fork = load_fork_database(Path(args.engine))
        all_ids = sorted(quest_ids_for_build(fork, build_items))
        sample = all_ids[:: max(1, len(all_ids) // args.sample)][: args.sample]
        result = fetch_quest_levels(args.build, sample, root=build_dir.parent)
        quest_levels = load_quest_levels(build_dir)
        stats = compare_with_classic_db(result.levels, quest_levels)
        print(
            f"verify-wowhead-quests: fetched {len(result.levels)}/{len(sample)} sample ids "
            f"({len(result.missing)} missing); compared {stats.compared} against "
            f"quest-levels.json: min_level {stats.min_level_agrees}/{stats.compared}, "
            f"level {stats.level_agrees}/{stats.compared}, both {stats.both_agree}/{stats.compared}"
        )
        for quest_id, w_min, w_level, c_min, c_level in stats.mismatches:
            print(
                f"  mismatch quest {quest_id}: wowhead min={w_min} level={w_level}, "
                f"classic-db min={c_min} level={c_level}"
            )
    elif args.command == "item-sources":
        from pathlib import Path

        from pipeline.item_sources import fetch_missing_from_wowhead
        from pipeline.loot.wowhead import (
            FOREVER_NEW_ID_THRESHOLD,
            load_class_item_rows,
            named_items_in_committed_loot,
            unsourced_real_item_ids,
        )

        build_dir = Path("builds") / args.build
        named = named_items_in_committed_loot(build_dir)
        # Ordered Forever-new-then-Classic, ascending required_level within
        # each bucket (see pipeline.loot.wowhead._crawl_priority_key);
        # --only-new narrows to just the Forever-new bucket.
        ids = unsourced_real_item_ids(load_class_item_rows(build_dir), named)
        if args.only_new:
            ids = [item_id for item_id in ids if item_id >= FOREVER_NEW_ID_THRESHOLD]
        stats = fetch_missing_from_wowhead(args.build, ids, max_pages=args.max_pages)
        print(
            f"item-sources: wowhead resolved a page for {stats.fetched}/{stats.needed} "
            f"unsourced real item ids ({stats.with_source}/{stats.fetched} fetched pages "
            f"named any source; {stats.still_missing} still missing; rerun later to resume "
            "from the warm cache)"
        )
    elif args.command == "itemnames":
        from pipeline.normalize.itemnames import write_item_names_from_build

        write_item_names_from_build(args.build)
    elif args.command == "hotfixes":
        from pathlib import Path

        from pipeline.hotfix_merge import write_build_hotfix_tables

        for result in write_build_hotfix_tables(args.build, Path(args.cache)):
            print(
                f"hotfixes {result.table}: {result.valid} valid, {result.removed} removed, "
                f"{result.invalid} invalid, {result.not_public} not_public records "
                f"({result.new_ids} new ids, {result.overriding_ids} overriding shipped "
                f"rows) -> {result.path}"
            )
    elif args.command == "normalize":
        from pathlib import Path

        from pipeline.normalize import normalize_build

        engine = Path(args.engine) if args.engine else None
        result = normalize_build(args.build, allow_shrink=args.allow_shrink, engine=engine)
        if result.skipped:
            # The build directory is incomplete. Exiting non-zero stops the CI
            # job before it can commit and push a build the site cannot render.
            for reason in result.skipped:
                logging.getLogger("pipeline").error("build %s is missing %s", args.build, reason)
            return 1
    elif args.command == "icons":
        from pathlib import Path

        from pipeline.icons import icons_for_build
        from pipeline.icons_fix import fix_icons_for_build

        engine = Path(args.engine) if args.engine else None
        for result in fix_icons_for_build(args.build, engine=engine):
            if result.before_placeholder:
                print(
                    f"icons: {result.class_slug} {result.before_placeholder} -> "
                    f"{result.after_placeholder} placeholder icons "
                    f"({result.fixed_by_fork} fixed via fork db, "
                    f"{result.fixed_by_wowhead} via wowhead)"
                )
        icons_for_build(args.build)
    elif args.command == "tree-art":
        from pipeline.art import backgrounds_for_build

        backgrounds_for_build(args.build)
    elif args.command == "gametables":
        from pipeline.gametables import write_game_tables

        print(write_game_tables(args.build))
    elif args.command == "forever-talents":
        from pipeline.forever import write_forever_talents

        write_forever_talents(args.snapshot, args.from_build, args.build)
        from pipeline.forever import fetch_missing_icons

        fetch_missing_icons(args.build)
    elif args.command == "diff":
        from pipeline.diff import diff_builds

        diff_builds(args.from_build, args.to_build)
    elif args.command == "wowhead-diff":
        from pipeline.wowhead_diff import write_snapshot_diff

        write_snapshot_diff(args.snapshot, args.build)
    elif args.command == "simconst":
        from pipeline.simconst import write_spell_constants

        print(write_spell_constants(args.build))
    elif args.command == "trainables":
        from pipeline.trainables import write_trainables

        print(write_trainables(args.build))
    elif args.command == "levels":
        from pipeline.levels import write_levels
        from pipeline.spellranks import write_spell_ranks

        levels_path = write_levels(args.build)
        if levels_path is None:
            # No wowhead payload for this build (Classic Era, or fetch-wowhead
            # has not run yet): levels.json needs a source it has no other
            # one for, so it is skipped rather than written empty or stale.
            logging.getLogger("pipeline").warning(
                "levels.json not written for build %s: no wowhead payload in raw/",
                args.build,
            )
        else:
            print(levels_path)
        print(write_spell_ranks(args.build))
    elif args.command == "simdb":
        from pipeline.simdb import write_sim_database

        print(write_sim_database(args.build))
    elif args.command == "loot":
        from pathlib import Path

        from pipeline.loot import write_loot_files

        for path in write_loot_files(
            args.build, Path(args.engine), allow_shrink=args.allow_shrink
        ):
            print(path)
    elif args.command == "loot-merge":
        from pathlib import Path

        from pipeline.loot import merge_loot_files

        for path in merge_loot_files(args.build, Path(args.engine)):
            print(path)
    elif args.command == "simproto":
        from pathlib import Path

        from pipeline.genproto import refresh

        print(refresh(Path(args.engine), Path("proto"), Path("pipeline/simproto")))
    elif args.command == "specs":
        from pathlib import Path

        from pipeline.specs import check_specs, write_specs

        if args.check:
            stale = check_specs(Path("curated"), Path(args.go), Path(args.ts))
            for path in stale:
                logging.getLogger("pipeline").error(
                    "%s does not match curated/specs.json; run `python -m pipeline specs`", path
                )
            return 1 if stale else 0
        for path in write_specs(Path("curated"), Path(args.go), Path(args.ts)):
            print(path)
    elif args.command == "phases":
        from pathlib import Path

        from pipeline.phases import check_phases, write_phases

        if args.check:
            if check_phases(Path("curated"), Path(args.web)):
                logging.getLogger("pipeline").error(
                    "%s does not match curated/phases.json; run `python -m pipeline phases`",
                    args.web,
                )
                return 1
            return 0
        print(write_phases(Path("curated"), Path(args.web)))
    elif args.command == "weights":
        from pipeline.weights import check_weights, write_weights

        if args.check:
            if check_weights(args.build):
                logging.getLogger("pipeline").error(
                    "builds/%s/stat-weights.json does not match curated/stat-weights.json; "
                    "run `python -m pipeline weights --build %s`",
                    args.build,
                    args.build,
                )
                return 1
            return 0
        print(write_weights(args.build))
    elif args.command == "addon-data":
        from pipeline.addondata import check_addon_data, write_addon_data
        from pipeline.addonlua import lua_has_drifted, write_lua

        if args.check:
            log = logging.getLogger("pipeline")
            stale = False
            if check_addon_data(args.build):
                log.error("builds/%s/addon-data.json has drifted", args.build)
                stale = True
            if lua_has_drifted(args.build):
                log.error("addon/ForeverSixty/Data.lua has drifted")
                stale = True
            if stale:
                log.error("run `python -m pipeline addon-data --build %s`", args.build)
                return 1
            return 0
        print(write_addon_data(args.build))
        print(write_lua(args.build))
    elif args.command == "audit":
        import tempfile
        from pathlib import Path

        from pipeline.audit import run_audit, summary_line, write_report

        out_dir = Path(args.out) if args.out else Path(tempfile.mkdtemp(prefix="forever-audit-"))
        dump = Path(args.classicdb_dump) if args.classicdb_dump else None
        engine = Path(args.engine) if args.engine else None
        results = run_audit(args.build, engine=engine, classicdb_dump=dump)
        json_path, md_path = write_report(results, out_dir, args.build)
        print(f"audit: wrote {json_path} and {md_path}")
        print(summary_line(results))
    return 0


if __name__ == "__main__":
    sys.exit(main())
