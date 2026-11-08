package gg.foreversixty.app.menu

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.offset
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.Person
import androidx.compose.material3.Icon
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import coil3.compose.AsyncImage
import gg.foreversixty.app.common.FactionMark
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.parityItem

/**
 * The Battle.net portrait in a faction-coloured ring with the faction emblem as a badge on its top right:
 * the profile button and the menu's characters only (DESIGN-SYSTEM.md tenet 2, the owner's scoped
 * exception; everywhere else is the class crest). Without a portrait, the person glyph fills the ring
 * and the badge stays. With no faction (signed out, no character), the glyph alone in a muted ring.
 * The parity inventory counts the portrait and the badge as images, as the mock's two `img`s.
 */
@Composable
fun CharacterPortrait(avatarUrl: String?, faction: String?, size: Dp, modifier: Modifier = Modifier) {
    val ring = faction?.let(::ringOf) ?: FsColor.muted
    Box(modifier.size(size)) {
        Box(
            Modifier
                .fillMaxSize()
                .border(RING, ring, CircleShape)
                .clip(CircleShape)
                .background(FsColor.raised),
            contentAlignment = Alignment.Center,
        ) {
            if (avatarUrl.isNullOrBlank()) {
                Icon(Icons.Outlined.Person, contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(size * GLYPH_RATIO))
            } else {
                AsyncImage(
                    model = avatarUrl,
                    contentDescription = null,
                    contentScale = ContentScale.Crop,
                    modifier = Modifier.fillMaxSize().parityItem(ParityItemKind.Image),
                )
            }
        }
        if (faction != null) {
            val badge = size * BADGE_RATIO
            Box(
                Modifier
                    .align(Alignment.TopEnd)
                    .offset(x = BADGE_OUT_X, y = -BADGE_OUT_Y)
                    .shadow(1.dp, CircleShape, ambientColor = FsColor.bg, spotColor = FsColor.bg)
                    .parityItem(ParityItemKind.Image),
            ) {
                FactionMark(faction, badge)
            }
        }
    }
}

/** The ring is the faction's bar colour (tokens `color.faction.allianceBar`, `hordeBar`). */
private fun ringOf(faction: String): Color = if (faction == "horde") FsColor.hordeBar else FsColor.allianceBar

private val RING = 2.dp
private val BADGE_OUT_X = 4.dp
private val BADGE_OUT_Y = 3.dp
private const val GLYPH_RATIO = 0.6f

/** The badge is 40% of the portrait's diameter (ws_app_frame.py BADGE). */
private const val BADGE_RATIO = 0.4f
