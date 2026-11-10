package gg.foreversixty.app.raids.flows

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.home.flows.RaidCardWords
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.model.GuildOfficerApplication
import gg.foreversixty.app.model.RaidBadge
import gg.foreversixty.app.model.ReadyCheck
import gg.foreversixty.app.phrase
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.raids.logic.PlannedRaids
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidsTabContent
import gg.foreversixty.app.raids.logic.RaidsTabState
import gg.foreversixty.app.raids.logic.SignupController
import gg.foreversixty.app.raids.logic.ShortRole
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.model.RaidSignup
import androidx.annotation.StringRes
import gg.foreversixty.app.home.flows.HomeFlowState
import java.time.Instant

/**
 * What the Raid tab reads beside the raids (spec 2026-11-08 4.2): the newest pending applicant (an
 * officer's), the guilds recruiting the spec (no guild), the Ready check of the raid on tonight.
 */
data class RaidTabExtras(
    val applicant: GuildOfficerApplication? = null,
    val recruiting: List<gg.foreversixty.app.model.RecruitingGuild> = emptyList(),
    val ready: ReadyCheck? = null,
)

/**
 * The Raid tab's view from the tab model's state (spec 2026-11-08 4.2). Null for the states the tab
 * still draws with its older views (loading, failed, waiting for approval).
 */
class RaidTabPresenter(private val res: Resources, private val clock: RaidClock, private val now: Instant) {
    fun view(state: RaidsTabState, controller: SignupController?, badge: RaidBadge?, hasExtras: Boolean, extras: RaidTabExtras = RaidTabExtras()): RaidTabView? =
        when (state) {
            RaidsTabState.SignedOut -> RaidTabView.SignedOut
            is RaidsTabState.NoGuild -> {
                val level = state.level
                if (level != null && level < BisRange.LAST_BAND) {
                    RaidTabView.Leveler(BisRange.LAST_BAND - level)
                } else {
                    RaidTabView.NoGuild(state.spec, extras.recruiting.map(::recruiting))
                }
            }
            is RaidsTabState.Loaded -> guild(state.content, controller, badge, hasExtras, extras)
            else -> null
        }

    /** `Alliance · Wed and Sat, 8:00 pm CST · needs damage`. */
    private fun recruiting(row: gg.foreversixty.app.model.RecruitingGuild): RecruitingGuild {
        val faction = res.getString(if (row.faction == HORDE) R.string.bis_faction_horde else R.string.bis_faction_alliance)
        val needs = res.getString(
            when (row.needs) {
                TANK -> R.string.home_officer_role_tanks
                HEALER -> R.string.home_officer_role_healers
                else -> R.string.home_officer_role_damage
            },
        )
        return RecruitingGuild(row.guildId, row.name, res.getString(R.string.raid_recruit_line, faction, row.schedule, needs))
    }

    private fun guild(content: RaidsTabContent, controller: SignupController?, badge: RaidBadge?, hasExtras: Boolean, extras: RaidTabExtras): RaidTabView.Guild {
        val words = RaidCardWords(res, clock, now)
        val key = content.character.key
        val live = content.planned.firstOrNull { HomeFlowState.isLive(it, now) }
        // Late's minutes stay open under the answers while Late is the answer, the chosen one selected (F1).
        val next = PlannedRaids(content.planned, now).next?.takeIf { live == null }?.let { raid ->
            val answer = controller?.takeIf { it.raid.id == raid.id }?.block(key)?.shown ?: words.ownAnswer(raid, key)
            words.card(raid, content.head.name, key, answer, lateOpen = true, badge = badge, lateMinutes = words.ownLateMinutes(raid, key))
        }
        val applicantLine = extras.applicant?.let { ApplicantPresenter(res, clock, now).line(it) }
        val (ready, stamp) = ReadyRows.of(res, extras.ready, now)
        return RaidTabView.Guild(
            guild = content.head.name,
            next = next,
            needs = next?.officer?.copy(applicantLine = applicantLine),
            applicantId = extras.applicant?.id,
            needsCount = badge?.count ?: 0,
            live = live?.let { words.live(it, content.head.name, key, words.ownAnswer(it, key)) },
            roles = (live ?: PlannedRaids(content.planned, now).next)?.let { RaidRoleLines(res).of(it) }.orEmpty(),
            ready = ready,
            readyStamp = stamp,
            pages = pages(content, hasExtras),
        )
    }

    /** The guild's raid pages: Roster and Loot history where the app has them, Calendar from the planned raids. */
    private fun pages(content: RaidsTabContent, hasExtras: Boolean): List<GuildPageRow> {
        val calendar = content.planned.take(CALENDAR_ITEMS).joinToString(", ") {
            res.getString(R.string.raid_tab_calendarItem, it.zoneName.substringBefore("'"), clock.shortDate(Instant.parse(it.startsAt)).substringBefore(" "))
        }.ifEmpty { null }
        return listOfNotNull(
            GuildPageRow(GuildPage.Roster, res.getString(R.string.raid_tab_roster), null).takeIf { hasExtras },
            GuildPageRow(GuildPage.Calendar, res.getString(R.string.raid_tab_calendar), calendar).takeIf { content.head.webUrl != null },
            GuildPageRow(GuildPage.Loot, res.getString(R.string.raid_tab_loot), null).takeIf { hasExtras },
            GuildPageRow(GuildPage.Sims, res.getString(R.string.raid_tab_sims), null).takeIf { content.head.webUrl != null },
        )
    }

    private companion object {
        const val TANK = "tank"
        const val HEALER = "healer"
        const val CALENDAR_ITEMS = 2
        const val HORDE = "horde"
    }
}

/** The Ready check's rows (ticked, missing, a warning, where to get it) and the export's age; empty without an export. */
object ReadyRows {
    private const val OK = "ok"
    private const val MISS = "miss"

    fun of(res: Resources, check: ReadyCheck?, now: Instant): Pair<List<ReadyItem>, String?> {
        val exported = check?.exportedAt ?: return emptyList<ReadyItem>() to null
        val items = check.items.map {
            val tone = when (it.state) {
                OK -> ReadyTone.Ok
                MISS -> ReadyTone.Miss
                else -> ReadyTone.Warn
            }
            ReadyItem(tone, it.title, it.line, it.whereToGet?.label, it.whereToGet?.link)
        }
        return items to res.getString(R.string.raid_tab_stamp, RelativeAge.of(Instant.parse(exported), now).phrase(res))
    }
}

/**
 * Who is coming (spec 4.2): tanks, healers and damage with the names of those In or Late (yours
 * first, a Late with its minutes), the count against the comp, and the short role's gap.
 */
class RaidRoleLines(private val res: Resources) {
    fun of(raid: RaidView): List<RoleLine> {
        val target = raid.compTarget
        val short = ShortRole.of(raid)
        val coming = raid.signups.filter { SignupState.of(it.state)?.isComing == true }.sortedByDescending { it.own }
        return listOf(
            Group(ShortRole.TANK, R.string.raid_group_tanks, target?.tanks),
            Group(ShortRole.HEALER, R.string.raid_group_healers, target?.healers),
            Group(ShortRole.DPS, R.string.raid_group_damage, target?.let { it.melee + it.ranged }),
        ).mapNotNull { group ->
            val people = coming.filter { group(it.role) == group.wire }
            if (people.isEmpty() && group.want == null) return@mapNotNull null
            val names = names(people)
            val gap = short?.takeIf { it.wire == group.wire }?.missing
            RoleLine(res.getString(group.label), people.size, group.want, gap?.let { res.getString(R.string.raid_tab_short, names, it) } ?: names)
        }
    }

    /** Up to [SHOWN_ALL] names in full; more read `Thandrel, Kaelra, Thalrin and 8 more`. */
    private fun names(people: List<RaidSignup>): String {
        val words = people.map { s ->
            val minutes = s.lateMinutes
            if (s.state == SignupState.Late.wire && minutes != null) res.getString(R.string.raid_tab_lateNote, s.name, minutes) else s.name
        }
        if (words.size <= SHOWN_ALL) return words.joinToString(", ")
        return res.getString(R.string.raid_tab_andMore, words.take(SHOWN_FEW).joinToString(", "), words.size - SHOWN_FEW)
    }

    private fun group(role: String) = ShortRole.group(role)

    private data class Group(val wire: String, @param:StringRes val label: Int, val want: Int?)

    private companion object {
        const val SHOWN_ALL = 5
        const val SHOWN_FEW = 3
    }
}
