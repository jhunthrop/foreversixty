import XCTest

/// Review captures for F6 Talents in the running app: Home's Talents row, the pushed screen, the
/// pinned compare bar, a rank sheet with the tapped cell left visible, and the same at AX5. Runs
/// on the shared simulator only; the PNGs land in design/reviews/2026-10-31-mobile-talents. Not
/// part of `make test`.
///
/// Skipped since the flows spec (design/specs/2026-11-08-app-flows.md): Home lost its rows
/// beyond the raid, Next upgrades and New for you (appendix C, "Home's ... Next rows"), the
/// Talents row with them, and no menu tool or link pushes the Talents compare screen (the
/// planner owns talents). The captures come back if a route to the screen does.
@MainActor
final class TalentsCaptureTests: XCTestCase {
    private static let cut = "Home's Talents row was cut by the flows spec (2026-11-08-app-flows.md appendix C) and no menu tool or link opens the Talents screen"

    func testTalentsScreen() throws {
        throw XCTSkip(Self.cut)
    }

    func testTalentsScreenAtLargestText() throws {
        throw XCTSkip(Self.cut)
    }
}
