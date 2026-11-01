package gg.foreversixty.app.system.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.model.WatchedCharacter
import gg.foreversixty.app.navigation.DeepLink
import java.time.Duration
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.Locale

/** The four static launcher shortcuts (spec F15.1) and the Assistant's features, by their stable ids. */
enum class AppShortcut(val id: String) {
    NextRaid("next_raid"),
    SignupIn("signup_in"),
    Upgrades("upgrades"),
    Search("search"),
    Bis("bis"),
    Guides("guides"),
    ;

    companion object {
        /** The extra a shortcut or Assistant intent carries to MainActivity. */
        const val EXTRA = "gg.foreversixty.app.SHORTCUT"

        fun of(id: String?): AppShortcut? = entries.firstOrNull { it.id == id }

        /** The Assistant's `OPEN_APP_FEATURE` feature names (shortcuts.xml), spoken or typed. */
        fun ofFeature(feature: String?): AppShortcut? = when (feature?.trim()?.lowercase()) {
            "next raid" -> NextRaid
            "my upgrades", "upgrades" -> Upgrades
            "best in slot", "bis" -> Bis
            "guides" -> Guides
            "search" -> Search
            else -> null
        }
    }
}

/** Which raid a shortcut or intent means (spec F15 Which raid). */
sealed interface RaidPick {
    data class One(val raid: RaidView) : RaidPick
    data class Which(val raids: List<RaidView>) : RaidPick
    data object None : RaidPick
}

object IntentRaid {
    private val WINDOW: Duration = Duration.ofHours(24)
    private val CLOSED = setOf("cancelled", "done")

    /**
     * The next raid starting within 24 hours; else the next planned one. Two or more within 24 hours
     * ask which. Cancelled, done and started raids are not candidates.
     */
    fun pick(planned: List<RaidView>, now: Instant): RaidPick {
        val upcoming = planned
            .filter { it.status !in CLOSED }
            .mapNotNull { raid -> parse(raid.startsAt)?.takeIf { !it.isBefore(now) }?.let { raid to it } }
            .sortedBy { it.second }
        val soon = upcoming.filter { Duration.between(now, it.second) <= WINDOW }.map { it.first }
        return when {
            soon.size >= 2 -> RaidPick.Which(soon)
            soon.size == 1 -> RaidPick.One(soon.single())
            upcoming.isNotEmpty() -> RaidPick.One(upcoming.first().first)
            else -> RaidPick.None
        }
    }

    private fun parse(at: String): Instant? = runCatching { Instant.parse(at) }.getOrNull()
}

/** What `Sign up In` did (spec F15.1): a toast, or a screen the app must open. */
sealed interface SignupShortcutResult {
    data class Signed(val zone: String) : SignupShortcutResult
    data object Locked : SignupShortcutResult
    data object Failed : SignupShortcutResult
    data object Offline : SignupShortcutResult
    data object NoRaid : SignupShortcutResult
    data object SignedOut : SignupShortcutResult
    data object NoCharacter : SignupShortcutResult
    data class NoGuild(val name: String) : SignupShortcutResult

    /** No spec saved: the app opens the raid page's spec ask (F4.3). */
    data class AskSpec(val guild: Long, val raid: Long) : SignupShortcutResult

    /** Two raids within 24 hours: the app opens the Raids tab, where both are (ruling, lane F15). */
    data object Which : SignupShortcutResult

    /** `system` is off: the shortcut opens the app as usual. */
    data object Off : SignupShortcutResult
}

/** The toast for a result, or the app link to open instead; exactly one is non-null. */
data class SignupShortcutAnswer(val toast: String?, val open: DeepLink?, val openApp: Boolean = false)

object SignupShortcutCopy {
    fun answer(result: SignupShortcutResult, res: Resources): SignupShortcutAnswer = when (result) {
        is SignupShortcutResult.Signed -> toast(res.getString(R.string.shortcut_signupToast, result.zone))
        SignupShortcutResult.Locked -> toast(res.getString(R.string.raid_locked))
        SignupShortcutResult.Failed -> toast(res.getString(R.string.common_failed))
        SignupShortcutResult.Offline -> toast(res.getString(R.string.intent_offline))
        SignupShortcutResult.NoRaid -> toast(res.getString(R.string.intent_noRaid))
        SignupShortcutResult.SignedOut -> toast(res.getString(R.string.intent_signedOut))
        SignupShortcutResult.NoCharacter -> toast(res.getString(R.string.intent_noCharacter))
        is SignupShortcutResult.NoGuild -> toast(res.getString(R.string.intent_noGuild, result.name))
        is SignupShortcutResult.AskSpec -> SignupShortcutAnswer(null, DeepLink.Raid(result.guild, result.raid, spec = ASK))
        SignupShortcutResult.Which -> SignupShortcutAnswer(null, null, openApp = true)
        SignupShortcutResult.Off -> SignupShortcutAnswer(null, null, openApp = true)
    }

    private fun toast(text: String) = SignupShortcutAnswer(text, null)

    private const val ASK = "ask"
}

/** One dynamic shortcut (spec F15.3): what the launcher lists and searches. */
data class DynamicShortcut(
    val id: String,
    val shortLabel: String,
    val longLabel: String,
    val link: DeepLink,
    /** The class crest to draw; null draws the calendar (the next raid). */
    val crestClass: String?,
)

/**
 * The dynamic shortcuts (spec F15 Android): the current character, the first two watched
 * characters (not the current one again) and the next raid, at most four, in that order.
 */
object DynamicShortcuts {
    const val MAX = 4
    private const val WATCHED = 2

    fun list(
        current: AccountCharacter?,
        watched: List<WatchedCharacter>,
        nextRaid: Pair<Long, RaidView>?,
        res: Resources,
        zone: ZoneId = ZoneId.systemDefault(),
    ): List<DynamicShortcut> {
        val sub = res.getString(R.string.shortcut_characterSub)
        // Rank order (lead ruling, review round 2): the current character, the next raid, then the
        // watched ones; the launcher's two dynamic slots then show the character and the raid.
        val mine = current?.let { character(it.key, it.region, it.ruleset, it.name, it.className, sub) }
        val raid = nextRaid?.let { (guild, view) -> raid(guild, view, res, zone) }
        val others = watched.filter { it.key != current?.key }.take(WATCHED).map { character(it.key, it.region, it.ruleset, it.name, it.className, sub) }
        return (listOfNotNull(mine, raid) + others).take(MAX)
    }

    private fun character(key: String, region: String, ruleset: String, name: String, className: String?, sub: String) = DynamicShortcut(
        id = "character:$key",
        shortLabel = name,
        longLabel = "$name · $sub",
        link = DeepLink.Player(region, ruleset, name),
        crestClass = className,
    )

    private fun raid(guild: Long, view: RaidView, res: Resources, zone: ZoneId): DynamicShortcut {
        val starts = Instant.parse(view.startsAt)
        return DynamicShortcut(
            id = "raid:next",
            shortLabel = res.getString(R.string.shortcut_nextRaid),
            longLabel = res.getString(R.string.spotlight_raidTitle, view.zoneName, IntentTime.weekdayShort(starts, zone), IntentTime.time(starts, zone)),
            link = DeepLink.Raid(guild, view.id),
            crestClass = null,
        )
    }
}

/** Times as answers and launcher labels say them, in the phone's zone: `Thu`, `Thursday`, `8 pm`, `8:30 pm`. */
object IntentTime {
    private val SHORT_DAY: DateTimeFormatter = DateTimeFormatter.ofPattern("EEE", Locale.US)
    private val DAY: DateTimeFormatter = DateTimeFormatter.ofPattern("EEEE", Locale.US)
    private val HOUR: DateTimeFormatter = DateTimeFormatter.ofPattern("h a", Locale.US)
    private val HOUR_MINUTE: DateTimeFormatter = DateTimeFormatter.ofPattern("h:mm a", Locale.US)

    fun weekdayShort(at: Instant, zone: ZoneId): String = at.atZone(zone).format(SHORT_DAY)

    fun weekday(at: Instant, zone: ZoneId): String = at.atZone(zone).format(DAY)

    fun time(at: Instant, zone: ZoneId): String {
        val local = at.atZone(zone)
        val pattern = if (local.minute == 0) HOUR else HOUR_MINUTE
        return local.format(pattern).replace("AM", "am").replace("PM", "pm")
    }
}
