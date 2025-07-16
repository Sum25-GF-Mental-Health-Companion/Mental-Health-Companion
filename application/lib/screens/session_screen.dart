import 'dart:async';
import 'package:flutter/material.dart';
import '../models/message.dart';
import '../services/api_service.dart';
import '../widgets/chat_input.dart';
import '../widgets/message_bubble.dart';
import '../utils/constants.dart';
import 'package:mental_health_companion/l10n/app_localizations.dart';

class SessionScreen extends StatefulWidget {
  final bool testMode;

  const SessionScreen({super.key, this.testMode = false});

  @override
  State<SessionScreen> createState() => _SessionScreenState();
}

class _SessionScreenState extends State<SessionScreen> {
  final List<Message> _messages = [];
  late Timer _timer;
  int _secondsLeft = sessionDuration.inSeconds;
  bool sessionEnded = false;
  int? _sessionId;

  @override
  void initState() {
    super.initState();
    if (widget.testMode) {
      _sessionId = 1;
      _startTimer();
    } else {
      _initSession();
    }
  }

  Future<void> _initSession() async {
    final sessionId = await ApiService.startSession();
    if (sessionId != null) {
      setState(() {
        _sessionId = sessionId;
      });
      _startTimer();
    }
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
    if (_sessionId != null) {
      await ApiService.endSession(_sessionId!, _messages);
    }
    if (context.mounted) {
      Navigator.pushReplacementNamed(context, '/summary');
    }
  }

  Future<void> _sendMessage(String text) async {
    if (_sessionId == null) return;

    setState(() {
      _messages
          .add(Message(sessionId: _sessionId, sender: 'user', content: text));
    });

    final response = await ApiService.sendMessage(text);

    setState(() {
      _messages
          .add(Message(sessionId: _sessionId, sender: 'ai', content: response));
    });
  }

  @override
  void dispose() {
    _timer.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final loc = AppLocalizations.of(context)!;
    final min = _secondsLeft ~/ 60;
    final sec = (_secondsLeft % 60).toString().padLeft(2, '0');

    return Scaffold(
      appBar: AppBar(title: Text('${loc.sessionTitle} — $min:$sec')),
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
          if (!sessionEnded && _sessionId != null)
            ChatInput(onSend: _sendMessage),
        ],
      ),
    );
  }
}
