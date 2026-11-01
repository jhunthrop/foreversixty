import DesignSystem
import SwiftUI

/// The new-data strip (spec §5.7, board F15.5): the v1 strip's 36 pt on a gold 10% ground, the
/// line and a trailing gold `Update` with a 44 pt target. It never moves what is being read: the
/// screen changes only on `Update`. F15 extends this with the stale state.
struct NewDataStrip: View {
    let text: String
    let action: () -> Void

    private static let height: CGFloat = 36
    private static let target: CGFloat = 44
    private static let groundOpacity = 0.10

    var body: some View {
        HStack(spacing: FSSpacing.space2) {
            Image(systemName: "arrow.down.circle").font(.system(size: 14)).accessibilityHidden(true)
            Text(text).font(ScaledType.meta).fixedSize(horizontal: false, vertical: true)
            Spacer(minLength: 0)
            Button(Copy.offlineUpdate, action: action)
                .font(ScaledType.bodyStrong)
                .foregroundStyle(FSColor.gold)
                .frame(minWidth: Self.target, minHeight: Self.target)
                .contentShape(Rectangle())
        }
        .foregroundStyle(FSColor.text)
        .padding(.horizontal, FSSpacing.space4)
        .frame(minHeight: Self.height)
        .background(FSColor.gold.opacity(Self.groundOpacity))
        .overlay(alignment: .bottom) { FSColor.line.frame(height: 1) }
        .accessibilityElement(children: .contain)
        .onAppear { AccessibilityNotification.Announcement(text).post() }
    }
}

extension NewDataStrip {
    /// The strip for a state, or nil while no newer build is waiting.
    init?(_ state: NewDataState, action: @escaping () -> Void) {
        switch state {
        case .none: return nil
        case .ready: self.init(text: Copy.offlineNewData, action: action)
        case .failed: self.init(text: Copy.offlineUpdateFailed, action: action)
        }
    }
}
