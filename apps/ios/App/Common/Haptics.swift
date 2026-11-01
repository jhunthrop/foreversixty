import SwiftUI
import UIKit

/// Every haptic the app plays, in one place (spec §5.7 Haptics, v1 §7): `.selection` for a
/// segment, chip, picker, tab section or character pick; `.success` when a sign-up, a share link
/// or a watch saves; `.error` when a sign-up is refused; `.impact(.medium)` for the planner's
/// long press. Pull to refresh, scrolls, meters and rows play nothing. Injected, so a test
/// records every call; the system mutes them as it decides (Reduce Motion, settings).
@MainActor
protocol Haptics {
    func play(_ moment: HapticMoment)
}

enum HapticMoment: Equatable, Sendable {
    case selection
    case success
    case error
    case impactMedium
}

extension Haptics {
    func selection() { play(.selection) }
    func success() { play(.success) }
    func error() { play(.error) }
}

struct SystemHaptics: Haptics {
    func play(_ moment: HapticMoment) {
        switch moment {
        case .selection: UISelectionFeedbackGenerator().selectionChanged()
        case .success: UINotificationFeedbackGenerator().notificationOccurred(.success)
        case .error: UINotificationFeedbackGenerator().notificationOccurred(.error)
        case .impactMedium: UIImpactFeedbackGenerator(style: .medium).impactOccurred()
        }
    }
}

private struct HapticsKey: EnvironmentKey {
    static var defaultValue: any Haptics { SystemHaptics() }
}

extension EnvironmentValues {
    /// The app's haptics; tests and captures inject a recorder.
    var haptics: any Haptics {
        get { self[HapticsKey.self] }
        set { self[HapticsKey.self] = newValue }
    }
}

/// Plays `moment` through the injected facade whenever `trigger` changes: the one way a view
/// asks for a haptic, in place of scattered `.sensoryFeedback` calls.
private struct HapticOnChange<Trigger: Equatable>: ViewModifier {
    let moment: HapticMoment
    let trigger: Trigger
    /// Plays only when this says so for the new value (a watch that turned on, not off).
    let when: (Trigger) -> Bool
    @Environment(\.haptics) private var haptics

    func body(content: Content) -> some View {
        content.onChange(of: trigger) { _, value in
            if when(value) { haptics.play(moment) }
        }
    }
}

extension View {
    func haptic<Trigger: Equatable>(_ moment: HapticMoment, trigger: Trigger, when: @escaping (Trigger) -> Bool = { _ in true }) -> some View {
        modifier(HapticOnChange(moment: moment, trigger: trigger, when: when))
    }
}
