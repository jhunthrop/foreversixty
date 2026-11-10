package gg.foreversixty.app.raids.flows

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.GuildRosterRow
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.raids.logic.SignupState

/**
 * The lines that close Who is coming on the Raid tab (round 2, guild-officer G1 and G4, info-density):
 * who else answered, what the comp misses, and for officers who signed In without syncing this week.
 */
class RaidTabFacts(private val res: Resources) {
    fun foot(raid: RaidView, officer: Boolean, roster: List<GuildRosterRow>): List<RoleFoot> = listOfNotNull(
        counts(raid)?.let(::RoleFoot),
        covered(raid)?.let { RoleFoot(it, warn = true) },
        notSynced(raid, roster).takeIf { officer && it > 0 }?.let {
            RoleFoot(res.getQuantityString(R.plurals.raid_tab_notSynced, it, it), warn = true)
        },
    )

    /** `Bench 1 · Out 4 · 14 not answered`; the parts with nobody are left out. */
    fun counts(raid: RaidView): String? = listOfNotNull(
        raid.counts.bench?.takeIf { it > 0 }?.let { res.getString(R.string.raid_tab_countBench, it) },
        raid.counts.signedOut?.takeIf { it > 0 }?.let { res.getString(R.string.raid_tab_countOut, it) },
        raid.notSigned.size.takeIf { it > 0 }?.let { res.getString(R.string.raid_tab_countNotAnswered, it) },
    ).joinToString(SEPARATOR).ifEmpty { null }

    /**
     * The Covered line (G4), one line a raid leader reads at a glance: the gaps nobody brings first, then
     * the short ones, at most [COVER_NAMES] names and `and N more` (`Missing Fortitude · Windfury Totem 1 of 2`,
     * `Missing Blessing of Kings, Blessing of Might, Judgement of Wisdom and 9 more`); null when all is covered.
     */
    fun covered(raid: RaidView): String? {
        val gaps = gaps(raid).sortedBy { it.have != null }
        if (gaps.isEmpty()) return null
        val shown = gaps.take(COVER_NAMES)
        val missing = shown.filter { it.have == null }.takeIf { it.isNotEmpty() }
            ?.let { list -> res.getString(R.string.raid_tab_coverMissing, list.joinToString(LIST) { it.name }) }
        val short = shown.mapNotNull { gap -> gap.have?.let { res.getString(R.string.raid_tab_coverShort, gap.name, it, gap.need) } }
        val line = listOfNotNull(missing, short.joinToString(LIST).ifEmpty { null }).joinToString(SEPARATOR)
        val hidden = gaps.size - shown.size
        return if (hidden > 0) res.getString(R.string.raid_tab_andMore, line, hidden) else line
    }

    /**
     * The cover rows that are short or missing (In and Late against need), in the cover's order; a raid
     * without cover rows falls back to the buffs nobody brings.
     */
    private fun gaps(raid: RaidView): List<CoverGap> {
        val rows = raid.cover.orEmpty()
        if (rows.isEmpty()) return raid.composition?.buffs.orEmpty().filter { !it.present }.map { CoverGap(it.name) }
        return rows.mapNotNull { row ->
            when (row.state) {
                MISSING -> CoverGap(row.name)
                SHORT -> CoverGap(row.name, row.providersIn + row.providersLate, row.need)
                else -> null
            }
        }
    }

    /** A gap in the cover: nobody brings [name] ([have] null), or [have] of the [need] the raid wants. */
    private data class CoverGap(val name: String, val have: Int? = null, val need: Int = 0)

    /** G1: the In and Late sign-ups whose roster row has not logged in this week. */
    fun notSynced(raid: RaidView, roster: List<GuildRosterRow>): Int {
        val stale = roster.filter { !it.loggedRecently }.map { it.characterKey }.toSet()
        return raid.signups.count { SignupState.of(it.state)?.isComing == true && it.characterKey in stale }
    }

    private companion object {
        const val SEPARATOR = " · "
        const val LIST = ", "

        /** The most names the Covered line spells out before `and N more`. */
        const val COVER_NAMES = 3
        const val MISSING = "missing"
        const val SHORT = "short"
    }
}
