package gg.foreversixty.app.planner

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onRoot
import androidx.compose.ui.unit.dp
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.ExperimentalRoborazziApi
import com.github.takahirom.roborazzi.captureRoboImage
import com.github.takahirom.roborazzi.captureScreenRoboImage
import gg.foreversixty.app.planner.ui.ResetDialog
import gg.foreversixty.app.bis.BisTestData
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.model.Build
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.planner.logic.BuildRow
import gg.foreversixty.app.planner.logic.EditorPhase
import gg.foreversixty.app.planner.logic.EditorState
import gg.foreversixty.app.planner.logic.EditorView
import gg.foreversixty.app.planner.logic.PlannerEditor
import gg.foreversixty.app.planner.logic.PlannerEditorModel
import gg.foreversixty.app.planner.logic.SharedBuildPhase
import gg.foreversixty.app.planner.logic.SharedBuildView
import gg.foreversixty.app.planner.logic.TalentIndex
import gg.foreversixty.app.planner.logic.YourBuildsPhase
import gg.foreversixty.app.planner.ui.BuildsActions
import gg.foreversixty.app.planner.ui.BuildsPageState
import gg.foreversixty.app.planner.ui.EditorActions
import gg.foreversixty.app.planner.ui.EditorPageState
import gg.foreversixty.app.planner.ui.NewBuildRow
import gg.foreversixty.app.planner.ui.PlannerEditorPage
import gg.foreversixty.app.planner.ui.ShareSheetContent
import gg.foreversixty.app.planner.ui.ShareSheetState
import gg.foreversixty.app.planner.ui.SharedActions
import gg.foreversixty.app.planner.ui.SharedBuildPage
import gg.foreversixty.app.planner.ui.SharedPageState
import gg.foreversixty.app.planner.ui.YourBuildsPage
import java.time.ZoneOffset
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * The planner's boards (spec F13.1 to F13.6) at Pixel 9 size, at the default font scale and at 2.0,
 * with the repo's own icon files. Record: ./gradlew :app:recordRoborazziDebug   Verify: :app:verifyRoborazziDebug
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class PlannerScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    private val shaman = PlannerTestData.shaman()
    private val index = TalentIndex(shaman)
    private val races = PlannerEditorModel.legalRaces(shaman.classId, PlannerTestData.races(), PlannerTestData.combos())
    private val full = PlannerTestData.elementalOrder()
    private fun talent(name: String) = index.trees.flatMap { it.talents }.first { it.name == name }.id

    private fun editor(
        order: List<Int>,
        touched: Int? = null,
        tree: Int = 0,
        refusalOf: Int? = null,
        undo: List<List<Int>> = listOf(emptyList()),
    ): @Composable () -> Unit {
        var state = EditorState("shaman", "orc", PlannerEditor.MAX_LEVEL, order, undo = undo, touched = touched)
        if (refusalOf != null) state = PlannerEditor.add(index, state, refusalOf).state
        val phase = EditorPhase.Ready(EditorView(shaman, index, "Shaman", races, state))
        return { PlannerEditorPage(EditorPageState(phase, BisTestData.BUILD, tree), EditorActions()) }
    }

    /** The Elemental talent with the longest rank text: the strip's clamp at its worst. */
    private val longest = index.trees[0].talents.maxBy { talent -> talent.ranks.maxOfOrNull { it.description.length } ?: 0 }.id

    /** 49 of the raid build's 51 points, the strip on the longest text: the board's `2 left`. */
    private val planned = editor(full.dropLast(2), touched = longest)

    @Test fun `F13_1 plan a build with points left`() = both("f13-1-plan", planned)

    /** Reset's alert over the editor: Clear in the error colour, and Cancel. */
    @OptIn(ExperimentalRoborazziApi::class)
    @Test fun `F13_1 the reset alert`() {
        RuntimeEnvironment.setFontScale(1f)
        compose.setContent {
            ForeverTheme {
                CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons) {
                    planned()
                    ResetDialog(full.size - 2, {}, {})
                }
            }
        }
        compose.waitForIdle()
        captureScreenRoboImage("src/test/screenshots/f13-1-reset_default.png", roborazziOptions = TOLERANT)
    }

    /** A fresh planner: nothing to undo, so Undo is at 40%. */
    @Test fun `F13_1 fresh`() = capture("f13-1-fresh_default", 1f, page = editor(emptyList(), undo = emptyList()))

    /** One tap later: the point, the strip on it, and Undo at full strength. */
    @Test fun `F13_1 after one tap`() {
        val tapped = PlannerEditor.add(index, EditorState("shaman", "orc", PlannerEditor.MAX_LEVEL, emptyList()), talent("Convection")).state
        capture("f13-1-undo_default", 1f, page = editor(tapped.order, touched = tapped.touched, undo = tapped.undo))
    }

    @Test fun `F13_1 at zero left`() = capture("f13-1-zero_default", 1f, page = editor(full, touched = talent("Elemental Mastery")))

    /** Elemental at 14 points: Lightning Overload is tier 4 and refuses (spec F13.2). */
    @Test fun `F13_2 a tier refusal`() {
        val elemental = index.trees[0].talents.map { it.id }.toSet()
        val fourteen = full.filter { it in elemental }.take(14) + full.filter { it !in elemental }
        both("f13-2-refusal", editor(fourteen, refusalOf = talent("Lightning Overload")))
    }

    @Test fun `F13_3 the share sheet with a title`() = capture("f13-3-share_default", 1f) {
        Box(Modifier.fillMaxSize()) {
            planned()
            Box(Modifier.fillMaxSize().background(Color.Black.copy(alpha = SCRIM)))
            Box(
                Modifier.align(Alignment.BottomCenter).fillMaxWidth().background(FsColor.raised, RoundedCornerShape(topStart = 28.dp, topEnd = 28.dp)),
            ) {
                Box(Modifier.align(Alignment.TopCenter).padding(top = 24.dp)) {
                    ShareSheetContent(ShareSheetState("Elemental raid", sharing = false, signedIn = true), {}, {}, {})
                }
            }
        }
    }

    private fun sharedView(build: Build = PlannerTestData.sharedBuild()) = SharedBuildView(build, shaman, "shaman", "Shaman", "Orc")

    private fun shared(phase: SharedBuildPhase, tree: Int = 0): @Composable () -> Unit = {
        SharedBuildPage(SharedPageState(phase, tree, zone = ZoneOffset.UTC), SharedActions())
    }

    @Test fun `F13_5 a shared build without an export of the class`() =
        capture("f13-5-shared-nocompare_default", 1f, page = shared(SharedBuildPhase.Loaded(sharedView())))

    @Test fun `F13_5 an unknown id`() = capture("f13-5-notfound_default", 1f, page = shared(SharedBuildPhase.NotFound))

    @Test fun `F13_5 offline with nothing saved`() = capture("f13-5-offline_default", 1f, page = shared(SharedBuildPhase.Offline))

    private fun row(build: Build, totals: List<Int>) = BuildRow(build, "shaman", "Shaman", "Orc", totals)

    private val rows = listOf(
        row(PlannerTestData.sharedBuild(), listOf(30, 3, 18)),
        row(PlannerTestData.sharedBuild(title = null, views = 3).copy(id = "a1", createdAt = "2026-10-29T10:00:00Z"), listOf(31, 20, 0)),
        row(PlannerTestData.sharedBuild(title = null, views = 1).copy(id = "a2", createdAt = "2026-10-21T10:00:00Z"), listOf(31, 0, 20)),
    )

    private fun builds(phase: YourBuildsPhase): @Composable () -> Unit = {
        YourBuildsPage(BuildsPageState(phase, NewBuildRow("shaman", "Shaman", 60), ZoneOffset.UTC), BuildsActions())
    }

    @Test fun `F13_6 your builds`() = both("f13-6-builds", builds(YourBuildsPhase.Loaded(rows)))

    @Test fun `F13_6 no builds yet`() = capture("f13-6-empty_default", 1f, page = builds(YourBuildsPhase.Loaded(emptyList())))

    private fun both(name: String, page: @Composable () -> Unit) {
        capture("${name}_default", 1f, page = page)
        capture("${name}_large", LARGE, page = page)
    }

    private fun capture(name: String, fontScale: Float, inShell: Boolean = name.startsWith("f13-5") || name.startsWith("f13-6"), page: @Composable () -> Unit) {
        RuntimeEnvironment.setFontScale(fontScale)
        compose.setContent {
            ForeverTheme {
                CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons) {
                    if (inShell) ForeverTabScaffold(selected = ForeverTab.Home, onSelect = {}) { page() } else page()
                }
            }
        }
        compose.onRoot().captureRoboImage("src/test/screenshots/$name.png", roborazziOptions = TOLERANT)
    }

    private companion object {
        const val LARGE = 2f
        const val SCRIM = 0.55f
    }
}

/** Goldens are recorded on macOS and verified on Linux CI; allow anti-aliasing noise only. */
private val TOLERANT = RoborazziOptions(compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F))

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
