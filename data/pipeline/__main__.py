import argparse
import logging
import sys


def build_parser() -> argparse.ArgumentParser:
    p = argparse.ArgumentParser(prog="pipeline", description="Forever Sixty data pipeline")
    sub = p.add_subparsers(dest="command", required=True)

    f = sub.add_parser("fetch", help="download DB2 CSV exports for a build")
    f.add_argument("--product", required=True, help="wago.tools product key, e.g. wow_classic_era")
    f.add_argument("--build", help="build string like 1.15.7.61582; default: latest for product")

    n = sub.add_parser("normalize", help="normalize raw CSVs into JSON")
    n.add_argument("--build", required=True)

    i = sub.add_parser("icons", help="download and convert the icons a build refers to")
    i.add_argument("--build", required=True)

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

    sd = sub.add_parser("simdb", help="build the engine's SimDatabase for a build")
    sd.add_argument("--build", required=True)

    lt = sub.add_parser("loot", help="build the Droptimizer and Top Gear data for a build")
    lt.add_argument("--build", required=True)
    lt.add_argument(
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
    return p


def main(argv: list[str] | None = None) -> int:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s %(name)s: %(message)s")
    args = build_parser().parse_args(argv)
    if args.command == "fetch":
        from pipeline.wago import fetch_build

        fetch_build(args.product, args.build)
    elif args.command == "normalize":
        from pipeline.normalize import normalize_build

        result = normalize_build(args.build)
        if result.skipped:
            # The build directory is incomplete. Exiting non-zero stops the CI
            # job before it can commit and push a build the site cannot render.
            for reason in result.skipped:
                logging.getLogger("pipeline").error("build %s is missing %s", args.build, reason)
            return 1
    elif args.command == "icons":
        from pipeline.icons import icons_for_build

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
    elif args.command == "simdb":
        from pipeline.simdb import write_sim_database

        print(write_sim_database(args.build))
    elif args.command == "loot":
        from pathlib import Path

        from pipeline.loot import write_loot_files

        for path in write_loot_files(args.build, Path(args.engine)):
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
    return 0


if __name__ == "__main__":
    sys.exit(main())
