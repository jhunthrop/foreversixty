package gg.foreversixty.app.menu

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Row
import androidx.compose.material3.Text
import androidx.compose.ui.Alignment
import androidx.compose.ui.semantics.heading
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.ScaledType
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.clickable
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
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
import gg.foreversixty.app.designsystem.LocalParityRecorder
import gg.foreversixty.app.model.AccountCharacter

/** Opens the profile menu; provided by the app shell on the tab roots, null elsewhere. */
val LocalOpenMenu = staticCompositionLocalOf<(() -> Unit)?> { null }

/**
 * The profile button at the top left of a tab root's bar (spec 2026-11-08 section 3): the selected
 * character's Battle.net portrait in its faction ring, the person glyph signed out. It opens the menu,
 * where characters switch.
 */
@Composable
fun ProfileButton(session: SessionManager?, modifier: Modifier = Modifier) {
    CharacterButton(currentCharacter(session), modifier)
}

/** The profile button for [character]: its portrait, else its class crest, in its faction ring; the glyph signed out. */
@Composable
fun CharacterButton(character: AccountCharacter?, modifier: Modifier = Modifier) {
    val open = LocalOpenMenu.current ?: return
    val label = stringResource(R.string.menu_a11y_open)
    Box(
        modifier
            .padding(horizontal = FsSpace.space3, vertical = FsSpace.space2)
            .size(TOUCH)
            .clip(CircleShape)
            .clickable(role = Role.Button, onClick = open)
            .semantics { contentDescription = label },
    ) {
        // The bar is chrome above the content, as the mocks' nav sits outside `main`: its images never count.
        CompositionLocalProvider(LocalParityRecorder provides null) {
            CharacterPortrait(character?.avatarUrl, character?.faction, SIZE, Modifier.padding((TOUCH - SIZE) / 2))
        }
    }
}

/** The session's selected character; null signed out or with no session. */
@Composable
private fun currentCharacter(session: SessionManager?): AccountCharacter? {
    if (session == null) return null
    val phase by session.phase.collectAsStateWithLifecycle()
    val picked by session.pickedKey.collectAsStateWithLifecycle()
    return session.currentCharacter(phase, picked)
}

/**
 * A tab root's top bar row: the profile button at the leading edge, full width for the trailing
 * actions. A [title] follows the button as Material 3's top app bar titles its tab (spec Mobile,
 * Android: `Raid` in the display face).
 */
@Composable
fun ProfileBar(session: SessionManager?, title: String? = null) {
    CharacterBar(currentCharacter(session), title)
}

/** [ProfileBar] for a known [character] (the shell resolves it from the session; tests pass a fixture). */
@Composable
fun CharacterBar(character: AccountCharacter?, title: String? = null) {
    Row(Modifier.fillMaxWidth(), verticalAlignment = Alignment.CenterVertically) {
        CharacterButton(character)
        if (title != null) Text(title, style = ScaledType.barTitle, color = FsColor.strong, modifier = Modifier.semantics { heading() })
    }
}


private val SIZE = 32.dp
private val TOUCH = 48.dp
