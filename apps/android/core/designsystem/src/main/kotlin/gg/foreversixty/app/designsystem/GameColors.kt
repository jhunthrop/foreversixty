package gg.foreversixty.app.designsystem

import androidx.compose.ui.graphics.Color

/** The class colours by the API's class word (`warrior`, `shaman`, ...). */
fun classColor(className: String?): Color = when (className?.lowercase()) {
    "warrior" -> FsColor.classWarrior
    "paladin" -> FsColor.classPaladin
    "hunter" -> FsColor.classHunter
    "rogue" -> FsColor.classRogue
    "priest" -> FsColor.classPriest
    "shaman" -> FsColor.classShaman
    "mage" -> FsColor.classMage
    "warlock" -> FsColor.classWarlock
    "druid" -> FsColor.classDruid
    else -> FsColor.text
}
