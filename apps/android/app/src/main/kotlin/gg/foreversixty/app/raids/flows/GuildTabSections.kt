package gg.foreversixty.app.raids.flows

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.designsystem.Codes
import gg.foreversixty.app.model.GuildHome
import gg.foreversixty.app.model.GuildMember
import gg.foreversixty.app.model.GuildOfficerApplication
import gg.foreversixty.app.model.GuildReportRow
import gg.foreversixty.app.model.GuildSetup
import gg.foreversixty.app.model.GuildSummary
import gg.foreversixty.app.model.LootNight
import gg.foreversixty.app.model.LootNights
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
    /** The guild's level spread: at 60, 55 to 59, and your level under 60 or the raiders at 60. */
    val spread: GuildSpread? = null,
    /** Under 60 the guild leads the tab (spec 2.1); at 60 the raid does (2.2). */
    val leveler: Boolean = false,
    val rosterMeta: String? = null,
    val recent: List<GuildMemberLine> = emptyList(),
    val close: List<GuildMemberLine> = emptyList(),
    val logs: List<GuildLogLine> = emptyList(),
    val progression: GuildProgressLine? = null,
    /** At 60: the newest night's drops and who won them; [loot] is the badge's count when the night is not loaded. */
    val lootNight: GuildLootNight? = null,
    val loot: String? = null,
    val officer: List<OfficerShortcut> = emptyList(),
    /** Officers: the raiders who missed most of the last 30 days' nights. */
    val missing: List<GuildMemberLine> = emptyList(),
    /** Officers: the raiders whose gear is furthest from their list (the addon's gear gap, DPS). */
    val gearGaps: List<GuildMemberLine> = emptyList(),
    /** Your own attendance over the last 30 days, `7 of 8 nights`. */
    val yourAttendance: GuildMemberLine? = null,
)

/** The newest loot night (spec 2.2): its raid and day, each drop with who won it (null while open). */
data class GuildLootNight(val zone: String, val day: String, val drops: List<Drop>) {
    /** [yours]: one of your characters won it. */
    data class Drop(
        val itemId: Int,
        val name: String,
        val icon: String,
        val quality: Int,
        val winner: String?,
        val winnerClass: String?,
        val yours: Boolean = false,
    )
}

data class GuildTabHead(val name: String, val faction: String?, val crestUrl: String?, val line: String)

/** `31 at 60 · 3 55 to 59 · 36 your level` under 60, or `24 raiders` at 60; [own] draws the third figure in gold. */
data class GuildSpread(val atCap: Int, val close: Int, val third: Int, val thirdLabel: String, val own: Boolean)

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
    val lootNights: LootNights? = null,
    /** Every character on the account, for your own loot. */
    val yourKeys: Set<String> = emptySet(),
)

/** Turns the guild home into the Guild tab's sections. Pure over [Resources], so each state is a test of values. */
class GuildTabPresenter(private val res: Resources, private val clock: RaidClock, private val now: Instant, private val zone: ZoneId = ZoneId.systemDefault()) {
    fun sections(inputs: GuildTabInputs, guildName: String): GuildTabSections? {
        val home = inputs.home ?: return null
        val leveler = (inputs.level ?: MAX_LEVEL) < MAX_LEVEL
        val levels = home.members.mapNotNull { it.level }
        val atCap = levels.count { it >= MAX_LEVEL }
        val others = home.members.filter { it.characterKey == null || it.characterKey != inputs.characterKey }
        val close = levels.count { it in CLOSE_FROM until MAX_LEVEL }
        val recent = recent(others, home.roster.mapNotNull { row -> row.itemLevel?.let { row.characterKey to it } }.toMap())
        val shown = recent.map { it.name }.toSet()
        val officerView = home.viewer.officerTools || inputs.badge?.officer == true
        val lootNight = if (leveler) null else inputs.lootNights?.nights?.firstOrNull()?.let { lootNight(it, inputs.yourKeys) }
        return GuildTabSections(
            head = head(home, guildName),
            spread = if (leveler) {
                GuildSpread(atCap, close, inputs.level ?: 0, res.getString(R.string.guild_tab_you), own = true)
            } else {
                GuildSpread(atCap, close, home.summary.raiders, res.getString(R.string.guild_tab_raiders), own = false)
            },
            leveler = leveler,
            rosterMeta = res.getString(R.string.guild_tab_rosterMeta, atCap, home.members.size, MAX_LEVEL),
            recent = recent,
            close = if (leveler) close(others).filterNot { it.name in shown } else emptyList(),
            logs = home.reports.take(LOGS_SHOWN).map(::log),
            progression = progression(home.summary, inputs.firstRaids),
            lootNight = lootNight,
            loot = if (leveler || lootNight != null) null else inputs.badge?.lastRaidLoot?.let { res.getQuantityString(R.plurals.raid_pages_lootLast, it.count, it.count) },
            officer = if (officerView) officer(inputs) else emptyList(),
            missing = if (officerView) missing(home) else emptyList(),
            gearGaps = if (officerView) gearGaps(home) else emptyList(),
            yourAttendance = yourAttendance(home, inputs.characterKey),
        )
    }

    private fun head(home: GuildHome, name: String): GuildTabHead {
        val members = res.getQuantityString(R.plurals.raid_pages_members, home.counts.onRoster, home.counts.onRoster)
        val faction = home.guild.faction?.replaceFirstChar { it.titlecase(Locale.US) }
        val line = faction?.let { res.getString(R.string.guild_tab_headLine, members, it, realm(home.guild.region, home.guild.ruleset)) } ?: members
        return GuildTabHead(home.guild.name.ifEmpty { name }, home.guild.faction, home.guild.crestUrl, line)
    }

    /** The newest night's gear drops, each with who won it. */
    private fun lootNight(night: LootNight, yours: Set<String>): GuildLootNight? {
        val drops = night.drops.filter { it.kind == GEAR }.take(LOOT_SHOWN).map {
            GuildLootNight.Drop(it.itemId, it.name, it.icon, it.quality, it.award?.name, it.award?.className, it.award?.characterKey in yours)
        }
        if (drops.isEmpty()) return null
        val day = runCatching { clock.shortDate(java.time.LocalDate.parse(night.date).atTime(12, 0).toInstant(java.time.ZoneOffset.UTC)) }.getOrDefault(night.date)
        return GuildLootNight(night.zone, day, drops)
    }

    /** Officers: the three raiders furthest from their list (the addon's gear gap), largest first. */
    fun gearGaps(home: GuildHome): List<GuildMemberLine> = home.roster
        .mapNotNull { row -> row.addon?.gearGap?.takeIf { it > 0 }?.let { row to it } }
        .sortedByDescending { it.second }
        .take(GEAR_GAPS_SHOWN)
        .map { (row, gap) ->
            GuildMemberLine(
                row.name, row.className, row.specIcon,
                listOfNotNull(row.specName, row.itemLevel?.let { res.getString(R.string.guild_tab_itemLevel, it) }).joinToString(SEPARATOR),
                res.getString(R.string.guild_tab_gearGap, gap),
            )
        }

    /** Your own character's attendance over the last 30 days, when its roster row carries one. */
    fun yourAttendance(home: GuildHome, key: String?): GuildMemberLine? {
        val row = home.roster.firstOrNull { it.characterKey == key } ?: return null
        val present = row.attendance30?.present ?: return null
        val nights = row.attendance30?.nights?.takeIf { it > 0 } ?: return null
        return GuildMemberLine(row.name, row.className, row.specIcon, row.specName.orEmpty(), res.getString(R.string.guild_tab_nights, present, nights))
    }

    /** Officers: the raiders present at under half of at least four nights in the last 30 days, fewest first. */
    fun missing(home: GuildHome): List<GuildMemberLine> = home.roster
        .mapNotNull { row ->
            val present = row.attendance30?.present ?: return@mapNotNull null
            val nights = row.attendance30?.nights ?: return@mapNotNull null
            if (nights < MISSING_NIGHTS || present * 2 >= nights) null else Triple(row, present, nights)
        }
        .sortedBy { (_, present, nights) -> present.toDouble() / nights }
        .take(MISSING_SHOWN)
        .map { (row, present, nights) ->
            GuildMemberLine(
                row.name, row.className, row.specIcon,
                listOfNotNull(row.specName, row.itemLevel?.let { res.getString(R.string.guild_tab_itemLevel, it) }).joinToString(SEPARATOR),
                res.getString(R.string.guild_tab_nights, present, nights),
            )
        }

    /** The five most recently seen members (the addon's `seen_at`, read in the viewer's zone as the site reads it). */
    fun recent(members: List<GuildMember>, itemLevels: Map<String, Int> = emptyMap()): List<GuildMemberLine> =
        members.mapNotNull { m -> seenAt(m.seenAt, zone)?.let { m to it } }
            .sortedByDescending { it.second }
            .take(RECENT_SHOWN)
            .map { (m, seen) ->
                val itemLevel = m.characterKey?.let(itemLevels::get)?.let { res.getString(R.string.guild_tab_itemLevel, it) }
                GuildMemberLine(
                    m.name, m.className, m.specIcon,
                    listOfNotNull(m.level?.toString(), m.specName, itemLevel).joinToString(SEPARATOR),
                    RelativeAge.of(seen, now).phrase(res),
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
        val raids = tier?.raids.orEmpty()
        val title = if (raids.isEmpty()) tier?.name else raids.joinToString(", ")
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
            val classWord = it.character.className?.replaceFirstChar { c -> c.titlecase(Locale.US) }
            val spec = it.character.specName
            // Some rows name the class in the spec already (`Holy Paladin`): never `Holy Paladin Paladin`.
            val kind = if (spec != null && classWord != null && spec.endsWith(classWord, ignoreCase = true)) spec else listOfNotNull(spec, classWord).joinToString(" ")
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
        const val LOOT_SHOWN = 4
        const val MISSING_SHOWN = 3
        const val GEAR_GAPS_SHOWN = 3
        /** Missing raids: at least this many of the last 30 days' nights, and present at under half. */
        const val MISSING_NIGHTS = 4
        private const val GEAR = "gear"

        /** `US PvP`, the site's realm words. */
        fun realm(region: String, ruleset: String): String {
            val rules = if (ruleset.equals("pvp", ignoreCase = true)) "PvP" else ruleset.replaceFirstChar { it.titlecase(Locale.US) }
            return "${Codes.region(region)} $rules"
        }
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
