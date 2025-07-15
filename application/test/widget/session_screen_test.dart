import 'package:flutter_test/flutter_test.dart';
import 'package:flutter/material.dart';
import 'package:mental_health_companion/screens/session_screen.dart';

void main() {
  testWidgets('Session screen displays timer and messages', (tester) async {
    await tester.pumpWidget(const MaterialApp(home: SessionScreen()));

    expect(find.byType(AppBar), findsOneWidget);
    expect(find.byType(ListView), findsOneWidget);
    expect(find.byType(TextField), findsOneWidget);
  });
}
