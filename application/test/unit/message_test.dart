import 'package:flutter_test/flutter_test.dart';
import 'package:mental_health_companion/models/message.dart';

void main() {
  test('Message serializes correctly', () {
    final msg = Message(sender: 'user', content: 'Hello');
    final json = msg.toJson();
    expect(json['sender'], 'user');
    expect(json['content'], 'Hello');
  });
}
