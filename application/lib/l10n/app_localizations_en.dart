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
}
