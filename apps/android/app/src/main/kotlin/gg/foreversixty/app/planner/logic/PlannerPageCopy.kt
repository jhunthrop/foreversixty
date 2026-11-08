package gg.foreversixty.app.planner.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.raids.logic.GameClass
import gg.foreversixty.app.talents.TalentGuide
import gg.foreversixty.app.talents.TalentRules
import java.util.Locale
import kotlin.math.roundToInt

/** The planner page's words (`copy.ts` `plannerCopy`, `bandCompareCopy`) from the shared catalog. */
class PlannerPageCopy(private val res: Resources) {
    /** `Level 60`, `Levels 30 to 39` (`bandPhrase`). */
    fun phrase(band: Int): String =
        if (band >= TalentGuide.LAST_BAND) res.getString(R.string.plannerPage_bandLevel, band)
        else res.getString(R.string.plannerPage_bandLevels, band, band + BAND_SPAN)

    /** The phrase inside a sentence: `the levels 40 to 49 build`. */
    private fun lowerPhrase(band: Int): String = phrase(band).replaceFirstChar { it.lowercase(Locale.ROOT) }

    fun heading(band: Int): String = res.getString(R.string.plannerPage_bandHeading, phrase(band))

    fun loadLabel(band: Int): String = res.getString(R.string.plannerPage_load, lowerPhrase(band))

    /** `30 to 39`, `60`: the band as the BiS lists name it. */
    fun range(band: Int): String = TalentGuide.label(band, res)

    /** The band panel's line: the coverage, Matches, the difference named, or the points still to take. */
    fun line(compare: BandCompare): String {
        val covered = compare.coveredOf
        if (covered != null) {
            return res.getString(
                R.string.plannerPage_coversFewer, compare.bandPoints, covered, covered - compare.bandPoints,
                lowerPhrase(PlannerBand.nextBand(compare.build.band)),
            )
        }
        return when (val status = compare.status) {
            is BandStatus.Matches -> withBeyond(res.getString(R.string.plannerPage_matches), status.beyond)
            is BandStatus.Differs -> withBeyond(differs(status.count, compare), status.beyond)
            is BandStatus.ToTake -> res.getQuantityString(R.plurals.plannerPage_toTake, status.count, status.count)
        }
    }

    private fun withBeyond(line: String, beyond: Int): String =
        if (beyond > 0) "$line ${res.getString(R.string.plannerPage_beyond, beyond)}" else line

    private fun differs(count: Int, compare: BandCompare): String {
        val sides = listOfNotNull(
            compare.mine.takeIf { it.isNotEmpty() }?.let { res.getString(R.string.plannerPage_mineTakes, joinNames(it)) },
            compare.theirs.takeIf { it.isNotEmpty() }?.let { res.getString(R.string.plannerPage_theirsTake, lowerPhrase(compare.build.band), joinNames(it)) },
        )
        val joined = sides.reduceOrNull { first, second -> res.getString(R.string.plannerPage_sidesJoin, first, second) }.orEmpty()
        return res.getString(R.string.plannerPage_differs, res.getQuantityString(R.plurals.plannerPage_points, count, count), joined)
    }

    /** `A, B and C`. */
    private fun joinNames(names: List<String>): String =
        if (names.size < 2) names.joinToString() else res.getString(R.string.plannerPage_listAnd, names.dropLast(1).joinToString(", "), names.last())

    /** Under Points left: `21 of 25 spent at level 34` for a character, else `51 of 51 spent · level 60`. */
    fun spent(spent: Int, level: Int, character: Boolean): String {
        val available = PlannerRules.pointsAvailable(level)
        return when {
            character -> res.getString(R.string.plannerPage_spentAtLevel, spent, available, level)
            spent == 0 -> res.getString(R.string.plannerPage_spentOf, spent, available)
            else -> res.getString(R.string.plannerPage_spentOfLevel, spent, available, spent + TalentRules.BASE_LEVEL)
        }
    }

    /** The set score's figure, for a build that is its band's build whole; null without one. */
    fun score(compare: BandCompare?): Pair<String, String>? {
        val score = compare?.takeIf { it.settled }?.build?.setScore ?: return null
        val label = res.getString(if (compare.build.healer) R.string.plannerPage_hps else R.string.plannerPage_dps)
        return label to score.roundToInt().toString()
    }

    fun scoreBasis(band: Int): String = res.getString(R.string.plannerPage_scoreBasis, lowerPhrase(band))

    fun note(note: EditorNote): String = when (note) {
        is EditorNote.Loaded ->
            if (note.order.size < note.fullPoints) res.getString(R.string.plannerPage_loadedFor, note.level, note.order.size, note.fullPoints)
            else res.getString(R.string.plannerPage_loadedAt, note.level, note.order.size)
        is EditorNote.Linked -> res.getString(R.string.plannerPage_linkedNote)
        is EditorNote.Imported -> res.getString(R.string.plannerPage_importedNote)
    }

    /** `Dwarf · Level 34 · Beast Mastery Hunter`, the spec once a point names one. */
    fun meta(race: String, level: Int, spec: String?, className: String): String =
        if (spec == null) res.getString(R.string.plannerPage_metaLevel, race, level)
        else res.getString(R.string.plannerPage_metaSpec, race, level, spec, className)

    /** `Dwarf Hunter · Beast Mastery 21/Marksmanship 0/Survival 0 · Level 34`. */
    fun summary(race: String, className: String, index: TalentIndex, totals: List<Int>, level: Int): String {
        val trees = index.trees.mapIndexed { i, tree -> res.getString(R.string.plannerPage_treePoints, tree.name, totals.getOrElse(i) { 0 }) }.joinToString("/")
        return res.getString(R.string.plannerPage_summary, race, className, trees, level)
    }

    /** A class's name before its files load (`hunter` to `Hunter`). */
    fun className(slug: String): String = GameClass.of(slug)?.let { res.getString(it.label) } ?: PlannerEditorModel.titleCase(slug)

    companion object {
        private const val BAND_SPAN = 9

        /** `night-elf` to `Night Elf`, for a race named only by its slug. */
        fun raceName(slug: String?): String =
            slug.orEmpty().split('-').filter { it.isNotEmpty() }.joinToString(" ") { part -> part.replaceFirstChar { it.titlecase(Locale.ROOT) } }
    }
}
