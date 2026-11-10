package gg.foreversixty.app.menu

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Row
import androidx.compose.material3.Text
import androidx.compose.ui.Alignment
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.clickable
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.staticCompositionLocalOf
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.SessionManager
import gg.foreversixty.app.designsystem.FsSpace

/** Opens the profile menu; provided by the app shell on the tab roots, null elsewhere. */
val LocalOpenMenu = staticCompositionLocalOf<(() -> Unit)?> { null }

/**
 * The profile button at the top left of a tab root's bar (spec 2026-11-08 section 3): the selected
 * character's Battle.net portrait in its faction ring, the person glyph signed out. It opens the menu,
 * where characters switch.
 */
@Composable
fun ProfileButton(session: SessionManager?, modifier: Modifier = Modifier) {
    val open = LocalOpenMenu.current ?: return
    if (session == null) {
        PortraitButton(null, null, open, modifier)
        return
    }
    val phase by session.phase.collectAsStateWithLifecycle()
    val picked by session.pickedKey.collectAsStateWithLifecycle()
    val current = session.currentCharacter(phase, picked)
    PortraitButton(current?.avatarUrl, current?.faction, open, modifier)
}

@Composable
private fun PortraitButton(avatarUrl: String?, faction: String?, open: () -> Unit, modifier: Modifier) {
    val label = stringResource(R.string.menu_a11y_open)
    Box(
        modifier
            .padding(horizontal = FsSpace.space3, vertical = FsSpace.space2)
            .size(TOUCH)
            .clip(CircleShape)
            .clickable(role = Role.Button, onClick = open)
            .semantics { contentDescription = label },
    ) {
        CharacterPortrait(avatarUrl, faction, SIZE, Modifier.padding((TOUCH - SIZE) / 2))
    }
}

/**
 * A tab root's top bar row: the profile button at the leading edge, full width for the trailing
 * actions. A [title] follows the button as Material 3's top app bar titles its tab (spec Mobile,
 * Android: `Raid` in the display face).
 */
@Composable
fun ProfileBar(session: SessionManager?, title: String? = null) {
    Row(Modifier.fillMaxWidth(), verticalAlignment = Alignment.CenterVertically) {
        ProfileButton(session)
        if (title != null) Text(title, style = BAR_TITLE, color = FsColor.strong, modifier = Modifier.semantics { heading() })
    }
}

/** The bar title: the display face at 22, Material 3's small top app bar size. */
private val BAR_TITLE = FsType.title.copy(fontSize = 22.sp, lineHeight = 28.sp, letterSpacing = 0.sp)

private val SIZE = 32.dp
private val TOUCH = 48.dp
