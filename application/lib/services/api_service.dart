import 'dart:convert';
import 'dart:developer';
import 'package:http/http.dart' as http;
import 'auth_service.dart';
import '../models/message.dart';

class ApiService {
  static const _baseUrl = 'http://localhost:8080';

  static Future<int?> startSession() async {
    final token = await AuthService.getToken();
    final response = await http.get(
      Uri.parse('$_baseUrl/session/start'),
      headers: {
        'Authorization': 'Bearer $token',
      },
    );

    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      return data['id'];
    }

    return null;
  }

  static Future<String> sendMessage(String text) async {
    final token = await AuthService.getToken();

    final response = await http.post(
      Uri.parse('$_baseUrl/message'),
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token',
      },
      body: jsonEncode({'text': text}),
    );

    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);

      final reply = data['response'];

      if (reply is String) return reply;
      return 'Unexpected format';
    }

    return 'AI error';
  }

  static Future<void> endSession(int sessionId, List<Message> messages) async {
    final token = await AuthService.getToken();

    final body = jsonEncode({
      'session_id': sessionId,
      'messages': messages.map((m) => {
        'sender': m.sender,
        'content': m.content,
      }).toList(),
    });

    final response = await http.post(
      Uri.parse('$_baseUrl/session/end'),
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token',
      },
      body: body,
    );

    if (response.statusCode != 200) {
      print('EndSession failed: ${response.statusCode} ${response.body}');
      throw Exception('End session failed');
    }
  }

  static Future<List<String>> fetchSummaries() async {
    final token = await AuthService.getToken();
    final response = await http.get(
      Uri.parse('$_baseUrl/sessions'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      final data = jsonDecode(response.body)['sessions'] as List;
      return data.map((e) => e['compressed_summary'] as String).toList();
    }

    return ['Error loading history'];
  }
}
