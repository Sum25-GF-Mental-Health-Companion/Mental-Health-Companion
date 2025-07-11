import 'package:flutter/material.dart';

class ChatInput extends StatefulWidget {
  final Function(String) onSend;

  const ChatInput({super.key, required this.onSend});

  @override
  State<ChatInput> createState() => _ChatInputState();
}

class _ChatInputState extends State<ChatInput> {
  final controller = TextEditingController();

  void _submit() {
    final text = controller.text.trim();
    if (text.isEmpty) return;
    widget.onSend(text);
    controller.clear();
  }

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Padding(
        padding: const EdgeInsets.all(8),
        child: Row(
          children: [
            Expanded(
              child: TextField(
                controller: controller,
                onSubmitted: (_) => _submit(),
                decoration: const InputDecoration(hintText: 'Type a message...'),
              ),
            ),
            IconButton(
              onPressed: _submit,
              icon: const Icon(Icons.send),
            ),
          ],
        ),
      ),
    );
  }
}
