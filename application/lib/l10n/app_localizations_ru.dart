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

  @override
  String get changeLanguageTooltip => 'Сменить язык';

  @override
  String get toggleThemeTooltip => 'Сменить тему';

  @override
  String get appTitle => 'Спутник психического здоровья';

  @override
  String get login => 'Вход';

  @override
  String get signUp => 'Регистрация';

  @override
  String get registrationFailed => 'Ошибка регистрации';

  @override
  String get invalidEmail => 'Неверный адрес электронной почты';

  @override
  String get weakPassword => 'Пароль должен содержать не менее 6 символов';

  @override
  String get signUpTitle => 'Регистрация';

  @override
  String get signUpButton => 'Зарегистрироваться';

  @override
  String get typeMessage => 'Введите сообщение...';
}
