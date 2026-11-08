package gg.foreversixty.app.account

import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.auth.session.AccountActionOutcome
import gg.foreversixty.app.auth.session.DeletionConsequence
import gg.foreversixty.app.auth.session.DeviceRow
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.flows.FlowBackBar
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode
import java.time.Instant

/** The account screens' mocked states (design/mocks/ws_app.py `devices`, `delete`, `expired`). */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class AccountParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun devices() = pushed("devices") { Devices() }

    @Test fun devicesAt2() = pushed("devices", fontScale = 2f) { Devices() }

    @Test fun deleteFree() = pushed("delete-free") { Delete(emptyList()) }

    @Test fun deleteFreeAt2() = pushed("delete-free", fontScale = 2f) { Delete(emptyList()) }

    @Test fun deletePremium() = pushed("delete-premium") { Delete(listOf(DeletionConsequence.PremiumPlan)) }

    @Test fun deleteFailed() = pushed("delete-failed") { Delete(emptyList(), AccountActionOutcome.Failed) }

    @Test fun deleteDone() = ParityCheck.state(compose, "delete-done") { AccountDeletedScreen {} }

    @Test fun deleteDoneAt2() = ParityCheck.state(compose, "delete-done", fontScale = 2f) { AccountDeletedScreen {} }

    @Test fun expired() = ParityCheck.state(compose, "expired") { SessionExpiredScreen {} }

    @Test fun expiredAt2() = ParityCheck.state(compose, "expired", fontScale = 2f) { SessionExpiredScreen {} }

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
    private fun Devices() = DevicesScreen(DevicesState.Loaded(AccountFixtures.devices), DevicesActions())

    @Composable
    private fun Delete(consequences: List<DeletionConsequence>, outcome: AccountActionOutcome? = null) =
        DeleteAccountScreen(DeleteState(consequences, AccountFixtures.BATTLETAG, "", outcome = outcome), DeleteActions())
}

/** The mock's account: Hunthrop#1894 on a Pixel 9, the companion on Windows and an iPad. */
object AccountFixtures {
    const val BATTLETAG = "Hunthrop#1894"

    val devices = listOf(
        DeviceRow("dvc_1", "Pixel 9", isThisDevice = true, signedIn = Instant.parse("2026-10-30T18:00:00Z"), lastUsed = RelativeAge.Now),
        DeviceRow("dvc_2", "Companion on Windows", isThisDevice = false, signedIn = Instant.parse("2026-10-24T18:00:00Z"), lastUsed = RelativeAge.Minutes(12), isCompanion = true),
        DeviceRow("dvc_3", "iPad Air", isThisDevice = false, signedIn = Instant.parse("2026-11-02T18:00:00Z"), lastUsed = RelativeAge.Days(3)),
    )
}
