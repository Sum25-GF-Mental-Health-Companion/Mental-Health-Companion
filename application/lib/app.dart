import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'screens/login_screen.dart';
import 'screens/session_screen.dart';
import 'screens/summary_screen.dart';

class MentalHealthApp extends StatelessWidget {
  const MentalHealthApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Mental Health Companion',
      theme: ThemeData.light(),
      darkTheme: ThemeData.dark(),
      supportedLocales: const [Locale('en'), Locale('ru')],
      localizationsDelegates: const [
        GlobalMaterialLocalizations.delegate,
        GlobalWidgetsLocalizations.delegate,
        GlobalCupertinoLocalizations.delegate,
      ],
      initialRoute: '/',
      routes: {
        '/': (_) => const LoginScreen(),
        '/session': (_) => const SessionScreen(),
        '/summary': (_) => const SummaryScreen(),
      },
    );
  }
}
