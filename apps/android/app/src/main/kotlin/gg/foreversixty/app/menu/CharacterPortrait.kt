package gg.foreversixty.app.menu

import androidx.compose.foundation.Image
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
import androidx.compose.runtime.staticCompositionLocalOf
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.res.painterResource
import coil3.compose.SubcomposeAsyncImage
import gg.foreversixty.app.common.crestOf
import gg.foreversixty.app.common.FactionMark
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.parityItem

/**
 * The Battle.net portrait in a faction-coloured ring with the faction emblem as a badge on its top right:
 * the profile button and the menu's characters only (DESIGN-SYSTEM.md tenet 2, the owner's scoped
 * exception; everywhere else is the class crest). Without a portrait, or when it fails to load, the
 * profile button and the menu's head draw the person glyph (flows spec 3: the class crest is Home's head
 * and is never drawn twice); the Switch character rows pass [className] so the crest tells them apart.
 * The parity inventory counts the portrait and the badge as images, as the mock's two `img`s.
 */
@Composable
fun CharacterPortrait(avatarUrl: String?, faction: String?, size: Dp, modifier: Modifier = Modifier, className: String? = null) {
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
            val crest = crestOf(className)
            val local = avatarUrl?.takeIf { it.isNotBlank() }?.let { LocalPortraitPainter.current?.painter(it) }
            when {
                avatarUrl.isNullOrBlank() -> Stand(crest, size)
                local != null -> Image(local, contentDescription = null, contentScale = ContentScale.Crop, modifier = Modifier.fillMaxSize().parityItem(ParityItemKind.Image))
                // Loading, or a portrait that fails (a Battle.net render 403, no network): the stand-in, never an empty ring.
                else -> SubcomposeAsyncImage(
                    model = avatarUrl,
                    contentDescription = null,
                    contentScale = ContentScale.Crop,
                    loading = { Stand(crest, size) },
                    error = { Stand(crest, size) },
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

/** What fills the ring without a portrait: the class crest when the caller names one, else the person glyph. */
@Composable
private fun Stand(crest: Int?, size: Dp) {
    if (crest != null) {
        Image(painterResource(crest), contentDescription = null, contentScale = ContentScale.Crop, modifier = Modifier.fillMaxSize().parityItem(ParityItemKind.Image))
    } else {
        Box(Modifier.fillMaxSize(), contentAlignment = Alignment.Center) {
            Icon(Icons.Outlined.Person, contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(size * GLYPH_RATIO))
        }
    }
}

/** Paints a Battle.net portrait from somewhere other than the network (the screenshot tests' repo avatars). */
fun interface PortraitPainterSource {
    @Composable
    fun painter(avatarUrl: String): Painter?
}

/** A portrait source that wins over the network when it has the picture, as [gg.foreversixty.app.items.LocalItemIconPainter] does for icons. */
val LocalPortraitPainter = staticCompositionLocalOf<PortraitPainterSource?> { null }

/** The ring is the faction's bar colour (tokens `color.faction.allianceBar`, `hordeBar`). */
private fun ringOf(faction: String): Color = if (faction == "horde") FsColor.hordeBar else FsColor.allianceBar

private val RING = 2.dp
private val BADGE_OUT_X = 4.dp
private val BADGE_OUT_Y = 3.dp
private const val GLYPH_RATIO = 0.6f

/** The badge is 40% of the portrait's diameter (ws_app_frame.py BADGE). */
private const val BADGE_RATIO = 0.4f
