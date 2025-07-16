import 'package:flutter_test/flutter_test.dart';
import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:mental_health_companion/screens/login_screen.dart';
import 'package:mental_health_companion/l10n/app_localizations.dart';

void main() {
  testWidgets('Login screen renders inputs and button', (tester) async {
    await tester.pumpWidget(
      const MaterialApp(
        home: LoginScreen(),
        localizationsDelegates: [
          AppLocalizations.delegate,
          GlobalMaterialLocalizations.delegate,
          GlobalWidgetsLocalizations.delegate,
          GlobalCupertinoLocalizations.delegate,
        ],
        supportedLocales: [Locale('en')],
      ),
    );

    await tester.pumpAndSettle();

    expect(find.byType(TextField), findsNWidgets(2));
    expect(find.text('Login'), findsOneWidget);
  });
}
