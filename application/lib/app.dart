import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:mental_health_companion/l10n/app_localizations.dart';
import 'screens/login_screen.dart';
import 'screens/sign_up_screen.dart';
import 'screens/session_screen.dart';
import 'screens/summary_screen.dart';

class MentalHealthApp extends StatefulWidget {
  const MentalHealthApp({super.key});

  @override
  State<MentalHealthApp> createState() => _MentalHealthAppState();
}

class _MentalHealthAppState extends State<MentalHealthApp> {
  Locale _locale = const Locale('en');

  void changeLanguage() {
    setState(() {
      _locale = _locale.languageCode == 'en' ? const Locale('ru') : const Locale('en');
    });
  }

  ThemeMode _themeMode = ThemeMode.system;

  void toggleTheme() {
    setState(() {
      _themeMode =
          _themeMode == ThemeMode.light ? ThemeMode.dark : ThemeMode.light;
    });
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Mental Health Companion',
      themeMode: _themeMode,
      theme: ThemeData.light(),
      darkTheme: ThemeData.dark(),
      locale: _locale,
      supportedLocales: const [Locale('en'), Locale('ru')],
      localizationsDelegates: const [
        GlobalMaterialLocalizations.delegate,
        GlobalWidgetsLocalizations.delegate,
        GlobalCupertinoLocalizations.delegate,
        AppLocalizations.delegate,
      ],
      home: Builder(
        builder: (context) => StartScreen(
          toggleTheme: toggleTheme,
          changeLanguage: changeLanguage,
        ),
      ),
      onGenerateRoute: (settings) {
        switch (settings.name) {
          case '/login':
            return MaterialPageRoute(builder: (_) => const LoginScreen());
          case '/signup':
            return MaterialPageRoute(builder: (_) => const SignUpScreen());
          case '/session':
            return MaterialPageRoute(builder: (_) => const SessionScreen());
          case '/summary':
            return MaterialPageRoute(builder: (_) => const SummaryScreen());
          default:
            return null;
        }
      },
    );
  }
}

class StartScreen extends StatelessWidget {
  final VoidCallback toggleTheme;
  final VoidCallback changeLanguage;

  const StartScreen({
    super.key,
    required this.toggleTheme,
    required this.changeLanguage,
  });

  @override
  Widget build(BuildContext context) {
    final loc = AppLocalizations.of(context)!;

    return Scaffold(
      appBar: AppBar(
        title: Text(loc.appTitle),
      ),
      body: Center(
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            IconButton(
              icon: const Icon(Icons.language),
              onPressed: changeLanguage,
              tooltip: loc.changeLanguageTooltip,
            ),
            IconButton(
              icon: const Icon(Icons.brightness_6),
              onPressed: toggleTheme,
              tooltip: loc.toggleThemeTooltip,
            ),
            const SizedBox(height: 24),
            SizedBox(
              width: 200,
              child: ElevatedButton(
                style: ElevatedButton.styleFrom(
                  foregroundColor: Colors.white,
                  backgroundColor: Colors.deepPurpleAccent,
                ),
                onPressed: () => Navigator.pushNamed(context, '/login'),
                child: Text(loc.login),
              ),
            ),
            const SizedBox(height: 16),
            SizedBox(
              width: 200,
              child: ElevatedButton(
                style: ElevatedButton.styleFrom(
                  foregroundColor: Colors.white,
                  backgroundColor: Colors.deepPurpleAccent,
                ),
                onPressed: () => Navigator.pushNamed(context, '/signup'),
                child: Text(loc.signUp),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
