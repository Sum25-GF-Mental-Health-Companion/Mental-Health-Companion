import 'package:flutter/material.dart';
import 'package:mental_health_companion/utils/validators.dart';
import '../services/auth_service.dart';
import 'package:mental_health_companion/l10n/app_localizations.dart';

class SignUpScreen extends StatefulWidget {
  const SignUpScreen({super.key});

  @override
  State<SignUpScreen> createState() => _SignUpScreenState();
}

class _SignUpScreenState extends State<SignUpScreen> {
  final emailController = TextEditingController();
  final passwordController = TextEditingController();
  bool isLoading = false;

  Future<void> _register() async {
    final loc = AppLocalizations.of(context)!;
    String email = emailController.text.trim();
    String password = passwordController.text;

    if (validateEmail(email) != null) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text(loc.invalidEmail)),
      );
      return;
    }

    if (validatePassword(password) != null) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text(loc.weakPassword)),
      );
      return;
    }

    setState(() => isLoading = true);
    final success = await AuthService.register(email, password);
    setState(() => isLoading = false);

    if (success && context.mounted) {
      Navigator.pushReplacementNamed(context, '/session');
    } else {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text(loc.registrationFailed)),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    final loc = AppLocalizations.of(context)!;

    return Scaffold(
      appBar: AppBar(title: Text(loc.signUpTitle)),
      body: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          children: [
            TextField(
              controller: emailController,
              decoration: InputDecoration(labelText: loc.email),
            ),
            TextField(
              controller: passwordController,
              obscureText: true,
              decoration: InputDecoration(labelText: loc.password),
            ),
            const SizedBox(height: 20),
            isLoading
                ? const CircularProgressIndicator()
                : SizedBox(
                    width: 200,
                    child: ElevatedButton(
                        style: ElevatedButton.styleFrom(
                          foregroundColor: Colors.white,
                          backgroundColor: Colors.deepPurpleAccent,
                        ),
                        onPressed: _register,
                        child: Text(loc.signUpButton),
                      ),
                  ),
          ],
        ),
      ),
    );
  }
}
