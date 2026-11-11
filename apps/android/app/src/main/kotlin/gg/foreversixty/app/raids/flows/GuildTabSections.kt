package gg.foreversixty.app.raids.flows

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.model.GuildHome
import gg.foreversixty.app.model.GuildMember
import gg.foreversixty.app.model.GuildOfficerApplication
import gg.foreversixty.app.model.GuildReportRow
import gg.foreversixty.app.model.GuildSetup
import gg.foreversixty.app.model.GuildSummary
import gg.foreversixty.app.model.RaidBadge
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.phrase
import gg.foreversixty.app.raids.logic.RaidClock
import java.time.Instant
import java.time.LocalDateTime
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.Locale

/**
 * The Guild tab's own sections (spec 2026-11-10-guild-tab.md section 2), around the raid the tab
 * already draws: the guild head and level spread, the next raid for a player who cannot answer it,
 * the roster's highlights, the recent logs, progression, the loot line and the officer's shortcuts.
 */
data class GuildTabSections(
    val head: GuildTabHead,
    /** Below 60: the level spread (the head leads the tab); null at 60. */
    val spread: GuildSpread? = null,
    val rosterMeta: String? = null,
    val recent: List<GuildMemberLine> = emptyList(),
    val close: List<GuildMemberLine> = emptyList(),
    val logs: List<GuildLogLine> = emptyList(),
    val progression: GuildProgressLine? = null,
    val loot: String? = null,
    val officer: List<OfficerShortcut> = emptyList(),
)

data class GuildTabHead(val name: String, val faction: String?, val crestUrl: String?, val line: String)

data class GuildSpread(val atCap: Int, val close: Int, val you: Int)

/** A roster highlight: crest, name in class colour, `60 · Protection Warrior`, and `Seen 12 min ago` or the level. */
data class GuildMemberLine(val name: String, val className: String?, val specIcon: String?, val line: String, val trail: String)

data class GuildLogLine(val id: String, val title: String, val line: String, val day: String, val clean: Boolean)

data class GuildProgressLine(val title: String, val line: String)

enum class OfficerKind { Plan, Applications, Setup }

/** One officer shortcut (spec 2.3), one tap from the tab. */
data class OfficerShortcut(val kind: OfficerKind, val title: String, val line: String?, val count: String?)

/** What the Guild tab reads for its own sections; each part is optional and a failed call leaves it out. */
data class GuildTabInputs(
    val home: GuildHome? = null,
    val setup: GuildSetup? = null,
    /** The pending applications, newest first. */
    val applicants: List<GuildOfficerApplication> = emptyList(),
    val badge: RaidBadge? = null,
    val level: Int? = null,
    val characterKey: String? = null,
    val planned: List<RaidView> = emptyList(),
    val firstRaids: Instant? = null,
)

/** Turns the guild home into the Guild tab's sections. Pure over [Resources], so each state is a test of values. */
class GuildTabPresenter(private val res: Resources, private val clock: RaidClock, private val now: Instant, private val zone: ZoneId = ZoneId.systemDefault()) {
    fun sections(inputs: GuildTabInputs, guildName: String): GuildTabSections? {
        val home = inputs.home ?: return null
        val leveler = (inputs.level ?: MAX_LEVEL) < MAX_LEVEL
        val levels = home.members.mapNotNull { it.level }
        val atCap = levels.count { it >= MAX_LEVEL }
        val others = home.members.filter { it.characterKey == null || it.characterKey != inputs.characterKey }
        return GuildTabSections(
            head = head(home, guildName),
            spread = if (leveler) GuildSpread(atCap, levels.count { it in CLOSE_FROM until MAX_LEVEL }, inputs.level ?: 0) else null,
            rosterMeta = res.getString(R.string.guild_tab_rosterMeta, atCap, home.members.size, MAX_LEVEL),
            recent = recent(others),
            close = if (leveler) close(others) else emptyList(),
            logs = home.reports.take(LOGS_SHOWN).map(::log),
            progression = progression(home.summary, inputs.firstRaids),
            loot = if (leveler) null else inputs.badge?.lastRaidLoot?.let { res.getQuantityString(R.plurals.raid_pages_lootLast, it.count, it.count) },
            officer = if (home.viewer.officerTools || inputs.badge?.officer == true) officer(inputs) else emptyList(),
        )
    }

    private fun head(home: GuildHome, name: String): GuildTabHead {
        val members = res.getQuantityString(R.plurals.raid_pages_members, home.counts.onRoster, home.counts.onRoster)
        val faction = home.guild.faction?.replaceFirstChar { it.titlecase(Locale.US) }
        val line = faction?.let { res.getString(R.string.guild_tab_headLine, members, it) } ?: members
        return GuildTabHead(home.guild.name.ifEmpty { name }, home.guild.faction, home.guild.crestUrl, line)
    }

    /** The five most recently seen members (the addon's `seen_at`, read in the viewer's zone as the site reads it). */
    fun recent(members: List<GuildMember>): List<GuildMemberLine> =
        members.mapNotNull { m -> seenAt(m.seenAt, zone)?.let { m to it } }
            .sortedByDescending { it.second }
            .take(RECENT_SHOWN)
            .map { (m, seen) ->
                GuildMemberLine(
                    m.name, m.className, m.specIcon,
                    listOfNotNull(m.level?.toString(), m.specName).joinToString(SEPARATOR),
                    res.getString(R.string.guild_tab_seen, RelativeAge.of(seen, now).phrase(res)),
                )
            }

    fun close(members: List<GuildMember>): List<GuildMemberLine> =
        members.filter { (it.level ?: 0) in CLOSE_FROM until MAX_LEVEL }
            .sortedWith(compareByDescending<GuildMember> { it.level ?: 0 }.thenBy { it.name })
            .take(CLOSE_SHOWN)
            .map { GuildMemberLine(it.name, it.className, it.specIcon, it.specName.orEmpty(), (it.level ?: 0).toString()) }

    private fun log(report: GuildReportRow): GuildLogLine = GuildLogLine(
        id = report.id,
        title = report.title,
        line = res.getString(R.string.guild_tab_logLine, report.zone, report.killCount, report.fightCount, duration(report.durationMs)),
        day = clock.shortDate(Instant.parse(report.createdAt)),
        clean = report.fightCount > 0 && report.killCount == report.fightCount,
    )

    /**
     * `1 of 1 down · 6 pulls this tier`; before the first pull, when the first raids open, and with no pull and no
     * date, nothing (an empty tier says nothing yet).
     */
    fun progression(summary: GuildSummary, firstRaids: Instant?): GuildProgressLine? {
        val tier = summary.tier
        val named = tier?.namedEncounters ?: 0
        val title = tier?.name ?: tier?.raids?.firstOrNull()
        if (title == null || named == 0) return null
        val pulls = tier?.pullsTotal ?: summary.pullsThisTier
        val down = tier?.down ?: summary.namedEncountersDown
        if (pulls == 0) {
            val opens = firstRaids?.takeIf { it.isAfter(now) } ?: return null
            return GuildProgressLine(title, res.getString(R.string.guild_tab_firstRaids, clock.shortDate(opens)))
        }
        val line = listOf(
            res.getString(R.string.guild_tab_down, down, named),
            res.getQuantityString(R.plurals.guild_tab_pulls, pulls, pulls),
        ).joinToString(SEPARATOR)
        return GuildProgressLine(title, line)
    }

    private fun officer(inputs: GuildTabInputs): List<OfficerShortcut> {
        val upcoming = inputs.planned.map { it to Instant.parse(it.startsAt) }.filter { it.second.isAfter(now) }.sortedBy { it.second }.take(PLANNED_SHOWN)
        val planned = if (upcoming.isEmpty()) {
            res.getString(R.string.guild_tab_nonePlanned)
        } else {
            upcoming.joinToString(", ") { (raid, at) -> res.getString(R.string.raid_pages_calendarItem, raid.title, clock.weekday(at).take(DAY_LETTERS)) }
        }
        val waiting = maxOf(inputs.applicants.size, inputs.badge?.applicants ?: 0)
        val newest = inputs.applicants.firstOrNull()?.let {
            val kind = listOfNotNull(it.character.specName, it.character.className?.replaceFirstChar { c -> c.titlecase(Locale.US) })
                .distinct().joinToString(" ")
            res.getString(R.string.guild_tab_applicant, it.character.name, kind, it.character.level ?: 0)
        }
        val line = newest?.let { if (waiting > 1) res.getString(R.string.guild_tab_applicantMore, it, waiting - 1) else it }
        val rows = mutableListOf(
            OfficerShortcut(OfficerKind.Plan, res.getString(R.string.guild_tab_planRaid), planned, null),
            OfficerShortcut(
                OfficerKind.Applications, res.getString(R.string.guild_tab_applications),
                line ?: res.getString(R.string.guild_tab_noApplications),
                if (waiting > 0) res.getString(R.string.guild_tab_waiting, waiting) else null,
            ),
        )
        inputs.setup?.takeIf { !it.finished }?.let { setup ->
            val todo = setup.steps.filter { it.state == SETUP_TODO }.mapNotNull { stepWord(it.key) }
            rows += OfficerShortcut(
                OfficerKind.Setup, res.getString(R.string.guild_tab_setup),
                todo.takeIf { it.isNotEmpty() }?.let { res.getString(R.string.guild_tab_setupToGo, list(it)).replaceFirstChar { c -> c.titlecase(Locale.US) } },
                res.getString(R.string.guild_tab_setupDone, setup.done, setup.total),
            )
        }
        return rows
    }

    /** `1 h 32 min`, `41 min`. */
    fun duration(ms: Long): String {
        val minutes = (ms / MS_PER_MINUTE).toInt()
        val hours = minutes / MINUTES_PER_HOUR
        return if (hours > 0) {
            res.getString(R.string.raid_tab_untilHours, hours, minutes % MINUTES_PER_HOUR)
        } else {
            res.getString(R.string.raid_tab_untilMinutes, minutes)
        }
    }

    private fun list(words: List<String>): String =
        if (words.size < 2) words.firstOrNull().orEmpty() else res.getString(R.string.raid_list_and, words.dropLast(1).joinToString(", "), words.last())

    private fun stepWord(key: String): String? = STEP_WORDS[key]?.let(res::getString)

    companion object {
        /**
         * The pages the Guild tab's sections do not already open (spec 2.1, 2.2): Roster has its section, Loot history
         * too at 60, Guild sims is a 60's page.
         */
        fun pagesLeft(pages: List<GuildPageRow>, leveler: Boolean): List<GuildPageRow> = pages.filter {
            when (it.page) {
                GuildPage.Roster -> false
                GuildPage.Loot -> leveler
                GuildPage.Sims -> !leveler
                GuildPage.Calendar -> true
            }
        }

        const val MAX_LEVEL = 60
        /** Close to 60 (spec 2.1): 55 to 59. */
        const val CLOSE_FROM = 55
        const val RECENT_SHOWN = 5
        const val CLOSE_SHOWN = 5
        const val LOGS_SHOWN = 3
        /** Plan a raid names the next two planned raids, as the Calendar row does. */
        const val PLANNED_SHOWN = 2
        const val SETUP_TODO = "todo"
        private const val SEPARATOR = " · "
        private const val DAY_LETTERS = 3
        private const val MS_PER_MINUTE = 60_000L
        private const val MINUTES_PER_HOUR = 60
        private val SEEN_FORMAT: DateTimeFormatter = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm", Locale.US)

        /** The setup step's name, as the site's checklist titles it. */
        private val STEP_WORDS = mapOf(
            "officer_ranks" to R.string.guild_tab_stepRanks,
            "discord_bot" to R.string.guild_tab_stepDiscord,
            "recruiting" to R.string.guild_tab_stepRecruiting,
            "raid_schedule" to R.string.guild_tab_stepSchedule,
            "invite_link" to R.string.guild_tab_stepInvite,
            "addon_sync" to R.string.guild_tab_stepAddon,
            "report_visibility" to R.string.guild_tab_stepVisibility,
            "linked_role" to R.string.guild_tab_stepRole,
        )

        /** The member's `seen_at`, their own local `YYYY-MM-DD HH:MM`, read in the viewer's zone (web's roster-addon.ts). */
        fun seenAt(text: String?, zone: ZoneId): Instant? = text?.let {
            runCatching { LocalDateTime.parse(it, SEEN_FORMAT).atZone(zone).toInstant() }.getOrNull()
        }
    }
}
