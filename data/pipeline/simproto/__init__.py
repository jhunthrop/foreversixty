"""The engine's protobuf messages, generated from its own .proto files.

`pipeline.genproto` writes the `*_pb2` modules beside this one; nothing here is
hand-written. Import the two short names rather than the generated modules, so
a later proto split is one edit:

    from pipeline.simproto import apl, pb
"""

from pathlib import Path

from pipeline.simproto import apl_pb2 as apl
from pipeline.simproto import common_pb2 as pb

ENGINE_SHA = (Path(__file__).resolve().parents[2] / "proto" / "ENGINE_SHA").read_text().strip()

__all__ = ["ENGINE_SHA", "apl", "pb"]
