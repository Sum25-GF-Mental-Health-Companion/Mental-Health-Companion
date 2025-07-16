// ignore: unused_import
import 'package:intl/intl.dart' as intl;
import 'app_localizations.dart';

// ignore_for_file: type=lint

/// The translations for Russian (`ru`).
class AppLocalizationsRu extends AppLocalizations {
  AppLocalizationsRu([String locale = 'ru']) : super(locale);

  @override
  String get loginTitle => 'Вход для студентов';

  @override
  String get email => 'Почта';

  @override
  String get password => 'Пароль';

  @override
  String get loginButton => 'Войти';

  @override
  String get loginFailed => 'Ошибка входа';

  @override
  String get sessionTitle => 'Сессия';

  @override
  String get summaryTitle => 'Итоги сессий';

  @override
  String sessionX(Object number) {
    return 'Сессия $number';
  }
}
