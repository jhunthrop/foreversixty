import argparse
import sys


def build_parser() -> argparse.ArgumentParser:
    p = argparse.ArgumentParser(prog="pipeline", description="Forever Sixty data pipeline")
    sub = p.add_subparsers(dest="command", required=True)

    f = sub.add_parser("fetch", help="download DB2 CSV exports for a build")
    f.add_argument("--product", required=True, help="wago.tools product key, e.g. wow_classic_era")
    f.add_argument("--build", help="build string like 1.15.7.61582; default: latest for product")

    n = sub.add_parser("normalize", help="normalize raw CSVs into JSON")
    n.add_argument("--build", required=True)

    d = sub.add_parser("diff", help="diff two normalized builds")
    d.add_argument("--from", dest="from_build", required=True)
    d.add_argument("--to", dest="to_build", required=True)
    return p


def main(argv: list[str] | None = None) -> int:
    args = build_parser().parse_args(argv)
    if args.command == "fetch":
        from pipeline.wago import fetch_build

        fetch_build(args.product, args.build)
    elif args.command == "normalize":
        from pipeline.normalize import normalize_build

        normalize_build(args.build)
    elif args.command == "diff":
        from pipeline.diff import diff_builds

        diff_builds(args.from_build, args.to_build)
    return 0


if __name__ == "__main__":
    sys.exit(main())
