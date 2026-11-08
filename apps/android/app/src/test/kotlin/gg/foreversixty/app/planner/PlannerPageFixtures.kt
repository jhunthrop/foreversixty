package gg.foreversixty.app.planner

import gg.foreversixty.app.bis.BisTestData
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.model.Build
import gg.foreversixty.app.model.TalentsFile
import gg.foreversixty.app.planner.logic.EditorNote
import gg.foreversixty.app.planner.logic.EditorPhase
import gg.foreversixty.app.planner.logic.EditorRequest
import gg.foreversixty.app.planner.logic.EditorState
import gg.foreversixty.app.planner.logic.EditorView
import gg.foreversixty.app.planner.logic.PlannerBand
import gg.foreversixty.app.planner.logic.PlannerEditorModel
import gg.foreversixty.app.planner.logic.PlannerRules
import gg.foreversixty.app.planner.logic.SharedBuildView
import gg.foreversixty.app.planner.logic.TalentIndex
import gg.foreversixty.app.planner.ui.EditorPageState
import gg.foreversixty.app.planner.ui.SharedPageState
import gg.foreversixty.app.planner.logic.SharedBuildPhase
import gg.foreversixty.app.talents.TalentRules
import java.io.File

/**
 * The planner mocks' states (ws_app.py `planner-*`, `build-shared`) on the repo's real hunter files:
 * Felsnap (Dwarf Beast Mastery Hunter 34, the 30 to 39 build loaded), a Night Elf Survival build from
 * a link that differs from the level 60 build by two points, the bare Warrior loading and failing,
 * and that Survival build shared.
 */
object PlannerPageFixtures {
    const val BUILD = BisTestData.BUILD

    private fun <T> data(path: String, strategy: kotlinx.serialization.DeserializationStrategy<T>): T =
        ApiJson.decodeFromString(strategy, File(BisTestData.repo, "data/builds/$BUILD/$path").readText())

    val hunter: TalentsFile = data("talents/hunter.json", TalentsFile.serializer())
    private val index = TalentIndex(hunter)
    private val catalog = BisTestData.catalog()
    private val bands: Map<String, BisFile> = listOf("hunter-beast-mastery", "hunter-marksmanship", "hunter-survival")
        .associateWith { data("bis/$it.json", BisFile.serializer()) }
    private val races = PlannerEditorModel.legalRaces(hunter.classId, PlannerTestData.races(), PlannerTestData.combos())

    private fun view(state: EditorState, note: EditorNote?) =
        EditorView(hunter, index, "Hunter", races, state, bands = bands, catalog = catalog, note = note)

    /** Felsnap at 34: the levels 30 to 39 build loaded, 21 of 25 points (the web's arrival). */
    val you: EditorPageState = run {
        val bm = requireNotNull(catalog.spec("hunter-beast-mastery"))
        val band = requireNotNull(PlannerBand.build(bm, bands.getValue(bm.spec), 30, "alliance"))
        val load = PlannerBand.loadAtLevel(index, band.talents, FELSNAP_LEVEL)
        val state = EditorState("hunter", "dwarf", FELSNAP_LEVEL, load.order)
        val request = EditorRequest.Draft("hunter", "dwarf", FELSNAP_LEVEL, "Beast Mastery")
        EditorPageState(
            EditorPhase.Ready(view(state, EditorNote.Loaded(state.order, FELSNAP_LEVEL, load.fullPoints))), BUILD, tree = 0, request = request,
            backTo = "Home", sendTo = "Felsnap", signedIn = true, forCharacter = true,
        )
    }

    /**
     * The level 60 Survival build with two points moved into Beast Mastery: a build from a link
     * that differs from the band's by two points.
     */
    val guideOrder: List<Int> = run {
        val survival = requireNotNull(catalog.spec("hunter-survival"))
        val band = requireNotNull(PlannerBand.build(survival, bands.getValue(survival.spec), 60, "alliance"))
        val ranks = TalentRules.ranksOfBand(band.talents).map { it.toMutableList() }
        val tree = index.trees[2]
        val prereqs = tree.talents.mapNotNull { it.prereqTalentId }.toSet()
        val moved = tree.talents.indexOfFirst { t -> ranks[2][tree.talents.indexOf(t)] >= 2 && t.maxRank > 2 && t.id !in prereqs && t.tier == MOVED_TIER }
        ranks[2][moved] -= 2
        ranks[0][index.trees[0].talents.indexOfFirst { it.tier == 0 && it.maxRank >= 2 }] += 2
        PlannerRules.orderFromRanks(index, ranks).order
    }

    val guide: EditorPageState = run {
        val state = EditorState("hunter", "night-elf", 60, guideOrder)
        val code = "FS1:$BUILD:hunter:night-elf"
        EditorPageState(
            EditorPhase.Ready(view(state, EditorNote.Linked(state.order))), BUILD, tree = 2,
            request = EditorRequest.Code("hunter", code, "night-elf"), backTo = "Home",
        )
    }

    private val bare = EditorRequest.Draft("warrior", "human")

    val loading = EditorPageState(EditorPhase.Loading, BUILD, tree = 0, request = bare, backTo = "Home")

    val failed = EditorPageState(EditorPhase.Failed, BUILD, tree = 0, request = bare, backTo = "Home")

    val shared: SharedPageState = run {
        val build = Build(
            id = "k7Qm2xa", classId = hunter.classId, raceId = NIGHT_ELF, treeVersion = BUILD, pointOrder = guideOrder,
            title = "Survival raid build", createdAt = "2026-11-04T19:00:00Z", views = 41,
        )
        val survival = requireNotNull(catalog.spec("hunter-survival"))
        val band = requireNotNull(PlannerBand.build(survival, bands.getValue(survival.spec), 60, "alliance"))
        val view = SharedBuildView(
            build, hunter, "hunter", "Hunter", "Night Elf",
            band = PlannerBand.compare(index, guideOrder, band, characterPoints = null), raceSlug = "night-elf",
        )
        SharedPageState(SharedBuildPhase.Loaded(view), tree = 2, backTo = "Home")
    }

    private const val FELSNAP_LEVEL = 34
    private const val NIGHT_ELF = 4
    /** Two points off a tier 4 talent keep every deeper row open. */
    private const val MOVED_TIER = 3
}
