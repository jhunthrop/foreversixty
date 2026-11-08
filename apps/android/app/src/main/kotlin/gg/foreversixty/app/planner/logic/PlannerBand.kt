package gg.foreversixty.app.planner.logic

import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.bis.SpecInfo
import gg.foreversixty.app.model.BisBand
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.talents.TalentGuide
import gg.foreversixty.app.talents.TalentRules

/** A level band's published build for one spec (a BiS band's talents and its nightly set score). */
data class BandBuild(
    val spec: SpecInfo,
    val band: Int,
    /** One base-36 digit per talent, trees dash-joined: the BiS band's `talents`. */
    val talents: String,
    /** The band set's simmed score with the band's own talents (`set_dps`), when the file has one. */
    val setScore: Double?,
    val healer: Boolean,
)

/** How a build stands against its band's build (`band-compare.ts` `bandDiffStatus`). */
sealed interface BandStatus {
    data class Matches(val beyond: Int) : BandStatus
    data class Differs(val count: Int, val beyond: Int) : BandStatus
    data class ToTake(val count: Int) : BandStatus
}

/**
 * The band panel's facts: the status, the talents each side takes above the other (ruling Q1), the
 * cells that differ, and the coverage line when the band build holds fewer points than the
 * character has (ruling MUST 3).
 */
data class BandCompare(
    val build: BandBuild,
    val status: BandStatus,
    val mine: List<String>,
    val theirs: List<String>,
    val differing: Set<Int>,
    val bandPoints: Int,
    /** The character's points when the band build covers fewer of them, else null. */
    val coveredOf: Int?,
) {
    /** A band build already taken whole: no Load, and its set score is this build's score. */
    val settled: Boolean get() = status is BandStatus.Matches && coveredOf == null
}

/** A band build cut at a level: the order to load and how many points the band holds in all. */
data class BandLoad(val order: List<Int>, val fullPoints: Int)

/** The web planner's band compare (`band-compare.ts`), ported. Pure: nothing here mutates. */
object PlannerBand {
    /** The band a visitor sees with nothing spent and no character (`VISITOR_DEFAULT_BAND`). */
    const val VISITOR_BAND = TalentGuide.LAST_BAND
    private const val BAND_STEP = 10
    private const val HEALER = "healer"

    /** The spec a build reads as: the tree with the most points, else the first; matched by name. */
    fun specFor(catalog: SpecCatalog, index: TalentIndex, order: List<Int>, fallback: String? = null): SpecInfo? {
        val points = PlannerRules.pointsPerTree(index, order)
        val tree = if (points.any { it > 0 }) index.trees.getOrNull(TalentRules.openingTree(points)) else null
        val classSlug = index.file.classSlug
        val name = tree?.name ?: fallback ?: index.trees.firstOrNull()?.name ?: return null
        return catalog.spec(classSlug, name) ?: catalog.specsOf(classSlug).firstOrNull { name.startsWith(it.name, ignoreCase = true) }
    }

    /** The band shown: the character's band while nothing is spent, else the band the build reaches. */
    fun bandFor(spent: Int, level: Int?): Int = when {
        spent == 0 -> level?.let(TalentGuide::bandForLevel) ?: VISITOR_BAND
        else -> TalentGuide.bandForLevel(spent + TalentRules.BASE_LEVEL) ?: TalentGuide.FIRST_BAND
    }

    /** The band's entry for the race's faction, as a build; null when the file has none. */
    fun build(spec: SpecInfo, file: BisFile, band: Int, faction: String?): BandBuild? {
        val entry: BisBand = TalentGuide.entry(file, band, faction?.lowercase()) ?: return null
        return BandBuild(spec, band, entry.talents, entry.setDps, entry.role == HEALER)
    }

    fun compare(index: TalentIndex, order: List<Int>, build: BandBuild, characterPoints: Int?): BandCompare {
        val ranks = PlannerRules.ranks(order)
        val bandRanks = bandRanks(index, build.talents)
        val diff = TalentRules.compare(index.file, ranks, bandRanks)
        val status = when {
            diff.missingCount == 0 -> BandStatus.Matches(diff.extraCount)
            diff.diffCount > 0 -> BandStatus.Differs(diff.diffCount, diff.extraCount)
            else -> BandStatus.ToTake(diff.missingCount)
        }
        val bandPoints = bandRanks.values.sum()
        val covered = characterPoints?.takeIf { status is BandStatus.Matches && bandPoints < it }
        val (mine, theirs) = names(index, ranks, bandRanks)
        return BandCompare(build, status, mine, theirs, diff.differing, bandPoints, covered)
    }

    /** The band's build as a legal order, cut at [level]'s points when that is fewer. */
    fun loadAtLevel(index: TalentIndex, talents: String, level: Int?): BandLoad {
        val full = PlannerRules.orderFromRanks(index, TalentRules.ranksOfBand(talents)).order
        val allowed = (level ?: TalentRules.BASE_LEVEL) - TalentRules.BASE_LEVEL
        return BandLoad(if (allowed > 0) full.take(allowed) else full, full.size)
    }

    /** The band after [band], for the coverage line (`the levels 40 to 49 build`). */
    fun nextBand(band: Int): Int = minOf(TalentGuide.LAST_BAND, band + BAND_STEP)

    /** Each talent's band rank, capped at its max rank. */
    private fun bandRanks(index: TalentIndex, talents: String): Map<Int, Int> {
        val trees = TalentRules.ranksOfBand(talents)
        return buildMap {
            index.trees.forEachIndexed { t, tree ->
                tree.talents.forEachIndexed { i, talent -> put(talent.id, minOf(trees.getOrNull(t)?.getOrNull(i) ?: 0, talent.maxRank)) }
            }
        }
    }

    /** `Rapid Killing 2/2` on the side that holds more; a one-rank talent by its name alone. */
    private fun names(index: TalentIndex, ranks: Map<Int, Int>, bandRanks: Map<Int, Int>): Pair<List<String>, List<String>> {
        val mine = mutableListOf<String>()
        val theirs = mutableListOf<String>()
        index.trees.flatMap { it.talents }.forEach { talent ->
            val built = ranks[talent.id] ?: 0
            val band = bandRanks[talent.id] ?: 0
            fun named(rank: Int) = if (talent.maxRank > 1) "${talent.name} $rank/${talent.maxRank}" else talent.name
            if (built > band) mine += named(built)
            if (band > built) theirs += named(band)
        }
        return mine to theirs
    }
}
