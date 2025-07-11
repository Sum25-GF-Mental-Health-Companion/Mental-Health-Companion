import 'dart:async';
import 'package:flutter/material.dart';
import '../models/message.dart';
import '../services/api_service.dart';
import '../widgets/chat_input.dart';
import '../widgets/message_bubble.dart';
import '../utils/constants.dart';

class SessionScreen extends StatefulWidget {
  const SessionScreen({super.key});

  @override
  State<SessionScreen> createState() => _SessionScreenState();
}

class _SessionScreenState extends State<SessionScreen> {
  final List<Message> _messages = [];
  late Timer _timer;
  int _secondsLeft = sessionDuration.inSeconds;
  bool sessionEnded = false;

  @override
  void initState() {
    super.initState();
    _startTimer();
  }

  void _startTimer() {
    _timer = Timer.periodic(const Duration(seconds: 1), (timer) {
      if (_secondsLeft == 0) {
        _endSession();
      } else {
        setState(() => _secondsLeft--);
      }
    });
  }

  Future<void> _endSession() async {
    _timer.cancel();
    setState(() => sessionEnded = true);
    await ApiService.endSession(_messages);
    if (context.mounted) {
      Navigator.pushReplacementNamed(context, '/summary');
    }
  }

  Future<void> _sendMessage(String text) async {
    setState(() {
      _messages.add(Message(sender: 'user', content: text));
    });

    final response = await ApiService.sendMessage(text);
    setState(() {
      _messages.add(Message(sender: 'ai', content: response));
    });
  }

  @override
  void dispose() {
    _timer.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final min = _secondsLeft ~/ 60;
    final sec = (_secondsLeft % 60).toString().padLeft(2, '0');

    return Scaffold(
      appBar: AppBar(title: Text('Session — $min:$sec')),
      body: Column(
        children: [
          Expanded(
            child: ListView.builder(
              padding: const EdgeInsets.all(8),
              itemCount: _messages.length,
              itemBuilder: (context, index) {
                final msg = _messages[index];
                return MessageBubble(
                  text: msg.content,
                  isUser: msg.sender == 'user',
                );
              },
            ),
          ),
          if (!sessionEnded)
            ChatInput(onSend: _sendMessage),
        ],
      ),
    );
  }
}
