import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../models/message.dart';

final sessionMessagesProvider = StateNotifierProvider<SessionNotifier, List<Message>>((ref) {
  return SessionNotifier();
});

class SessionNotifier extends StateNotifier<List<Message>> {
  SessionNotifier() : super([]);

  void addUserMessage(String text) {
    state = [...state, Message(sender: 'user', content: text)];
  }

  void addAiResponse(String text) {
    state = [...state, Message(sender: 'ai', content: text)];
  }

  void clearSession() {
    state = [];
  }
}
