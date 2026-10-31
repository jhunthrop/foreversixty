package gg.foreversixty.app.widgets.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.WidgetNextRaid
import gg.foreversixty.app.model.WidgetUpgrades
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.SignupState
import java.time.Duration
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.Locale
import kotlin.math.ceil

/** Which words the raid surfaces show (spec §5.6, F14.2). */
enum class RaidSurfaceState { SignedOut, NoCharacter, NoGuild, NoRaid, Raid }

/** The foot tag's kind, which picks its colour: kill for In and Late, gold to act, wipe when cancelled. */
enum class TagKind { In, Late, Bench, Out, NotSignedUp, ChooseSpec, Locked, Cancelled }

/** Every word and figure the raid surfaces draw, from one snapshot and one clock. */
data class RaidFigures(
    val state: RaidSurfaceState,
    /** The headline line of a wordy state (`No guild yet`), else the zone. */
    val title: String,
    /** The second line of a wordy state, null for a raid. */
    val subtitle: String? = null,
    val guildId: Long? = null,
    val raidId: Long? = null,
    val characterName: String? = null,
    val tag: String? = null,
    val tagKind: TagKind? = null,
    val countdown: String? = null,
    val short: String? = null,
    val time: String? = null,
    val timeShort: String? = null,
    val inline: String,
    val you: String? = null,
    /** The answer alone (`In`, `Not signed up`, `Locked`): the inline widget's and the tile's last word. */
    val inlineAnswer: String? = null,
    val signedUp: String? = null,
    val signedUpCount: Int = 0,
    val size: Int? = null,
    val stale: String? = null,
    val buttons: Boolean = false,
    /** The answer checked on the 4 x 2 widget: the pending press, else the snapshot's. */
    val checked: SignupState? = null,
    /** The button whose press the server refused, saying `Not saved` for 10 s. */
    val notSaved: SignupState? = null,
    /** No spec saved: every button opens the app's spec ask instead of answering. */
    val asksSpec: Boolean = false,
    val liveActive: Boolean = false,
    val liveTitle: String? = null,
    val liveUntil: String? = null,
    /** The instant the countdown runs to: invites, else the pull. */
    val liveTarget: Instant? = null,
    val liveYou: String? = null,
    /** Cancelled inside the countdown's window: the notification says `Cancelled` for 5 minutes. */
    val liveCancelled: Boolean = false,
    val guildName: String? = null,
    val description: String,
)

/** One upgrade row: item, quality, icon and its gain words. */
data class UpgradeLine(val itemId: Int, val name: String, val quality: Int, val icon: String, val slot: String, val gain: String)

/** The upgrade surfaces' words. [sentence] replaces the rows for any state but ok. */
data class UpgradeFigures(
    val header: String,
    val rows: List<UpgradeLine>,
    val footer: String?,
    val rectTitle: String?,
    val sentence: String?,
    /** No Home answer behind the snapshot yet: the widget says `Open the app`. */
    val opensApp: Boolean,
    val description: String,
    /** The character's class, for the crest beside the label (one crest language). */
    val className: String? = null,
)

/**
 * The ONE derivation of every live surface's words (spec §5.6; golden vectors
 * apps/shared/widgets/figures.json, shared with iOS). The widgets, the ongoing raid notification
 * and the quick-settings tile call this with the same snapshot and clock, so the countdown, the
 * count and the answer agree everywhere (the wow-player's F14 SHOULD).
 */
class WidgetFigures(private val res: Resources, private val zone: ZoneId = ZoneId.systemDefault()) {
    private val clock = RaidClock(res, zone)

    fun raid(snapshot: WidgetSnapshot?, now: Instant, overlay: SignupOverlay? = null): RaidFigures {
        if (snapshot == null) return wordy(RaidSurfaceState.SignedOut, res.getString(R.string.widget_signedOut), null)
        val widget = snapshot.widget
        val stale = staleLine(snapshot.fetchedInstant, now)
        if (widget.character == null) return wordy(RaidSurfaceState.NoCharacter, res.getString(R.string.widget_noCharacter), null, stale)
        if (widget.guild == null) {
            return wordy(RaidSurfaceState.NoGuild, res.getString(R.string.widget_noGuild), res.getString(R.string.widget_noGuildSub), stale)
        }
        val raid = widget.nextRaid ?: return wordy(
            RaidSurfaceState.NoRaid,
            res.getString(R.string.widget_noRaid),
            widget.lastRaid?.let { res.getString(R.string.widget_lastRaid, it.zoneName, dayMonth(Instant.parse(it.startsAt))) },
            stale,
        )
        return raidFigures(raid, widget.character?.name.orEmpty(), now, stale, overlay?.takeIf { it.raidId == raid.raidId })
    }

    private fun wordy(state: RaidSurfaceState, title: String, subtitle: String?, stale: String? = null) =
        RaidFigures(state = state, title = title, subtitle = subtitle, inline = title, stale = stale,
            description = listOfNotNull(res.getString(R.string.widget_nextRaid), title, subtitle).joinToString(". "))

    private fun raidFigures(raid: WidgetNextRaid, name: String, now: Instant, stale: String?, overlay: SignupOverlay?): RaidFigures {
        val start = Instant.parse(raid.startsAt)
        val end = Instant.parse(raid.endsAt)
        val pending = overlay?.takeIf { it.isPending }?.signupState
        val state = pending ?: SignupState.of(raid.my.state)
        val role = roleLabel(raid.my.role)
        val asksSpec = raid.my.spec == null && state != SignupState.Out
        val answerShort = when {
            asksSpec -> res.getString(R.string.widget_chooseSpec)
            state == null -> res.getString(R.string.widget_notSignedUp)
            else -> res.getString(state.label)
        }
        val answer = if (!asksSpec && state != null) res.getString(R.string.widget_answer, answerShort, role) else answerShort
        val (tag, kind) = when {
            raid.status == STATUS_CANCELLED -> res.getString(R.string.widget_cancelled) to TagKind.Cancelled
            raid.status == STATUS_LOCKED -> res.getString(R.string.widget_locked) to TagKind.Locked
            asksSpec -> answer to TagKind.ChooseSpec
            state == null -> answer to TagKind.NotSignedUp
            // `In · DPS` names the role you come as; Late, Bench and Out stand alone (spec §5.6).
            state == SignupState.In -> answer to TagKind.In
            else -> answerShort to answerKind(state)
        }
        val inlineAnswer = if (kind == TagKind.Cancelled || kind == TagKind.Locked) tag else answerShort
        val cancelled = raid.status == STATUS_CANCELLED
        val span = if (cancelled) null else Countdown.of(start, end, now)
        val signedUp = raid.size?.let { res.getString(R.string.widget_ofSize, raid.signedUp, it) } ?: raid.signedUp.toString()
        val live = liveFigures(raid, start, now, state, cancelled)
        val inWindow = state?.isComing == true && inLiveWindow(raid, start, now)
        val time = weekday(start) + " " + clock.clock(start)
        val timeShort = weekday(start) + " " + shortClock(start)
        return RaidFigures(
            state = RaidSurfaceState.Raid, title = raid.zoneName, guildId = raid.guildId, raidId = raid.raidId,
            characterName = name, tag = tag, tagKind = kind,
            countdown = span?.let(::countdownWords), short = span?.let(::shortWords),
            time = time, timeShort = timeShort,
            inline = res.getString(R.string.widget_inline, raid.zoneName, timeShort, inlineAnswer), inlineAnswer = inlineAnswer,
            you = res.getString(R.string.widget_you, answer), signedUp = signedUp, signedUpCount = raid.signedUp, size = raid.size,
            stale = stale, buttons = raid.canSignUp && raid.status == STATUS_OPEN, checked = state,
            notSaved = overlay?.takeIf { it.isFailed(now) }?.signupState, asksSpec = asksSpec,
            liveActive = live != null, liveTitle = live?.first, liveUntil = live?.second,
            liveCancelled = cancelled && inWindow, guildName = raid.guildName,
            liveTarget = Instant.parse(raid.inviteAt ?: raid.startsAt),
            liveYou = state?.let { res.getString(R.string.live_you, res.getString(it.label), role, name) },
            description = description(raid, start, span, state, role, signedUp),
        )
    }

    /** The countdown's title and `until` line while the Live window is open, else null. */
    private fun liveFigures(raid: WidgetNextRaid, start: Instant, now: Instant, state: SignupState?, cancelled: Boolean): Pair<String, String>? {
        val target = Instant.parse(raid.inviteAt ?: raid.startsAt)
        if (cancelled || state?.isComing != true || !inLiveWindow(raid, start, now)) return null
        val minutes = ceilMinutes(Duration.between(now, target))
        // The title names invites only when the raid has an invite time; else the pull.
        val invites = raid.inviteAt != null
        val title = when {
            minutes <= 0 -> res.getString(if (invites) R.string.live_notifTitleNow else R.string.live_pullTitleNow, raid.zoneName)
            minutes == 1L -> res.getString(if (invites) R.string.live_notifTitleOne else R.string.live_pullTitleOne, raid.zoneName)
            else -> res.getString(if (invites) R.string.live_notifTitle else R.string.live_pullTitle, raid.zoneName, minutes.toInt())
        }
        val until = raid.inviteAt?.let { res.getString(R.string.live_invites, clock.clock(Instant.parse(it))) }
            ?: res.getString(R.string.live_starts, clock.clock(start))
        return title to until
    }

    private fun inLiveWindow(raid: WidgetNextRaid, start: Instant, now: Instant): Boolean {
        val target = Instant.parse(raid.inviteAt ?: raid.startsAt)
        return now >= target.minus(LIVE_LEAD) && now < start.plus(LIVE_TAIL)
    }

    private fun description(raid: WidgetNextRaid, start: Instant, span: Countdown?, state: SignupState?, role: String, signedUp: String): String {
        val whenLong = DateTimeFormatter.ofPattern("EEEE", LOCALE).format(start.atZone(zone)) + " " + shortClock(start)
        val countdown = span?.let { res.getString(R.string.widget_a11y_in, spokenSpan(it)) } ?: res.getString(R.string.widget_cancelled)
        val parts = mutableListOf(res.getString(R.string.widget_a11y_raid, raid.zoneName, whenLong, countdown))
        state?.let { parts += res.getString(R.string.widget_a11y_you, res.getString(it.label), role) }
        raid.size?.let { parts += res.getString(R.string.widget_a11y_signedUp, raid.signedUp, it) } ?: run { parts += signedUp }
        return parts.joinToString(" ")
    }

    fun upgrades(snapshot: WidgetSnapshot?): UpgradeFigures {
        val name = snapshot?.widget?.character?.name.orEmpty()
        val className = snapshot?.widget?.character?.className
        val header = res.getString(R.string.widget_upgrades, name)
        val upgrades = snapshot?.widget?.upgrades
            ?: return UpgradeFigures(header, emptyList(), null, null, res.getString(R.string.widget_openApp), opensApp = true,
                description = res.getString(R.string.widget_nextUpgrade), className = className)
        if (upgrades.state != STATE_OK) {
            val sentence = upgrades.sentence ?: res.getString(R.string.widget_noGear)
            return UpgradeFigures(header, emptyList(), null, null, sentence, opensApp = false,
                description = res.getString(R.string.widget_nextUpgrade) + ". " + sentence, className = className)
        }
        val unit = unitWord(upgrades.unit)
        val rows = upgrades.rows.map { row ->
            UpgradeLine(row.pick.id, row.pick.name, row.pick.quality, row.pick.icon, row.slotLabel,
                row.gain?.let { res.getString(R.string.widget_gain, tenths(it), unit) } ?: res.getString(R.string.widget_noGain))
        }
        val first = upgrades.rows.firstOrNull()
        return UpgradeFigures(
            header = header, rows = rows, footer = footer(upgrades, unit),
            rectTitle = first?.let { res.getString(R.string.widget_answer, res.getString(R.string.widget_nextUpgrade), it.slotLabel) },
            sentence = null, opensApp = false,
            description = rows.firstOrNull()?.let { res.getString(R.string.widget_a11y_upgrade, it.slot, it.name, it.gain) }
                ?: res.getString(R.string.widget_nextUpgrade),
            className = className,
        )
    }

    private fun footer(upgrades: WidgetUpgrades, unit: String): String {
        val count = res.getQuantityString(R.plurals.home_answer_upgrades, upgrades.count, upgrades.count)
        return upgrades.totalGain?.let { res.getString(R.string.widget_upgradesFooter, count, tenths(it), unit) } ?: count
    }

    private fun unitWord(unit: String): String = when (unit) {
        "dps" -> res.getString(R.string.bis_unit_dps)
        "hps" -> res.getString(R.string.bis_unit_hps)
        else -> unit.uppercase(LOCALE)
    }

    private fun roleLabel(role: String): String = when (role) {
        "tank" -> res.getString(R.string.raid_role_tank)
        "healer" -> res.getString(R.string.raid_role_healer)
        else -> res.getString(R.string.raid_role_dps)
    }

    private fun answerKind(state: SignupState): TagKind = when (state) {
        SignupState.In -> TagKind.In
        SignupState.Late -> TagKind.Late
        SignupState.Bench -> TagKind.Bench
        SignupState.Out -> TagKind.Out
    }

    private fun staleLine(fetched: Instant, now: Instant): String? {
        val hours = Duration.between(fetched, now).toHours()
        return if (hours >= STALE_HOURS) res.getString(R.string.widget_updated, hours.toInt()) else null
    }

    private fun countdownWords(span: Countdown): String = when (span) {
        Countdown.Now -> res.getString(R.string.widget_now)
        is Countdown.Left -> when {
            span.days >= 1 && span.hours > 0 && span.days == 1 -> res.getString(R.string.widget_in_dayHours, span.hours)
            span.days >= 1 && span.hours > 0 -> res.getString(R.string.widget_in_daysHours, span.days, span.hours)
            span.days >= 1 -> res.getQuantityString(R.plurals.widget_in_days, span.days, span.days)
            span.hours >= 1 && span.minutes > 0 -> res.getString(R.string.widget_in_hoursMinutes, span.hours, span.minutes)
            span.hours >= 1 -> res.getString(R.string.widget_in_hours, span.hours)
            else -> res.getString(R.string.widget_in_minutes, span.minutes)
        }
    }

    private fun shortWords(span: Countdown): String? = when (span) {
        Countdown.Now -> null
        is Countdown.Left -> when {
            span.days >= 1 -> res.getString(R.string.widget_short_days, span.days)
            span.hours >= 1 -> res.getString(R.string.widget_short_hours, span.hours)
            else -> res.getString(R.string.widget_short_minutes, span.minutes)
        }
    }

    private fun spokenSpan(span: Countdown): String = when (span) {
        Countdown.Now -> res.getString(R.string.widget_now)
        is Countdown.Left -> buildList {
            if (span.days > 0) add(res.getQuantityString(R.plurals.widget_a11y_days, span.days, span.days))
            if (span.hours > 0) add(res.getQuantityString(R.plurals.widget_a11y_hours, span.hours, span.hours))
            if (span.days == 0 && span.minutes > 0) add(res.getQuantityString(R.plurals.widget_a11y_minutes, span.minutes, span.minutes))
        }.joinToString(" ")
    }

    /** `3 Dec`, the day and month held together by a no-break space so the month never wraps alone. */
    private fun dayMonth(at: Instant): String = DateTimeFormatter.ofPattern("d MMM", LOCALE).format(at.atZone(zone)).replace(' ', '\u00A0')

    private fun weekday(at: Instant): String = DateTimeFormatter.ofPattern("EEE", LOCALE).format(at.atZone(zone))

    /** `8 pm` on the hour, else `8:30 pm`. */
    private fun shortClock(at: Instant): String {
        val pattern = if (at.atZone(zone).minute == 0) "h a" else "h:mm a"
        return DateTimeFormatter.ofPattern(pattern, LOCALE).format(at.atZone(zone)).replace("AM", "am").replace("PM", "pm")
    }

    private fun tenths(value: Double): String = String.format(Locale.US, "%.1f", value)

    companion object {
        const val STATUS_OPEN = "open"
        const val STATUS_LOCKED = "locked"
        const val STATUS_CANCELLED = "cancelled"
        private const val STATE_OK = "ok"
        private const val STALE_HOURS = 12L
        private val LOCALE: Locale = Locale.UK

        /** The raid countdown runs from an hour before invites to 15 minutes past the pull (spec §5.6). */
        val LIVE_LEAD: Duration = Duration.ofMinutes(60)
        val LIVE_TAIL: Duration = Duration.ofMinutes(15)

        fun ceilMinutes(duration: Duration): Long = ceil(duration.toMillis() / MILLIS_PER_MINUTE).toLong()

        private const val MILLIS_PER_MINUTE = 60_000.0
    }
}

/** The time to the pull: whole minutes rounded up, split into days, hours and minutes. */
sealed interface Countdown {
    data object Now : Countdown
    data class Left(val days: Int, val hours: Int, val minutes: Int) : Countdown

    companion object {
        private const val MINUTES_PER_DAY = 1440L
        private const val MINUTES_PER_HOUR = 60L

        /** Null once the raid has ended; [Now] between the pull and the end. */
        fun of(start: Instant, end: Instant, now: Instant): Countdown? {
            if (now >= end) return null
            if (now >= start) return Now
            val total = WidgetFigures.ceilMinutes(Duration.between(now, start))
            return Left(
                days = (total / MINUTES_PER_DAY).toInt(),
                hours = (total % MINUTES_PER_DAY / MINUTES_PER_HOUR).toInt(),
                minutes = (total % MINUTES_PER_HOUR).toInt(),
            )
        }
    }
}
