package gg.foreversixty.app.navigation

import androidx.annotation.StringRes
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.List
import androidx.compose.material.icons.filled.DateRange
import androidx.compose.material.icons.filled.Home
import androidx.compose.material.icons.filled.Person
import androidx.compose.ui.graphics.vector.ImageVector
import gg.foreversixty.app.R

/**
 * The app's four tabs, in the spec's order (§3.1): Home, Raids, BiS, You. Feature lanes plug their
 * root screen into `ForeverAppScaffold`'s `when`, one branch each; the enum itself stays fixed.
 */
enum class ForeverTab(@param:StringRes val label: Int, val icon: ImageVector) {
    Home(R.string.tab_home, Icons.Filled.Home),
    Raids(R.string.tab_raids, Icons.Filled.DateRange),
    Bis(R.string.tab_bis, Icons.AutoMirrored.Filled.List),
    You(R.string.tab_you, Icons.Filled.Person),
}
