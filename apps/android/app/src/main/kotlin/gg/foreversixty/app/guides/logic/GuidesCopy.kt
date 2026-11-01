package gg.foreversixty.app.guides.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.GuideFile
import java.util.Locale

/** F9's sentences from the shared copy (apps/shared/copy/en.json); every figure comes from the data. */
class GuidesCopy(private val res: Resources) {
    fun role(role: String): String = res.getString(
        when (role) {
            TierChoice.HEALER -> R.string.guides_role_healer
            TierChoice.TANK -> R.string.guides_role_tank
            else -> R.string.guides_role_dps
        },
    )

    /** `DPS · Updated 24 Sep 2026`. */
    fun cardMeta(role: String, date: String): String = res.getString(R.string.guides_cardMeta, role(role), date)

    fun footer(summary: SavedSummary?): String = if (summary == null) {
        res.getString(R.string.guides_footerNone)
    } else {
        res.getString(R.string.guides_footer, summary.classNames.joinToString(LIST_SEPARATOR), Guides.formatSize(summary.bytes))
    }

    /** `Updated 24 Sep 2026 · Recommended races: Dwarf, Orc`. */
    fun headMeta(guide: GuideFile): String {
        val date = GuideDates.long(guide.updated)
        val races = guide.recommendedRaces.map { it.name }
        return if (races.isEmpty()) {
            res.getString(R.string.guides_headMetaNoRaces, date)
        } else {
            res.getString(R.string.guides_headMeta, date, races.joinToString(LIST_SEPARATOR))
        }
    }

    fun section(section: GuideSection): String = res.getString(
        when (section) {
            GuideSection.Overview -> R.string.guides_section_overview
            GuideSection.Talents -> R.string.guides_section_talents
            GuideSection.Rotation -> R.string.guides_section_rotation
            GuideSection.Stats -> R.string.guides_section_stats
            GuideSection.Gear -> R.string.guides_section_gear
            GuideSection.Races -> R.string.guides_section_races
        },
    )

    /** `Talents, section 2 of 6` (spec §6). */
    fun sectionLabel(sections: List<GuideSection>, section: GuideSection): String =
        res.getString(R.string.guides_a11y_section, section(section), sections.indexOf(section) + 1, sections.size)

    /** `Priority 1, Intellect` for the chip at [position] (from 0). */
    fun statLabel(position: Int, stat: String): String = res.getString(R.string.guides_a11y_stat, position + 1, stat)

    /** `5 ranks`, `1 rank`; null without a rank count. */
    fun ranks(maxRank: Int?): String? = maxRank?.let { res.getQuantityString(R.plurals.guides_ranks, it, it) }

    /** `Rank 10`; null for a spell with no rank. */
    fun rank(rank: Int?): String? = rank?.let { res.getString(R.string.guides_rank, it) }

    fun rotation(kind: RotationKind): String = res.getString(
        when (kind) {
            RotationKind.Opener -> R.string.guides_opener
            RotationKind.Priority -> R.string.guides_priority
            RotationKind.Cooldowns -> R.string.guides_cooldowns
        },
    )

    fun build(kind: BuildKind): String = res.getString(
        when (kind) {
            BuildKind.Raid -> R.string.guides_raidBuild
            BuildKind.Leveling -> R.string.guides_levelingBuild
        },
    )

    fun faction(faction: String): String =
        res.getString(if (faction == TierChoice.ALLIANCE) R.string.tiers_faction_alliance else R.string.tiers_faction_horde)

    fun tag(tag: ChangeTag): String = res.getString(
        when (tag) {
            ChangeTag.Site -> R.string.changes_tag_site
            ChangeTag.Game -> R.string.changes_tag_game
        },
    )

    /** A tier row to TalkBack (spec §6): name, rank of count, tier, gap, race, and a tie when there is one. */
    fun tierLabel(model: TierRowModel, count: Int): String {
        val row = model.row
        val tier = row.tier?.let { res.getString(R.string.tiers_a11y_tier, it) }.orEmpty()
        val base = if (row.rank == 1) {
            res.getString(R.string.tiers_a11y_top, row.name, count.toString(), tier, row.race)
        } else {
            val gap = res.getString(R.string.tiers_a11y_gap, String.format(Locale.US, GAP_FORMAT, row.gapPercent))
            res.getString(R.string.tiers_a11y_behind, row.name, row.rank.toString(), count.toString(), tier, gap, row.race)
        }
        return if (row.tie != null) "$base$LIST_SEPARATOR${res.getString(R.string.tiers_a11y_tie)}" else base
    }

    private companion object {
        const val LIST_SEPARATOR = ", "
        const val GAP_FORMAT = "%.1f"
    }
}
