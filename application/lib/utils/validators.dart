String? validateEmail(String? value) {
  if (value == null || value.isEmpty) return 'Email is required';
  if (!value.contains('@')) return 'Invalid email';
  return null;
}

String? validatePassword(String? value) {
  if (value == null || value.length < 6) return 'Minimum 6 characters';
  return null;
}
