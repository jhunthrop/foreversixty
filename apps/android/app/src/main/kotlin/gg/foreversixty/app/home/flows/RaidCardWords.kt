package gg.foreversixty.app.home.flows

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.flows.ClockPart
import gg.foreversixty.app.flows.TallyItem
import gg.foreversixty.app.model.RaidBadge
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.phrase
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidFacts
import gg.foreversixty.app.raids.logic.ShortRole
import gg.foreversixty.app.raids.logic.SignupState
import java.time.Duration
import java.time.Instant
import java.time.ZoneId

/**
 * The next raid in words (spec 2026-11-08 F1, E7): the countdown, who the answer is for, who is coming
 * and an officer's open items. Formats only; the raid view and the badge are the server's.
 */
class RaidCardWords(private val res: Resources, private val clock: RaidClock, private val now: Instant) {
    /** [answer] is the shown answer (the controller's, else the raid's own sign-up); [badge] an officer's counts. */
    fun card(raid: RaidView, guild: String, characterKey: String?, answer: SignupState?, lateOpen: Boolean, badge: RaidBadge?): RaidCardModel {
        val facts = RaidFacts(raid, now)
        val zone = runCatching { ZoneId.of(raid.guildTimezone) }.getOrDefault(clock.zone)
        val whenText = res.getString(R.string.home_raid_when, clock.shortDate(facts.startsAt), clock.guildClock(facts.startsAt, zone))
        val aside = whenText
        val until = Duration.between(now, facts.startsAt).coerceAtLeast(Duration.ZERO)
        val hours = until.toHours().toInt()
        val minutes = (until.toMinutes() % MINUTES_PER_HOUR).toInt()
        val who = raid.viewer.characters.firstOrNull { it.characterKey == characterKey } ?: raid.viewer.characters.firstOrNull()
        val name = who?.name.orEmpty()
        val role = roleWord(who?.role)
        val me = answer?.let { res.getString(R.string.home_raid_answered, res.getString(it.label), name, role) }
            ?: res.getString(R.string.home_raid_notAnswered, name, role)
        return RaidCardModel(
            raidId = raid.id,
            icon = RaidIcons.of(raid),
            eyebrow = res.getString(R.string.home_raid_tonight, guild, raid.zoneName),
            clock = listOf(ClockPart(twoDigits(hours), res.getString(R.string.home_raid_hours)), ClockPart(twoDigits(minutes), res.getString(R.string.home_raid_min))),
            clockLabel = res.getString(R.string.home_raid_clock, hours, minutes),
            aside = aside,
            title = res.getString(R.string.home_raid_lineTitle, raid.zoneName, whenText),
            sub = if (answer == null) res.getString(R.string.home_raid_lineSub, counts(raid), name, role) else me,
            me = me,
            answer = answer,
            lateOpen = lateOpen && answer == SignupState.Late,
            tally = tally(raid),
            ready = null,
            officer = badge?.takeIf { it.officer }?.let { officer(raid, it) },
        )
    }

    /** The raid live card: pulled how long ago, the counts, your answer. */
    fun live(raid: RaidView, guild: String, characterKey: String?, answer: SignupState?): RaidLiveModel {
        val pulled = RelativeAge.of(RaidFacts(raid, now).startsAt, now).phrase(res)
        val who = raid.viewer.characters.firstOrNull { it.characterKey == characterKey } ?: raid.viewer.characters.firstOrNull()
        val line = if (answer != null && who != null) {
            res.getString(R.string.home_live_line, pulled, counts(raid), res.getString(answer.label), who.name, roleWord(who.role))
        } else {
            res.getString(R.string.home_live_lineNoAnswer, pulled, counts(raid))
        }
        return RaidLiveModel(raid.id, res.getString(R.string.home_raid_live, guild, raid.zoneName), line)
    }

    /** The viewer's own answer on the raid, for [characterKey]. */
    fun ownAnswer(raid: RaidView, characterKey: String?): SignupState? =
        raid.signups.firstOrNull { it.own && (characterKey == null || it.characterKey == characterKey) }?.let { SignupState.of(it.state) }

    private fun counts(raid: RaidView) = res.getString(R.string.home_raid_counts, raid.counts.signedIn, raid.counts.late)

    private fun tally(raid: RaidView): List<TallyItem> {
        val target = raid.compTarget ?: return emptyList()
        val c = raid.counts
        return listOf(
            TallyItem(res.getString(R.string.raid_group_tanks), c.tanks, target.tanks),
            TallyItem(res.getString(R.string.raid_group_healers), c.healers, target.healers),
            TallyItem(res.getString(R.string.raid_group_damage), c.melee + c.ranged, target.melee + target.ranged),
        )
    }

    /** `5 have not answered, 2 of them healers`, `Healers short by 3`, `1 applicant`. */
    private fun officer(raid: RaidView, badge: RaidBadge): OfficerItems {
        val short = ShortRole.of(raid)
        val unansweredInRole = short?.let { role -> raid.notSigned.count { it.role == role.wire } } ?: 0
        val unanswered = when {
            badge.unanswered <= 0 -> null
            short != null && unansweredInRole > 0 ->
                res.getString(R.string.home_officer_unansweredRole, badge.unanswered, unansweredInRole, res.getString(short.lower))
            else -> res.getString(R.string.home_officer_unanswered, badge.unanswered)
        }
        val shortLine = short?.let { res.getString(R.string.home_officer_short, res.getString(it.group), it.missing) }
        val applicants = badge.applicants.takeIf { it > 0 }?.let { res.getQuantityString(R.plurals.home_officer_applicants, it, it) }
        return OfficerItems(unanswered, shortLine?.takeIf { unanswered != null }, applicants, null)
    }

    private fun roleWord(role: String?): String = res.getString(
        when (role) {
            TANK -> R.string.raid_role_tank
            HEALER -> R.string.raid_role_healer
            else -> R.string.raid_group_damage
        },
    )

    private fun twoDigits(n: Int) = n.toString().padStart(2, '0')

    private companion object {
        const val TANK = "tank"
        const val HEALER = "healer"
        const val MINUTES_PER_HOUR = 60L
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
