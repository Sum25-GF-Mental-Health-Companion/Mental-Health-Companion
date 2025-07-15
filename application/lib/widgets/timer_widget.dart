import 'package:flutter/material.dart';

class TimerWidget extends StatelessWidget {
  final int secondsLeft;

  const TimerWidget({super.key, required this.secondsLeft});

  @override
  Widget build(BuildContext context) {
    final min = secondsLeft ~/ 60;
    final sec = (secondsLeft % 60).toString().padLeft(2, '0');
    return Text('$min:$sec', style: const TextStyle(fontSize: 20));
  }
}
