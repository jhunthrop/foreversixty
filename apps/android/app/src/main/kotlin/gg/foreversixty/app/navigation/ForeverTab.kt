package gg.foreversixty.app.navigation

import androidx.annotation.StringRes
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.List
import androidx.compose.material.icons.filled.DateRange
import androidx.compose.material.icons.filled.Home
import androidx.compose.material.icons.filled.Person
import androidx.compose.runtime.Composable
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.compose.ui.graphics.vector.rememberVectorPainter
import androidx.compose.ui.res.painterResource
import gg.foreversixty.app.R

/** A tab's glyph: a Material icon, or a bundled vector for one Material core lacks. */
sealed interface TabIcon {
    @Composable
    fun painter(): Painter

    data class Vector(val image: ImageVector) : TabIcon {
        @Composable
        override fun painter(): Painter = rememberVectorPainter(image)
    }

    data class Drawable(val id: Int) : TabIcon {
        @Composable
        override fun painter(): Painter = painterResource(id)
    }
}

/**
 * The app's tabs, in the spec's order (W3 §3.1): Home, Raids, BiS, Guides, You. Guides shows only
 * while its `guides` switch is on ([visible]). Feature lanes plug their root screen into
 * `ForeverAppScaffold`'s `when`, one branch each.
 */
enum class ForeverTab(@param:StringRes val label: Int, val icon: TabIcon) {
    Home(R.string.tab_home, TabIcon.Vector(Icons.Filled.Home)),
    Raids(R.string.tab_raids, TabIcon.Vector(Icons.Filled.DateRange)),
    Bis(R.string.tab_bis, TabIcon.Vector(Icons.AutoMirrored.Filled.List)),
    Guides(R.string.tab_guides, TabIcon.Drawable(R.drawable.ic_guides)),
    You(R.string.tab_you, TabIcon.Vector(Icons.Filled.Person)),
    ;

    companion object {
        /** The tabs drawn: the five, or four while the Guides switch is off. */
        fun visible(guides: Boolean): List<ForeverTab> = if (guides) entries else entries - Guides
    }
}
