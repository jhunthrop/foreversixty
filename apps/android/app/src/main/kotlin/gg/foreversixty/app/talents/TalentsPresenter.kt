package gg.foreversixty.app.talents

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.Talent
import gg.foreversixty.app.model.TalentTree
import gg.foreversixty.app.model.TalentsFile

/** Which build the tree draws (spec F6 §5.1 segments). */
enum class TalentBuildKind { Guide, Yours }

/** Your export, as the screen can use it (spec F6 §4.3, §6). */
sealed interface TalentExport {
    /** Ranks per tree, read on the active build's trees. */
    data class Ranks(val trees: List<List<Int>>) : TalentExport
    data object None : TalentExport
    /** Made on a build whose trees differ from the active build's. */
    data object Stale : TalentExport
}

/** The guide build, as the screen can use it (spec F6 §4.2, §6). */
sealed interface TalentGuideBuild {
    data class Build(val talents: String, val band: Int) : TalentGuideBuild
    /** No band, no file or no entry: `talents.noList`. */
    data object None : TalentGuideBuild
    /** A shared build's ranks per tree (lane F15, F13.5 `Compare with mine`): the other side as is. */
    data class Ranks(val trees: List<List<Int>>) : TalentGuideBuild
    /** The character has no spec, so there is no BiS file to read. */
    data object NoSpec : TalentGuideBuild
    /** The BiS file failed to load. */
    data object Failed : TalentGuideBuild
}

/** The tag kinds the status line uses (spec F6 §5.1 item 3). */
enum class TalentTagKind { Matches, Differs, Muted }

/** The status block: a tag (or none) and the sentence; [isError] draws it as the failed line. */
data class TalentStatusLine(
    val tag: String?,
    val tagKind: TalentTagKind,
    val sentence: String,
    val isError: Boolean = false,
) {
    val accessibilityLabel: String get() = listOfNotNull(tag, sentence).joinToString(". ")
}

/** One tree tab: `{Tree} {points}`, plus the dot for a tree holding a dashed cell. */
data class TalentTabModel(
    val index: Int,
    val name: String,
    val points: Int,
    val differing: Int,
    val against: CompareAgainst = CompareAgainst.Guide,
)

/** One drawn cell: the talent, its rank and state, and on Your build its signed difference. */
data class TalentCellModel(
    val talent: Talent,
    val rank: Int,
    val state: TalentCellState,
    /** Your rank minus the guide's, on Your build when they differ; null otherwise. */
    val delta: Int?,
    val guideRank: Int?,
    val against: CompareAgainst = CompareAgainst.Guide,
) {
    val badge: String get() = "$rank/${talent.maxRank}"
    val chip: String? get() = delta?.let { if (it > 0) "+$it" else "−${-it}" }
}

/** The rank sheet (spec F6 §7). */
data class TalentSheetModel(
    val talent: Talent,
    val rank: Int,
    val state: TalentCellState,
    val current: String?,
    val next: String?,
    val requirements: List<String>,
    val guideLine: String?,
    /** True when your rank is over the guide's: the guide line's ring is wipe, else gold. */
    val isOver: Boolean,
)

/**
 * Everything the Talents screen shows, pure over the files and the character (spec F6). The
 * counts, dashes and totals all come from one pass over the drawn ranks. iOS: TalentsPresenter.swift.
 */
class TalentsPresenter(
    val file: TalentsFile,
    val guide: TalentGuideBuild,
    val export: TalentExport,
    val level: Int?,
    val specName: String,
    /** What the other side is, and so the compare's words (a band's guide build, or a shared build). */
    val against: CompareAgainst = CompareAgainst.Guide,
) {
    val trees: List<TalentTree> = TalentRules.trees(file)
    val guideRanks: Map<Int, Int>? = when (guide) {
        is TalentGuideBuild.Build -> TalentRules.rankById(file, TalentRules.ranksOfBand(guide.talents))
        is TalentGuideBuild.Ranks -> TalentRules.rankById(file, guide.trees)
        else -> null
    }
    val yourRanks: Map<Int, Int>? = (export as? TalentExport.Ranks)?.let { TalentRules.rankById(file, it.trees) }

    val diff: TalentDiff? = guideRanks?.let { g -> yourRanks?.let { y -> TalentRules.compare(file, y, g) } }

    fun bandLabel(res: Resources): String? = (guide as? TalentGuideBuild.Build)?.let { TalentGuide.label(it.band, res) }

    val yourTotal: Int? get() = yourRanks?.values?.sum()

    /** The builds that can be drawn; the segments show only when both can. */
    val builds: List<TalentBuildKind>
        get() = listOfNotNull(guideRanks?.let { TalentBuildKind.Guide }, yourRanks?.let { TalentBuildKind.Yours })
    val showsSegments: Boolean get() = builds.size == 2
    val hasTree: Boolean get() = builds.isNotEmpty()

    /** Your build when it can be drawn, else the guide (spec F6 §5.4). */
    val initialBuild: TalentBuildKind get() = if (yourRanks == null) TalentBuildKind.Guide else TalentBuildKind.Yours

    fun ranks(build: TalentBuildKind): Map<Int, Int> =
        (if (build == TalentBuildKind.Yours) yourRanks else guideRanks) ?: yourRanks ?: guideRanks ?: emptyMap()

    // The status and notes

    fun status(res: Resources): TalentStatusLine = when (export) {
        is TalentExport.Ranks -> when (guide) {
            TalentGuideBuild.Failed -> TalentStatusLine(null, TalentTagKind.Muted, res.getString(R.string.talents_guideFailed), isError = true)
            TalentGuideBuild.NoSpec -> muted(res, R.string.talents_yours, R.string.talents_pickSpec)
            TalentGuideBuild.None -> muted(res, R.string.talents_yours, R.string.talents_noList)
            is TalentGuideBuild.Build, is TalentGuideBuild.Ranks -> compareStatus(res)
        }
        TalentExport.Stale -> TalentStatusLine(guideTag(res), TalentTagKind.Muted, res.getString(R.string.talents_staleExport))
        TalentExport.None -> TalentStatusLine(guideTag(res), TalentTagKind.Muted, res.getString(R.string.home_next_talentsNone))
    }

    private fun guideTag(res: Resources): String? = guideRanks?.let { res.getString(against.side) }

    private fun muted(res: Resources, tag: Int, sentence: Int) =
        TalentStatusLine(res.getString(tag), TalentTagKind.Muted, res.getString(sentence))

    private fun compareStatus(res: Resources): TalentStatusLine {
        val diff = diff ?: return muted(res, against.yours, R.string.talents_noList)
        val band = bandLabel(res).orEmpty()
        val differs = res.getString(R.string.talents_differs)
        return when (val status = diff.status) {
            TalentStatus.Matches -> {
                val sentence = if (diff.extraCount > 0) against.matchBeyond else against.match
                TalentStatusLine(res.getString(R.string.talents_matches), TalentTagKind.Matches, res.getString(sentence, specName, band))
            }
            is TalentStatus.Differs ->
                TalentStatusLine(differs, TalentTagKind.Differs, res.getString(against.differ, status.count, yourTotal ?: 0, band))
            is TalentStatus.ToTake ->
                TalentStatusLine(differs, TalentTagKind.Differs, res.getString(against.toTake, status.count, band))
        }
    }

    /** The secondary lines under the status, in order (spec F6 §5.1 item 4). */
    fun notes(res: Resources): List<String> = buildList {
        val yours = yourTotal
        if (yours != null && level != null) {
            val unspent = (TalentRules.pointsAvailable(level) - yours).coerceAtLeast(0)
            if (unspent > 0) add(res.getQuantityString(R.plurals.talents_unspent, unspent, unspent))
        }
        val diff = diff
        if (diff != null && diff.diffCount > 0 && diff.missingCount > 0) {
            add(res.getQuantityString(against.missing, diff.missingCount, diff.missingCount))
        }
        if (diff != null && diff.extraCount > 0) add(res.getString(against.beyond, diff.extraCount))
        // No export and no guide either: say both, and draw no tree (spec F6 §6, No export).
        if (!hasTree) add(res.getString(R.string.talents_noList))
    }

    // Tabs and the tree

    fun tabs(build: TalentBuildKind): List<TalentTabModel> {
        val points = TalentRules.pointsPerTree(file, ranks(build))
        val differing = differingIds(build)
        return trees.mapIndexed { index, tree ->
            TalentTabModel(index, tree.name, points[index], tree.talents.count { it.id in differing }, against)
        }
    }

    /** The tree that opens: the most points in the build first shown (spec F6 §5.3). */
    val openingTree: Int get() = TalentRules.openingTree(TalentRules.pointsPerTree(file, ranks(initialBuild)))

    fun cells(build: TalentBuildKind, tree: Int): List<TalentCellModel> {
        val drawn = trees.getOrNull(tree) ?: return emptyList()
        val ranks = ranks(build)
        val marks = if (build == TalentBuildKind.Yours) guideRanks else null
        return drawn.talents.map { talent ->
            val rank = ranks[talent.id] ?: 0
            val guideRank = marks?.get(talent.id)
            val delta = guideRank?.let { if (rank == it) null else rank - it }
            TalentCellModel(talent, rank, TalentRules.cellState(talent, drawn, ranks), delta, guideRank, against)
        }
    }

    fun connectors(build: TalentBuildKind, tree: Int): List<TalentConnector> =
        trees.getOrNull(tree)?.let { TalentRules.connectors(it, ranks(build)) }.orEmpty()

    fun legend(build: TalentBuildKind, res: Resources): String =
        res.getString(if (differingIds(build).isEmpty()) R.string.talents_tapHint else against.legend)

    private fun differingIds(build: TalentBuildKind): Set<Int> =
        if (build == TalentBuildKind.Yours) diff?.differing.orEmpty() else emptySet()

    // The sheet

    fun sheet(build: TalentBuildKind, talentId: Int, res: Resources): TalentSheetModel? {
        val tree = trees.firstOrNull { tree -> tree.talents.any { it.id == talentId } } ?: return null
        val talent = tree.talents.first { it.id == talentId }
        val ranks = ranks(build)
        val rank = ranks[talent.id] ?: 0
        val guideRank = if (build == TalentBuildKind.Yours) guideRanks?.get(talent.id) else null
        val differs = guideRank != null && guideRank != rank
        val guideLine = guideRank?.takeIf { differs }?.let { res.getString(against.rank, it, talent.maxRank) }
        return TalentSheetModel(
            talent = talent,
            rank = rank,
            state = TalentRules.cellState(talent, tree, ranks),
            current = if (rank > 0) talent.ranks.getOrNull(rank - 1)?.description else null,
            next = if (rank < talent.maxRank) talent.ranks.getOrNull(rank)?.description else null,
            requirements = requirements(talent, tree, ranks, res),
            guideLine = guideLine,
            isOver = guideRank != null && differs && rank > guideRank,
        )
    }

    /** The game's requirement lines, only when unmet in the drawn build (spec F6 §7). */
    private fun requirements(talent: Talent, tree: TalentTree, ranks: Map<Int, Int>, res: Resources): List<String> = buildList {
        val needed = TalentRules.POINTS_PER_TIER * talent.tier
        if (tree.talents.sumOf { ranks[it.id] ?: 0 } < needed) add(res.getString(R.string.talents_requiresTree, needed, tree.name))
        val prereq = talent.prereqTalentId?.let { id -> tree.talents.firstOrNull { it.id == id } }
        if (prereq != null && !TalentRules.prerequisiteMet(talent, ranks)) {
            val points = talent.prereqRank ?: 1
            add(
                if (points == 1) res.getString(R.string.talents_requiresTalentOne, prereq.name)
                else res.getString(R.string.talents_requiresTalent, points, prereq.name),
            )
        }
    }
}

/** A cell's accessibility label (spec F6 §11). */
fun TalentCellModel.accessibilityLabel(res: Resources): String =
    if (guideRank != null && delta != null) res.getString(against.a11yCell, talent.name, rank, talent.maxRank, guideRank)
    else res.getString(R.string.talents_a11yCell, talent.name, rank, talent.maxRank)

/** A tab's accessibility value: its points, and how many talents differ when it carries the dot. */
fun TalentTabModel.accessibilityValue(res: Resources): String =
    if (differing > 0) res.getString(against.a11yTab, points, differing)
    else res.getQuantityString(R.plurals.talents_points, points, points)
