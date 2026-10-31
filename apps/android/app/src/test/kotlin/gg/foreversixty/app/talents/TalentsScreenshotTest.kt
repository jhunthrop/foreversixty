package gg.foreversixty.app.talents

import android.content.Context
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onRoot
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.bis.BisTestData
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * The Talents screen's states (spec F6 §6, §7) at Pixel 9 size, at the default font scale and at
 * 2.0, with the repo's own icon files. iOS: TalentsSnapshotTests.swift.
 * Record: ./gradlew :app:recordRoborazziDebug   Verify: ./gradlew :app:verifyRoborazziDebug
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class TalentsScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    private val res = ApplicationProvider.getApplicationContext<Context>().resources
    private val shaman = TalentTestData.shaman()

    private fun presenter(
        yours: String? = TalentTestData.SHAMAN_YOURS,
        guide: TalentGuideBuild = TalentGuideBuild.Build(TalentTestData.SHAMAN_GUIDE, 60),
    ) = TalentsPresenter(shaman, guide, yours?.let { TalentExport.Ranks(TalentRules.ranksOfBand(it)) } ?: TalentExport.None, 60, "Elemental")

    private fun page(phase: TalentsPhase, shown: TalentBuildKind = TalentBuildKind.Yours, offline: String? = null, tree: Int = 0) = @Composable {
        TalentsPage(phase, BisTestData.BUILD, offline, shown, tree = tree)
    }

    private fun sheet(name: String) = @Composable {
        val id = TalentRules.trees(shaman).flatMap { it.talents }.first { it.name == name }.id
        TalentRankContent(requireNotNull(presenter().sheet(TalentBuildKind.Yours, id, res)), BisTestData.BUILD)
    }

    @Test fun `differs`() = both("talents_differs", page(TalentsPhase.Loaded(presenter(), null)))

    @Test fun `matches`() = capture("talents_matches_default", 1f, page = page(TalentsPhase.Loaded(presenter(TalentTestData.SHAMAN_GUIDE), null)))

    @Test fun `guide segment`() = capture("talents_guide_default", 1f, page = page(TalentsPhase.Loaded(presenter(), null), TalentBuildKind.Guide))

    @Test fun `no export`() =
        capture("talents_noexport_default", 1f, page = page(TalentsPhase.Loaded(presenter(yours = null), null), TalentBuildKind.Guide))

    @Test fun `no build`() = capture("talents_nobuild_default", 1f, page = page(TalentsPhase.Loaded(presenter(guide = TalentGuideBuild.None), null)))

    @Test fun `no spec`() = capture("talents_nospec_default", 1f, page = page(TalentsPhase.Loaded(presenter(guide = TalentGuideBuild.NoSpec), null)))

    @Test fun `loading`() = capture("talents_loading_default", 1f, page = page(TalentsPhase.Loading))

    @Test fun `offline`() = capture("talents_offline_default", 1f, page = page(TalentsPhase.Loaded(presenter(), BisTestData.NOW), offline = "9:41 pm"))

    @Test fun `error`() = capture("talents_error_default", 1f, page = page(TalentsPhase.Failed))

    @Test fun `rank sheet on the bottom-row capstone`() = capture("talents_rank_bottom_row_default", 1f, inShell = false, page = sheet("Elemental Mastery"))

    @Test fun `rank sheet on the right column with the guide line`() = both("talents_rank_right_column", sheet("Call of Flame"), inShell = false)

    @Test fun `restoration tab`() = capture("talents_restoration_default", 1f, page = page(TalentsPhase.Loaded(presenter(), null), tree = 2))

    /** The smallest supported Android phone (spec F6 §1). */
    @Config(qualifiers = SMALL_PHONE)
    @Test fun `first screen at 360 dp`() = capture("talents_differs_360_default", 1f, page = page(TalentsPhase.Loaded(presenter(), null)))

    @Test fun `rank sheet maxed`() = capture("talents_rank_maxed_default", 1f, inShell = false, page = sheet("Convection"))

    private fun both(name: String, page: @Composable () -> Unit, inShell: Boolean = true) {
        capture("${name}_default", 1f, inShell, page)
        capture("${name}_large", LARGE, inShell, page)
    }

    private fun capture(name: String, fontScale: Float, inShell: Boolean = true, page: @Composable () -> Unit) {
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
    }
}

/** Goldens are recorded on macOS and verified on Linux CI; allow anti-aliasing noise only. */
private val TOLERANT = RoborazziOptions(compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F))

/** 360 x 640 dp, the smallest supported phone. */
private const val SMALL_PHONE = "w360dp-h640dp-normal-notlong-notround-any-420dpi-keyshidden-nonav"

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
