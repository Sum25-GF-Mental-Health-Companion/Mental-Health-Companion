import 'package:flutter_test/flutter_test.dart';
import 'package:flutter/material.dart';
import 'package:mental_health_companion/screens/login_screen.dart';

void main() {
  testWidgets('Login screen renders inputs and button', (tester) async {
    await tester.pumpWidget(const MaterialApp(home: LoginScreen()));

    expect(find.byType(TextField), findsNWidgets(2));
    expect(find.text('Login'), findsOneWidget);
  });
}
