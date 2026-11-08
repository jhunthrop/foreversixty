package gg.foreversixty.app.planner.logic

import android.net.Uri

/**
 * The codes and site links a planner build travels as: the FS1 head the site's planner and sim read
 * (`fs1.ts` `encodeFS1`, talents only: the planner owns no gear), the FSB1 code the addon's build
 * dialog reads (`fsb1.ts` `encodeFSB1`), and the site pages that take them.
 */
object PlannerCodes {
    private const val FS1 = "FS1"
    private const val FSB1 = "FSB1"
    private const val RADIX = 36
    private const val SITE = "https://foreversixty.gg"

    /** `FS1:<build>:<class>:<race>:<t1>/<t2>/<t3>:`: one base-36 digit per talent, in tree order. */
    fun fs1(index: TalentIndex, order: List<Int>, dataBuild: String, raceSlug: String): String {
        val ranks = PlannerRules.ranks(order)
        val trees = index.trees.joinToString("/") { tree -> tree.talents.joinToString("") { (ranks[it.id] ?: 0).toString(RADIX) } }
        return listOf(FS1, dataBuild, index.file.classSlug, raceSlug, trees, "").joinToString(":")
    }

    /** `FSB1:<build>:<class>:<tab tier column per point>:`: what `Copy for the game` puts on the clipboard. */
    fun fsb1(index: TalentIndex, order: List<Int>, dataBuild: String): String {
        val points = order.mapNotNull { id ->
            val talent = index.byId[id] ?: return@mapNotNull null
            val tree = index.treeOf[id] ?: return@mapNotNull null
            listOf(tree.position + 1, talent.tier + 1, talent.column + 1).joinToString("") { it.toString(RADIX) }
        }
        return listOf(FSB1, dataBuild, index.file.classSlug, points.joinToString(""), "").joinToString(":")
    }

    /** The site's sim for a build code (`/sim?code=`). */
    fun simFor(code: String): String = "$SITE/sim?code=${Uri.encode(code)}"

    /** The site's sim for a saved build (`/sim?source=build&ref=`). */
    fun simForBuild(id: String): String = "$SITE/sim?source=build&ref=${Uri.encode(id)}"

    /** The site's planner on a build code: where `Send this build to` a character lives. */
    fun sitePlanner(code: String): String = "$SITE/planner?code=${Uri.encode(code)}"

    /** The site's sim with nothing loaded. */
    const val SIM = "$SITE/sim"

    /** Where the addon and the companion are set up. */
    const val ADDON = "$SITE/setup#addon"
}
