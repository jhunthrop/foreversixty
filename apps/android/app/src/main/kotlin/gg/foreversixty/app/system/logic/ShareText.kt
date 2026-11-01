package gg.foreversixty.app.system.logic

import android.content.Context
import android.content.Intent
import android.content.res.Resources
import android.net.Uri
import gg.foreversixty.app.R
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.Locale

/** One share (spec F15 Share): the sentence and the page link whose Open Graph card the site makes. */
data class ShareMessage(val text: String, val url: String) {
    /** What the share sheet sends: the sentence, then the link on its own line. */
    val body: String get() = "$text\n$url"
}

/** The tier lists a share can name (`/tiers`, `/tiers/tank`, `/tiers/healer`). */
enum class TierRole(val wire: String, val path: String, val word: Int) {
    Dps("dps", "", R.string.guides_role_dps),
    Tank("tank", "/tank", R.string.guides_role_tank),
    Healer("healer", "/healer", R.string.guides_role_healer),
    ;

    companion object {
        /** The tier list role of the Guides tab's chip (`dps`, `tank`, `healer`); DPS for anything else. */
        fun of(role: String): TierRole = entries.firstOrNull { it.wire == role } ?: Dps
    }
}

/** The share texts (spec §5.7 Share table), pure over their inputs. Only things with a working web page share. */
object ShareText {
    private const val SITE = "https://foreversixty.gg"
    private val DAY_MONTH: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM", Locale.US)

    /**
     * `Lightningbolltz, Elemental Shaman, level 60` and the character page [url]
     * (`/character/{region}/{ruleset}/{name}`). Missing a spec, class or level, the name alone.
     */
    fun character(res: Resources, url: String, name: String, spec: String?, className: String?, level: Int?): ShareMessage {
        val text = if (spec != null && className != null && level != null) {
            res.getString(R.string.share_character, name, spec, className, level)
        } else {
            name
        }
        return ShareMessage(text, url)
    }

    /** `Elemental Shaman guide` and `/guides/{class}/{spec}`. */
    fun guide(res: Resources, classSlug: String, specSlug: String, specName: String, className: String) = ShareMessage(
        res.getString(R.string.share_guide, specName, className),
        "$SITE/guides/${seg(classSlug)}/${seg(specSlug)}",
    )

    /** `DPS tier list` and `/tiers`, `/tiers/tank` or `/tiers/healer`. */
    fun tiers(res: Resources, role: TierRole) = ShareMessage(
        res.getString(R.string.share_tiers, res.getString(role.word)),
        "$SITE/tiers${role.path}",
    )

    /** `Onyxia's Lair · 6 Nov` and `/reports/{id}`; null for a guild or private report (no share action). */
    fun report(res: Resources, id: String, zone: String, at: Instant, visibility: String, tz: ZoneId = ZoneId.systemDefault()): ShareMessage? {
        if (visibility !in SHAREABLE) return null
        return ShareMessage(res.getString(R.string.share_report, zone, at.atZone(tz).format(DAY_MONTH)), "$SITE/reports/${seg(id)}")
    }

    /** The system share sheet with [message]. */
    fun send(context: Context, message: ShareMessage) {
        val send = Intent(Intent.ACTION_SEND).setType("text/plain").putExtra(Intent.EXTRA_TEXT, message.body)
        context.startActivity(Intent.createChooser(send, null))
    }

    private val SHAREABLE = setOf("public", "unlisted")

    private fun seg(part: String): String = Uri.encode(part)
}
