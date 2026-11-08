import DesignSystem
import SwiftUI

/// The profile menu as a modal drawer (flows spec 3, iOS): 85% of the width from the leading edge
/// over the dimmed screen and the tab bar. A tap on the dim, a swipe back or Escape closes it.
struct ProfileDrawer<Menu: View>: ViewModifier {
    @Bindable var shell: AppShell
    @ViewBuilder let menu: () -> Menu
    @Environment(\.accessibilityReduceMotion) private var reduceMotion

    /// `.drawer-wrap { grid-template-columns: 85% 15% }`.
    static var widthShare: CGFloat { 0.85 }
    /// `.dim { background: #000000a6 }`.
    static var dimOpacity: Double { 0.65 }
    /// How far a swipe back travels before the drawer closes.
    static var closeDistance: CGFloat { 60 }

    func body(content: Content) -> some View {
        content
            .overlay {
                GeometryReader { geo in
                    ZStack(alignment: .leading) {
                        if shell.menuOpen {
                            Color.black.opacity(Self.dimOpacity)
                                .ignoresSafeArea()
                                .onTapGesture { shell.closeMenu() }
                                .accessibilityHidden(true)
                                .transition(.opacity)
                            menu()
                                .frame(width: geo.size.width * Self.widthShare)
                                .overlay(alignment: .trailing) {
                                    Rectangle().fill(FSColor.line).frame(width: FSStroke.hairline)
                                }
                                .background(FSColor.raised.ignoresSafeArea())
                                .gesture(DragGesture(minimumDistance: 12).onEnded { value in
                                    if value.translation.width < -Self.closeDistance { shell.closeMenu() }
                                })
                                .accessibilityAddTraits(.isModal)
                                .accessibilityAction(.escape) { shell.closeMenu() }
                                .transition(.move(edge: .leading))
                        }
                    }
                }
                .animation(reduceMotion ? nil : .easeOut(duration: FSMotion.base), value: shell.menuOpen)
            }
    }
}
