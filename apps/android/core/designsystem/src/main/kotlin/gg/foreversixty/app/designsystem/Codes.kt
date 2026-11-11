package gg.foreversixty.app.designsystem

import java.util.Locale

/**
 * The only words the apps set in capitals: codes that are written that way wherever players see
 * them (a realm region `US`, a metric unit `DPS`). Labels and headings are sentence case (ruling
 * N1); `CaseGuardTest` fails on uppercasing anywhere else. iOS: `Codes`.
 */
object Codes {
    /** A realm region as the game writes it: `us` to `US`. */
    fun region(region: String): String = region.uppercase(Locale.ROOT)

    /** A metric's unit as the site writes it: `dps` to `DPS`, `hps` to `HPS`. */
    fun unit(unit: String): String = unit.uppercase(Locale.ROOT)
}
