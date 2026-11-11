import Foundation

/// The only words the apps set in capitals: codes written that way wherever players see them (a realm
/// region `US`, a metric unit `DPS`). Labels and headings are sentence case (ruling N1);
/// `CaseGuardTests` fails on uppercasing anywhere else. Android: `Codes`.
public enum Codes {
    /// A realm region as the game writes it: `us` to `US`.
    public static func region(_ region: String) -> String { region.uppercased() }

    /// A metric's unit as the site writes it: `dps` to `DPS`, `hps` to `HPS`.
    public static func unit(_ unit: String) -> String { unit.uppercased() }
}
