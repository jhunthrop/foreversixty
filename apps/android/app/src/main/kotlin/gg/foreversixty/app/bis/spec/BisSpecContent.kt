package gg.foreversixty.app.bis.spec

import gg.foreversixty.app.bis.BisSelection
import gg.foreversixty.app.items.ItemLook

/**
 * One spec's best-in-slot page (mocks `bis-spec-60`, `bis-spec-healer-60`, `bis-spec-20`): pushed under
 * Gear for a player who is not signed in, after the class picker. Every word is composed already.
 */
data class BisSpecContent(
    val title: String,
    val classSlug: String,
    val meta: String,
    val answer: BisSpecAnswer,
    val controls: BisSpecControls,
    val summary: BisSpecSummary,
    val notes: List<String>,
    val cells: List<BisSpecCell>,
    val groups: List<BisSpecGroup>,
    val weights: BisSpecWeights?,
    val guide: String,
)

/** The head's answer: the set's figure, a healer's three facts, the line on how it was simmed, the links. */
data class BisSpecAnswer(
    val label: String,
    val value: String?,
    val unit: String,
    val facts: List<BisSpecFact>,
    val line: String,
    val rankingSim: String?,
    val openSim: Boolean,
)

/** A healer's fact under the figure (`OVERHEAL 6%`). */
data class BisSpecFact(val label: String, val value: String)

/** One choice of a segmented control in the Change panel. */
data class BisSpecOption(val label: String, val on: Boolean, val icon: String? = null, val pick: BisSelection? = null)

/** The bar under the head and the Change panel it opens: spec, level, faction, profile, race and sources. */
data class BisSpecControls(
    val bar: String,
    val specs: List<BisSpecOption>,
    val levels: List<BisSpecOption>,
    val factions: List<BisSpecOption>,
    val profiles: List<BisSpecOption>,
    val race: String,
    val sources: String,
)

/** `16 of 17 slots have a pick at this level. Sign in to see them against your gear.` and the next band. */
data class BisSpecSummary(val count: String, val nextBand: String?, val nextPick: BisSelection?)

/** One slot's cell: the pick, its stats and source, and its gain over the runner-up or the gap to it. */
data class BisSpecCell(
    val slot: String,
    val level: String?,
    val look: ItemLook?,
    val stats: String,
    val source: String,
    val gain: String?,
    val unit: String?,
    val versus: String?,
    val versusItem: ItemLook?,
)

/** Where to go: the cells grouped by the place they come from. */
data class BisSpecGroup(val place: String, val pieces: String, val cells: List<BisSpecCell>)

/** What to stack: the sentence, the legend, the weights in two columns and the addon line. */
data class BisSpecWeights(val sentence: String, val legend: String, val stats: List<BisSpecWeight>)

/** A stat's scale factor (null when not significant), its bar and its value per point. */
data class BisSpecWeight(val name: String, val factor: String?, val fraction: Float, val perPoint: String?)

/** Which list the plate shows. */
enum class BisSpecView { BySlot, WhereToGo }

/** What the page hands back. */
data class BisSpecActions(
    val onPick: (BisSelection) -> Unit = {},
    val onView: (BisSpecView) -> Unit = {},
    val onSources: () -> Unit = {},
    val onItem: (ItemLook) -> Unit = {},
    val onSignIn: () -> Unit = {},
    val onRankingSim: () -> Unit = {},
    val onOpenSim: () -> Unit = {},
    val onPlanner: () -> Unit = {},
    val onAddon: () -> Unit = {},
    val onGuide: () -> Unit = {},
)

/** The cells whose pick or runner-up names contain [query] (the list's search), groups with none left out. */
fun BisSpecContent.filtered(query: String): BisSpecContent {
    val needle = query.trim()
    if (needle.isEmpty()) return this
    fun BisSpecCell.matches() = listOfNotNull(look?.name, versusItem?.name).any { it.contains(needle, ignoreCase = true) }
    return copy(
        cells = cells.filter { it.matches() },
        groups = groups.mapNotNull { group -> group.cells.filter { it.matches() }.takeIf { it.isNotEmpty() }?.let { group.copy(cells = it) } },
    )
}
