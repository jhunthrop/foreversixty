package gg.foreversixty.app.planner.logic

import gg.foreversixty.app.common.Haptics
import gg.foreversixty.app.model.Build
import gg.foreversixty.app.model.RaceClassCombo
import gg.foreversixty.app.model.RaceInfo
import gg.foreversixty.app.model.TalentsFile
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.talents.TalentLinks
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow

/** How the editor was opened (spec §3.3): each fixes the class; a code or a copy also brings points. */
sealed interface EditorRequest {
    val classSlug: String

    /** `Plan a build` (Talents, Your builds, a bare planner link): the class's draft, else empty. */
    data class Draft(override val classSlug: String, val raceSlug: String? = null, val level: Int? = null) : EditorRequest

    /** A guide's build or a rankings row: an FS1 code read as one legal order. */
    data class Code(override val classSlug: String, val code: String, val raceSlug: String? = null) : EditorRequest

    /** A shared build's `Edit a copy`: its points; the shared record is never changed. */
    data class Copy(override val classSlug: String, val build: Build) : EditorRequest
}

/** What the editor draws once its files are in. */
data class EditorView(
    val file: TalentsFile,
    val index: TalentIndex,
    /** `Shaman`, from classes.json; the crest's accessibility label. */
    val className: String,
    /** The class's legal, non-placeholder races (combos.json and races.json). */
    val races: List<RaceInfo>,
    val state: EditorState,
    /** The last share's problem, shown under the editor until the next edit. */
    val shareProblem: ShareOutcome? = null,
    val sharing: Boolean = false,
) {
    val totals: List<Int> get() = PlannerRules.pointsPerTree(index, state.order)
    val race: RaceInfo? get() = races.firstOrNull { it.slug == state.raceSlug }
}

sealed interface EditorPhase {
    data object Loading : EditorPhase
    /** The talent file is neither saved nor fetchable. */
    data object Failed : EditorPhase
    data class Ready(val view: EditorView) : EditorPhase
}

/**
 * The editor's controller (spec F13.1 to F13.3): loads the class's files, applies each edit through
 * [PlannerEditor], plays its haptic, saves the draft as it goes and shares. iOS: PlannerEditorModel.swift.
 */
class PlannerEditorModel(
    private val source: PlannerSource,
    private val drafts: PlannerDrafts,
    val request: EditorRequest,
    private val haptics: Haptics,
) {
    private val mutable = MutableStateFlow<EditorPhase>(EditorPhase.Loading)
    val phase: StateFlow<EditorPhase> = mutable.asStateFlow()
    var build: String = ""
        private set

    private val ready: EditorView? get() = (mutable.value as? EditorPhase.Ready)?.view

    suspend fun load() {
        if (ready != null) return
        mutable.value = EditorPhase.Loading
        build = source.activeBuild()
        val file = (source.talents(build, request.classSlug) as? ApiResult.Success)?.data?.value
        if (file == null) {
            mutable.value = EditorPhase.Failed
            return
        }
        val index = TalentIndex(file)
        val allRaces = (source.races(build) as? ApiResult.Success)?.data?.value.orEmpty()
        val combos = (source.combos(build) as? ApiResult.Success)?.data?.value.orEmpty()
        val races = legalRaces(file.classId, allRaces, combos)
        val classes = (source.classes(build) as? ApiResult.Success)?.data?.value.orEmpty()
        val className = classes.firstOrNull { it.id == file.classId }?.name ?: titleCase(file.classSlug)
        val opened = opening(index, allRaces)
        val state = opened.copy(raceSlug = opened.raceSlug?.takeIf { slug -> races.any { it.slug == slug } } ?: races.firstOrNull()?.slug)
        drafts.save(state.draft)
        mutable.value = EditorPhase.Ready(EditorView(file, index, className, races, state))
    }

    /** The state the editor opens on: the draft, or the request's points over it. */
    private fun opening(index: TalentIndex, races: List<RaceInfo>): EditorState {
        val draft = drafts.load(request.classSlug)
        val base = draft?.let { EditorState(it.classSlug, it.raceSlug, it.level, it.order) }
            ?: EditorState(request.classSlug, null, PlannerEditor.MAX_LEVEL, emptyList())
        return when (val r = request) {
            is EditorRequest.Draft -> if (draft != null) {
                base.copy(raceSlug = base.raceSlug ?: r.raceSlug)
            } else {
                PlannerEditor.setLevel(base.copy(raceSlug = r.raceSlug), r.level ?: PlannerEditor.MAX_LEVEL)
            }
            is EditorRequest.Code -> {
                val decoded = Fs1.decode(r.code)?.takeIf { it.classSlug == r.classSlug }
                if (decoded == null) base else PlannerEditor.load(base, PlannerRules.orderFromRanks(index, decoded.treeRanks).order, r.raceSlug ?: decoded.raceSlug)
            }
            is EditorRequest.Copy -> PlannerEditor.load(base, copiedOrder(index, r.build.pointOrder), races.firstOrNull { it.id == r.build.raceId }?.slug)
        }
    }

    // Edits

    fun add(talentId: Int) = edit { PlannerEditor.add(it.index, it.state, talentId) }

    fun remove(talentId: Int) = edit { PlannerEditor.remove(it.index, it.state, talentId) }

    fun undo() = edit { PlannerEditor.undo(it.state) }

    fun reset() = edit { PlannerEditor.reset(it.state) }

    fun setLevel(level: Int) = edit { Edit(PlannerEditor.setLevel(it.state, level), EditOutcome.Unchanged) }

    fun setRace(raceSlug: String) = edit { Edit(PlannerEditor.setRace(it.state, raceSlug), EditOutcome.Unchanged) }

    /** The strip's talent, chosen without a point (the hardware keyboard's arrows). */
    fun touch(talentId: Int) = edit { Edit(it.state.copy(touched = talentId, refusal = null), EditOutcome.Unchanged) }

    private fun edit(change: (EditorView) -> Edit) {
        val view = ready ?: return
        val result = change(view)
        haptics.play(result.outcome)
        if (result.state.draft != view.state.draft) drafts.save(result.state.draft)
        mutable.value = EditorPhase.Ready(view.copy(state = result.state, shareProblem = null))
    }

    // Share

    /** `Share anyway`: one POST, never queued; the URL on success, else the problem under the editor. */
    suspend fun share(title: String): ShareOutcome {
        val view = ready ?: return ShareOutcome.Failed
        val race = view.race ?: return ShareOutcome.Failed.also(::problem)
        mutable.value = EditorPhase.Ready(view.copy(sharing = true, shareProblem = null))
        val input = PlannerShare.input(view.file.classId, race.id, build, view.state.order, title)
        val outcome = PlannerShare.outcome(source.save(input))
        if (outcome is ShareOutcome.Shared) haptics.success() else haptics.error()
        problem(outcome.takeUnless { it is ShareOutcome.Shared })
        return outcome
    }

    private fun problem(outcome: ShareOutcome?) {
        val view = ready ?: return
        mutable.value = EditorPhase.Ready(view.copy(sharing = false, shareProblem = outcome))
    }

    companion object {
        /** The class's playable races, in races.json's order (spec F13.1 Race chip). */
        fun legalRaces(classId: Int, races: List<RaceInfo>, combos: List<RaceClassCombo>): List<RaceInfo> {
            val legal = combos.filter { it.classId == classId }.map { it.raceId }.toSet()
            return races.filter { it.id in legal && !it.placeholder }
        }

        /**
         * A shared build's order as it was saved when every point is legal on the active trees,
         * else one legal order to the same ranks (its own tree version may differ from today's).
         */
        fun copiedOrder(index: TalentIndex, order: List<Int>): List<Int> {
            if (PlannerRules.validate(index, order).isEmpty()) return order
            val ranks = PlannerRules.ranks(order)
            val treeRanks = index.trees.map { tree -> tree.talents.map { ranks[it.id] ?: 0 } }
            return PlannerRules.orderFromRanks(index, treeRanks).order
        }

        /** `shaman` to `Shaman`, for a class missing from classes.json. */
        fun titleCase(slug: String): String = slug.replaceFirstChar { it.titlecase(java.util.Locale.ROOT) }

        /** "Night Elf" from an account character to the planner's `night-elf`. */
        fun raceSlug(race: String?): String? = race?.takeIf { it.isNotBlank() }?.let(TalentLinks::raceSlug)
    }
}
