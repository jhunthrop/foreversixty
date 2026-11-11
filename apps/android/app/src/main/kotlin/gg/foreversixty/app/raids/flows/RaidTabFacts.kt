package gg.foreversixty.app.raids.flows

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.GuildRosterRow
import gg.foreversixty.app.model.RaidSignup
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.raids.logic.SignupState
import java.time.Duration
import java.time.Instant

/**
 * The lines that close Who is coming on the Raid tab (round 2, guild-officer G1 and G4, info-density):
 * who else answered, what the comp misses, and for officers who signed In without syncing this week.
 */
class RaidTabFacts(private val res: Resources) {
    fun foot(raid: RaidView, officer: Boolean, roster: List<GuildRosterRow>, now: Instant): List<RoleFoot> = listOfNotNull(
        counts(raid)?.let(::RoleFoot),
        coverLine(raid)?.let { RoleFoot(text(it), cover = it) },
        notSynced(raid, roster, now).takeIf { officer && it.isNotEmpty() }?.let { stale ->
            RoleFoot(notSyncedLine(stale), onlyKeys = stale.map { it.characterKey }.toSet())
        },
    )

    /** `Bench 1 · Out 4 · 14 not answered`; the parts with nobody are left out. */
    fun counts(raid: RaidView): String? = listOfNotNull(
        raid.counts.bench?.takeIf { it > 0 }?.let { res.getString(R.string.raid_tab_countBench, it) },
        raid.counts.signedOut?.takeIf { it > 0 }?.let { res.getString(R.string.raid_tab_countOut, it) },
        raid.notSigned.size.takeIf { it > 0 }?.let { res.getString(R.string.raid_tab_countNotAnswered, it) },
    ).joinToString(SEPARATOR).ifEmpty { null }

    /**
     * The Covered line (G4, round 3 row 4), one line a raid leader reads at a glance: `Missing` once, the
     * gaps nobody brings first, then the short ones, by the buff's short name (`Fort`, `MotW`) else its
     * name, [shown] names then `and N more` (`Missing Fort · WF 1 of 2`); null when all is covered.
     */
    fun covered(raid: RaidView): String? = coverLine(raid)?.let(::text)

    /** The Covered line's parts for the row that draws each name in wipe red beside its spell icon. */
    fun coverLine(raid: RaidView): CoverLine? {
        val gaps = gaps(raid).sortedBy { it.have != null }
        if (gaps.isEmpty()) return null
        val shown = gaps.take(shown(gaps.size))
        return CoverLine(shown.filter { it.have == null }, shown.filter { it.have != null }, gaps.size - shown.size)
    }

    /**
     * The line in words, also what TalkBack reads for the drawn row: every name, missing then short, on
     * the one separator (round 4 row 1, `Missing Kings · Might · WF 1 of 2 and 2 more`).
     */
    fun text(line: CoverLine): String {
        val names = names(line).joinToString(SEPARATOR) { it.text }
        val text = if (line.missing.isEmpty()) names else res.getString(R.string.raid_tab_coverMissing, names)
        return if (line.more > 0) res.getString(R.string.raid_tab_andMore, text, line.more) else text
    }

    /** The names the line spells out, in its order: the missing ones (`Kings`), then the short ones (`WF 1 of 2`). */
    fun names(line: CoverLine): List<CoverName> =
        line.missing.map { CoverName(it.name, it.icon) } + line.short.map { CoverName(shortText(it), it.icon) }

    /** `WF 1 of 2`. */
    fun shortText(item: CoverItem): String = res.getString(R.string.raid_tab_coverShort, item.name, item.have ?: 0, item.need)

    /**
     * The cover rows that are short or missing (In and Late against need), in the cover's order; a raid
     * without cover rows falls back to the buffs nobody brings.
     */
    private fun gaps(raid: RaidView): List<CoverItem> {
        val buffs = raid.composition?.buffs.orEmpty()
        val rows = raid.cover.orEmpty()
        if (rows.isEmpty()) return buffs.filter { !it.present }.map { CoverItem(it.label.ifBlank { it.name }, it.icon) }
        val labels = buffs.associate { it.key to it.label }
        return rows.mapNotNull { row ->
            val name = labels[row.key]?.takeIf { it.isNotBlank() } ?: row.name
            when (row.state) {
                MISSING -> CoverItem(name, row.icon)
                SHORT -> CoverItem(name, row.icon, row.providersIn + row.providersLate, row.need)
                else -> null
            }
        }
    }

    /**
     * G1 (round 4 s5): the In and Late sign-ups whose addon has not synced this week, `logged_at` (the
     * addon's last upload, not a game login) missing or more than [SYNC_WEEK_DAYS] days before [now].
     */
    fun notSynced(raid: RaidView, roster: List<GuildRosterRow>, now: Instant): List<RaidSignup> {
        val weekAgo = now.minus(Duration.ofDays(SYNC_WEEK_DAYS))
        val stale = roster.filter { row -> row.syncedAt()?.isBefore(weekAgo) ?: true }.map { it.characterKey }.toSet()
        return raid.signups.filter { SignupState.of(it.state)?.isComing == true && it.characterKey in stale }
    }

    /** Row 5 (GO N4): `Not synced this week: Kaelra, Thalrin, Brakka and 1 more`. */
    fun notSyncedLine(stale: List<RaidSignup>): String {
        val names = stale.take(shown(stale.size)).joinToString(LIST) { it.name }
        val hidden = stale.size - shown(stale.size)
        val list = if (hidden > 0) res.getString(R.string.raid_tab_andMore, names, hidden) else names
        return res.getString(R.string.raid_tab_notSyncedNames, list)
    }

    /**
     * How many names a line spells out of [total]: all of them up to one past [COVER_NAMES] (never
     * `and 1 more`), else [COVER_NAMES] then `and N more` with N of at least two (as iOS).
     */
    private fun shown(total: Int) = if (total <= COVER_NAMES + 1) total else COVER_NAMES

    companion object {
        /** The one separator of the tab's lines, the Covered names included. */
        const val SEPARATOR = " · "
        private const val LIST = ", "

        /** The most names the Covered line spells out before `and N more`. */
        private const val COVER_NAMES = 3
        private const val MISSING = "missing"
        private const val SHORT = "short"

        /** `Not synced this week`: the addon's last upload is older than this. */
        private const val SYNC_WEEK_DAYS = 7L

        private fun GuildRosterRow.syncedAt(): Instant? = runCatching { Instant.parse(loggedAt) }.getOrNull()
    }
}

/** A gap in the cover: nobody brings [name] ([have] null), or [have] of the [need] the raid wants. */
data class CoverItem(val name: String, val icon: String?, val have: Int? = null, val need: Int = 0)

/** A name the Covered line draws: its words (`Kings`, `WF 1 of 2`) and its spell icon, if it has one. */
data class CoverName(val text: String, val icon: String?) {
    /** A name without an icon draws none, and its text closes up (round 4 row 1). */
    val hasIcon: Boolean get() = !icon.isNullOrBlank()
}

/** The Covered line: the missing names, the short ones, and how many more are not spelled out. */
data class CoverLine(val missing: List<CoverItem>, val short: List<CoverItem>, val more: Int)
