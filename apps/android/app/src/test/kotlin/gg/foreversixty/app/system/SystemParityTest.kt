package gg.foreversixty.app.system

import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.flows.FlowBackBar
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.system.offline.OfflineCategory
import gg.foreversixty.app.system.offline.OfflineDataActions
import gg.foreversixty.app.system.offline.OfflineDataScreen
import gg.foreversixty.app.system.offline.OfflineDataState
import gg.foreversixty.app.system.offline.OfflineReport
import gg.foreversixty.app.system.offline.OfflineRow
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import gg.foreversixty.app.update.UpdateRequiredScreen
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode
import java.time.Instant
import java.time.ZoneId

/** Offline data and Update required against their mocks (design/mocks/ws_app.py `offline`, `update_required`). */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class SystemParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun offlineData() = pushed("offline-data") { Offline(SystemFixtures.report) }

    @Test fun offlineDataAt2() = pushed("offline-data", fontScale = 2f) { Offline(SystemFixtures.report) }

    @Test fun offlineDataEmpty() = pushed("offline-data-empty") { Offline(OfflineReport(emptyList())) }

    @Test fun offlineStrip() = pushed("offline-strip") { Offline(SystemFixtures.report, online = false) }

    @Test fun updateRequired() = ParityCheck.state(compose, "update-required") { UpdateRequiredScreen {} }

    @Test fun updateRequiredAt2() = ParityCheck.state(compose, "update-required", fontScale = 2f) { UpdateRequiredScreen {} }

    private fun pushed(screen: String, fontScale: Float = 1f, content: @Composable () -> Unit) =
        ParityCheck.state(compose, screen, fontScale) {
            ForeverTabScaffold(selected = ForeverTab.Home, onSelect = {}) {
                Column {
                    FlowBackBar("Home", {})
                    content()
                }
            }
        }

    @Composable
    private fun Offline(report: OfflineReport, online: Boolean = true) = OfflineDataScreen(
        OfflineDataState(report, keepClassesReady = true, wifiOnly = true, online = online, updating = false, note = null, checkedAt = SystemFixtures.CHECKED),
        SystemFixtures.NOW,
        OfflineDataActions(),
        SystemFixtures.ZONE,
    )
}

/** The mock's phone: two guides, the lists, talents, the search index and the icons, checked today at 4:12 pm. */
object SystemFixtures {
    val CHECKED: Instant = Instant.parse("2026-11-06T22:12:00Z")
    val NOW: Instant = Instant.parse("2026-11-06T23:00:00Z")
    val ZONE: ZoneId = ZoneId.of("America/Chicago")

    val report = OfflineReport(
        listOf(
            OfflineRow(OfflineCategory.Guides, listOf("warrior", "hunter"), 4_200_000, CHECKED),
            OfflineRow(OfflineCategory.Bis, emptyList(), 1_100_000, CHECKED),
            OfflineRow(OfflineCategory.Talents, emptyList(), 600_000, CHECKED.minusSeconds(DAY)),
            OfflineRow(OfflineCategory.Search, emptyList(), 2_300_000, CHECKED),
            OfflineRow(OfflineCategory.Icons, emptyList(), 18_400_000, CHECKED),
        ),
    )

    private const val DAY = 86_400L
}
