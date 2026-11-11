package gg.foreversixty.app.home.flows

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.flows.ClockPart
import gg.foreversixty.app.flows.TallyItem
import gg.foreversixty.app.model.RaidBadge
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.model.RaidViewerCharacter
import gg.foreversixty.app.phrase
import gg.foreversixty.app.raids.logic.ComingCounts
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidFacts
import gg.foreversixty.app.raids.logic.ShortRole
import gg.foreversixty.app.raids.logic.SignupController
import gg.foreversixty.app.raids.logic.SignupState
import java.time.Duration
import java.time.Instant
import java.time.ZoneId

/**
 * The next raid in words (spec 2026-11-08 F1, E7): the countdown, who the answer is for, who is coming
 * and an officer's open items. Formats only; the raid view and the badge are the server's.
 */
class RaidCardWords(private val res: Resources, private val clock: RaidClock, private val now: Instant) {
    /**
     * Home's and the Raid tab's card for [raid]: while [controller] holds the raid, its answer (drawn at
     * once), the minutes its last write stored and its live state; else the raid's own sign-up. Late's
     * minutes stay open under the answers while Late is the answer, the chosen one selected (F1).
     */
    fun card(raid: RaidView, guild: String, characterKey: String?, controller: SignupController?, badge: RaidBadge?): RaidCardModel {
        val held = controller?.takeIf { it.raid.id == raid.id }
        val view = held?.raid ?: raid
        val answer = characterKey?.let { held?.block(it)?.shown } ?: ownAnswer(view, characterKey)
        val control = held?.let { c -> characterKey?.let { AnswerControl(c.segmentsEnabled(it), c.answerStatus(it), c.canRetry(it)) } }
        val card = card(view, guild, characterKey, answer, lateOpen = true, badge = badge, lateMinutes = ownLateMinutes(view, characterKey))
        return control?.let { card.copy(control = it) } ?: card
    }

    /**
     * [answer] is the shown answer (the controller's, else the raid's own sign-up), [lateMinutes] how
     * late a Late answer said; [badge] an officer's counts.
     */
    fun card(
        raid: RaidView,
        guild: String,
        characterKey: String?,
        answer: SignupState?,
        lateOpen: Boolean,
        badge: RaidBadge?,
        lateMinutes: Int? = null,
    ): RaidCardModel {
        val facts = RaidFacts(raid, now)
        val zone = runCatching { ZoneId.of(raid.guildTimezone) }.getOrDefault(clock.zone)
        val whenText = res.getString(R.string.home_raid_when, clock.shortDate(facts.startsAt), clock.guildClock(facts.startsAt, zone))
        val aside = whenText
        val (clock, clockLabel) = countdown(facts.startsAt)
        val who = namedCharacter(raid, characterKey)
        val name = who?.name.orEmpty()
        val role = roleWord(who?.role)
        val minutes = lateMinutes.takeIf { answer == SignupState.Late }
        val said = answer?.let { answerWord(it, minutes) }
        val asWho = who?.let { res.getString(R.string.raid_tab_as, it.name, role) }
        // AS5: once answered, the segmented control is the answer; the line only names the character (E7).
        val me = when {
            said != null -> asWho
            who == null -> res.getString(R.string.home_raid_notAnsweredAlone)
            else -> res.getString(R.string.home_raid_notAnswered, name, role)
        }
        val unansweredSub = if (who == null) {
            res.getString(R.string.home_raid_lineSubAlone, counts(raid))
        } else {
            res.getString(R.string.home_raid_lineSub, counts(raid), name, role)
        }
        return RaidCardModel(
            raidId = raid.id,
            icon = RaidIcons.of(raid),
            eyebrow = res.getString(R.string.home_raid_tonight, guild, raid.zoneName),
            clock = clock,
            clockLabel = clockLabel,
            aside = aside,
            title = res.getString(R.string.home_raid_lineTitle, raid.zoneName, whenText),
            sub = if (answer == null) unansweredSub else listOfNotNull(counts(raid), asWho).joinToString(SEPARATOR),
            me = me,
            answer = answer,
            lateOpen = lateOpen && answer == SignupState.Late,
            tally = roleTally(raid),
            ready = null,
            officer = badge?.takeIf { it.officer }?.let { officer(raid, it) },
            zone = raid.zoneName,
            place = res.getString(R.string.home_raid_place, guild, whenText),
            headLine = headLine(facts.startsAt, zone, whenText),
            facts = headFacts(raid, zone),
            who = asWho,
            lateMinutes = minutes,
        )
    }

    /** `In`, or `Late · 30 min` when a Late answer said how late (F1). */
    fun answerWord(answer: SignupState, lateMinutes: Int?): String =
        if (answer == SignupState.Late && lateMinutes != null) {
            res.getString(R.string.raid_answer_lateMinutes, lateMinutes)
        } else {
            res.getString(answer.label)
        }

    /** Within a day: `8:00 pm CST · pull in 2 h 48 min`; further out the date and time ([whenText]). */
    private fun headLine(start: Instant, zone: ZoneId, whenText: String): String {
        val until = Duration.between(now, start).coerceAtLeast(Duration.ZERO)
        if (until >= Duration.ofDays(1)) return whenText
        val hours = until.toHours().toInt()
        val minutes = (until.toMinutes() % MINUTES_PER_HOUR).toInt()
        val left = if (hours > 0) res.getString(R.string.raid_tab_untilHours, hours, minutes) else res.getString(R.string.raid_tab_untilMinutes, minutes)
        return res.getString(R.string.raid_tab_pullIn, clock.guildClock(start, zone), left)
    }

    /**
     * E7 (owner, 2026-11-09): the character the answer is for, named only when the account has more
     * than one character in the raid's guild; null when the answer alone says it.
     */
    fun namedCharacter(raid: RaidView, characterKey: String?): RaidViewerCharacter? {
        val mine = raid.viewer.characters
        if (mine.size < 2) return null
        return mine.firstOrNull { it.characterKey == characterKey } ?: mine.first()
    }

    /** The clock to [start] (`02 hours 48 min`) and the timer read aloud. */
    fun countdown(start: Instant): Pair<List<ClockPart>, String> {
        val until = Duration.between(now, start).coerceAtLeast(Duration.ZERO)
        val hours = until.toHours().toInt()
        val minutes = (until.toMinutes() % MINUTES_PER_HOUR).toInt()
        val parts = listOf(ClockPart(twoDigits(hours), res.getString(R.string.home_raid_hours)), ClockPart(twoDigits(minutes), res.getString(R.string.home_raid_min)))
        return parts to res.getString(R.string.home_raid_clock, hours, minutes)
    }

    /** `Damage`, `Healer`, `Tank`: how the answer names the role. */
    fun roleWord(role: String?): String = res.getString(
        when (role) {
            TANK -> R.string.raid_role_tank
            HEALER -> R.string.raid_role_healer
            else -> R.string.raid_group_damage
        },
    )

    /** The raid live card: pulled how long ago, the counts, your answer. */
    fun live(raid: RaidView, guild: String, characterKey: String?, answer: SignupState?): RaidLiveModel {
        val pulled = RelativeAge.of(RaidFacts(raid, now).startsAt, now).phrase(res)
        val who = namedCharacter(raid, characterKey)
        val said = answer?.let { answerWord(it, ownLateMinutes(raid, characterKey)) }
        val line = when {
            said == null -> res.getString(R.string.home_live_lineNoAnswer, pulled, counts(raid))
            who == null -> res.getString(R.string.home_live_lineAlone, pulled, counts(raid), said)
            else -> res.getString(R.string.home_live_line, pulled, counts(raid), said, who.name, roleWord(who.role))
        }
        return RaidLiveModel(raid.id, res.getString(R.string.home_raid_live, guild, raid.zoneName), line)
    }

    /** The viewer's own answer on the raid, for [characterKey]. */
    fun ownAnswer(raid: RaidView, characterKey: String?): SignupState? = own(raid, characterKey)?.let { SignupState.of(it.state) }

    /** How late the viewer's own Late answer said it would be; null for any other answer. */
    fun ownLateMinutes(raid: RaidView, characterKey: String?): Int? = own(raid, characterKey)?.lateMinutes

    private fun own(raid: RaidView, characterKey: String?) =
        raid.signups.firstOrNull { it.own && (characterKey == null || it.characterKey == characterKey) }

    private fun counts(raid: RaidView) = res.getString(R.string.home_raid_counts, raid.counts.signedIn, raid.counts.late)

    /** Tanks, healers and damage against the comp target (none without one): the one tally model (S6). */
    fun roleTally(raid: RaidView): List<TallyItem> {
        val target = raid.compTarget
        val c = ComingCounts.of(raid)
        return listOf(
            TallyItem(res.getString(R.string.raid_group_tanks), c.tanks, target?.tanks),
            TallyItem(res.getString(R.string.raid_group_healers), c.healers, target?.healers),
            TallyItem(res.getString(R.string.raid_group_damage), c.damage, target?.let { it.melee + it.ranged }),
        )
    }

    /** `Invites 7:45 pm · 40` (spec 4.2 head facts); null without an invite time or a size. */
    fun headFacts(raid: RaidView, zone: ZoneId): String? {
        val invite = raid.inviteAt?.let { runCatching { Instant.parse(it) }.getOrNull() } ?: return null
        val size = raid.size ?: return null
        return res.getString(R.string.raid_tab_inviteSize, clock.guildClock(invite, zone), size)
    }

    /**
     * `5 have not answered, 2 of them healers`, `Healers short by 3`, `1 applicant`. The not-answered count
     * is the raid's own list, the same people Who is coming's `N not answered` counts (row 6).
     */
    private fun officer(raid: RaidView, badge: RaidBadge): OfficerItems {
        val short = ShortRole.of(raid)
        val notAnswered = raid.notSigned.size
        val unansweredInRole = short?.let { role -> raid.notSigned.count { it.role == role.wire } } ?: 0
        val unanswered = when {
            notAnswered <= 0 -> null
            short != null && unansweredInRole > 0 ->
                res.getString(R.string.home_officer_unansweredRole, notAnswered, unansweredInRole, res.getString(short.lower))
            else -> res.getString(R.string.home_officer_unanswered, notAnswered)
        }
        val shortLine = short?.let { res.getString(R.string.home_officer_short, res.getString(it.group), it.missing) }
        val applicants = badge.applicants.takeIf { it > 0 }?.let { res.getQuantityString(R.plurals.home_officer_applicants, it, it) }
        return OfficerItems(unanswered, shortLine?.takeIf { unanswered != null }, applicants, null)
    }

    private fun twoDigits(n: Int) = n.toString().padStart(2, '0')

    private companion object {
        const val TANK = "tank"
        const val HEALER = "healer"
        const val MINUTES_PER_HOUR = 60L
        const val SEPARATOR = " · "
    }
}

/** A raid's zone icon: the server's `zone_icon`, else the icons of api/internal/zoneicons, else the generic book. */
object RaidIcons {
    private const val GENERIC = "inv_misc_book_09"
    private val ZONES = mapOf(
        "raid:barrow-deeps" to "spell_shadow_raisedead",
        "raid:hyjal-summit" to "spell_nature_timestop",
        "raid:onyxias-lair" to "inv_misc_head_dragon_black",
    )

    fun of(raid: RaidView): String = of(raid.zone, raid.zoneIcon)

    fun of(zoneKey: String, zoneIcon: String?): String = zoneIcon ?: ZONES[zoneKey] ?: GENERIC
}
