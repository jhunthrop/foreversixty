package gg.foreversixty.app.system.platform

import android.content.Context
import gg.foreversixty.app.BuildConfig
import gg.foreversixty.app.system.logic.ShareText
import java.time.Instant
import java.time.ZoneId

/**
 * Debug builds only: opens the real system share chooser with the app's own share of the spec's one
 * story (Lightningbolltz, the Elemental Shaman guide, the Onyxia's Lair report), built by the same
 * [ShareText] the screens use, for the review captures on the emulator. Started by
 * `adb shell am start ... --es fs.debug.share character|guide|report`.
 */
object DebugShare {
    const val EXTRA = "fs.debug.share"
    private const val REPORT_ID = "rpt-onyxia-1105"
    private val REPORT_AT: Instant = Instant.parse("2026-11-06T03:00:00Z")

    fun open(context: Context, what: String) {
        if (!BuildConfig.DEBUG) return
        val res = context.resources
        val message = when (what) {
            "character" -> ShareText.character(
                res, "https://foreversixty.gg/character/us/normal/lightningbolltz", "Lightningbolltz", "Elemental", "Shaman", 60,
            )
            "guide" -> ShareText.guide(res, "shaman", "elemental", "Elemental", "Shaman")
            "report" -> ShareText.report(res, REPORT_ID, "Onyxia's Lair", REPORT_AT, "public", ZoneId.of("America/Chicago"))
            else -> null
        } ?: return
        ShareText.send(context, message)
    }
}
