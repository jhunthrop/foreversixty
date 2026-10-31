package gg.foreversixty.app.talents

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.BisBand
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.model.TalentsFile

/**
 * Which guide build the screen compares with, and whether an export can be drawn on the active
 * build's trees (spec F6 §4.2, §4.3). iOS: TalentGuide.swift.
 */
object TalentGuide {
    const val FIRST_BAND = 20
    const val LAST_BAND = 60
    private const val BAND_STEP = 10
    private const val RAID_PRESET = "raid"
    private const val BARE_PRESET = "bare"

    /** The band of a level: 60 from 60, `floor(level / 10) * 10` from 20 to 59, none under 20. */
    fun bandForLevel(level: Int?): Int? =
        if (level == null || level < FIRST_BAND) null else minOf(LAST_BAND, level / BAND_STEP * BAND_STEP)

    /** The planner's `selectBand`: the band's raid entry for the faction, else bare, else the first. */
    fun entry(file: BisFile, band: Int, faction: String?): BisBand? {
        val entries = file.bands.filter { it.band == band && it.faction == faction }
        return entries.firstOrNull { it.preset == RAID_PRESET }
            ?: entries.firstOrNull { it.preset == BARE_PRESET }
            ?: entries.firstOrNull()
    }

    /** The band's copy label: `60`, or `50 to 59` (the planner's `bandLabel`). */
    fun label(band: Int, res: Resources): String =
        if (band >= LAST_BAND) res.getString(R.string.bis_range_top, band)
        else res.getString(R.string.bis_range_span, band, band + BAND_STEP - 1)

    /** True when an export made on [older] can be read on [active]'s trees, digit for talent. */
    fun sameLayout(older: TalentsFile, active: TalentsFile): Boolean = layout(older) == layout(active)

    private fun layout(file: TalentsFile): List<List<Int>> = TalentRules.trees(file).map { tree -> tree.talents.map { it.id } }
}
