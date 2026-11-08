package gg.foreversixty.app.planner.logic

import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.bis.SpecInfo
import gg.foreversixty.app.common.Haptics
import gg.foreversixty.app.model.BisFile
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
    data class Draft(
        override val classSlug: String,
        val raceSlug: String? = null,
        val level: Int? = null,
        /** The character's spec (`Beast Mastery`): with no draft, its band build at [level] loads. */
        val specName: String? = null,
    ) : EditorRequest

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
    /** The class's BiS files by spec key: their bands carry the level bands' builds. */
    val bands: Map<String, BisFile> = emptyMap(),
    val catalog: SpecCatalog = SpecCatalog(emptyList()),
    /** Where the build on screen came from; shown while its order is unchanged. */
    val note: EditorNote? = null,
    /** The spec the editor opened for, before any point says otherwise. */
    val specName: String? = null,
) {
    val totals: List<Int> get() = PlannerRules.pointsPerTree(index, state.order)
    val race: RaceInfo? get() = races.firstOrNull { it.slug == state.raceSlug }

    /** The spec the build reads as: the tree with the most points. */
    val spec: SpecInfo? get() = PlannerBand.specFor(catalog, index, state.order, specName)

    /** The build against its band's published build, when the spec's file has that band. */
    val band: BandCompare?
        get() {
            val spec = spec ?: return null
            val file = bands[spec.spec] ?: return null
            val build = PlannerBand.build(spec, file, PlannerBand.bandFor(state.spent, state.level), race?.faction) ?: return null
            return PlannerBand.compare(index, state.order, build, PlannerRules.pointsAvailable(state.level))
        }

    /** The note, while the order it was written for is still the one on screen. */
    val shownNote: EditorNote? get() = note?.takeIf { it.order == state.order }
}

/** Where the build on screen came from (`copy.ts` notes): each holds the order it describes. */
sealed interface EditorNote {
    val order: List<Int>

    /** `Loaded at level 34 · 21 points`, or `Loaded for level 34 · 25 of 31 points` when cut. */
    data class Loaded(override val order: List<Int>, val level: Int, val fullPoints: Int) : EditorNote

    /** A code from a link (a guide's build, a rankings row). */
    data class Linked(override val order: List<Int>) : EditorNote

    /** An addon export: the game records ranks, never the order. */
    data class Imported(override val order: List<Int>) : EditorNote
}

/** What `Import` did with a pasted export. */
sealed interface ImportOutcome {
    data object Loaded : ImportOutcome

    /** Another class's export: the editor reopens on that class with the code. */
    data class OtherClass(val classSlug: String, val code: String) : ImportOutcome

    data object Unreadable : ImportOutcome
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
    private val catalog: SpecCatalog = SpecCatalog(emptyList()),
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
        val arrivesOnBand = arrivesOnBand(file.classSlug)
        val opened = opening(index, allRaces)
        val state = opened.copy(raceSlug = opened.raceSlug?.takeIf { slug -> races.any { it.slug == slug } } ?: races.firstOrNull()?.slug)
        val bands = catalog.specsOf(file.classSlug).mapNotNull { spec ->
            (source.bis(build, spec.spec) as? ApiResult.Success)?.data?.value?.let { spec.spec to it }
        }.toMap()
        val specName = (request as? EditorRequest.Draft)?.specName
        val note = (request as? EditorRequest.Code)?.let { EditorNote.Linked(state.order) }
        val view = EditorView(file, index, className, races, state, bands = bands, catalog = catalog, note = note, specName = specName)
        val arrived = if (arrivesOnBand) withBand(view) ?: view else view
        drafts.save(arrived.state.draft)
        mutable.value = EditorPhase.Ready(arrived)
    }

    /** A character's planner with nothing saved opens on its band's build at its level (the web's arrival). */
    private fun arrivesOnBand(classSlug: String): Boolean {
        val draft = request as? EditorRequest.Draft ?: return false
        return draft.level != null && drafts.load(classSlug) == null
    }

    /** The view with its band's build loaded, cut at its level; null when it has no band build. */
    private fun withBand(view: EditorView): EditorView? {
        val band = view.band ?: return null
        val load = PlannerBand.loadAtLevel(view.index, band.build.talents, view.state.level)
        val state = PlannerEditor.load(view.state, load.order, null)
        return view.copy(state = state, note = EditorNote.Loaded(state.order, state.level, load.fullPoints))
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

    /** `Load the levels 30 to 39 build`: the band's build replaces the points, one undo away. */
    fun loadBand() {
        val view = ready ?: return
        val loaded = withBand(view) ?: return
        haptics.play(EditOutcome.Added)
        drafts.save(loaded.state.draft)
        mutable.value = EditorPhase.Ready(loaded.copy(shareProblem = null))
    }

    /** `Import`: an addon export of this class loads as one legal order; another class reopens the editor. */
    fun import(code: String): ImportOutcome {
        val view = ready ?: return ImportOutcome.Unreadable
        val decoded = Fs1.decode(code) ?: return ImportOutcome.Unreadable
        if (decoded.classSlug != view.file.classSlug) return ImportOutcome.OtherClass(decoded.classSlug, code.trim())
        val order = PlannerRules.orderFromRanks(view.index, decoded.treeRanks).order
        val race = decoded.raceSlug.takeIf { slug -> view.races.any { it.slug == slug } }
        val state = PlannerEditor.load(view.state, order, race)
        drafts.save(state.draft)
        mutable.value = EditorPhase.Ready(view.copy(state = state, note = EditorNote.Imported(state.order), shareProblem = null))
        return ImportOutcome.Loaded
    }

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
