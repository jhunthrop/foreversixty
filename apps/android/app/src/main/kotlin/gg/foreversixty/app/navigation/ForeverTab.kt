package gg.foreversixty.app.navigation

import androidx.annotation.StringRes
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.List
import androidx.compose.material.icons.filled.AccountCircle
import androidx.compose.material.icons.filled.DateRange
import androidx.compose.material.icons.filled.Home
import androidx.compose.material.icons.filled.PlayArrow
import androidx.compose.ui.graphics.vector.ImageVector
import gg.foreversixty.app.R

/**
 * The top-level destinations, from the lane plan's features: F2 character, F3 BiS, F4 guild and
 * raids, F5 sims, F1/F8 account. Placeholders until lane D's spec fixes the IA and the icons
 * (the house crests and game icons replace these Material stand-ins).
 */
enum class ForeverTab(@param:StringRes val label: Int, val icon: ImageVector) {
    Character(R.string.tab_character, Icons.Filled.Home),
    Bis(R.string.tab_bis, Icons.AutoMirrored.Filled.List),
    Raids(R.string.tab_raids, Icons.Filled.DateRange),
    Sims(R.string.tab_sims, Icons.Filled.PlayArrow),
    Account(R.string.tab_account, Icons.Filled.AccountCircle),
}
