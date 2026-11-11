package gg.foreversixty.app.raids.flows

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.home.flows.RaidCardWords
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.model.GuildOfficerApplication
import gg.foreversixty.app.model.RaidBadge
import gg.foreversixty.app.model.RaidNightRow
import gg.foreversixty.app.model.GuildHome
import gg.foreversixty.app.model.ReadyCheck
import gg.foreversixty.app.phrase
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.raids.logic.ComingCounts
import gg.foreversixty.app.raids.logic.PlannedRaids
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidsTabContent
import gg.foreversixty.app.raids.logic.RaidsTabState
import gg.foreversixty.app.raids.logic.SignupController
import gg.foreversixty.app.raids.logic.ShortRole
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.model.RaidSignup
import gg.foreversixty.app.flows.TallyItem
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
    /** The guild's home (roster and counts), for an officer's readiness line and the Roster row's figure. */
    val home: GuildHome? = null,
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
        val next = PlannedRaids(content.planned, now).next?.takeIf { live == null }?.let { raid ->
            words.card(raid, content.head.name, key, controller, badge)
        }
        val applicantLine = extras.applicant?.let { ApplicantPresenter(res, clock, now).line(it) }
        val (ready, stamp) = ReadyRows.of(res, extras.ready, now)
        val shown = live ?: PlannedRaids(content.planned, now).next
        val officer = badge?.officer == true
        return RaidTabView.Guild(
            guild = content.head.name,
            next = next,
            needs = next?.officer?.copy(applicantLine = applicantLine),
            applicantId = extras.applicant?.id,
            needsCount = badge?.count ?: 0,
            live = live?.let { words.live(it, content.head.name, key, words.ownAnswer(it, key)) },
            roles = shown?.let { RaidRoleLines(res).of(it) }.orEmpty(),
            foot = shown?.let { RaidTabFacts(res).foot(it, officer, extras.home?.roster.orEmpty(), now) }.orEmpty(),
            notes = next?.let { shown?.notesHtml?.takeIf { html -> html.isNotBlank() } },
            needsFirst = officer,
            ready = ready,
            readyStamp = stamp,
            pages = pages(content, hasExtras, extras.home, badge),
        )
    }

    /** The guild's raid pages: Roster and Loot history where the app has them, Calendar from the planned raids. */
    private fun pages(content: RaidsTabContent, hasExtras: Boolean, home: GuildHome?, badge: RaidBadge?): List<GuildPageRow> {
        val members = home?.counts?.onRoster?.let { res.getQuantityString(R.plurals.raids_rosterRow_sub, it, it) }
        val calendar = content.planned.take(CALENDAR_ITEMS).joinToString(", ") {
            res.getString(R.string.raid_tab_calendarItem, it.zoneName.substringBefore("'"), clock.shortDate(Instant.parse(it.startsAt)).substringBefore(" "))
        }.ifEmpty { null }
        return listOfNotNull(
            GuildPageRow(GuildPage.Roster, res.getString(R.string.raid_tab_roster), members).takeIf { hasExtras },
            GuildPageRow(GuildPage.Calendar, res.getString(R.string.raid_tab_calendar), calendar).takeIf { content.head.webUrl != null },
            GuildPageRow(GuildPage.Loot, res.getString(R.string.raid_tab_loot), lootLine(content.nights.firstOrNull(), badge)).takeIf { hasExtras },
            GuildPageRow(GuildPage.Sims, res.getString(R.string.raid_tab_sims), null).takeIf { content.head.webUrl != null },
        )
    }

    /** Loot history's line: `Last: Onyxia's Lair, 1 kill, 2 wipes`, then the last raid's pieces; null with neither. */
    fun lootLine(lastNight: RaidNightRow?, badge: RaidBadge?): String? {
        val night = lastNight?.let {
            res.getString(
                R.string.raid_pages_lastNight,
                it.zone,
                res.getQuantityString(R.plurals.raid_pages_kills, it.killCount, it.killCount),
                res.getQuantityString(R.plurals.raid_pages_wipes, it.wipeCount, it.wipeCount),
            )
        }
        val pieces = badge?.lastRaidLoot?.let { res.getQuantityString(R.plurals.raid_pages_lootLast, it.count, it.count) }
        return listOfNotNull(night, pieces).joinToString(LINE_SEPARATOR).ifEmpty { null }
    }

    private companion object {
        const val TANK = "tank"
        const val HEALER = "healer"
        const val CALENDAR_ITEMS = 2
        const val HORDE = "horde"
        const val LINE_SEPARATOR = " · "
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
            ReadyItem(tone, it.title, it.line, it.whereToGet?.label, it.whereToGet?.link, it.icon)
        }
        return items to res.getString(R.string.raid_tab_stamp, RelativeAge.of(Instant.parse(exported), now).phrase(res))
    }
}

/**
 * Who is coming (spec 4.2): tanks, healers and damage with the names of those In or Late (yours
 * first, a Late with its minutes, each in its class colour), the role's tally against the comp
 * (S6: `Healers 5 of 8 · 3 short`), and on Damage the melee and ranged split.
 */
class RaidRoleLines(private val res: Resources) {
    fun of(raid: RaidView): List<RoleLine> {
        val target = raid.compTarget
        val counts = ComingCounts.of(raid)
        val coming = raid.signups.filter { SignupState.of(it.state)?.isComing == true }.sortedByDescending { it.own }
        return listOf(
            Group(ShortRole.TANK, R.string.raid_group_tanks, counts.tanks, target?.tanks),
            Group(ShortRole.HEALER, R.string.raid_group_healers, counts.healers, target?.healers),
            Group(ShortRole.DPS, R.string.raid_group_damage, counts.damage, target?.let { it.melee + it.ranged }),
        ).mapNotNull { group ->
            val people = coming.filter { group(it.role) == group.wire }
            if (group.have == 0 && group.want == null) return@mapNotNull null
            val tally = TallyItem(res.getString(group.label), group.have, group.want)
            val split = if (group.wire == ShortRole.DPS) split(raid) else null
            val hidden = people.size - shown(people.size)
            // `and 8 more` after the names (the names themselves are drawn in class colours).
            val more = res.getString(R.string.raid_tab_andMore, "", hidden).takeIf { hidden > 0 }
            RoleLine(tally, people.take(shown(people.size)).map(::name), more, split)
        }
    }

    private fun shown(count: Int) = if (count <= SHOWN_ALL) count else SHOWN_FEW

    private fun name(s: RaidSignup): RoleName {
        val minutes = s.lateMinutes?.takeIf { s.state == SignupState.Late.wire }
        val word = minutes?.let { res.getString(R.string.raid_tab_lateNote, s.name, it) } ?: s.name
        return RoleName(word, s.className)
    }

    /** `5 melee, 6 ranged`, from the same coming counts as the tally (row 1); null when neither has anyone. */
    private fun split(raid: RaidView): String? {
        val c = ComingCounts.of(raid)
        if (c.damage == 0) return null
        return res.getString(R.string.raid_tab_damageSplit, c.melee, c.ranged)
    }

    private fun group(role: String) = ShortRole.group(role)

    private data class Group(val wire: String, @param:StringRes val label: Int, val have: Int, val want: Int?)

    private companion object {
        const val SHOWN_ALL = 5
        const val SHOWN_FEW = 3
    }
}
