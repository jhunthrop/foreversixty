"""Summarises jobs.sh output: per job, effective HPS and its error, overheal, mana left, casts a
minute and each heal's casts a minute. Usage: python3 -I summarize.py <jobs output>..."""
import re
import sys

NUM = r"([0-9][0-9,]*\.?[0-9]*)"
# Rows that are not a cast the healer chose: totals, mana returns, consumables, a totem's ticks.
SKIP = ("total", "---", "OtherAction", "item ", "Healing Stream (10461)")


def jobs(text):
    for block in text.split("##### ")[1:]:
        name, _, body = block.partition("\n")
        yield name.strip(), body


def figure(body, label):
    m = re.search(r"^\| " + re.escape(label) + r" \| " + NUM, body, re.M)
    return float(m.group(1).replace(",", "")) if m else float("nan")


def casts(body, seconds):
    out, spend = {}, False
    for line in body.splitlines():
        if line.startswith("| Mana spend") or line.startswith("| Cast without a mana row"):
            spend = True
            continue
        if spend and not line.startswith("|"):
            spend = False
        if spend:
            m = re.match(r"^\| (.+?) \| " + NUM + r" \|", line)
            if m and not m.group(1).startswith(SKIP):
                out[m.group(1)] = out.get(m.group(1), 0) + float(m.group(2)) * 60 / seconds
    return out


def main():
    for path in sys.argv[1:]:
        text = open(path).read()
        for name, body in jobs(text):
            m = re.search(r"healer fight, \d+ iterations, (\d+) s", body)
            seconds = int(m.group(1)) if m else 0
            e = re.search(r"^\| effective HPS \| " + NUM + r" \+/- " + NUM, body, re.M)
            if not e or not seconds:
                print(f"{name}: no result")
                continue
            left = re.search(r"mana left at the end \| \S+ of \S+ \(([0-9.]+)%\)", body)
            c = casts(body, seconds)
            heals = {k: v for k, v in c.items() if v >= 0.05}
            print(f"{name}: eff {float(e.group(1)):.1f} +/- {float(e.group(2)):.2f}, raw {figure(body, 'raw HPS'):.0f}, "
                  f"overheal {figure(body, 'overheal'):.1f}%, mana left {left.group(1) if left else '-'}%, "
                  f"casts/min {sum(c.values()):.1f} [" + ", ".join(f"{k} {v:.1f}" for k, v in
                                                            sorted(heals.items(), key=lambda kv: -kv[1])) + "]")


if __name__ == "__main__":
    main()
