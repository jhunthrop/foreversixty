package gg.foreversixty.app.raids.flows

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.flows.ClockPart
import gg.foreversixty.app.home.flows.RaidCardWords
import gg.foreversixty.app.home.flows.RaidIcons
import gg.foreversixty.app.model.RaidCompTarget
import gg.foreversixty.app.model.RaidNightPerson
import gg.foreversixty.app.model.RaidNightSummary
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.SignupState
import java.time.Instant

/** What the raid page on raid night draws (mocks raid-night-live, raid-night-signup). */
data class RaidNightPageContent(
    val title: String,
    val icon: String,
    val facts: List<String>,
    val eyebrow: String,
    val clock: List<ClockPart>,
    val clockLabel: String,
    /** The green dot: the raid is on or starts within the hour (`raid.starting`). */
    val live: Boolean,
    val aside: String,
    /** `You are In` (kill green) or `Not signed up` (muted). */
    val tag: String,
    val answered: Boolean,
    val selected: SignupState?,
    val lateOpen: Boolean,
    /** Tanks, healers and damage against the comp target, with names. */
    val roles: List<Pair<String, String>>,
    val ready: List<ReadyItem>,
    val readyStamp: String?,
)

/**
 * The raid page on raid night (spec F2) from `GET /v1/me/raid-nights/{raid}` before the log: the
 * countdown with your answer, Your sign-up, who is coming by role, Before you pull. [target] is the
 * raid's comp target when the raid view is at hand.
 */
class RaidNightPagePresenter(private val res: Resources, private val clock: RaidClock, private val now: Instant) {
    fun content(night: RaidNightSummary, target: RaidCompTarget?, pressed: SignupState? = null, lateOpen: Boolean = false): RaidNightPageContent {
        val start = Instant.parse(night.startsAt)
        val time = res.getString(R.string.home_raid_when, clock.shortDate(start), clock.guildClock(start, clock.zone))
        val words = RaidCardWords(res, clock, now)
        val (parts, label) = if (night.phase == LIVE) emptyList<ClockPart>() to "" else words.countdown(start)
        val selected = pressed ?: SignupState.of(night.you.state)
        val coming = night.coming
        val (ready, stamp) = ReadyRows.of(res, night.ready, now)
        return RaidNightPageContent(
            title = night.zoneName,
            icon = RaidIcons.of(night.zoneKey, night.zoneIcon),
            facts = listOf(night.guildName, time),
            eyebrow = res.getString(if (night.phase == LIVE) R.string.home_raid_live else R.string.home_raid_tonight, night.guildName, night.zoneName),
            clock = parts,
            clockLabel = label,
            live = night.phase == LIVE || java.time.Duration.between(now, start) <= STARTING,
            aside = res.getString(R.string.raid_night_aside, time, res.getString(R.string.home_raid_counts, coming.inCount, coming.late)),
            tag = selected?.let { res.getString(R.string.raid_night_youAre, res.getString(it.label)) } ?: res.getString(R.string.raid_chip_notSigned),
            answered = selected != null,
            selected = selected,
            lateOpen = lateOpen && selected == SignupState.Late,
            roles = listOf(
                role(res.getString(R.string.raid_group_tanks), coming.tanks, target?.tanks, short = true),
                role(res.getString(R.string.raid_group_healers), coming.healers, target?.healers, short = true),
                role(res.getString(R.string.raid_group_damage), coming.damage, target?.let { it.melee + it.ranged }, short = false),
            ),
            ready = ready,
            readyStamp = stamp,
        )
    }

    /** `Healers 5 of 8`, `Mirethil, Selka, Hollowmend, Aenwyn, Brightwen · 3 short`; past four names, `and 8 more`. */
    private fun role(label: String, people: List<RaidNightPerson>, want: Int?, short: Boolean): Pair<String, String> {
        val names = people.map { it.name }
        var line = if (names.size > NAMED) res.getString(R.string.raid_night_more, names.take(NAMED_WHEN_LONG).joinToString(", "), names.size - NAMED_WHEN_LONG) else names.joinToString(", ")
        if (short && want != null && want > people.size) line = res.getString(R.string.raid_night_short, line, want - people.size)
        val title = want?.let { res.getString(R.string.raid_tab_roleLine, label, people.size, it) } ?: label
        return title to line
    }

    private companion object {
        const val LIVE = "live"
        const val NAMED = 5
        const val NAMED_WHEN_LONG = 3
        val STARTING: java.time.Duration = java.time.Duration.ofHours(1)
    }
}
