// PLACEHOLDER until lane T's generator lands. Lane T replaces this whole file with the output
// generated from web/src/styles/tokens.css; do not hand-edit beyond that. The values below are
// copied from tokens.css (2026-10-30) so the scaffold renders in the house palette meanwhile.
// ForeverTheme reads only the names declared here; if the generator names them differently,
// update ForeverTheme.kt in the same commit.
package gg.foreversixty.app.designsystem

import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp

object Tokens {
    object Colors {
        val Bg = Color(0xFF07090D)
        val Raised = Color(0xFF0D111A)
        val CardTop = Color(0xFF131824)
        val Line = Color(0xFF262E40)
        val Text = Color(0xFFE9E4D8)
        val Strong = Color(0xFFF2EEE4)
        val Muted = Color(0xFF9A9484)
        val Gold = Color(0xFFE5B955)
        val GoldDeep = Color(0xFFA8762A)
        val Ember = Color(0xFFD66E28)
        val Wipe = Color(0xFFF0736B)
    }

    object Radius {
        val Panel = 12.dp
        val Control = 8.dp
        val Pill = 6.dp
    }

    object TypeScale {
        val H1 = 34.sp
        val H2 = 26.sp
        val CardTitle = 18.sp
        val Body = 15.sp
        val BodyLineHeight = 22.sp
        val Meta = 13.sp
        val Micro = 11.sp
    }
}
