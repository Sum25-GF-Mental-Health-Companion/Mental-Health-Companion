// ignore: unused_import
import 'package:intl/intl.dart' as intl;
import 'app_localizations.dart';

// ignore_for_file: type=lint

/// The translations for English (`en`).
class AppLocalizationsEn extends AppLocalizations {
  AppLocalizationsEn([String locale = 'en']) : super(locale);

  @override
  String get loginTitle => 'Login as Student';

  @override
  String get email => 'Email';

  @override
  String get password => 'Password';

  @override
  String get loginButton => 'Login';

  @override
  String get loginFailed => 'Login failed';

  @override
  String get sessionTitle => 'Session';

  @override
  String get summaryTitle => 'Session Summaries';

  @override
  String sessionX(Object number) {
    return 'Session $number';
  }

  @override
  String get changeLanguageTooltip => 'Change language';

  @override
  String get toggleThemeTooltip => 'Toggle theme';

  @override
  String get appTitle => 'Mental Health Companion';

  @override
  String get login => 'Login';

  @override
  String get signUp => 'Sign Up';

  @override
  String get registrationFailed => 'Registration Failed';

  @override
  String get invalidEmail => 'Invalid email';

  @override
  String get weakPassword => 'Password should contain at least 6 characters';

  @override
  String get signUpTitle => 'Sign Up';

  @override
  String get signUpButton => 'Sign Up';

  @override
  String get typeMessage => 'Type a message...';
}
