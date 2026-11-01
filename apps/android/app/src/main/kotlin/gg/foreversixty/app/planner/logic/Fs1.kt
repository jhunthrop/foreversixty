package gg.foreversixty.app.planner.logic

/** The head of an FS1 code the planner reads: whose build it is and the ranks of its three trees. */
data class Fs1Build(val dataBuild: String, val classSlug: String, val raceSlug: String, val treeRanks: List<List<Int>>)

/**
 * The talent half of `fs1.ts` `decodeFS1`: `FS1:<build>:<class>:<race>:<t1>/<t2>/<t3>:<gear>[|...]`,
 * each tree one base-36 digit per talent. Gear and the version-2 sections are not the planner's
 * (one job per page), so they are not read. Null for anything that is not a readable FS1 head.
 */
object Fs1 {
    private const val PREFIX = "FS1"
    private const val TREES = 3
    private const val RADIX = 36
    private const val MIN_FIELDS = 6
    private const val MAX_LENGTH = 16_384

    fun decode(code: String): Fs1Build? {
        if (code.length > MAX_LENGTH) return null
        val parts = code.trim().substringBefore('|').split(':')
        if (parts.size < MIN_FIELDS || parts[0] != PREFIX) return null
        val trees = parts[4].split('/')
        if (trees.size != TREES) return null
        val ranks = trees.map { tree -> tree.map { digit -> digit.digitToIntOrNull(RADIX) ?: return null } }
        return Fs1Build(dataBuild = parts[1], classSlug = parts[2], raceSlug = parts[3], treeRanks = ranks)
    }
}
