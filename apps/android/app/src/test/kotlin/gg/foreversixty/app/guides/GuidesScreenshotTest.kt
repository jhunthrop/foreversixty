package gg.foreversixty.app.guides

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Modifier
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onRoot
import androidx.compose.ui.unit.dp
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.github.takahirom.roborazzi.RoborazziOptions
import com.github.takahirom.roborazzi.captureRoboImage
import gg.foreversixty.app.account.CharacterSelector
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.guides.logic.GuideSection
import gg.foreversixty.app.guides.logic.Guides
import gg.foreversixty.app.guides.logic.GuidesSegment
import gg.foreversixty.app.guides.logic.Loadable
import gg.foreversixty.app.guides.logic.TierChoice
import gg.foreversixty.app.guides.logic.Tiers
import gg.foreversixty.app.guides.ui.ClassPage
import gg.foreversixty.app.guides.ui.GuideReader
import gg.foreversixty.app.guides.ui.GuideReaderActions
import gg.foreversixty.app.guides.ui.GuideReaderState
import gg.foreversixty.app.guides.ui.GuidesHome
import gg.foreversixty.app.guides.ui.GuidesHomeActions
import gg.foreversixty.app.guides.ui.GuidesHomeState
import gg.foreversixty.app.guides.ui.NewDataState
import gg.foreversixty.app.guides.ui.TierListActions
import gg.foreversixty.app.guides.ui.TierListContent
import gg.foreversixty.app.guides.ui.TierListState
import gg.foreversixty.app.guides.ui.WhatChangedActions
import gg.foreversixty.app.guides.ui.WhatChangedContent
import gg.foreversixty.app.guides.ui.WhatChangedState
import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.GuideFile
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import java.io.File
import java.time.Instant
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RuntimeEnvironment
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * Every F9 screen and state on a Pixel 9 (412 x 915 dp) with the five-tab bar, at the default size
 * and at font scale 2.0 (spec W3 §9), from the W3-C1 fixtures. Record: ./gradlew :app:recordRoborazziDebug
 * Each PNG is also copied to design/reviews/2026-11-01-mobile-guides/build/and for the design review.
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PIXEL_9)
class GuidesScreenshotTest {
    @get:Rule
    val compose = createComposeRule()

    private val now = GuidesTestData.now
    private val offlineAt: Instant = Instant.ofEpochMilli(FakeGuideFiles.SAVED_AT)
    private val index = GuidesTestData.index()

    /** Elemental and Restoration saved, as on the board: Enhancement is the unsaved one. */
    private val saved = mapOf(GuidesTestData.ELEMENTAL to ELEMENTAL_BYTES, "shaman-restoration" to RESTORATION_BYTES, "warrior-protection" to PROTECTION_BYTES)

    private fun capture(name: String, large: Boolean = false, screen: @Composable () -> Unit) {
        if (large) RuntimeEnvironment.setFontScale(LARGE_SCALE)
        compose.setContent {
            ForeverTheme {
                CompositionLocalProvider(LocalItemIconPainter provides HomeFixtures.repoIcons) {
                    ForeverTabScaffold(selected = ForeverTab.Guides, onSelect = {}, tabs = ForeverTab.visible(guides = true)) {
                        Box(Modifier.fillMaxSize().background(FsColor.bg)) { screen() }
                    }
                }
            }
        }
        val path = "src/test/screenshots/guides/$name.png"
        compose.onRoot().captureRoboImage(path, roborazziOptions = TOLERANT)
        copyForReview(path, name)
    }

    /** The review copy, written only when the capture exists (record runs; verify runs compare). */
    private fun copyForReview(path: String, name: String) {
        val png = File(path)
        if (!png.exists()) return
        val folder = File(REVIEW_DIR).apply { mkdirs() }
        png.copyTo(File(folder, "$name.png"), overwrite = true)
    }

    @Composable
    private fun Root(segment: GuidesSegment, character: AccountCharacter?, content: @Composable () -> Unit) {
        GuidesRootLayout(GuidesSegment.visible(tiersOn = true), segment, {}, selector = { Selector(character) }, content = content)
    }

    @Composable
    private fun Selector(character: AccountCharacter?) {
        if (character != null) Row(Modifier.padding(horizontal = 16.dp, vertical = 8.dp)) { CharacterSelector(character, open = false) {} }
    }

    private fun home(name: String, character: AccountCharacter?, large: Boolean = false, phase: Loadable<gg.foreversixty.app.model.GuideIndex> = Loadable.Loaded(index, null)) =
        capture(name, large) {
            Root(GuidesSegment.Guides, character) {
                GuidesHome(GuidesHomeState(phase, saved, character?.let { Guides.yourGuide(index, it) }), GuidesHomeActions())
            }
        }

    private fun guide(
        name: String,
        file: GuideFile = GuidesTestData.elemental(),
        section: GuideSection = GuideSection.Overview,
        large: Boolean = false,
        phase: Loadable<GuideFile> = Loadable.Loaded(file, null),
        newData: NewDataState? = null,
    ) = capture(name, large) {
        GuideReader(GuideReaderState(phase, "Elemental Shaman", "shaman", newData, section), now, GuideReaderActions())
    }

    private fun tiers(name: String, character: AccountCharacter?, choice: TierChoice? = null, large: Boolean = false) = capture(name, large) {
        Root(GuidesSegment.Tiers, character) {
            TierListContent(
                TierListState(Loadable.Loaded(GuidesTestData.tiers(), null), choice ?: Tiers.defaultChoice(index, character), character?.let { Guides.yourGuide(index, it)?.spec }),
                now,
                TierListActions(),
            )
        }
    }

    // F9.1 the Guides tab, F9.7 loading

    @Test fun guides() = home("f9-guides-and", GuidesTestData.lightningbolltz)

    @Test fun guidesNoCharacter() = home("f9-guides-nochar-and", null)

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun guidesLarge() = home("f9-guides-large-and", GuidesTestData.lightningbolltz, large = true)

    @Test fun loading() = home("f9-loading-and", GuidesTestData.lightningbolltz, phase = Loadable.Loading)

    @Test fun guidesOffline() = home("f9-guides-offline-and", GuidesTestData.lightningbolltz, phase = Loadable.Loaded(index, offlineAt))

    // F9.2 a guide, F9.3 talents and rotation

    @Test fun guideTop() = guide("f9-guide-top-and")

    @Test fun guideTalents() = guide("f9-guide-talents-and", section = GuideSection.Talents)

    @Test fun guideRotation() = guide("f9-guide-rotation-and", section = GuideSection.Rotation)

    @Test fun guideStats() = guide("f9-guide-stats-and", section = GuideSection.Stats)

    @Test fun guideRaces() = guide("f9-guide-races-and", section = GuideSection.Races)

    @Test fun guideFireTalents() = guide("f9-guide-fire-talents-and", GuidesTestData.fire(), GuideSection.Talents)

    @Test fun guideProtection() = guide("f9-guide-protection-and", GuidesTestData.protection(), GuideSection.Talents)

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun guideLarge() = guide("f9-guide-large-and", large = true)

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun guideRotationLarge() = guide("f9-guide-rotation-large-and", section = GuideSection.Rotation, large = true)

    @Test fun guideLoading() = guide("f9-guide-loading-and", phase = Loadable.Loading)

    @Test fun guideNotCached() = guide("f9-guide-notcached-and", phase = Loadable.NotCached)

    @Test fun guideError() = guide("f9-guide-error-and", phase = Loadable.Failed)

    @Test fun guideOffline() = guide("f9-guide-offline-and", phase = Loadable.Loaded(GuidesTestData.elemental(), offlineAt))

    @Test fun newData() = guide("f9-newdata-and", newData = NewDataState())

    @Test fun newDataFailed() = guide("f9-newdata-failed-and", newData = NewDataState(failed = true))

    // F9.4 tier lists

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun tiersDps() = tiers("f9-tiers-dps-and", GuidesTestData.emberquill)

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun tiersFirst() = tiers("f9-tiers-first-and", GuidesTestData.lightningbolltz.copy(className = "warrior", spec = "Fury"))

    @Test fun tiersTank() = tiers("f9-tiers-tank-and", GuidesTestData.kraggor)

    @Test fun tiersHealer() = tiers("f9-tiers-healer-and", null, TierChoice(TierChoice.HEALER, TierChoice.HORDE))

    @Test fun tiersNoCharacter() = tiers("f9-tiers-nochar-and", null)

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun tiersLarge() = tiers("f9-tiers-large-and", GuidesTestData.emberquill, large = true)

    // F9.5 what changed

    @Test @Config(qualifiers = PIXEL_9_TALL)
    fun changes() = capture("f9-changes-and") {
        Root(GuidesSegment.Changes, GuidesTestData.lightningbolltz) {
            WhatChangedContent(
                WhatChangedState(Loadable.Loaded(GuidesTestData.dates(), null), Loadable.Loaded(GuidesTestData.changelog(), null), now.atZone(GuidesTestData.zone).toLocalDate()),
                now,
                WhatChangedActions(),
            )
        }
    }

    @Test fun changesFailed() = capture("f9-changes-error-and") {
        Root(GuidesSegment.Changes, GuidesTestData.lightningbolltz) {
            WhatChangedContent(WhatChangedState(Loadable.Loaded(GuidesTestData.dates(), null), Loadable.Failed, now.atZone(GuidesTestData.zone).toLocalDate()), now, WhatChangedActions())
        }
    }

    @Test fun changeEntry() = capture("f9-change-entry-and") {
        gg.foreversixty.app.guides.ui.ChangeReader(GuidesTestData.changelog().entries.first()) {}
    }

    // F9.6 a class page offline, one guide not saved

    @Test fun classOffline() = capture("f9-class-offline-and") {
        ClassPage("shaman", Loadable.Loaded(index, offlineAt), saved, now, back = {}) {}
    }

    @Test @Config(qualifiers = PIXEL_9_XL)
    fun classLarge() = capture("f9-class-large-and", large = true) {
        ClassPage("shaman", Loadable.Loaded(index, offlineAt), saved, now, back = {}) {}
    }
}

/** Goldens are recorded on macOS and verified on Linux CI; allow anti-aliasing noise only. */
private val TOLERANT = RoborazziOptions(compareOptions = RoborazziOptions.CompareOptions(changeThreshold = 0.005F))

private const val LARGE_SCALE = 2f
private const val ELEMENTAL_BYTES = 21_800L
private const val RESTORATION_BYTES = 20_400L
private const val PROTECTION_BYTES = 24_100L

/** From apps/android/app (Gradle's test working directory) to the review folder. */
private const val REVIEW_DIR = "../../../design/reviews/2026-11-01-mobile-guides/build/and"

/** Pixel 9: 412 x 915 dp at 420 dpi. */
private const val PIXEL_9 = "w412dp-h915dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private const val PIXEL_9_TALL = "w412dp-h1600dp-normal-long-notround-any-420dpi-keyshidden-nonav"
private const val PIXEL_9_XL = "w412dp-h2600dp-normal-long-notround-any-420dpi-keyshidden-nonav"
