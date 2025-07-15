import 'dart:convert';
import 'package:http/http.dart' as http;
import 'auth_service.dart';
import '../models/message.dart';

class ApiService {
  static const _baseUrl = 'http://localhost:8080';

  static Future<String> sendMessage(String text) async {
    final token = await AuthService.getToken();
    final response = await http.post(
      Uri.parse('$_baseUrl/message'),
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token',
      },
      body: jsonEncode({'message': text}),
    );

    if (response.statusCode == 200) {
      return jsonDecode(response.body)['reply'];
    }

    return 'Error when receiving a response from a LLM';
  }

  static Future<void> endSession(List<Message> messages) async {
    final token = await AuthService.getToken();
    await http.post(
      Uri.parse('$_baseUrl/session/end'),
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token',
      },
      body: jsonEncode({
        'messages': messages.map((m) => m.toJson()).toList(),
      }),
    );
  }

  static Future<List<String>> fetchSummaries() async {
    final token = await AuthService.getToken();
    final response = await http.get(
      Uri.parse('$_baseUrl/sessions'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      final data = jsonDecode(response.body) as List;
      return data.map((item) => item['compressed_summary'] as String).toList();
    }

    return ['Error while loading history'];
  }
}
