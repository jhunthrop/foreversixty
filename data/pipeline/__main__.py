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
    elif args.command == "forever-talents":
        from pipeline.forever import write_forever_talents

        write_forever_talents(args.snapshot, args.from_build, args.build)
    elif args.command == "diff":
        from pipeline.diff import diff_builds

        diff_builds(args.from_build, args.to_build)
    return 0


if __name__ == "__main__":
    sys.exit(main())
