"""The gem signal.

The parity design's section 4.4 states outright that gems, sockets, socket
bonuses and the whole gem economy do not exist here: the 1.60 client's item
table carries vanilla's socket columns and **none of its 19,171 rows sets
one**. Every consumer downstream -- the planner's item cards, `simdb.bin`,
the Top Gear expander -- is built on that measurement.

The day a build arrives with a socketed item, that is a design decision to
make, not a column to ignore. So the pipeline stops here and names the
items, rather than emitting a build whose sockets nothing reads.
"""

from __future__ import annotations

from pipeline.csvio import populated

#: The ItemSparse columns that would hold a gem socket's colour. A client
#: that does not export them at all reads as "no sockets", which is what
#: `populated` returning None means.
SOCKET_COLUMNS = ("SocketType_0", "SocketType_1", "SocketType_2")

#: How many ids the error message spells out before it just counts.
SHOWN = 20


class SocketError(SystemExit):
    """An item has a gem socket, which nothing in this repository models."""


def socketed_item_ids(sparse_rows: list[dict[str, str]]) -> list[int]:
    """Every ItemSparse id with a non-zero socket colour, in id order."""
    found: list[int] = []
    for row in sparse_rows:
        for column in SOCKET_COLUMNS:
            value = populated(row, column)
            if value is not None and int(value) != 0:
                found.append(int(row["ID"]))
                break
    return sorted(found)


def check_no_sockets(sparse_rows: list[dict[str, str]], build: str) -> None:
    found = socketed_item_ids(sparse_rows)
    if not found:
        return
    shown = ", ".join(str(item_id) for item_id in found[:SHOWN])
    more = f" and {len(found) - SHOWN} more" if len(found) > SHOWN else ""
    raise SocketError(
        f"build {build} has {len(found)} item(s) with a gem socket: {shown}{more}. "
        "Nothing here models gems -- see the parity design's section 4.4, which "
        "records that no item on 1.60.1.69893 had one. Design gems before "
        "emitting this build."
    )
