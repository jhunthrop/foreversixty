import Foundation
import Models

/// The raid leader's read (coordinator ruling, player #5), drawn from the API's `composition`
/// block. The app only words it; every count and every present/missing comes from the API.
struct CompositionRead: Equatable, Sendable {
    struct ClassLine: Equatable, Sendable, Identifiable {
        let className: String
        let name: String
        let count: Int
        var id: String { className }
    }

    struct BuffLine: Equatable, Sendable, Identifiable {
        let key: String
        let name: String
        let label: String
        let icon: String
        let present: Bool
        let accessibilityLabel: String
        var id: String { key }
    }

    let classes: [ClassLine]
    /// `Tanks 3/4` against the raid's targets, else `Tanks 3`.
    let roles: [String]
    let buffs: [BuffLine]
    /// `Missing: Blessings`, or `Every raid buff is covered.`
    let buffSummary: String
    /// `Classes coming: Warrior 5, Paladin 0, …`
    let classesAccessibilityLabel: String

    init?(_ composition: RaidComposition?) {
        guard let composition else { return nil }
        classes = composition.classes.map { entry in
            ClassLine(className: entry.className, name: GameClass(rawValue: entry.className.lowercased())?.name ?? entry.className, count: entry.count)
        }
        let groups = [Copy.raidGroupTanks, Copy.raidGroupHealers, Copy.raidGroupDps]
        let counts = [composition.roles.tanks, composition.roles.healers, composition.roles.dps]
        let targets = composition.targets.map { [$0.tanks, $0.healers, $0.dps] }
        roles = groups.indices.map { index in
            targets.map { Copy.raidsCompRole(role: groups[index], n: counts[index], target: $0[index]) }
                ?? Copy.raidsCompRoleCount(role: groups[index], n: counts[index])
        }
        buffs = composition.buffs.map { buff in
            BuffLine(
                key: buff.key, name: buff.name, label: buff.label, icon: buff.icon, present: buff.present,
                accessibilityLabel: buff.present ? Copy.raidsA11yBuffPresent(name: buff.name, n: buff.count) : Copy.raidsA11yBuffMissing(name: buff.name)
            )
        }
        let missing = composition.buffs.filter { !$0.present }.map(\.name)
        buffSummary = missing.isEmpty ? Copy.raidsCompCovered : Copy.raidsCompMissing(names: missing.joined(separator: ", "))
        classesAccessibilityLabel = Copy.raidsA11yClasses(
            list: classes.map { Copy.raidsA11yClassCount(className: $0.name, n: $0.count) }.joined(separator: ", ")
        )
    }
}
