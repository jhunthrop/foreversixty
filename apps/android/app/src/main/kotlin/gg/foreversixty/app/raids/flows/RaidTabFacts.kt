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
     * `Missing Blessings · Windfury Totem 1 of 2` (G4): the cover rows that are short or missing (In and
     * Late against need), then the buffs nobody brings; null when all is covered.
     */
    fun covered(raid: RaidView): String? {
        val rows = raid.cover.orEmpty().mapNotNull { row ->
            val have = row.providersIn + row.providersLate
            when (row.state) {
                MISSING -> res.getString(R.string.raid_tab_coverMissing, row.name)
                SHORT -> res.getString(R.string.raid_tab_coverShort, row.name, have, row.need)
                else -> null
            }
        }
        val buffs = raid.composition?.buffs.orEmpty().filter { !it.present }.map { res.getString(R.string.raid_tab_coverMissing, it.name) }
        return (rows + buffs).distinct().joinToString(SEPARATOR).ifEmpty { null }
    }

    /** G1: the In and Late sign-ups whose roster row has not logged in this week. */
    fun notSynced(raid: RaidView, roster: List<GuildRosterRow>): Int {
        val stale = roster.filter { !it.loggedRecently }.map { it.characterKey }.toSet()
        return raid.signups.count { SignupState.of(it.state)?.isComing == true && it.characterKey in stale }
    }

    private companion object {
        const val SEPARATOR = " · "
        const val MISSING = "missing"
        const val SHORT = "short"
    }
}
